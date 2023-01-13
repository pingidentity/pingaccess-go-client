/*
PingAccess Administrative API

The PingAccess Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingAccess.

API version: 7.1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AcmeCertificateRequest A reference to a Key Pair to be signed by the ACME protocol.
type AcmeCertificateRequest struct {
	// When creating a new AcmeCertificateRequest, this is the ID for the AcmeCertificateRequest. If not specified, an ID will be automatically assigned. When updating an existing AcmeCertificateRequest, this field is ignored and the ID is determined by the path parameter.
	Id *string `json:"id,omitempty"`
	// The ID of the Key Pair for which a signed certificate will be requested.
	KeyPairId int32 `json:"keyPairId"`
	AcmeCertStatus *AcmeCertStatus `json:"acmeCertStatus,omitempty"`
	// The URL at the ACME server for the associated ACME order.
	Url *string `json:"url,omitempty"`
	// The ID of the ACME Server to be used for the ACME protocol. This is read-only.
	AcmeServerId *string `json:"acmeServerId,omitempty"`
	// The ID of the ACME Account to be used for the ACME protocol. This is read-only.
	AcmeAccountId *string `json:"acmeAccountId,omitempty"`
}

// NewAcmeCertificateRequest instantiates a new AcmeCertificateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcmeCertificateRequest(keyPairId int32) *AcmeCertificateRequest {
	this := AcmeCertificateRequest{}
	this.KeyPairId = keyPairId
	return &this
}

// NewAcmeCertificateRequestWithDefaults instantiates a new AcmeCertificateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcmeCertificateRequestWithDefaults() *AcmeCertificateRequest {
	this := AcmeCertificateRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AcmeCertificateRequest) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeCertificateRequest) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AcmeCertificateRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AcmeCertificateRequest) SetId(v string) {
	o.Id = &v
}

// GetKeyPairId returns the KeyPairId field value
func (o *AcmeCertificateRequest) GetKeyPairId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.KeyPairId
}

// GetKeyPairIdOk returns a tuple with the KeyPairId field value
// and a boolean to check if the value has been set.
func (o *AcmeCertificateRequest) GetKeyPairIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.KeyPairId, true
}

// SetKeyPairId sets field value
func (o *AcmeCertificateRequest) SetKeyPairId(v int32) {
	o.KeyPairId = v
}

// GetAcmeCertStatus returns the AcmeCertStatus field value if set, zero value otherwise.
func (o *AcmeCertificateRequest) GetAcmeCertStatus() AcmeCertStatus {
	if o == nil || isNil(o.AcmeCertStatus) {
		var ret AcmeCertStatus
		return ret
	}
	return *o.AcmeCertStatus
}

// GetAcmeCertStatusOk returns a tuple with the AcmeCertStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeCertificateRequest) GetAcmeCertStatusOk() (*AcmeCertStatus, bool) {
	if o == nil || isNil(o.AcmeCertStatus) {
    return nil, false
	}
	return o.AcmeCertStatus, true
}

// HasAcmeCertStatus returns a boolean if a field has been set.
func (o *AcmeCertificateRequest) HasAcmeCertStatus() bool {
	if o != nil && !isNil(o.AcmeCertStatus) {
		return true
	}

	return false
}

// SetAcmeCertStatus gets a reference to the given AcmeCertStatus and assigns it to the AcmeCertStatus field.
func (o *AcmeCertificateRequest) SetAcmeCertStatus(v AcmeCertStatus) {
	o.AcmeCertStatus = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AcmeCertificateRequest) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeCertificateRequest) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AcmeCertificateRequest) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AcmeCertificateRequest) SetUrl(v string) {
	o.Url = &v
}

// GetAcmeServerId returns the AcmeServerId field value if set, zero value otherwise.
func (o *AcmeCertificateRequest) GetAcmeServerId() string {
	if o == nil || isNil(o.AcmeServerId) {
		var ret string
		return ret
	}
	return *o.AcmeServerId
}

// GetAcmeServerIdOk returns a tuple with the AcmeServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeCertificateRequest) GetAcmeServerIdOk() (*string, bool) {
	if o == nil || isNil(o.AcmeServerId) {
    return nil, false
	}
	return o.AcmeServerId, true
}

// HasAcmeServerId returns a boolean if a field has been set.
func (o *AcmeCertificateRequest) HasAcmeServerId() bool {
	if o != nil && !isNil(o.AcmeServerId) {
		return true
	}

	return false
}

// SetAcmeServerId gets a reference to the given string and assigns it to the AcmeServerId field.
func (o *AcmeCertificateRequest) SetAcmeServerId(v string) {
	o.AcmeServerId = &v
}

// GetAcmeAccountId returns the AcmeAccountId field value if set, zero value otherwise.
func (o *AcmeCertificateRequest) GetAcmeAccountId() string {
	if o == nil || isNil(o.AcmeAccountId) {
		var ret string
		return ret
	}
	return *o.AcmeAccountId
}

// GetAcmeAccountIdOk returns a tuple with the AcmeAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeCertificateRequest) GetAcmeAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AcmeAccountId) {
    return nil, false
	}
	return o.AcmeAccountId, true
}

// HasAcmeAccountId returns a boolean if a field has been set.
func (o *AcmeCertificateRequest) HasAcmeAccountId() bool {
	if o != nil && !isNil(o.AcmeAccountId) {
		return true
	}

	return false
}

// SetAcmeAccountId gets a reference to the given string and assigns it to the AcmeAccountId field.
func (o *AcmeCertificateRequest) SetAcmeAccountId(v string) {
	o.AcmeAccountId = &v
}

func (o AcmeCertificateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["keyPairId"] = o.KeyPairId
	}
	if !isNil(o.AcmeCertStatus) {
		toSerialize["acmeCertStatus"] = o.AcmeCertStatus
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.AcmeServerId) {
		toSerialize["acmeServerId"] = o.AcmeServerId
	}
	if !isNil(o.AcmeAccountId) {
		toSerialize["acmeAccountId"] = o.AcmeAccountId
	}
	return json.Marshal(toSerialize)
}

type NullableAcmeCertificateRequest struct {
	value *AcmeCertificateRequest
	isSet bool
}

func (v NullableAcmeCertificateRequest) Get() *AcmeCertificateRequest {
	return v.value
}

func (v *NullableAcmeCertificateRequest) Set(val *AcmeCertificateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAcmeCertificateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAcmeCertificateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcmeCertificateRequest(val *AcmeCertificateRequest) *NullableAcmeCertificateRequest {
	return &NullableAcmeCertificateRequest{value: val, isSet: true}
}

func (v NullableAcmeCertificateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcmeCertificateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


