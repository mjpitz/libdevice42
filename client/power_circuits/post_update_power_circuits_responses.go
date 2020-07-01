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

// PostUpdatePowerCircuitsReader is a Reader for the PostUpdatePowerCircuits structure.
type PostUpdatePowerCircuitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUpdatePowerCircuitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUpdatePowerCircuitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUpdatePowerCircuitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUpdatePowerCircuitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUpdatePowerCircuitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUpdatePowerCircuitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostUpdatePowerCircuitsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostUpdatePowerCircuitsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUpdatePowerCircuitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUpdatePowerCircuitsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostUpdatePowerCircuitsOK creates a PostUpdatePowerCircuitsOK with default headers values
func NewPostUpdatePowerCircuitsOK() *PostUpdatePowerCircuitsOK {
	return &PostUpdatePowerCircuitsOK{}
}

/*PostUpdatePowerCircuitsOK handles this case with default header values.

The above command returns results like this:
*/
type PostUpdatePowerCircuitsOK struct {
	Payload *PostUpdatePowerCircuitsOKBody
}

func (o *PostUpdatePowerCircuitsOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/powercircuits/][%d] postUpdatePowerCircuitsOK  %+v", 200, o.Payload)
}

func (o *PostUpdatePowerCircuitsOK) GetPayload() *PostUpdatePowerCircuitsOKBody {
	return o.Payload
}

func (o *PostUpdatePowerCircuitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostUpdatePowerCircuitsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUpdatePowerCircuitsBadRequest creates a PostUpdatePowerCircuitsBadRequest with default headers values
func NewPostUpdatePowerCircuitsBadRequest() *PostUpdatePowerCircuitsBadRequest {
	return &PostUpdatePowerCircuitsBadRequest{}
}

/*PostUpdatePowerCircuitsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostUpdatePowerCircuitsBadRequest struct {
}

func (o *PostUpdatePowerCircuitsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/powercircuits/][%d] postUpdatePowerCircuitsBadRequest ", 400)
}

func (o *PostUpdatePowerCircuitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdatePowerCircuitsUnauthorized creates a PostUpdatePowerCircuitsUnauthorized with default headers values
func NewPostUpdatePowerCircuitsUnauthorized() *PostUpdatePowerCircuitsUnauthorized {
	return &PostUpdatePowerCircuitsUnauthorized{}
}

/*PostUpdatePowerCircuitsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostUpdatePowerCircuitsUnauthorized struct {
}

func (o *PostUpdatePowerCircuitsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/powercircuits/][%d] postUpdatePowerCircuitsUnauthorized ", 401)
}

func (o *PostUpdatePowerCircuitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdatePowerCircuitsForbidden creates a PostUpdatePowerCircuitsForbidden with default headers values
func NewPostUpdatePowerCircuitsForbidden() *PostUpdatePowerCircuitsForbidden {
	return &PostUpdatePowerCircuitsForbidden{}
}

/*PostUpdatePowerCircuitsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostUpdatePowerCircuitsForbidden struct {
}

func (o *PostUpdatePowerCircuitsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/powercircuits/][%d] postUpdatePowerCircuitsForbidden ", 403)
}

func (o *PostUpdatePowerCircuitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdatePowerCircuitsNotFound creates a PostUpdatePowerCircuitsNotFound with default headers values
func NewPostUpdatePowerCircuitsNotFound() *PostUpdatePowerCircuitsNotFound {
	return &PostUpdatePowerCircuitsNotFound{}
}

/*PostUpdatePowerCircuitsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostUpdatePowerCircuitsNotFound struct {
}

func (o *PostUpdatePowerCircuitsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/powercircuits/][%d] postUpdatePowerCircuitsNotFound ", 404)
}

func (o *PostUpdatePowerCircuitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdatePowerCircuitsMethodNotAllowed creates a PostUpdatePowerCircuitsMethodNotAllowed with default headers values
func NewPostUpdatePowerCircuitsMethodNotAllowed() *PostUpdatePowerCircuitsMethodNotAllowed {
	return &PostUpdatePowerCircuitsMethodNotAllowed{}
}

/*PostUpdatePowerCircuitsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostUpdatePowerCircuitsMethodNotAllowed struct {
}

func (o *PostUpdatePowerCircuitsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/powercircuits/][%d] postUpdatePowerCircuitsMethodNotAllowed ", 405)
}

func (o *PostUpdatePowerCircuitsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdatePowerCircuitsGone creates a PostUpdatePowerCircuitsGone with default headers values
func NewPostUpdatePowerCircuitsGone() *PostUpdatePowerCircuitsGone {
	return &PostUpdatePowerCircuitsGone{}
}

/*PostUpdatePowerCircuitsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostUpdatePowerCircuitsGone struct {
}

func (o *PostUpdatePowerCircuitsGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/powercircuits/][%d] postUpdatePowerCircuitsGone ", 410)
}

func (o *PostUpdatePowerCircuitsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdatePowerCircuitsInternalServerError creates a PostUpdatePowerCircuitsInternalServerError with default headers values
func NewPostUpdatePowerCircuitsInternalServerError() *PostUpdatePowerCircuitsInternalServerError {
	return &PostUpdatePowerCircuitsInternalServerError{}
}

/*PostUpdatePowerCircuitsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type PostUpdatePowerCircuitsInternalServerError struct {
}

func (o *PostUpdatePowerCircuitsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/powercircuits/][%d] postUpdatePowerCircuitsInternalServerError ", 500)
}

func (o *PostUpdatePowerCircuitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdatePowerCircuitsServiceUnavailable creates a PostUpdatePowerCircuitsServiceUnavailable with default headers values
func NewPostUpdatePowerCircuitsServiceUnavailable() *PostUpdatePowerCircuitsServiceUnavailable {
	return &PostUpdatePowerCircuitsServiceUnavailable{}
}

/*PostUpdatePowerCircuitsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostUpdatePowerCircuitsServiceUnavailable struct {
}

func (o *PostUpdatePowerCircuitsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/powercircuits/][%d] postUpdatePowerCircuitsServiceUnavailable ", 503)
}

func (o *PostUpdatePowerCircuitsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostUpdatePowerCircuitsOKBody post update power circuits o k body
swagger:model PostUpdatePowerCircuitsOKBody
*/
type PostUpdatePowerCircuitsOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post update power circuits o k body
func (o *PostUpdatePowerCircuitsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostUpdatePowerCircuitsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostUpdatePowerCircuitsOKBody) UnmarshalBinary(b []byte) error {
	var res PostUpdatePowerCircuitsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
