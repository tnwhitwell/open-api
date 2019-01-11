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

// NewRestoreSiteDeployParams creates a new RestoreSiteDeployParams object
// with the default values initialized.
func NewRestoreSiteDeployParams() *RestoreSiteDeployParams {
	var ()
	return &RestoreSiteDeployParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRestoreSiteDeployParamsWithTimeout creates a new RestoreSiteDeployParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRestoreSiteDeployParamsWithTimeout(timeout time.Duration) *RestoreSiteDeployParams {
	var ()
	return &RestoreSiteDeployParams{

		timeout: timeout,
	}
}

// NewRestoreSiteDeployParamsWithContext creates a new RestoreSiteDeployParams object
// with the default values initialized, and the ability to set a context for a request
func NewRestoreSiteDeployParamsWithContext(ctx context.Context) *RestoreSiteDeployParams {
	var ()
	return &RestoreSiteDeployParams{

		Context: ctx,
	}
}

// NewRestoreSiteDeployParamsWithHTTPClient creates a new RestoreSiteDeployParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRestoreSiteDeployParamsWithHTTPClient(client *http.Client) *RestoreSiteDeployParams {
	var ()
	return &RestoreSiteDeployParams{
		HTTPClient: client,
	}
}

/*RestoreSiteDeployParams contains all the parameters to send to the API endpoint
for the restore site deploy operation typically these are written to a http.Request
*/
type RestoreSiteDeployParams struct {

	/*DeployID*/
	DeployID string
	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the restore site deploy params
func (o *RestoreSiteDeployParams) WithTimeout(timeout time.Duration) *RestoreSiteDeployParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restore site deploy params
func (o *RestoreSiteDeployParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restore site deploy params
func (o *RestoreSiteDeployParams) WithContext(ctx context.Context) *RestoreSiteDeployParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restore site deploy params
func (o *RestoreSiteDeployParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restore site deploy params
func (o *RestoreSiteDeployParams) WithHTTPClient(client *http.Client) *RestoreSiteDeployParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restore site deploy params
func (o *RestoreSiteDeployParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeployID adds the deployID to the restore site deploy params
func (o *RestoreSiteDeployParams) WithDeployID(deployID string) *RestoreSiteDeployParams {
	o.SetDeployID(deployID)
	return o
}

// SetDeployID adds the deployId to the restore site deploy params
func (o *RestoreSiteDeployParams) SetDeployID(deployID string) {
	o.DeployID = deployID
}

// WithSiteID adds the siteID to the restore site deploy params
func (o *RestoreSiteDeployParams) WithSiteID(siteID string) *RestoreSiteDeployParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the restore site deploy params
func (o *RestoreSiteDeployParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *RestoreSiteDeployParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deploy_id
	if err := r.SetPathParam("deploy_id", o.DeployID); err != nil {
		return err
	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
