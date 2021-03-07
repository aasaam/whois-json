package main

import (
	"time"

	"github.com/araddon/dateparse"
	whois "github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

// DomainWhois will return structured whois data
func DomainWhois(domain Domain) (WhoisData, error) {
	t := time.Now()

	whoisRaw, e1 := whois.Whois(domain.Unicode)
	if e1 != nil {
		return WhoisData{}, e1
	}

	whoisInfo, e2 := whoisparser.Parse(whoisRaw)
	if e2 != nil {
		return WhoisData{}, e2
	}

	createdDate, e := dateparse.ParseLocal(whoisInfo.Domain.CreatedDate)
	if e == nil {
		whoisInfo.Domain.CreatedDate = createdDate.Format(time.RFC3339)
	}

	expirationDate, e := dateparse.ParseLocal(whoisInfo.Domain.ExpirationDate)
	if e == nil {
		whoisInfo.Domain.ExpirationDate = expirationDate.Format(time.RFC3339)
	}

	updatedDate, e := dateparse.ParseLocal(whoisInfo.Domain.UpdatedDate)
	if e == nil {
		whoisInfo.Domain.UpdatedDate = updatedDate.Format(time.RFC3339)
	}

	result := WhoisData{}

	result.Date = t.Format(time.RFC3339)
	result.WhoIs = whoisInfo
	result.Domain = domain
	result.Raw = whoisRaw

	return result, nil
}
