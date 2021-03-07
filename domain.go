package main

import (
	"errors"

	"golang.org/x/net/idna"
	"golang.org/x/net/publicsuffix"
)

// NewDomain will create new valid public suffix base icann registered domain
func NewDomain(domain string) (Domain, error) {
	var p *idna.Profile
	p = idna.New()
	idnDomain, e := p.ToASCII(domain)
	if e != nil {
		return Domain{}, e
	}

	eTLD, icann := publicsuffix.PublicSuffix(idnDomain)
	if icann == false {
		return Domain{}, errors.New("Domain not valid ICANN")
	}

	dt := Domain{}

	dt.ASCII, _ = publicsuffix.EffectiveTLDPlusOne(idnDomain)
	dt.TLDASCII = eTLD

	dt.Unicode, _ = p.ToUnicode(dt.ASCII)
	dt.TLDUnicode, _ = p.ToUnicode(eTLD)

	return dt, nil
}
