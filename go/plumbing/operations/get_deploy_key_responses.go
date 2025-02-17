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

// GetDeployKeyReader is a Reader for the GetDeployKey structure.
type GetDeployKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeployKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDeployKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDeployKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeployKeyOK creates a GetDeployKeyOK with default headers values
func NewGetDeployKeyOK() *GetDeployKeyOK {
	return &GetDeployKeyOK{}
}

/*GetDeployKeyOK handles this case with default header values.

OK
*/
type GetDeployKeyOK struct {
	Payload *models.DeployKey
}

func (o *GetDeployKeyOK) Error() string {
	return fmt.Sprintf("[GET /deploy_keys/{key_id}][%d] getDeployKeyOK  %+v", 200, o.Payload)
}

func (o *GetDeployKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeployKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeployKeyDefault creates a GetDeployKeyDefault with default headers values
func NewGetDeployKeyDefault(code int) *GetDeployKeyDefault {
	return &GetDeployKeyDefault{
		_statusCode: code,
	}
}

/*GetDeployKeyDefault handles this case with default header values.

error
*/
type GetDeployKeyDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get deploy key default response
func (o *GetDeployKeyDefault) Code() int {
	return o._statusCode
}

func (o *GetDeployKeyDefault) Error() string {
	return fmt.Sprintf("[GET /deploy_keys/{key_id}][%d] getDeployKey default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeployKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
