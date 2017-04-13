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

// NewEventServiceRootParams creates a new EventServiceRootParams object
// with the default values initialized.
func NewEventServiceRootParams() *EventServiceRootParams {

	return &EventServiceRootParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEventServiceRootParamsWithTimeout creates a new EventServiceRootParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEventServiceRootParamsWithTimeout(timeout time.Duration) *EventServiceRootParams {

	return &EventServiceRootParams{

		timeout: timeout,
	}
}

// NewEventServiceRootParamsWithContext creates a new EventServiceRootParams object
// with the default values initialized, and the ability to set a context for a request
func NewEventServiceRootParamsWithContext(ctx context.Context) *EventServiceRootParams {

	return &EventServiceRootParams{

		Context: ctx,
	}
}

/*EventServiceRootParams contains all the parameters to send to the API endpoint
for the event service root operation typically these are written to a http.Request
*/
type EventServiceRootParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the event service root params
func (o *EventServiceRootParams) WithTimeout(timeout time.Duration) *EventServiceRootParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the event service root params
func (o *EventServiceRootParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the event service root params
func (o *EventServiceRootParams) WithContext(ctx context.Context) *EventServiceRootParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the event service root params
func (o *EventServiceRootParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *EventServiceRootParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
