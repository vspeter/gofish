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

// ListSimpleStorageReader is a Reader for the ListSimpleStorage structure.
type ListSimpleStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSimpleStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSimpleStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListSimpleStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListSimpleStorageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListSimpleStorageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListSimpleStorageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListSimpleStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListSimpleStorageOK creates a ListSimpleStorageOK with default headers values
func NewListSimpleStorageOK() *ListSimpleStorageOK {
	return &ListSimpleStorageOK{}
}

/*ListSimpleStorageOK handles this case with default header values.

Success
*/
type ListSimpleStorageOK struct {
	Payload *models.SimpleStorageCollectionSimpleStorageCollection
}

func (o *ListSimpleStorageOK) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/SimpleStorage][%d] listSimpleStorageOK  %+v", 200, o.Payload)
}

func (o *ListSimpleStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleStorageCollectionSimpleStorageCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSimpleStorageBadRequest creates a ListSimpleStorageBadRequest with default headers values
func NewListSimpleStorageBadRequest() *ListSimpleStorageBadRequest {
	return &ListSimpleStorageBadRequest{}
}

/*ListSimpleStorageBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information (such as validation error on an input field, a missing required value, and so on). An extended error shall be returned in the response body, as defined in section Extended Error Handling.

*/
type ListSimpleStorageBadRequest struct {
}

func (o *ListSimpleStorageBadRequest) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/SimpleStorage][%d] listSimpleStorageBadRequest ", 400)
}

func (o *ListSimpleStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListSimpleStorageUnauthorized creates a ListSimpleStorageUnauthorized with default headers values
func NewListSimpleStorageUnauthorized() *ListSimpleStorageUnauthorized {
	return &ListSimpleStorageUnauthorized{}
}

/*ListSimpleStorageUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type ListSimpleStorageUnauthorized struct {
}

func (o *ListSimpleStorageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/SimpleStorage][%d] listSimpleStorageUnauthorized ", 401)
}

func (o *ListSimpleStorageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListSimpleStorageForbidden creates a ListSimpleStorageForbidden with default headers values
func NewListSimpleStorageForbidden() *ListSimpleStorageForbidden {
	return &ListSimpleStorageForbidden{}
}

/*ListSimpleStorageForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type ListSimpleStorageForbidden struct {
}

func (o *ListSimpleStorageForbidden) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/SimpleStorage][%d] listSimpleStorageForbidden ", 403)
}

func (o *ListSimpleStorageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListSimpleStorageNotFound creates a ListSimpleStorageNotFound with default headers values
func NewListSimpleStorageNotFound() *ListSimpleStorageNotFound {
	return &ListSimpleStorageNotFound{}
}

/*ListSimpleStorageNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type ListSimpleStorageNotFound struct {
}

func (o *ListSimpleStorageNotFound) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/SimpleStorage][%d] listSimpleStorageNotFound ", 404)
}

func (o *ListSimpleStorageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListSimpleStorageInternalServerError creates a ListSimpleStorageInternalServerError with default headers values
func NewListSimpleStorageInternalServerError() *ListSimpleStorageInternalServerError {
	return &ListSimpleStorageInternalServerError{}
}

/*ListSimpleStorageInternalServerError handles this case with default header values.

Error
*/
type ListSimpleStorageInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ListSimpleStorageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/SimpleStorage][%d] listSimpleStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *ListSimpleStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
