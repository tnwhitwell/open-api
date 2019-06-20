// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/tnwhitwell/open-api/go/models"
)

// CreateTeamAccountReader is a Reader for the CreateTeamAccount structure.
type CreateTeamAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTeamAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateTeamAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateTeamAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateTeamAccountOK creates a CreateTeamAccountOK with default headers values
func NewCreateTeamAccountOK() *CreateTeamAccountOK {
	return &CreateTeamAccountOK{}
}

/*CreateTeamAccountOK handles this case with default header values.

Created
*/
type CreateTeamAccountOK struct {
	Payload *models.AccountMembership
}

func (o *CreateTeamAccountOK) Error() string {
	return fmt.Sprintf("[POST /accounts][%d] createTeamAccountOK  %+v", 200, o.Payload)
}

func (o *CreateTeamAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountMembership)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamAccountDefault creates a CreateTeamAccountDefault with default headers values
func NewCreateTeamAccountDefault(code int) *CreateTeamAccountDefault {
	return &CreateTeamAccountDefault{
		_statusCode: code,
	}
}

/*CreateTeamAccountDefault handles this case with default header values.

error
*/
type CreateTeamAccountDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create team account default response
func (o *CreateTeamAccountDefault) Code() int {
	return o._statusCode
}

func (o *CreateTeamAccountDefault) Error() string {
	return fmt.Sprintf("[POST /accounts][%d] createTeamAccount default  %+v", o._statusCode, o.Payload)
}

func (o *CreateTeamAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
