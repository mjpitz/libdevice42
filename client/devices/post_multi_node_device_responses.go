// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/mjpitz/libdevice42/models"
)

// PostMultiNodeDeviceReader is a Reader for the PostMultiNodeDevice structure.
type PostMultiNodeDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMultiNodeDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostMultiNodeDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostMultiNodeDeviceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostMultiNodeDeviceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostMultiNodeDeviceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostMultiNodeDeviceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostMultiNodeDeviceMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostMultiNodeDeviceGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostMultiNodeDeviceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostMultiNodeDeviceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostMultiNodeDeviceOK creates a PostMultiNodeDeviceOK with default headers values
func NewPostMultiNodeDeviceOK() *PostMultiNodeDeviceOK {
	return &PostMultiNodeDeviceOK{}
}

/*PostMultiNodeDeviceOK handles this case with default header values.

The above command returns results like this:
*/
type PostMultiNodeDeviceOK struct {
	Payload *PostMultiNodeDeviceOKBody
}

func (o *PostMultiNodeDeviceOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/multinode_device/][%d] postMultiNodeDeviceOK  %+v", 200, o.Payload)
}

func (o *PostMultiNodeDeviceOK) GetPayload() *PostMultiNodeDeviceOKBody {
	return o.Payload
}

func (o *PostMultiNodeDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostMultiNodeDeviceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMultiNodeDeviceBadRequest creates a PostMultiNodeDeviceBadRequest with default headers values
func NewPostMultiNodeDeviceBadRequest() *PostMultiNodeDeviceBadRequest {
	return &PostMultiNodeDeviceBadRequest{}
}

/*PostMultiNodeDeviceBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostMultiNodeDeviceBadRequest struct {
}

func (o *PostMultiNodeDeviceBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/multinode_device/][%d] postMultiNodeDeviceBadRequest ", 400)
}

func (o *PostMultiNodeDeviceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMultiNodeDeviceUnauthorized creates a PostMultiNodeDeviceUnauthorized with default headers values
func NewPostMultiNodeDeviceUnauthorized() *PostMultiNodeDeviceUnauthorized {
	return &PostMultiNodeDeviceUnauthorized{}
}

/*PostMultiNodeDeviceUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostMultiNodeDeviceUnauthorized struct {
}

func (o *PostMultiNodeDeviceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/multinode_device/][%d] postMultiNodeDeviceUnauthorized ", 401)
}

func (o *PostMultiNodeDeviceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMultiNodeDeviceForbidden creates a PostMultiNodeDeviceForbidden with default headers values
func NewPostMultiNodeDeviceForbidden() *PostMultiNodeDeviceForbidden {
	return &PostMultiNodeDeviceForbidden{}
}

/*PostMultiNodeDeviceForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostMultiNodeDeviceForbidden struct {
}

func (o *PostMultiNodeDeviceForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/multinode_device/][%d] postMultiNodeDeviceForbidden ", 403)
}

func (o *PostMultiNodeDeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMultiNodeDeviceNotFound creates a PostMultiNodeDeviceNotFound with default headers values
func NewPostMultiNodeDeviceNotFound() *PostMultiNodeDeviceNotFound {
	return &PostMultiNodeDeviceNotFound{}
}

/*PostMultiNodeDeviceNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostMultiNodeDeviceNotFound struct {
}

func (o *PostMultiNodeDeviceNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/multinode_device/][%d] postMultiNodeDeviceNotFound ", 404)
}

func (o *PostMultiNodeDeviceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMultiNodeDeviceMethodNotAllowed creates a PostMultiNodeDeviceMethodNotAllowed with default headers values
func NewPostMultiNodeDeviceMethodNotAllowed() *PostMultiNodeDeviceMethodNotAllowed {
	return &PostMultiNodeDeviceMethodNotAllowed{}
}

/*PostMultiNodeDeviceMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostMultiNodeDeviceMethodNotAllowed struct {
}

func (o *PostMultiNodeDeviceMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/multinode_device/][%d] postMultiNodeDeviceMethodNotAllowed ", 405)
}

func (o *PostMultiNodeDeviceMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMultiNodeDeviceGone creates a PostMultiNodeDeviceGone with default headers values
func NewPostMultiNodeDeviceGone() *PostMultiNodeDeviceGone {
	return &PostMultiNodeDeviceGone{}
}

/*PostMultiNodeDeviceGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostMultiNodeDeviceGone struct {
}

func (o *PostMultiNodeDeviceGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/multinode_device/][%d] postMultiNodeDeviceGone ", 410)
}

func (o *PostMultiNodeDeviceGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMultiNodeDeviceInternalServerError creates a PostMultiNodeDeviceInternalServerError with default headers values
func NewPostMultiNodeDeviceInternalServerError() *PostMultiNodeDeviceInternalServerError {
	return &PostMultiNodeDeviceInternalServerError{}
}

/*PostMultiNodeDeviceInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type PostMultiNodeDeviceInternalServerError struct {
}

func (o *PostMultiNodeDeviceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/multinode_device/][%d] postMultiNodeDeviceInternalServerError ", 500)
}

func (o *PostMultiNodeDeviceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMultiNodeDeviceServiceUnavailable creates a PostMultiNodeDeviceServiceUnavailable with default headers values
func NewPostMultiNodeDeviceServiceUnavailable() *PostMultiNodeDeviceServiceUnavailable {
	return &PostMultiNodeDeviceServiceUnavailable{}
}

/*PostMultiNodeDeviceServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostMultiNodeDeviceServiceUnavailable struct {
}

func (o *PostMultiNodeDeviceServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/multinode_device/][%d] postMultiNodeDeviceServiceUnavailable ", 503)
}

func (o *PostMultiNodeDeviceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostMultiNodeDeviceOKBody post multi node device o k body
swagger:model PostMultiNodeDeviceOKBody
*/
type PostMultiNodeDeviceOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg models.DeviceName `json:"msg,omitempty"`
}

// Validate validates this post multi node device o k body
func (o *PostMultiNodeDeviceOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostMultiNodeDeviceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostMultiNodeDeviceOKBody) UnmarshalBinary(b []byte) error {
	var res PostMultiNodeDeviceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}