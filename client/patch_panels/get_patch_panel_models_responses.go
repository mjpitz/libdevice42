// Code generated by go-swagger; DO NOT EDIT.

package patch_panels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetPatchPanelModelsReader is a Reader for the GetPatchPanelModels structure.
type GetPatchPanelModelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPatchPanelModelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPatchPanelModelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPatchPanelModelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPatchPanelModelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPatchPanelModelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPatchPanelModelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetPatchPanelModelsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetPatchPanelModelsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPatchPanelModelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetPatchPanelModelsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPatchPanelModelsOK creates a GetPatchPanelModelsOK with default headers values
func NewGetPatchPanelModelsOK() *GetPatchPanelModelsOK {
	return &GetPatchPanelModelsOK{}
}

/*GetPatchPanelModelsOK handles this case with default header values.

The above command returns results like this:
*/
type GetPatchPanelModelsOK struct {
	Payload []*GetPatchPanelModelsOKBodyItems0
}

func (o *GetPatchPanelModelsOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/patch_panel_models/][%d] getPatchPanelModelsOK  %+v", 200, o.Payload)
}

func (o *GetPatchPanelModelsOK) GetPayload() []*GetPatchPanelModelsOKBodyItems0 {
	return o.Payload
}

func (o *GetPatchPanelModelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPatchPanelModelsBadRequest creates a GetPatchPanelModelsBadRequest with default headers values
func NewGetPatchPanelModelsBadRequest() *GetPatchPanelModelsBadRequest {
	return &GetPatchPanelModelsBadRequest{}
}

/*GetPatchPanelModelsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetPatchPanelModelsBadRequest struct {
}

func (o *GetPatchPanelModelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/patch_panel_models/][%d] getPatchPanelModelsBadRequest ", 400)
}

func (o *GetPatchPanelModelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPatchPanelModelsUnauthorized creates a GetPatchPanelModelsUnauthorized with default headers values
func NewGetPatchPanelModelsUnauthorized() *GetPatchPanelModelsUnauthorized {
	return &GetPatchPanelModelsUnauthorized{}
}

/*GetPatchPanelModelsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetPatchPanelModelsUnauthorized struct {
}

func (o *GetPatchPanelModelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/patch_panel_models/][%d] getPatchPanelModelsUnauthorized ", 401)
}

func (o *GetPatchPanelModelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPatchPanelModelsForbidden creates a GetPatchPanelModelsForbidden with default headers values
func NewGetPatchPanelModelsForbidden() *GetPatchPanelModelsForbidden {
	return &GetPatchPanelModelsForbidden{}
}

/*GetPatchPanelModelsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetPatchPanelModelsForbidden struct {
}

func (o *GetPatchPanelModelsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/patch_panel_models/][%d] getPatchPanelModelsForbidden ", 403)
}

func (o *GetPatchPanelModelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPatchPanelModelsNotFound creates a GetPatchPanelModelsNotFound with default headers values
func NewGetPatchPanelModelsNotFound() *GetPatchPanelModelsNotFound {
	return &GetPatchPanelModelsNotFound{}
}

/*GetPatchPanelModelsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetPatchPanelModelsNotFound struct {
}

func (o *GetPatchPanelModelsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/patch_panel_models/][%d] getPatchPanelModelsNotFound ", 404)
}

func (o *GetPatchPanelModelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPatchPanelModelsMethodNotAllowed creates a GetPatchPanelModelsMethodNotAllowed with default headers values
func NewGetPatchPanelModelsMethodNotAllowed() *GetPatchPanelModelsMethodNotAllowed {
	return &GetPatchPanelModelsMethodNotAllowed{}
}

/*GetPatchPanelModelsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetPatchPanelModelsMethodNotAllowed struct {
}

func (o *GetPatchPanelModelsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/patch_panel_models/][%d] getPatchPanelModelsMethodNotAllowed ", 405)
}

func (o *GetPatchPanelModelsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPatchPanelModelsGone creates a GetPatchPanelModelsGone with default headers values
func NewGetPatchPanelModelsGone() *GetPatchPanelModelsGone {
	return &GetPatchPanelModelsGone{}
}

/*GetPatchPanelModelsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetPatchPanelModelsGone struct {
}

func (o *GetPatchPanelModelsGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/patch_panel_models/][%d] getPatchPanelModelsGone ", 410)
}

func (o *GetPatchPanelModelsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPatchPanelModelsInternalServerError creates a GetPatchPanelModelsInternalServerError with default headers values
func NewGetPatchPanelModelsInternalServerError() *GetPatchPanelModelsInternalServerError {
	return &GetPatchPanelModelsInternalServerError{}
}

/*GetPatchPanelModelsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetPatchPanelModelsInternalServerError struct {
}

func (o *GetPatchPanelModelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/patch_panel_models/][%d] getPatchPanelModelsInternalServerError ", 500)
}

func (o *GetPatchPanelModelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPatchPanelModelsServiceUnavailable creates a GetPatchPanelModelsServiceUnavailable with default headers values
func NewGetPatchPanelModelsServiceUnavailable() *GetPatchPanelModelsServiceUnavailable {
	return &GetPatchPanelModelsServiceUnavailable{}
}

/*GetPatchPanelModelsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetPatchPanelModelsServiceUnavailable struct {
}

func (o *GetPatchPanelModelsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/patch_panel_models/][%d] getPatchPanelModelsServiceUnavailable ", 503)
}

func (o *GetPatchPanelModelsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetPatchPanelModelsOKBodyItems0 get patch panel models o k body items0
swagger:model GetPatchPanelModelsOKBodyItems0
*/
type GetPatchPanelModelsOKBodyItems0 struct {

	// module position
	ModulePosition interface{} `json:"module_position,omitempty"`

	// name
	Name interface{} `json:"name,omitempty"`

	// numbering direction
	NumberingDirection interface{} `json:"numbering_direction,omitempty"`

	// numbering start location
	NumberingStartLocation interface{} `json:"numbering_start_location,omitempty"`

	// patch panel model id
	PatchPanelModelID interface{} `json:"patch_panel_model_id,omitempty"`

	// type
	Type interface{} `json:"type,omitempty"`

	// type id
	TypeID interface{} `json:"type_id,omitempty"`
}

// Validate validates this get patch panel models o k body items0
func (o *GetPatchPanelModelsOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPatchPanelModelsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPatchPanelModelsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetPatchPanelModelsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}