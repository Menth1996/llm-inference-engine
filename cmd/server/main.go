package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/Menth1996/llm-inference-engine/pkg/inference"
	"github.com/Menth1996/llm-inference-engine/pkg/model"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func main() {
	app := fiber.New()

	modelPath := viper.GetString("model.path")
	modelType := viper.GetString("model.type")

	llmModel, err := model.LoadModel(modelPath, modelType)
	if err != nil {
		log.Fatalf("Failed to load model: %v", err)
	}

	inferenceEngine := inference.NewInferenceEngine(llmModel)

	app.Post("/predict", func(c *fiber.Ctx) error {
		type Request struct {
			Prompt string `json:"prompt"`
		}
		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		result, err := inferenceEngine.Predict(req.Prompt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"prediction": result,
		})
	})

	log.Fatal(app.Listen(":" + viper.GetString("server.port")))
}
