// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/libdevice42/models"
)

// GetDevicesCustomerIDReader is a Reader for the GetDevicesCustomerID structure.
type GetDevicesCustomerIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesCustomerIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesCustomerIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesCustomerIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesCustomerIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesCustomerIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesCustomerIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetDevicesCustomerIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetDevicesCustomerIDGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesCustomerIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetDevicesCustomerIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDevicesCustomerIDOK creates a GetDevicesCustomerIDOK with default headers values
func NewGetDevicesCustomerIDOK() *GetDevicesCustomerIDOK {
	return &GetDevicesCustomerIDOK{}
}

/*GetDevicesCustomerIDOK handles this case with default header values.

The above command returns results like this:
*/
type GetDevicesCustomerIDOK struct {
	Payload *models.DevicesCustomerID
}

func (o *GetDevicesCustomerIDOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/customer_id/{customer-id}/][%d] getDevicesCustomerIdOK  %+v", 200, o.Payload)
}

func (o *GetDevicesCustomerIDOK) GetPayload() *models.DevicesCustomerID {
	return o.Payload
}

func (o *GetDevicesCustomerIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DevicesCustomerID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesCustomerIDBadRequest creates a GetDevicesCustomerIDBadRequest with default headers values
func NewGetDevicesCustomerIDBadRequest() *GetDevicesCustomerIDBadRequest {
	return &GetDevicesCustomerIDBadRequest{}
}

/*GetDevicesCustomerIDBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetDevicesCustomerIDBadRequest struct {
}

func (o *GetDevicesCustomerIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/customer_id/{customer-id}/][%d] getDevicesCustomerIdBadRequest ", 400)
}

func (o *GetDevicesCustomerIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesCustomerIDUnauthorized creates a GetDevicesCustomerIDUnauthorized with default headers values
func NewGetDevicesCustomerIDUnauthorized() *GetDevicesCustomerIDUnauthorized {
	return &GetDevicesCustomerIDUnauthorized{}
}

/*GetDevicesCustomerIDUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetDevicesCustomerIDUnauthorized struct {
}

func (o *GetDevicesCustomerIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/customer_id/{customer-id}/][%d] getDevicesCustomerIdUnauthorized ", 401)
}

func (o *GetDevicesCustomerIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesCustomerIDForbidden creates a GetDevicesCustomerIDForbidden with default headers values
func NewGetDevicesCustomerIDForbidden() *GetDevicesCustomerIDForbidden {
	return &GetDevicesCustomerIDForbidden{}
}

/*GetDevicesCustomerIDForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetDevicesCustomerIDForbidden struct {
}

func (o *GetDevicesCustomerIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/customer_id/{customer-id}/][%d] getDevicesCustomerIdForbidden ", 403)
}

func (o *GetDevicesCustomerIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesCustomerIDNotFound creates a GetDevicesCustomerIDNotFound with default headers values
func NewGetDevicesCustomerIDNotFound() *GetDevicesCustomerIDNotFound {
	return &GetDevicesCustomerIDNotFound{}
}

/*GetDevicesCustomerIDNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetDevicesCustomerIDNotFound struct {
}

func (o *GetDevicesCustomerIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/customer_id/{customer-id}/][%d] getDevicesCustomerIdNotFound ", 404)
}

func (o *GetDevicesCustomerIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesCustomerIDMethodNotAllowed creates a GetDevicesCustomerIDMethodNotAllowed with default headers values
func NewGetDevicesCustomerIDMethodNotAllowed() *GetDevicesCustomerIDMethodNotAllowed {
	return &GetDevicesCustomerIDMethodNotAllowed{}
}

/*GetDevicesCustomerIDMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetDevicesCustomerIDMethodNotAllowed struct {
}

func (o *GetDevicesCustomerIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/customer_id/{customer-id}/][%d] getDevicesCustomerIdMethodNotAllowed ", 405)
}

func (o *GetDevicesCustomerIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesCustomerIDGone creates a GetDevicesCustomerIDGone with default headers values
func NewGetDevicesCustomerIDGone() *GetDevicesCustomerIDGone {
	return &GetDevicesCustomerIDGone{}
}

/*GetDevicesCustomerIDGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetDevicesCustomerIDGone struct {
}

func (o *GetDevicesCustomerIDGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/customer_id/{customer-id}/][%d] getDevicesCustomerIdGone ", 410)
}

func (o *GetDevicesCustomerIDGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesCustomerIDInternalServerError creates a GetDevicesCustomerIDInternalServerError with default headers values
func NewGetDevicesCustomerIDInternalServerError() *GetDevicesCustomerIDInternalServerError {
	return &GetDevicesCustomerIDInternalServerError{}
}

/*GetDevicesCustomerIDInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetDevicesCustomerIDInternalServerError struct {
}

func (o *GetDevicesCustomerIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/customer_id/{customer-id}/][%d] getDevicesCustomerIdInternalServerError ", 500)
}

func (o *GetDevicesCustomerIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesCustomerIDServiceUnavailable creates a GetDevicesCustomerIDServiceUnavailable with default headers values
func NewGetDevicesCustomerIDServiceUnavailable() *GetDevicesCustomerIDServiceUnavailable {
	return &GetDevicesCustomerIDServiceUnavailable{}
}

/*GetDevicesCustomerIDServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetDevicesCustomerIDServiceUnavailable struct {
}

func (o *GetDevicesCustomerIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/customer_id/{customer-id}/][%d] getDevicesCustomerIdServiceUnavailable ", 503)
}

func (o *GetDevicesCustomerIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
