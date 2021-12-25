package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"word-letters-shuffle/word"

	"github.com/gofiber/fiber/v2"
)

// TODO: wystawić jako serwis na Heroku

type InputTextBody struct {
	Text string `json:"text"`
}

func main() {

	port := os.Getenv("PORT")

	fmt.Println("Port:", port)

	if port == "" {
		port = "5000"
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		inputTextBody := new(InputTextBody)

		if err := c.BodyParser(inputTextBody); err != nil {
			return err
		}

		inputText := inputTextBody.Text

		rand.Seed(time.Now().UnixNano()) // to ensure unique shuffle each time it runs

		inputTextSplitted := strings.Fields(inputText)
		outputTextSplitted := make([]string, 0, 100)

		for i := range inputTextSplitted {
			shuffledWord := word.ShuffleWordInnerials(inputTextSplitted[i])
			outputTextSplitted = append(outputTextSplitted, string(shuffledWord))
		}

		outputText := strings.Join(outputTextSplitted, " ")

		return c.JSON(fiber.Map{
			"modifiedText": outputText,
		})

		// return c.SendString(outputText)
	})

	app.Listen(":" + port)

}
