package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSystemParams creates a new GetSystemParams object
// with the default values initialized.
func NewGetSystemParams() *GetSystemParams {
	var ()
	return &GetSystemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemParamsWithTimeout creates a new GetSystemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSystemParamsWithTimeout(timeout time.Duration) *GetSystemParams {
	var ()
	return &GetSystemParams{

		timeout: timeout,
	}
}

// NewGetSystemParamsWithContext creates a new GetSystemParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSystemParamsWithContext(ctx context.Context) *GetSystemParams {
	var ()
	return &GetSystemParams{

		Context: ctx,
	}
}

/*GetSystemParams contains all the parameters to send to the API endpoint
for the get system operation typically these are written to a http.Request
*/
type GetSystemParams struct {

	/*Identifier
	  node identifier

	*/
	Identifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get system params
func (o *GetSystemParams) WithTimeout(timeout time.Duration) *GetSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system params
func (o *GetSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system params
func (o *GetSystemParams) WithContext(ctx context.Context) *GetSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system params
func (o *GetSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithIdentifier adds the identifier to the get system params
func (o *GetSystemParams) WithIdentifier(identifier string) *GetSystemParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the get system params
func (o *GetSystemParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
