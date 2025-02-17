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

// ListFormSubmissionReader is a Reader for the ListFormSubmission structure.
type ListFormSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListFormSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListFormSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListFormSubmissionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListFormSubmissionOK creates a ListFormSubmissionOK with default headers values
func NewListFormSubmissionOK() *ListFormSubmissionOK {
	return &ListFormSubmissionOK{}
}

/*ListFormSubmissionOK handles this case with default header values.

OK
*/
type ListFormSubmissionOK struct {
	Payload []*models.Submission
}

func (o *ListFormSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /submissions/{submission_id}][%d] listFormSubmissionOK  %+v", 200, o.Payload)
}

func (o *ListFormSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFormSubmissionDefault creates a ListFormSubmissionDefault with default headers values
func NewListFormSubmissionDefault(code int) *ListFormSubmissionDefault {
	return &ListFormSubmissionDefault{
		_statusCode: code,
	}
}

/*ListFormSubmissionDefault handles this case with default header values.

error
*/
type ListFormSubmissionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list form submission default response
func (o *ListFormSubmissionDefault) Code() int {
	return o._statusCode
}

func (o *ListFormSubmissionDefault) Error() string {
	return fmt.Sprintf("[GET /submissions/{submission_id}][%d] listFormSubmission default  %+v", o._statusCode, o.Payload)
}

func (o *ListFormSubmissionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
