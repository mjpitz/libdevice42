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

// GetAdmingroupsReader is a Reader for the GetAdmingroups structure.
type GetAdmingroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdmingroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAdmingroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAdmingroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAdmingroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAdmingroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAdmingroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetAdmingroupsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetAdmingroupsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAdmingroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAdmingroupsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAdmingroupsOK creates a GetAdmingroupsOK with default headers values
func NewGetAdmingroupsOK() *GetAdmingroupsOK {
	return &GetAdmingroupsOK{}
}

/*GetAdmingroupsOK handles this case with default header values.

The above command returns results like this:
*/
type GetAdmingroupsOK struct {
	Payload *GetAdmingroupsOKBody
}

func (o *GetAdmingroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/admingroups/][%d] getAdmingroupsOK  %+v", 200, o.Payload)
}

func (o *GetAdmingroupsOK) GetPayload() *GetAdmingroupsOKBody {
	return o.Payload
}

func (o *GetAdmingroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAdmingroupsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdmingroupsBadRequest creates a GetAdmingroupsBadRequest with default headers values
func NewGetAdmingroupsBadRequest() *GetAdmingroupsBadRequest {
	return &GetAdmingroupsBadRequest{}
}

/*GetAdmingroupsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetAdmingroupsBadRequest struct {
}

func (o *GetAdmingroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/admingroups/][%d] getAdmingroupsBadRequest ", 400)
}

func (o *GetAdmingroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdmingroupsUnauthorized creates a GetAdmingroupsUnauthorized with default headers values
func NewGetAdmingroupsUnauthorized() *GetAdmingroupsUnauthorized {
	return &GetAdmingroupsUnauthorized{}
}

/*GetAdmingroupsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetAdmingroupsUnauthorized struct {
}

func (o *GetAdmingroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/admingroups/][%d] getAdmingroupsUnauthorized ", 401)
}

func (o *GetAdmingroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdmingroupsForbidden creates a GetAdmingroupsForbidden with default headers values
func NewGetAdmingroupsForbidden() *GetAdmingroupsForbidden {
	return &GetAdmingroupsForbidden{}
}

/*GetAdmingroupsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetAdmingroupsForbidden struct {
}

func (o *GetAdmingroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/admingroups/][%d] getAdmingroupsForbidden ", 403)
}

func (o *GetAdmingroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdmingroupsNotFound creates a GetAdmingroupsNotFound with default headers values
func NewGetAdmingroupsNotFound() *GetAdmingroupsNotFound {
	return &GetAdmingroupsNotFound{}
}

/*GetAdmingroupsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetAdmingroupsNotFound struct {
}

func (o *GetAdmingroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/admingroups/][%d] getAdmingroupsNotFound ", 404)
}

func (o *GetAdmingroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdmingroupsMethodNotAllowed creates a GetAdmingroupsMethodNotAllowed with default headers values
func NewGetAdmingroupsMethodNotAllowed() *GetAdmingroupsMethodNotAllowed {
	return &GetAdmingroupsMethodNotAllowed{}
}

/*GetAdmingroupsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetAdmingroupsMethodNotAllowed struct {
}

func (o *GetAdmingroupsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/admingroups/][%d] getAdmingroupsMethodNotAllowed ", 405)
}

func (o *GetAdmingroupsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdmingroupsGone creates a GetAdmingroupsGone with default headers values
func NewGetAdmingroupsGone() *GetAdmingroupsGone {
	return &GetAdmingroupsGone{}
}

/*GetAdmingroupsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetAdmingroupsGone struct {
}

func (o *GetAdmingroupsGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/admingroups/][%d] getAdmingroupsGone ", 410)
}

func (o *GetAdmingroupsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdmingroupsInternalServerError creates a GetAdmingroupsInternalServerError with default headers values
func NewGetAdmingroupsInternalServerError() *GetAdmingroupsInternalServerError {
	return &GetAdmingroupsInternalServerError{}
}

/*GetAdmingroupsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetAdmingroupsInternalServerError struct {
}

func (o *GetAdmingroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/admingroups/][%d] getAdmingroupsInternalServerError ", 500)
}

func (o *GetAdmingroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdmingroupsServiceUnavailable creates a GetAdmingroupsServiceUnavailable with default headers values
func NewGetAdmingroupsServiceUnavailable() *GetAdmingroupsServiceUnavailable {
	return &GetAdmingroupsServiceUnavailable{}
}

/*GetAdmingroupsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetAdmingroupsServiceUnavailable struct {
}

func (o *GetAdmingroupsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/admingroups/][%d] getAdmingroupsServiceUnavailable ", 503)
}

func (o *GetAdmingroupsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*AdminusersItems0 adminusers items0
swagger:model AdminusersItems0
*/
type AdminusersItems0 struct {

	// groups
	Groups interface{} `json:"groups,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`

	// username
	Username interface{} `json:"username,omitempty"`
}

// Validate validates this adminusers items0
func (o *AdminusersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AdminusersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AdminusersItems0) UnmarshalBinary(b []byte) error {
	var res AdminusersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAdmingroupsOKBody get admingroups o k body
swagger:model GetAdmingroupsOKBody
*/
type GetAdmingroupsOKBody struct {

	// adminusers
	Adminusers []*AdminusersItems0 `json:"adminusers"`
}

// Validate validates this get admingroups o k body
func (o *GetAdmingroupsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAdminusers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAdmingroupsOKBody) validateAdminusers(formats strfmt.Registry) error {

	if swag.IsZero(o.Adminusers) { // not required
		return nil
	}

	for i := 0; i < len(o.Adminusers); i++ {
		if swag.IsZero(o.Adminusers[i]) { // not required
			continue
		}

		if o.Adminusers[i] != nil {
			if err := o.Adminusers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getAdmingroupsOK" + "." + "adminusers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAdmingroupsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAdmingroupsOKBody) UnmarshalBinary(b []byte) error {
	var res GetAdmingroupsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
