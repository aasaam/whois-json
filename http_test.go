package main

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber"
)

func TestHTTPBasicAuthFailed(t *testing.T) {
	app, _ := HTTPServer("/", "user", "pass", true)

	req := httptest.NewRequest("GET", "/", nil)

	resp, _ := app.Test(req)
	if resp.StatusCode != 401 {
		t.Errorf("Status must be 403")
	}
}

func TestHTTPSuccessTest(t *testing.T) {
	app, _ := HTTPServer("/", "user", "pass", true)

	req := httptest.NewRequest("GET", "/not-found", nil)
	req.SetBasicAuth("user", "pass")
	resp, _ := app.Test(req)

	if resp.StatusCode != 404 {
		t.Errorf("Status must be 404")
	}

	req = httptest.NewRequest("GET", "/whois/localhost", nil)
	req.SetBasicAuth("user", "pass")
	resp, _ = app.Test(req)

	if resp.StatusCode != 400 {
		t.Errorf("Status must be 400")
	}

	req = httptest.NewRequest("GET", "/whois/nic.ir", nil)
	req.SetBasicAuth("user", "pass")
	resp, _ = app.Test(req)

	if resp.StatusCode != 200 {
		t.Errorf("Status must be 200")
	}

	req = httptest.NewRequest("GET", "/validate/nic.ir", nil)
	req.SetBasicAuth("user", "pass")
	resp, _ = app.Test(req)

	if resp.StatusCode != 200 {
		t.Errorf("Status must be 200")
	}

	req = httptest.NewRequest("GET", "/validate/nic.not-exist", nil)
	req.SetBasicAuth("user", "pass")
	resp, _ = app.Test(req)

	if resp.StatusCode != 400 {
		t.Errorf("Status must be 400")
	}
}

func TestHTTPBasicAuthSuccess(t *testing.T) {
	app, api := HTTPServer("/base", "user", "pass", false)

	api.Get("/ok", func(c *fiber.Ctx) {
		c.Send("OK")
	})

	api.Get("/error", func(c *fiber.Ctx) {
		err := fiber.NewError(500, "This is an error")
		c.Next(err)
	})

	req := httptest.NewRequest("GET", "/base/ok", nil)
	req.SetBasicAuth("user", "pass")
	resp, _ := app.Test(req)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	body := buf.String()
	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		t.Errorf("Status must be 200")
	}

	if body != "OK" {
		fmt.Println(resp.StatusCode)
		t.Errorf("Status must be 200")
	}

	req = httptest.NewRequest("GET", "/base/error", nil)
	req.SetBasicAuth("user", "pass")
	resp, _ = app.Test(req)
	if resp.StatusCode != 500 {
		fmt.Println(resp.StatusCode)
		t.Errorf("Status must be 500")
	}

	return
}
