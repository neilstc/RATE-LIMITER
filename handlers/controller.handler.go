package handlers

import (
	"rate-limiter/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UrlDto struct {
	Url string `validate:"required,url,min=7"`
}

func ResponseHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	urlDto := new(UrlDto)

	// content type validation
	if err := c.BodyParser(urlDto); err != nil {
		return c.Status(fiber.StatusUnsupportedMediaType).JSON(fiber.Map{"message": "Content-type should be application/json! ", "data": err})
	}

	// url validation
	if err := ValidateStruct(*urlDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid url! "})
	}

	url := urlDto.Url

	isBlocked := services.AddEntry(url)

	var statusCode int
	if isBlocked {
		statusCode = fiber.StatusServiceUnavailable
	} else {
		statusCode = fiber.StatusOK
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"block": isBlocked,
	})
}

func ValidateStruct(urlDto UrlDto) error {

	validate := validator.New()
	err := validate.Struct(urlDto)
	return err

}
