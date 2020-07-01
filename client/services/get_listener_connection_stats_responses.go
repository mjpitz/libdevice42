// Code generated by go-swagger; DO NOT EDIT.

package services

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

// GetListenerConnectionStatsReader is a Reader for the GetListenerConnectionStats structure.
type GetListenerConnectionStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetListenerConnectionStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetListenerConnectionStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetListenerConnectionStatsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetListenerConnectionStatsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetListenerConnectionStatsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetListenerConnectionStatsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetListenerConnectionStatsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetListenerConnectionStatsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetListenerConnectionStatsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetListenerConnectionStatsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetListenerConnectionStatsOK creates a GetListenerConnectionStatsOK with default headers values
func NewGetListenerConnectionStatsOK() *GetListenerConnectionStatsOK {
	return &GetListenerConnectionStatsOK{}
}

/*GetListenerConnectionStatsOK handles this case with default header values.

The above command returns results like this:
*/
type GetListenerConnectionStatsOK struct {
	Payload *GetListenerConnectionStatsOKBody
}

func (o *GetListenerConnectionStatsOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/listener_connection_stats/][%d] getListenerConnectionStatsOK  %+v", 200, o.Payload)
}

func (o *GetListenerConnectionStatsOK) GetPayload() *GetListenerConnectionStatsOKBody {
	return o.Payload
}

func (o *GetListenerConnectionStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetListenerConnectionStatsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListenerConnectionStatsBadRequest creates a GetListenerConnectionStatsBadRequest with default headers values
func NewGetListenerConnectionStatsBadRequest() *GetListenerConnectionStatsBadRequest {
	return &GetListenerConnectionStatsBadRequest{}
}

/*GetListenerConnectionStatsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetListenerConnectionStatsBadRequest struct {
}

func (o *GetListenerConnectionStatsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/listener_connection_stats/][%d] getListenerConnectionStatsBadRequest ", 400)
}

func (o *GetListenerConnectionStatsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListenerConnectionStatsUnauthorized creates a GetListenerConnectionStatsUnauthorized with default headers values
func NewGetListenerConnectionStatsUnauthorized() *GetListenerConnectionStatsUnauthorized {
	return &GetListenerConnectionStatsUnauthorized{}
}

/*GetListenerConnectionStatsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetListenerConnectionStatsUnauthorized struct {
}

func (o *GetListenerConnectionStatsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/listener_connection_stats/][%d] getListenerConnectionStatsUnauthorized ", 401)
}

func (o *GetListenerConnectionStatsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListenerConnectionStatsForbidden creates a GetListenerConnectionStatsForbidden with default headers values
func NewGetListenerConnectionStatsForbidden() *GetListenerConnectionStatsForbidden {
	return &GetListenerConnectionStatsForbidden{}
}

/*GetListenerConnectionStatsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetListenerConnectionStatsForbidden struct {
}

func (o *GetListenerConnectionStatsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/listener_connection_stats/][%d] getListenerConnectionStatsForbidden ", 403)
}

func (o *GetListenerConnectionStatsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListenerConnectionStatsNotFound creates a GetListenerConnectionStatsNotFound with default headers values
func NewGetListenerConnectionStatsNotFound() *GetListenerConnectionStatsNotFound {
	return &GetListenerConnectionStatsNotFound{}
}

/*GetListenerConnectionStatsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetListenerConnectionStatsNotFound struct {
}

func (o *GetListenerConnectionStatsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/listener_connection_stats/][%d] getListenerConnectionStatsNotFound ", 404)
}

func (o *GetListenerConnectionStatsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListenerConnectionStatsMethodNotAllowed creates a GetListenerConnectionStatsMethodNotAllowed with default headers values
func NewGetListenerConnectionStatsMethodNotAllowed() *GetListenerConnectionStatsMethodNotAllowed {
	return &GetListenerConnectionStatsMethodNotAllowed{}
}

/*GetListenerConnectionStatsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetListenerConnectionStatsMethodNotAllowed struct {
}

func (o *GetListenerConnectionStatsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/listener_connection_stats/][%d] getListenerConnectionStatsMethodNotAllowed ", 405)
}

func (o *GetListenerConnectionStatsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListenerConnectionStatsGone creates a GetListenerConnectionStatsGone with default headers values
func NewGetListenerConnectionStatsGone() *GetListenerConnectionStatsGone {
	return &GetListenerConnectionStatsGone{}
}

/*GetListenerConnectionStatsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetListenerConnectionStatsGone struct {
}

func (o *GetListenerConnectionStatsGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/listener_connection_stats/][%d] getListenerConnectionStatsGone ", 410)
}

func (o *GetListenerConnectionStatsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListenerConnectionStatsInternalServerError creates a GetListenerConnectionStatsInternalServerError with default headers values
func NewGetListenerConnectionStatsInternalServerError() *GetListenerConnectionStatsInternalServerError {
	return &GetListenerConnectionStatsInternalServerError{}
}

/*GetListenerConnectionStatsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetListenerConnectionStatsInternalServerError struct {
}

func (o *GetListenerConnectionStatsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/listener_connection_stats/][%d] getListenerConnectionStatsInternalServerError ", 500)
}

func (o *GetListenerConnectionStatsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListenerConnectionStatsServiceUnavailable creates a GetListenerConnectionStatsServiceUnavailable with default headers values
func NewGetListenerConnectionStatsServiceUnavailable() *GetListenerConnectionStatsServiceUnavailable {
	return &GetListenerConnectionStatsServiceUnavailable{}
}

/*GetListenerConnectionStatsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetListenerConnectionStatsServiceUnavailable struct {
}

func (o *GetListenerConnectionStatsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/listener_connection_stats/][%d] getListenerConnectionStatsServiceUnavailable ", 503)
}

func (o *GetListenerConnectionStatsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetListenerConnectionStatsOKBody get listener connection stats o k body
swagger:model GetListenerConnectionStatsOKBody
*/
type GetListenerConnectionStatsOKBody struct {

	// client ips
	ClientIps interface{} `json:"client_ips,omitempty"`

	// client stats
	ClientStats []*models.ClientStats `json:"client_stats"`

	// description
	Description interface{} `json:"description,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`

	// listener device id
	ListenerDeviceID interface{} `json:"listener_device_id,omitempty"`

	// listener device name
	ListenerDeviceName interface{} `json:"listener_device_name,omitempty"`

	// listener service
	ListenerService interface{} `json:"listener_service,omitempty"`

	// listening ip
	ListeningIP interface{} `json:"listening_ip,omitempty"`

	// port
	Port interface{} `json:"port,omitempty"`
}

// Validate validates this get listener connection stats o k body
func (o *GetListenerConnectionStatsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClientStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetListenerConnectionStatsOKBody) validateClientStats(formats strfmt.Registry) error {

	if swag.IsZero(o.ClientStats) { // not required
		return nil
	}

	for i := 0; i < len(o.ClientStats); i++ {
		if swag.IsZero(o.ClientStats[i]) { // not required
			continue
		}

		if o.ClientStats[i] != nil {
			if err := o.ClientStats[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getListenerConnectionStatsOK" + "." + "client_stats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetListenerConnectionStatsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetListenerConnectionStatsOKBody) UnmarshalBinary(b []byte) error {
	var res GetListenerConnectionStatsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
