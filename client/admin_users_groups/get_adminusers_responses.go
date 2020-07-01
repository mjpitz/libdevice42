// Code generated by go-swagger; DO NOT EDIT.

package admin_users_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetAdminusersReader is a Reader for the GetAdminusers structure.
type GetAdminusersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdminusersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAdminusersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAdminusersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAdminusersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAdminusersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAdminusersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetAdminusersMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetAdminusersGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAdminusersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAdminusersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAdminusersOK creates a GetAdminusersOK with default headers values
func NewGetAdminusersOK() *GetAdminusersOK {
	return &GetAdminusersOK{}
}

/*GetAdminusersOK handles this case with default header values.

The above command returns results like this:
*/
type GetAdminusersOK struct {
	Payload *GetAdminusersOKBody
}

func (o *GetAdminusersOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/adminusers/][%d] getAdminusersOK  %+v", 200, o.Payload)
}

func (o *GetAdminusersOK) GetPayload() *GetAdminusersOKBody {
	return o.Payload
}

func (o *GetAdminusersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAdminusersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdminusersBadRequest creates a GetAdminusersBadRequest with default headers values
func NewGetAdminusersBadRequest() *GetAdminusersBadRequest {
	return &GetAdminusersBadRequest{}
}

/*GetAdminusersBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetAdminusersBadRequest struct {
}

func (o *GetAdminusersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/adminusers/][%d] getAdminusersBadRequest ", 400)
}

func (o *GetAdminusersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdminusersUnauthorized creates a GetAdminusersUnauthorized with default headers values
func NewGetAdminusersUnauthorized() *GetAdminusersUnauthorized {
	return &GetAdminusersUnauthorized{}
}

/*GetAdminusersUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetAdminusersUnauthorized struct {
}

func (o *GetAdminusersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/adminusers/][%d] getAdminusersUnauthorized ", 401)
}

func (o *GetAdminusersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdminusersForbidden creates a GetAdminusersForbidden with default headers values
func NewGetAdminusersForbidden() *GetAdminusersForbidden {
	return &GetAdminusersForbidden{}
}

/*GetAdminusersForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetAdminusersForbidden struct {
}

func (o *GetAdminusersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/adminusers/][%d] getAdminusersForbidden ", 403)
}

func (o *GetAdminusersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdminusersNotFound creates a GetAdminusersNotFound with default headers values
func NewGetAdminusersNotFound() *GetAdminusersNotFound {
	return &GetAdminusersNotFound{}
}

/*GetAdminusersNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetAdminusersNotFound struct {
}

func (o *GetAdminusersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/adminusers/][%d] getAdminusersNotFound ", 404)
}

func (o *GetAdminusersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdminusersMethodNotAllowed creates a GetAdminusersMethodNotAllowed with default headers values
func NewGetAdminusersMethodNotAllowed() *GetAdminusersMethodNotAllowed {
	return &GetAdminusersMethodNotAllowed{}
}

/*GetAdminusersMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetAdminusersMethodNotAllowed struct {
}

func (o *GetAdminusersMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/adminusers/][%d] getAdminusersMethodNotAllowed ", 405)
}

func (o *GetAdminusersMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdminusersGone creates a GetAdminusersGone with default headers values
func NewGetAdminusersGone() *GetAdminusersGone {
	return &GetAdminusersGone{}
}

/*GetAdminusersGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetAdminusersGone struct {
}

func (o *GetAdminusersGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/adminusers/][%d] getAdminusersGone ", 410)
}

func (o *GetAdminusersGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdminusersInternalServerError creates a GetAdminusersInternalServerError with default headers values
func NewGetAdminusersInternalServerError() *GetAdminusersInternalServerError {
	return &GetAdminusersInternalServerError{}
}

/*GetAdminusersInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetAdminusersInternalServerError struct {
}

func (o *GetAdminusersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/adminusers/][%d] getAdminusersInternalServerError ", 500)
}

func (o *GetAdminusersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdminusersServiceUnavailable creates a GetAdminusersServiceUnavailable with default headers values
func NewGetAdminusersServiceUnavailable() *GetAdminusersServiceUnavailable {
	return &GetAdminusersServiceUnavailable{}
}

/*GetAdminusersServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetAdminusersServiceUnavailable struct {
}

func (o *GetAdminusersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/adminusers/][%d] getAdminusersServiceUnavailable ", 503)
}

func (o *GetAdminusersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*AdmingroupsItems0 admingroups items0
swagger:model AdmingroupsItems0
*/
type AdmingroupsItems0 struct {

	// id
	ID interface{} `json:"id,omitempty"`

	// members
	Members interface{} `json:"members,omitempty"`

	// name
	Name interface{} `json:"name,omitempty"`
}

// Validate validates this admingroups items0
func (o *AdmingroupsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AdmingroupsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AdmingroupsItems0) UnmarshalBinary(b []byte) error {
	var res AdmingroupsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAdminusersOKBody get adminusers o k body
swagger:model GetAdminusersOKBody
*/
type GetAdminusersOKBody struct {

	// admingroups
	Admingroups []*AdmingroupsItems0 `json:"admingroups"`
}

// Validate validates this get adminusers o k body
func (o *GetAdminusersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAdmingroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAdminusersOKBody) validateAdmingroups(formats strfmt.Registry) error {

	if swag.IsZero(o.Admingroups) { // not required
		return nil
	}

	for i := 0; i < len(o.Admingroups); i++ {
		if swag.IsZero(o.Admingroups[i]) { // not required
			continue
		}

		if o.Admingroups[i] != nil {
			if err := o.Admingroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getAdminusersOK" + "." + "admingroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAdminusersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAdminusersOKBody) UnmarshalBinary(b []byte) error {
	var res GetAdminusersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
