package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/danehans/gofish/models"
)

// GetSessionInfoReader is a Reader for the GetSessionInfo structure.
type GetSessionInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSessionInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSessionInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetSessionInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetSessionInfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSessionInfoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSessionInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSessionInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSessionInfoOK creates a GetSessionInfoOK with default headers values
func NewGetSessionInfoOK() *GetSessionInfoOK {
	return &GetSessionInfoOK{}
}

/*GetSessionInfoOK handles this case with default header values.

Success
*/
type GetSessionInfoOK struct {
	Payload *models.Session100Session
}

func (o *GetSessionInfoOK) Error() string {
	return fmt.Sprintf("[GET /SessionService/Sessions/{identifier}][%d] getSessionInfoOK  %+v", 200, o.Payload)
}

func (o *GetSessionInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Session100Session)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSessionInfoBadRequest creates a GetSessionInfoBadRequest with default headers values
func NewGetSessionInfoBadRequest() *GetSessionInfoBadRequest {
	return &GetSessionInfoBadRequest{}
}

/*GetSessionInfoBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information (such as validation error on an input field, a missing required value, and so on). An extended error shall be returned in the response body, as defined in section Extended Error Handling.

*/
type GetSessionInfoBadRequest struct {
}

func (o *GetSessionInfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /SessionService/Sessions/{identifier}][%d] getSessionInfoBadRequest ", 400)
}

func (o *GetSessionInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSessionInfoUnauthorized creates a GetSessionInfoUnauthorized with default headers values
func NewGetSessionInfoUnauthorized() *GetSessionInfoUnauthorized {
	return &GetSessionInfoUnauthorized{}
}

/*GetSessionInfoUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetSessionInfoUnauthorized struct {
}

func (o *GetSessionInfoUnauthorized) Error() string {
	return fmt.Sprintf("[GET /SessionService/Sessions/{identifier}][%d] getSessionInfoUnauthorized ", 401)
}

func (o *GetSessionInfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSessionInfoForbidden creates a GetSessionInfoForbidden with default headers values
func NewGetSessionInfoForbidden() *GetSessionInfoForbidden {
	return &GetSessionInfoForbidden{}
}

/*GetSessionInfoForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type GetSessionInfoForbidden struct {
}

func (o *GetSessionInfoForbidden) Error() string {
	return fmt.Sprintf("[GET /SessionService/Sessions/{identifier}][%d] getSessionInfoForbidden ", 403)
}

func (o *GetSessionInfoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSessionInfoNotFound creates a GetSessionInfoNotFound with default headers values
func NewGetSessionInfoNotFound() *GetSessionInfoNotFound {
	return &GetSessionInfoNotFound{}
}

/*GetSessionInfoNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type GetSessionInfoNotFound struct {
}

func (o *GetSessionInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /SessionService/Sessions/{identifier}][%d] getSessionInfoNotFound ", 404)
}

func (o *GetSessionInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSessionInfoInternalServerError creates a GetSessionInfoInternalServerError with default headers values
func NewGetSessionInfoInternalServerError() *GetSessionInfoInternalServerError {
	return &GetSessionInfoInternalServerError{}
}

/*GetSessionInfoInternalServerError handles this case with default header values.

Error
*/
type GetSessionInfoInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetSessionInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /SessionService/Sessions/{identifier}][%d] getSessionInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSessionInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
