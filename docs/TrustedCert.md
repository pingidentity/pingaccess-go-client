# TrustedCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | When creating a new TrustedCert, this is the ID for the TrustedCert. If not specified, an ID will be automatically assigned. When updating an existing TrustedCert, this field is ignored and the ID is determined by the path parameter. | [optional] 
**SerialNumber** | **string** | (sortable) The serial number for the certificate. | 
**Alias** | **string** | (sortable) The alias for the certificate. | 
**SubjectDn** | **string** | (sortable) The subject DN for the certificate. | 
**SubjectCn** | Pointer to **string** | (sortable) The subject CN for the certificate. | [optional] 
**IssuerDn** | **string** | (sortable) The issuer DN for the certificate. | 
**ValidFrom** | Pointer to **time.Time** | (sortable) The date at which the certificate is valid from as the number of milliseconds since January 1, 1970, 00:00:00 GMT. | [optional] 
**Expires** | Pointer to **time.Time** | (sortable) The date at which the certificate expires as the number of milliseconds since January 1, 1970, 00:00:00 GMT. | [optional] 
**SignatureAlgorithm** | **string** | (sortable) The algorithm used to sign the certificate. | 
**Status** | [**CertStatus**](CertStatus.md) |  | 
**Sha1sum** | **string** | The SHA1 sum for the certificate. | 
**Md5sum** | **string** | The MD5 sum for the certificate. The value will be set to \&quot;\&quot; when in FIPS mode. | 
**Sha256sum** | **string** | The SHA256 sum for the certificate. | 
**SubjectAlternativeNames** | Pointer to [**[]GeneralName**](GeneralName.md) | A collection of subject alternative names for the certificate. | [optional] 

## Methods

### NewTrustedCert

`func NewTrustedCert(serialNumber string, alias string, subjectDn string, issuerDn string, signatureAlgorithm string, status CertStatus, sha1sum string, md5sum string, sha256sum string, ) *TrustedCert`

NewTrustedCert instantiates a new TrustedCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustedCertWithDefaults

`func NewTrustedCertWithDefaults() *TrustedCert`

NewTrustedCertWithDefaults instantiates a new TrustedCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrustedCert) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrustedCert) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrustedCert) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TrustedCert) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *TrustedCert) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *TrustedCert) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *TrustedCert) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetAlias

`func (o *TrustedCert) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *TrustedCert) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *TrustedCert) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetSubjectDn

`func (o *TrustedCert) GetSubjectDn() string`

GetSubjectDn returns the SubjectDn field if non-nil, zero value otherwise.

### GetSubjectDnOk

`func (o *TrustedCert) GetSubjectDnOk() (*string, bool)`

GetSubjectDnOk returns a tuple with the SubjectDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDn

`func (o *TrustedCert) SetSubjectDn(v string)`

SetSubjectDn sets SubjectDn field to given value.


### GetSubjectCn

`func (o *TrustedCert) GetSubjectCn() string`

GetSubjectCn returns the SubjectCn field if non-nil, zero value otherwise.

### GetSubjectCnOk

`func (o *TrustedCert) GetSubjectCnOk() (*string, bool)`

GetSubjectCnOk returns a tuple with the SubjectCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCn

`func (o *TrustedCert) SetSubjectCn(v string)`

SetSubjectCn sets SubjectCn field to given value.

### HasSubjectCn

`func (o *TrustedCert) HasSubjectCn() bool`

HasSubjectCn returns a boolean if a field has been set.

### GetIssuerDn

`func (o *TrustedCert) GetIssuerDn() string`

GetIssuerDn returns the IssuerDn field if non-nil, zero value otherwise.

### GetIssuerDnOk

`func (o *TrustedCert) GetIssuerDnOk() (*string, bool)`

GetIssuerDnOk returns a tuple with the IssuerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDn

`func (o *TrustedCert) SetIssuerDn(v string)`

SetIssuerDn sets IssuerDn field to given value.


### GetValidFrom

`func (o *TrustedCert) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *TrustedCert) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *TrustedCert) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *TrustedCert) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetExpires

`func (o *TrustedCert) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TrustedCert) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TrustedCert) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *TrustedCert) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *TrustedCert) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *TrustedCert) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *TrustedCert) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.


### GetStatus

`func (o *TrustedCert) GetStatus() CertStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrustedCert) GetStatusOk() (*CertStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrustedCert) SetStatus(v CertStatus)`

SetStatus sets Status field to given value.


### GetSha1sum

`func (o *TrustedCert) GetSha1sum() string`

GetSha1sum returns the Sha1sum field if non-nil, zero value otherwise.

### GetSha1sumOk

`func (o *TrustedCert) GetSha1sumOk() (*string, bool)`

GetSha1sumOk returns a tuple with the Sha1sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1sum

`func (o *TrustedCert) SetSha1sum(v string)`

SetSha1sum sets Sha1sum field to given value.


### GetMd5sum

`func (o *TrustedCert) GetMd5sum() string`

GetMd5sum returns the Md5sum field if non-nil, zero value otherwise.

### GetMd5sumOk

`func (o *TrustedCert) GetMd5sumOk() (*string, bool)`

GetMd5sumOk returns a tuple with the Md5sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5sum

`func (o *TrustedCert) SetMd5sum(v string)`

SetMd5sum sets Md5sum field to given value.


### GetSha256sum

`func (o *TrustedCert) GetSha256sum() string`

GetSha256sum returns the Sha256sum field if non-nil, zero value otherwise.

### GetSha256sumOk

`func (o *TrustedCert) GetSha256sumOk() (*string, bool)`

GetSha256sumOk returns a tuple with the Sha256sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256sum

`func (o *TrustedCert) SetSha256sum(v string)`

SetSha256sum sets Sha256sum field to given value.


### GetSubjectAlternativeNames

`func (o *TrustedCert) GetSubjectAlternativeNames() []GeneralName`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *TrustedCert) GetSubjectAlternativeNamesOk() (*[]GeneralName, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *TrustedCert) SetSubjectAlternativeNames(v []GeneralName)`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.

### HasSubjectAlternativeNames

`func (o *TrustedCert) HasSubjectAlternativeNames() bool`

HasSubjectAlternativeNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


