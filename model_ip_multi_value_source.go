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

// IpMultiValueSource Configuration for the IP source.
type IpMultiValueSource struct {
	// An array of header names used to identify the source IP address.
	HeaderNameList []string `json:"headerNameList"`
	ListValueLocation ListValueLocation `json:"listValueLocation"`
	// Indicator if the upstream IP address should be used as the client address if none of the expected headers is returned.
	FallbackToLastHopIp *bool `json:"fallbackToLastHopIp,omitempty"`
}

// NewIpMultiValueSource instantiates a new IpMultiValueSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpMultiValueSource(headerNameList []string, listValueLocation ListValueLocation) *IpMultiValueSource {
	this := IpMultiValueSource{}
	this.HeaderNameList = headerNameList
	this.ListValueLocation = listValueLocation
	return &this
}

// NewIpMultiValueSourceWithDefaults instantiates a new IpMultiValueSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpMultiValueSourceWithDefaults() *IpMultiValueSource {
	this := IpMultiValueSource{}
	return &this
}

// GetHeaderNameList returns the HeaderNameList field value
func (o *IpMultiValueSource) GetHeaderNameList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.HeaderNameList
}

// GetHeaderNameListOk returns a tuple with the HeaderNameList field value
// and a boolean to check if the value has been set.
func (o *IpMultiValueSource) GetHeaderNameListOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.HeaderNameList, true
}

// SetHeaderNameList sets field value
func (o *IpMultiValueSource) SetHeaderNameList(v []string) {
	o.HeaderNameList = v
}

// GetListValueLocation returns the ListValueLocation field value
func (o *IpMultiValueSource) GetListValueLocation() ListValueLocation {
	if o == nil {
		var ret ListValueLocation
		return ret
	}

	return o.ListValueLocation
}

// GetListValueLocationOk returns a tuple with the ListValueLocation field value
// and a boolean to check if the value has been set.
func (o *IpMultiValueSource) GetListValueLocationOk() (*ListValueLocation, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ListValueLocation, true
}

// SetListValueLocation sets field value
func (o *IpMultiValueSource) SetListValueLocation(v ListValueLocation) {
	o.ListValueLocation = v
}

// GetFallbackToLastHopIp returns the FallbackToLastHopIp field value if set, zero value otherwise.
func (o *IpMultiValueSource) GetFallbackToLastHopIp() bool {
	if o == nil || isNil(o.FallbackToLastHopIp) {
		var ret bool
		return ret
	}
	return *o.FallbackToLastHopIp
}

// GetFallbackToLastHopIpOk returns a tuple with the FallbackToLastHopIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpMultiValueSource) GetFallbackToLastHopIpOk() (*bool, bool) {
	if o == nil || isNil(o.FallbackToLastHopIp) {
    return nil, false
	}
	return o.FallbackToLastHopIp, true
}

// HasFallbackToLastHopIp returns a boolean if a field has been set.
func (o *IpMultiValueSource) HasFallbackToLastHopIp() bool {
	if o != nil && !isNil(o.FallbackToLastHopIp) {
		return true
	}

	return false
}

// SetFallbackToLastHopIp gets a reference to the given bool and assigns it to the FallbackToLastHopIp field.
func (o *IpMultiValueSource) SetFallbackToLastHopIp(v bool) {
	o.FallbackToLastHopIp = &v
}

func (o IpMultiValueSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["headerNameList"] = o.HeaderNameList
	}
	if true {
		toSerialize["listValueLocation"] = o.ListValueLocation
	}
	if !isNil(o.FallbackToLastHopIp) {
		toSerialize["fallbackToLastHopIp"] = o.FallbackToLastHopIp
	}
	return json.Marshal(toSerialize)
}

type NullableIpMultiValueSource struct {
	value *IpMultiValueSource
	isSet bool
}

func (v NullableIpMultiValueSource) Get() *IpMultiValueSource {
	return v.value
}

func (v *NullableIpMultiValueSource) Set(val *IpMultiValueSource) {
	v.value = val
	v.isSet = true
}

func (v NullableIpMultiValueSource) IsSet() bool {
	return v.isSet
}

func (v *NullableIpMultiValueSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpMultiValueSource(val *IpMultiValueSource) *NullableIpMultiValueSource {
	return &NullableIpMultiValueSource{value: val, isSet: true}
}

func (v NullableIpMultiValueSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpMultiValueSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


