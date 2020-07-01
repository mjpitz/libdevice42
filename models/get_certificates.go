// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetCertificates get certificates
//
// swagger:model Get_certificates
type GetCertificates struct {

	// content commitment usage
	ContentCommitmentUsage interface{} `json:"content_commitment_usage,omitempty"`

	// crl sign usage
	CrlSignUsage interface{} `json:"crl_sign_usage,omitempty"`

	// custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// data encipherment usage
	DataEnciphermentUsage interface{} `json:"data_encipherment_usage,omitempty"`

	// days to expiry
	DaysToExpiry interface{} `json:"days_to_expiry,omitempty"`

	// decipher only usage
	DecipherOnlyUsage interface{} `json:"decipher_only_usage,omitempty"`

	// digital signature usage
	DigitalSignatureUsage interface{} `json:"digital_signature_usage,omitempty"`

	// encipher only usage
	EncipherOnlyUsage interface{} `json:"encipher_only_usage,omitempty"`

	// extended key usage
	ExtendedKeyUsage interface{} `json:"extended_key_usage,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`

	// issued by
	IssuedBy interface{} `json:"issued_by,omitempty"`

	// issued to
	IssuedTo interface{} `json:"issued_to,omitempty"`

	// key agreement usage
	KeyAgreementUsage interface{} `json:"key_agreement_usage,omitempty"`

	// key cert sign usage
	KeyCertSignUsage interface{} `json:"key_cert_sign_usage,omitempty"`

	// key encipherment usage
	KeyEnciphermentUsage interface{} `json:"key_encipherment_usage,omitempty"`

	// parent cert
	ParentCert interface{} `json:"parent_cert,omitempty"`

	// serial number
	SerialNumber interface{} `json:"serial_number,omitempty"`

	// signature algorithm
	SignatureAlgorithm interface{} `json:"signature_algorithm,omitempty"`

	// signature hash
	SignatureHash interface{} `json:"signature_hash,omitempty"`

	// subject
	Subject interface{} `json:"subject,omitempty"`

	// valid from
	ValidFrom interface{} `json:"valid_from,omitempty"`

	// valid to
	ValidTo interface{} `json:"valid_to,omitempty"`

	// vendor
	Vendor interface{} `json:"vendor,omitempty"`

	// version
	Version interface{} `json:"version,omitempty"`
}

// Validate validates this get certificates
func (m *GetCertificates) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetCertificates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCertificates) UnmarshalBinary(b []byte) error {
	var res GetCertificates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
