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

// GetRoleReader is a Reader for the GetRole structure.
type GetRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRoleOK creates a GetRoleOK with default headers values
func NewGetRoleOK() *GetRoleOK {
	return &GetRoleOK{}
}

/*GetRoleOK handles this case with default header values.

Success
*/
type GetRoleOK struct {
	Payload *models.Role100Role
}

func (o *GetRoleOK) Error() string {
	return fmt.Sprintf("[GET /AccountService/Roles/{identifier}][%d] getRoleOK  %+v", 200, o.Payload)
}

func (o *GetRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role100Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleBadRequest creates a GetRoleBadRequest with default headers values
func NewGetRoleBadRequest() *GetRoleBadRequest {
	return &GetRoleBadRequest{}
}

/*GetRoleBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information (such as validation error on an input field, a missing required value, and so on). An extended error shall be returned in the response body, as defined in section Extended Error Handling.

*/
type GetRoleBadRequest struct {
}

func (o *GetRoleBadRequest) Error() string {
	return fmt.Sprintf("[GET /AccountService/Roles/{identifier}][%d] getRoleBadRequest ", 400)
}

func (o *GetRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleUnauthorized creates a GetRoleUnauthorized with default headers values
func NewGetRoleUnauthorized() *GetRoleUnauthorized {
	return &GetRoleUnauthorized{}
}

/*GetRoleUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetRoleUnauthorized struct {
}

func (o *GetRoleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /AccountService/Roles/{identifier}][%d] getRoleUnauthorized ", 401)
}

func (o *GetRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleForbidden creates a GetRoleForbidden with default headers values
func NewGetRoleForbidden() *GetRoleForbidden {
	return &GetRoleForbidden{}
}

/*GetRoleForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type GetRoleForbidden struct {
}

func (o *GetRoleForbidden) Error() string {
	return fmt.Sprintf("[GET /AccountService/Roles/{identifier}][%d] getRoleForbidden ", 403)
}

func (o *GetRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleNotFound creates a GetRoleNotFound with default headers values
func NewGetRoleNotFound() *GetRoleNotFound {
	return &GetRoleNotFound{}
}

/*GetRoleNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type GetRoleNotFound struct {
}

func (o *GetRoleNotFound) Error() string {
	return fmt.Sprintf("[GET /AccountService/Roles/{identifier}][%d] getRoleNotFound ", 404)
}

func (o *GetRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleInternalServerError creates a GetRoleInternalServerError with default headers values
func NewGetRoleInternalServerError() *GetRoleInternalServerError {
	return &GetRoleInternalServerError{}
}

/*GetRoleInternalServerError handles this case with default header values.

Error
*/
type GetRoleInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetRoleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /AccountService/Roles/{identifier}][%d] getRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
