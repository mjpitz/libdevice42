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

// PostIPAMSubnetsCreateChildReader is a Reader for the PostIPAMSubnetsCreateChild structure.
type PostIPAMSubnetsCreateChildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIPAMSubnetsCreateChildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostIPAMSubnetsCreateChildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostIPAMSubnetsCreateChildBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostIPAMSubnetsCreateChildUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostIPAMSubnetsCreateChildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostIPAMSubnetsCreateChildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostIPAMSubnetsCreateChildMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostIPAMSubnetsCreateChildGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostIPAMSubnetsCreateChildInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostIPAMSubnetsCreateChildServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostIPAMSubnetsCreateChildOK creates a PostIPAMSubnetsCreateChildOK with default headers values
func NewPostIPAMSubnetsCreateChildOK() *PostIPAMSubnetsCreateChildOK {
	return &PostIPAMSubnetsCreateChildOK{}
}

/*PostIPAMSubnetsCreateChildOK handles this case with default header values.

The above command returns results like this:
*/
type PostIPAMSubnetsCreateChildOK struct {
	Payload *PostIPAMSubnetsCreateChildOKBody
}

func (o *PostIPAMSubnetsCreateChildOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/subnets/create_child/][%d] postIpAMSubnetsCreateChildOK  %+v", 200, o.Payload)
}

func (o *PostIPAMSubnetsCreateChildOK) GetPayload() *PostIPAMSubnetsCreateChildOKBody {
	return o.Payload
}

func (o *PostIPAMSubnetsCreateChildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostIPAMSubnetsCreateChildOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIPAMSubnetsCreateChildBadRequest creates a PostIPAMSubnetsCreateChildBadRequest with default headers values
func NewPostIPAMSubnetsCreateChildBadRequest() *PostIPAMSubnetsCreateChildBadRequest {
	return &PostIPAMSubnetsCreateChildBadRequest{}
}

/*PostIPAMSubnetsCreateChildBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostIPAMSubnetsCreateChildBadRequest struct {
}

func (o *PostIPAMSubnetsCreateChildBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/subnets/create_child/][%d] postIpAMSubnetsCreateChildBadRequest ", 400)
}

func (o *PostIPAMSubnetsCreateChildBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMSubnetsCreateChildUnauthorized creates a PostIPAMSubnetsCreateChildUnauthorized with default headers values
func NewPostIPAMSubnetsCreateChildUnauthorized() *PostIPAMSubnetsCreateChildUnauthorized {
	return &PostIPAMSubnetsCreateChildUnauthorized{}
}

/*PostIPAMSubnetsCreateChildUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostIPAMSubnetsCreateChildUnauthorized struct {
}

func (o *PostIPAMSubnetsCreateChildUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/subnets/create_child/][%d] postIpAMSubnetsCreateChildUnauthorized ", 401)
}

func (o *PostIPAMSubnetsCreateChildUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMSubnetsCreateChildForbidden creates a PostIPAMSubnetsCreateChildForbidden with default headers values
func NewPostIPAMSubnetsCreateChildForbidden() *PostIPAMSubnetsCreateChildForbidden {
	return &PostIPAMSubnetsCreateChildForbidden{}
}

/*PostIPAMSubnetsCreateChildForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostIPAMSubnetsCreateChildForbidden struct {
}

func (o *PostIPAMSubnetsCreateChildForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/subnets/create_child/][%d] postIpAMSubnetsCreateChildForbidden ", 403)
}

func (o *PostIPAMSubnetsCreateChildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMSubnetsCreateChildNotFound creates a PostIPAMSubnetsCreateChildNotFound with default headers values
func NewPostIPAMSubnetsCreateChildNotFound() *PostIPAMSubnetsCreateChildNotFound {
	return &PostIPAMSubnetsCreateChildNotFound{}
}

/*PostIPAMSubnetsCreateChildNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostIPAMSubnetsCreateChildNotFound struct {
}

func (o *PostIPAMSubnetsCreateChildNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/subnets/create_child/][%d] postIpAMSubnetsCreateChildNotFound ", 404)
}

func (o *PostIPAMSubnetsCreateChildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMSubnetsCreateChildMethodNotAllowed creates a PostIPAMSubnetsCreateChildMethodNotAllowed with default headers values
func NewPostIPAMSubnetsCreateChildMethodNotAllowed() *PostIPAMSubnetsCreateChildMethodNotAllowed {
	return &PostIPAMSubnetsCreateChildMethodNotAllowed{}
}

/*PostIPAMSubnetsCreateChildMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostIPAMSubnetsCreateChildMethodNotAllowed struct {
}

func (o *PostIPAMSubnetsCreateChildMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/subnets/create_child/][%d] postIpAMSubnetsCreateChildMethodNotAllowed ", 405)
}

func (o *PostIPAMSubnetsCreateChildMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMSubnetsCreateChildGone creates a PostIPAMSubnetsCreateChildGone with default headers values
func NewPostIPAMSubnetsCreateChildGone() *PostIPAMSubnetsCreateChildGone {
	return &PostIPAMSubnetsCreateChildGone{}
}

/*PostIPAMSubnetsCreateChildGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostIPAMSubnetsCreateChildGone struct {
}

func (o *PostIPAMSubnetsCreateChildGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/subnets/create_child/][%d] postIpAMSubnetsCreateChildGone ", 410)
}

func (o *PostIPAMSubnetsCreateChildGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMSubnetsCreateChildInternalServerError creates a PostIPAMSubnetsCreateChildInternalServerError with default headers values
func NewPostIPAMSubnetsCreateChildInternalServerError() *PostIPAMSubnetsCreateChildInternalServerError {
	return &PostIPAMSubnetsCreateChildInternalServerError{}
}

/*PostIPAMSubnetsCreateChildInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type PostIPAMSubnetsCreateChildInternalServerError struct {
}

func (o *PostIPAMSubnetsCreateChildInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/subnets/create_child/][%d] postIpAMSubnetsCreateChildInternalServerError ", 500)
}

func (o *PostIPAMSubnetsCreateChildInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMSubnetsCreateChildServiceUnavailable creates a PostIPAMSubnetsCreateChildServiceUnavailable with default headers values
func NewPostIPAMSubnetsCreateChildServiceUnavailable() *PostIPAMSubnetsCreateChildServiceUnavailable {
	return &PostIPAMSubnetsCreateChildServiceUnavailable{}
}

/*PostIPAMSubnetsCreateChildServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostIPAMSubnetsCreateChildServiceUnavailable struct {
}

func (o *PostIPAMSubnetsCreateChildServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/subnets/create_child/][%d] postIpAMSubnetsCreateChildServiceUnavailable ", 503)
}

func (o *PostIPAMSubnetsCreateChildServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostIPAMSubnetsCreateChildOKBody post IP a m subnets create child o k body
swagger:model PostIPAMSubnetsCreateChildOKBody
*/
type PostIPAMSubnetsCreateChildOKBody struct {

	// mask bits
	MaskBits interface{} `json:"mask_bits,omitempty"`

	// network
	Network interface{} `json:"network,omitempty"`

	// subnet id
	SubnetID interface{} `json:"subnet_id,omitempty"`
}

// Validate validates this post IP a m subnets create child o k body
func (o *PostIPAMSubnetsCreateChildOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostIPAMSubnetsCreateChildOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostIPAMSubnetsCreateChildOKBody) UnmarshalBinary(b []byte) error {
	var res PostIPAMSubnetsCreateChildOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
