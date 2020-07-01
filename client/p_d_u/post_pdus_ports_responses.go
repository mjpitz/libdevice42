// Code generated by go-swagger; DO NOT EDIT.

package p_d_u

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostPdusPortsReader is a Reader for the PostPdusPorts structure.
type PostPdusPortsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPdusPortsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostPdusPortsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPdusPortsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostPdusPortsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostPdusPortsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostPdusPortsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostPdusPortsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostPdusPortsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPdusPortsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostPdusPortsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPdusPortsOK creates a PostPdusPortsOK with default headers values
func NewPostPdusPortsOK() *PostPdusPortsOK {
	return &PostPdusPortsOK{}
}

/*PostPdusPortsOK handles this case with default header values.

The above command returns results like this:
*/
type PostPdusPortsOK struct {
	Payload *PostPdusPortsOKBody
}

func (o *PostPdusPortsOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/pdus/ports/][%d] postPdusPortsOK  %+v", 200, o.Payload)
}

func (o *PostPdusPortsOK) GetPayload() *PostPdusPortsOKBody {
	return o.Payload
}

func (o *PostPdusPortsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostPdusPortsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPdusPortsBadRequest creates a PostPdusPortsBadRequest with default headers values
func NewPostPdusPortsBadRequest() *PostPdusPortsBadRequest {
	return &PostPdusPortsBadRequest{}
}

/*PostPdusPortsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostPdusPortsBadRequest struct {
}

func (o *PostPdusPortsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/pdus/ports/][%d] postPdusPortsBadRequest ", 400)
}

func (o *PostPdusPortsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPdusPortsUnauthorized creates a PostPdusPortsUnauthorized with default headers values
func NewPostPdusPortsUnauthorized() *PostPdusPortsUnauthorized {
	return &PostPdusPortsUnauthorized{}
}

/*PostPdusPortsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostPdusPortsUnauthorized struct {
}

func (o *PostPdusPortsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/pdus/ports/][%d] postPdusPortsUnauthorized ", 401)
}

func (o *PostPdusPortsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPdusPortsForbidden creates a PostPdusPortsForbidden with default headers values
func NewPostPdusPortsForbidden() *PostPdusPortsForbidden {
	return &PostPdusPortsForbidden{}
}

/*PostPdusPortsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostPdusPortsForbidden struct {
}

func (o *PostPdusPortsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/pdus/ports/][%d] postPdusPortsForbidden ", 403)
}

func (o *PostPdusPortsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPdusPortsNotFound creates a PostPdusPortsNotFound with default headers values
func NewPostPdusPortsNotFound() *PostPdusPortsNotFound {
	return &PostPdusPortsNotFound{}
}

/*PostPdusPortsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostPdusPortsNotFound struct {
}

func (o *PostPdusPortsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/pdus/ports/][%d] postPdusPortsNotFound ", 404)
}

func (o *PostPdusPortsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPdusPortsMethodNotAllowed creates a PostPdusPortsMethodNotAllowed with default headers values
func NewPostPdusPortsMethodNotAllowed() *PostPdusPortsMethodNotAllowed {
	return &PostPdusPortsMethodNotAllowed{}
}

/*PostPdusPortsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostPdusPortsMethodNotAllowed struct {
}

func (o *PostPdusPortsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/pdus/ports/][%d] postPdusPortsMethodNotAllowed ", 405)
}

func (o *PostPdusPortsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPdusPortsGone creates a PostPdusPortsGone with default headers values
func NewPostPdusPortsGone() *PostPdusPortsGone {
	return &PostPdusPortsGone{}
}

/*PostPdusPortsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostPdusPortsGone struct {
}

func (o *PostPdusPortsGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/pdus/ports/][%d] postPdusPortsGone ", 410)
}

func (o *PostPdusPortsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPdusPortsInternalServerError creates a PostPdusPortsInternalServerError with default headers values
func NewPostPdusPortsInternalServerError() *PostPdusPortsInternalServerError {
	return &PostPdusPortsInternalServerError{}
}

/*PostPdusPortsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type PostPdusPortsInternalServerError struct {
}

func (o *PostPdusPortsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/pdus/ports/][%d] postPdusPortsInternalServerError ", 500)
}

func (o *PostPdusPortsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPdusPortsServiceUnavailable creates a PostPdusPortsServiceUnavailable with default headers values
func NewPostPdusPortsServiceUnavailable() *PostPdusPortsServiceUnavailable {
	return &PostPdusPortsServiceUnavailable{}
}

/*PostPdusPortsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostPdusPortsServiceUnavailable struct {
}

func (o *PostPdusPortsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/pdus/ports/][%d] postPdusPortsServiceUnavailable ", 503)
}

func (o *PostPdusPortsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostPdusPortsOKBody post pdus ports o k body
swagger:model PostPdusPortsOKBody
*/
type PostPdusPortsOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post pdus ports o k body
func (o *PostPdusPortsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostPdusPortsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostPdusPortsOKBody) UnmarshalBinary(b []byte) error {
	var res PostPdusPortsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}