package main

import (
	"github.com/gofiber/basicauth"
	"github.com/gofiber/fiber"
)

func domainWhoIS(c *fiber.Ctx) {
	domainType, e := DomainValidation(c.Params("domain"))
	if e != nil || domainType.TLDASCII == "" {
		err := fiber.NewError(400, e.Error())
		c.Next(err)
		return
	}

	result, e := DomainParse(domainType)
	if e != nil {
		err := fiber.NewError(500, e.Error())
		c.Next(err)
		return
	}
	if err := c.JSON(result); err != nil {
		err := fiber.NewError(500, "Internal Server Error")
		c.Next(err)
		return
	}
}

func domainValidation(c *fiber.Ctx) {
	domainType, e := DomainValidation(c.Params("domain"))
	if e != nil || domainType.TLDASCII == "" {
		err := fiber.NewError(400, e.Error())
		c.Next(err)
		return
	}

	if err := c.JSON(domainType); err != nil {
		err := fiber.NewError(500, "Internal Server Error")
		c.Next(err)
		return
	}
}

// HTTPServer return http server
func HTTPServer(baseURL string, username string, password string, set404 bool) (application *fiber.App, router fiber.Router) {
	app := fiber.New()

	app.Settings.DisableStartupMessage = true
	app.Settings.Prefork = false
	app.Settings.UnescapePath = true
	app.Settings.CaseSensitive = true
	app.Settings.StrictRouting = false
	app.Settings.BodyLimit = 1 * 1024
	app.Settings.ServerHeader = "aasaam"

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			username: password,
		},
	}))

	app.Settings.ErrorHandler = func(ctx *fiber.Ctx, err error) {
		code := fiber.StatusInternalServerError

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		// Return HTTP response
		ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
		ctx.Status(code).SendString(err.Error())
	}

	api := app.Group(baseURL)

	api.Get("/whois/:domain", domainWhoIS)
	api.Get("/validate/:domain", domainValidation)

	if set404 {
		app.Use(func(c *fiber.Ctx) {
			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			c.Status(404).SendString("Not found: " + c.OriginalURL())
		})
	}

	return app, api
}
