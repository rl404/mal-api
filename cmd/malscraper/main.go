package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/rl404/go-malscraper"
	_ "github.com/rl404/mal-api/api"
	"github.com/rl404/mal-api/internal/cacher"
	"github.com/rl404/mal-api/internal/config"
	"github.com/rl404/mal-api/internal/pkg/http"
	"github.com/rl404/mal-api/internal/pkg/middleware"
	"github.com/rl404/mal-api/internal/router/api"
	"github.com/rl404/mal-api/internal/router/ping"
	"github.com/rl404/mal-api/internal/router/swagger"
	"github.com/rl404/mal-plugin/log/elasticsearch"
	"github.com/rl404/mal-plugin/log/mallogger"
	"github.com/rs/cors"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

// @title Malscraper API
// @description Scraping/parsing MyAnimeList website to a useful and easy-to-use data.
// @contact.name Axel
// @contact.url https://github.com/rl404
// @contact.email axel.rl.404@gmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @BasePath /
// @schemes http https
func main() {
	// Get config.
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	// Init logger.
	l := mallogger.New(cfg.Log.Level, cfg.Log.Color)
	l.Info("logger initialized")

	// Init cache.
	c, err := cacher.New(l, cfg.Cache.Dialect, cfg.Cache.Address, cfg.Cache.Password, time.Duration(cfg.Cache.Time)*time.Second)
	if err != nil {
		l.Fatal(err.Error())
	}
	l.Info("cache initialized")

	// Init malscraper.
	mal, err := malscraper.New(malscraper.Config{
		Cacher:        c,
		CleanImageURL: cfg.Clean.Image,
		CleanVideoURL: cfg.Clean.Video,
		Logger:        l,
	})
	if err != nil {
		l.Fatal(err.Error())
	}
	l.Info("malscraper initialized")

	// Init web server.
	server := http.New(http.Config{
		Port:            cfg.Web.Port,
		ReadTimeout:     cfg.Web.ReadTimeout,
		WriteTimeout:    cfg.Web.WriteTimeout,
		GracefulTimeout: cfg.Web.GracefulTimeout,
	})
	r := server.Router()

	// Init web router middleware.
	r.Use(cors.AllowAll().Handler)
	r.Use(middleware.RealIP)
	if cfg.ES.Address != "" {
		// Init elasticsearch.
		es, err := elasticsearch.New(strings.Split(cfg.ES.Address, ","), cfg.ES.User, cfg.ES.Password)
		if err != nil {
			l.Fatal(err.Error())
		}
		r.Use(middleware.Logger(l, es))
		l.Info("elasticsearch initialized")
	}
	r.Use(middleware.Recoverer)
	l.Info("middleware initialized")

	// Register ping route.
	ping.New().Register(r)
	l.Info("base routes initialized")

	// Register swagger route.
	swagger.New().Register(r)
	l.Info("swagger routes initialized")

	// Register api routes.
	api.New(mal).Register(r)
	l.Info("api routes initialized")

	// Run web server.
	serverChan := server.Run()
	l.Info("web server initialized")
	l.Info("server listening at %s", cfg.Web.Port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case err := <-serverChan:
		if err != nil {
			l.Fatal(err.Error())
		}
	case <-sigChan:
	}

	if err = server.Close(); err != nil {
		l.Error(err.Error())
	} else {
		l.Info("web server stopped")
	}

	if err = c.Close(); err != nil {
		l.Error(err.Error())
	} else {
		l.Info("cache stopped")
	}

	l.Info("server has been running for %s", time.Since(startTime).Truncate(time.Millisecond))
}
