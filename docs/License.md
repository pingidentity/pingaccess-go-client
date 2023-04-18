# License

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tier** | **string** | The tier value from the license file. | 
**MaxApplications** | **int64** | The maximum number of applications from the license file. | 
**Product** | **string** | The Ping Identity product value from the license file. | 
**Version** | **string** | The Ping Identity product version from the license file. | 
**Organization** | **string** | The organization value from the license file. | 
**EnforcementType** | **int64** | The enforcement type value from the license file. | 
**ExpirationDate** | **string** | The expiration date value from the license file. | 
**IssueDate** | **string** | The issue date value from the license file. | 
**Name** | **string** | The name value from the license file.  Name is required if the organization value is not set. | 
**Id** | **int64** | The ID value from the license file. | 

## Methods

### NewLicense

`func NewLicense(tier string, maxApplications int64, product string, version string, organization string, enforcementType int64, expirationDate string, issueDate string, name string, id int64, ) *License`

NewLicense instantiates a new License object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseWithDefaults

`func NewLicenseWithDefaults() *License`

NewLicenseWithDefaults instantiates a new License object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTier

`func (o *License) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *License) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *License) SetTier(v string)`

SetTier sets Tier field to given value.


### GetMaxApplications

`func (o *License) GetMaxApplications() int64`

GetMaxApplications returns the MaxApplications field if non-nil, zero value otherwise.

### GetMaxApplicationsOk

`func (o *License) GetMaxApplicationsOk() (*int64, bool)`

GetMaxApplicationsOk returns a tuple with the MaxApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxApplications

`func (o *License) SetMaxApplications(v int64)`

SetMaxApplications sets MaxApplications field to given value.


### GetProduct

`func (o *License) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *License) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *License) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetVersion

`func (o *License) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *License) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *License) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetOrganization

`func (o *License) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *License) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *License) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetEnforcementType

`func (o *License) GetEnforcementType() int64`

GetEnforcementType returns the EnforcementType field if non-nil, zero value otherwise.

### GetEnforcementTypeOk

`func (o *License) GetEnforcementTypeOk() (*int64, bool)`

GetEnforcementTypeOk returns a tuple with the EnforcementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementType

`func (o *License) SetEnforcementType(v int64)`

SetEnforcementType sets EnforcementType field to given value.


### GetExpirationDate

`func (o *License) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *License) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *License) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.


### GetIssueDate

`func (o *License) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *License) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *License) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetName

`func (o *License) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *License) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *License) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *License) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *License) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *License) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


