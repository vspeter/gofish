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

// GetSystemProcessorReader is a Reader for the GetSystemProcessor structure.
type GetSystemProcessorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemProcessorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSystemProcessorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetSystemProcessorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetSystemProcessorUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSystemProcessorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSystemProcessorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSystemProcessorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSystemProcessorOK creates a GetSystemProcessorOK with default headers values
func NewGetSystemProcessorOK() *GetSystemProcessorOK {
	return &GetSystemProcessorOK{}
}

/*GetSystemProcessorOK handles this case with default header values.

Success
*/
type GetSystemProcessorOK struct {
	Payload *models.Processor100Processor
}

func (o *GetSystemProcessorOK) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/Processors/{socket}][%d] getSystemProcessorOK  %+v", 200, o.Payload)
}

func (o *GetSystemProcessorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Processor100Processor)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystemProcessorBadRequest creates a GetSystemProcessorBadRequest with default headers values
func NewGetSystemProcessorBadRequest() *GetSystemProcessorBadRequest {
	return &GetSystemProcessorBadRequest{}
}

/*GetSystemProcessorBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information (such as validation error on an input field, a missing required value, and so on). An extended error shall be returned in the response body, as defined in section Extended Error Handling.

*/
type GetSystemProcessorBadRequest struct {
}

func (o *GetSystemProcessorBadRequest) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/Processors/{socket}][%d] getSystemProcessorBadRequest ", 400)
}

func (o *GetSystemProcessorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemProcessorUnauthorized creates a GetSystemProcessorUnauthorized with default headers values
func NewGetSystemProcessorUnauthorized() *GetSystemProcessorUnauthorized {
	return &GetSystemProcessorUnauthorized{}
}

/*GetSystemProcessorUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetSystemProcessorUnauthorized struct {
}

func (o *GetSystemProcessorUnauthorized) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/Processors/{socket}][%d] getSystemProcessorUnauthorized ", 401)
}

func (o *GetSystemProcessorUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemProcessorForbidden creates a GetSystemProcessorForbidden with default headers values
func NewGetSystemProcessorForbidden() *GetSystemProcessorForbidden {
	return &GetSystemProcessorForbidden{}
}

/*GetSystemProcessorForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type GetSystemProcessorForbidden struct {
}

func (o *GetSystemProcessorForbidden) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/Processors/{socket}][%d] getSystemProcessorForbidden ", 403)
}

func (o *GetSystemProcessorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemProcessorNotFound creates a GetSystemProcessorNotFound with default headers values
func NewGetSystemProcessorNotFound() *GetSystemProcessorNotFound {
	return &GetSystemProcessorNotFound{}
}

/*GetSystemProcessorNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type GetSystemProcessorNotFound struct {
}

func (o *GetSystemProcessorNotFound) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/Processors/{socket}][%d] getSystemProcessorNotFound ", 404)
}

func (o *GetSystemProcessorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemProcessorInternalServerError creates a GetSystemProcessorInternalServerError with default headers values
func NewGetSystemProcessorInternalServerError() *GetSystemProcessorInternalServerError {
	return &GetSystemProcessorInternalServerError{}
}

/*GetSystemProcessorInternalServerError handles this case with default header values.

Error
*/
type GetSystemProcessorInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetSystemProcessorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/Processors/{socket}][%d] getSystemProcessorInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSystemProcessorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
