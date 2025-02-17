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

// ListAccountsForUserReader is a Reader for the ListAccountsForUser structure.
type ListAccountsForUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAccountsForUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAccountsForUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListAccountsForUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAccountsForUserOK creates a ListAccountsForUserOK with default headers values
func NewListAccountsForUserOK() *ListAccountsForUserOK {
	return &ListAccountsForUserOK{}
}

/*ListAccountsForUserOK handles this case with default header values.

OK
*/
type ListAccountsForUserOK struct {
	Payload []*models.AccountMembership
}

func (o *ListAccountsForUserOK) Error() string {
	return fmt.Sprintf("[GET /accounts][%d] listAccountsForUserOK  %+v", 200, o.Payload)
}

func (o *ListAccountsForUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAccountsForUserDefault creates a ListAccountsForUserDefault with default headers values
func NewListAccountsForUserDefault(code int) *ListAccountsForUserDefault {
	return &ListAccountsForUserDefault{
		_statusCode: code,
	}
}

/*ListAccountsForUserDefault handles this case with default header values.

error
*/
type ListAccountsForUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list accounts for user default response
func (o *ListAccountsForUserDefault) Code() int {
	return o._statusCode
}

func (o *ListAccountsForUserDefault) Error() string {
	return fmt.Sprintf("[GET /accounts][%d] listAccountsForUser default  %+v", o._statusCode, o.Payload)
}

func (o *ListAccountsForUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
