package main

import (
	"encoding/json"
	"fmt"
	"os"

	"testing"
)

func skipCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}

func TestValidDomainParse(t *testing.T) {
	skipCI(t)
	domain, _ := NewDomain("nic.ir")
	whois, e := DomainWhois(domain)
	whoisJSON, _ := json.Marshal(whois)
	fmt.Println(string(whoisJSON))
	if e != nil {
		t.Error(e)
	}
}

func TestNewDomain(t *testing.T) {
	domains := []string{
		"amazon.co.uk",
		"google.com",
		"www.google.com",
		"بخر.ایران",
		"ساب.بخر.ایران",
		"sample.gov.ir",
		"www.sub.sample.gov.ir",
	}
	for _, domain := range domains {
		_, e := NewDomain(domain)
		if e != nil {
			t.Error(e)
		}
	}

	invalidDomains := []string{
		"AAAAAAAAAAA",
		"محمد.آیران",
		"😀",
		"",
		"東京\uFF0EFjp",
		"1",
	}

	for _, domain := range invalidDomains {
		_, e := NewDomain(domain)
		if e == nil {
			t.Errorf("Domain is invalid so we need an error")
		}
	}
}
