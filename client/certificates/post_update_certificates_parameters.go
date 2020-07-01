// Code generated by go-swagger; DO NOT EDIT.

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPostUpdateCertificatesParams creates a new PostUpdateCertificatesParams object
// with the default values initialized.
func NewPostUpdateCertificatesParams() *PostUpdateCertificatesParams {
	var ()
	return &PostUpdateCertificatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUpdateCertificatesParamsWithTimeout creates a new PostUpdateCertificatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUpdateCertificatesParamsWithTimeout(timeout time.Duration) *PostUpdateCertificatesParams {
	var ()
	return &PostUpdateCertificatesParams{

		timeout: timeout,
	}
}

// NewPostUpdateCertificatesParamsWithContext creates a new PostUpdateCertificatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUpdateCertificatesParamsWithContext(ctx context.Context) *PostUpdateCertificatesParams {
	var ()
	return &PostUpdateCertificatesParams{

		Context: ctx,
	}
}

// NewPostUpdateCertificatesParamsWithHTTPClient creates a new PostUpdateCertificatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUpdateCertificatesParamsWithHTTPClient(client *http.Client) *PostUpdateCertificatesParams {
	var ()
	return &PostUpdateCertificatesParams{
		HTTPClient: client,
	}
}

/*PostUpdateCertificatesParams contains all the parameters to send to the API endpoint
for the post update certificates operation typically these are written to a http.Request
*/
type PostUpdateCertificatesParams struct {

	/*ContentCommitmentUsage*/
	ContentCommitmentUsage *string
	/*CrlSignUsage
	  True when the subject public key is used for verifying signatures on certificate revocation lists.

	*/
	CrlSignUsage *string
	/*DataEnciphermentUsage
	  True when the subject public key is used for directly enciphering raw user data without the use of an intermediate symmetric cipher.

	*/
	DataEnciphermentUsage *string
	/*DecipherOnlyUsage
	  When the decipherOnly bit is asserted and the keyAgreement bit is also set, the subject public key may be used only for deciphering data while performing key agreement.

	*/
	DecipherOnlyUsage *string
	/*DigitalSignatureUsage
	  True when the subject public key is used for verifying digital signatures

	*/
	DigitalSignatureUsage *string
	/*DNS
	  dns address of the site certificate is issued for. If entered, is the only required field

	*/
	DNS *string
	/*EncipherOnlyUsage
	  When the encipherOnly bit is asserted and the keyAgreement bit is also set, the subject public key may be used only for enciphering data while performing key agreement.

	*/
	EncipherOnlyUsage *string
	/*EndPointID*/
	EndPointID *string
	/*EndPointType*/
	EndPointType *string
	/*ExtendedKeyUsage
	  Indicates the purpose of the public key contained in the certificate. It contains a list of OIDs, each of which indicates an allowed use.

	*/
	ExtendedKeyUsage *string
	/*Groups
	  Use only if Multitenancy is on. List of groups separated by commas. Each group has name followed by colon followed by yes or no for change_permission, eg. CorpUS:yes, CorpNY:no.

	*/
	Groups *string
	/*IssuedBy
	  The entity issuing the certificate

	*/
	IssuedBy *string
	/*IssuedTo
	  Entity certificate is issued to, required if no dns value

	*/
	IssuedTo *string
	/*KeyAgreementUsage
	  True when the subject public key is used for key agreement.

	*/
	KeyAgreementUsage *string
	/*KeyCertSignUsage
	  True when the subject public key is used for verifying signatures on public key certificates.

	*/
	KeyCertSignUsage *string
	/*KeyEnciphermentUsage
	  True when the subject public key is used for enciphering private or secret keys

	*/
	KeyEnciphermentUsage *string
	/*SerialNumber
	  Used to uniquely identify the certificate

	*/
	SerialNumber *string
	/*SignatureAlgorithm
	  The algorithm used to create the signature.

	*/
	SignatureAlgorithm *string
	/*SignatureHash
	  The actual signature to verify that it came from the issuer

	*/
	SignatureHash *string
	/*Subject
	  The person, or entity identified. Required if no dns value

	*/
	Subject *string
	/*ValidFrom
	  The start date of the certificate, required if no dns value. YYYY-MM-DD format

	*/
	ValidFrom *string
	/*ValidTo
	  The expiration date of the certificate, required if no dns value. YYYY-MM-DD format

	*/
	ValidTo *string
	/*Vendor
	  The name of the vendor that provided this certificate

	*/
	Vendor *string
	/*Version
	  The version of the encoded certificate

	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post update certificates params
func (o *PostUpdateCertificatesParams) WithTimeout(timeout time.Duration) *PostUpdateCertificatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post update certificates params
func (o *PostUpdateCertificatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post update certificates params
func (o *PostUpdateCertificatesParams) WithContext(ctx context.Context) *PostUpdateCertificatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post update certificates params
func (o *PostUpdateCertificatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post update certificates params
func (o *PostUpdateCertificatesParams) WithHTTPClient(client *http.Client) *PostUpdateCertificatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post update certificates params
func (o *PostUpdateCertificatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentCommitmentUsage adds the contentCommitmentUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) WithContentCommitmentUsage(contentCommitmentUsage *string) *PostUpdateCertificatesParams {
	o.SetContentCommitmentUsage(contentCommitmentUsage)
	return o
}

// SetContentCommitmentUsage adds the contentCommitmentUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) SetContentCommitmentUsage(contentCommitmentUsage *string) {
	o.ContentCommitmentUsage = contentCommitmentUsage
}

// WithCrlSignUsage adds the crlSignUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) WithCrlSignUsage(crlSignUsage *string) *PostUpdateCertificatesParams {
	o.SetCrlSignUsage(crlSignUsage)
	return o
}

// SetCrlSignUsage adds the crlSignUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) SetCrlSignUsage(crlSignUsage *string) {
	o.CrlSignUsage = crlSignUsage
}

// WithDataEnciphermentUsage adds the dataEnciphermentUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) WithDataEnciphermentUsage(dataEnciphermentUsage *string) *PostUpdateCertificatesParams {
	o.SetDataEnciphermentUsage(dataEnciphermentUsage)
	return o
}

// SetDataEnciphermentUsage adds the dataEnciphermentUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) SetDataEnciphermentUsage(dataEnciphermentUsage *string) {
	o.DataEnciphermentUsage = dataEnciphermentUsage
}

// WithDecipherOnlyUsage adds the decipherOnlyUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) WithDecipherOnlyUsage(decipherOnlyUsage *string) *PostUpdateCertificatesParams {
	o.SetDecipherOnlyUsage(decipherOnlyUsage)
	return o
}

// SetDecipherOnlyUsage adds the decipherOnlyUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) SetDecipherOnlyUsage(decipherOnlyUsage *string) {
	o.DecipherOnlyUsage = decipherOnlyUsage
}

// WithDigitalSignatureUsage adds the digitalSignatureUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) WithDigitalSignatureUsage(digitalSignatureUsage *string) *PostUpdateCertificatesParams {
	o.SetDigitalSignatureUsage(digitalSignatureUsage)
	return o
}

// SetDigitalSignatureUsage adds the digitalSignatureUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) SetDigitalSignatureUsage(digitalSignatureUsage *string) {
	o.DigitalSignatureUsage = digitalSignatureUsage
}

// WithDNS adds the dns to the post update certificates params
func (o *PostUpdateCertificatesParams) WithDNS(dns *string) *PostUpdateCertificatesParams {
	o.SetDNS(dns)
	return o
}

// SetDNS adds the dns to the post update certificates params
func (o *PostUpdateCertificatesParams) SetDNS(dns *string) {
	o.DNS = dns
}

// WithEncipherOnlyUsage adds the encipherOnlyUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) WithEncipherOnlyUsage(encipherOnlyUsage *string) *PostUpdateCertificatesParams {
	o.SetEncipherOnlyUsage(encipherOnlyUsage)
	return o
}

// SetEncipherOnlyUsage adds the encipherOnlyUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) SetEncipherOnlyUsage(encipherOnlyUsage *string) {
	o.EncipherOnlyUsage = encipherOnlyUsage
}

// WithEndPointID adds the endPointID to the post update certificates params
func (o *PostUpdateCertificatesParams) WithEndPointID(endPointID *string) *PostUpdateCertificatesParams {
	o.SetEndPointID(endPointID)
	return o
}

// SetEndPointID adds the endPointId to the post update certificates params
func (o *PostUpdateCertificatesParams) SetEndPointID(endPointID *string) {
	o.EndPointID = endPointID
}

// WithEndPointType adds the endPointType to the post update certificates params
func (o *PostUpdateCertificatesParams) WithEndPointType(endPointType *string) *PostUpdateCertificatesParams {
	o.SetEndPointType(endPointType)
	return o
}

// SetEndPointType adds the endPointType to the post update certificates params
func (o *PostUpdateCertificatesParams) SetEndPointType(endPointType *string) {
	o.EndPointType = endPointType
}

// WithExtendedKeyUsage adds the extendedKeyUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) WithExtendedKeyUsage(extendedKeyUsage *string) *PostUpdateCertificatesParams {
	o.SetExtendedKeyUsage(extendedKeyUsage)
	return o
}

// SetExtendedKeyUsage adds the extendedKeyUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) SetExtendedKeyUsage(extendedKeyUsage *string) {
	o.ExtendedKeyUsage = extendedKeyUsage
}

// WithGroups adds the groups to the post update certificates params
func (o *PostUpdateCertificatesParams) WithGroups(groups *string) *PostUpdateCertificatesParams {
	o.SetGroups(groups)
	return o
}

// SetGroups adds the groups to the post update certificates params
func (o *PostUpdateCertificatesParams) SetGroups(groups *string) {
	o.Groups = groups
}

// WithIssuedBy adds the issuedBy to the post update certificates params
func (o *PostUpdateCertificatesParams) WithIssuedBy(issuedBy *string) *PostUpdateCertificatesParams {
	o.SetIssuedBy(issuedBy)
	return o
}

// SetIssuedBy adds the issuedBy to the post update certificates params
func (o *PostUpdateCertificatesParams) SetIssuedBy(issuedBy *string) {
	o.IssuedBy = issuedBy
}

// WithIssuedTo adds the issuedTo to the post update certificates params
func (o *PostUpdateCertificatesParams) WithIssuedTo(issuedTo *string) *PostUpdateCertificatesParams {
	o.SetIssuedTo(issuedTo)
	return o
}

// SetIssuedTo adds the issuedTo to the post update certificates params
func (o *PostUpdateCertificatesParams) SetIssuedTo(issuedTo *string) {
	o.IssuedTo = issuedTo
}

// WithKeyAgreementUsage adds the keyAgreementUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) WithKeyAgreementUsage(keyAgreementUsage *string) *PostUpdateCertificatesParams {
	o.SetKeyAgreementUsage(keyAgreementUsage)
	return o
}

// SetKeyAgreementUsage adds the keyAgreementUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) SetKeyAgreementUsage(keyAgreementUsage *string) {
	o.KeyAgreementUsage = keyAgreementUsage
}

// WithKeyCertSignUsage adds the keyCertSignUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) WithKeyCertSignUsage(keyCertSignUsage *string) *PostUpdateCertificatesParams {
	o.SetKeyCertSignUsage(keyCertSignUsage)
	return o
}

// SetKeyCertSignUsage adds the keyCertSignUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) SetKeyCertSignUsage(keyCertSignUsage *string) {
	o.KeyCertSignUsage = keyCertSignUsage
}

// WithKeyEnciphermentUsage adds the keyEnciphermentUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) WithKeyEnciphermentUsage(keyEnciphermentUsage *string) *PostUpdateCertificatesParams {
	o.SetKeyEnciphermentUsage(keyEnciphermentUsage)
	return o
}

// SetKeyEnciphermentUsage adds the keyEnciphermentUsage to the post update certificates params
func (o *PostUpdateCertificatesParams) SetKeyEnciphermentUsage(keyEnciphermentUsage *string) {
	o.KeyEnciphermentUsage = keyEnciphermentUsage
}

// WithSerialNumber adds the serialNumber to the post update certificates params
func (o *PostUpdateCertificatesParams) WithSerialNumber(serialNumber *string) *PostUpdateCertificatesParams {
	o.SetSerialNumber(serialNumber)
	return o
}

// SetSerialNumber adds the serialNumber to the post update certificates params
func (o *PostUpdateCertificatesParams) SetSerialNumber(serialNumber *string) {
	o.SerialNumber = serialNumber
}

// WithSignatureAlgorithm adds the signatureAlgorithm to the post update certificates params
func (o *PostUpdateCertificatesParams) WithSignatureAlgorithm(signatureAlgorithm *string) *PostUpdateCertificatesParams {
	o.SetSignatureAlgorithm(signatureAlgorithm)
	return o
}

// SetSignatureAlgorithm adds the signatureAlgorithm to the post update certificates params
func (o *PostUpdateCertificatesParams) SetSignatureAlgorithm(signatureAlgorithm *string) {
	o.SignatureAlgorithm = signatureAlgorithm
}

// WithSignatureHash adds the signatureHash to the post update certificates params
func (o *PostUpdateCertificatesParams) WithSignatureHash(signatureHash *string) *PostUpdateCertificatesParams {
	o.SetSignatureHash(signatureHash)
	return o
}

// SetSignatureHash adds the signatureHash to the post update certificates params
func (o *PostUpdateCertificatesParams) SetSignatureHash(signatureHash *string) {
	o.SignatureHash = signatureHash
}

// WithSubject adds the subject to the post update certificates params
func (o *PostUpdateCertificatesParams) WithSubject(subject *string) *PostUpdateCertificatesParams {
	o.SetSubject(subject)
	return o
}

// SetSubject adds the subject to the post update certificates params
func (o *PostUpdateCertificatesParams) SetSubject(subject *string) {
	o.Subject = subject
}

// WithValidFrom adds the validFrom to the post update certificates params
func (o *PostUpdateCertificatesParams) WithValidFrom(validFrom *string) *PostUpdateCertificatesParams {
	o.SetValidFrom(validFrom)
	return o
}

// SetValidFrom adds the validFrom to the post update certificates params
func (o *PostUpdateCertificatesParams) SetValidFrom(validFrom *string) {
	o.ValidFrom = validFrom
}

// WithValidTo adds the validTo to the post update certificates params
func (o *PostUpdateCertificatesParams) WithValidTo(validTo *string) *PostUpdateCertificatesParams {
	o.SetValidTo(validTo)
	return o
}

// SetValidTo adds the validTo to the post update certificates params
func (o *PostUpdateCertificatesParams) SetValidTo(validTo *string) {
	o.ValidTo = validTo
}

// WithVendor adds the vendor to the post update certificates params
func (o *PostUpdateCertificatesParams) WithVendor(vendor *string) *PostUpdateCertificatesParams {
	o.SetVendor(vendor)
	return o
}

// SetVendor adds the vendor to the post update certificates params
func (o *PostUpdateCertificatesParams) SetVendor(vendor *string) {
	o.Vendor = vendor
}

// WithVersion adds the version to the post update certificates params
func (o *PostUpdateCertificatesParams) WithVersion(version *string) *PostUpdateCertificatesParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the post update certificates params
func (o *PostUpdateCertificatesParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PostUpdateCertificatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentCommitmentUsage != nil {

		// form param content_commitment_usage
		var frContentCommitmentUsage string
		if o.ContentCommitmentUsage != nil {
			frContentCommitmentUsage = *o.ContentCommitmentUsage
		}
		fContentCommitmentUsage := frContentCommitmentUsage
		if fContentCommitmentUsage != "" {
			if err := r.SetFormParam("content_commitment_usage", fContentCommitmentUsage); err != nil {
				return err
			}
		}

	}

	if o.CrlSignUsage != nil {

		// form param crl_sign_usage
		var frCrlSignUsage string
		if o.CrlSignUsage != nil {
			frCrlSignUsage = *o.CrlSignUsage
		}
		fCrlSignUsage := frCrlSignUsage
		if fCrlSignUsage != "" {
			if err := r.SetFormParam("crl_sign_usage", fCrlSignUsage); err != nil {
				return err
			}
		}

	}

	if o.DataEnciphermentUsage != nil {

		// form param data_encipherment_usage
		var frDataEnciphermentUsage string
		if o.DataEnciphermentUsage != nil {
			frDataEnciphermentUsage = *o.DataEnciphermentUsage
		}
		fDataEnciphermentUsage := frDataEnciphermentUsage
		if fDataEnciphermentUsage != "" {
			if err := r.SetFormParam("data_encipherment_usage", fDataEnciphermentUsage); err != nil {
				return err
			}
		}

	}

	if o.DecipherOnlyUsage != nil {

		// form param decipher_only_usage
		var frDecipherOnlyUsage string
		if o.DecipherOnlyUsage != nil {
			frDecipherOnlyUsage = *o.DecipherOnlyUsage
		}
		fDecipherOnlyUsage := frDecipherOnlyUsage
		if fDecipherOnlyUsage != "" {
			if err := r.SetFormParam("decipher_only_usage", fDecipherOnlyUsage); err != nil {
				return err
			}
		}

	}

	if o.DigitalSignatureUsage != nil {

		// form param digital_signature_usage
		var frDigitalSignatureUsage string
		if o.DigitalSignatureUsage != nil {
			frDigitalSignatureUsage = *o.DigitalSignatureUsage
		}
		fDigitalSignatureUsage := frDigitalSignatureUsage
		if fDigitalSignatureUsage != "" {
			if err := r.SetFormParam("digital_signature_usage", fDigitalSignatureUsage); err != nil {
				return err
			}
		}

	}

	if o.DNS != nil {

		// form param dns
		var frDNS string
		if o.DNS != nil {
			frDNS = *o.DNS
		}
		fDNS := frDNS
		if fDNS != "" {
			if err := r.SetFormParam("dns", fDNS); err != nil {
				return err
			}
		}

	}

	if o.EncipherOnlyUsage != nil {

		// form param encipher_only_usage
		var frEncipherOnlyUsage string
		if o.EncipherOnlyUsage != nil {
			frEncipherOnlyUsage = *o.EncipherOnlyUsage
		}
		fEncipherOnlyUsage := frEncipherOnlyUsage
		if fEncipherOnlyUsage != "" {
			if err := r.SetFormParam("encipher_only_usage", fEncipherOnlyUsage); err != nil {
				return err
			}
		}

	}

	if o.EndPointID != nil {

		// form param end_point_id
		var frEndPointID string
		if o.EndPointID != nil {
			frEndPointID = *o.EndPointID
		}
		fEndPointID := frEndPointID
		if fEndPointID != "" {
			if err := r.SetFormParam("end_point_id", fEndPointID); err != nil {
				return err
			}
		}

	}

	if o.EndPointType != nil {

		// form param end_point_type
		var frEndPointType string
		if o.EndPointType != nil {
			frEndPointType = *o.EndPointType
		}
		fEndPointType := frEndPointType
		if fEndPointType != "" {
			if err := r.SetFormParam("end_point_type", fEndPointType); err != nil {
				return err
			}
		}

	}

	if o.ExtendedKeyUsage != nil {

		// form param extended_key_usage
		var frExtendedKeyUsage string
		if o.ExtendedKeyUsage != nil {
			frExtendedKeyUsage = *o.ExtendedKeyUsage
		}
		fExtendedKeyUsage := frExtendedKeyUsage
		if fExtendedKeyUsage != "" {
			if err := r.SetFormParam("extended_key_usage", fExtendedKeyUsage); err != nil {
				return err
			}
		}

	}

	if o.Groups != nil {

		// form param groups
		var frGroups string
		if o.Groups != nil {
			frGroups = *o.Groups
		}
		fGroups := frGroups
		if fGroups != "" {
			if err := r.SetFormParam("groups", fGroups); err != nil {
				return err
			}
		}

	}

	if o.IssuedBy != nil {

		// form param issued_by
		var frIssuedBy string
		if o.IssuedBy != nil {
			frIssuedBy = *o.IssuedBy
		}
		fIssuedBy := frIssuedBy
		if fIssuedBy != "" {
			if err := r.SetFormParam("issued_by", fIssuedBy); err != nil {
				return err
			}
		}

	}

	if o.IssuedTo != nil {

		// form param issued_to
		var frIssuedTo string
		if o.IssuedTo != nil {
			frIssuedTo = *o.IssuedTo
		}
		fIssuedTo := frIssuedTo
		if fIssuedTo != "" {
			if err := r.SetFormParam("issued_to", fIssuedTo); err != nil {
				return err
			}
		}

	}

	if o.KeyAgreementUsage != nil {

		// form param key_agreement_usage
		var frKeyAgreementUsage string
		if o.KeyAgreementUsage != nil {
			frKeyAgreementUsage = *o.KeyAgreementUsage
		}
		fKeyAgreementUsage := frKeyAgreementUsage
		if fKeyAgreementUsage != "" {
			if err := r.SetFormParam("key_agreement_usage", fKeyAgreementUsage); err != nil {
				return err
			}
		}

	}

	if o.KeyCertSignUsage != nil {

		// form param key_cert_sign_usage
		var frKeyCertSignUsage string
		if o.KeyCertSignUsage != nil {
			frKeyCertSignUsage = *o.KeyCertSignUsage
		}
		fKeyCertSignUsage := frKeyCertSignUsage
		if fKeyCertSignUsage != "" {
			if err := r.SetFormParam("key_cert_sign_usage", fKeyCertSignUsage); err != nil {
				return err
			}
		}

	}

	if o.KeyEnciphermentUsage != nil {

		// form param key_encipherment_usage
		var frKeyEnciphermentUsage string
		if o.KeyEnciphermentUsage != nil {
			frKeyEnciphermentUsage = *o.KeyEnciphermentUsage
		}
		fKeyEnciphermentUsage := frKeyEnciphermentUsage
		if fKeyEnciphermentUsage != "" {
			if err := r.SetFormParam("key_encipherment_usage", fKeyEnciphermentUsage); err != nil {
				return err
			}
		}

	}

	if o.SerialNumber != nil {

		// form param serial_number
		var frSerialNumber string
		if o.SerialNumber != nil {
			frSerialNumber = *o.SerialNumber
		}
		fSerialNumber := frSerialNumber
		if fSerialNumber != "" {
			if err := r.SetFormParam("serial_number", fSerialNumber); err != nil {
				return err
			}
		}

	}

	if o.SignatureAlgorithm != nil {

		// form param signature_algorithm
		var frSignatureAlgorithm string
		if o.SignatureAlgorithm != nil {
			frSignatureAlgorithm = *o.SignatureAlgorithm
		}
		fSignatureAlgorithm := frSignatureAlgorithm
		if fSignatureAlgorithm != "" {
			if err := r.SetFormParam("signature_algorithm", fSignatureAlgorithm); err != nil {
				return err
			}
		}

	}

	if o.SignatureHash != nil {

		// form param signature_hash
		var frSignatureHash string
		if o.SignatureHash != nil {
			frSignatureHash = *o.SignatureHash
		}
		fSignatureHash := frSignatureHash
		if fSignatureHash != "" {
			if err := r.SetFormParam("signature_hash", fSignatureHash); err != nil {
				return err
			}
		}

	}

	if o.Subject != nil {

		// form param subject
		var frSubject string
		if o.Subject != nil {
			frSubject = *o.Subject
		}
		fSubject := frSubject
		if fSubject != "" {
			if err := r.SetFormParam("subject", fSubject); err != nil {
				return err
			}
		}

	}

	if o.ValidFrom != nil {

		// form param valid_from
		var frValidFrom string
		if o.ValidFrom != nil {
			frValidFrom = *o.ValidFrom
		}
		fValidFrom := frValidFrom
		if fValidFrom != "" {
			if err := r.SetFormParam("valid_from", fValidFrom); err != nil {
				return err
			}
		}

	}

	if o.ValidTo != nil {

		// form param valid_to
		var frValidTo string
		if o.ValidTo != nil {
			frValidTo = *o.ValidTo
		}
		fValidTo := frValidTo
		if fValidTo != "" {
			if err := r.SetFormParam("valid_to", fValidTo); err != nil {
				return err
			}
		}

	}

	if o.Vendor != nil {

		// form param vendor
		var frVendor string
		if o.Vendor != nil {
			frVendor = *o.Vendor
		}
		fVendor := frVendor
		if fVendor != "" {
			if err := r.SetFormParam("vendor", fVendor); err != nil {
				return err
			}
		}

	}

	if o.Version != nil {

		// form param version
		var frVersion string
		if o.Version != nil {
			frVersion = *o.Version
		}
		fVersion := frVersion
		if fVersion != "" {
			if err := r.SetFormParam("version", fVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
