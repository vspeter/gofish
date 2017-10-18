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

// GetSystemTasksReader is a Reader for the GetSystemTasks structure.
type GetSystemTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSystemTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetSystemTasksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetSystemTasksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSystemTasksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSystemTasksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSystemTasksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSystemTasksOK creates a GetSystemTasksOK with default headers values
func NewGetSystemTasksOK() *GetSystemTasksOK {
	return &GetSystemTasksOK{}
}

/*GetSystemTasksOK handles this case with default header values.

Success
*/
type GetSystemTasksOK struct {
	Payload *models.TaskCollectionTaskCollection
}

func (o *GetSystemTasksOK) Error() string {
	return fmt.Sprintf("[GET /TaskService/Oem/Tasks/{identifier}][%d] getSystemTasksOK  %+v", 200, o.Payload)
}

func (o *GetSystemTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskCollectionTaskCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystemTasksBadRequest creates a GetSystemTasksBadRequest with default headers values
func NewGetSystemTasksBadRequest() *GetSystemTasksBadRequest {
	return &GetSystemTasksBadRequest{}
}

/*GetSystemTasksBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information (such as validation error on an input field, a missing required value, and so on). An extended error shall be returned in the response body, as defined in section Extended Error Handling.

*/
type GetSystemTasksBadRequest struct {
}

func (o *GetSystemTasksBadRequest) Error() string {
	return fmt.Sprintf("[GET /TaskService/Oem/Tasks/{identifier}][%d] getSystemTasksBadRequest ", 400)
}

func (o *GetSystemTasksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemTasksUnauthorized creates a GetSystemTasksUnauthorized with default headers values
func NewGetSystemTasksUnauthorized() *GetSystemTasksUnauthorized {
	return &GetSystemTasksUnauthorized{}
}

/*GetSystemTasksUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetSystemTasksUnauthorized struct {
}

func (o *GetSystemTasksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /TaskService/Oem/Tasks/{identifier}][%d] getSystemTasksUnauthorized ", 401)
}

func (o *GetSystemTasksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemTasksForbidden creates a GetSystemTasksForbidden with default headers values
func NewGetSystemTasksForbidden() *GetSystemTasksForbidden {
	return &GetSystemTasksForbidden{}
}

/*GetSystemTasksForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type GetSystemTasksForbidden struct {
}

func (o *GetSystemTasksForbidden) Error() string {
	return fmt.Sprintf("[GET /TaskService/Oem/Tasks/{identifier}][%d] getSystemTasksForbidden ", 403)
}

func (o *GetSystemTasksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemTasksNotFound creates a GetSystemTasksNotFound with default headers values
func NewGetSystemTasksNotFound() *GetSystemTasksNotFound {
	return &GetSystemTasksNotFound{}
}

/*GetSystemTasksNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type GetSystemTasksNotFound struct {
}

func (o *GetSystemTasksNotFound) Error() string {
	return fmt.Sprintf("[GET /TaskService/Oem/Tasks/{identifier}][%d] getSystemTasksNotFound ", 404)
}

func (o *GetSystemTasksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemTasksInternalServerError creates a GetSystemTasksInternalServerError with default headers values
func NewGetSystemTasksInternalServerError() *GetSystemTasksInternalServerError {
	return &GetSystemTasksInternalServerError{}
}

/*GetSystemTasksInternalServerError handles this case with default header values.

Error
*/
type GetSystemTasksInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetSystemTasksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /TaskService/Oem/Tasks/{identifier}][%d] getSystemTasksInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSystemTasksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
