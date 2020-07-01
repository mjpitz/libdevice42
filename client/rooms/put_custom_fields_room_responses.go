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

// PutCustomFieldsRoomReader is a Reader for the PutCustomFieldsRoom structure.
type PutCustomFieldsRoomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomFieldsRoomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomFieldsRoomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutCustomFieldsRoomBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutCustomFieldsRoomUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutCustomFieldsRoomForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutCustomFieldsRoomNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPutCustomFieldsRoomMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPutCustomFieldsRoomGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutCustomFieldsRoomInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutCustomFieldsRoomServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCustomFieldsRoomOK creates a PutCustomFieldsRoomOK with default headers values
func NewPutCustomFieldsRoomOK() *PutCustomFieldsRoomOK {
	return &PutCustomFieldsRoomOK{}
}

/*PutCustomFieldsRoomOK handles this case with default header values.

The above command returns results like this:
*/
type PutCustomFieldsRoomOK struct {
	Payload *PutCustomFieldsRoomOKBody
}

func (o *PutCustomFieldsRoomOK) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/room/][%d] putCustomFieldsRoomOK  %+v", 200, o.Payload)
}

func (o *PutCustomFieldsRoomOK) GetPayload() *PutCustomFieldsRoomOKBody {
	return o.Payload
}

func (o *PutCustomFieldsRoomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutCustomFieldsRoomOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCustomFieldsRoomBadRequest creates a PutCustomFieldsRoomBadRequest with default headers values
func NewPutCustomFieldsRoomBadRequest() *PutCustomFieldsRoomBadRequest {
	return &PutCustomFieldsRoomBadRequest{}
}

/*PutCustomFieldsRoomBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PutCustomFieldsRoomBadRequest struct {
}

func (o *PutCustomFieldsRoomBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/room/][%d] putCustomFieldsRoomBadRequest ", 400)
}

func (o *PutCustomFieldsRoomBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsRoomUnauthorized creates a PutCustomFieldsRoomUnauthorized with default headers values
func NewPutCustomFieldsRoomUnauthorized() *PutCustomFieldsRoomUnauthorized {
	return &PutCustomFieldsRoomUnauthorized{}
}

/*PutCustomFieldsRoomUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PutCustomFieldsRoomUnauthorized struct {
}

func (o *PutCustomFieldsRoomUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/room/][%d] putCustomFieldsRoomUnauthorized ", 401)
}

func (o *PutCustomFieldsRoomUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsRoomForbidden creates a PutCustomFieldsRoomForbidden with default headers values
func NewPutCustomFieldsRoomForbidden() *PutCustomFieldsRoomForbidden {
	return &PutCustomFieldsRoomForbidden{}
}

/*PutCustomFieldsRoomForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PutCustomFieldsRoomForbidden struct {
}

func (o *PutCustomFieldsRoomForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/room/][%d] putCustomFieldsRoomForbidden ", 403)
}

func (o *PutCustomFieldsRoomForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsRoomNotFound creates a PutCustomFieldsRoomNotFound with default headers values
func NewPutCustomFieldsRoomNotFound() *PutCustomFieldsRoomNotFound {
	return &PutCustomFieldsRoomNotFound{}
}

/*PutCustomFieldsRoomNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PutCustomFieldsRoomNotFound struct {
}

func (o *PutCustomFieldsRoomNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/room/][%d] putCustomFieldsRoomNotFound ", 404)
}

func (o *PutCustomFieldsRoomNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsRoomMethodNotAllowed creates a PutCustomFieldsRoomMethodNotAllowed with default headers values
func NewPutCustomFieldsRoomMethodNotAllowed() *PutCustomFieldsRoomMethodNotAllowed {
	return &PutCustomFieldsRoomMethodNotAllowed{}
}

/*PutCustomFieldsRoomMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PutCustomFieldsRoomMethodNotAllowed struct {
}

func (o *PutCustomFieldsRoomMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/room/][%d] putCustomFieldsRoomMethodNotAllowed ", 405)
}

func (o *PutCustomFieldsRoomMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsRoomGone creates a PutCustomFieldsRoomGone with default headers values
func NewPutCustomFieldsRoomGone() *PutCustomFieldsRoomGone {
	return &PutCustomFieldsRoomGone{}
}

/*PutCustomFieldsRoomGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PutCustomFieldsRoomGone struct {
}

func (o *PutCustomFieldsRoomGone) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/room/][%d] putCustomFieldsRoomGone ", 410)
}

func (o *PutCustomFieldsRoomGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsRoomInternalServerError creates a PutCustomFieldsRoomInternalServerError with default headers values
func NewPutCustomFieldsRoomInternalServerError() *PutCustomFieldsRoomInternalServerError {
	return &PutCustomFieldsRoomInternalServerError{}
}

/*PutCustomFieldsRoomInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type PutCustomFieldsRoomInternalServerError struct {
}

func (o *PutCustomFieldsRoomInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/room/][%d] putCustomFieldsRoomInternalServerError ", 500)
}

func (o *PutCustomFieldsRoomInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsRoomServiceUnavailable creates a PutCustomFieldsRoomServiceUnavailable with default headers values
func NewPutCustomFieldsRoomServiceUnavailable() *PutCustomFieldsRoomServiceUnavailable {
	return &PutCustomFieldsRoomServiceUnavailable{}
}

/*PutCustomFieldsRoomServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PutCustomFieldsRoomServiceUnavailable struct {
}

func (o *PutCustomFieldsRoomServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/room/][%d] putCustomFieldsRoomServiceUnavailable ", 503)
}

func (o *PutCustomFieldsRoomServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PutCustomFieldsRoomOKBody put custom fields room o k body
swagger:model PutCustomFieldsRoomOKBody
*/
type PutCustomFieldsRoomOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this put custom fields room o k body
func (o *PutCustomFieldsRoomOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutCustomFieldsRoomOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCustomFieldsRoomOKBody) UnmarshalBinary(b []byte) error {
	var res PutCustomFieldsRoomOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
