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

// NewDeleteEventParams creates a new DeleteEventParams object
// with the default values initialized.
func NewDeleteEventParams() *DeleteEventParams {
	var ()
	return &DeleteEventParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEventParamsWithTimeout creates a new DeleteEventParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteEventParamsWithTimeout(timeout time.Duration) *DeleteEventParams {
	var ()
	return &DeleteEventParams{

		timeout: timeout,
	}
}

// NewDeleteEventParamsWithContext creates a new DeleteEventParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteEventParamsWithContext(ctx context.Context) *DeleteEventParams {
	var ()
	return &DeleteEventParams{

		Context: ctx,
	}
}

/*DeleteEventParams contains all the parameters to send to the API endpoint
for the delete event operation typically these are written to a http.Request
*/
type DeleteEventParams struct {

	/*Index*/
	Index string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete event params
func (o *DeleteEventParams) WithTimeout(timeout time.Duration) *DeleteEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete event params
func (o *DeleteEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete event params
func (o *DeleteEventParams) WithContext(ctx context.Context) *DeleteEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete event params
func (o *DeleteEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithIndex adds the index to the delete event params
func (o *DeleteEventParams) WithIndex(index string) *DeleteEventParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the delete event params
func (o *DeleteEventParams) SetIndex(index string) {
	o.Index = index
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param index
	if err := r.SetPathParam("index", o.Index); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
