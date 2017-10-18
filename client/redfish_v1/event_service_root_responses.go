package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vspeter/gofish/models"
)

// EventServiceRootReader is a Reader for the EventServiceRoot structure.
type EventServiceRootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EventServiceRootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEventServiceRootOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewEventServiceRootUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewEventServiceRootForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEventServiceRootOK creates a EventServiceRootOK with default headers values
func NewEventServiceRootOK() *EventServiceRootOK {
	return &EventServiceRootOK{}
}

/*EventServiceRootOK handles this case with default header values.

Success
*/
type EventServiceRootOK struct {
	Payload *models.EventService100EventService
}

func (o *EventServiceRootOK) Error() string {
	return fmt.Sprintf("[GET /EventService][%d] eventServiceRootOK  %+v", 200, o.Payload)
}

func (o *EventServiceRootOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventService100EventService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEventServiceRootUnauthorized creates a EventServiceRootUnauthorized with default headers values
func NewEventServiceRootUnauthorized() *EventServiceRootUnauthorized {
	return &EventServiceRootUnauthorized{}
}

/*EventServiceRootUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type EventServiceRootUnauthorized struct {
}

func (o *EventServiceRootUnauthorized) Error() string {
	return fmt.Sprintf("[GET /EventService][%d] eventServiceRootUnauthorized ", 401)
}

func (o *EventServiceRootUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEventServiceRootForbidden creates a EventServiceRootForbidden with default headers values
func NewEventServiceRootForbidden() *EventServiceRootForbidden {
	return &EventServiceRootForbidden{}
}

/*EventServiceRootForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type EventServiceRootForbidden struct {
}

func (o *EventServiceRootForbidden) Error() string {
	return fmt.Sprintf("[GET /EventService][%d] eventServiceRootForbidden ", 403)
}

func (o *EventServiceRootForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
