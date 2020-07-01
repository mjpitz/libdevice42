// Code generated by go-swagger; DO NOT EDIT.

package power_circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeletePowerCircuitReader is a Reader for the DeletePowerCircuit structure.
type DeletePowerCircuitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePowerCircuitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePowerCircuitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeletePowerCircuitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeletePowerCircuitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePowerCircuitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePowerCircuitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeletePowerCircuitMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewDeletePowerCircuitGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeletePowerCircuitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeletePowerCircuitServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePowerCircuitOK creates a DeletePowerCircuitOK with default headers values
func NewDeletePowerCircuitOK() *DeletePowerCircuitOK {
	return &DeletePowerCircuitOK{}
}

/*DeletePowerCircuitOK handles this case with default header values.

The above command returns results like this:
*/
type DeletePowerCircuitOK struct {
	Payload *DeletePowerCircuitOKBody
}

func (o *DeletePowerCircuitOK) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/powercircuits/{id}/][%d] deletePowerCircuitOK  %+v", 200, o.Payload)
}

func (o *DeletePowerCircuitOK) GetPayload() *DeletePowerCircuitOKBody {
	return o.Payload
}

func (o *DeletePowerCircuitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeletePowerCircuitOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePowerCircuitBadRequest creates a DeletePowerCircuitBadRequest with default headers values
func NewDeletePowerCircuitBadRequest() *DeletePowerCircuitBadRequest {
	return &DeletePowerCircuitBadRequest{}
}

/*DeletePowerCircuitBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type DeletePowerCircuitBadRequest struct {
}

func (o *DeletePowerCircuitBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/powercircuits/{id}/][%d] deletePowerCircuitBadRequest ", 400)
}

func (o *DeletePowerCircuitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePowerCircuitUnauthorized creates a DeletePowerCircuitUnauthorized with default headers values
func NewDeletePowerCircuitUnauthorized() *DeletePowerCircuitUnauthorized {
	return &DeletePowerCircuitUnauthorized{}
}

/*DeletePowerCircuitUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type DeletePowerCircuitUnauthorized struct {
}

func (o *DeletePowerCircuitUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/powercircuits/{id}/][%d] deletePowerCircuitUnauthorized ", 401)
}

func (o *DeletePowerCircuitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePowerCircuitForbidden creates a DeletePowerCircuitForbidden with default headers values
func NewDeletePowerCircuitForbidden() *DeletePowerCircuitForbidden {
	return &DeletePowerCircuitForbidden{}
}

/*DeletePowerCircuitForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type DeletePowerCircuitForbidden struct {
}

func (o *DeletePowerCircuitForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/powercircuits/{id}/][%d] deletePowerCircuitForbidden ", 403)
}

func (o *DeletePowerCircuitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePowerCircuitNotFound creates a DeletePowerCircuitNotFound with default headers values
func NewDeletePowerCircuitNotFound() *DeletePowerCircuitNotFound {
	return &DeletePowerCircuitNotFound{}
}

/*DeletePowerCircuitNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type DeletePowerCircuitNotFound struct {
}

func (o *DeletePowerCircuitNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/powercircuits/{id}/][%d] deletePowerCircuitNotFound ", 404)
}

func (o *DeletePowerCircuitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePowerCircuitMethodNotAllowed creates a DeletePowerCircuitMethodNotAllowed with default headers values
func NewDeletePowerCircuitMethodNotAllowed() *DeletePowerCircuitMethodNotAllowed {
	return &DeletePowerCircuitMethodNotAllowed{}
}

/*DeletePowerCircuitMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type DeletePowerCircuitMethodNotAllowed struct {
}

func (o *DeletePowerCircuitMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/powercircuits/{id}/][%d] deletePowerCircuitMethodNotAllowed ", 405)
}

func (o *DeletePowerCircuitMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePowerCircuitGone creates a DeletePowerCircuitGone with default headers values
func NewDeletePowerCircuitGone() *DeletePowerCircuitGone {
	return &DeletePowerCircuitGone{}
}

/*DeletePowerCircuitGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type DeletePowerCircuitGone struct {
}

func (o *DeletePowerCircuitGone) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/powercircuits/{id}/][%d] deletePowerCircuitGone ", 410)
}

func (o *DeletePowerCircuitGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePowerCircuitInternalServerError creates a DeletePowerCircuitInternalServerError with default headers values
func NewDeletePowerCircuitInternalServerError() *DeletePowerCircuitInternalServerError {
	return &DeletePowerCircuitInternalServerError{}
}

/*DeletePowerCircuitInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type DeletePowerCircuitInternalServerError struct {
}

func (o *DeletePowerCircuitInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/powercircuits/{id}/][%d] deletePowerCircuitInternalServerError ", 500)
}

func (o *DeletePowerCircuitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePowerCircuitServiceUnavailable creates a DeletePowerCircuitServiceUnavailable with default headers values
func NewDeletePowerCircuitServiceUnavailable() *DeletePowerCircuitServiceUnavailable {
	return &DeletePowerCircuitServiceUnavailable{}
}

/*DeletePowerCircuitServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type DeletePowerCircuitServiceUnavailable struct {
}

func (o *DeletePowerCircuitServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/powercircuits/{id}/][%d] deletePowerCircuitServiceUnavailable ", 503)
}

func (o *DeletePowerCircuitServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeletePowerCircuitOKBody delete power circuit o k body
swagger:model DeletePowerCircuitOKBody
*/
type DeletePowerCircuitOKBody struct {

	// deleted
	Deleted interface{} `json:"deleted,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`
}

// Validate validates this delete power circuit o k body
func (o *DeletePowerCircuitOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeletePowerCircuitOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePowerCircuitOKBody) UnmarshalBinary(b []byte) error {
	var res DeletePowerCircuitOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
