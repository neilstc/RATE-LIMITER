package services

import (
	"fmt"
	"sync"
	"time"
)

var UrlCache UrlCacheStruct

type UrlCacheStruct struct {
	RWMutex   sync.RWMutex
	Urls      map[string]int
	Ttl       int
	Threshold int
}

func AddEntry(url string) bool {

	isBlocked := false
	UrlCache.RWMutex.Lock()
	dt := time.Now()
	if count, ok := UrlCache.Urls[url]; ok {
		if count < UrlCache.Threshold {
			UrlCache.Urls[url] = count + 1
		} else {
			isBlocked = true
		}
	} else {
		UrlCache.Urls[url] = 1
	}
	UrlCache.RWMutex.Unlock()
	fmt.Println(dt.Format("00:00:00"), "URL: ", url, "is reported, ", "count:=", UrlCache.Urls[url], ", blocked? ", isBlocked)

	return isBlocked
}
