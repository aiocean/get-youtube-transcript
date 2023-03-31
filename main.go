package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/utils"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	cacheModule := cache.New(cache.Config{
		ExpirationGenerator: func(c *fiber.Ctx, cfg *cache.Config) time.Duration {
			newCacheTime, _ := strconv.Atoi(c.GetRespHeader("Cache-Time", "600"))
			return time.Second * time.Duration(newCacheTime)
		},
		KeyGenerator: func(c *fiber.Ctx) string {
			return utils.CopyString(c.Path())
		}})

	app.Use(cacheModule)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Get("/transcripts/:id", getTranscriptHandler)

	log.Println("listening on", port)
	log.Fatal(app.Listen(":" + port))
}

func getTranscriptHandler(c *fiber.Ctx) error {
	videoID := c.Params("id")
	c.Response().Header.Add("Cache-Time", "6000")

	return c.JSON(fiber.Map{
		"videoID": videoID,
	})
}
