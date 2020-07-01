// Code generated by go-swagger; DO NOT EDIT.

package application_components

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

	"github.com/mjpitz/libdevice42/models"
)

// GetAppcompsAppcompIDReader is a Reader for the GetAppcompsAppcompID structure.
type GetAppcompsAppcompIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppcompsAppcompIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppcompsAppcompIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAppcompsAppcompIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAppcompsAppcompIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAppcompsAppcompIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAppcompsAppcompIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetAppcompsAppcompIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetAppcompsAppcompIDGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAppcompsAppcompIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAppcompsAppcompIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAppcompsAppcompIDOK creates a GetAppcompsAppcompIDOK with default headers values
func NewGetAppcompsAppcompIDOK() *GetAppcompsAppcompIDOK {
	return &GetAppcompsAppcompIDOK{}
}

/*GetAppcompsAppcompIDOK handles this case with default header values.

The above command returns results like this:
*/
type GetAppcompsAppcompIDOK struct {
	Payload *GetAppcompsAppcompIDOKBody
}

func (o *GetAppcompsAppcompIDOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/appcomps/{appcomp_id}/][%d] getAppcompsAppcompIdOK  %+v", 200, o.Payload)
}

func (o *GetAppcompsAppcompIDOK) GetPayload() *GetAppcompsAppcompIDOKBody {
	return o.Payload
}

func (o *GetAppcompsAppcompIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAppcompsAppcompIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppcompsAppcompIDBadRequest creates a GetAppcompsAppcompIDBadRequest with default headers values
func NewGetAppcompsAppcompIDBadRequest() *GetAppcompsAppcompIDBadRequest {
	return &GetAppcompsAppcompIDBadRequest{}
}

/*GetAppcompsAppcompIDBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetAppcompsAppcompIDBadRequest struct {
}

func (o *GetAppcompsAppcompIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/appcomps/{appcomp_id}/][%d] getAppcompsAppcompIdBadRequest ", 400)
}

func (o *GetAppcompsAppcompIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAppcompsAppcompIDUnauthorized creates a GetAppcompsAppcompIDUnauthorized with default headers values
func NewGetAppcompsAppcompIDUnauthorized() *GetAppcompsAppcompIDUnauthorized {
	return &GetAppcompsAppcompIDUnauthorized{}
}

/*GetAppcompsAppcompIDUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetAppcompsAppcompIDUnauthorized struct {
}

func (o *GetAppcompsAppcompIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/appcomps/{appcomp_id}/][%d] getAppcompsAppcompIdUnauthorized ", 401)
}

func (o *GetAppcompsAppcompIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAppcompsAppcompIDForbidden creates a GetAppcompsAppcompIDForbidden with default headers values
func NewGetAppcompsAppcompIDForbidden() *GetAppcompsAppcompIDForbidden {
	return &GetAppcompsAppcompIDForbidden{}
}

/*GetAppcompsAppcompIDForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetAppcompsAppcompIDForbidden struct {
}

func (o *GetAppcompsAppcompIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/appcomps/{appcomp_id}/][%d] getAppcompsAppcompIdForbidden ", 403)
}

func (o *GetAppcompsAppcompIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAppcompsAppcompIDNotFound creates a GetAppcompsAppcompIDNotFound with default headers values
func NewGetAppcompsAppcompIDNotFound() *GetAppcompsAppcompIDNotFound {
	return &GetAppcompsAppcompIDNotFound{}
}

/*GetAppcompsAppcompIDNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetAppcompsAppcompIDNotFound struct {
}

func (o *GetAppcompsAppcompIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/appcomps/{appcomp_id}/][%d] getAppcompsAppcompIdNotFound ", 404)
}

func (o *GetAppcompsAppcompIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAppcompsAppcompIDMethodNotAllowed creates a GetAppcompsAppcompIDMethodNotAllowed with default headers values
func NewGetAppcompsAppcompIDMethodNotAllowed() *GetAppcompsAppcompIDMethodNotAllowed {
	return &GetAppcompsAppcompIDMethodNotAllowed{}
}

/*GetAppcompsAppcompIDMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetAppcompsAppcompIDMethodNotAllowed struct {
}

func (o *GetAppcompsAppcompIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/appcomps/{appcomp_id}/][%d] getAppcompsAppcompIdMethodNotAllowed ", 405)
}

func (o *GetAppcompsAppcompIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAppcompsAppcompIDGone creates a GetAppcompsAppcompIDGone with default headers values
func NewGetAppcompsAppcompIDGone() *GetAppcompsAppcompIDGone {
	return &GetAppcompsAppcompIDGone{}
}

/*GetAppcompsAppcompIDGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetAppcompsAppcompIDGone struct {
}

func (o *GetAppcompsAppcompIDGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/appcomps/{appcomp_id}/][%d] getAppcompsAppcompIdGone ", 410)
}

func (o *GetAppcompsAppcompIDGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAppcompsAppcompIDInternalServerError creates a GetAppcompsAppcompIDInternalServerError with default headers values
func NewGetAppcompsAppcompIDInternalServerError() *GetAppcompsAppcompIDInternalServerError {
	return &GetAppcompsAppcompIDInternalServerError{}
}

/*GetAppcompsAppcompIDInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetAppcompsAppcompIDInternalServerError struct {
}

func (o *GetAppcompsAppcompIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/appcomps/{appcomp_id}/][%d] getAppcompsAppcompIdInternalServerError ", 500)
}

func (o *GetAppcompsAppcompIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAppcompsAppcompIDServiceUnavailable creates a GetAppcompsAppcompIDServiceUnavailable with default headers values
func NewGetAppcompsAppcompIDServiceUnavailable() *GetAppcompsAppcompIDServiceUnavailable {
	return &GetAppcompsAppcompIDServiceUnavailable{}
}

/*GetAppcompsAppcompIDServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetAppcompsAppcompIDServiceUnavailable struct {
}

func (o *GetAppcompsAppcompIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/appcomps/{appcomp_id}/][%d] getAppcompsAppcompIdServiceUnavailable ", 503)
}

func (o *GetAppcompsAppcompIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetAppcompsAppcompIDOKBody get appcomps appcomp ID o k body
swagger:model GetAppcompsAppcompIDOKBody
*/
type GetAppcompsAppcompIDOKBody struct {

	// appcomp id
	AppcompID interface{} `json:"appcomp_id,omitempty"`

	// custom fields
	CustomFields []*models.AppcompsCustomFields `json:"custom_fields"`

	// dependent
	Dependent interface{} `json:"dependent,omitempty"`

	// depends on
	DependsOn interface{} `json:"depends_on,omitempty"`

	// device
	Device interface{} `json:"device,omitempty"`

	// group owner
	GroupOwner interface{} `json:"group_owner,omitempty"`

	// groups affected
	GroupsAffected interface{} `json:"groups_affected,omitempty"`

	// name
	Name interface{} `json:"name,omitempty"`

	// notes
	Notes interface{} `json:"notes,omitempty"`

	// what
	What interface{} `json:"what,omitempty"`
}

// Validate validates this get appcomps appcomp ID o k body
func (o *GetAppcompsAppcompIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCustomFields(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAppcompsAppcompIDOKBody) validateCustomFields(formats strfmt.Registry) error {

	if swag.IsZero(o.CustomFields) { // not required
		return nil
	}

	for i := 0; i < len(o.CustomFields); i++ {
		if swag.IsZero(o.CustomFields[i]) { // not required
			continue
		}

		if o.CustomFields[i] != nil {
			if err := o.CustomFields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getAppcompsAppcompIdOK" + "." + "custom_fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAppcompsAppcompIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAppcompsAppcompIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetAppcompsAppcompIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
