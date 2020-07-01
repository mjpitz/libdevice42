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

// DeletePdusRackReader is a Reader for the DeletePdusRack structure.
type DeletePdusRackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePdusRackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePdusRackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeletePdusRackBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeletePdusRackUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePdusRackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePdusRackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeletePdusRackMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewDeletePdusRackGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeletePdusRackInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeletePdusRackServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePdusRackOK creates a DeletePdusRackOK with default headers values
func NewDeletePdusRackOK() *DeletePdusRackOK {
	return &DeletePdusRackOK{}
}

/*DeletePdusRackOK handles this case with default header values.

The above command returns results like this:
*/
type DeletePdusRackOK struct {
	Payload *DeletePdusRackOKBody
}

func (o *DeletePdusRackOK) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/pdus/rack/{id}/][%d] deletePdusRackOK  %+v", 200, o.Payload)
}

func (o *DeletePdusRackOK) GetPayload() *DeletePdusRackOKBody {
	return o.Payload
}

func (o *DeletePdusRackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeletePdusRackOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePdusRackBadRequest creates a DeletePdusRackBadRequest with default headers values
func NewDeletePdusRackBadRequest() *DeletePdusRackBadRequest {
	return &DeletePdusRackBadRequest{}
}

/*DeletePdusRackBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type DeletePdusRackBadRequest struct {
}

func (o *DeletePdusRackBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/pdus/rack/{id}/][%d] deletePdusRackBadRequest ", 400)
}

func (o *DeletePdusRackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePdusRackUnauthorized creates a DeletePdusRackUnauthorized with default headers values
func NewDeletePdusRackUnauthorized() *DeletePdusRackUnauthorized {
	return &DeletePdusRackUnauthorized{}
}

/*DeletePdusRackUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type DeletePdusRackUnauthorized struct {
}

func (o *DeletePdusRackUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/pdus/rack/{id}/][%d] deletePdusRackUnauthorized ", 401)
}

func (o *DeletePdusRackUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePdusRackForbidden creates a DeletePdusRackForbidden with default headers values
func NewDeletePdusRackForbidden() *DeletePdusRackForbidden {
	return &DeletePdusRackForbidden{}
}

/*DeletePdusRackForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type DeletePdusRackForbidden struct {
}

func (o *DeletePdusRackForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/pdus/rack/{id}/][%d] deletePdusRackForbidden ", 403)
}

func (o *DeletePdusRackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePdusRackNotFound creates a DeletePdusRackNotFound with default headers values
func NewDeletePdusRackNotFound() *DeletePdusRackNotFound {
	return &DeletePdusRackNotFound{}
}

/*DeletePdusRackNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type DeletePdusRackNotFound struct {
}

func (o *DeletePdusRackNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/pdus/rack/{id}/][%d] deletePdusRackNotFound ", 404)
}

func (o *DeletePdusRackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePdusRackMethodNotAllowed creates a DeletePdusRackMethodNotAllowed with default headers values
func NewDeletePdusRackMethodNotAllowed() *DeletePdusRackMethodNotAllowed {
	return &DeletePdusRackMethodNotAllowed{}
}

/*DeletePdusRackMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type DeletePdusRackMethodNotAllowed struct {
}

func (o *DeletePdusRackMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/pdus/rack/{id}/][%d] deletePdusRackMethodNotAllowed ", 405)
}

func (o *DeletePdusRackMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePdusRackGone creates a DeletePdusRackGone with default headers values
func NewDeletePdusRackGone() *DeletePdusRackGone {
	return &DeletePdusRackGone{}
}

/*DeletePdusRackGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type DeletePdusRackGone struct {
}

func (o *DeletePdusRackGone) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/pdus/rack/{id}/][%d] deletePdusRackGone ", 410)
}

func (o *DeletePdusRackGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePdusRackInternalServerError creates a DeletePdusRackInternalServerError with default headers values
func NewDeletePdusRackInternalServerError() *DeletePdusRackInternalServerError {
	return &DeletePdusRackInternalServerError{}
}

/*DeletePdusRackInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type DeletePdusRackInternalServerError struct {
}

func (o *DeletePdusRackInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/pdus/rack/{id}/][%d] deletePdusRackInternalServerError ", 500)
}

func (o *DeletePdusRackInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePdusRackServiceUnavailable creates a DeletePdusRackServiceUnavailable with default headers values
func NewDeletePdusRackServiceUnavailable() *DeletePdusRackServiceUnavailable {
	return &DeletePdusRackServiceUnavailable{}
}

/*DeletePdusRackServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type DeletePdusRackServiceUnavailable struct {
}

func (o *DeletePdusRackServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/pdus/rack/{id}/][%d] deletePdusRackServiceUnavailable ", 503)
}

func (o *DeletePdusRackServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeletePdusRackOKBody delete pdus rack o k body
swagger:model DeletePdusRackOKBody
*/
type DeletePdusRackOKBody struct {

	// deleted
	Deleted interface{} `json:"deleted,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`
}

// Validate validates this delete pdus rack o k body
func (o *DeletePdusRackOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeletePdusRackOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePdusRackOKBody) UnmarshalBinary(b []byte) error {
	var res DeletePdusRackOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
