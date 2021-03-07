package main

import "github.com/gofiber/fiber/v2"

// HTTPDomainWhois is http endpoint for whois domain
func HTTPDomainWhois(c *fiber.Ctx) error {
	domain, e := NewDomain(c.Params("domain"))

	if e != nil {
		return fiber.ErrBadRequest
	}

	result, e := DomainWhois(domain)
	if e != nil {
		return fiber.ErrInternalServerError
	}

	if err := c.JSON(result); err != nil {
		return fiber.ErrInternalServerError
	}

	return nil
}
