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
)

// DeleteDeviceMountpointsReader is a Reader for the DeleteDeviceMountpoints structure.
type DeleteDeviceMountpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeviceMountpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDeviceMountpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDeviceMountpointsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteDeviceMountpointsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDeviceMountpointsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDeviceMountpointsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeleteDeviceMountpointsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewDeleteDeviceMountpointsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDeviceMountpointsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteDeviceMountpointsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDeviceMountpointsOK creates a DeleteDeviceMountpointsOK with default headers values
func NewDeleteDeviceMountpointsOK() *DeleteDeviceMountpointsOK {
	return &DeleteDeviceMountpointsOK{}
}

/*DeleteDeviceMountpointsOK handles this case with default header values.

The above command returns results like this:
*/
type DeleteDeviceMountpointsOK struct {
	Payload *DeleteDeviceMountpointsOKBody
}

func (o *DeleteDeviceMountpointsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device/mountpoints/{id}/][%d] deleteDeviceMountpointsOK  %+v", 200, o.Payload)
}

func (o *DeleteDeviceMountpointsOK) GetPayload() *DeleteDeviceMountpointsOKBody {
	return o.Payload
}

func (o *DeleteDeviceMountpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteDeviceMountpointsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceMountpointsBadRequest creates a DeleteDeviceMountpointsBadRequest with default headers values
func NewDeleteDeviceMountpointsBadRequest() *DeleteDeviceMountpointsBadRequest {
	return &DeleteDeviceMountpointsBadRequest{}
}

/*DeleteDeviceMountpointsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type DeleteDeviceMountpointsBadRequest struct {
}

func (o *DeleteDeviceMountpointsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device/mountpoints/{id}/][%d] deleteDeviceMountpointsBadRequest ", 400)
}

func (o *DeleteDeviceMountpointsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceMountpointsUnauthorized creates a DeleteDeviceMountpointsUnauthorized with default headers values
func NewDeleteDeviceMountpointsUnauthorized() *DeleteDeviceMountpointsUnauthorized {
	return &DeleteDeviceMountpointsUnauthorized{}
}

/*DeleteDeviceMountpointsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type DeleteDeviceMountpointsUnauthorized struct {
}

func (o *DeleteDeviceMountpointsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device/mountpoints/{id}/][%d] deleteDeviceMountpointsUnauthorized ", 401)
}

func (o *DeleteDeviceMountpointsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceMountpointsForbidden creates a DeleteDeviceMountpointsForbidden with default headers values
func NewDeleteDeviceMountpointsForbidden() *DeleteDeviceMountpointsForbidden {
	return &DeleteDeviceMountpointsForbidden{}
}

/*DeleteDeviceMountpointsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type DeleteDeviceMountpointsForbidden struct {
}

func (o *DeleteDeviceMountpointsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device/mountpoints/{id}/][%d] deleteDeviceMountpointsForbidden ", 403)
}

func (o *DeleteDeviceMountpointsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceMountpointsNotFound creates a DeleteDeviceMountpointsNotFound with default headers values
func NewDeleteDeviceMountpointsNotFound() *DeleteDeviceMountpointsNotFound {
	return &DeleteDeviceMountpointsNotFound{}
}

/*DeleteDeviceMountpointsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type DeleteDeviceMountpointsNotFound struct {
}

func (o *DeleteDeviceMountpointsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device/mountpoints/{id}/][%d] deleteDeviceMountpointsNotFound ", 404)
}

func (o *DeleteDeviceMountpointsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceMountpointsMethodNotAllowed creates a DeleteDeviceMountpointsMethodNotAllowed with default headers values
func NewDeleteDeviceMountpointsMethodNotAllowed() *DeleteDeviceMountpointsMethodNotAllowed {
	return &DeleteDeviceMountpointsMethodNotAllowed{}
}

/*DeleteDeviceMountpointsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type DeleteDeviceMountpointsMethodNotAllowed struct {
}

func (o *DeleteDeviceMountpointsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device/mountpoints/{id}/][%d] deleteDeviceMountpointsMethodNotAllowed ", 405)
}

func (o *DeleteDeviceMountpointsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceMountpointsGone creates a DeleteDeviceMountpointsGone with default headers values
func NewDeleteDeviceMountpointsGone() *DeleteDeviceMountpointsGone {
	return &DeleteDeviceMountpointsGone{}
}

/*DeleteDeviceMountpointsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type DeleteDeviceMountpointsGone struct {
}

func (o *DeleteDeviceMountpointsGone) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device/mountpoints/{id}/][%d] deleteDeviceMountpointsGone ", 410)
}

func (o *DeleteDeviceMountpointsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceMountpointsInternalServerError creates a DeleteDeviceMountpointsInternalServerError with default headers values
func NewDeleteDeviceMountpointsInternalServerError() *DeleteDeviceMountpointsInternalServerError {
	return &DeleteDeviceMountpointsInternalServerError{}
}

/*DeleteDeviceMountpointsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type DeleteDeviceMountpointsInternalServerError struct {
}

func (o *DeleteDeviceMountpointsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device/mountpoints/{id}/][%d] deleteDeviceMountpointsInternalServerError ", 500)
}

func (o *DeleteDeviceMountpointsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceMountpointsServiceUnavailable creates a DeleteDeviceMountpointsServiceUnavailable with default headers values
func NewDeleteDeviceMountpointsServiceUnavailable() *DeleteDeviceMountpointsServiceUnavailable {
	return &DeleteDeviceMountpointsServiceUnavailable{}
}

/*DeleteDeviceMountpointsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type DeleteDeviceMountpointsServiceUnavailable struct {
}

func (o *DeleteDeviceMountpointsServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device/mountpoints/{id}/][%d] deleteDeviceMountpointsServiceUnavailable ", 503)
}

func (o *DeleteDeviceMountpointsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeleteDeviceMountpointsOKBody delete device mountpoints o k body
swagger:model DeleteDeviceMountpointsOKBody
*/
type DeleteDeviceMountpointsOKBody struct {

	// deleted
	Deleted interface{} `json:"deleted,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`
}

// Validate validates this delete device mountpoints o k body
func (o *DeleteDeviceMountpointsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteDeviceMountpointsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteDeviceMountpointsOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteDeviceMountpointsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
