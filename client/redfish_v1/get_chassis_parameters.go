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

// NewGetChassisParams creates a new GetChassisParams object
// with the default values initialized.
func NewGetChassisParams() *GetChassisParams {
	var ()
	return &GetChassisParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetChassisParamsWithTimeout creates a new GetChassisParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetChassisParamsWithTimeout(timeout time.Duration) *GetChassisParams {
	var ()
	return &GetChassisParams{

		timeout: timeout,
	}
}

// NewGetChassisParamsWithContext creates a new GetChassisParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetChassisParamsWithContext(ctx context.Context) *GetChassisParams {
	var ()
	return &GetChassisParams{

		Context: ctx,
	}
}

/*GetChassisParams contains all the parameters to send to the API endpoint
for the get chassis operation typically these are written to a http.Request
*/
type GetChassisParams struct {

	/*Identifier*/
	Identifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get chassis params
func (o *GetChassisParams) WithTimeout(timeout time.Duration) *GetChassisParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get chassis params
func (o *GetChassisParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get chassis params
func (o *GetChassisParams) WithContext(ctx context.Context) *GetChassisParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get chassis params
func (o *GetChassisParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithIdentifier adds the identifier to the get chassis params
func (o *GetChassisParams) WithIdentifier(identifier string) *GetChassisParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the get chassis params
func (o *GetChassisParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WriteToRequest writes these params to a swagger request
func (o *GetChassisParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
