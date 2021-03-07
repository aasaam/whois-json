package main

import "github.com/gofiber/fiber/v2"

// HTTPDomainValidation is http endpoint for validate domain
func HTTPDomainValidation(c *fiber.Ctx) error {
	domain, e := NewDomain(c.Params("domain"))
	if e != nil {
		return fiber.ErrBadRequest
	}

	if err := c.JSON(domain); err != nil {
		return fiber.ErrInternalServerError
	}

	return nil
}
