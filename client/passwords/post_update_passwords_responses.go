// Code generated by go-swagger; DO NOT EDIT.

package passwords

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostUpdatePasswordsReader is a Reader for the PostUpdatePasswords structure.
type PostUpdatePasswordsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUpdatePasswordsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUpdatePasswordsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUpdatePasswordsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUpdatePasswordsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUpdatePasswordsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUpdatePasswordsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostUpdatePasswordsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostUpdatePasswordsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUpdatePasswordsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUpdatePasswordsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostUpdatePasswordsOK creates a PostUpdatePasswordsOK with default headers values
func NewPostUpdatePasswordsOK() *PostUpdatePasswordsOK {
	return &PostUpdatePasswordsOK{}
}

/*PostUpdatePasswordsOK handles this case with default header values.

The above command returns results like this:
*/
type PostUpdatePasswordsOK struct {
	Payload *PostUpdatePasswordsOKBody
}

func (o *PostUpdatePasswordsOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/passwords/][%d] postUpdatePasswordsOK  %+v", 200, o.Payload)
}

func (o *PostUpdatePasswordsOK) GetPayload() *PostUpdatePasswordsOKBody {
	return o.Payload
}

func (o *PostUpdatePasswordsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostUpdatePasswordsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUpdatePasswordsBadRequest creates a PostUpdatePasswordsBadRequest with default headers values
func NewPostUpdatePasswordsBadRequest() *PostUpdatePasswordsBadRequest {
	return &PostUpdatePasswordsBadRequest{}
}

/*PostUpdatePasswordsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostUpdatePasswordsBadRequest struct {
}

func (o *PostUpdatePasswordsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/passwords/][%d] postUpdatePasswordsBadRequest ", 400)
}

func (o *PostUpdatePasswordsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdatePasswordsUnauthorized creates a PostUpdatePasswordsUnauthorized with default headers values
func NewPostUpdatePasswordsUnauthorized() *PostUpdatePasswordsUnauthorized {
	return &PostUpdatePasswordsUnauthorized{}
}

/*PostUpdatePasswordsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostUpdatePasswordsUnauthorized struct {
}

func (o *PostUpdatePasswordsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/passwords/][%d] postUpdatePasswordsUnauthorized ", 401)
}

func (o *PostUpdatePasswordsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdatePasswordsForbidden creates a PostUpdatePasswordsForbidden with default headers values
func NewPostUpdatePasswordsForbidden() *PostUpdatePasswordsForbidden {
	return &PostUpdatePasswordsForbidden{}
}

/*PostUpdatePasswordsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostUpdatePasswordsForbidden struct {
}

func (o *PostUpdatePasswordsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/passwords/][%d] postUpdatePasswordsForbidden ", 403)
}

func (o *PostUpdatePasswordsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdatePasswordsNotFound creates a PostUpdatePasswordsNotFound with default headers values
func NewPostUpdatePasswordsNotFound() *PostUpdatePasswordsNotFound {
	return &PostUpdatePasswordsNotFound{}
}

/*PostUpdatePasswordsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostUpdatePasswordsNotFound struct {
}

func (o *PostUpdatePasswordsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/passwords/][%d] postUpdatePasswordsNotFound ", 404)
}

func (o *PostUpdatePasswordsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdatePasswordsMethodNotAllowed creates a PostUpdatePasswordsMethodNotAllowed with default headers values
func NewPostUpdatePasswordsMethodNotAllowed() *PostUpdatePasswordsMethodNotAllowed {
	return &PostUpdatePasswordsMethodNotAllowed{}
}

/*PostUpdatePasswordsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostUpdatePasswordsMethodNotAllowed struct {
}

func (o *PostUpdatePasswordsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/passwords/][%d] postUpdatePasswordsMethodNotAllowed ", 405)
}

func (o *PostUpdatePasswordsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdatePasswordsGone creates a PostUpdatePasswordsGone with default headers values
func NewPostUpdatePasswordsGone() *PostUpdatePasswordsGone {
	return &PostUpdatePasswordsGone{}
}

/*PostUpdatePasswordsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostUpdatePasswordsGone struct {
}

func (o *PostUpdatePasswordsGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/passwords/][%d] postUpdatePasswordsGone ", 410)
}

func (o *PostUpdatePasswordsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdatePasswordsInternalServerError creates a PostUpdatePasswordsInternalServerError with default headers values
func NewPostUpdatePasswordsInternalServerError() *PostUpdatePasswordsInternalServerError {
	return &PostUpdatePasswordsInternalServerError{}
}

/*PostUpdatePasswordsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type PostUpdatePasswordsInternalServerError struct {
}

func (o *PostUpdatePasswordsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/passwords/][%d] postUpdatePasswordsInternalServerError ", 500)
}

func (o *PostUpdatePasswordsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdatePasswordsServiceUnavailable creates a PostUpdatePasswordsServiceUnavailable with default headers values
func NewPostUpdatePasswordsServiceUnavailable() *PostUpdatePasswordsServiceUnavailable {
	return &PostUpdatePasswordsServiceUnavailable{}
}

/*PostUpdatePasswordsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostUpdatePasswordsServiceUnavailable struct {
}

func (o *PostUpdatePasswordsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/passwords/][%d] postUpdatePasswordsServiceUnavailable ", 503)
}

func (o *PostUpdatePasswordsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostUpdatePasswordsOKBody post update passwords o k body
swagger:model PostUpdatePasswordsOKBody
*/
type PostUpdatePasswordsOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post update passwords o k body
func (o *PostUpdatePasswordsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostUpdatePasswordsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostUpdatePasswordsOKBody) UnmarshalBinary(b []byte) error {
	var res PostUpdatePasswordsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
