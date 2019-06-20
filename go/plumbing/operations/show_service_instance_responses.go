// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tnwhitwell/open-api/go/models"
)

// ShowServiceInstanceReader is a Reader for the ShowServiceInstance structure.
type ShowServiceInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowServiceInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewShowServiceInstanceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewShowServiceInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewShowServiceInstanceCreated creates a ShowServiceInstanceCreated with default headers values
func NewShowServiceInstanceCreated() *ShowServiceInstanceCreated {
	return &ShowServiceInstanceCreated{}
}

/*ShowServiceInstanceCreated handles this case with default header values.

Created
*/
type ShowServiceInstanceCreated struct {
	Payload *models.ServiceInstance
}

func (o *ShowServiceInstanceCreated) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/services/{addon}/instances][%d] showServiceInstanceCreated  %+v", 201, o.Payload)
}

func (o *ShowServiceInstanceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowServiceInstanceDefault creates a ShowServiceInstanceDefault with default headers values
func NewShowServiceInstanceDefault(code int) *ShowServiceInstanceDefault {
	return &ShowServiceInstanceDefault{
		_statusCode: code,
	}
}

/*ShowServiceInstanceDefault handles this case with default header values.

error
*/
type ShowServiceInstanceDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the show service instance default response
func (o *ShowServiceInstanceDefault) Code() int {
	return o._statusCode
}

func (o *ShowServiceInstanceDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/services/{addon}/instances][%d] showServiceInstance default  %+v", o._statusCode, o.Payload)
}

func (o *ShowServiceInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
