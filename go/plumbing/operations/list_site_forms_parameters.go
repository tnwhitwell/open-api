// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListSiteFormsParams creates a new ListSiteFormsParams object
// with the default values initialized.
func NewListSiteFormsParams() *ListSiteFormsParams {
	var ()
	return &ListSiteFormsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSiteFormsParamsWithTimeout creates a new ListSiteFormsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSiteFormsParamsWithTimeout(timeout time.Duration) *ListSiteFormsParams {
	var ()
	return &ListSiteFormsParams{

		timeout: timeout,
	}
}

// NewListSiteFormsParamsWithContext creates a new ListSiteFormsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSiteFormsParamsWithContext(ctx context.Context) *ListSiteFormsParams {
	var ()
	return &ListSiteFormsParams{

		Context: ctx,
	}
}

// NewListSiteFormsParamsWithHTTPClient creates a new ListSiteFormsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSiteFormsParamsWithHTTPClient(client *http.Client) *ListSiteFormsParams {
	var ()
	return &ListSiteFormsParams{
		HTTPClient: client,
	}
}

/*ListSiteFormsParams contains all the parameters to send to the API endpoint
for the list site forms operation typically these are written to a http.Request
*/
type ListSiteFormsParams struct {

	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list site forms params
func (o *ListSiteFormsParams) WithTimeout(timeout time.Duration) *ListSiteFormsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list site forms params
func (o *ListSiteFormsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list site forms params
func (o *ListSiteFormsParams) WithContext(ctx context.Context) *ListSiteFormsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list site forms params
func (o *ListSiteFormsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list site forms params
func (o *ListSiteFormsParams) WithHTTPClient(client *http.Client) *ListSiteFormsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list site forms params
func (o *ListSiteFormsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSiteID adds the siteID to the list site forms params
func (o *ListSiteFormsParams) WithSiteID(siteID string) *ListSiteFormsParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the list site forms params
func (o *ListSiteFormsParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *ListSiteFormsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
