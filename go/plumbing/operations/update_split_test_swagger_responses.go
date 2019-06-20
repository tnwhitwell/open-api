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

// UpdateSplitTestReader is a Reader for the UpdateSplitTest structure.
type UpdateSplitTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSplitTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewUpdateSplitTestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateSplitTestDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSplitTestCreated creates a UpdateSplitTestCreated with default headers values
func NewUpdateSplitTestCreated() *UpdateSplitTestCreated {
	return &UpdateSplitTestCreated{}
}

/*UpdateSplitTestCreated handles this case with default header values.

Created
*/
type UpdateSplitTestCreated struct {
	Payload *models.SplitTest
}

func (o *UpdateSplitTestCreated) Error() string {
	return fmt.Sprintf("[PUT /site/{site_id}/traffic_splits/{split_test_id}][%d] updateSplitTestCreated  %+v", 201, o.Payload)
}

func (o *UpdateSplitTestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SplitTest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSplitTestDefault creates a UpdateSplitTestDefault with default headers values
func NewUpdateSplitTestDefault(code int) *UpdateSplitTestDefault {
	return &UpdateSplitTestDefault{
		_statusCode: code,
	}
}

/*UpdateSplitTestDefault handles this case with default header values.

error
*/
type UpdateSplitTestDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update split test default response
func (o *UpdateSplitTestDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSplitTestDefault) Error() string {
	return fmt.Sprintf("[PUT /site/{site_id}/traffic_splits/{split_test_id}][%d] updateSplitTest default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSplitTestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
