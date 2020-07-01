// Code generated by go-swagger; DO NOT EDIT.

package auto_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetAutoDiscoveryCertificateReader is a Reader for the GetAutoDiscoveryCertificate structure.
type GetAutoDiscoveryCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAutoDiscoveryCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAutoDiscoveryCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAutoDiscoveryCertificateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAutoDiscoveryCertificateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAutoDiscoveryCertificateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAutoDiscoveryCertificateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetAutoDiscoveryCertificateMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetAutoDiscoveryCertificateGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAutoDiscoveryCertificateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAutoDiscoveryCertificateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAutoDiscoveryCertificateOK creates a GetAutoDiscoveryCertificateOK with default headers values
func NewGetAutoDiscoveryCertificateOK() *GetAutoDiscoveryCertificateOK {
	return &GetAutoDiscoveryCertificateOK{}
}

/*GetAutoDiscoveryCertificateOK handles this case with default header values.

The above command returns results like this:
*/
type GetAutoDiscoveryCertificateOK struct {
	Payload *GetAutoDiscoveryCertificateOKBody
}

func (o *GetAutoDiscoveryCertificateOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/auto_discovery/certificate/][%d] getAutoDiscoveryCertificateOK  %+v", 200, o.Payload)
}

func (o *GetAutoDiscoveryCertificateOK) GetPayload() *GetAutoDiscoveryCertificateOKBody {
	return o.Payload
}

func (o *GetAutoDiscoveryCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAutoDiscoveryCertificateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAutoDiscoveryCertificateBadRequest creates a GetAutoDiscoveryCertificateBadRequest with default headers values
func NewGetAutoDiscoveryCertificateBadRequest() *GetAutoDiscoveryCertificateBadRequest {
	return &GetAutoDiscoveryCertificateBadRequest{}
}

/*GetAutoDiscoveryCertificateBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetAutoDiscoveryCertificateBadRequest struct {
}

func (o *GetAutoDiscoveryCertificateBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/auto_discovery/certificate/][%d] getAutoDiscoveryCertificateBadRequest ", 400)
}

func (o *GetAutoDiscoveryCertificateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAutoDiscoveryCertificateUnauthorized creates a GetAutoDiscoveryCertificateUnauthorized with default headers values
func NewGetAutoDiscoveryCertificateUnauthorized() *GetAutoDiscoveryCertificateUnauthorized {
	return &GetAutoDiscoveryCertificateUnauthorized{}
}

/*GetAutoDiscoveryCertificateUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetAutoDiscoveryCertificateUnauthorized struct {
}

func (o *GetAutoDiscoveryCertificateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/auto_discovery/certificate/][%d] getAutoDiscoveryCertificateUnauthorized ", 401)
}

func (o *GetAutoDiscoveryCertificateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAutoDiscoveryCertificateForbidden creates a GetAutoDiscoveryCertificateForbidden with default headers values
func NewGetAutoDiscoveryCertificateForbidden() *GetAutoDiscoveryCertificateForbidden {
	return &GetAutoDiscoveryCertificateForbidden{}
}

/*GetAutoDiscoveryCertificateForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetAutoDiscoveryCertificateForbidden struct {
}

func (o *GetAutoDiscoveryCertificateForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/auto_discovery/certificate/][%d] getAutoDiscoveryCertificateForbidden ", 403)
}

func (o *GetAutoDiscoveryCertificateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAutoDiscoveryCertificateNotFound creates a GetAutoDiscoveryCertificateNotFound with default headers values
func NewGetAutoDiscoveryCertificateNotFound() *GetAutoDiscoveryCertificateNotFound {
	return &GetAutoDiscoveryCertificateNotFound{}
}

/*GetAutoDiscoveryCertificateNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetAutoDiscoveryCertificateNotFound struct {
}

func (o *GetAutoDiscoveryCertificateNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/auto_discovery/certificate/][%d] getAutoDiscoveryCertificateNotFound ", 404)
}

func (o *GetAutoDiscoveryCertificateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAutoDiscoveryCertificateMethodNotAllowed creates a GetAutoDiscoveryCertificateMethodNotAllowed with default headers values
func NewGetAutoDiscoveryCertificateMethodNotAllowed() *GetAutoDiscoveryCertificateMethodNotAllowed {
	return &GetAutoDiscoveryCertificateMethodNotAllowed{}
}

/*GetAutoDiscoveryCertificateMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetAutoDiscoveryCertificateMethodNotAllowed struct {
}

func (o *GetAutoDiscoveryCertificateMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/auto_discovery/certificate/][%d] getAutoDiscoveryCertificateMethodNotAllowed ", 405)
}

func (o *GetAutoDiscoveryCertificateMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAutoDiscoveryCertificateGone creates a GetAutoDiscoveryCertificateGone with default headers values
func NewGetAutoDiscoveryCertificateGone() *GetAutoDiscoveryCertificateGone {
	return &GetAutoDiscoveryCertificateGone{}
}

/*GetAutoDiscoveryCertificateGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetAutoDiscoveryCertificateGone struct {
}

func (o *GetAutoDiscoveryCertificateGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/auto_discovery/certificate/][%d] getAutoDiscoveryCertificateGone ", 410)
}

func (o *GetAutoDiscoveryCertificateGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAutoDiscoveryCertificateInternalServerError creates a GetAutoDiscoveryCertificateInternalServerError with default headers values
func NewGetAutoDiscoveryCertificateInternalServerError() *GetAutoDiscoveryCertificateInternalServerError {
	return &GetAutoDiscoveryCertificateInternalServerError{}
}

/*GetAutoDiscoveryCertificateInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetAutoDiscoveryCertificateInternalServerError struct {
}

func (o *GetAutoDiscoveryCertificateInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/auto_discovery/certificate/][%d] getAutoDiscoveryCertificateInternalServerError ", 500)
}

func (o *GetAutoDiscoveryCertificateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAutoDiscoveryCertificateServiceUnavailable creates a GetAutoDiscoveryCertificateServiceUnavailable with default headers values
func NewGetAutoDiscoveryCertificateServiceUnavailable() *GetAutoDiscoveryCertificateServiceUnavailable {
	return &GetAutoDiscoveryCertificateServiceUnavailable{}
}

/*GetAutoDiscoveryCertificateServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetAutoDiscoveryCertificateServiceUnavailable struct {
}

func (o *GetAutoDiscoveryCertificateServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/auto_discovery/certificate/][%d] getAutoDiscoveryCertificateServiceUnavailable ", 503)
}

func (o *GetAutoDiscoveryCertificateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetAutoDiscoveryCertificateOKBody get auto discovery certificate o k body
swagger:model GetAutoDiscoveryCertificateOKBody
*/
type GetAutoDiscoveryCertificateOKBody struct {

	// jobs
	Jobs []*JobsItems0 `json:"jobs"`
}

// Validate validates this get auto discovery certificate o k body
func (o *GetAutoDiscoveryCertificateOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateJobs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAutoDiscoveryCertificateOKBody) validateJobs(formats strfmt.Registry) error {

	if swag.IsZero(o.Jobs) { // not required
		return nil
	}

	for i := 0; i < len(o.Jobs); i++ {
		if swag.IsZero(o.Jobs[i]) { // not required
			continue
		}

		if o.Jobs[i] != nil {
			if err := o.Jobs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getAutoDiscoveryCertificateOK" + "." + "jobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAutoDiscoveryCertificateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAutoDiscoveryCertificateOKBody) UnmarshalBinary(b []byte) error {
	var res GetAutoDiscoveryCertificateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*JobsItems0 jobs items0
swagger:model JobsItems0
*/
type JobsItems0 struct {

	// end ip address
	EndIPAddress interface{} `json:"end_ip_address,omitempty"`

	// follow chain
	FollowChain interface{} `json:"follow_chain,omitempty"`

	// job id
	JobID interface{} `json:"job_id,omitempty"`

	// name
	Name interface{} `json:"name,omitempty"`

	// ports
	Ports interface{} `json:"ports,omitempty"`

	// start ip address
	StartIPAddress interface{} `json:"start_ip_address,omitempty"`
}

// Validate validates this jobs items0
func (o *JobsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *JobsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *JobsItems0) UnmarshalBinary(b []byte) error {
	var res JobsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
