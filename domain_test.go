package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDomainValidationValid(t *testing.T) {
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
		_, e := DomainValidation(domain)
		if e != nil {
			t.Error(e)
		}
	}
}

func TestDomainValidationNotValid(t *testing.T) {
	invalidDomains := []string{
		"AAAAAAAAAAA",
		"محمد.آیران",
		"😀",
		"",
		"1",
	}
	for _, domain := range invalidDomains {
		_, e := DomainValidation(domain)
		if e == nil {
			t.Errorf("Domain is invalid so we need an error")
		}
	}
}

func TestGetStructureWhoIsDataNotValid(t *testing.T) {
	_, e := GetStructureWhoIsData("AAA")
	if e == nil {
		t.Errorf("Domain is invalid so we need an error")
	}
}
func TestDomainParse(t *testing.T) {
	domainType, _ := DomainValidation("www.nic.ir")
	result, e := DomainParse(domainType)
	if e != nil {
		t.Error(e)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}
