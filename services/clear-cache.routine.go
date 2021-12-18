package services

import (
	"encoding/json"
	"fmt"
	"time"
)

func CleanUrlCacheRoutine() {

	for {
		time.Sleep(time.Duration(UrlCache.Ttl) * time.Millisecond)
		formattedMap, _ := json.MarshalIndent(UrlCache.Urls, "", " ")
		fmt.Println("map before clearing c", string(formattedMap))
		UrlCache.Urls = make(map[string]int)
	}
}
