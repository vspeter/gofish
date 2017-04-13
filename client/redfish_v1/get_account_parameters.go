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

// NewGetAccountParams creates a new GetAccountParams object
// with the default values initialized.
func NewGetAccountParams() *GetAccountParams {
	var ()
	return &GetAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountParamsWithTimeout creates a new GetAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountParamsWithTimeout(timeout time.Duration) *GetAccountParams {
	var ()
	return &GetAccountParams{

		timeout: timeout,
	}
}

// NewGetAccountParamsWithContext creates a new GetAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountParamsWithContext(ctx context.Context) *GetAccountParams {
	var ()
	return &GetAccountParams{

		Context: ctx,
	}
}

/*GetAccountParams contains all the parameters to send to the API endpoint
for the get account operation typically these are written to a http.Request
*/
type GetAccountParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get account params
func (o *GetAccountParams) WithTimeout(timeout time.Duration) *GetAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account params
func (o *GetAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account params
func (o *GetAccountParams) WithContext(ctx context.Context) *GetAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account params
func (o *GetAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithName adds the name to the get account params
func (o *GetAccountParams) WithName(name string) *GetAccountParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get account params
func (o *GetAccountParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
