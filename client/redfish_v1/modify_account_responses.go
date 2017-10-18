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

// ModifyAccountReader is a Reader for the ModifyAccount structure.
type ModifyAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewModifyAccountAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewModifyAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewModifyAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewModifyAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewModifyAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewModifyAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyAccountAccepted creates a ModifyAccountAccepted with default headers values
func NewModifyAccountAccepted() *ModifyAccountAccepted {
	return &ModifyAccountAccepted{}
}

/*ModifyAccountAccepted handles this case with default header values.

Success
*/
type ModifyAccountAccepted struct {
	Payload *models.ErrorResponse
}

func (o *ModifyAccountAccepted) Error() string {
	return fmt.Sprintf("[PATCH /AccountService/Accounts/{name}][%d] modifyAccountAccepted  %+v", 202, o.Payload)
}

func (o *ModifyAccountAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyAccountBadRequest creates a ModifyAccountBadRequest with default headers values
func NewModifyAccountBadRequest() *ModifyAccountBadRequest {
	return &ModifyAccountBadRequest{}
}

/*ModifyAccountBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information (such as validation error on an input field, a missing required value, and so on). An extended error shall be returned in the response body, as defined in section Extended Error Handling.

*/
type ModifyAccountBadRequest struct {
}

func (o *ModifyAccountBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /AccountService/Accounts/{name}][%d] modifyAccountBadRequest ", 400)
}

func (o *ModifyAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyAccountUnauthorized creates a ModifyAccountUnauthorized with default headers values
func NewModifyAccountUnauthorized() *ModifyAccountUnauthorized {
	return &ModifyAccountUnauthorized{}
}

/*ModifyAccountUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type ModifyAccountUnauthorized struct {
}

func (o *ModifyAccountUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /AccountService/Accounts/{name}][%d] modifyAccountUnauthorized ", 401)
}

func (o *ModifyAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyAccountForbidden creates a ModifyAccountForbidden with default headers values
func NewModifyAccountForbidden() *ModifyAccountForbidden {
	return &ModifyAccountForbidden{}
}

/*ModifyAccountForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type ModifyAccountForbidden struct {
}

func (o *ModifyAccountForbidden) Error() string {
	return fmt.Sprintf("[PATCH /AccountService/Accounts/{name}][%d] modifyAccountForbidden ", 403)
}

func (o *ModifyAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyAccountNotFound creates a ModifyAccountNotFound with default headers values
func NewModifyAccountNotFound() *ModifyAccountNotFound {
	return &ModifyAccountNotFound{}
}

/*ModifyAccountNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type ModifyAccountNotFound struct {
}

func (o *ModifyAccountNotFound) Error() string {
	return fmt.Sprintf("[PATCH /AccountService/Accounts/{name}][%d] modifyAccountNotFound ", 404)
}

func (o *ModifyAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyAccountInternalServerError creates a ModifyAccountInternalServerError with default headers values
func NewModifyAccountInternalServerError() *ModifyAccountInternalServerError {
	return &ModifyAccountInternalServerError{}
}

/*ModifyAccountInternalServerError handles this case with default header values.

Error
*/
type ModifyAccountInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ModifyAccountInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /AccountService/Accounts/{name}][%d] modifyAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *ModifyAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
