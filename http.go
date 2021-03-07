package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

// HTTPServer return http server
func HTTPServer(
	baseURL string,
	username string,
	password string,
) (application *fiber.App, router fiber.Router) {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Prefork:               true,
		UnescapePath:          true,
		CaseSensitive:         true,
		StrictRouting:         true,
		BodyLimit:             1 * 1024,
		ServerHeader:          "aasaam-whois-json",

		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			ctx.Status(code).SendString(err.Error())

			return nil
		},
	})

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			username: password,
		},
	}))

	api := app.Group(baseURL)

	api.Get("/whois/:domain", HTTPDomainWhois)
	api.Get("/validate/:domain", HTTPDomainValidation)

	return app, api
}
