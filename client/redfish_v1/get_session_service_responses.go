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

// GetSessionServiceReader is a Reader for the GetSessionService structure.
type GetSessionServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSessionServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSessionServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetSessionServiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetSessionServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSessionServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSessionServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSessionServiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSessionServiceOK creates a GetSessionServiceOK with default headers values
func NewGetSessionServiceOK() *GetSessionServiceOK {
	return &GetSessionServiceOK{}
}

/*GetSessionServiceOK handles this case with default header values.

Success
*/
type GetSessionServiceOK struct {
	Payload *models.SessionService100SessionService
}

func (o *GetSessionServiceOK) Error() string {
	return fmt.Sprintf("[GET /SessionService][%d] getSessionServiceOK  %+v", 200, o.Payload)
}

func (o *GetSessionServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SessionService100SessionService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSessionServiceBadRequest creates a GetSessionServiceBadRequest with default headers values
func NewGetSessionServiceBadRequest() *GetSessionServiceBadRequest {
	return &GetSessionServiceBadRequest{}
}

/*GetSessionServiceBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information (such as validation error on an input field, a missing required value, and so on). An extended error shall be returned in the response body, as defined in section Extended Error Handling.

*/
type GetSessionServiceBadRequest struct {
}

func (o *GetSessionServiceBadRequest) Error() string {
	return fmt.Sprintf("[GET /SessionService][%d] getSessionServiceBadRequest ", 400)
}

func (o *GetSessionServiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSessionServiceUnauthorized creates a GetSessionServiceUnauthorized with default headers values
func NewGetSessionServiceUnauthorized() *GetSessionServiceUnauthorized {
	return &GetSessionServiceUnauthorized{}
}

/*GetSessionServiceUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetSessionServiceUnauthorized struct {
}

func (o *GetSessionServiceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /SessionService][%d] getSessionServiceUnauthorized ", 401)
}

func (o *GetSessionServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSessionServiceForbidden creates a GetSessionServiceForbidden with default headers values
func NewGetSessionServiceForbidden() *GetSessionServiceForbidden {
	return &GetSessionServiceForbidden{}
}

/*GetSessionServiceForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type GetSessionServiceForbidden struct {
}

func (o *GetSessionServiceForbidden) Error() string {
	return fmt.Sprintf("[GET /SessionService][%d] getSessionServiceForbidden ", 403)
}

func (o *GetSessionServiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSessionServiceNotFound creates a GetSessionServiceNotFound with default headers values
func NewGetSessionServiceNotFound() *GetSessionServiceNotFound {
	return &GetSessionServiceNotFound{}
}

/*GetSessionServiceNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type GetSessionServiceNotFound struct {
}

func (o *GetSessionServiceNotFound) Error() string {
	return fmt.Sprintf("[GET /SessionService][%d] getSessionServiceNotFound ", 404)
}

func (o *GetSessionServiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSessionServiceInternalServerError creates a GetSessionServiceInternalServerError with default headers values
func NewGetSessionServiceInternalServerError() *GetSessionServiceInternalServerError {
	return &GetSessionServiceInternalServerError{}
}

/*GetSessionServiceInternalServerError handles this case with default header values.

Error
*/
type GetSessionServiceInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetSessionServiceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /SessionService][%d] getSessionServiceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSessionServiceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
