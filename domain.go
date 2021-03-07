package main

import (
	"errors"
	"time"

	"github.com/araddon/dateparse"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
	"golang.org/x/net/idna"
	"golang.org/x/net/publicsuffix"
)

// DomainType is domain and tld type
type DomainType struct {
	ASCII      string `json:"ascii"`
	Unicode    string `json:"unicode"`
	TLDASCII   string `json:"tld_ascii"`
	TLDUnicode string `json:"tld_unicode"`
}

// WhoIsData is structure who is data
type WhoIsData struct {
	Date   string                `json:"date"`
	Domain DomainType            `json:"domain"`
	WhoIs  whoisparser.WhoisInfo `json:"whois"`
	Raw    string                `json:"raw"`
}

// DomainValidation verify idn and publicsuffix
func DomainValidation(domain string) (domainType DomainType, err error) {
	var p *idna.Profile
	p = idna.New()
	idnDomain, e := p.ToASCII(domain)
	if e != nil {
		return DomainType{}, e
	}

	eTLD, icann := publicsuffix.PublicSuffix(idnDomain)
	if icann == false {
		return DomainType{}, errors.New("Domain not valid ICANN")
	}

	dt := DomainType{}

	dt.ASCII, _ = publicsuffix.EffectiveTLDPlusOne(idnDomain)
	dt.TLDASCII = eTLD

	dt.Unicode, _ = p.ToUnicode(dt.ASCII)
	dt.TLDUnicode, _ = p.ToUnicode(eTLD)

	return dt, nil
}

// GetStructureWhoIsData parse raw who is data and get who is
func GetStructureWhoIsData(rawWhoIs string) (result whoisparser.WhoisInfo, err error) {
	r, err := whoisparser.Parse(rawWhoIs)
	if err != nil {
		return whoisparser.WhoisInfo{}, err
	}

	return r, nil
}

// DomainParse get domain return parsed data
func DomainParse(domainType DomainType) (result WhoIsData, err error) {
	t := time.Now()

	rawWhoIs, e := whois.Whois(domainType.Unicode)
	if e != nil {
		return WhoIsData{}, e
	}

	rawWhoIsString := string(rawWhoIs)

	whoIsParser, e := GetStructureWhoIsData(rawWhoIsString)
	if e != nil {
		return WhoIsData{}, e
	}

	createdDate, err := dateparse.ParseLocal(whoIsParser.Domain.CreatedDate)
	if err == nil {
		whoIsParser.Domain.CreatedDate = createdDate.Format(time.RFC3339)
	}

	expirationDate, err := dateparse.ParseLocal(whoIsParser.Domain.ExpirationDate)
	if err == nil {
		whoIsParser.Domain.ExpirationDate = expirationDate.Format(time.RFC3339)
	}

	updatedDate, err := dateparse.ParseLocal(whoIsParser.Domain.UpdatedDate)
	if err == nil {
		whoIsParser.Domain.UpdatedDate = updatedDate.Format(time.RFC3339)
	}

	result.Date = t.Format(time.RFC3339)
	result.WhoIs = whoIsParser
	result.Domain = domainType
	result.Raw = rawWhoIsString

	return result, nil
}
