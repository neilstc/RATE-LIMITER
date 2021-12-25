package services

import (
	"encoding/base64"
	"fmt"
	"time"
)

func CleanUrlCacheRoutine() {
	for {
		go cleanUrlCache()
		time.Sleep(300 * time.Microsecond)
	}
}

func cleanUrlCache() {
	cleaningTime := time.Now().UnixMilli()
	UrlCache.RWMutex.RLock()
	defer UrlCache.RWMutex.RUnlock()

	// iterate only on relevant timestamps
	if urlArray, ok := UrlTtlTracker[cleaningTime]; ok {
		for _, url := range urlArray {
			delete(UrlCache.Urls, url)
			// clean form ttl tracker
			delete(UrlTtlTracker, cleaningTime)

			decoded, _ := base64.StdEncoding.DecodeString(url)
			fmt.Printf("%q finished thier session!\n", decoded)
		}
	}

}
