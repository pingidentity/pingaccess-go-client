# AgentCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The id for the agent certificate. | [optional] 
**SerialNumber** | **string** | (sortable) The Serial Number for the agent certificate. | 
**Alias** | **string** | (sortable) The alias for the agent certificate. | 
**SubjectDn** | **string** | (sortable) The Subject DN for the agent certificate. | 
**SubjectCn** | Pointer to **string** | (sortable) The common name (CN) identifying the certificate. | [optional] 
**IssuerDn** | **string** | (sortable) The issuer DN for the agent certificate. | 
**ValidFrom** | Pointer to **time.Time** | (sortable) The date at which the agent certificate is valid from as the number of milliseconds since January 1, 1970, 00:00:00 GMT. | [optional] 
**Expires** | Pointer to **time.Time** | (sortable) The date at which the agent certificate expires as the number of milliseconds since January 1, 1970, 00:00:00 GMT. | [optional] 
**SignatureAlgorithm** | **string** | (sortable) The Signature Algorithm used by the agent certificate. | 
**Status** | [**CertStatus**](CertStatus.md) |  | 
**Sha1sum** | **string** | The SHA1 checksum of the agent certificate. | 
**Md5sum** | **string** | The MD5 checksum of the agent certificate. The value will be set to \&quot;\&quot; when in FIPS mode. | 
**Sha256sum** | **string** | The SHA256 checksum of the agent certificate. | 
**SubjectAlternativeNames** | Pointer to [**[]GeneralName**](GeneralName.md) | A collection of subject alternative names for the agent certificate. | [optional] 
**TrustedCertificate** | **bool** | A flag indicating whether the agent certificate is a trusted certificate. | 
**KeyPair** | **bool** | A flag indicating whether the agent certificate is a key pair. | 
**ChainCertificate** | **bool** | A flag indicating whether the agent certificate is a chain certificate. | 

## Methods

### NewAgentCertificate

`func NewAgentCertificate(serialNumber string, alias string, subjectDn string, issuerDn string, signatureAlgorithm string, status CertStatus, sha1sum string, md5sum string, sha256sum string, trustedCertificate bool, keyPair bool, chainCertificate bool, ) *AgentCertificate`

NewAgentCertificate instantiates a new AgentCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCertificateWithDefaults

`func NewAgentCertificateWithDefaults() *AgentCertificate`

NewAgentCertificateWithDefaults instantiates a new AgentCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentCertificate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentCertificate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentCertificate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AgentCertificate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *AgentCertificate) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *AgentCertificate) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *AgentCertificate) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetAlias

`func (o *AgentCertificate) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *AgentCertificate) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *AgentCertificate) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetSubjectDn

`func (o *AgentCertificate) GetSubjectDn() string`

GetSubjectDn returns the SubjectDn field if non-nil, zero value otherwise.

### GetSubjectDnOk

`func (o *AgentCertificate) GetSubjectDnOk() (*string, bool)`

GetSubjectDnOk returns a tuple with the SubjectDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDn

`func (o *AgentCertificate) SetSubjectDn(v string)`

SetSubjectDn sets SubjectDn field to given value.


### GetSubjectCn

`func (o *AgentCertificate) GetSubjectCn() string`

GetSubjectCn returns the SubjectCn field if non-nil, zero value otherwise.

### GetSubjectCnOk

`func (o *AgentCertificate) GetSubjectCnOk() (*string, bool)`

GetSubjectCnOk returns a tuple with the SubjectCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCn

`func (o *AgentCertificate) SetSubjectCn(v string)`

SetSubjectCn sets SubjectCn field to given value.

### HasSubjectCn

`func (o *AgentCertificate) HasSubjectCn() bool`

HasSubjectCn returns a boolean if a field has been set.

### GetIssuerDn

`func (o *AgentCertificate) GetIssuerDn() string`

GetIssuerDn returns the IssuerDn field if non-nil, zero value otherwise.

### GetIssuerDnOk

`func (o *AgentCertificate) GetIssuerDnOk() (*string, bool)`

GetIssuerDnOk returns a tuple with the IssuerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDn

`func (o *AgentCertificate) SetIssuerDn(v string)`

SetIssuerDn sets IssuerDn field to given value.


### GetValidFrom

`func (o *AgentCertificate) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *AgentCertificate) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *AgentCertificate) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *AgentCertificate) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetExpires

`func (o *AgentCertificate) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *AgentCertificate) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *AgentCertificate) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *AgentCertificate) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *AgentCertificate) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *AgentCertificate) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *AgentCertificate) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.


### GetStatus

`func (o *AgentCertificate) GetStatus() CertStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentCertificate) GetStatusOk() (*CertStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentCertificate) SetStatus(v CertStatus)`

SetStatus sets Status field to given value.


### GetSha1sum

`func (o *AgentCertificate) GetSha1sum() string`

GetSha1sum returns the Sha1sum field if non-nil, zero value otherwise.

### GetSha1sumOk

`func (o *AgentCertificate) GetSha1sumOk() (*string, bool)`

GetSha1sumOk returns a tuple with the Sha1sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1sum

`func (o *AgentCertificate) SetSha1sum(v string)`

SetSha1sum sets Sha1sum field to given value.


### GetMd5sum

`func (o *AgentCertificate) GetMd5sum() string`

GetMd5sum returns the Md5sum field if non-nil, zero value otherwise.

### GetMd5sumOk

`func (o *AgentCertificate) GetMd5sumOk() (*string, bool)`

GetMd5sumOk returns a tuple with the Md5sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5sum

`func (o *AgentCertificate) SetMd5sum(v string)`

SetMd5sum sets Md5sum field to given value.


### GetSha256sum

`func (o *AgentCertificate) GetSha256sum() string`

GetSha256sum returns the Sha256sum field if non-nil, zero value otherwise.

### GetSha256sumOk

`func (o *AgentCertificate) GetSha256sumOk() (*string, bool)`

GetSha256sumOk returns a tuple with the Sha256sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256sum

`func (o *AgentCertificate) SetSha256sum(v string)`

SetSha256sum sets Sha256sum field to given value.


### GetSubjectAlternativeNames

`func (o *AgentCertificate) GetSubjectAlternativeNames() []GeneralName`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *AgentCertificate) GetSubjectAlternativeNamesOk() (*[]GeneralName, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *AgentCertificate) SetSubjectAlternativeNames(v []GeneralName)`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.

### HasSubjectAlternativeNames

`func (o *AgentCertificate) HasSubjectAlternativeNames() bool`

HasSubjectAlternativeNames returns a boolean if a field has been set.

### GetTrustedCertificate

`func (o *AgentCertificate) GetTrustedCertificate() bool`

GetTrustedCertificate returns the TrustedCertificate field if non-nil, zero value otherwise.

### GetTrustedCertificateOk

`func (o *AgentCertificate) GetTrustedCertificateOk() (*bool, bool)`

GetTrustedCertificateOk returns a tuple with the TrustedCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCertificate

`func (o *AgentCertificate) SetTrustedCertificate(v bool)`

SetTrustedCertificate sets TrustedCertificate field to given value.


### GetKeyPair

`func (o *AgentCertificate) GetKeyPair() bool`

GetKeyPair returns the KeyPair field if non-nil, zero value otherwise.

### GetKeyPairOk

`func (o *AgentCertificate) GetKeyPairOk() (*bool, bool)`

GetKeyPairOk returns a tuple with the KeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPair

`func (o *AgentCertificate) SetKeyPair(v bool)`

SetKeyPair sets KeyPair field to given value.


### GetChainCertificate

`func (o *AgentCertificate) GetChainCertificate() bool`

GetChainCertificate returns the ChainCertificate field if non-nil, zero value otherwise.

### GetChainCertificateOk

`func (o *AgentCertificate) GetChainCertificateOk() (*bool, bool)`

GetChainCertificateOk returns a tuple with the ChainCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainCertificate

`func (o *AgentCertificate) SetChainCertificate(v bool)`

SetChainCertificate sets ChainCertificate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


