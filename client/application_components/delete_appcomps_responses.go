// Code generated by go-swagger; DO NOT EDIT.

package application_components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteAppcompsReader is a Reader for the DeleteAppcomps structure.
type DeleteAppcompsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAppcompsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAppcompsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAppcompsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteAppcompsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAppcompsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAppcompsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeleteAppcompsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewDeleteAppcompsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAppcompsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteAppcompsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAppcompsOK creates a DeleteAppcompsOK with default headers values
func NewDeleteAppcompsOK() *DeleteAppcompsOK {
	return &DeleteAppcompsOK{}
}

/*DeleteAppcompsOK handles this case with default header values.

The above command returns results like this:
*/
type DeleteAppcompsOK struct {
	Payload *DeleteAppcompsOKBody
}

func (o *DeleteAppcompsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/appcomps/{appcomp_id}/][%d] deleteAppcompsOK  %+v", 200, o.Payload)
}

func (o *DeleteAppcompsOK) GetPayload() *DeleteAppcompsOKBody {
	return o.Payload
}

func (o *DeleteAppcompsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteAppcompsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAppcompsBadRequest creates a DeleteAppcompsBadRequest with default headers values
func NewDeleteAppcompsBadRequest() *DeleteAppcompsBadRequest {
	return &DeleteAppcompsBadRequest{}
}

/*DeleteAppcompsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type DeleteAppcompsBadRequest struct {
}

func (o *DeleteAppcompsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/appcomps/{appcomp_id}/][%d] deleteAppcompsBadRequest ", 400)
}

func (o *DeleteAppcompsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAppcompsUnauthorized creates a DeleteAppcompsUnauthorized with default headers values
func NewDeleteAppcompsUnauthorized() *DeleteAppcompsUnauthorized {
	return &DeleteAppcompsUnauthorized{}
}

/*DeleteAppcompsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type DeleteAppcompsUnauthorized struct {
}

func (o *DeleteAppcompsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/appcomps/{appcomp_id}/][%d] deleteAppcompsUnauthorized ", 401)
}

func (o *DeleteAppcompsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAppcompsForbidden creates a DeleteAppcompsForbidden with default headers values
func NewDeleteAppcompsForbidden() *DeleteAppcompsForbidden {
	return &DeleteAppcompsForbidden{}
}

/*DeleteAppcompsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type DeleteAppcompsForbidden struct {
}

func (o *DeleteAppcompsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/appcomps/{appcomp_id}/][%d] deleteAppcompsForbidden ", 403)
}

func (o *DeleteAppcompsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAppcompsNotFound creates a DeleteAppcompsNotFound with default headers values
func NewDeleteAppcompsNotFound() *DeleteAppcompsNotFound {
	return &DeleteAppcompsNotFound{}
}

/*DeleteAppcompsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type DeleteAppcompsNotFound struct {
}

func (o *DeleteAppcompsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/appcomps/{appcomp_id}/][%d] deleteAppcompsNotFound ", 404)
}

func (o *DeleteAppcompsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAppcompsMethodNotAllowed creates a DeleteAppcompsMethodNotAllowed with default headers values
func NewDeleteAppcompsMethodNotAllowed() *DeleteAppcompsMethodNotAllowed {
	return &DeleteAppcompsMethodNotAllowed{}
}

/*DeleteAppcompsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type DeleteAppcompsMethodNotAllowed struct {
}

func (o *DeleteAppcompsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/appcomps/{appcomp_id}/][%d] deleteAppcompsMethodNotAllowed ", 405)
}

func (o *DeleteAppcompsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAppcompsGone creates a DeleteAppcompsGone with default headers values
func NewDeleteAppcompsGone() *DeleteAppcompsGone {
	return &DeleteAppcompsGone{}
}

/*DeleteAppcompsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type DeleteAppcompsGone struct {
}

func (o *DeleteAppcompsGone) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/appcomps/{appcomp_id}/][%d] deleteAppcompsGone ", 410)
}

func (o *DeleteAppcompsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAppcompsInternalServerError creates a DeleteAppcompsInternalServerError with default headers values
func NewDeleteAppcompsInternalServerError() *DeleteAppcompsInternalServerError {
	return &DeleteAppcompsInternalServerError{}
}

/*DeleteAppcompsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type DeleteAppcompsInternalServerError struct {
}

func (o *DeleteAppcompsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/appcomps/{appcomp_id}/][%d] deleteAppcompsInternalServerError ", 500)
}

func (o *DeleteAppcompsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAppcompsServiceUnavailable creates a DeleteAppcompsServiceUnavailable with default headers values
func NewDeleteAppcompsServiceUnavailable() *DeleteAppcompsServiceUnavailable {
	return &DeleteAppcompsServiceUnavailable{}
}

/*DeleteAppcompsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type DeleteAppcompsServiceUnavailable struct {
}

func (o *DeleteAppcompsServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/appcomps/{appcomp_id}/][%d] deleteAppcompsServiceUnavailable ", 503)
}

func (o *DeleteAppcompsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeleteAppcompsOKBody delete appcomps o k body
swagger:model DeleteAppcompsOKBody
*/
type DeleteAppcompsOKBody struct {

	// deleted
	Deleted interface{} `json:"deleted,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`
}

// Validate validates this delete appcomps o k body
func (o *DeleteAppcompsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteAppcompsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteAppcompsOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteAppcompsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
