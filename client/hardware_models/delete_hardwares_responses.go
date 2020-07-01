// Code generated by go-swagger; DO NOT EDIT.

package hardware_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteHardwaresReader is a Reader for the DeleteHardwares structure.
type DeleteHardwaresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteHardwaresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteHardwaresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteHardwaresBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteHardwaresUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteHardwaresForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteHardwaresNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeleteHardwaresMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewDeleteHardwaresGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteHardwaresInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteHardwaresServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteHardwaresOK creates a DeleteHardwaresOK with default headers values
func NewDeleteHardwaresOK() *DeleteHardwaresOK {
	return &DeleteHardwaresOK{}
}

/*DeleteHardwaresOK handles this case with default header values.

The above command returns results like this:
*/
type DeleteHardwaresOK struct {
	Payload *DeleteHardwaresOKBody
}

func (o *DeleteHardwaresOK) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/hardwares/{id}/][%d] deleteHardwaresOK  %+v", 200, o.Payload)
}

func (o *DeleteHardwaresOK) GetPayload() *DeleteHardwaresOKBody {
	return o.Payload
}

func (o *DeleteHardwaresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteHardwaresOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteHardwaresBadRequest creates a DeleteHardwaresBadRequest with default headers values
func NewDeleteHardwaresBadRequest() *DeleteHardwaresBadRequest {
	return &DeleteHardwaresBadRequest{}
}

/*DeleteHardwaresBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type DeleteHardwaresBadRequest struct {
}

func (o *DeleteHardwaresBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/hardwares/{id}/][%d] deleteHardwaresBadRequest ", 400)
}

func (o *DeleteHardwaresBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHardwaresUnauthorized creates a DeleteHardwaresUnauthorized with default headers values
func NewDeleteHardwaresUnauthorized() *DeleteHardwaresUnauthorized {
	return &DeleteHardwaresUnauthorized{}
}

/*DeleteHardwaresUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type DeleteHardwaresUnauthorized struct {
}

func (o *DeleteHardwaresUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/hardwares/{id}/][%d] deleteHardwaresUnauthorized ", 401)
}

func (o *DeleteHardwaresUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHardwaresForbidden creates a DeleteHardwaresForbidden with default headers values
func NewDeleteHardwaresForbidden() *DeleteHardwaresForbidden {
	return &DeleteHardwaresForbidden{}
}

/*DeleteHardwaresForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type DeleteHardwaresForbidden struct {
}

func (o *DeleteHardwaresForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/hardwares/{id}/][%d] deleteHardwaresForbidden ", 403)
}

func (o *DeleteHardwaresForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHardwaresNotFound creates a DeleteHardwaresNotFound with default headers values
func NewDeleteHardwaresNotFound() *DeleteHardwaresNotFound {
	return &DeleteHardwaresNotFound{}
}

/*DeleteHardwaresNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type DeleteHardwaresNotFound struct {
}

func (o *DeleteHardwaresNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/hardwares/{id}/][%d] deleteHardwaresNotFound ", 404)
}

func (o *DeleteHardwaresNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHardwaresMethodNotAllowed creates a DeleteHardwaresMethodNotAllowed with default headers values
func NewDeleteHardwaresMethodNotAllowed() *DeleteHardwaresMethodNotAllowed {
	return &DeleteHardwaresMethodNotAllowed{}
}

/*DeleteHardwaresMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type DeleteHardwaresMethodNotAllowed struct {
}

func (o *DeleteHardwaresMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/hardwares/{id}/][%d] deleteHardwaresMethodNotAllowed ", 405)
}

func (o *DeleteHardwaresMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHardwaresGone creates a DeleteHardwaresGone with default headers values
func NewDeleteHardwaresGone() *DeleteHardwaresGone {
	return &DeleteHardwaresGone{}
}

/*DeleteHardwaresGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type DeleteHardwaresGone struct {
}

func (o *DeleteHardwaresGone) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/hardwares/{id}/][%d] deleteHardwaresGone ", 410)
}

func (o *DeleteHardwaresGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHardwaresInternalServerError creates a DeleteHardwaresInternalServerError with default headers values
func NewDeleteHardwaresInternalServerError() *DeleteHardwaresInternalServerError {
	return &DeleteHardwaresInternalServerError{}
}

/*DeleteHardwaresInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type DeleteHardwaresInternalServerError struct {
}

func (o *DeleteHardwaresInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/hardwares/{id}/][%d] deleteHardwaresInternalServerError ", 500)
}

func (o *DeleteHardwaresInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHardwaresServiceUnavailable creates a DeleteHardwaresServiceUnavailable with default headers values
func NewDeleteHardwaresServiceUnavailable() *DeleteHardwaresServiceUnavailable {
	return &DeleteHardwaresServiceUnavailable{}
}

/*DeleteHardwaresServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type DeleteHardwaresServiceUnavailable struct {
}

func (o *DeleteHardwaresServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/hardwares/{id}/][%d] deleteHardwaresServiceUnavailable ", 503)
}

func (o *DeleteHardwaresServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeleteHardwaresOKBody delete hardwares o k body
swagger:model DeleteHardwaresOKBody
*/
type DeleteHardwaresOKBody struct {

	// deleted
	Deleted interface{} `json:"deleted,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`
}

// Validate validates this delete hardwares o k body
func (o *DeleteHardwaresOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteHardwaresOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteHardwaresOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteHardwaresOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}