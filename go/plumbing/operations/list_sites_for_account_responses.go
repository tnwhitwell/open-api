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

// ListSitesForAccountReader is a Reader for the ListSitesForAccount structure.
type ListSitesForAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSitesForAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSitesForAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListSitesForAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSitesForAccountOK creates a ListSitesForAccountOK with default headers values
func NewListSitesForAccountOK() *ListSitesForAccountOK {
	return &ListSitesForAccountOK{}
}

/*ListSitesForAccountOK handles this case with default header values.

OK
*/
type ListSitesForAccountOK struct {
	Payload []*models.Site
}

func (o *ListSitesForAccountOK) Error() string {
	return fmt.Sprintf("[GET /{account_slug}/sites][%d] listSitesForAccountOK  %+v", 200, o.Payload)
}

func (o *ListSitesForAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSitesForAccountDefault creates a ListSitesForAccountDefault with default headers values
func NewListSitesForAccountDefault(code int) *ListSitesForAccountDefault {
	return &ListSitesForAccountDefault{
		_statusCode: code,
	}
}

/*ListSitesForAccountDefault handles this case with default header values.

error
*/
type ListSitesForAccountDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list sites for account default response
func (o *ListSitesForAccountDefault) Code() int {
	return o._statusCode
}

func (o *ListSitesForAccountDefault) Error() string {
	return fmt.Sprintf("[GET /{account_slug}/sites][%d] listSitesForAccount default  %+v", o._statusCode, o.Payload)
}

func (o *ListSitesForAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
