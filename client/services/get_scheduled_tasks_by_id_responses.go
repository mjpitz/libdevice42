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

// GetScheduledTasksByIDReader is a Reader for the GetScheduledTasksByID structure.
type GetScheduledTasksByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScheduledTasksByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScheduledTasksByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetScheduledTasksByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetScheduledTasksByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScheduledTasksByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScheduledTasksByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetScheduledTasksByIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetScheduledTasksByIDGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScheduledTasksByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetScheduledTasksByIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetScheduledTasksByIDOK creates a GetScheduledTasksByIDOK with default headers values
func NewGetScheduledTasksByIDOK() *GetScheduledTasksByIDOK {
	return &GetScheduledTasksByIDOK{}
}

/*GetScheduledTasksByIDOK handles this case with default header values.

The above command returns results like this:
*/
type GetScheduledTasksByIDOK struct {
	Payload *models.ServiceSchedule
}

func (o *GetScheduledTasksByIDOK) Error() string {
	return fmt.Sprintf("[GET /api/2.0/scheduled_tasks/{id}][%d] getScheduledTasksByIdOK  %+v", 200, o.Payload)
}

func (o *GetScheduledTasksByIDOK) GetPayload() *models.ServiceSchedule {
	return o.Payload
}

func (o *GetScheduledTasksByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScheduledTasksByIDBadRequest creates a GetScheduledTasksByIDBadRequest with default headers values
func NewGetScheduledTasksByIDBadRequest() *GetScheduledTasksByIDBadRequest {
	return &GetScheduledTasksByIDBadRequest{}
}

/*GetScheduledTasksByIDBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetScheduledTasksByIDBadRequest struct {
}

func (o *GetScheduledTasksByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.0/scheduled_tasks/{id}][%d] getScheduledTasksByIdBadRequest ", 400)
}

func (o *GetScheduledTasksByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScheduledTasksByIDUnauthorized creates a GetScheduledTasksByIDUnauthorized with default headers values
func NewGetScheduledTasksByIDUnauthorized() *GetScheduledTasksByIDUnauthorized {
	return &GetScheduledTasksByIDUnauthorized{}
}

/*GetScheduledTasksByIDUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetScheduledTasksByIDUnauthorized struct {
}

func (o *GetScheduledTasksByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/2.0/scheduled_tasks/{id}][%d] getScheduledTasksByIdUnauthorized ", 401)
}

func (o *GetScheduledTasksByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScheduledTasksByIDForbidden creates a GetScheduledTasksByIDForbidden with default headers values
func NewGetScheduledTasksByIDForbidden() *GetScheduledTasksByIDForbidden {
	return &GetScheduledTasksByIDForbidden{}
}

/*GetScheduledTasksByIDForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetScheduledTasksByIDForbidden struct {
}

func (o *GetScheduledTasksByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/2.0/scheduled_tasks/{id}][%d] getScheduledTasksByIdForbidden ", 403)
}

func (o *GetScheduledTasksByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScheduledTasksByIDNotFound creates a GetScheduledTasksByIDNotFound with default headers values
func NewGetScheduledTasksByIDNotFound() *GetScheduledTasksByIDNotFound {
	return &GetScheduledTasksByIDNotFound{}
}

/*GetScheduledTasksByIDNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetScheduledTasksByIDNotFound struct {
}

func (o *GetScheduledTasksByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/2.0/scheduled_tasks/{id}][%d] getScheduledTasksByIdNotFound ", 404)
}

func (o *GetScheduledTasksByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScheduledTasksByIDMethodNotAllowed creates a GetScheduledTasksByIDMethodNotAllowed with default headers values
func NewGetScheduledTasksByIDMethodNotAllowed() *GetScheduledTasksByIDMethodNotAllowed {
	return &GetScheduledTasksByIDMethodNotAllowed{}
}

/*GetScheduledTasksByIDMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetScheduledTasksByIDMethodNotAllowed struct {
}

func (o *GetScheduledTasksByIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/2.0/scheduled_tasks/{id}][%d] getScheduledTasksByIdMethodNotAllowed ", 405)
}

func (o *GetScheduledTasksByIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScheduledTasksByIDGone creates a GetScheduledTasksByIDGone with default headers values
func NewGetScheduledTasksByIDGone() *GetScheduledTasksByIDGone {
	return &GetScheduledTasksByIDGone{}
}

/*GetScheduledTasksByIDGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetScheduledTasksByIDGone struct {
}

func (o *GetScheduledTasksByIDGone) Error() string {
	return fmt.Sprintf("[GET /api/2.0/scheduled_tasks/{id}][%d] getScheduledTasksByIdGone ", 410)
}

func (o *GetScheduledTasksByIDGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScheduledTasksByIDInternalServerError creates a GetScheduledTasksByIDInternalServerError with default headers values
func NewGetScheduledTasksByIDInternalServerError() *GetScheduledTasksByIDInternalServerError {
	return &GetScheduledTasksByIDInternalServerError{}
}

/*GetScheduledTasksByIDInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetScheduledTasksByIDInternalServerError struct {
}

func (o *GetScheduledTasksByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/2.0/scheduled_tasks/{id}][%d] getScheduledTasksByIdInternalServerError ", 500)
}

func (o *GetScheduledTasksByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScheduledTasksByIDServiceUnavailable creates a GetScheduledTasksByIDServiceUnavailable with default headers values
func NewGetScheduledTasksByIDServiceUnavailable() *GetScheduledTasksByIDServiceUnavailable {
	return &GetScheduledTasksByIDServiceUnavailable{}
}

/*GetScheduledTasksByIDServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetScheduledTasksByIDServiceUnavailable struct {
}

func (o *GetScheduledTasksByIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/2.0/scheduled_tasks/{id}][%d] getScheduledTasksByIdServiceUnavailable ", 503)
}

func (o *GetScheduledTasksByIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
