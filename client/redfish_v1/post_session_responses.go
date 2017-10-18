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

// PostSessionReader is a Reader for the PostSession structure.
type PostSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostSessionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostSessionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostSessionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostSessionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostSessionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostSessionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSessionCreated creates a PostSessionCreated with default headers values
func NewPostSessionCreated() *PostSessionCreated {
	return &PostSessionCreated{}
}

/*PostSessionCreated handles this case with default header values.

Created
*/
type PostSessionCreated struct {
	Payload *models.Session100Session
}

func (o *PostSessionCreated) Error() string {
	return fmt.Sprintf("[POST /SessionService/Sessions][%d] postSessionCreated  %+v", 201, o.Payload)
}

func (o *PostSessionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Session100Session)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSessionBadRequest creates a PostSessionBadRequest with default headers values
func NewPostSessionBadRequest() *PostSessionBadRequest {
	return &PostSessionBadRequest{}
}

/*PostSessionBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information (such as validation error on an input field, a missing required value, and so on). An extended error shall be returned in the response body, as defined in section Extended Error Handling.

*/
type PostSessionBadRequest struct {
}

func (o *PostSessionBadRequest) Error() string {
	return fmt.Sprintf("[POST /SessionService/Sessions][%d] postSessionBadRequest ", 400)
}

func (o *PostSessionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSessionUnauthorized creates a PostSessionUnauthorized with default headers values
func NewPostSessionUnauthorized() *PostSessionUnauthorized {
	return &PostSessionUnauthorized{}
}

/*PostSessionUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type PostSessionUnauthorized struct {
}

func (o *PostSessionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /SessionService/Sessions][%d] postSessionUnauthorized ", 401)
}

func (o *PostSessionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSessionForbidden creates a PostSessionForbidden with default headers values
func NewPostSessionForbidden() *PostSessionForbidden {
	return &PostSessionForbidden{}
}

/*PostSessionForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type PostSessionForbidden struct {
}

func (o *PostSessionForbidden) Error() string {
	return fmt.Sprintf("[POST /SessionService/Sessions][%d] postSessionForbidden ", 403)
}

func (o *PostSessionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSessionNotFound creates a PostSessionNotFound with default headers values
func NewPostSessionNotFound() *PostSessionNotFound {
	return &PostSessionNotFound{}
}

/*PostSessionNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type PostSessionNotFound struct {
}

func (o *PostSessionNotFound) Error() string {
	return fmt.Sprintf("[POST /SessionService/Sessions][%d] postSessionNotFound ", 404)
}

func (o *PostSessionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSessionInternalServerError creates a PostSessionInternalServerError with default headers values
func NewPostSessionInternalServerError() *PostSessionInternalServerError {
	return &PostSessionInternalServerError{}
}

/*PostSessionInternalServerError handles this case with default header values.

Error
*/
type PostSessionInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PostSessionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /SessionService/Sessions][%d] postSessionInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSessionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
