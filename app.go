package main

import (
	"flag"
	"fmt"
	"os"
	"rate-limiter/handlers"
	"rate-limiter/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {

	port := appInit()

	go services.CleanUrlCacheRoutine()

	app := fiber.New()

	app.Post("/response", handlers.ResponseHandler)

	err := app.Listen(":" + port)

	if err != nil {
		panic(err)
	}
}

func appInit() string {
	thresholdPtr := flag.Int("threshold", 10, "url count threshold")
	ttlPtr := flag.Int("ttl", 60000, "server's ttl")
	portPtr := flag.String("port", "8081", "server's port")
	flag.Parse()

	checkport, _ := strconv.Atoi(*portPtr)
	if *thresholdPtr < 0 || *ttlPtr < 0 || checkport <= 1023 || checkport >= 49151 {
		fmt.Fprintf(os.Stderr, "error: %v\n", "illegal command line arguments!")
		os.Exit(1)
	}

	services.UrlCache.Threshold = *thresholdPtr
	services.UrlCache.Ttl = *ttlPtr
	services.UrlCache.Urls = make(map[string]int)
	port := *portPtr

	fmt.Println("app initialized")

	return port
}
