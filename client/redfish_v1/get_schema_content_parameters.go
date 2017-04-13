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

// NewGetSchemaContentParams creates a new GetSchemaContentParams object
// with the default values initialized.
func NewGetSchemaContentParams() *GetSchemaContentParams {
	var ()
	return &GetSchemaContentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSchemaContentParamsWithTimeout creates a new GetSchemaContentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSchemaContentParamsWithTimeout(timeout time.Duration) *GetSchemaContentParams {
	var ()
	return &GetSchemaContentParams{

		timeout: timeout,
	}
}

// NewGetSchemaContentParamsWithContext creates a new GetSchemaContentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSchemaContentParamsWithContext(ctx context.Context) *GetSchemaContentParams {
	var ()
	return &GetSchemaContentParams{

		Context: ctx,
	}
}

/*GetSchemaContentParams contains all the parameters to send to the API endpoint
for the get schema content operation typically these are written to a http.Request
*/
type GetSchemaContentParams struct {

	/*Identifier*/
	Identifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get schema content params
func (o *GetSchemaContentParams) WithTimeout(timeout time.Duration) *GetSchemaContentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schema content params
func (o *GetSchemaContentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schema content params
func (o *GetSchemaContentParams) WithContext(ctx context.Context) *GetSchemaContentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schema content params
func (o *GetSchemaContentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithIdentifier adds the identifier to the get schema content params
func (o *GetSchemaContentParams) WithIdentifier(identifier string) *GetSchemaContentParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the get schema content params
func (o *GetSchemaContentParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WriteToRequest writes these params to a swagger request
func (o *GetSchemaContentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
