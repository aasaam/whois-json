package main

import (
	whoisparser "github.com/likexian/whois-parser"
)

// Domain is domain and tld type
type Domain struct {
	ASCII      string `json:"ascii"`
	Unicode    string `json:"unicode"`
	TLDASCII   string `json:"tld_ascii"`
	TLDUnicode string `json:"tld_unicode"`
}

// WhoisData is structure whois data
type WhoisData struct {
	Date   string                `json:"date"`
	Domain Domain                `json:"domain"`
	WhoIs  whoisparser.WhoisInfo `json:"whois"`
	Raw    string                `json:"raw"`
}
