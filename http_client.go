package utils

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/pkg/errors"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

// NewHttpClientTrustCA Returns a http client with a 20-second timeout.
// The certServerName arg is so that we can tell the client what hostname it should
// expect the server to respond with. The caCert arg is so that we
// can specify the certificate authority cert so that we can verify the self-signed ssl cert
func NewHttpClientTrustCA(certServerName string, caCert string) (*http.Client, error) {

	// Get the SystemCertPool, continue with an empty pool on error
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	// Read in the cert file
	certs, err := ioutil.ReadFile(caCert)
	if err != nil {
		return nil, errors.Errorf("Failed to append %q to RootCAs: %v\n", caCert, err)
	}

	// Append our cert to the system pool
	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		return nil, errors.New("Could not append CA to trusted system certs when creating http client")
	}

	var netTransport *http.Transport

	netTransport = &http.Transport{
		TLSClientConfig: &tls.Config{
			ServerName: certServerName, // "example.com"
			RootCAs: rootCAs,
		},
		DialContext: (&net.Dialer{
			Timeout: 20 * time.Second, // TCP connect timeout
		}).DialContext,
		TLSHandshakeTimeout: 20 * time.Second, // TLS handshake timeout
	}

	var client = &http.Client{
		Timeout:   time.Second * 20, // response timeout
		Transport: netTransport,
		// make sure headers are copied on redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Honor golangs default of maximum of 10 redirects
			if len(via) >= 10 {
				return http.ErrUseLastResponse
			}

			// Copy the headers from last request
			req.Header = via[len(via)-1].Header
			return nil
		},
	}

	return client, nil
}

func NewHttpClientWithTimeout() *http.Client {

	var netTransport *http.Transport

	netTransport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 20 * time.Second, // TCP connect timeout
		}).DialContext,
		TLSHandshakeTimeout: 20 * time.Second, // TLS handshake timeout
	}

	var client = &http.Client{
		Timeout:   time.Second * 20, // response timeout
		Transport: netTransport,
		// make sure headers are copied on redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Honor golangs default of maximum of 10 redirects
			if len(via) >= 10 {
				return http.ErrUseLastResponse
			}

			// Copy the headers from last request
			req.Header = via[len(via)-1].Header
			return nil
		},
	}

	return client
}

func NewHttpClientWithCustomTimeout(timeout time.Duration) *http.Client {

	var netTransport *http.Transport

	netTransport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: timeout, // TCP connect timeout
		}).DialContext,
		TLSHandshakeTimeout: timeout, // TLS handshake timeout
	}

	var client = &http.Client{
		Timeout:   timeout, // response timeout
		Transport: netTransport,
		// make sure headers are copied on redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Honor golangs default of maximum of 10 redirects
			if len(via) >= 10 {
				return http.ErrUseLastResponse
			}

			// Copy the headers from last request
			req.Header = via[len(via)-1].Header
			return nil
		},
	}

	return client
}

// CloneRequest Returns a clone of the provided *http.Request. The clone is a
// shallow copy of the struct and its Header map.
// https://stackoverflow.com/questions/43447405/change-http-client-transport
// https://github.com/google/go-github/blob/d23570d44313ca73dbcaadec71fc43eca4d29f8b/github/github.go#L841-L875
func CloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}