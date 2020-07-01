// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/libdevice42/models"
)

// GetServiceListenerPortsReader is a Reader for the GetServiceListenerPorts structure.
type GetServiceListenerPortsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceListenerPortsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceListenerPortsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetServiceListenerPortsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetServiceListenerPortsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetServiceListenerPortsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServiceListenerPortsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetServiceListenerPortsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetServiceListenerPortsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetServiceListenerPortsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetServiceListenerPortsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceListenerPortsOK creates a GetServiceListenerPortsOK with default headers values
func NewGetServiceListenerPortsOK() *GetServiceListenerPortsOK {
	return &GetServiceListenerPortsOK{}
}

/*GetServiceListenerPortsOK handles this case with default header values.

The above command returns results like this:
*/
type GetServiceListenerPortsOK struct {
	Payload *models.ServiceListenerPorts
}

func (o *GetServiceListenerPortsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.0/service_listener_ports/][%d] getServiceListenerPortsOK  %+v", 200, o.Payload)
}

func (o *GetServiceListenerPortsOK) GetPayload() *models.ServiceListenerPorts {
	return o.Payload
}

func (o *GetServiceListenerPortsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceListenerPorts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceListenerPortsBadRequest creates a GetServiceListenerPortsBadRequest with default headers values
func NewGetServiceListenerPortsBadRequest() *GetServiceListenerPortsBadRequest {
	return &GetServiceListenerPortsBadRequest{}
}

/*GetServiceListenerPortsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetServiceListenerPortsBadRequest struct {
}

func (o *GetServiceListenerPortsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.0/service_listener_ports/][%d] getServiceListenerPortsBadRequest ", 400)
}

func (o *GetServiceListenerPortsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceListenerPortsUnauthorized creates a GetServiceListenerPortsUnauthorized with default headers values
func NewGetServiceListenerPortsUnauthorized() *GetServiceListenerPortsUnauthorized {
	return &GetServiceListenerPortsUnauthorized{}
}

/*GetServiceListenerPortsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetServiceListenerPortsUnauthorized struct {
}

func (o *GetServiceListenerPortsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/2.0/service_listener_ports/][%d] getServiceListenerPortsUnauthorized ", 401)
}

func (o *GetServiceListenerPortsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceListenerPortsForbidden creates a GetServiceListenerPortsForbidden with default headers values
func NewGetServiceListenerPortsForbidden() *GetServiceListenerPortsForbidden {
	return &GetServiceListenerPortsForbidden{}
}

/*GetServiceListenerPortsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetServiceListenerPortsForbidden struct {
}

func (o *GetServiceListenerPortsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/2.0/service_listener_ports/][%d] getServiceListenerPortsForbidden ", 403)
}

func (o *GetServiceListenerPortsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceListenerPortsNotFound creates a GetServiceListenerPortsNotFound with default headers values
func NewGetServiceListenerPortsNotFound() *GetServiceListenerPortsNotFound {
	return &GetServiceListenerPortsNotFound{}
}

/*GetServiceListenerPortsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetServiceListenerPortsNotFound struct {
}

func (o *GetServiceListenerPortsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/2.0/service_listener_ports/][%d] getServiceListenerPortsNotFound ", 404)
}

func (o *GetServiceListenerPortsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceListenerPortsMethodNotAllowed creates a GetServiceListenerPortsMethodNotAllowed with default headers values
func NewGetServiceListenerPortsMethodNotAllowed() *GetServiceListenerPortsMethodNotAllowed {
	return &GetServiceListenerPortsMethodNotAllowed{}
}

/*GetServiceListenerPortsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetServiceListenerPortsMethodNotAllowed struct {
}

func (o *GetServiceListenerPortsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/2.0/service_listener_ports/][%d] getServiceListenerPortsMethodNotAllowed ", 405)
}

func (o *GetServiceListenerPortsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceListenerPortsGone creates a GetServiceListenerPortsGone with default headers values
func NewGetServiceListenerPortsGone() *GetServiceListenerPortsGone {
	return &GetServiceListenerPortsGone{}
}

/*GetServiceListenerPortsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetServiceListenerPortsGone struct {
}

func (o *GetServiceListenerPortsGone) Error() string {
	return fmt.Sprintf("[GET /api/2.0/service_listener_ports/][%d] getServiceListenerPortsGone ", 410)
}

func (o *GetServiceListenerPortsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceListenerPortsInternalServerError creates a GetServiceListenerPortsInternalServerError with default headers values
func NewGetServiceListenerPortsInternalServerError() *GetServiceListenerPortsInternalServerError {
	return &GetServiceListenerPortsInternalServerError{}
}

/*GetServiceListenerPortsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetServiceListenerPortsInternalServerError struct {
}

func (o *GetServiceListenerPortsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/2.0/service_listener_ports/][%d] getServiceListenerPortsInternalServerError ", 500)
}

func (o *GetServiceListenerPortsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceListenerPortsServiceUnavailable creates a GetServiceListenerPortsServiceUnavailable with default headers values
func NewGetServiceListenerPortsServiceUnavailable() *GetServiceListenerPortsServiceUnavailable {
	return &GetServiceListenerPortsServiceUnavailable{}
}

/*GetServiceListenerPortsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetServiceListenerPortsServiceUnavailable struct {
}

func (o *GetServiceListenerPortsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/2.0/service_listener_ports/][%d] getServiceListenerPortsServiceUnavailable ", 503)
}

func (o *GetServiceListenerPortsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
