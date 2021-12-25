package services

import (
	"encoding/base64"
	"fmt"
	"sync"
	"time"
)

type UrlCacheStruct struct {
	RWMutex   sync.RWMutex
	Urls      map[string]int
	Ttl       int
	Threshold int
}

var UrlCache UrlCacheStruct
var UrlTtlTracker map[int64][]string

func AddEntry(url string) bool {
	urlHash := base64.StdEncoding.EncodeToString([]byte(url))
	isBlocked := false
	dt := time.Now()
	if count, ok := UrlCache.Urls[urlHash]; ok {
		if count < UrlCache.Threshold {
			UrlCache.RWMutex.Lock()
			UrlCache.Urls[urlHash] = count + 1
			UrlCache.RWMutex.Unlock()
		} else {
			isBlocked = true
		}
	} else {
		fmt.Println("new entry:  ", url)
		urlTtl := time.Now().Add(time.Duration(UrlCache.Ttl) * time.Millisecond).UnixMilli()

		UrlCache.RWMutex.Lock()
		UrlCache.Urls[urlHash] = 1
		UrlCache.RWMutex.Unlock()

		//put inside tracker
		UrlTtlTracker[urlTtl] = append(UrlTtlTracker[urlTtl], urlHash)
	}
	fmt.Println(dt.Format("15:04:05"), "URL: ", url, "is reported, ", "count:=", UrlCache.Urls[urlHash], ", blocked? ", isBlocked)
	return isBlocked
}
