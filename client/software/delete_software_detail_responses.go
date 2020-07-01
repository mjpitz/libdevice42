// Code generated by go-swagger; DO NOT EDIT.

package software

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteSoftwareDetailReader is a Reader for the DeleteSoftwareDetail structure.
type DeleteSoftwareDetailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSoftwareDetailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSoftwareDetailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSoftwareDetailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteSoftwareDetailUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteSoftwareDetailForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSoftwareDetailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeleteSoftwareDetailMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewDeleteSoftwareDetailGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSoftwareDetailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteSoftwareDetailServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSoftwareDetailOK creates a DeleteSoftwareDetailOK with default headers values
func NewDeleteSoftwareDetailOK() *DeleteSoftwareDetailOK {
	return &DeleteSoftwareDetailOK{}
}

/*DeleteSoftwareDetailOK handles this case with default header values.

The above command returns results like this:
*/
type DeleteSoftwareDetailOK struct {
	Payload *DeleteSoftwareDetailOKBody
}

func (o *DeleteSoftwareDetailOK) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/software_details/{id}/][%d] deleteSoftwareDetailOK  %+v", 200, o.Payload)
}

func (o *DeleteSoftwareDetailOK) GetPayload() *DeleteSoftwareDetailOKBody {
	return o.Payload
}

func (o *DeleteSoftwareDetailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteSoftwareDetailOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSoftwareDetailBadRequest creates a DeleteSoftwareDetailBadRequest with default headers values
func NewDeleteSoftwareDetailBadRequest() *DeleteSoftwareDetailBadRequest {
	return &DeleteSoftwareDetailBadRequest{}
}

/*DeleteSoftwareDetailBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type DeleteSoftwareDetailBadRequest struct {
}

func (o *DeleteSoftwareDetailBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/software_details/{id}/][%d] deleteSoftwareDetailBadRequest ", 400)
}

func (o *DeleteSoftwareDetailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSoftwareDetailUnauthorized creates a DeleteSoftwareDetailUnauthorized with default headers values
func NewDeleteSoftwareDetailUnauthorized() *DeleteSoftwareDetailUnauthorized {
	return &DeleteSoftwareDetailUnauthorized{}
}

/*DeleteSoftwareDetailUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type DeleteSoftwareDetailUnauthorized struct {
}

func (o *DeleteSoftwareDetailUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/software_details/{id}/][%d] deleteSoftwareDetailUnauthorized ", 401)
}

func (o *DeleteSoftwareDetailUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSoftwareDetailForbidden creates a DeleteSoftwareDetailForbidden with default headers values
func NewDeleteSoftwareDetailForbidden() *DeleteSoftwareDetailForbidden {
	return &DeleteSoftwareDetailForbidden{}
}

/*DeleteSoftwareDetailForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type DeleteSoftwareDetailForbidden struct {
}

func (o *DeleteSoftwareDetailForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/software_details/{id}/][%d] deleteSoftwareDetailForbidden ", 403)
}

func (o *DeleteSoftwareDetailForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSoftwareDetailNotFound creates a DeleteSoftwareDetailNotFound with default headers values
func NewDeleteSoftwareDetailNotFound() *DeleteSoftwareDetailNotFound {
	return &DeleteSoftwareDetailNotFound{}
}

/*DeleteSoftwareDetailNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type DeleteSoftwareDetailNotFound struct {
}

func (o *DeleteSoftwareDetailNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/software_details/{id}/][%d] deleteSoftwareDetailNotFound ", 404)
}

func (o *DeleteSoftwareDetailNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSoftwareDetailMethodNotAllowed creates a DeleteSoftwareDetailMethodNotAllowed with default headers values
func NewDeleteSoftwareDetailMethodNotAllowed() *DeleteSoftwareDetailMethodNotAllowed {
	return &DeleteSoftwareDetailMethodNotAllowed{}
}

/*DeleteSoftwareDetailMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type DeleteSoftwareDetailMethodNotAllowed struct {
}

func (o *DeleteSoftwareDetailMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/software_details/{id}/][%d] deleteSoftwareDetailMethodNotAllowed ", 405)
}

func (o *DeleteSoftwareDetailMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSoftwareDetailGone creates a DeleteSoftwareDetailGone with default headers values
func NewDeleteSoftwareDetailGone() *DeleteSoftwareDetailGone {
	return &DeleteSoftwareDetailGone{}
}

/*DeleteSoftwareDetailGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type DeleteSoftwareDetailGone struct {
}

func (o *DeleteSoftwareDetailGone) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/software_details/{id}/][%d] deleteSoftwareDetailGone ", 410)
}

func (o *DeleteSoftwareDetailGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSoftwareDetailInternalServerError creates a DeleteSoftwareDetailInternalServerError with default headers values
func NewDeleteSoftwareDetailInternalServerError() *DeleteSoftwareDetailInternalServerError {
	return &DeleteSoftwareDetailInternalServerError{}
}

/*DeleteSoftwareDetailInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type DeleteSoftwareDetailInternalServerError struct {
}

func (o *DeleteSoftwareDetailInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/software_details/{id}/][%d] deleteSoftwareDetailInternalServerError ", 500)
}

func (o *DeleteSoftwareDetailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSoftwareDetailServiceUnavailable creates a DeleteSoftwareDetailServiceUnavailable with default headers values
func NewDeleteSoftwareDetailServiceUnavailable() *DeleteSoftwareDetailServiceUnavailable {
	return &DeleteSoftwareDetailServiceUnavailable{}
}

/*DeleteSoftwareDetailServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type DeleteSoftwareDetailServiceUnavailable struct {
}

func (o *DeleteSoftwareDetailServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/software_details/{id}/][%d] deleteSoftwareDetailServiceUnavailable ", 503)
}

func (o *DeleteSoftwareDetailServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeleteSoftwareDetailOKBody delete software detail o k body
swagger:model DeleteSoftwareDetailOKBody
*/
type DeleteSoftwareDetailOKBody struct {

	// deleted
	Deleted interface{} `json:"deleted,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`
}

// Validate validates this delete software detail o k body
func (o *DeleteSoftwareDetailOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteSoftwareDetailOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteSoftwareDetailOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteSoftwareDetailOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
