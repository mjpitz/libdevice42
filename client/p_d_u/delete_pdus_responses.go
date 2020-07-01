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

// DeletePdusReader is a Reader for the DeletePdus structure.
type DeletePdusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePdusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePdusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeletePdusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeletePdusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePdusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePdusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeletePdusMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewDeletePdusGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeletePdusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeletePdusServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePdusOK creates a DeletePdusOK with default headers values
func NewDeletePdusOK() *DeletePdusOK {
	return &DeletePdusOK{}
}

/*DeletePdusOK handles this case with default header values.

The above command returns results like this:
*/
type DeletePdusOK struct {
	Payload *DeletePdusOKBody
}

func (o *DeletePdusOK) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/pdus/{ID}/][%d] deletePdusOK  %+v", 200, o.Payload)
}

func (o *DeletePdusOK) GetPayload() *DeletePdusOKBody {
	return o.Payload
}

func (o *DeletePdusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeletePdusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePdusBadRequest creates a DeletePdusBadRequest with default headers values
func NewDeletePdusBadRequest() *DeletePdusBadRequest {
	return &DeletePdusBadRequest{}
}

/*DeletePdusBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type DeletePdusBadRequest struct {
}

func (o *DeletePdusBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/pdus/{ID}/][%d] deletePdusBadRequest ", 400)
}

func (o *DeletePdusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePdusUnauthorized creates a DeletePdusUnauthorized with default headers values
func NewDeletePdusUnauthorized() *DeletePdusUnauthorized {
	return &DeletePdusUnauthorized{}
}

/*DeletePdusUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type DeletePdusUnauthorized struct {
}

func (o *DeletePdusUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/pdus/{ID}/][%d] deletePdusUnauthorized ", 401)
}

func (o *DeletePdusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePdusForbidden creates a DeletePdusForbidden with default headers values
func NewDeletePdusForbidden() *DeletePdusForbidden {
	return &DeletePdusForbidden{}
}

/*DeletePdusForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type DeletePdusForbidden struct {
}

func (o *DeletePdusForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/pdus/{ID}/][%d] deletePdusForbidden ", 403)
}

func (o *DeletePdusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePdusNotFound creates a DeletePdusNotFound with default headers values
func NewDeletePdusNotFound() *DeletePdusNotFound {
	return &DeletePdusNotFound{}
}

/*DeletePdusNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type DeletePdusNotFound struct {
}

func (o *DeletePdusNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/pdus/{ID}/][%d] deletePdusNotFound ", 404)
}

func (o *DeletePdusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePdusMethodNotAllowed creates a DeletePdusMethodNotAllowed with default headers values
func NewDeletePdusMethodNotAllowed() *DeletePdusMethodNotAllowed {
	return &DeletePdusMethodNotAllowed{}
}

/*DeletePdusMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type DeletePdusMethodNotAllowed struct {
}

func (o *DeletePdusMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/pdus/{ID}/][%d] deletePdusMethodNotAllowed ", 405)
}

func (o *DeletePdusMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePdusGone creates a DeletePdusGone with default headers values
func NewDeletePdusGone() *DeletePdusGone {
	return &DeletePdusGone{}
}

/*DeletePdusGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type DeletePdusGone struct {
}

func (o *DeletePdusGone) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/pdus/{ID}/][%d] deletePdusGone ", 410)
}

func (o *DeletePdusGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePdusInternalServerError creates a DeletePdusInternalServerError with default headers values
func NewDeletePdusInternalServerError() *DeletePdusInternalServerError {
	return &DeletePdusInternalServerError{}
}

/*DeletePdusInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type DeletePdusInternalServerError struct {
}

func (o *DeletePdusInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/pdus/{ID}/][%d] deletePdusInternalServerError ", 500)
}

func (o *DeletePdusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePdusServiceUnavailable creates a DeletePdusServiceUnavailable with default headers values
func NewDeletePdusServiceUnavailable() *DeletePdusServiceUnavailable {
	return &DeletePdusServiceUnavailable{}
}

/*DeletePdusServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type DeletePdusServiceUnavailable struct {
}

func (o *DeletePdusServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/pdus/{ID}/][%d] deletePdusServiceUnavailable ", 503)
}

func (o *DeletePdusServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeletePdusOKBody delete pdus o k body
swagger:model DeletePdusOKBody
*/
type DeletePdusOKBody struct {

	// deleted
	Deleted interface{} `json:"deleted,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`
}

// Validate validates this delete pdus o k body
func (o *DeletePdusOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeletePdusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePdusOKBody) UnmarshalBinary(b []byte) error {
	var res DeletePdusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
