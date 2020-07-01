// Code generated by go-swagger; DO NOT EDIT.

package ip_a_m

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

// GetIPAMvlansReader is a Reader for the GetIPAMvlans structure.
type GetIPAMvlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPAMvlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIPAMvlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIPAMvlansBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIPAMvlansUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIPAMvlansForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIPAMvlansNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetIPAMvlansMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetIPAMvlansGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIPAMvlansInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIPAMvlansServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIPAMvlansOK creates a GetIPAMvlansOK with default headers values
func NewGetIPAMvlansOK() *GetIPAMvlansOK {
	return &GetIPAMvlansOK{}
}

/*GetIPAMvlansOK handles this case with default header values.

The above command returns results like this:
*/
type GetIPAMvlansOK struct {
	Payload *GetIPAMvlansOKBody
}

func (o *GetIPAMvlansOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/vlans/][%d] getIpAMvlansOK  %+v", 200, o.Payload)
}

func (o *GetIPAMvlansOK) GetPayload() *GetIPAMvlansOKBody {
	return o.Payload
}

func (o *GetIPAMvlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetIPAMvlansOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPAMvlansBadRequest creates a GetIPAMvlansBadRequest with default headers values
func NewGetIPAMvlansBadRequest() *GetIPAMvlansBadRequest {
	return &GetIPAMvlansBadRequest{}
}

/*GetIPAMvlansBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetIPAMvlansBadRequest struct {
}

func (o *GetIPAMvlansBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/vlans/][%d] getIpAMvlansBadRequest ", 400)
}

func (o *GetIPAMvlansBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMvlansUnauthorized creates a GetIPAMvlansUnauthorized with default headers values
func NewGetIPAMvlansUnauthorized() *GetIPAMvlansUnauthorized {
	return &GetIPAMvlansUnauthorized{}
}

/*GetIPAMvlansUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetIPAMvlansUnauthorized struct {
}

func (o *GetIPAMvlansUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/vlans/][%d] getIpAMvlansUnauthorized ", 401)
}

func (o *GetIPAMvlansUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMvlansForbidden creates a GetIPAMvlansForbidden with default headers values
func NewGetIPAMvlansForbidden() *GetIPAMvlansForbidden {
	return &GetIPAMvlansForbidden{}
}

/*GetIPAMvlansForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetIPAMvlansForbidden struct {
}

func (o *GetIPAMvlansForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/vlans/][%d] getIpAMvlansForbidden ", 403)
}

func (o *GetIPAMvlansForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMvlansNotFound creates a GetIPAMvlansNotFound with default headers values
func NewGetIPAMvlansNotFound() *GetIPAMvlansNotFound {
	return &GetIPAMvlansNotFound{}
}

/*GetIPAMvlansNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetIPAMvlansNotFound struct {
}

func (o *GetIPAMvlansNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/vlans/][%d] getIpAMvlansNotFound ", 404)
}

func (o *GetIPAMvlansNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMvlansMethodNotAllowed creates a GetIPAMvlansMethodNotAllowed with default headers values
func NewGetIPAMvlansMethodNotAllowed() *GetIPAMvlansMethodNotAllowed {
	return &GetIPAMvlansMethodNotAllowed{}
}

/*GetIPAMvlansMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetIPAMvlansMethodNotAllowed struct {
}

func (o *GetIPAMvlansMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/vlans/][%d] getIpAMvlansMethodNotAllowed ", 405)
}

func (o *GetIPAMvlansMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMvlansGone creates a GetIPAMvlansGone with default headers values
func NewGetIPAMvlansGone() *GetIPAMvlansGone {
	return &GetIPAMvlansGone{}
}

/*GetIPAMvlansGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetIPAMvlansGone struct {
}

func (o *GetIPAMvlansGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/vlans/][%d] getIpAMvlansGone ", 410)
}

func (o *GetIPAMvlansGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMvlansInternalServerError creates a GetIPAMvlansInternalServerError with default headers values
func NewGetIPAMvlansInternalServerError() *GetIPAMvlansInternalServerError {
	return &GetIPAMvlansInternalServerError{}
}

/*GetIPAMvlansInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetIPAMvlansInternalServerError struct {
}

func (o *GetIPAMvlansInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/vlans/][%d] getIpAMvlansInternalServerError ", 500)
}

func (o *GetIPAMvlansInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMvlansServiceUnavailable creates a GetIPAMvlansServiceUnavailable with default headers values
func NewGetIPAMvlansServiceUnavailable() *GetIPAMvlansServiceUnavailable {
	return &GetIPAMvlansServiceUnavailable{}
}

/*GetIPAMvlansServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetIPAMvlansServiceUnavailable struct {
}

func (o *GetIPAMvlansServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/vlans/][%d] getIpAMvlansServiceUnavailable ", 503)
}

func (o *GetIPAMvlansServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetIPAMvlansOKBody get IP a mvlans o k body
swagger:model GetIPAMvlansOKBody
*/
type GetIPAMvlansOKBody struct {

	// vlans
	Vlans []*models.IPAMvlans `json:"vlans"`
}

// Validate validates this get IP a mvlans o k body
func (o *GetIPAMvlansOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVlans(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetIPAMvlansOKBody) validateVlans(formats strfmt.Registry) error {

	if swag.IsZero(o.Vlans) { // not required
		return nil
	}

	for i := 0; i < len(o.Vlans); i++ {
		if swag.IsZero(o.Vlans[i]) { // not required
			continue
		}

		if o.Vlans[i] != nil {
			if err := o.Vlans[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getIpAMvlansOK" + "." + "vlans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetIPAMvlansOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIPAMvlansOKBody) UnmarshalBinary(b []byte) error {
	var res GetIPAMvlansOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
