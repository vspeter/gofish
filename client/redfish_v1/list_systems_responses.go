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

// ListSystemsReader is a Reader for the ListSystems structure.
type ListSystemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSystemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSystemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListSystemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListSystemsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListSystemsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListSystemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListSystemsOK creates a ListSystemsOK with default headers values
func NewListSystemsOK() *ListSystemsOK {
	return &ListSystemsOK{}
}

/*ListSystemsOK handles this case with default header values.

Success
*/
type ListSystemsOK struct {
	Payload *models.ComputerSystemCollectionComputerSystemCollection
}

func (o *ListSystemsOK) Error() string {
	return fmt.Sprintf("[GET /Systems][%d] listSystemsOK  %+v", 200, o.Payload)
}

func (o *ListSystemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputerSystemCollectionComputerSystemCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSystemsBadRequest creates a ListSystemsBadRequest with default headers values
func NewListSystemsBadRequest() *ListSystemsBadRequest {
	return &ListSystemsBadRequest{}
}

/*ListSystemsBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information (such as validation error on an input field, a missing required value, and so on). An extended error shall be returned in the response body, as defined in section Extended Error Handling.

*/
type ListSystemsBadRequest struct {
}

func (o *ListSystemsBadRequest) Error() string {
	return fmt.Sprintf("[GET /Systems][%d] listSystemsBadRequest ", 400)
}

func (o *ListSystemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListSystemsUnauthorized creates a ListSystemsUnauthorized with default headers values
func NewListSystemsUnauthorized() *ListSystemsUnauthorized {
	return &ListSystemsUnauthorized{}
}

/*ListSystemsUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type ListSystemsUnauthorized struct {
}

func (o *ListSystemsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /Systems][%d] listSystemsUnauthorized ", 401)
}

func (o *ListSystemsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListSystemsForbidden creates a ListSystemsForbidden with default headers values
func NewListSystemsForbidden() *ListSystemsForbidden {
	return &ListSystemsForbidden{}
}

/*ListSystemsForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type ListSystemsForbidden struct {
}

func (o *ListSystemsForbidden) Error() string {
	return fmt.Sprintf("[GET /Systems][%d] listSystemsForbidden ", 403)
}

func (o *ListSystemsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListSystemsInternalServerError creates a ListSystemsInternalServerError with default headers values
func NewListSystemsInternalServerError() *ListSystemsInternalServerError {
	return &ListSystemsInternalServerError{}
}

/*ListSystemsInternalServerError handles this case with default header values.

Error
*/
type ListSystemsInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ListSystemsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /Systems][%d] listSystemsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListSystemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
