package client

import (
	"crypto/tls"
	"encoding/base64"
	"net/http"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/ontap-go/api/client"
)

type Config struct {
	AdminUser     string
	AdminPassword string
	Host          string
	InsecureTLS   bool
}

func NewAPIClient(cfg Config) (*client.Ontap, error) {
	transport := httptransport.New(cfg.Host, "/api", []string{"https"})

	//  API call error: no consumer: "application/hal+json", need to fix this
	jsonConsumer := runtime.JSONConsumer()
	transport.Consumers["application/hal+json"] = jsonConsumer

	// Insecure skip verify if needed just for testing
	if httpTrans, ok := transport.Transport.(*http.Transport); ok && cfg.InsecureTLS {
		httpTrans.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //nolint:all
	}

	// Basic Auth for now
	userpass := base64.StdEncoding.EncodeToString([]byte(cfg.AdminUser + ":" + cfg.AdminPassword))
	transport.DefaultAuthentication = runtime.ClientAuthInfoWriterFunc(
		func(req runtime.ClientRequest, reg strfmt.Registry) error {
			return req.SetHeaderParam("Authorization", "Basic "+userpass)
		},
	)

	client := client.New(transport, strfmt.Default)
	return client, nil
}
