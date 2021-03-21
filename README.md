# mal-api

REST API to scrap and parse [MyAnimeList](https://myanimelist.net/) to get information.

Powered by my [go-malscraper](https://github.com/rl404/go-malscraper).

## Features

* Get anime information (details, characters, episodes, pictures, etc)
* Get manga information (details, characters, pictures, recommendations, etc)
* Get character information (details, pictures, etc)
* Get people information (details, pictures, etc)
* Get list of all anime/manga's genres
* Get list of all anime/manga's producers/studios/licensors/magazines/serializations
* Get anime/manga's recommendations
* Get anime/manga's reviews
* Search anime, manga, character and people
* Get seasonal anime list
* Get anime, manga, character and people top list
* Get user information (profile, friends, histories, recommendations, reviews, etc)
* Get news list and details
* Get featured article list and details
* Get club list and details
* Caching (in-memory, [redis](https://redis.io/), or [memcache](https://memcached.org/))
* Logging ([elasticsearch](https://www.elastic.co/))
* [Swagger](https://mal-rest-api.herokuapp.com/swagger/index.html)
* [Docker](https://www.docker.com/)

_More will be coming soon..._

## Installation

1. Clone the repo.
```
git clone github.com/rl404/mal-api
```
2. Update the `.env`.
3. Build and run
```
# With Go
make

# With Docker
make docker

# To stop docker container
make docker-stop
```

Or just pull docker image.
```
docker pull rl404/mal-api:latest
```

## Config

> All env are optional. Use what you have and what you want.

Env | Default | Description
--- | :---: | ---
`MAL_WEB_PORT` | `8005` | HTTP port
`MAL_WEB_READ_TIMEOUT` | `5` | HTTP read timeout (in seconds)
`MAL_WEB_WRITE_TIMEOUT` | `5` | HTTP write timeout (in seconds)
`MAL_WEB_GRACEFUL_TIMEOUT` | `10` | HTTP server shutdown timeout (in seconds)
`MAL_CLEAN_IMAGE` | `true` | Cleaning MyAnimelist image URL
`MAL_CLEAN_VIDEO` | `true` | Cleaning MyAnimelist image URL
`MAL_CACHE_DIALECT` | `inmemory` | Cache type (`nocache`, `inmemory`, `redis`, `memcache`)
`MAL_CACHE_ADDRESS` |  | Cache address
`MAL_CACHE_PASSWORD` |  | Cache password
`MAL_CACHE_TIME` | `86400` | Cache time (in seconds)
`MAL_LOG_LEVEL` | `4` | Log all
`MAL_LOG_COLOR` | `true` | Log color
`MAL_ES_ADDRESS` |  | Elasticsearch address
`MAL_ES_USER` |  | Elasticsearch user
`MAL_ES_PASSWORD` |  | Elasticsearch password

### Log Level

Level | Trace | Debug | Info | Warn | Error | Fatal
:---: | :---: | :---: | :---: | :---: | :---: | :---: |
`0` | :x: | :x: | :x: | :x: | :x: | :x:
`1` | :x: | :x: | :x: | :x: | :heavy_check_mark: | :heavy_check_mark:
`2` | :x: | :x: | :heavy_check_mark: | :x: | :heavy_check_mark: | :heavy_check_mark:
`3` | :x: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:
`4` | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:

## Disclamer

_mal-api_ is meant for educational purpose and personal usage only. Although there is no limit in using the API, do remember that every scraper method is accessing MyAnimeList page so use it responsibly according to MyAnimeList's [Terms Of Service](https://myanimelist.net/about/terms_of_use).

All data (including anime, manga, people, etc) belong to their respective copyrights owners. mal-api does not have any affiliation with content providers.

## License

MIT License

Copyright (c) 2021 Axel