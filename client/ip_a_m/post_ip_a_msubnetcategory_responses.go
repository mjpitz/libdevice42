// Code generated by go-swagger; DO NOT EDIT.

package ip_a_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostIPAMsubnetcategoryReader is a Reader for the PostIPAMsubnetcategory structure.
type PostIPAMsubnetcategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIPAMsubnetcategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostIPAMsubnetcategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostIPAMsubnetcategoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostIPAMsubnetcategoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostIPAMsubnetcategoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostIPAMsubnetcategoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostIPAMsubnetcategoryMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostIPAMsubnetcategoryGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostIPAMsubnetcategoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostIPAMsubnetcategoryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostIPAMsubnetcategoryOK creates a PostIPAMsubnetcategoryOK with default headers values
func NewPostIPAMsubnetcategoryOK() *PostIPAMsubnetcategoryOK {
	return &PostIPAMsubnetcategoryOK{}
}

/*PostIPAMsubnetcategoryOK handles this case with default header values.

The above command returns JSON structured like this:
*/
type PostIPAMsubnetcategoryOK struct {
}

func (o *PostIPAMsubnetcategoryOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/subnet_category/][%d] postIpAMsubnetcategoryOK ", 200)
}

func (o *PostIPAMsubnetcategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMsubnetcategoryBadRequest creates a PostIPAMsubnetcategoryBadRequest with default headers values
func NewPostIPAMsubnetcategoryBadRequest() *PostIPAMsubnetcategoryBadRequest {
	return &PostIPAMsubnetcategoryBadRequest{}
}

/*PostIPAMsubnetcategoryBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostIPAMsubnetcategoryBadRequest struct {
}

func (o *PostIPAMsubnetcategoryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/subnet_category/][%d] postIpAMsubnetcategoryBadRequest ", 400)
}

func (o *PostIPAMsubnetcategoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMsubnetcategoryUnauthorized creates a PostIPAMsubnetcategoryUnauthorized with default headers values
func NewPostIPAMsubnetcategoryUnauthorized() *PostIPAMsubnetcategoryUnauthorized {
	return &PostIPAMsubnetcategoryUnauthorized{}
}

/*PostIPAMsubnetcategoryUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostIPAMsubnetcategoryUnauthorized struct {
}

func (o *PostIPAMsubnetcategoryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/subnet_category/][%d] postIpAMsubnetcategoryUnauthorized ", 401)
}

func (o *PostIPAMsubnetcategoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMsubnetcategoryForbidden creates a PostIPAMsubnetcategoryForbidden with default headers values
func NewPostIPAMsubnetcategoryForbidden() *PostIPAMsubnetcategoryForbidden {
	return &PostIPAMsubnetcategoryForbidden{}
}

/*PostIPAMsubnetcategoryForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostIPAMsubnetcategoryForbidden struct {
}

func (o *PostIPAMsubnetcategoryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/subnet_category/][%d] postIpAMsubnetcategoryForbidden ", 403)
}

func (o *PostIPAMsubnetcategoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMsubnetcategoryNotFound creates a PostIPAMsubnetcategoryNotFound with default headers values
func NewPostIPAMsubnetcategoryNotFound() *PostIPAMsubnetcategoryNotFound {
	return &PostIPAMsubnetcategoryNotFound{}
}

/*PostIPAMsubnetcategoryNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostIPAMsubnetcategoryNotFound struct {
}

func (o *PostIPAMsubnetcategoryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/subnet_category/][%d] postIpAMsubnetcategoryNotFound ", 404)
}

func (o *PostIPAMsubnetcategoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMsubnetcategoryMethodNotAllowed creates a PostIPAMsubnetcategoryMethodNotAllowed with default headers values
func NewPostIPAMsubnetcategoryMethodNotAllowed() *PostIPAMsubnetcategoryMethodNotAllowed {
	return &PostIPAMsubnetcategoryMethodNotAllowed{}
}

/*PostIPAMsubnetcategoryMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostIPAMsubnetcategoryMethodNotAllowed struct {
}

func (o *PostIPAMsubnetcategoryMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/subnet_category/][%d] postIpAMsubnetcategoryMethodNotAllowed ", 405)
}

func (o *PostIPAMsubnetcategoryMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMsubnetcategoryGone creates a PostIPAMsubnetcategoryGone with default headers values
func NewPostIPAMsubnetcategoryGone() *PostIPAMsubnetcategoryGone {
	return &PostIPAMsubnetcategoryGone{}
}

/*PostIPAMsubnetcategoryGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostIPAMsubnetcategoryGone struct {
}

func (o *PostIPAMsubnetcategoryGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/subnet_category/][%d] postIpAMsubnetcategoryGone ", 410)
}

func (o *PostIPAMsubnetcategoryGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMsubnetcategoryInternalServerError creates a PostIPAMsubnetcategoryInternalServerError with default headers values
func NewPostIPAMsubnetcategoryInternalServerError() *PostIPAMsubnetcategoryInternalServerError {
	return &PostIPAMsubnetcategoryInternalServerError{}
}

/*PostIPAMsubnetcategoryInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type PostIPAMsubnetcategoryInternalServerError struct {
}

func (o *PostIPAMsubnetcategoryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/subnet_category/][%d] postIpAMsubnetcategoryInternalServerError ", 500)
}

func (o *PostIPAMsubnetcategoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMsubnetcategoryServiceUnavailable creates a PostIPAMsubnetcategoryServiceUnavailable with default headers values
func NewPostIPAMsubnetcategoryServiceUnavailable() *PostIPAMsubnetcategoryServiceUnavailable {
	return &PostIPAMsubnetcategoryServiceUnavailable{}
}

/*PostIPAMsubnetcategoryServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostIPAMsubnetcategoryServiceUnavailable struct {
}

func (o *PostIPAMsubnetcategoryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/subnet_category/][%d] postIpAMsubnetcategoryServiceUnavailable ", 503)
}

func (o *PostIPAMsubnetcategoryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
