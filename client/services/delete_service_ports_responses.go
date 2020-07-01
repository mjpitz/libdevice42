// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteServicePortsReader is a Reader for the DeleteServicePorts structure.
type DeleteServicePortsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServicePortsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteServicePortsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteServicePortsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteServicePortsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteServicePortsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteServicePortsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeleteServicePortsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewDeleteServicePortsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteServicePortsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteServicePortsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteServicePortsOK creates a DeleteServicePortsOK with default headers values
func NewDeleteServicePortsOK() *DeleteServicePortsOK {
	return &DeleteServicePortsOK{}
}

/*DeleteServicePortsOK handles this case with default header values.

The above command returns results like this:
*/
type DeleteServicePortsOK struct {
	Payload *DeleteServicePortsOKBody
}

func (o *DeleteServicePortsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/service_ports/{id}/][%d] deleteServicePortsOK  %+v", 200, o.Payload)
}

func (o *DeleteServicePortsOK) GetPayload() *DeleteServicePortsOKBody {
	return o.Payload
}

func (o *DeleteServicePortsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteServicePortsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServicePortsBadRequest creates a DeleteServicePortsBadRequest with default headers values
func NewDeleteServicePortsBadRequest() *DeleteServicePortsBadRequest {
	return &DeleteServicePortsBadRequest{}
}

/*DeleteServicePortsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type DeleteServicePortsBadRequest struct {
}

func (o *DeleteServicePortsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/service_ports/{id}/][%d] deleteServicePortsBadRequest ", 400)
}

func (o *DeleteServicePortsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServicePortsUnauthorized creates a DeleteServicePortsUnauthorized with default headers values
func NewDeleteServicePortsUnauthorized() *DeleteServicePortsUnauthorized {
	return &DeleteServicePortsUnauthorized{}
}

/*DeleteServicePortsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type DeleteServicePortsUnauthorized struct {
}

func (o *DeleteServicePortsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/service_ports/{id}/][%d] deleteServicePortsUnauthorized ", 401)
}

func (o *DeleteServicePortsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServicePortsForbidden creates a DeleteServicePortsForbidden with default headers values
func NewDeleteServicePortsForbidden() *DeleteServicePortsForbidden {
	return &DeleteServicePortsForbidden{}
}

/*DeleteServicePortsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type DeleteServicePortsForbidden struct {
}

func (o *DeleteServicePortsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/service_ports/{id}/][%d] deleteServicePortsForbidden ", 403)
}

func (o *DeleteServicePortsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServicePortsNotFound creates a DeleteServicePortsNotFound with default headers values
func NewDeleteServicePortsNotFound() *DeleteServicePortsNotFound {
	return &DeleteServicePortsNotFound{}
}

/*DeleteServicePortsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type DeleteServicePortsNotFound struct {
}

func (o *DeleteServicePortsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/service_ports/{id}/][%d] deleteServicePortsNotFound ", 404)
}

func (o *DeleteServicePortsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServicePortsMethodNotAllowed creates a DeleteServicePortsMethodNotAllowed with default headers values
func NewDeleteServicePortsMethodNotAllowed() *DeleteServicePortsMethodNotAllowed {
	return &DeleteServicePortsMethodNotAllowed{}
}

/*DeleteServicePortsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type DeleteServicePortsMethodNotAllowed struct {
}

func (o *DeleteServicePortsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/service_ports/{id}/][%d] deleteServicePortsMethodNotAllowed ", 405)
}

func (o *DeleteServicePortsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServicePortsGone creates a DeleteServicePortsGone with default headers values
func NewDeleteServicePortsGone() *DeleteServicePortsGone {
	return &DeleteServicePortsGone{}
}

/*DeleteServicePortsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type DeleteServicePortsGone struct {
}

func (o *DeleteServicePortsGone) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/service_ports/{id}/][%d] deleteServicePortsGone ", 410)
}

func (o *DeleteServicePortsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServicePortsInternalServerError creates a DeleteServicePortsInternalServerError with default headers values
func NewDeleteServicePortsInternalServerError() *DeleteServicePortsInternalServerError {
	return &DeleteServicePortsInternalServerError{}
}

/*DeleteServicePortsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type DeleteServicePortsInternalServerError struct {
}

func (o *DeleteServicePortsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/service_ports/{id}/][%d] deleteServicePortsInternalServerError ", 500)
}

func (o *DeleteServicePortsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServicePortsServiceUnavailable creates a DeleteServicePortsServiceUnavailable with default headers values
func NewDeleteServicePortsServiceUnavailable() *DeleteServicePortsServiceUnavailable {
	return &DeleteServicePortsServiceUnavailable{}
}

/*DeleteServicePortsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type DeleteServicePortsServiceUnavailable struct {
}

func (o *DeleteServicePortsServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/service_ports/{id}/][%d] deleteServicePortsServiceUnavailable ", 503)
}

func (o *DeleteServicePortsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeleteServicePortsOKBody delete service ports o k body
swagger:model DeleteServicePortsOKBody
*/
type DeleteServicePortsOKBody struct {

	// deleted
	Deleted interface{} `json:"deleted,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`
}

// Validate validates this delete service ports o k body
func (o *DeleteServicePortsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteServicePortsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteServicePortsOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteServicePortsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
