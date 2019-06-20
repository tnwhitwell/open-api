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

// ShowTicketReader is a Reader for the ShowTicket structure.
type ShowTicketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowTicketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewShowTicketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewShowTicketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewShowTicketOK creates a ShowTicketOK with default headers values
func NewShowTicketOK() *ShowTicketOK {
	return &ShowTicketOK{}
}

/*ShowTicketOK handles this case with default header values.

ok
*/
type ShowTicketOK struct {
	Payload *models.Ticket
}

func (o *ShowTicketOK) Error() string {
	return fmt.Sprintf("[GET /oauth/tickets/{ticket_id}][%d] showTicketOK  %+v", 200, o.Payload)
}

func (o *ShowTicketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Ticket)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowTicketDefault creates a ShowTicketDefault with default headers values
func NewShowTicketDefault(code int) *ShowTicketDefault {
	return &ShowTicketDefault{
		_statusCode: code,
	}
}

/*ShowTicketDefault handles this case with default header values.

error
*/
type ShowTicketDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the show ticket default response
func (o *ShowTicketDefault) Code() int {
	return o._statusCode
}

func (o *ShowTicketDefault) Error() string {
	return fmt.Sprintf("[GET /oauth/tickets/{ticket_id}][%d] showTicket default  %+v", o._statusCode, o.Payload)
}

func (o *ShowTicketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
