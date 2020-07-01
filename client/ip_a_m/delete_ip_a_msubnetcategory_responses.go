// Code generated by go-swagger; DO NOT EDIT.

package ip_a_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteIPAMsubnetcategoryReader is a Reader for the DeleteIPAMsubnetcategory structure.
type DeleteIPAMsubnetcategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIPAMsubnetcategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIPAMsubnetcategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIPAMsubnetcategoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteIPAMsubnetcategoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteIPAMsubnetcategoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteIPAMsubnetcategoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeleteIPAMsubnetcategoryMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewDeleteIPAMsubnetcategoryGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteIPAMsubnetcategoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteIPAMsubnetcategoryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteIPAMsubnetcategoryOK creates a DeleteIPAMsubnetcategoryOK with default headers values
func NewDeleteIPAMsubnetcategoryOK() *DeleteIPAMsubnetcategoryOK {
	return &DeleteIPAMsubnetcategoryOK{}
}

/*DeleteIPAMsubnetcategoryOK handles this case with default header values.

The above command returns JSON structured like this:
*/
type DeleteIPAMsubnetcategoryOK struct {
	Payload *DeleteIPAMsubnetcategoryOKBody
}

func (o *DeleteIPAMsubnetcategoryOK) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/subnet_category/{id}/][%d] deleteIpAMsubnetcategoryOK  %+v", 200, o.Payload)
}

func (o *DeleteIPAMsubnetcategoryOK) GetPayload() *DeleteIPAMsubnetcategoryOKBody {
	return o.Payload
}

func (o *DeleteIPAMsubnetcategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteIPAMsubnetcategoryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPAMsubnetcategoryBadRequest creates a DeleteIPAMsubnetcategoryBadRequest with default headers values
func NewDeleteIPAMsubnetcategoryBadRequest() *DeleteIPAMsubnetcategoryBadRequest {
	return &DeleteIPAMsubnetcategoryBadRequest{}
}

/*DeleteIPAMsubnetcategoryBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type DeleteIPAMsubnetcategoryBadRequest struct {
}

func (o *DeleteIPAMsubnetcategoryBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/subnet_category/{id}/][%d] deleteIpAMsubnetcategoryBadRequest ", 400)
}

func (o *DeleteIPAMsubnetcategoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIPAMsubnetcategoryUnauthorized creates a DeleteIPAMsubnetcategoryUnauthorized with default headers values
func NewDeleteIPAMsubnetcategoryUnauthorized() *DeleteIPAMsubnetcategoryUnauthorized {
	return &DeleteIPAMsubnetcategoryUnauthorized{}
}

/*DeleteIPAMsubnetcategoryUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type DeleteIPAMsubnetcategoryUnauthorized struct {
}

func (o *DeleteIPAMsubnetcategoryUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/subnet_category/{id}/][%d] deleteIpAMsubnetcategoryUnauthorized ", 401)
}

func (o *DeleteIPAMsubnetcategoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIPAMsubnetcategoryForbidden creates a DeleteIPAMsubnetcategoryForbidden with default headers values
func NewDeleteIPAMsubnetcategoryForbidden() *DeleteIPAMsubnetcategoryForbidden {
	return &DeleteIPAMsubnetcategoryForbidden{}
}

/*DeleteIPAMsubnetcategoryForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type DeleteIPAMsubnetcategoryForbidden struct {
}

func (o *DeleteIPAMsubnetcategoryForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/subnet_category/{id}/][%d] deleteIpAMsubnetcategoryForbidden ", 403)
}

func (o *DeleteIPAMsubnetcategoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIPAMsubnetcategoryNotFound creates a DeleteIPAMsubnetcategoryNotFound with default headers values
func NewDeleteIPAMsubnetcategoryNotFound() *DeleteIPAMsubnetcategoryNotFound {
	return &DeleteIPAMsubnetcategoryNotFound{}
}

/*DeleteIPAMsubnetcategoryNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type DeleteIPAMsubnetcategoryNotFound struct {
}

func (o *DeleteIPAMsubnetcategoryNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/subnet_category/{id}/][%d] deleteIpAMsubnetcategoryNotFound ", 404)
}

func (o *DeleteIPAMsubnetcategoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIPAMsubnetcategoryMethodNotAllowed creates a DeleteIPAMsubnetcategoryMethodNotAllowed with default headers values
func NewDeleteIPAMsubnetcategoryMethodNotAllowed() *DeleteIPAMsubnetcategoryMethodNotAllowed {
	return &DeleteIPAMsubnetcategoryMethodNotAllowed{}
}

/*DeleteIPAMsubnetcategoryMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type DeleteIPAMsubnetcategoryMethodNotAllowed struct {
}

func (o *DeleteIPAMsubnetcategoryMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/subnet_category/{id}/][%d] deleteIpAMsubnetcategoryMethodNotAllowed ", 405)
}

func (o *DeleteIPAMsubnetcategoryMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIPAMsubnetcategoryGone creates a DeleteIPAMsubnetcategoryGone with default headers values
func NewDeleteIPAMsubnetcategoryGone() *DeleteIPAMsubnetcategoryGone {
	return &DeleteIPAMsubnetcategoryGone{}
}

/*DeleteIPAMsubnetcategoryGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type DeleteIPAMsubnetcategoryGone struct {
}

func (o *DeleteIPAMsubnetcategoryGone) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/subnet_category/{id}/][%d] deleteIpAMsubnetcategoryGone ", 410)
}

func (o *DeleteIPAMsubnetcategoryGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIPAMsubnetcategoryInternalServerError creates a DeleteIPAMsubnetcategoryInternalServerError with default headers values
func NewDeleteIPAMsubnetcategoryInternalServerError() *DeleteIPAMsubnetcategoryInternalServerError {
	return &DeleteIPAMsubnetcategoryInternalServerError{}
}

/*DeleteIPAMsubnetcategoryInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type DeleteIPAMsubnetcategoryInternalServerError struct {
}

func (o *DeleteIPAMsubnetcategoryInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/subnet_category/{id}/][%d] deleteIpAMsubnetcategoryInternalServerError ", 500)
}

func (o *DeleteIPAMsubnetcategoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIPAMsubnetcategoryServiceUnavailable creates a DeleteIPAMsubnetcategoryServiceUnavailable with default headers values
func NewDeleteIPAMsubnetcategoryServiceUnavailable() *DeleteIPAMsubnetcategoryServiceUnavailable {
	return &DeleteIPAMsubnetcategoryServiceUnavailable{}
}

/*DeleteIPAMsubnetcategoryServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type DeleteIPAMsubnetcategoryServiceUnavailable struct {
}

func (o *DeleteIPAMsubnetcategoryServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/subnet_category/{id}/][%d] deleteIpAMsubnetcategoryServiceUnavailable ", 503)
}

func (o *DeleteIPAMsubnetcategoryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeleteIPAMsubnetcategoryOKBody delete IP a msubnetcategory o k body
swagger:model DeleteIPAMsubnetcategoryOKBody
*/
type DeleteIPAMsubnetcategoryOKBody struct {

	// deleted
	Deleted interface{} `json:"deleted,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`
}

// Validate validates this delete IP a msubnetcategory o k body
func (o *DeleteIPAMsubnetcategoryOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteIPAMsubnetcategoryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteIPAMsubnetcategoryOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteIPAMsubnetcategoryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
