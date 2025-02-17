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

// ListAccountTypesForUserReader is a Reader for the ListAccountTypesForUser structure.
type ListAccountTypesForUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAccountTypesForUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAccountTypesForUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListAccountTypesForUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAccountTypesForUserOK creates a ListAccountTypesForUserOK with default headers values
func NewListAccountTypesForUserOK() *ListAccountTypesForUserOK {
	return &ListAccountTypesForUserOK{}
}

/*ListAccountTypesForUserOK handles this case with default header values.

OK
*/
type ListAccountTypesForUserOK struct {
	Payload []*models.AccountType
}

func (o *ListAccountTypesForUserOK) Error() string {
	return fmt.Sprintf("[GET /accounts/types][%d] listAccountTypesForUserOK  %+v", 200, o.Payload)
}

func (o *ListAccountTypesForUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAccountTypesForUserDefault creates a ListAccountTypesForUserDefault with default headers values
func NewListAccountTypesForUserDefault(code int) *ListAccountTypesForUserDefault {
	return &ListAccountTypesForUserDefault{
		_statusCode: code,
	}
}

/*ListAccountTypesForUserDefault handles this case with default header values.

error
*/
type ListAccountTypesForUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list account types for user default response
func (o *ListAccountTypesForUserDefault) Code() int {
	return o._statusCode
}

func (o *ListAccountTypesForUserDefault) Error() string {
	return fmt.Sprintf("[GET /accounts/types][%d] listAccountTypesForUser default  %+v", o._statusCode, o.Payload)
}

func (o *ListAccountTypesForUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
