package routes

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/safwatimran/url-shortener-service/api/database"
)

func ResolveURL(c *fiber.Ctx) error {

	url := c.Params("url")

	r := database.GetClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short url not found.",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "server error.",
		})
	}

	rateIncrement := database.GetClient(1)
	defer rateIncrement.Close()

	_ = rateIncrement.Incr(database.Ctx, "counter")

	return c.Redirect(value, http.StatusMovedPermanently)
}
