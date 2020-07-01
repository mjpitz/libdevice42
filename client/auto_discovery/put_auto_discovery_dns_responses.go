// Code generated by go-swagger; DO NOT EDIT.

package auto_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PutAutoDiscoveryDNSReader is a Reader for the PutAutoDiscoveryDNS structure.
type PutAutoDiscoveryDNSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAutoDiscoveryDNSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAutoDiscoveryDNSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAutoDiscoveryDNSBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutAutoDiscoveryDNSUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutAutoDiscoveryDNSForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAutoDiscoveryDNSNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPutAutoDiscoveryDNSMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPutAutoDiscoveryDNSGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutAutoDiscoveryDNSInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutAutoDiscoveryDNSServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutAutoDiscoveryDNSOK creates a PutAutoDiscoveryDNSOK with default headers values
func NewPutAutoDiscoveryDNSOK() *PutAutoDiscoveryDNSOK {
	return &PutAutoDiscoveryDNSOK{}
}

/*PutAutoDiscoveryDNSOK handles this case with default header values.

The above command returns results like this:
*/
type PutAutoDiscoveryDNSOK struct {
	Payload *PutAutoDiscoveryDNSOKBody
}

func (o *PutAutoDiscoveryDNSOK) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/auto_discovery/dns/][%d] putAutoDiscoveryDnsOK  %+v", 200, o.Payload)
}

func (o *PutAutoDiscoveryDNSOK) GetPayload() *PutAutoDiscoveryDNSOKBody {
	return o.Payload
}

func (o *PutAutoDiscoveryDNSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutAutoDiscoveryDNSOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAutoDiscoveryDNSBadRequest creates a PutAutoDiscoveryDNSBadRequest with default headers values
func NewPutAutoDiscoveryDNSBadRequest() *PutAutoDiscoveryDNSBadRequest {
	return &PutAutoDiscoveryDNSBadRequest{}
}

/*PutAutoDiscoveryDNSBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PutAutoDiscoveryDNSBadRequest struct {
}

func (o *PutAutoDiscoveryDNSBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/auto_discovery/dns/][%d] putAutoDiscoveryDnsBadRequest ", 400)
}

func (o *PutAutoDiscoveryDNSBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAutoDiscoveryDNSUnauthorized creates a PutAutoDiscoveryDNSUnauthorized with default headers values
func NewPutAutoDiscoveryDNSUnauthorized() *PutAutoDiscoveryDNSUnauthorized {
	return &PutAutoDiscoveryDNSUnauthorized{}
}

/*PutAutoDiscoveryDNSUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PutAutoDiscoveryDNSUnauthorized struct {
}

func (o *PutAutoDiscoveryDNSUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/auto_discovery/dns/][%d] putAutoDiscoveryDnsUnauthorized ", 401)
}

func (o *PutAutoDiscoveryDNSUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAutoDiscoveryDNSForbidden creates a PutAutoDiscoveryDNSForbidden with default headers values
func NewPutAutoDiscoveryDNSForbidden() *PutAutoDiscoveryDNSForbidden {
	return &PutAutoDiscoveryDNSForbidden{}
}

/*PutAutoDiscoveryDNSForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PutAutoDiscoveryDNSForbidden struct {
}

func (o *PutAutoDiscoveryDNSForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/auto_discovery/dns/][%d] putAutoDiscoveryDnsForbidden ", 403)
}

func (o *PutAutoDiscoveryDNSForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAutoDiscoveryDNSNotFound creates a PutAutoDiscoveryDNSNotFound with default headers values
func NewPutAutoDiscoveryDNSNotFound() *PutAutoDiscoveryDNSNotFound {
	return &PutAutoDiscoveryDNSNotFound{}
}

/*PutAutoDiscoveryDNSNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PutAutoDiscoveryDNSNotFound struct {
}

func (o *PutAutoDiscoveryDNSNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/auto_discovery/dns/][%d] putAutoDiscoveryDnsNotFound ", 404)
}

func (o *PutAutoDiscoveryDNSNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAutoDiscoveryDNSMethodNotAllowed creates a PutAutoDiscoveryDNSMethodNotAllowed with default headers values
func NewPutAutoDiscoveryDNSMethodNotAllowed() *PutAutoDiscoveryDNSMethodNotAllowed {
	return &PutAutoDiscoveryDNSMethodNotAllowed{}
}

/*PutAutoDiscoveryDNSMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PutAutoDiscoveryDNSMethodNotAllowed struct {
}

func (o *PutAutoDiscoveryDNSMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/auto_discovery/dns/][%d] putAutoDiscoveryDnsMethodNotAllowed ", 405)
}

func (o *PutAutoDiscoveryDNSMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAutoDiscoveryDNSGone creates a PutAutoDiscoveryDNSGone with default headers values
func NewPutAutoDiscoveryDNSGone() *PutAutoDiscoveryDNSGone {
	return &PutAutoDiscoveryDNSGone{}
}

/*PutAutoDiscoveryDNSGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PutAutoDiscoveryDNSGone struct {
}

func (o *PutAutoDiscoveryDNSGone) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/auto_discovery/dns/][%d] putAutoDiscoveryDnsGone ", 410)
}

func (o *PutAutoDiscoveryDNSGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAutoDiscoveryDNSInternalServerError creates a PutAutoDiscoveryDNSInternalServerError with default headers values
func NewPutAutoDiscoveryDNSInternalServerError() *PutAutoDiscoveryDNSInternalServerError {
	return &PutAutoDiscoveryDNSInternalServerError{}
}

/*PutAutoDiscoveryDNSInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type PutAutoDiscoveryDNSInternalServerError struct {
}

func (o *PutAutoDiscoveryDNSInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/auto_discovery/dns/][%d] putAutoDiscoveryDnsInternalServerError ", 500)
}

func (o *PutAutoDiscoveryDNSInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAutoDiscoveryDNSServiceUnavailable creates a PutAutoDiscoveryDNSServiceUnavailable with default headers values
func NewPutAutoDiscoveryDNSServiceUnavailable() *PutAutoDiscoveryDNSServiceUnavailable {
	return &PutAutoDiscoveryDNSServiceUnavailable{}
}

/*PutAutoDiscoveryDNSServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PutAutoDiscoveryDNSServiceUnavailable struct {
}

func (o *PutAutoDiscoveryDNSServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/auto_discovery/dns/][%d] putAutoDiscoveryDnsServiceUnavailable ", 503)
}

func (o *PutAutoDiscoveryDNSServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PutAutoDiscoveryDNSOKBody put auto discovery DNS o k body
swagger:model PutAutoDiscoveryDNSOKBody
*/
type PutAutoDiscoveryDNSOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this put auto discovery DNS o k body
func (o *PutAutoDiscoveryDNSOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutAutoDiscoveryDNSOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutAutoDiscoveryDNSOKBody) UnmarshalBinary(b []byte) error {
	var res PutAutoDiscoveryDNSOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}