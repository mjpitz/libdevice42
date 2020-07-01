// Code generated by go-swagger; DO NOT EDIT.

package parts_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostPartsReader is a Reader for the PostParts structure.
type PostPartsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPartsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostPartsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPartsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostPartsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostPartsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostPartsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostPartsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostPartsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPartsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostPartsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPartsOK creates a PostPartsOK with default headers values
func NewPostPartsOK() *PostPartsOK {
	return &PostPartsOK{}
}

/*PostPartsOK handles this case with default header values.

The above command returns results like this:
*/
type PostPartsOK struct {
	Payload *PostPartsOKBody
}

func (o *PostPartsOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/parts/][%d] postPartsOK  %+v", 200, o.Payload)
}

func (o *PostPartsOK) GetPayload() *PostPartsOKBody {
	return o.Payload
}

func (o *PostPartsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostPartsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPartsBadRequest creates a PostPartsBadRequest with default headers values
func NewPostPartsBadRequest() *PostPartsBadRequest {
	return &PostPartsBadRequest{}
}

/*PostPartsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostPartsBadRequest struct {
}

func (o *PostPartsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/parts/][%d] postPartsBadRequest ", 400)
}

func (o *PostPartsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPartsUnauthorized creates a PostPartsUnauthorized with default headers values
func NewPostPartsUnauthorized() *PostPartsUnauthorized {
	return &PostPartsUnauthorized{}
}

/*PostPartsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostPartsUnauthorized struct {
}

func (o *PostPartsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/parts/][%d] postPartsUnauthorized ", 401)
}

func (o *PostPartsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPartsForbidden creates a PostPartsForbidden with default headers values
func NewPostPartsForbidden() *PostPartsForbidden {
	return &PostPartsForbidden{}
}

/*PostPartsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostPartsForbidden struct {
}

func (o *PostPartsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/parts/][%d] postPartsForbidden ", 403)
}

func (o *PostPartsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPartsNotFound creates a PostPartsNotFound with default headers values
func NewPostPartsNotFound() *PostPartsNotFound {
	return &PostPartsNotFound{}
}

/*PostPartsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostPartsNotFound struct {
}

func (o *PostPartsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/parts/][%d] postPartsNotFound ", 404)
}

func (o *PostPartsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPartsMethodNotAllowed creates a PostPartsMethodNotAllowed with default headers values
func NewPostPartsMethodNotAllowed() *PostPartsMethodNotAllowed {
	return &PostPartsMethodNotAllowed{}
}

/*PostPartsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostPartsMethodNotAllowed struct {
}

func (o *PostPartsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/parts/][%d] postPartsMethodNotAllowed ", 405)
}

func (o *PostPartsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPartsGone creates a PostPartsGone with default headers values
func NewPostPartsGone() *PostPartsGone {
	return &PostPartsGone{}
}

/*PostPartsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostPartsGone struct {
}

func (o *PostPartsGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/parts/][%d] postPartsGone ", 410)
}

func (o *PostPartsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPartsInternalServerError creates a PostPartsInternalServerError with default headers values
func NewPostPartsInternalServerError() *PostPartsInternalServerError {
	return &PostPartsInternalServerError{}
}

/*PostPartsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type PostPartsInternalServerError struct {
}

func (o *PostPartsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/parts/][%d] postPartsInternalServerError ", 500)
}

func (o *PostPartsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPartsServiceUnavailable creates a PostPartsServiceUnavailable with default headers values
func NewPostPartsServiceUnavailable() *PostPartsServiceUnavailable {
	return &PostPartsServiceUnavailable{}
}

/*PostPartsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostPartsServiceUnavailable struct {
}

func (o *PostPartsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/parts/][%d] postPartsServiceUnavailable ", 503)
}

func (o *PostPartsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostPartsOKBody post parts o k body
swagger:model PostPartsOKBody
*/
type PostPartsOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post parts o k body
func (o *PostPartsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostPartsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostPartsOKBody) UnmarshalBinary(b []byte) error {
	var res PostPartsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
