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

// GetSchemaReader is a Reader for the GetSchema structure.
type GetSchemaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSchemaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSchemaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetSchemaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSchemaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSchemaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSchemaOK creates a GetSchemaOK with default headers values
func NewGetSchemaOK() *GetSchemaOK {
	return &GetSchemaOK{}
}

/*GetSchemaOK handles this case with default header values.

Success
*/
type GetSchemaOK struct {
	Payload *models.JSONSchemaFile100JSONSchemaFile
}

func (o *GetSchemaOK) Error() string {
	return fmt.Sprintf("[GET /Schemas/{identifier}][%d] getSchemaOK  %+v", 200, o.Payload)
}

func (o *GetSchemaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONSchemaFile100JSONSchemaFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemaUnauthorized creates a GetSchemaUnauthorized with default headers values
func NewGetSchemaUnauthorized() *GetSchemaUnauthorized {
	return &GetSchemaUnauthorized{}
}

/*GetSchemaUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetSchemaUnauthorized struct {
}

func (o *GetSchemaUnauthorized) Error() string {
	return fmt.Sprintf("[GET /Schemas/{identifier}][%d] getSchemaUnauthorized ", 401)
}

func (o *GetSchemaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSchemaForbidden creates a GetSchemaForbidden with default headers values
func NewGetSchemaForbidden() *GetSchemaForbidden {
	return &GetSchemaForbidden{}
}

/*GetSchemaForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type GetSchemaForbidden struct {
}

func (o *GetSchemaForbidden) Error() string {
	return fmt.Sprintf("[GET /Schemas/{identifier}][%d] getSchemaForbidden ", 403)
}

func (o *GetSchemaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSchemaInternalServerError creates a GetSchemaInternalServerError with default headers values
func NewGetSchemaInternalServerError() *GetSchemaInternalServerError {
	return &GetSchemaInternalServerError{}
}

/*GetSchemaInternalServerError handles this case with default header values.

Error
*/
type GetSchemaInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetSchemaInternalServerError) Error() string {
	return fmt.Sprintf("[GET /Schemas/{identifier}][%d] getSchemaInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSchemaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
