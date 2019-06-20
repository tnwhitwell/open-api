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

// EnableSplitTestReader is a Reader for the EnableSplitTest structure.
type EnableSplitTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableSplitTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewEnableSplitTestNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewEnableSplitTestDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEnableSplitTestNoContent creates a EnableSplitTestNoContent with default headers values
func NewEnableSplitTestNoContent() *EnableSplitTestNoContent {
	return &EnableSplitTestNoContent{}
}

/*EnableSplitTestNoContent handles this case with default header values.

enable
*/
type EnableSplitTestNoContent struct {
}

func (o *EnableSplitTestNoContent) Error() string {
	return fmt.Sprintf("[POST /site/{site_id}/traffic_splits/{split_test_id}/publish][%d] enableSplitTestNoContent ", 204)
}

func (o *EnableSplitTestNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableSplitTestDefault creates a EnableSplitTestDefault with default headers values
func NewEnableSplitTestDefault(code int) *EnableSplitTestDefault {
	return &EnableSplitTestDefault{
		_statusCode: code,
	}
}

/*EnableSplitTestDefault handles this case with default header values.

error
*/
type EnableSplitTestDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the enable split test default response
func (o *EnableSplitTestDefault) Code() int {
	return o._statusCode
}

func (o *EnableSplitTestDefault) Error() string {
	return fmt.Sprintf("[POST /site/{site_id}/traffic_splits/{split_test_id}/publish][%d] enableSplitTest default  %+v", o._statusCode, o.Payload)
}

func (o *EnableSplitTestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
