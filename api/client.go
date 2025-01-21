package api

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"

	"software.sslmate.com/src/go-pkcs12"
)

var baseApiUrl = "https://idplab.hemma.local:9443"
var pfxPath = "client.pfx"
var pfxPassword = "12345"
var rootCAPath = "ca.crt"

func InitApiClient() (*http.Client, error) {
	pfxData, err := os.ReadFile(pfxPath)
	if err != nil {
		return nil, fmt.Errorf("error reading PFX file: %w", err)
	}

	// Decode PFX file
	privateKey, certificate, _, err := pkcs12.DecodeChain(pfxData, pfxPassword)
	if err != nil {
		return nil, fmt.Errorf("error decoding PFX file: %w", err)
	}

	// Load CA certificate
	caCert, err := os.ReadFile(rootCAPath)
	if err != nil {
		return nil, fmt.Errorf("error reading CA certificate: %w", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS certificate
	cert := tls.Certificate{
		Certificate: [][]byte{certificate.Raw},
		PrivateKey:  privateKey,
	}

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: true,
	}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	return client, nil
}
