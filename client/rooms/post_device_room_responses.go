// Code generated by go-swagger; DO NOT EDIT.

package rooms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostDeviceRoomReader is a Reader for the PostDeviceRoom structure.
type PostDeviceRoomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDeviceRoomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDeviceRoomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDeviceRoomBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDeviceRoomUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDeviceRoomForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDeviceRoomNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostDeviceRoomMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostDeviceRoomGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDeviceRoomInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostDeviceRoomServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostDeviceRoomOK creates a PostDeviceRoomOK with default headers values
func NewPostDeviceRoomOK() *PostDeviceRoomOK {
	return &PostDeviceRoomOK{}
}

/*PostDeviceRoomOK handles this case with default header values.

The above command returns results like this:
*/
type PostDeviceRoomOK struct {
	Payload *PostDeviceRoomOKBody
}

func (o *PostDeviceRoomOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/device/room/][%d] postDeviceRoomOK  %+v", 200, o.Payload)
}

func (o *PostDeviceRoomOK) GetPayload() *PostDeviceRoomOKBody {
	return o.Payload
}

func (o *PostDeviceRoomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostDeviceRoomOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDeviceRoomBadRequest creates a PostDeviceRoomBadRequest with default headers values
func NewPostDeviceRoomBadRequest() *PostDeviceRoomBadRequest {
	return &PostDeviceRoomBadRequest{}
}

/*PostDeviceRoomBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostDeviceRoomBadRequest struct {
}

func (o *PostDeviceRoomBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/device/room/][%d] postDeviceRoomBadRequest ", 400)
}

func (o *PostDeviceRoomBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDeviceRoomUnauthorized creates a PostDeviceRoomUnauthorized with default headers values
func NewPostDeviceRoomUnauthorized() *PostDeviceRoomUnauthorized {
	return &PostDeviceRoomUnauthorized{}
}

/*PostDeviceRoomUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostDeviceRoomUnauthorized struct {
}

func (o *PostDeviceRoomUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/device/room/][%d] postDeviceRoomUnauthorized ", 401)
}

func (o *PostDeviceRoomUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDeviceRoomForbidden creates a PostDeviceRoomForbidden with default headers values
func NewPostDeviceRoomForbidden() *PostDeviceRoomForbidden {
	return &PostDeviceRoomForbidden{}
}

/*PostDeviceRoomForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostDeviceRoomForbidden struct {
}

func (o *PostDeviceRoomForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/device/room/][%d] postDeviceRoomForbidden ", 403)
}

func (o *PostDeviceRoomForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDeviceRoomNotFound creates a PostDeviceRoomNotFound with default headers values
func NewPostDeviceRoomNotFound() *PostDeviceRoomNotFound {
	return &PostDeviceRoomNotFound{}
}

/*PostDeviceRoomNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostDeviceRoomNotFound struct {
}

func (o *PostDeviceRoomNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/device/room/][%d] postDeviceRoomNotFound ", 404)
}

func (o *PostDeviceRoomNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDeviceRoomMethodNotAllowed creates a PostDeviceRoomMethodNotAllowed with default headers values
func NewPostDeviceRoomMethodNotAllowed() *PostDeviceRoomMethodNotAllowed {
	return &PostDeviceRoomMethodNotAllowed{}
}

/*PostDeviceRoomMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostDeviceRoomMethodNotAllowed struct {
}

func (o *PostDeviceRoomMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/device/room/][%d] postDeviceRoomMethodNotAllowed ", 405)
}

func (o *PostDeviceRoomMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDeviceRoomGone creates a PostDeviceRoomGone with default headers values
func NewPostDeviceRoomGone() *PostDeviceRoomGone {
	return &PostDeviceRoomGone{}
}

/*PostDeviceRoomGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostDeviceRoomGone struct {
}

func (o *PostDeviceRoomGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/device/room/][%d] postDeviceRoomGone ", 410)
}

func (o *PostDeviceRoomGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDeviceRoomInternalServerError creates a PostDeviceRoomInternalServerError with default headers values
func NewPostDeviceRoomInternalServerError() *PostDeviceRoomInternalServerError {
	return &PostDeviceRoomInternalServerError{}
}

/*PostDeviceRoomInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type PostDeviceRoomInternalServerError struct {
}

func (o *PostDeviceRoomInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/device/room/][%d] postDeviceRoomInternalServerError ", 500)
}

func (o *PostDeviceRoomInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDeviceRoomServiceUnavailable creates a PostDeviceRoomServiceUnavailable with default headers values
func NewPostDeviceRoomServiceUnavailable() *PostDeviceRoomServiceUnavailable {
	return &PostDeviceRoomServiceUnavailable{}
}

/*PostDeviceRoomServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostDeviceRoomServiceUnavailable struct {
}

func (o *PostDeviceRoomServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/device/room/][%d] postDeviceRoomServiceUnavailable ", 503)
}

func (o *PostDeviceRoomServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostDeviceRoomOKBody post device room o k body
swagger:model PostDeviceRoomOKBody
*/
type PostDeviceRoomOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post device room o k body
func (o *PostDeviceRoomOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostDeviceRoomOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDeviceRoomOKBody) UnmarshalBinary(b []byte) error {
	var res PostDeviceRoomOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
