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

// UpdateSiteMetadataReader is a Reader for the UpdateSiteMetadata structure.
type UpdateSiteMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSiteMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateSiteMetadataNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateSiteMetadataDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSiteMetadataNoContent creates a UpdateSiteMetadataNoContent with default headers values
func NewUpdateSiteMetadataNoContent() *UpdateSiteMetadataNoContent {
	return &UpdateSiteMetadataNoContent{}
}

/*UpdateSiteMetadataNoContent handles this case with default header values.

No content
*/
type UpdateSiteMetadataNoContent struct {
}

func (o *UpdateSiteMetadataNoContent) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/metadata][%d] updateSiteMetadataNoContent ", 204)
}

func (o *UpdateSiteMetadataNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSiteMetadataDefault creates a UpdateSiteMetadataDefault with default headers values
func NewUpdateSiteMetadataDefault(code int) *UpdateSiteMetadataDefault {
	return &UpdateSiteMetadataDefault{
		_statusCode: code,
	}
}

/*UpdateSiteMetadataDefault handles this case with default header values.

error
*/
type UpdateSiteMetadataDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update site metadata default response
func (o *UpdateSiteMetadataDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSiteMetadataDefault) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/metadata][%d] updateSiteMetadata default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSiteMetadataDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
