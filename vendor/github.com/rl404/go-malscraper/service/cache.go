package service

// Cacher is caching interface required for malscraper
// caching system. If you use custom caching, try to
// implement this interface to your cacher.
type Cacher interface {
	// Get data from cache. The returned value will be
	// assigned to param `data`. Param `data` should
	// be a pointer just like when using json.Unmarshal.
	Get(key string, data interface{}) error
	// Save data to cache. Set and Get should be using
	// the same encoding method for example, json.Marshal
	// for Set and json.Unmarshal for Get.
	Set(key string, data interface{}) error
	// Delete data from cache.
	Delete(key string) error
	// Close cache connection.
	Close() error
}
