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

// NewGetEventsCollectionParams creates a new GetEventsCollectionParams object
// with the default values initialized.
func NewGetEventsCollectionParams() *GetEventsCollectionParams {

	return &GetEventsCollectionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventsCollectionParamsWithTimeout creates a new GetEventsCollectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEventsCollectionParamsWithTimeout(timeout time.Duration) *GetEventsCollectionParams {

	return &GetEventsCollectionParams{

		timeout: timeout,
	}
}

// NewGetEventsCollectionParamsWithContext creates a new GetEventsCollectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEventsCollectionParamsWithContext(ctx context.Context) *GetEventsCollectionParams {

	return &GetEventsCollectionParams{

		Context: ctx,
	}
}

/*GetEventsCollectionParams contains all the parameters to send to the API endpoint
for the get events collection operation typically these are written to a http.Request
*/
type GetEventsCollectionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get events collection params
func (o *GetEventsCollectionParams) WithTimeout(timeout time.Duration) *GetEventsCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get events collection params
func (o *GetEventsCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get events collection params
func (o *GetEventsCollectionParams) WithContext(ctx context.Context) *GetEventsCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get events collection params
func (o *GetEventsCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventsCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
