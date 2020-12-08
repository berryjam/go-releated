package main

import (
	"fmt"
	"github.com/allegro/bigcache"
	"time"
)

func main() {
	config := bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 1024,

		// time after which entry can be evicted
		LifeWindow: 5 * time.Second,

		// Interval between removing expired entries (clean up).
		// If set to <= 0 then no action is performed.
		// Setting to < 1 second is counterproductive — bigcache has a one second resolution.
		CleanWindow: 500 * time.Millisecond,

		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,

		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,

		// prints information about additional memory allocation
		Verbose: true,

		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 8192,

		// callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A bitmask representing the reason will be returned.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		OnRemove: nil,

		// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A constant representing the reason will be passed through.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		// Ignored if OnRemove is specified.
		OnRemoveWithReason: nil,
	}
	cache, _ := bigcache.NewBigCache(config)

	cache.Set("my-unique-key", nil)

	entry, err := cache.Get("my-unique-key")
	if err != nil && err== bigcache.ErrEntryNotFound {
		fmt.Println("entry not found")
	} else {
		fmt.Println("entry val" + string(entry))
	}

	time.Sleep(6 * time.Second)
	anEntry, _ := cache.Get("my-unique-key")
	if anEntry == nil {
		fmt.Println("nil")
	} else if len(anEntry) == 0 {
		fmt.Println("empty")
	} else {
		fmt.Println("val:" + string(anEntry))
	}

}
