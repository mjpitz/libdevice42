// Code generated by go-swagger; DO NOT EDIT.

package auto_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPutAutoDiscoveryNetworksParams creates a new PutAutoDiscoveryNetworksParams object
// with the default values initialized.
func NewPutAutoDiscoveryNetworksParams() *PutAutoDiscoveryNetworksParams {
	var ()
	return &PutAutoDiscoveryNetworksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAutoDiscoveryNetworksParamsWithTimeout creates a new PutAutoDiscoveryNetworksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAutoDiscoveryNetworksParamsWithTimeout(timeout time.Duration) *PutAutoDiscoveryNetworksParams {
	var ()
	return &PutAutoDiscoveryNetworksParams{

		timeout: timeout,
	}
}

// NewPutAutoDiscoveryNetworksParamsWithContext creates a new PutAutoDiscoveryNetworksParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAutoDiscoveryNetworksParamsWithContext(ctx context.Context) *PutAutoDiscoveryNetworksParams {
	var ()
	return &PutAutoDiscoveryNetworksParams{

		Context: ctx,
	}
}

// NewPutAutoDiscoveryNetworksParamsWithHTTPClient creates a new PutAutoDiscoveryNetworksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAutoDiscoveryNetworksParamsWithHTTPClient(client *http.Client) *PutAutoDiscoveryNetworksParams {
	var ()
	return &PutAutoDiscoveryNetworksParams{
		HTTPClient: client,
	}
}

/*PutAutoDiscoveryNetworksParams contains all the parameters to send to the API endpoint
for the put auto discovery networks operation typically these are written to a http.Request
*/
type PutAutoDiscoveryNetworksParams struct {

	/*JobID
	  D42 ID for the job - required if no name

	*/
	JobID *string
	/*Name
	  name of the job - required if no job_id

	*/
	Name *string
	/*Run
	  yes to start

	*/
	Run string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put auto discovery networks params
func (o *PutAutoDiscoveryNetworksParams) WithTimeout(timeout time.Duration) *PutAutoDiscoveryNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put auto discovery networks params
func (o *PutAutoDiscoveryNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put auto discovery networks params
func (o *PutAutoDiscoveryNetworksParams) WithContext(ctx context.Context) *PutAutoDiscoveryNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put auto discovery networks params
func (o *PutAutoDiscoveryNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put auto discovery networks params
func (o *PutAutoDiscoveryNetworksParams) WithHTTPClient(client *http.Client) *PutAutoDiscoveryNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put auto discovery networks params
func (o *PutAutoDiscoveryNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the put auto discovery networks params
func (o *PutAutoDiscoveryNetworksParams) WithJobID(jobID *string) *PutAutoDiscoveryNetworksParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the put auto discovery networks params
func (o *PutAutoDiscoveryNetworksParams) SetJobID(jobID *string) {
	o.JobID = jobID
}

// WithName adds the name to the put auto discovery networks params
func (o *PutAutoDiscoveryNetworksParams) WithName(name *string) *PutAutoDiscoveryNetworksParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put auto discovery networks params
func (o *PutAutoDiscoveryNetworksParams) SetName(name *string) {
	o.Name = name
}

// WithRun adds the run to the put auto discovery networks params
func (o *PutAutoDiscoveryNetworksParams) WithRun(run string) *PutAutoDiscoveryNetworksParams {
	o.SetRun(run)
	return o
}

// SetRun adds the run to the put auto discovery networks params
func (o *PutAutoDiscoveryNetworksParams) SetRun(run string) {
	o.Run = run
}

// WriteToRequest writes these params to a swagger request
func (o *PutAutoDiscoveryNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JobID != nil {

		// form param job_id
		var frJobID string
		if o.JobID != nil {
			frJobID = *o.JobID
		}
		fJobID := frJobID
		if fJobID != "" {
			if err := r.SetFormParam("job_id", fJobID); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// form param name
		var frName string
		if o.Name != nil {
			frName = *o.Name
		}
		fName := frName
		if fName != "" {
			if err := r.SetFormParam("name", fName); err != nil {
				return err
			}
		}

	}

	// form param run
	frRun := o.Run
	fRun := frRun
	if fRun != "" {
		if err := r.SetFormParam("run", fRun); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}