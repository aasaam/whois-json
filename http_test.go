package main

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"strings"
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

	req1 := httptest.NewRequest("GET", "/not-found", nil)
	req1.SetBasicAuth("user", "pass")
	resp1, _ := app.Test(req1)

	if resp1.StatusCode != 404 {
		t.Errorf("Status must be 404")
	}

	req2 := httptest.NewRequest("GET", "/whois/localhost", nil)
	req2.SetBasicAuth("user", "pass")
	resp2, _ := app.Test(req2)

	if resp2.StatusCode != 400 {
		t.Errorf("Status must be 400")
	}

	req5 := httptest.NewRequest("GET", "/validate/nic.not-exist", nil)
	req5.SetBasicAuth("user", "pass")
	resp5, _ := app.Test(req5)

	if resp5.StatusCode != 400 {
		t.Errorf("Status must be 400")
	}
}

func TestHTTPSuccessTest2(t *testing.T) {
	app, _ := HTTPServer("/", "user", "pass", true)

	domains := []string{"google.com", "github.net", "mozilla.org", "nic.ir", "ایرنیک.ایران"}

	for _, domain := range domains {
		req3 := httptest.NewRequest("GET", "/whois/"+domain, nil)
		req3.SetBasicAuth("user", "pass")
		resp3, err3 := app.Test(req3)
		if err3 != nil {
			if !strings.Contains(err3.Error(), "timeout") {
				t.Error(err3)
			}
		} else {
			if resp3.StatusCode != 200 {
				t.Errorf("Status must be 200")
			}
		}
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
