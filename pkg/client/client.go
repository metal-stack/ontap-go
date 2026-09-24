package client

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/ontap-go/api/client"
)

type Config struct {
	// ApiURL points to the ontap api endpoint, usually ends with "/api"
	ApiURL string
	// BasicAuth configures basic auth for api communication if provided
	BasicAuth *BasicAuthConfig
	// TLS sets client tls configuration if provided
	TLS *TLSConfig
}

type BasicAuthConfig struct {
	User     string
	Password string
}

type TLSConfig struct {
	// either set directly as bytes
	Cert []byte
	Key  []byte
	Ca   []byte

	// or set from file paths
	CertPath string
	KeyPath  string
	CaPath   string

	// InsecureTLS should only be used for devel purposes
	InsecureTLS *bool
}

func NewAPIClient(cfg Config) (*client.Ontap, error) {
	parsedURL, err := url.Parse(cfg.ApiURL)
	if err != nil {
		return nil, err
	}
	if parsedURL.Host == "" {
		return nil, fmt.Errorf("invalid ontap api url: %s, must be in the form scheme://host[:port]/basepath", cfg.ApiURL)
	}

	httpClient := http.DefaultClient

	if cfg.TLS != nil {
		tlsOptions := httptransport.TLSClientOptions{}

		if cfg.TLS.CaPath != "" {
			tlsOptions.CA = cfg.TLS.CaPath
		}
		if cfg.TLS.CertPath != "" {
			tlsOptions.Certificate = cfg.TLS.CertPath
		}
		if cfg.TLS.KeyPath != "" {
			tlsOptions.Key = cfg.TLS.KeyPath
		}
		if cfg.TLS.InsecureTLS != nil {
			tlsOptions.InsecureSkipVerify = *cfg.TLS.InsecureTLS
		}

		if len(cfg.TLS.Ca) > 0 {
			block, _ := pem.Decode([]byte(cfg.TLS.Ca))
			if block == nil {
				return nil, fmt.Errorf("failed to decode PEM block: %w", err)
			}

			ca, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return nil, err
			}

			tlsOptions.LoadedCA = ca
		}
		if len(cfg.TLS.Cert) > 0 || len(cfg.TLS.Key) > 0 {
			pair, err := tls.X509KeyPair(cfg.TLS.Cert, cfg.TLS.Key)
			if err != nil {
				return nil, err
			}

			tlsOptions.LoadedCertificate = pair.Leaf
			tlsOptions.LoadedKey = pair.PrivateKey
		}

		tlsConfig, err := httptransport.TLSClientAuth(tlsOptions)
		if err != nil {
			return nil, err
		}

		httpClient.Transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	}

	transport := httptransport.NewWithClient(parsedURL.Host, parsedURL.Path, []string{parsedURL.Scheme}, httpClient)

	//  API call error: no consumer: "application/hal+json", need to fix this
	transport.Consumers["application/hal+json"] = runtime.JSONConsumer()

	if cfg.BasicAuth != nil {
		transport.DefaultAuthentication = httptransport.BasicAuth(cfg.BasicAuth.User, cfg.BasicAuth.Password)
	}

	return client.New(transport, strfmt.Default), nil
}

// MetroClusterConfig holds the configuration for n clusters in a metro cluster.
type MetroClusterConfig []Config

// MetroClusterClient holds n API clients, one for each cluster in a metro cluster.
type MetroClusterClient []client.Ontap

// NewMetroClusterClient creates a new client for a metro cluster, which contains a client for each of the n clusters.
func NewMetroClusterClient(cfg MetroClusterConfig) (*MetroClusterClient, error) {
	var (
		metroclients MetroClusterClient
	)

	for _, config := range cfg {
		mccclient, err := NewAPIClient(config)
		if err != nil {
			return nil, err
		}
		metroclients = append(metroclients, *mccclient)
	}

	return &metroclients, nil
}
