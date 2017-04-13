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

// NewListSystemsParams creates a new ListSystemsParams object
// with the default values initialized.
func NewListSystemsParams() *ListSystemsParams {

	return &ListSystemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSystemsParamsWithTimeout creates a new ListSystemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSystemsParamsWithTimeout(timeout time.Duration) *ListSystemsParams {

	return &ListSystemsParams{

		timeout: timeout,
	}
}

// NewListSystemsParamsWithContext creates a new ListSystemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSystemsParamsWithContext(ctx context.Context) *ListSystemsParams {

	return &ListSystemsParams{

		Context: ctx,
	}
}

/*ListSystemsParams contains all the parameters to send to the API endpoint
for the list systems operation typically these are written to a http.Request
*/
type ListSystemsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list systems params
func (o *ListSystemsParams) WithTimeout(timeout time.Duration) *ListSystemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list systems params
func (o *ListSystemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list systems params
func (o *ListSystemsParams) WithContext(ctx context.Context) *ListSystemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list systems params
func (o *ListSystemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *ListSystemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
