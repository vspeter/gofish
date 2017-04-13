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

// CreateAccountReader is a Reader for the CreateAccount structure.
type CreateAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAccountCreated creates a CreateAccountCreated with default headers values
func NewCreateAccountCreated() *CreateAccountCreated {
	return &CreateAccountCreated{}
}

/*CreateAccountCreated handles this case with default header values.

Success
*/
type CreateAccountCreated struct {
	Payload *models.ErrorResponse
}

func (o *CreateAccountCreated) Error() string {
	return fmt.Sprintf("[POST /AccountService/Accounts][%d] createAccountCreated  %+v", 201, o.Payload)
}

func (o *CreateAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountBadRequest creates a CreateAccountBadRequest with default headers values
func NewCreateAccountBadRequest() *CreateAccountBadRequest {
	return &CreateAccountBadRequest{}
}

/*CreateAccountBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information (such as validation error on an input field, a missing required value, and so on). An extended error shall be returned in the response body, as defined in section Extended Error Handling.

*/
type CreateAccountBadRequest struct {
}

func (o *CreateAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /AccountService/Accounts][%d] createAccountBadRequest ", 400)
}

func (o *CreateAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccountUnauthorized creates a CreateAccountUnauthorized with default headers values
func NewCreateAccountUnauthorized() *CreateAccountUnauthorized {
	return &CreateAccountUnauthorized{}
}

/*CreateAccountUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type CreateAccountUnauthorized struct {
}

func (o *CreateAccountUnauthorized) Error() string {
	return fmt.Sprintf("[POST /AccountService/Accounts][%d] createAccountUnauthorized ", 401)
}

func (o *CreateAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccountForbidden creates a CreateAccountForbidden with default headers values
func NewCreateAccountForbidden() *CreateAccountForbidden {
	return &CreateAccountForbidden{}
}

/*CreateAccountForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type CreateAccountForbidden struct {
}

func (o *CreateAccountForbidden) Error() string {
	return fmt.Sprintf("[POST /AccountService/Accounts][%d] createAccountForbidden ", 403)
}

func (o *CreateAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccountNotFound creates a CreateAccountNotFound with default headers values
func NewCreateAccountNotFound() *CreateAccountNotFound {
	return &CreateAccountNotFound{}
}

/*CreateAccountNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type CreateAccountNotFound struct {
}

func (o *CreateAccountNotFound) Error() string {
	return fmt.Sprintf("[POST /AccountService/Accounts][%d] createAccountNotFound ", 404)
}

func (o *CreateAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccountInternalServerError creates a CreateAccountInternalServerError with default headers values
func NewCreateAccountInternalServerError() *CreateAccountInternalServerError {
	return &CreateAccountInternalServerError{}
}

/*CreateAccountInternalServerError handles this case with default header values.

Error
*/
type CreateAccountInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CreateAccountInternalServerError) Error() string {
	return fmt.Sprintf("[POST /AccountService/Accounts][%d] createAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
