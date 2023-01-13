# KeyPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The Id for the key pair. | [optional] 
**SerialNumber** | **string** | (sortable) The Serial Number for the key pair. | 
**Alias** | **string** | (sortable) The Alias for the key pair. | 
**SubjectDn** | **string** | (sortable) The Subject DN for the key pair. | 
**SubjectCn** | Pointer to **string** | (sortable) The common name (CN) identifying the certificate. | [optional] 
**IssuerDn** | **string** | (sortable) The Issuer DN for the key pair. | 
**ValidFrom** | Pointer to **time.Time** | (sortable) The date at which the key pair is valid from as the number of milliseconds since January 1, 1970, 00:00:00 GMT. | [optional] 
**Expires** | Pointer to **time.Time** | (sortable) The date at which the key pair expires as the number of milliseconds since January 1, 1970, 00:00:00 GMT. | [optional] 
**SignatureAlgorithm** | **string** | (sortable) The Signature Algorithm used by the key pair. | 
**Status** | [**CertStatus**](CertStatus.md) |  | 
**Sha1sum** | **string** | The SHA1 checksum of the key pair. | 
**Md5sum** | **string** | The MD5 checksum of the key pair. The value will be set to \&quot;\&quot; when in FIPS mode. | 
**Sha256sum** | **string** | The SHA256 checksum of the key pair. | 
**SubjectAlternativeNames** | Pointer to [**[]GeneralName**](GeneralName.md) | A collection of subject alternative names for the certificate. | [optional] 
**CsrPending** | **bool** | (sortable) True if a CSR is generated for this key pair. | 
**HsmProviderId** | Pointer to **int32** | The HSM Provider ID.  The default value is 0 indicating an HSM is not used for this KeyPair. | [optional] 
**ChainCertificates** | Pointer to [**[]ChainCertificate**](ChainCertificate.md) | The complete list of certificates in the key pair certificate chain. | [optional] 

## Methods

### NewKeyPair

`func NewKeyPair(serialNumber string, alias string, subjectDn string, issuerDn string, signatureAlgorithm string, status CertStatus, sha1sum string, md5sum string, sha256sum string, csrPending bool, ) *KeyPair`

NewKeyPair instantiates a new KeyPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyPairWithDefaults

`func NewKeyPairWithDefaults() *KeyPair`

NewKeyPairWithDefaults instantiates a new KeyPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyPair) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyPair) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyPair) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyPair) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *KeyPair) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *KeyPair) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *KeyPair) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetAlias

`func (o *KeyPair) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *KeyPair) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *KeyPair) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetSubjectDn

`func (o *KeyPair) GetSubjectDn() string`

GetSubjectDn returns the SubjectDn field if non-nil, zero value otherwise.

### GetSubjectDnOk

`func (o *KeyPair) GetSubjectDnOk() (*string, bool)`

GetSubjectDnOk returns a tuple with the SubjectDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDn

`func (o *KeyPair) SetSubjectDn(v string)`

SetSubjectDn sets SubjectDn field to given value.


### GetSubjectCn

`func (o *KeyPair) GetSubjectCn() string`

GetSubjectCn returns the SubjectCn field if non-nil, zero value otherwise.

### GetSubjectCnOk

`func (o *KeyPair) GetSubjectCnOk() (*string, bool)`

GetSubjectCnOk returns a tuple with the SubjectCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCn

`func (o *KeyPair) SetSubjectCn(v string)`

SetSubjectCn sets SubjectCn field to given value.

### HasSubjectCn

`func (o *KeyPair) HasSubjectCn() bool`

HasSubjectCn returns a boolean if a field has been set.

### GetIssuerDn

`func (o *KeyPair) GetIssuerDn() string`

GetIssuerDn returns the IssuerDn field if non-nil, zero value otherwise.

### GetIssuerDnOk

`func (o *KeyPair) GetIssuerDnOk() (*string, bool)`

GetIssuerDnOk returns a tuple with the IssuerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDn

`func (o *KeyPair) SetIssuerDn(v string)`

SetIssuerDn sets IssuerDn field to given value.


### GetValidFrom

`func (o *KeyPair) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *KeyPair) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *KeyPair) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *KeyPair) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetExpires

`func (o *KeyPair) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *KeyPair) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *KeyPair) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *KeyPair) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *KeyPair) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *KeyPair) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *KeyPair) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.


### GetStatus

`func (o *KeyPair) GetStatus() CertStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KeyPair) GetStatusOk() (*CertStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KeyPair) SetStatus(v CertStatus)`

SetStatus sets Status field to given value.


### GetSha1sum

`func (o *KeyPair) GetSha1sum() string`

GetSha1sum returns the Sha1sum field if non-nil, zero value otherwise.

### GetSha1sumOk

`func (o *KeyPair) GetSha1sumOk() (*string, bool)`

GetSha1sumOk returns a tuple with the Sha1sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1sum

`func (o *KeyPair) SetSha1sum(v string)`

SetSha1sum sets Sha1sum field to given value.


### GetMd5sum

`func (o *KeyPair) GetMd5sum() string`

GetMd5sum returns the Md5sum field if non-nil, zero value otherwise.

### GetMd5sumOk

`func (o *KeyPair) GetMd5sumOk() (*string, bool)`

GetMd5sumOk returns a tuple with the Md5sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5sum

`func (o *KeyPair) SetMd5sum(v string)`

SetMd5sum sets Md5sum field to given value.


### GetSha256sum

`func (o *KeyPair) GetSha256sum() string`

GetSha256sum returns the Sha256sum field if non-nil, zero value otherwise.

### GetSha256sumOk

`func (o *KeyPair) GetSha256sumOk() (*string, bool)`

GetSha256sumOk returns a tuple with the Sha256sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256sum

`func (o *KeyPair) SetSha256sum(v string)`

SetSha256sum sets Sha256sum field to given value.


### GetSubjectAlternativeNames

`func (o *KeyPair) GetSubjectAlternativeNames() []GeneralName`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *KeyPair) GetSubjectAlternativeNamesOk() (*[]GeneralName, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *KeyPair) SetSubjectAlternativeNames(v []GeneralName)`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.

### HasSubjectAlternativeNames

`func (o *KeyPair) HasSubjectAlternativeNames() bool`

HasSubjectAlternativeNames returns a boolean if a field has been set.

### GetCsrPending

`func (o *KeyPair) GetCsrPending() bool`

GetCsrPending returns the CsrPending field if non-nil, zero value otherwise.

### GetCsrPendingOk

`func (o *KeyPair) GetCsrPendingOk() (*bool, bool)`

GetCsrPendingOk returns a tuple with the CsrPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrPending

`func (o *KeyPair) SetCsrPending(v bool)`

SetCsrPending sets CsrPending field to given value.


### GetHsmProviderId

`func (o *KeyPair) GetHsmProviderId() int32`

GetHsmProviderId returns the HsmProviderId field if non-nil, zero value otherwise.

### GetHsmProviderIdOk

`func (o *KeyPair) GetHsmProviderIdOk() (*int32, bool)`

GetHsmProviderIdOk returns a tuple with the HsmProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmProviderId

`func (o *KeyPair) SetHsmProviderId(v int32)`

SetHsmProviderId sets HsmProviderId field to given value.

### HasHsmProviderId

`func (o *KeyPair) HasHsmProviderId() bool`

HasHsmProviderId returns a boolean if a field has been set.

### GetChainCertificates

`func (o *KeyPair) GetChainCertificates() []ChainCertificate`

GetChainCertificates returns the ChainCertificates field if non-nil, zero value otherwise.

### GetChainCertificatesOk

`func (o *KeyPair) GetChainCertificatesOk() (*[]ChainCertificate, bool)`

GetChainCertificatesOk returns a tuple with the ChainCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainCertificates

`func (o *KeyPair) SetChainCertificates(v []ChainCertificate)`

SetChainCertificates sets ChainCertificates field to given value.

### HasChainCertificates

`func (o *KeyPair) HasChainCertificates() bool`

HasChainCertificates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


