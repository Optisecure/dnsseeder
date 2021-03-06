package main

import (
	"net"
	"time"

	"github.com/gcash/bchd/wire"
)

// Node struct contains details on one client
type node struct {
	NA           *wire.NetAddress // holds ip address & port details
	LastConnect  time.Time        // last time we successfully connected to this client
	LastTry      time.Time        // last time we tried to connect to this client
	CrawlStart   time.Time        // time when we started the last crawl
	NonstdIP     net.IP           // if not using the default port then this is the encoded ip containing the actual port
	StatusStr    string           // string with last error or OK details
	StrVersion   string           // remote client user agent
	Services     wire.ServiceFlag // remote client supported services
	ConnectFails uint32           // number of times we have failed to connect to this client
	Version      int32            // remote client protocol version
	LastBlock    int32            // remote client last block
	Status       uint32           // rg,cg,wg,ng
	Rating       uint32           // if it reaches 100 then we mark them statusNG
	DNSType      uint32           // what dns type this client is
	CrawlActive  bool             // are we currently crawling this client
}

// dns2str will return the string description of the dns type
func (nd node) dns2str() string {
	switch nd.DNSType {
	case dnsV4Std:
		return "v4 standard port"
	case dnsV4Non:
		return "v4 non-standard port"
	case dnsV6Std:
		return "v6 standard port"
	case dnsV6Non:
		return "v6 non-standard port"
	default:
		return "Unknown DNS Type"
	}
}
