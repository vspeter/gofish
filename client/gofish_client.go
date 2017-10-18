package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vspeter/gofish/client/redfish_v1"
)

// Default gofish HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new gofish HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Gofish {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("localhost", "/redfish/v1", []string{"http", "https"})
	return New(transport, formats)
}

// New creates a new gofish client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Gofish {
	cli := new(Gofish)
	cli.Transport = transport

	cli.RedfishV1 = redfish_v1.New(transport, formats)

	return cli
}

// Gofish is a client for gofish
type Gofish struct {
	RedfishV1 *redfish_v1.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Gofish) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.RedfishV1.SetTransport(transport)

}
