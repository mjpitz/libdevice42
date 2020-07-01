// Code generated by go-swagger; DO NOT EDIT.

package object_categories

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

// GetObjectCategoriesReader is a Reader for the GetObjectCategories structure.
type GetObjectCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetObjectCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetObjectCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetObjectCategoriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetObjectCategoriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetObjectCategoriesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetObjectCategoriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetObjectCategoriesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetObjectCategoriesGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetObjectCategoriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetObjectCategoriesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetObjectCategoriesOK creates a GetObjectCategoriesOK with default headers values
func NewGetObjectCategoriesOK() *GetObjectCategoriesOK {
	return &GetObjectCategoriesOK{}
}

/*GetObjectCategoriesOK handles this case with default header values.

The above command returns results like this:
*/
type GetObjectCategoriesOK struct {
	Payload *GetObjectCategoriesOKBody
}

func (o *GetObjectCategoriesOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/object_categories/][%d] getObjectCategoriesOK  %+v", 200, o.Payload)
}

func (o *GetObjectCategoriesOK) GetPayload() *GetObjectCategoriesOKBody {
	return o.Payload
}

func (o *GetObjectCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetObjectCategoriesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetObjectCategoriesBadRequest creates a GetObjectCategoriesBadRequest with default headers values
func NewGetObjectCategoriesBadRequest() *GetObjectCategoriesBadRequest {
	return &GetObjectCategoriesBadRequest{}
}

/*GetObjectCategoriesBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetObjectCategoriesBadRequest struct {
}

func (o *GetObjectCategoriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/object_categories/][%d] getObjectCategoriesBadRequest ", 400)
}

func (o *GetObjectCategoriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetObjectCategoriesUnauthorized creates a GetObjectCategoriesUnauthorized with default headers values
func NewGetObjectCategoriesUnauthorized() *GetObjectCategoriesUnauthorized {
	return &GetObjectCategoriesUnauthorized{}
}

/*GetObjectCategoriesUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetObjectCategoriesUnauthorized struct {
}

func (o *GetObjectCategoriesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/object_categories/][%d] getObjectCategoriesUnauthorized ", 401)
}

func (o *GetObjectCategoriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetObjectCategoriesForbidden creates a GetObjectCategoriesForbidden with default headers values
func NewGetObjectCategoriesForbidden() *GetObjectCategoriesForbidden {
	return &GetObjectCategoriesForbidden{}
}

/*GetObjectCategoriesForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetObjectCategoriesForbidden struct {
}

func (o *GetObjectCategoriesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/object_categories/][%d] getObjectCategoriesForbidden ", 403)
}

func (o *GetObjectCategoriesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetObjectCategoriesNotFound creates a GetObjectCategoriesNotFound with default headers values
func NewGetObjectCategoriesNotFound() *GetObjectCategoriesNotFound {
	return &GetObjectCategoriesNotFound{}
}

/*GetObjectCategoriesNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetObjectCategoriesNotFound struct {
}

func (o *GetObjectCategoriesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/object_categories/][%d] getObjectCategoriesNotFound ", 404)
}

func (o *GetObjectCategoriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetObjectCategoriesMethodNotAllowed creates a GetObjectCategoriesMethodNotAllowed with default headers values
func NewGetObjectCategoriesMethodNotAllowed() *GetObjectCategoriesMethodNotAllowed {
	return &GetObjectCategoriesMethodNotAllowed{}
}

/*GetObjectCategoriesMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetObjectCategoriesMethodNotAllowed struct {
}

func (o *GetObjectCategoriesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/object_categories/][%d] getObjectCategoriesMethodNotAllowed ", 405)
}

func (o *GetObjectCategoriesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetObjectCategoriesGone creates a GetObjectCategoriesGone with default headers values
func NewGetObjectCategoriesGone() *GetObjectCategoriesGone {
	return &GetObjectCategoriesGone{}
}

/*GetObjectCategoriesGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetObjectCategoriesGone struct {
}

func (o *GetObjectCategoriesGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/object_categories/][%d] getObjectCategoriesGone ", 410)
}

func (o *GetObjectCategoriesGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetObjectCategoriesInternalServerError creates a GetObjectCategoriesInternalServerError with default headers values
func NewGetObjectCategoriesInternalServerError() *GetObjectCategoriesInternalServerError {
	return &GetObjectCategoriesInternalServerError{}
}

/*GetObjectCategoriesInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetObjectCategoriesInternalServerError struct {
}

func (o *GetObjectCategoriesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/object_categories/][%d] getObjectCategoriesInternalServerError ", 500)
}

func (o *GetObjectCategoriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetObjectCategoriesServiceUnavailable creates a GetObjectCategoriesServiceUnavailable with default headers values
func NewGetObjectCategoriesServiceUnavailable() *GetObjectCategoriesServiceUnavailable {
	return &GetObjectCategoriesServiceUnavailable{}
}

/*GetObjectCategoriesServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetObjectCategoriesServiceUnavailable struct {
}

func (o *GetObjectCategoriesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/object_categories/][%d] getObjectCategoriesServiceUnavailable ", 503)
}

func (o *GetObjectCategoriesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetObjectCategoriesOKBody get object categories o k body
swagger:model GetObjectCategoriesOKBody
*/
type GetObjectCategoriesOKBody struct {

	// limit
	Limit interface{} `json:"limit,omitempty"`

	// object categories
	ObjectCategories []*models.ObjectCategories `json:"object_categories"`

	// offset
	Offset interface{} `json:"offset,omitempty"`

	// total count
	TotalCount interface{} `json:"total_count,omitempty"`
}

// Validate validates this get object categories o k body
func (o *GetObjectCategoriesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateObjectCategories(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetObjectCategoriesOKBody) validateObjectCategories(formats strfmt.Registry) error {

	if swag.IsZero(o.ObjectCategories) { // not required
		return nil
	}

	for i := 0; i < len(o.ObjectCategories); i++ {
		if swag.IsZero(o.ObjectCategories[i]) { // not required
			continue
		}

		if o.ObjectCategories[i] != nil {
			if err := o.ObjectCategories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getObjectCategoriesOK" + "." + "object_categories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetObjectCategoriesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetObjectCategoriesOKBody) UnmarshalBinary(b []byte) error {
	var res GetObjectCategoriesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}