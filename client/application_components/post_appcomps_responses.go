// Code generated by go-swagger; DO NOT EDIT.

package application_components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostAppcompsReader is a Reader for the PostAppcomps structure.
type PostAppcompsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAppcompsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAppcompsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAppcompsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAppcompsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAppcompsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAppcompsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostAppcompsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostAppcompsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAppcompsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAppcompsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAppcompsOK creates a PostAppcompsOK with default headers values
func NewPostAppcompsOK() *PostAppcompsOK {
	return &PostAppcompsOK{}
}

/*PostAppcompsOK handles this case with default header values.

The above command returns results like this:
*/
type PostAppcompsOK struct {
	Payload *PostAppcompsOKBody
}

func (o *PostAppcompsOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/appcomps/][%d] postAppcompsOK  %+v", 200, o.Payload)
}

func (o *PostAppcompsOK) GetPayload() *PostAppcompsOKBody {
	return o.Payload
}

func (o *PostAppcompsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostAppcompsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAppcompsBadRequest creates a PostAppcompsBadRequest with default headers values
func NewPostAppcompsBadRequest() *PostAppcompsBadRequest {
	return &PostAppcompsBadRequest{}
}

/*PostAppcompsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostAppcompsBadRequest struct {
}

func (o *PostAppcompsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/appcomps/][%d] postAppcompsBadRequest ", 400)
}

func (o *PostAppcompsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAppcompsUnauthorized creates a PostAppcompsUnauthorized with default headers values
func NewPostAppcompsUnauthorized() *PostAppcompsUnauthorized {
	return &PostAppcompsUnauthorized{}
}

/*PostAppcompsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostAppcompsUnauthorized struct {
}

func (o *PostAppcompsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/appcomps/][%d] postAppcompsUnauthorized ", 401)
}

func (o *PostAppcompsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAppcompsForbidden creates a PostAppcompsForbidden with default headers values
func NewPostAppcompsForbidden() *PostAppcompsForbidden {
	return &PostAppcompsForbidden{}
}

/*PostAppcompsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostAppcompsForbidden struct {
}

func (o *PostAppcompsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/appcomps/][%d] postAppcompsForbidden ", 403)
}

func (o *PostAppcompsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAppcompsNotFound creates a PostAppcompsNotFound with default headers values
func NewPostAppcompsNotFound() *PostAppcompsNotFound {
	return &PostAppcompsNotFound{}
}

/*PostAppcompsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostAppcompsNotFound struct {
}

func (o *PostAppcompsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/appcomps/][%d] postAppcompsNotFound ", 404)
}

func (o *PostAppcompsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAppcompsMethodNotAllowed creates a PostAppcompsMethodNotAllowed with default headers values
func NewPostAppcompsMethodNotAllowed() *PostAppcompsMethodNotAllowed {
	return &PostAppcompsMethodNotAllowed{}
}

/*PostAppcompsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostAppcompsMethodNotAllowed struct {
}

func (o *PostAppcompsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/appcomps/][%d] postAppcompsMethodNotAllowed ", 405)
}

func (o *PostAppcompsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAppcompsGone creates a PostAppcompsGone with default headers values
func NewPostAppcompsGone() *PostAppcompsGone {
	return &PostAppcompsGone{}
}

/*PostAppcompsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostAppcompsGone struct {
}

func (o *PostAppcompsGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/appcomps/][%d] postAppcompsGone ", 410)
}

func (o *PostAppcompsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAppcompsInternalServerError creates a PostAppcompsInternalServerError with default headers values
func NewPostAppcompsInternalServerError() *PostAppcompsInternalServerError {
	return &PostAppcompsInternalServerError{}
}

/*PostAppcompsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type PostAppcompsInternalServerError struct {
}

func (o *PostAppcompsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/appcomps/][%d] postAppcompsInternalServerError ", 500)
}

func (o *PostAppcompsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAppcompsServiceUnavailable creates a PostAppcompsServiceUnavailable with default headers values
func NewPostAppcompsServiceUnavailable() *PostAppcompsServiceUnavailable {
	return &PostAppcompsServiceUnavailable{}
}

/*PostAppcompsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostAppcompsServiceUnavailable struct {
}

func (o *PostAppcompsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/appcomps/][%d] postAppcompsServiceUnavailable ", 503)
}

func (o *PostAppcompsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostAppcompsOKBody post appcomps o k body
swagger:model PostAppcompsOKBody
*/
type PostAppcompsOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post appcomps o k body
func (o *PostAppcompsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAppcompsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAppcompsOKBody) UnmarshalBinary(b []byte) error {
	var res PostAppcompsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
