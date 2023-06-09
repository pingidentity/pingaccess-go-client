/*
Administrative API Documentation

The PingAccess Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingAccess as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API.

API version: 7.3.0.2-SNAPSHOT
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"time"
)

// checks if the KeyPair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyPair{}

// KeyPair A key pair.
type KeyPair struct {
	// The Id for the key pair.
	Id *int64 `json:"id,omitempty"`
	// (sortable) The Serial Number for the key pair.
	SerialNumber string `json:"serialNumber"`
	// (sortable) The Alias for the key pair.
	Alias string `json:"alias"`
	// (sortable) The Subject DN for the key pair.
	SubjectDn string `json:"subjectDn"`
	// (sortable) The common name (CN) identifying the certificate.
	SubjectCn *string `json:"subjectCn,omitempty"`
	// (sortable) The Issuer DN for the key pair.
	IssuerDn string `json:"issuerDn"`
	// (sortable) The date at which the key pair is valid from as the number of milliseconds since January 1, 1970, 00:00:00 GMT.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// (sortable) The date at which the key pair expires as the number of milliseconds since January 1, 1970, 00:00:00 GMT.
	Expires *time.Time `json:"expires,omitempty"`
	// (sortable) The Signature Algorithm used by the key pair.
	SignatureAlgorithm string     `json:"signatureAlgorithm"`
	Status             CertStatus `json:"status"`
	// The SHA1 checksum of the key pair.
	Sha1sum string `json:"sha1sum"`
	// The MD5 checksum of the key pair. The value will be set to \"\" when in FIPS mode.
	Md5sum string `json:"md5sum"`
	// The SHA256 checksum of the key pair.
	Sha256sum string `json:"sha256sum"`
	// A collection of subject alternative names for the certificate.
	SubjectAlternativeNames []GeneralName `json:"subjectAlternativeNames,omitempty"`
	// (sortable) True if a CSR is generated for this key pair.
	CsrPending bool `json:"csrPending"`
	// The HSM Provider ID.  The default value is 0 indicating an HSM is not used for this KeyPair.
	HsmProviderId *int64 `json:"hsmProviderId,omitempty"`
	// The complete list of certificates in the key pair certificate chain.
	ChainCertificates []ChainCertificate `json:"chainCertificates,omitempty"`
}

// NewKeyPair instantiates a new KeyPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyPair(serialNumber string, alias string, subjectDn string, issuerDn string, signatureAlgorithm string, status CertStatus, sha1sum string, md5sum string, sha256sum string, csrPending bool) *KeyPair {
	this := KeyPair{}
	this.SerialNumber = serialNumber
	this.Alias = alias
	this.SubjectDn = subjectDn
	this.IssuerDn = issuerDn
	this.SignatureAlgorithm = signatureAlgorithm
	this.Status = status
	this.Sha1sum = sha1sum
	this.Md5sum = md5sum
	this.Sha256sum = sha256sum
	this.CsrPending = csrPending
	return &this
}

// NewKeyPairWithDefaults instantiates a new KeyPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyPairWithDefaults() *KeyPair {
	this := KeyPair{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KeyPair) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPair) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KeyPair) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *KeyPair) SetId(v int64) {
	o.Id = &v
}

// GetSerialNumber returns the SerialNumber field value
func (o *KeyPair) GetSerialNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value
// and a boolean to check if the value has been set.
func (o *KeyPair) GetSerialNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SerialNumber, true
}

// SetSerialNumber sets field value
func (o *KeyPair) SetSerialNumber(v string) {
	o.SerialNumber = v
}

// GetAlias returns the Alias field value
func (o *KeyPair) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *KeyPair) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *KeyPair) SetAlias(v string) {
	o.Alias = v
}

// GetSubjectDn returns the SubjectDn field value
func (o *KeyPair) GetSubjectDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectDn
}

// GetSubjectDnOk returns a tuple with the SubjectDn field value
// and a boolean to check if the value has been set.
func (o *KeyPair) GetSubjectDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectDn, true
}

// SetSubjectDn sets field value
func (o *KeyPair) SetSubjectDn(v string) {
	o.SubjectDn = v
}

// GetSubjectCn returns the SubjectCn field value if set, zero value otherwise.
func (o *KeyPair) GetSubjectCn() string {
	if o == nil || IsNil(o.SubjectCn) {
		var ret string
		return ret
	}
	return *o.SubjectCn
}

// GetSubjectCnOk returns a tuple with the SubjectCn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPair) GetSubjectCnOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectCn) {
		return nil, false
	}
	return o.SubjectCn, true
}

// HasSubjectCn returns a boolean if a field has been set.
func (o *KeyPair) HasSubjectCn() bool {
	if o != nil && !IsNil(o.SubjectCn) {
		return true
	}

	return false
}

// SetSubjectCn gets a reference to the given string and assigns it to the SubjectCn field.
func (o *KeyPair) SetSubjectCn(v string) {
	o.SubjectCn = &v
}

// GetIssuerDn returns the IssuerDn field value
func (o *KeyPair) GetIssuerDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuerDn
}

// GetIssuerDnOk returns a tuple with the IssuerDn field value
// and a boolean to check if the value has been set.
func (o *KeyPair) GetIssuerDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuerDn, true
}

// SetIssuerDn sets field value
func (o *KeyPair) SetIssuerDn(v string) {
	o.IssuerDn = v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *KeyPair) GetValidFrom() time.Time {
	if o == nil || IsNil(o.ValidFrom) {
		var ret time.Time
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPair) GetValidFromOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidFrom) {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *KeyPair) HasValidFrom() bool {
	if o != nil && !IsNil(o.ValidFrom) {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.
func (o *KeyPair) SetValidFrom(v time.Time) {
	o.ValidFrom = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *KeyPair) GetExpires() time.Time {
	if o == nil || IsNil(o.Expires) {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPair) GetExpiresOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expires) {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *KeyPair) HasExpires() bool {
	if o != nil && !IsNil(o.Expires) {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *KeyPair) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value
func (o *KeyPair) GetSignatureAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value
// and a boolean to check if the value has been set.
func (o *KeyPair) GetSignatureAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignatureAlgorithm, true
}

// SetSignatureAlgorithm sets field value
func (o *KeyPair) SetSignatureAlgorithm(v string) {
	o.SignatureAlgorithm = v
}

// GetStatus returns the Status field value
func (o *KeyPair) GetStatus() CertStatus {
	if o == nil {
		var ret CertStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *KeyPair) GetStatusOk() (*CertStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *KeyPair) SetStatus(v CertStatus) {
	o.Status = v
}

// GetSha1sum returns the Sha1sum field value
func (o *KeyPair) GetSha1sum() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha1sum
}

// GetSha1sumOk returns a tuple with the Sha1sum field value
// and a boolean to check if the value has been set.
func (o *KeyPair) GetSha1sumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha1sum, true
}

// SetSha1sum sets field value
func (o *KeyPair) SetSha1sum(v string) {
	o.Sha1sum = v
}

// GetMd5sum returns the Md5sum field value
func (o *KeyPair) GetMd5sum() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Md5sum
}

// GetMd5sumOk returns a tuple with the Md5sum field value
// and a boolean to check if the value has been set.
func (o *KeyPair) GetMd5sumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Md5sum, true
}

// SetMd5sum sets field value
func (o *KeyPair) SetMd5sum(v string) {
	o.Md5sum = v
}

// GetSha256sum returns the Sha256sum field value
func (o *KeyPair) GetSha256sum() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha256sum
}

// GetSha256sumOk returns a tuple with the Sha256sum field value
// and a boolean to check if the value has been set.
func (o *KeyPair) GetSha256sumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha256sum, true
}

// SetSha256sum sets field value
func (o *KeyPair) SetSha256sum(v string) {
	o.Sha256sum = v
}

// GetSubjectAlternativeNames returns the SubjectAlternativeNames field value if set, zero value otherwise.
func (o *KeyPair) GetSubjectAlternativeNames() []GeneralName {
	if o == nil || IsNil(o.SubjectAlternativeNames) {
		var ret []GeneralName
		return ret
	}
	return o.SubjectAlternativeNames
}

// GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPair) GetSubjectAlternativeNamesOk() ([]GeneralName, bool) {
	if o == nil || IsNil(o.SubjectAlternativeNames) {
		return nil, false
	}
	return o.SubjectAlternativeNames, true
}

// HasSubjectAlternativeNames returns a boolean if a field has been set.
func (o *KeyPair) HasSubjectAlternativeNames() bool {
	if o != nil && !IsNil(o.SubjectAlternativeNames) {
		return true
	}

	return false
}

// SetSubjectAlternativeNames gets a reference to the given []GeneralName and assigns it to the SubjectAlternativeNames field.
func (o *KeyPair) SetSubjectAlternativeNames(v []GeneralName) {
	o.SubjectAlternativeNames = v
}

// GetCsrPending returns the CsrPending field value
func (o *KeyPair) GetCsrPending() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CsrPending
}

// GetCsrPendingOk returns a tuple with the CsrPending field value
// and a boolean to check if the value has been set.
func (o *KeyPair) GetCsrPendingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CsrPending, true
}

// SetCsrPending sets field value
func (o *KeyPair) SetCsrPending(v bool) {
	o.CsrPending = v
}

// GetHsmProviderId returns the HsmProviderId field value if set, zero value otherwise.
func (o *KeyPair) GetHsmProviderId() int64 {
	if o == nil || IsNil(o.HsmProviderId) {
		var ret int64
		return ret
	}
	return *o.HsmProviderId
}

// GetHsmProviderIdOk returns a tuple with the HsmProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPair) GetHsmProviderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.HsmProviderId) {
		return nil, false
	}
	return o.HsmProviderId, true
}

// HasHsmProviderId returns a boolean if a field has been set.
func (o *KeyPair) HasHsmProviderId() bool {
	if o != nil && !IsNil(o.HsmProviderId) {
		return true
	}

	return false
}

// SetHsmProviderId gets a reference to the given int64 and assigns it to the HsmProviderId field.
func (o *KeyPair) SetHsmProviderId(v int64) {
	o.HsmProviderId = &v
}

// GetChainCertificates returns the ChainCertificates field value if set, zero value otherwise.
func (o *KeyPair) GetChainCertificates() []ChainCertificate {
	if o == nil || IsNil(o.ChainCertificates) {
		var ret []ChainCertificate
		return ret
	}
	return o.ChainCertificates
}

// GetChainCertificatesOk returns a tuple with the ChainCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPair) GetChainCertificatesOk() ([]ChainCertificate, bool) {
	if o == nil || IsNil(o.ChainCertificates) {
		return nil, false
	}
	return o.ChainCertificates, true
}

// HasChainCertificates returns a boolean if a field has been set.
func (o *KeyPair) HasChainCertificates() bool {
	if o != nil && !IsNil(o.ChainCertificates) {
		return true
	}

	return false
}

// SetChainCertificates gets a reference to the given []ChainCertificate and assigns it to the ChainCertificates field.
func (o *KeyPair) SetChainCertificates(v []ChainCertificate) {
	o.ChainCertificates = v
}

func (o KeyPair) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyPair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["serialNumber"] = o.SerialNumber
	toSerialize["alias"] = o.Alias
	toSerialize["subjectDn"] = o.SubjectDn
	if !IsNil(o.SubjectCn) {
		toSerialize["subjectCn"] = o.SubjectCn
	}
	toSerialize["issuerDn"] = o.IssuerDn
	if !IsNil(o.ValidFrom) {
		toSerialize["validFrom"] = o.ValidFrom
	}
	if !IsNil(o.Expires) {
		toSerialize["expires"] = o.Expires
	}
	toSerialize["signatureAlgorithm"] = o.SignatureAlgorithm
	toSerialize["status"] = o.Status
	toSerialize["sha1sum"] = o.Sha1sum
	toSerialize["md5sum"] = o.Md5sum
	toSerialize["sha256sum"] = o.Sha256sum
	if !IsNil(o.SubjectAlternativeNames) {
		toSerialize["subjectAlternativeNames"] = o.SubjectAlternativeNames
	}
	toSerialize["csrPending"] = o.CsrPending
	if !IsNil(o.HsmProviderId) {
		toSerialize["hsmProviderId"] = o.HsmProviderId
	}
	if !IsNil(o.ChainCertificates) {
		toSerialize["chainCertificates"] = o.ChainCertificates
	}
	return toSerialize, nil
}

type NullableKeyPair struct {
	value *KeyPair
	isSet bool
}

func (v NullableKeyPair) Get() *KeyPair {
	return v.value
}

func (v *NullableKeyPair) Set(val *KeyPair) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyPair) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyPair(val *KeyPair) *NullableKeyPair {
	return &NullableKeyPair{value: val, isSet: true}
}

func (v NullableKeyPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
