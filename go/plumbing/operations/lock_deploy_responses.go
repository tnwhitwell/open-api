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

// LockDeployReader is a Reader for the LockDeploy structure.
type LockDeployReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LockDeployReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLockDeployOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewLockDeployDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLockDeployOK creates a LockDeployOK with default headers values
func NewLockDeployOK() *LockDeployOK {
	return &LockDeployOK{}
}

/*LockDeployOK handles this case with default header values.

OK
*/
type LockDeployOK struct {
	Payload *models.Deploy
}

func (o *LockDeployOK) Error() string {
	return fmt.Sprintf("[POST /deploys/{deploy_id}/lock][%d] lockDeployOK  %+v", 200, o.Payload)
}

func (o *LockDeployOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Deploy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLockDeployDefault creates a LockDeployDefault with default headers values
func NewLockDeployDefault(code int) *LockDeployDefault {
	return &LockDeployDefault{
		_statusCode: code,
	}
}

/*LockDeployDefault handles this case with default header values.

error
*/
type LockDeployDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the lock deploy default response
func (o *LockDeployDefault) Code() int {
	return o._statusCode
}

func (o *LockDeployDefault) Error() string {
	return fmt.Sprintf("[POST /deploys/{deploy_id}/lock][%d] lockDeploy default  %+v", o._statusCode, o.Payload)
}

func (o *LockDeployDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
