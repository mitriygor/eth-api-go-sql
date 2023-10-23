package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

// LoggingMiddleware is a Fiber middleware function that logs information about each HTTP request.
func LoggingMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()

	log.Printf("%v - %v - %v - %v", c.IP(), c.Method(), c.Path(), time.Since(start))

	return err
}
