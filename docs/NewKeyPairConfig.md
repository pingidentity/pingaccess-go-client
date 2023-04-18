# NewKeyPairConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID for the key pair. If not specified, an ID will be automatically assigned. | [optional] 
**Alias** | **string** | A unique alias name to identify the key pair. Special characters and spaces are allowed. | 
**CommonName** | **string** | The common name (CN) identifying the certificate. | 
**Organization** | **string** | The organization (O) or company name creating the certificate. | 
**OrganizationUnit** | **string** | The specific unit within the organization (OU). | 
**City** | **string** | The city or other primary location (L) where the company operates. | 
**State** | **string** | The state (ST) or other political unit encompassing the location. | 
**Country** | **string** | The country (C) where the company is based, using two capital letters. | 
**ValidDays** | **int64** | The number of days the certificate is valid. | 
**KeyAlgorithm** | **string** | The key algorithm to use to generate a key. | 
**SignatureAlgorithm** | Pointer to **string** | The Signature Algorithm to use for the key. | [optional] 
**SubjectAlternativeNames** | Pointer to [**[]GeneralName**](GeneralName.md) | Any additional DNS names or IP addresses that are valid for this certificate. | [optional] 
**KeySize** | **int64** | The number of bits used in the key.  Choices depend on selected key algorithm. | 
**HsmProviderId** | **int64** | The HSM Provider ID. The default value is 0 indicating an HSM is not used for this key pair. | 

## Methods

### NewNewKeyPairConfig

`func NewNewKeyPairConfig(alias string, commonName string, organization string, organizationUnit string, city string, state string, country string, validDays int64, keyAlgorithm string, keySize int64, hsmProviderId int64, ) *NewKeyPairConfig`

NewNewKeyPairConfig instantiates a new NewKeyPairConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewKeyPairConfigWithDefaults

`func NewNewKeyPairConfigWithDefaults() *NewKeyPairConfig`

NewNewKeyPairConfigWithDefaults instantiates a new NewKeyPairConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NewKeyPairConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NewKeyPairConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NewKeyPairConfig) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NewKeyPairConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlias

`func (o *NewKeyPairConfig) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *NewKeyPairConfig) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *NewKeyPairConfig) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetCommonName

`func (o *NewKeyPairConfig) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *NewKeyPairConfig) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *NewKeyPairConfig) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetOrganization

`func (o *NewKeyPairConfig) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *NewKeyPairConfig) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *NewKeyPairConfig) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetOrganizationUnit

`func (o *NewKeyPairConfig) GetOrganizationUnit() string`

GetOrganizationUnit returns the OrganizationUnit field if non-nil, zero value otherwise.

### GetOrganizationUnitOk

`func (o *NewKeyPairConfig) GetOrganizationUnitOk() (*string, bool)`

GetOrganizationUnitOk returns a tuple with the OrganizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUnit

`func (o *NewKeyPairConfig) SetOrganizationUnit(v string)`

SetOrganizationUnit sets OrganizationUnit field to given value.


### GetCity

`func (o *NewKeyPairConfig) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *NewKeyPairConfig) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *NewKeyPairConfig) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *NewKeyPairConfig) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NewKeyPairConfig) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NewKeyPairConfig) SetState(v string)`

SetState sets State field to given value.


### GetCountry

`func (o *NewKeyPairConfig) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *NewKeyPairConfig) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *NewKeyPairConfig) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetValidDays

`func (o *NewKeyPairConfig) GetValidDays() int64`

GetValidDays returns the ValidDays field if non-nil, zero value otherwise.

### GetValidDaysOk

`func (o *NewKeyPairConfig) GetValidDaysOk() (*int64, bool)`

GetValidDaysOk returns a tuple with the ValidDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDays

`func (o *NewKeyPairConfig) SetValidDays(v int64)`

SetValidDays sets ValidDays field to given value.


### GetKeyAlgorithm

`func (o *NewKeyPairConfig) GetKeyAlgorithm() string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *NewKeyPairConfig) GetKeyAlgorithmOk() (*string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *NewKeyPairConfig) SetKeyAlgorithm(v string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.


### GetSignatureAlgorithm

`func (o *NewKeyPairConfig) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *NewKeyPairConfig) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *NewKeyPairConfig) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *NewKeyPairConfig) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetSubjectAlternativeNames

`func (o *NewKeyPairConfig) GetSubjectAlternativeNames() []GeneralName`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *NewKeyPairConfig) GetSubjectAlternativeNamesOk() (*[]GeneralName, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *NewKeyPairConfig) SetSubjectAlternativeNames(v []GeneralName)`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.

### HasSubjectAlternativeNames

`func (o *NewKeyPairConfig) HasSubjectAlternativeNames() bool`

HasSubjectAlternativeNames returns a boolean if a field has been set.

### GetKeySize

`func (o *NewKeyPairConfig) GetKeySize() int64`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *NewKeyPairConfig) GetKeySizeOk() (*int64, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *NewKeyPairConfig) SetKeySize(v int64)`

SetKeySize sets KeySize field to given value.


### GetHsmProviderId

`func (o *NewKeyPairConfig) GetHsmProviderId() int64`

GetHsmProviderId returns the HsmProviderId field if non-nil, zero value otherwise.

### GetHsmProviderIdOk

`func (o *NewKeyPairConfig) GetHsmProviderIdOk() (*int64, bool)`

GetHsmProviderIdOk returns a tuple with the HsmProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmProviderId

`func (o *NewKeyPairConfig) SetHsmProviderId(v int64)`

SetHsmProviderId sets HsmProviderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


