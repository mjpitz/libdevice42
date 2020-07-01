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

// GetIPAMMacsReader is a Reader for the GetIPAMMacs structure.
type GetIPAMMacsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPAMMacsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIPAMMacsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIPAMMacsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIPAMMacsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIPAMMacsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIPAMMacsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetIPAMMacsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetIPAMMacsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIPAMMacsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIPAMMacsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIPAMMacsOK creates a GetIPAMMacsOK with default headers values
func NewGetIPAMMacsOK() *GetIPAMMacsOK {
	return &GetIPAMMacsOK{}
}

/*GetIPAMMacsOK handles this case with default header values.

The above command returns results like this:
*/
type GetIPAMMacsOK struct {
	Payload *GetIPAMMacsOKBody
}

func (o *GetIPAMMacsOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/macs/][%d] getIpAMMacsOK  %+v", 200, o.Payload)
}

func (o *GetIPAMMacsOK) GetPayload() *GetIPAMMacsOKBody {
	return o.Payload
}

func (o *GetIPAMMacsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetIPAMMacsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPAMMacsBadRequest creates a GetIPAMMacsBadRequest with default headers values
func NewGetIPAMMacsBadRequest() *GetIPAMMacsBadRequest {
	return &GetIPAMMacsBadRequest{}
}

/*GetIPAMMacsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetIPAMMacsBadRequest struct {
}

func (o *GetIPAMMacsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/macs/][%d] getIpAMMacsBadRequest ", 400)
}

func (o *GetIPAMMacsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMMacsUnauthorized creates a GetIPAMMacsUnauthorized with default headers values
func NewGetIPAMMacsUnauthorized() *GetIPAMMacsUnauthorized {
	return &GetIPAMMacsUnauthorized{}
}

/*GetIPAMMacsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetIPAMMacsUnauthorized struct {
}

func (o *GetIPAMMacsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/macs/][%d] getIpAMMacsUnauthorized ", 401)
}

func (o *GetIPAMMacsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMMacsForbidden creates a GetIPAMMacsForbidden with default headers values
func NewGetIPAMMacsForbidden() *GetIPAMMacsForbidden {
	return &GetIPAMMacsForbidden{}
}

/*GetIPAMMacsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetIPAMMacsForbidden struct {
}

func (o *GetIPAMMacsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/macs/][%d] getIpAMMacsForbidden ", 403)
}

func (o *GetIPAMMacsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMMacsNotFound creates a GetIPAMMacsNotFound with default headers values
func NewGetIPAMMacsNotFound() *GetIPAMMacsNotFound {
	return &GetIPAMMacsNotFound{}
}

/*GetIPAMMacsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetIPAMMacsNotFound struct {
}

func (o *GetIPAMMacsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/macs/][%d] getIpAMMacsNotFound ", 404)
}

func (o *GetIPAMMacsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMMacsMethodNotAllowed creates a GetIPAMMacsMethodNotAllowed with default headers values
func NewGetIPAMMacsMethodNotAllowed() *GetIPAMMacsMethodNotAllowed {
	return &GetIPAMMacsMethodNotAllowed{}
}

/*GetIPAMMacsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetIPAMMacsMethodNotAllowed struct {
}

func (o *GetIPAMMacsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/macs/][%d] getIpAMMacsMethodNotAllowed ", 405)
}

func (o *GetIPAMMacsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMMacsGone creates a GetIPAMMacsGone with default headers values
func NewGetIPAMMacsGone() *GetIPAMMacsGone {
	return &GetIPAMMacsGone{}
}

/*GetIPAMMacsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetIPAMMacsGone struct {
}

func (o *GetIPAMMacsGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/macs/][%d] getIpAMMacsGone ", 410)
}

func (o *GetIPAMMacsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMMacsInternalServerError creates a GetIPAMMacsInternalServerError with default headers values
func NewGetIPAMMacsInternalServerError() *GetIPAMMacsInternalServerError {
	return &GetIPAMMacsInternalServerError{}
}

/*GetIPAMMacsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetIPAMMacsInternalServerError struct {
}

func (o *GetIPAMMacsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/macs/][%d] getIpAMMacsInternalServerError ", 500)
}

func (o *GetIPAMMacsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMMacsServiceUnavailable creates a GetIPAMMacsServiceUnavailable with default headers values
func NewGetIPAMMacsServiceUnavailable() *GetIPAMMacsServiceUnavailable {
	return &GetIPAMMacsServiceUnavailable{}
}

/*GetIPAMMacsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetIPAMMacsServiceUnavailable struct {
}

func (o *GetIPAMMacsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/macs/][%d] getIpAMMacsServiceUnavailable ", 503)
}

func (o *GetIPAMMacsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetIPAMMacsOKBody get IP a m macs o k body
swagger:model GetIPAMMacsOKBody
*/
type GetIPAMMacsOKBody struct {

	// limit
	Limit interface{} `json:"limit,omitempty"`

	// macaddresses
	Macaddresses []*models.IPAMmacs `json:"macaddresses"`

	// offset
	Offset interface{} `json:"offset,omitempty"`

	// total count
	TotalCount interface{} `json:"total_count,omitempty"`
}

// Validate validates this get IP a m macs o k body
func (o *GetIPAMMacsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMacaddresses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetIPAMMacsOKBody) validateMacaddresses(formats strfmt.Registry) error {

	if swag.IsZero(o.Macaddresses) { // not required
		return nil
	}

	for i := 0; i < len(o.Macaddresses); i++ {
		if swag.IsZero(o.Macaddresses[i]) { // not required
			continue
		}

		if o.Macaddresses[i] != nil {
			if err := o.Macaddresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getIpAMMacsOK" + "." + "macaddresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetIPAMMacsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIPAMMacsOKBody) UnmarshalBinary(b []byte) error {
	var res GetIPAMMacsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
