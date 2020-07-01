// Code generated by go-swagger; DO NOT EDIT.

package assets

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

// GetAssetsReader is a Reader for the GetAssets structure.
type GetAssetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAssetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAssetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAssetsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAssetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAssetsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetAssetsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetAssetsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAssetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAssetsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAssetsOK creates a GetAssetsOK with default headers values
func NewGetAssetsOK() *GetAssetsOK {
	return &GetAssetsOK{}
}

/*GetAssetsOK handles this case with default header values.

The above command returns results like this:
*/
type GetAssetsOK struct {
	Payload *GetAssetsOKBody
}

func (o *GetAssetsOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/assets/][%d] getAssetsOK  %+v", 200, o.Payload)
}

func (o *GetAssetsOK) GetPayload() *GetAssetsOKBody {
	return o.Payload
}

func (o *GetAssetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAssetsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssetsBadRequest creates a GetAssetsBadRequest with default headers values
func NewGetAssetsBadRequest() *GetAssetsBadRequest {
	return &GetAssetsBadRequest{}
}

/*GetAssetsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetAssetsBadRequest struct {
}

func (o *GetAssetsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/assets/][%d] getAssetsBadRequest ", 400)
}

func (o *GetAssetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAssetsUnauthorized creates a GetAssetsUnauthorized with default headers values
func NewGetAssetsUnauthorized() *GetAssetsUnauthorized {
	return &GetAssetsUnauthorized{}
}

/*GetAssetsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetAssetsUnauthorized struct {
}

func (o *GetAssetsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/assets/][%d] getAssetsUnauthorized ", 401)
}

func (o *GetAssetsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAssetsForbidden creates a GetAssetsForbidden with default headers values
func NewGetAssetsForbidden() *GetAssetsForbidden {
	return &GetAssetsForbidden{}
}

/*GetAssetsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetAssetsForbidden struct {
}

func (o *GetAssetsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/assets/][%d] getAssetsForbidden ", 403)
}

func (o *GetAssetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAssetsNotFound creates a GetAssetsNotFound with default headers values
func NewGetAssetsNotFound() *GetAssetsNotFound {
	return &GetAssetsNotFound{}
}

/*GetAssetsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetAssetsNotFound struct {
}

func (o *GetAssetsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/assets/][%d] getAssetsNotFound ", 404)
}

func (o *GetAssetsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAssetsMethodNotAllowed creates a GetAssetsMethodNotAllowed with default headers values
func NewGetAssetsMethodNotAllowed() *GetAssetsMethodNotAllowed {
	return &GetAssetsMethodNotAllowed{}
}

/*GetAssetsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetAssetsMethodNotAllowed struct {
}

func (o *GetAssetsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/assets/][%d] getAssetsMethodNotAllowed ", 405)
}

func (o *GetAssetsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAssetsGone creates a GetAssetsGone with default headers values
func NewGetAssetsGone() *GetAssetsGone {
	return &GetAssetsGone{}
}

/*GetAssetsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetAssetsGone struct {
}

func (o *GetAssetsGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/assets/][%d] getAssetsGone ", 410)
}

func (o *GetAssetsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAssetsInternalServerError creates a GetAssetsInternalServerError with default headers values
func NewGetAssetsInternalServerError() *GetAssetsInternalServerError {
	return &GetAssetsInternalServerError{}
}

/*GetAssetsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetAssetsInternalServerError struct {
}

func (o *GetAssetsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/assets/][%d] getAssetsInternalServerError ", 500)
}

func (o *GetAssetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAssetsServiceUnavailable creates a GetAssetsServiceUnavailable with default headers values
func NewGetAssetsServiceUnavailable() *GetAssetsServiceUnavailable {
	return &GetAssetsServiceUnavailable{}
}

/*GetAssetsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetAssetsServiceUnavailable struct {
}

func (o *GetAssetsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/assets/][%d] getAssetsServiceUnavailable ", 503)
}

func (o *GetAssetsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetAssetsOKBody get assets o k body
swagger:model GetAssetsOKBody
*/
type GetAssetsOKBody struct {

	// assets
	Assets []*models.Assets `json:"assets"`

	// limit
	Limit interface{} `json:"limit,omitempty"`

	// offset
	Offset interface{} `json:"offset,omitempty"`

	// total count
	TotalCount interface{} `json:"total_count,omitempty"`
}

// Validate validates this get assets o k body
func (o *GetAssetsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAssets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAssetsOKBody) validateAssets(formats strfmt.Registry) error {

	if swag.IsZero(o.Assets) { // not required
		return nil
	}

	for i := 0; i < len(o.Assets); i++ {
		if swag.IsZero(o.Assets[i]) { // not required
			continue
		}

		if o.Assets[i] != nil {
			if err := o.Assets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getAssetsOK" + "." + "assets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAssetsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAssetsOKBody) UnmarshalBinary(b []byte) error {
	var res GetAssetsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
