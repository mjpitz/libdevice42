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

// PutIPAMCustomFIvrfgroupReader is a Reader for the PutIPAMCustomFIvrfgroup structure.
type PutIPAMCustomFIvrfgroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutIPAMCustomFIvrfgroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutIPAMCustomFIvrfgroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutIPAMCustomFIvrfgroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutIPAMCustomFIvrfgroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutIPAMCustomFIvrfgroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutIPAMCustomFIvrfgroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPutIPAMCustomFIvrfgroupMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPutIPAMCustomFIvrfgroupGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutIPAMCustomFIvrfgroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutIPAMCustomFIvrfgroupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutIPAMCustomFIvrfgroupOK creates a PutIPAMCustomFIvrfgroupOK with default headers values
func NewPutIPAMCustomFIvrfgroupOK() *PutIPAMCustomFIvrfgroupOK {
	return &PutIPAMCustomFIvrfgroupOK{}
}

/*PutIPAMCustomFIvrfgroupOK handles this case with default header values.

The above command returns results like this:
*/
type PutIPAMCustomFIvrfgroupOK struct {
	Payload *PutIPAMCustomFIvrfgroupOKBody
}

func (o *PutIPAMCustomFIvrfgroupOK) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/vrfgroup/][%d] putIpAMCustomFIvrfgroupOK  %+v", 200, o.Payload)
}

func (o *PutIPAMCustomFIvrfgroupOK) GetPayload() *PutIPAMCustomFIvrfgroupOKBody {
	return o.Payload
}

func (o *PutIPAMCustomFIvrfgroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutIPAMCustomFIvrfgroupOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIPAMCustomFIvrfgroupBadRequest creates a PutIPAMCustomFIvrfgroupBadRequest with default headers values
func NewPutIPAMCustomFIvrfgroupBadRequest() *PutIPAMCustomFIvrfgroupBadRequest {
	return &PutIPAMCustomFIvrfgroupBadRequest{}
}

/*PutIPAMCustomFIvrfgroupBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PutIPAMCustomFIvrfgroupBadRequest struct {
}

func (o *PutIPAMCustomFIvrfgroupBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/vrfgroup/][%d] putIpAMCustomFIvrfgroupBadRequest ", 400)
}

func (o *PutIPAMCustomFIvrfgroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutIPAMCustomFIvrfgroupUnauthorized creates a PutIPAMCustomFIvrfgroupUnauthorized with default headers values
func NewPutIPAMCustomFIvrfgroupUnauthorized() *PutIPAMCustomFIvrfgroupUnauthorized {
	return &PutIPAMCustomFIvrfgroupUnauthorized{}
}

/*PutIPAMCustomFIvrfgroupUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PutIPAMCustomFIvrfgroupUnauthorized struct {
}

func (o *PutIPAMCustomFIvrfgroupUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/vrfgroup/][%d] putIpAMCustomFIvrfgroupUnauthorized ", 401)
}

func (o *PutIPAMCustomFIvrfgroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutIPAMCustomFIvrfgroupForbidden creates a PutIPAMCustomFIvrfgroupForbidden with default headers values
func NewPutIPAMCustomFIvrfgroupForbidden() *PutIPAMCustomFIvrfgroupForbidden {
	return &PutIPAMCustomFIvrfgroupForbidden{}
}

/*PutIPAMCustomFIvrfgroupForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PutIPAMCustomFIvrfgroupForbidden struct {
}

func (o *PutIPAMCustomFIvrfgroupForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/vrfgroup/][%d] putIpAMCustomFIvrfgroupForbidden ", 403)
}

func (o *PutIPAMCustomFIvrfgroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutIPAMCustomFIvrfgroupNotFound creates a PutIPAMCustomFIvrfgroupNotFound with default headers values
func NewPutIPAMCustomFIvrfgroupNotFound() *PutIPAMCustomFIvrfgroupNotFound {
	return &PutIPAMCustomFIvrfgroupNotFound{}
}

/*PutIPAMCustomFIvrfgroupNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PutIPAMCustomFIvrfgroupNotFound struct {
}

func (o *PutIPAMCustomFIvrfgroupNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/vrfgroup/][%d] putIpAMCustomFIvrfgroupNotFound ", 404)
}

func (o *PutIPAMCustomFIvrfgroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutIPAMCustomFIvrfgroupMethodNotAllowed creates a PutIPAMCustomFIvrfgroupMethodNotAllowed with default headers values
func NewPutIPAMCustomFIvrfgroupMethodNotAllowed() *PutIPAMCustomFIvrfgroupMethodNotAllowed {
	return &PutIPAMCustomFIvrfgroupMethodNotAllowed{}
}

/*PutIPAMCustomFIvrfgroupMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PutIPAMCustomFIvrfgroupMethodNotAllowed struct {
}

func (o *PutIPAMCustomFIvrfgroupMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/vrfgroup/][%d] putIpAMCustomFIvrfgroupMethodNotAllowed ", 405)
}

func (o *PutIPAMCustomFIvrfgroupMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutIPAMCustomFIvrfgroupGone creates a PutIPAMCustomFIvrfgroupGone with default headers values
func NewPutIPAMCustomFIvrfgroupGone() *PutIPAMCustomFIvrfgroupGone {
	return &PutIPAMCustomFIvrfgroupGone{}
}

/*PutIPAMCustomFIvrfgroupGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PutIPAMCustomFIvrfgroupGone struct {
}

func (o *PutIPAMCustomFIvrfgroupGone) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/vrfgroup/][%d] putIpAMCustomFIvrfgroupGone ", 410)
}

func (o *PutIPAMCustomFIvrfgroupGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutIPAMCustomFIvrfgroupInternalServerError creates a PutIPAMCustomFIvrfgroupInternalServerError with default headers values
func NewPutIPAMCustomFIvrfgroupInternalServerError() *PutIPAMCustomFIvrfgroupInternalServerError {
	return &PutIPAMCustomFIvrfgroupInternalServerError{}
}

/*PutIPAMCustomFIvrfgroupInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type PutIPAMCustomFIvrfgroupInternalServerError struct {
}

func (o *PutIPAMCustomFIvrfgroupInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/vrfgroup/][%d] putIpAMCustomFIvrfgroupInternalServerError ", 500)
}

func (o *PutIPAMCustomFIvrfgroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutIPAMCustomFIvrfgroupServiceUnavailable creates a PutIPAMCustomFIvrfgroupServiceUnavailable with default headers values
func NewPutIPAMCustomFIvrfgroupServiceUnavailable() *PutIPAMCustomFIvrfgroupServiceUnavailable {
	return &PutIPAMCustomFIvrfgroupServiceUnavailable{}
}

/*PutIPAMCustomFIvrfgroupServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PutIPAMCustomFIvrfgroupServiceUnavailable struct {
}

func (o *PutIPAMCustomFIvrfgroupServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/vrfgroup/][%d] putIpAMCustomFIvrfgroupServiceUnavailable ", 503)
}

func (o *PutIPAMCustomFIvrfgroupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PutIPAMCustomFIvrfgroupOKBody put IP a m custom f ivrfgroup o k body
swagger:model PutIPAMCustomFIvrfgroupOKBody
*/
type PutIPAMCustomFIvrfgroupOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this put IP a m custom f ivrfgroup o k body
func (o *PutIPAMCustomFIvrfgroupOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutIPAMCustomFIvrfgroupOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutIPAMCustomFIvrfgroupOKBody) UnmarshalBinary(b []byte) error {
	var res PutIPAMCustomFIvrfgroupOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
