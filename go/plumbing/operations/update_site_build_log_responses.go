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

// UpdateSiteBuildLogReader is a Reader for the UpdateSiteBuildLog structure.
type UpdateSiteBuildLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSiteBuildLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateSiteBuildLogNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateSiteBuildLogDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSiteBuildLogNoContent creates a UpdateSiteBuildLogNoContent with default headers values
func NewUpdateSiteBuildLogNoContent() *UpdateSiteBuildLogNoContent {
	return &UpdateSiteBuildLogNoContent{}
}

/*UpdateSiteBuildLogNoContent handles this case with default header values.

No content
*/
type UpdateSiteBuildLogNoContent struct {
}

func (o *UpdateSiteBuildLogNoContent) Error() string {
	return fmt.Sprintf("[POST /builds/{build_id}/log][%d] updateSiteBuildLogNoContent ", 204)
}

func (o *UpdateSiteBuildLogNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSiteBuildLogDefault creates a UpdateSiteBuildLogDefault with default headers values
func NewUpdateSiteBuildLogDefault(code int) *UpdateSiteBuildLogDefault {
	return &UpdateSiteBuildLogDefault{
		_statusCode: code,
	}
}

/*UpdateSiteBuildLogDefault handles this case with default header values.

error
*/
type UpdateSiteBuildLogDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update site build log default response
func (o *UpdateSiteBuildLogDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSiteBuildLogDefault) Error() string {
	return fmt.Sprintf("[POST /builds/{build_id}/log][%d] updateSiteBuildLog default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSiteBuildLogDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
