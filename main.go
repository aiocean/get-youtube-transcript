package main

import (
	"github.com/aiocean/get-youtube-transcript/pkg/youtube"
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

	app.Get("/", homeHandler)
	app.Get("/transcripts/:id", getTranscriptHandler)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}

var transcriptCache = make(map[string]*youtube.Transcript)

func getTranscriptHandler(c *fiber.Ctx) error {
	videoID := c.Params("id")

	if transcript, ok := transcriptCache[videoID]; ok {
		return c.JSON(transcript)
	}

	transcript, err := youtube.GetTranscript(videoID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	transcriptCache[videoID] = transcript
	return c.JSON(transcript)
}

func homeHandler(c *fiber.Ctx) error {
	return c.Redirect("https://github.com/aiocean/get-youtube-transcript")
}
