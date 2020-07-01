// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostUpdateCustomerContactsReader is a Reader for the PostUpdateCustomerContacts structure.
type PostUpdateCustomerContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUpdateCustomerContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUpdateCustomerContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUpdateCustomerContactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUpdateCustomerContactsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUpdateCustomerContactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUpdateCustomerContactsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostUpdateCustomerContactsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostUpdateCustomerContactsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUpdateCustomerContactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUpdateCustomerContactsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostUpdateCustomerContactsOK creates a PostUpdateCustomerContactsOK with default headers values
func NewPostUpdateCustomerContactsOK() *PostUpdateCustomerContactsOK {
	return &PostUpdateCustomerContactsOK{}
}

/*PostUpdateCustomerContactsOK handles this case with default header values.

The above command returns results like this:
*/
type PostUpdateCustomerContactsOK struct {
	Payload *PostUpdateCustomerContactsOKBody
}

func (o *PostUpdateCustomerContactsOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/customers/contacts/][%d] postUpdateCustomerContactsOK  %+v", 200, o.Payload)
}

func (o *PostUpdateCustomerContactsOK) GetPayload() *PostUpdateCustomerContactsOKBody {
	return o.Payload
}

func (o *PostUpdateCustomerContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostUpdateCustomerContactsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUpdateCustomerContactsBadRequest creates a PostUpdateCustomerContactsBadRequest with default headers values
func NewPostUpdateCustomerContactsBadRequest() *PostUpdateCustomerContactsBadRequest {
	return &PostUpdateCustomerContactsBadRequest{}
}

/*PostUpdateCustomerContactsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostUpdateCustomerContactsBadRequest struct {
}

func (o *PostUpdateCustomerContactsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/customers/contacts/][%d] postUpdateCustomerContactsBadRequest ", 400)
}

func (o *PostUpdateCustomerContactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdateCustomerContactsUnauthorized creates a PostUpdateCustomerContactsUnauthorized with default headers values
func NewPostUpdateCustomerContactsUnauthorized() *PostUpdateCustomerContactsUnauthorized {
	return &PostUpdateCustomerContactsUnauthorized{}
}

/*PostUpdateCustomerContactsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostUpdateCustomerContactsUnauthorized struct {
}

func (o *PostUpdateCustomerContactsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/customers/contacts/][%d] postUpdateCustomerContactsUnauthorized ", 401)
}

func (o *PostUpdateCustomerContactsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdateCustomerContactsForbidden creates a PostUpdateCustomerContactsForbidden with default headers values
func NewPostUpdateCustomerContactsForbidden() *PostUpdateCustomerContactsForbidden {
	return &PostUpdateCustomerContactsForbidden{}
}

/*PostUpdateCustomerContactsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostUpdateCustomerContactsForbidden struct {
}

func (o *PostUpdateCustomerContactsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/customers/contacts/][%d] postUpdateCustomerContactsForbidden ", 403)
}

func (o *PostUpdateCustomerContactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdateCustomerContactsNotFound creates a PostUpdateCustomerContactsNotFound with default headers values
func NewPostUpdateCustomerContactsNotFound() *PostUpdateCustomerContactsNotFound {
	return &PostUpdateCustomerContactsNotFound{}
}

/*PostUpdateCustomerContactsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostUpdateCustomerContactsNotFound struct {
}

func (o *PostUpdateCustomerContactsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/customers/contacts/][%d] postUpdateCustomerContactsNotFound ", 404)
}

func (o *PostUpdateCustomerContactsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdateCustomerContactsMethodNotAllowed creates a PostUpdateCustomerContactsMethodNotAllowed with default headers values
func NewPostUpdateCustomerContactsMethodNotAllowed() *PostUpdateCustomerContactsMethodNotAllowed {
	return &PostUpdateCustomerContactsMethodNotAllowed{}
}

/*PostUpdateCustomerContactsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostUpdateCustomerContactsMethodNotAllowed struct {
}

func (o *PostUpdateCustomerContactsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/customers/contacts/][%d] postUpdateCustomerContactsMethodNotAllowed ", 405)
}

func (o *PostUpdateCustomerContactsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdateCustomerContactsGone creates a PostUpdateCustomerContactsGone with default headers values
func NewPostUpdateCustomerContactsGone() *PostUpdateCustomerContactsGone {
	return &PostUpdateCustomerContactsGone{}
}

/*PostUpdateCustomerContactsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostUpdateCustomerContactsGone struct {
}

func (o *PostUpdateCustomerContactsGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/customers/contacts/][%d] postUpdateCustomerContactsGone ", 410)
}

func (o *PostUpdateCustomerContactsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdateCustomerContactsInternalServerError creates a PostUpdateCustomerContactsInternalServerError with default headers values
func NewPostUpdateCustomerContactsInternalServerError() *PostUpdateCustomerContactsInternalServerError {
	return &PostUpdateCustomerContactsInternalServerError{}
}

/*PostUpdateCustomerContactsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type PostUpdateCustomerContactsInternalServerError struct {
}

func (o *PostUpdateCustomerContactsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/customers/contacts/][%d] postUpdateCustomerContactsInternalServerError ", 500)
}

func (o *PostUpdateCustomerContactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdateCustomerContactsServiceUnavailable creates a PostUpdateCustomerContactsServiceUnavailable with default headers values
func NewPostUpdateCustomerContactsServiceUnavailable() *PostUpdateCustomerContactsServiceUnavailable {
	return &PostUpdateCustomerContactsServiceUnavailable{}
}

/*PostUpdateCustomerContactsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostUpdateCustomerContactsServiceUnavailable struct {
}

func (o *PostUpdateCustomerContactsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/customers/contacts/][%d] postUpdateCustomerContactsServiceUnavailable ", 503)
}

func (o *PostUpdateCustomerContactsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostUpdateCustomerContactsOKBody post update customer contacts o k body
swagger:model PostUpdateCustomerContactsOKBody
*/
type PostUpdateCustomerContactsOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post update customer contacts o k body
func (o *PostUpdateCustomerContactsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostUpdateCustomerContactsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostUpdateCustomerContactsOKBody) UnmarshalBinary(b []byte) error {
	var res PostUpdateCustomerContactsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
