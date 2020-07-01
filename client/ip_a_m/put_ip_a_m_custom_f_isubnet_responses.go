// Code generated by go-swagger; DO NOT EDIT.

package ip_a_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PutIPAMCustomFIsubnetReader is a Reader for the PutIPAMCustomFIsubnet structure.
type PutIPAMCustomFIsubnetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutIPAMCustomFIsubnetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutIPAMCustomFIsubnetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutIPAMCustomFIsubnetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutIPAMCustomFIsubnetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutIPAMCustomFIsubnetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutIPAMCustomFIsubnetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPutIPAMCustomFIsubnetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPutIPAMCustomFIsubnetGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutIPAMCustomFIsubnetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutIPAMCustomFIsubnetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutIPAMCustomFIsubnetOK creates a PutIPAMCustomFIsubnetOK with default headers values
func NewPutIPAMCustomFIsubnetOK() *PutIPAMCustomFIsubnetOK {
	return &PutIPAMCustomFIsubnetOK{}
}

/*PutIPAMCustomFIsubnetOK handles this case with default header values.

The above command returns results like this:
*/
type PutIPAMCustomFIsubnetOK struct {
	Payload *PutIPAMCustomFIsubnetOKBody
}

func (o *PutIPAMCustomFIsubnetOK) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/subnet/][%d] putIpAMCustomFIsubnetOK  %+v", 200, o.Payload)
}

func (o *PutIPAMCustomFIsubnetOK) GetPayload() *PutIPAMCustomFIsubnetOKBody {
	return o.Payload
}

func (o *PutIPAMCustomFIsubnetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutIPAMCustomFIsubnetOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIPAMCustomFIsubnetBadRequest creates a PutIPAMCustomFIsubnetBadRequest with default headers values
func NewPutIPAMCustomFIsubnetBadRequest() *PutIPAMCustomFIsubnetBadRequest {
	return &PutIPAMCustomFIsubnetBadRequest{}
}

/*PutIPAMCustomFIsubnetBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PutIPAMCustomFIsubnetBadRequest struct {
}

func (o *PutIPAMCustomFIsubnetBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/subnet/][%d] putIpAMCustomFIsubnetBadRequest ", 400)
}

func (o *PutIPAMCustomFIsubnetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutIPAMCustomFIsubnetUnauthorized creates a PutIPAMCustomFIsubnetUnauthorized with default headers values
func NewPutIPAMCustomFIsubnetUnauthorized() *PutIPAMCustomFIsubnetUnauthorized {
	return &PutIPAMCustomFIsubnetUnauthorized{}
}

/*PutIPAMCustomFIsubnetUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PutIPAMCustomFIsubnetUnauthorized struct {
}

func (o *PutIPAMCustomFIsubnetUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/subnet/][%d] putIpAMCustomFIsubnetUnauthorized ", 401)
}

func (o *PutIPAMCustomFIsubnetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutIPAMCustomFIsubnetForbidden creates a PutIPAMCustomFIsubnetForbidden with default headers values
func NewPutIPAMCustomFIsubnetForbidden() *PutIPAMCustomFIsubnetForbidden {
	return &PutIPAMCustomFIsubnetForbidden{}
}

/*PutIPAMCustomFIsubnetForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PutIPAMCustomFIsubnetForbidden struct {
}

func (o *PutIPAMCustomFIsubnetForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/subnet/][%d] putIpAMCustomFIsubnetForbidden ", 403)
}

func (o *PutIPAMCustomFIsubnetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutIPAMCustomFIsubnetNotFound creates a PutIPAMCustomFIsubnetNotFound with default headers values
func NewPutIPAMCustomFIsubnetNotFound() *PutIPAMCustomFIsubnetNotFound {
	return &PutIPAMCustomFIsubnetNotFound{}
}

/*PutIPAMCustomFIsubnetNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PutIPAMCustomFIsubnetNotFound struct {
}

func (o *PutIPAMCustomFIsubnetNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/subnet/][%d] putIpAMCustomFIsubnetNotFound ", 404)
}

func (o *PutIPAMCustomFIsubnetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutIPAMCustomFIsubnetMethodNotAllowed creates a PutIPAMCustomFIsubnetMethodNotAllowed with default headers values
func NewPutIPAMCustomFIsubnetMethodNotAllowed() *PutIPAMCustomFIsubnetMethodNotAllowed {
	return &PutIPAMCustomFIsubnetMethodNotAllowed{}
}

/*PutIPAMCustomFIsubnetMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PutIPAMCustomFIsubnetMethodNotAllowed struct {
}

func (o *PutIPAMCustomFIsubnetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/subnet/][%d] putIpAMCustomFIsubnetMethodNotAllowed ", 405)
}

func (o *PutIPAMCustomFIsubnetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutIPAMCustomFIsubnetGone creates a PutIPAMCustomFIsubnetGone with default headers values
func NewPutIPAMCustomFIsubnetGone() *PutIPAMCustomFIsubnetGone {
	return &PutIPAMCustomFIsubnetGone{}
}

/*PutIPAMCustomFIsubnetGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PutIPAMCustomFIsubnetGone struct {
}

func (o *PutIPAMCustomFIsubnetGone) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/subnet/][%d] putIpAMCustomFIsubnetGone ", 410)
}

func (o *PutIPAMCustomFIsubnetGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutIPAMCustomFIsubnetInternalServerError creates a PutIPAMCustomFIsubnetInternalServerError with default headers values
func NewPutIPAMCustomFIsubnetInternalServerError() *PutIPAMCustomFIsubnetInternalServerError {
	return &PutIPAMCustomFIsubnetInternalServerError{}
}

/*PutIPAMCustomFIsubnetInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type PutIPAMCustomFIsubnetInternalServerError struct {
}

func (o *PutIPAMCustomFIsubnetInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/subnet/][%d] putIpAMCustomFIsubnetInternalServerError ", 500)
}

func (o *PutIPAMCustomFIsubnetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutIPAMCustomFIsubnetServiceUnavailable creates a PutIPAMCustomFIsubnetServiceUnavailable with default headers values
func NewPutIPAMCustomFIsubnetServiceUnavailable() *PutIPAMCustomFIsubnetServiceUnavailable {
	return &PutIPAMCustomFIsubnetServiceUnavailable{}
}

/*PutIPAMCustomFIsubnetServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PutIPAMCustomFIsubnetServiceUnavailable struct {
}

func (o *PutIPAMCustomFIsubnetServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/subnet/][%d] putIpAMCustomFIsubnetServiceUnavailable ", 503)
}

func (o *PutIPAMCustomFIsubnetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PutIPAMCustomFIsubnetOKBody put IP a m custom f isubnet o k body
swagger:model PutIPAMCustomFIsubnetOKBody
*/
type PutIPAMCustomFIsubnetOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this put IP a m custom f isubnet o k body
func (o *PutIPAMCustomFIsubnetOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutIPAMCustomFIsubnetOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutIPAMCustomFIsubnetOKBody) UnmarshalBinary(b []byte) error {
	var res PutIPAMCustomFIsubnetOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
