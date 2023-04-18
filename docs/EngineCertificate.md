# EngineCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The id for the engine certificate. | [optional] 
**SerialNumber** | **string** | (sortable) The Serial Number for the engine certificate. | 
**Alias** | **string** | (sortable) The alias for the engine certificate. | 
**SubjectDn** | **string** | (sortable) The Subject DN for the engine certificate. | 
**SubjectCn** | Pointer to **string** | (sortable) The common name (CN) identifying the certificate. | [optional] 
**IssuerDn** | **string** | (sortable) The issuer DN for the engine certificate. | 
**ValidFrom** | Pointer to **time.Time** | (sortable) The date at which the engine certificate is valid from as the number of milliseconds since January 1, 1970, 00:00:00 GMT. | [optional] 
**Expires** | Pointer to **time.Time** | (sortable) The date at which the engine certificate expires as the number of milliseconds since January 1, 1970, 00:00:00 GMT. | [optional] 
**SignatureAlgorithm** | **string** | (sortable) The Signature Algorithm used by the engine certificate. | 
**Status** | [**CertStatus**](CertStatus.md) |  | 
**Sha1sum** | **string** | The SHA1 checksum of the engine certificate. | 
**Md5sum** | **string** | The MD5 checksum of the engine certificate. The value will be set to \&quot;\&quot; when in FIPS mode. | 
**Sha256sum** | **string** | The SHA256 checksum of the engine certificate. | 
**SubjectAlternativeNames** | Pointer to [**[]GeneralName**](GeneralName.md) | A collection of subject alternative names for the engine certificate. | [optional] 
**TrustedCertificate** | **bool** | A flag indicating whether the engine certificate is a trusted certificate. | 
**KeyPair** | **bool** | A flag indicating whether the engine certificate is a key pair. | 
**ChainCertificate** | **bool** | A flag indicating whether the engine certificate is a chain certificate. | 

## Methods

### NewEngineCertificate

`func NewEngineCertificate(serialNumber string, alias string, subjectDn string, issuerDn string, signatureAlgorithm string, status CertStatus, sha1sum string, md5sum string, sha256sum string, trustedCertificate bool, keyPair bool, chainCertificate bool, ) *EngineCertificate`

NewEngineCertificate instantiates a new EngineCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineCertificateWithDefaults

`func NewEngineCertificateWithDefaults() *EngineCertificate`

NewEngineCertificateWithDefaults instantiates a new EngineCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EngineCertificate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EngineCertificate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EngineCertificate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EngineCertificate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *EngineCertificate) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *EngineCertificate) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *EngineCertificate) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetAlias

`func (o *EngineCertificate) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *EngineCertificate) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *EngineCertificate) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetSubjectDn

`func (o *EngineCertificate) GetSubjectDn() string`

GetSubjectDn returns the SubjectDn field if non-nil, zero value otherwise.

### GetSubjectDnOk

`func (o *EngineCertificate) GetSubjectDnOk() (*string, bool)`

GetSubjectDnOk returns a tuple with the SubjectDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDn

`func (o *EngineCertificate) SetSubjectDn(v string)`

SetSubjectDn sets SubjectDn field to given value.


### GetSubjectCn

`func (o *EngineCertificate) GetSubjectCn() string`

GetSubjectCn returns the SubjectCn field if non-nil, zero value otherwise.

### GetSubjectCnOk

`func (o *EngineCertificate) GetSubjectCnOk() (*string, bool)`

GetSubjectCnOk returns a tuple with the SubjectCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCn

`func (o *EngineCertificate) SetSubjectCn(v string)`

SetSubjectCn sets SubjectCn field to given value.

### HasSubjectCn

`func (o *EngineCertificate) HasSubjectCn() bool`

HasSubjectCn returns a boolean if a field has been set.

### GetIssuerDn

`func (o *EngineCertificate) GetIssuerDn() string`

GetIssuerDn returns the IssuerDn field if non-nil, zero value otherwise.

### GetIssuerDnOk

`func (o *EngineCertificate) GetIssuerDnOk() (*string, bool)`

GetIssuerDnOk returns a tuple with the IssuerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDn

`func (o *EngineCertificate) SetIssuerDn(v string)`

SetIssuerDn sets IssuerDn field to given value.


### GetValidFrom

`func (o *EngineCertificate) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *EngineCertificate) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *EngineCertificate) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *EngineCertificate) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetExpires

`func (o *EngineCertificate) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *EngineCertificate) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *EngineCertificate) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *EngineCertificate) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *EngineCertificate) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *EngineCertificate) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *EngineCertificate) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.


### GetStatus

`func (o *EngineCertificate) GetStatus() CertStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EngineCertificate) GetStatusOk() (*CertStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EngineCertificate) SetStatus(v CertStatus)`

SetStatus sets Status field to given value.


### GetSha1sum

`func (o *EngineCertificate) GetSha1sum() string`

GetSha1sum returns the Sha1sum field if non-nil, zero value otherwise.

### GetSha1sumOk

`func (o *EngineCertificate) GetSha1sumOk() (*string, bool)`

GetSha1sumOk returns a tuple with the Sha1sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1sum

`func (o *EngineCertificate) SetSha1sum(v string)`

SetSha1sum sets Sha1sum field to given value.


### GetMd5sum

`func (o *EngineCertificate) GetMd5sum() string`

GetMd5sum returns the Md5sum field if non-nil, zero value otherwise.

### GetMd5sumOk

`func (o *EngineCertificate) GetMd5sumOk() (*string, bool)`

GetMd5sumOk returns a tuple with the Md5sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5sum

`func (o *EngineCertificate) SetMd5sum(v string)`

SetMd5sum sets Md5sum field to given value.


### GetSha256sum

`func (o *EngineCertificate) GetSha256sum() string`

GetSha256sum returns the Sha256sum field if non-nil, zero value otherwise.

### GetSha256sumOk

`func (o *EngineCertificate) GetSha256sumOk() (*string, bool)`

GetSha256sumOk returns a tuple with the Sha256sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256sum

`func (o *EngineCertificate) SetSha256sum(v string)`

SetSha256sum sets Sha256sum field to given value.


### GetSubjectAlternativeNames

`func (o *EngineCertificate) GetSubjectAlternativeNames() []GeneralName`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *EngineCertificate) GetSubjectAlternativeNamesOk() (*[]GeneralName, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *EngineCertificate) SetSubjectAlternativeNames(v []GeneralName)`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.

### HasSubjectAlternativeNames

`func (o *EngineCertificate) HasSubjectAlternativeNames() bool`

HasSubjectAlternativeNames returns a boolean if a field has been set.

### GetTrustedCertificate

`func (o *EngineCertificate) GetTrustedCertificate() bool`

GetTrustedCertificate returns the TrustedCertificate field if non-nil, zero value otherwise.

### GetTrustedCertificateOk

`func (o *EngineCertificate) GetTrustedCertificateOk() (*bool, bool)`

GetTrustedCertificateOk returns a tuple with the TrustedCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCertificate

`func (o *EngineCertificate) SetTrustedCertificate(v bool)`

SetTrustedCertificate sets TrustedCertificate field to given value.


### GetKeyPair

`func (o *EngineCertificate) GetKeyPair() bool`

GetKeyPair returns the KeyPair field if non-nil, zero value otherwise.

### GetKeyPairOk

`func (o *EngineCertificate) GetKeyPairOk() (*bool, bool)`

GetKeyPairOk returns a tuple with the KeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPair

`func (o *EngineCertificate) SetKeyPair(v bool)`

SetKeyPair sets KeyPair field to given value.


### GetChainCertificate

`func (o *EngineCertificate) GetChainCertificate() bool`

GetChainCertificate returns the ChainCertificate field if non-nil, zero value otherwise.

### GetChainCertificateOk

`func (o *EngineCertificate) GetChainCertificateOk() (*bool, bool)`

GetChainCertificateOk returns a tuple with the ChainCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainCertificate

`func (o *EngineCertificate) SetChainCertificate(v bool)`

SetChainCertificate sets ChainCertificate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


