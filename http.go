package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func domainWhoIS(c *fiber.Ctx) error {
	domainType, e := DomainValidation(c.Params("domain"))

	if e != nil || domainType.TLDASCII == "" {
		return fiber.ErrBadRequest
	}

	result, e := DomainParse(domainType)
	if e != nil {
		return fiber.ErrInternalServerError
	}

	if err := c.JSON(result); err != nil {
		return fiber.ErrInternalServerError
	}

	return nil
}

func domainValidation(c *fiber.Ctx) error {
	domainType, e := DomainValidation(c.Params("domain"))
	if e != nil || domainType.TLDASCII == "" {
		return fiber.ErrBadRequest
	}

	if err := c.JSON(domainType); err != nil {
		return fiber.ErrInternalServerError
	}

	return nil
}

// HTTPServer return http server
func HTTPServer(baseURL string, username string, password string, set404 bool) (application *fiber.App, router fiber.Router) {
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

	api.Get("/whois/:domain", domainWhoIS)
	api.Get("/validate/:domain", domainValidation)

	if set404 {
		app.Use(func(c *fiber.Ctx) error {
			return fiber.ErrNotFound
		})
	}

	return app, api
}
