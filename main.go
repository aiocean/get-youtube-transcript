package main

import (
	"github.com/aiocean/get-youtube-transcript/pkg/youtube"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/utils"
	tokenizer "github.com/samber/go-gpt-3-encoder"
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
			return utils.CopyString(c.Params("id"))
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

func getTranscriptHandler(c *fiber.Ctx) error {
	videoID := c.Params("id")

	transcript, err := youtube.GetTranscript(videoID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if len(transcript.Segments) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No transcript found",
		})
	}

	format := c.Query("format")
	if format == "" || format == "json" {
		return c.JSON(transcript)
	}

	includeTime := c.Query("include-time", "false") == "true"

	var plainText string
	for _, segment := range transcript.Segments {
		if includeTime {
			plainText += "[" + segment.Time + "]"
		}
		plainText += segment.Text + " "
	}

	chunkSize := c.QueryInt("chunk-size", 0)
	if chunkSize == 0 {
		return c.JSON([]string{plainText})
	}

	var chunks []string

	encoder, err := getTokenEncoder()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	encodedText, err := encoder.Encode(plainText)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	for i := 0; i < len(encodedText); i += chunkSize {
		end := i + chunkSize
		if end > len(encodedText) {
			end = len(encodedText)
		}

		chunks = append(chunks, encoder.Decode(encodedText[i:end]))
	}

	return c.JSON(chunks)
}

var cacheTokenEncoder *tokenizer.Encoder

func getTokenEncoder() (*tokenizer.Encoder, error) {
	if cacheTokenEncoder == nil {
		var err error
		cacheTokenEncoder, err = tokenizer.NewEncoder()
		if err != nil {
			return nil, err
		}
	}

	return cacheTokenEncoder, nil
}

func homeHandler(c *fiber.Ctx) error {
	return c.Redirect("https://github.com/aiocean/get-youtube-transcript")
}
