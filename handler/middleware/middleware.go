package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	appCtx *fiber.App
	// another stuff , may be needed by middleware
}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS() fiber.Handler {
	crs := os.Getenv("SIPD_CORS_WHITELISTS")

	if crs == "*" {
		return cors.New(cors.Config{
			AllowOrigins:  "*",
			AllowHeaders:  "Content-Type, Accept, Authorization",
			AllowMethods:  "GET, HEAD, PUT, PATCH, POST, DELETE",
			ExposeHeaders: "*", //"X-Pagination-Current-Page,X-Pagination-Next-Page,X-Pagination-Page-Count,X-Pagination-Page-Size,X-Pagination-Total-Count"
		})
	}

	return cors.New(cors.Config{
		AllowOrigins:     crs,
		AllowCredentials: true,
		AllowHeaders:     "Content-Type, Accept, Authorization",
		AllowMethods:     "GET, HEAD, PUT, PATCH, POST, DELETE",
		ExposeHeaders:    "*", //"X-Pagination-Current-Page,X-Pagination-Next-Page,X-Pagination-Page-Count,X-Pagination-Page-Size,X-Pagination-Total-Count"
	})
}

// LOGGER simple logger.
func (m *GoMiddleware) LOGGER() fiber.Handler {
	return logger.New()
}

// JWT jwt.
func (m *GoMiddleware) JWT() fiber.Handler {
	// Create config for JWT authentication middleware.
	config := jwtware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey:   "jwt", // used in private routes
		ErrorHandler: jwtError,
	}

	return jwtware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 400 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}

func (m *GoMiddleware) RateLimiter() fiber.Handler {
	limiterCfg := limiter.Config{
		Max:        10,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			authKey := c.Get("Authorization")
			if authKey == "" {
				return c.IP() + c.Get("User-Agent")
			} else {
				return authKey
			}
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
	}

	return limiter.New(limiterCfg)
}

// InitMiddleware initialize the middleware
func InitMiddleware(ctx *fiber.App) *GoMiddleware {
	return &GoMiddleware{appCtx: ctx}
}
