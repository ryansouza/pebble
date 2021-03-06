// addDNS01 handles an HTTP POST request to add a new DNS-01 challenge TXT
package main

import (
	"net/http"
	"strings"

	"github.com/letsencrypt/challtestsrv"
)

// setDefaultDNSIPv4 handles an HTTP POST request to set the default IPv4
// address used for all A query responses that do not match more-specific mocked
// responses.
//
// The POST body is expected to have one parameter:
// "ip" - the string representation of an IPv4 address to use for all A queries
//        that do not match more specific mocks.
//
// Providing an empty string as the IP value will disable the default
// A responses.
//
// A successful POST will write http.StatusOK to the client.
func (srv *managementServer) setDefaultDNSIPv4(w http.ResponseWriter, r *http.Request) {
	var request struct {
		IP string
	}
	if err := mustParsePOST(&request, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the challenge server's default IPv4 address - we allow request.IP to be
	// the empty string so that the default can be be cleared using the same
	// method.
	srv.challSrv.SetDefaultDNSIPv4(request.IP)
	srv.log.Printf("Set default IPv4 address for DNS A queries to %q\n", request.IP)
	w.WriteHeader(http.StatusOK)
}

// setDefaultDNSIPv6 handles an HTTP POST request to set the default IPv6
// address used for all AAAA query responses that do not match more-specific
// mocked responses.
//
// The POST body is expected to have one parameter:
// "ip" - the string representation of an IPv6 address to use for all AAAA
//        queries that do not match more specific mocks.
//
// Providing an empty string as the IP value will disable the default
// A responses.
//
// A successful POST will write http.StatusOK to the client.
func (srv *managementServer) setDefaultDNSIPv6(w http.ResponseWriter, r *http.Request) {
	var request struct {
		IP string
	}
	if err := mustParsePOST(&request, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the challenge server's default IPv6 address - we allow request.IP to be
	// the empty string so that the default can be be cleared using the same
	// method.
	srv.challSrv.SetDefaultDNSIPv6(request.IP)
	srv.log.Printf("Set default IPv6 address for DNS AAAA queries to %q\n", request.IP)
	w.WriteHeader(http.StatusOK)
}

// addDNSARecord handles an HTTP POST request to add a mock A query response record
// for a host.
//
// The POST body is expected to have two non-empty parameters:
// "host" - the hostname that when queried should return the mocked A record.
// "addresses" - an array of IPv4 addresses in string representation that should
// be used for the A records returned for the query.
//
// A successful POST will write http.StatusOK to the client.
func (srv *managementServer) addDNSARecord(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Host      string
		Addresses []string
	}
	if err := mustParsePOST(&request, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// If the request has no addresses or an empty host it's a bad request
	if len(request.Addresses) == 0 || request.Host == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	srv.challSrv.AddDNSARecord(request.Host, request.Addresses)
	srv.log.Printf("Added response for DNS A queries to %q : %s\n",
		request.Host, strings.Join(request.Addresses, ", "))
	w.WriteHeader(http.StatusOK)
}

// delDNSARecord handles an HTTP POST request to delete an existing mock A
// policy record for a host.
//
// The POST body is expected to have one non-empty parameter:
// "host" - the hostname to remove the mock A record for.
//
// A successful POST will write http.StatusOK to the client.
func (srv *managementServer) delDNSARecord(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Host string
	}
	if err := mustParsePOST(&request, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// If the request has an empty host it's a bad request
	if request.Host == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	srv.challSrv.DeleteDNSARecord(request.Host)
	srv.log.Printf("Removed response for DNS A queries to %q", request.Host)
	w.WriteHeader(http.StatusOK)
}

// addDNSAAAARecord handles an HTTP POST request to add a mock AAAA query
// response record for a host.
//
// The POST body is expected to have two non-empty parameters:
// "host" - the hostname that when queried should return the mocked A record.
// "addresses" - an array of IPv6 addresses in string representation that should
// be used for the AAAA records returned for the query.
//
// A successful POST will write http.StatusOK to the client.
func (srv *managementServer) addDNSAAAARecord(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Host      string
		Addresses []string
	}
	if err := mustParsePOST(&request, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// If the request has no addresses or an empty host it's a bad request
	if len(request.Addresses) == 0 || request.Host == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	srv.challSrv.AddDNSAAAARecord(request.Host, request.Addresses)
	srv.log.Printf("Added response for DNS AAAA queries to %q : %s\n",
		request.Host, strings.Join(request.Addresses, ", "))
	w.WriteHeader(http.StatusOK)
}

// delDNSAAAARecord handles an HTTP POST request to delete an existing mock AAAA
// policy record for a host.
//
// The POST body is expected to have one non-empty parameter:
// "host" - the hostname to remove the mock AAAA record for.
//
// A successful POST will write http.StatusOK to the client.
func (srv *managementServer) delDNSAAAARecord(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Host string
	}
	if err := mustParsePOST(&request, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// If the request has an empty host it's a bad request
	if request.Host == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	srv.challSrv.DeleteDNSAAAARecord(request.Host)
	srv.log.Printf("Removed response for DNS AAAA queries to %q", request.Host)
	w.WriteHeader(http.StatusOK)
}

// addDNSCAARecord handles an HTTP POST request to add a mock CAA query
// response record for a host.
//
// The POST body is expected to have two non-empty parameters:
// "host" - the hostname that when queried should return the mocked CAA record.
// "policies" - an array of CAA policy objects. Each policy object is expected
// to have two non-empty keys, "tag" and "value".
//
// A successful POST will write http.StatusOK to the client.
func (srv *managementServer) addDNSCAARecord(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Host     string
		Policies []challtestsrv.MockCAAPolicy
	}
	if err := mustParsePOST(&request, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// If the request has no host or no caa policies it's a bad request
	if request.Host == "" || len(request.Policies) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	srv.challSrv.AddDNSCAARecord(request.Host, request.Policies)
	srv.log.Printf("Added response for DNS CAA queries to %q", request.Host)
	w.WriteHeader(http.StatusOK)
}

// delDNSCAARecord handles an HTTP POST request to delete an existing mock CAA
// policy record for a host.
//
// The POST body is expected to have one non-empty parameter:
// "host" - the hostname to remove the mock CAA policy for.
//
// A successful POST will write http.StatusOK to the client.
func (srv *managementServer) delDNSCAARecord(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Host string
	}
	if err := mustParsePOST(&request, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// If the request has an empty host it's a bad request
	if request.Host == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	srv.challSrv.DeleteDNSCAARecord(request.Host)
	srv.log.Printf("Removed response for DNS CAA queries to %q", request.Host)
	w.WriteHeader(http.StatusOK)
}
