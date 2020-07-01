// Code generated by go-swagger; DO NOT EDIT.

package passwords

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostCustomFieldsReader is a Reader for the PostCustomFields structure.
type PostCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCustomFieldsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostCustomFieldsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostCustomFieldsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostCustomFieldsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostCustomFieldsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostCustomFieldsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostCustomFieldsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostCustomFieldsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostCustomFieldsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCustomFieldsOK creates a PostCustomFieldsOK with default headers values
func NewPostCustomFieldsOK() *PostCustomFieldsOK {
	return &PostCustomFieldsOK{}
}

/*PostCustomFieldsOK handles this case with default header values.

The above command returns results like this:
*/
type PostCustomFieldsOK struct {
	Payload *PostCustomFieldsOKBody
}

func (o *PostCustomFieldsOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/custom_fields/password/][%d] postCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *PostCustomFieldsOK) GetPayload() *PostCustomFieldsOKBody {
	return o.Payload
}

func (o *PostCustomFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostCustomFieldsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCustomFieldsBadRequest creates a PostCustomFieldsBadRequest with default headers values
func NewPostCustomFieldsBadRequest() *PostCustomFieldsBadRequest {
	return &PostCustomFieldsBadRequest{}
}

/*PostCustomFieldsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostCustomFieldsBadRequest struct {
}

func (o *PostCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/custom_fields/password/][%d] postCustomFieldsBadRequest ", 400)
}

func (o *PostCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomFieldsUnauthorized creates a PostCustomFieldsUnauthorized with default headers values
func NewPostCustomFieldsUnauthorized() *PostCustomFieldsUnauthorized {
	return &PostCustomFieldsUnauthorized{}
}

/*PostCustomFieldsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostCustomFieldsUnauthorized struct {
}

func (o *PostCustomFieldsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/custom_fields/password/][%d] postCustomFieldsUnauthorized ", 401)
}

func (o *PostCustomFieldsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomFieldsForbidden creates a PostCustomFieldsForbidden with default headers values
func NewPostCustomFieldsForbidden() *PostCustomFieldsForbidden {
	return &PostCustomFieldsForbidden{}
}

/*PostCustomFieldsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostCustomFieldsForbidden struct {
}

func (o *PostCustomFieldsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/custom_fields/password/][%d] postCustomFieldsForbidden ", 403)
}

func (o *PostCustomFieldsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomFieldsNotFound creates a PostCustomFieldsNotFound with default headers values
func NewPostCustomFieldsNotFound() *PostCustomFieldsNotFound {
	return &PostCustomFieldsNotFound{}
}

/*PostCustomFieldsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostCustomFieldsNotFound struct {
}

func (o *PostCustomFieldsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/custom_fields/password/][%d] postCustomFieldsNotFound ", 404)
}

func (o *PostCustomFieldsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomFieldsMethodNotAllowed creates a PostCustomFieldsMethodNotAllowed with default headers values
func NewPostCustomFieldsMethodNotAllowed() *PostCustomFieldsMethodNotAllowed {
	return &PostCustomFieldsMethodNotAllowed{}
}

/*PostCustomFieldsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostCustomFieldsMethodNotAllowed struct {
}

func (o *PostCustomFieldsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/custom_fields/password/][%d] postCustomFieldsMethodNotAllowed ", 405)
}

func (o *PostCustomFieldsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomFieldsGone creates a PostCustomFieldsGone with default headers values
func NewPostCustomFieldsGone() *PostCustomFieldsGone {
	return &PostCustomFieldsGone{}
}

/*PostCustomFieldsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostCustomFieldsGone struct {
}

func (o *PostCustomFieldsGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/custom_fields/password/][%d] postCustomFieldsGone ", 410)
}

func (o *PostCustomFieldsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomFieldsInternalServerError creates a PostCustomFieldsInternalServerError with default headers values
func NewPostCustomFieldsInternalServerError() *PostCustomFieldsInternalServerError {
	return &PostCustomFieldsInternalServerError{}
}

/*PostCustomFieldsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type PostCustomFieldsInternalServerError struct {
}

func (o *PostCustomFieldsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/custom_fields/password/][%d] postCustomFieldsInternalServerError ", 500)
}

func (o *PostCustomFieldsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomFieldsServiceUnavailable creates a PostCustomFieldsServiceUnavailable with default headers values
func NewPostCustomFieldsServiceUnavailable() *PostCustomFieldsServiceUnavailable {
	return &PostCustomFieldsServiceUnavailable{}
}

/*PostCustomFieldsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostCustomFieldsServiceUnavailable struct {
}

func (o *PostCustomFieldsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/custom_fields/password/][%d] postCustomFieldsServiceUnavailable ", 503)
}

func (o *PostCustomFieldsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostCustomFieldsOKBody post custom fields o k body
swagger:model PostCustomFieldsOKBody
*/
type PostCustomFieldsOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post custom fields o k body
func (o *PostCustomFieldsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCustomFieldsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCustomFieldsOKBody) UnmarshalBinary(b []byte) error {
	var res PostCustomFieldsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}