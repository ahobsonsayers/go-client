# BTUserSummaryInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to [**BTCompanySummaryInfo**](BTCompanySummaryInfo.md) |  | [optional] 
**DocumentationName** | Pointer to **string** |  | [optional] 
**DocumentationNameOverride** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**GlobalPermissions** | Pointer to [**GlobalPermissionInfo**](GlobalPermissionInfo.md) |  | [optional] 
**IsGuest** | Pointer to **bool** |  | [optional] 
**IsLight** | Pointer to **bool** |  | [optional] 
**LastLoginTime** | Pointer to **JSONTime** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**PersonalMessageAllowed** | Pointer to **bool** |  | [optional] 
**Source** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTUserSummaryInfoAllOf

`func NewBTUserSummaryInfoAllOf() *BTUserSummaryInfoAllOf`

NewBTUserSummaryInfoAllOf instantiates a new BTUserSummaryInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTUserSummaryInfoAllOfWithDefaults

`func NewBTUserSummaryInfoAllOfWithDefaults() *BTUserSummaryInfoAllOf`

NewBTUserSummaryInfoAllOfWithDefaults instantiates a new BTUserSummaryInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *BTUserSummaryInfoAllOf) GetCompany() BTCompanySummaryInfo`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *BTUserSummaryInfoAllOf) GetCompanyOk() (*BTCompanySummaryInfo, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *BTUserSummaryInfoAllOf) SetCompany(v BTCompanySummaryInfo)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *BTUserSummaryInfoAllOf) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetDocumentationName

`func (o *BTUserSummaryInfoAllOf) GetDocumentationName() string`

GetDocumentationName returns the DocumentationName field if non-nil, zero value otherwise.

### GetDocumentationNameOk

`func (o *BTUserSummaryInfoAllOf) GetDocumentationNameOk() (*string, bool)`

GetDocumentationNameOk returns a tuple with the DocumentationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationName

`func (o *BTUserSummaryInfoAllOf) SetDocumentationName(v string)`

SetDocumentationName sets DocumentationName field to given value.

### HasDocumentationName

`func (o *BTUserSummaryInfoAllOf) HasDocumentationName() bool`

HasDocumentationName returns a boolean if a field has been set.

### GetDocumentationNameOverride

`func (o *BTUserSummaryInfoAllOf) GetDocumentationNameOverride() string`

GetDocumentationNameOverride returns the DocumentationNameOverride field if non-nil, zero value otherwise.

### GetDocumentationNameOverrideOk

`func (o *BTUserSummaryInfoAllOf) GetDocumentationNameOverrideOk() (*string, bool)`

GetDocumentationNameOverrideOk returns a tuple with the DocumentationNameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationNameOverride

`func (o *BTUserSummaryInfoAllOf) SetDocumentationNameOverride(v string)`

SetDocumentationNameOverride sets DocumentationNameOverride field to given value.

### HasDocumentationNameOverride

`func (o *BTUserSummaryInfoAllOf) HasDocumentationNameOverride() bool`

HasDocumentationNameOverride returns a boolean if a field has been set.

### GetEmail

`func (o *BTUserSummaryInfoAllOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BTUserSummaryInfoAllOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BTUserSummaryInfoAllOf) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BTUserSummaryInfoAllOf) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *BTUserSummaryInfoAllOf) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BTUserSummaryInfoAllOf) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BTUserSummaryInfoAllOf) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *BTUserSummaryInfoAllOf) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGlobalPermissions

`func (o *BTUserSummaryInfoAllOf) GetGlobalPermissions() GlobalPermissionInfo`

GetGlobalPermissions returns the GlobalPermissions field if non-nil, zero value otherwise.

### GetGlobalPermissionsOk

`func (o *BTUserSummaryInfoAllOf) GetGlobalPermissionsOk() (*GlobalPermissionInfo, bool)`

GetGlobalPermissionsOk returns a tuple with the GlobalPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPermissions

`func (o *BTUserSummaryInfoAllOf) SetGlobalPermissions(v GlobalPermissionInfo)`

SetGlobalPermissions sets GlobalPermissions field to given value.

### HasGlobalPermissions

`func (o *BTUserSummaryInfoAllOf) HasGlobalPermissions() bool`

HasGlobalPermissions returns a boolean if a field has been set.

### GetIsGuest

`func (o *BTUserSummaryInfoAllOf) GetIsGuest() bool`

GetIsGuest returns the IsGuest field if non-nil, zero value otherwise.

### GetIsGuestOk

`func (o *BTUserSummaryInfoAllOf) GetIsGuestOk() (*bool, bool)`

GetIsGuestOk returns a tuple with the IsGuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGuest

`func (o *BTUserSummaryInfoAllOf) SetIsGuest(v bool)`

SetIsGuest sets IsGuest field to given value.

### HasIsGuest

`func (o *BTUserSummaryInfoAllOf) HasIsGuest() bool`

HasIsGuest returns a boolean if a field has been set.

### GetIsLight

`func (o *BTUserSummaryInfoAllOf) GetIsLight() bool`

GetIsLight returns the IsLight field if non-nil, zero value otherwise.

### GetIsLightOk

`func (o *BTUserSummaryInfoAllOf) GetIsLightOk() (*bool, bool)`

GetIsLightOk returns a tuple with the IsLight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLight

`func (o *BTUserSummaryInfoAllOf) SetIsLight(v bool)`

SetIsLight sets IsLight field to given value.

### HasIsLight

`func (o *BTUserSummaryInfoAllOf) HasIsLight() bool`

HasIsLight returns a boolean if a field has been set.

### GetLastLoginTime

`func (o *BTUserSummaryInfoAllOf) GetLastLoginTime() JSONTime`

GetLastLoginTime returns the LastLoginTime field if non-nil, zero value otherwise.

### GetLastLoginTimeOk

`func (o *BTUserSummaryInfoAllOf) GetLastLoginTimeOk() (*JSONTime, bool)`

GetLastLoginTimeOk returns a tuple with the LastLoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTime

`func (o *BTUserSummaryInfoAllOf) SetLastLoginTime(v JSONTime)`

SetLastLoginTime sets LastLoginTime field to given value.

### HasLastLoginTime

`func (o *BTUserSummaryInfoAllOf) HasLastLoginTime() bool`

HasLastLoginTime returns a boolean if a field has been set.

### GetLastName

`func (o *BTUserSummaryInfoAllOf) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BTUserSummaryInfoAllOf) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BTUserSummaryInfoAllOf) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *BTUserSummaryInfoAllOf) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPersonalMessageAllowed

`func (o *BTUserSummaryInfoAllOf) GetPersonalMessageAllowed() bool`

GetPersonalMessageAllowed returns the PersonalMessageAllowed field if non-nil, zero value otherwise.

### GetPersonalMessageAllowedOk

`func (o *BTUserSummaryInfoAllOf) GetPersonalMessageAllowedOk() (*bool, bool)`

GetPersonalMessageAllowedOk returns a tuple with the PersonalMessageAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalMessageAllowed

`func (o *BTUserSummaryInfoAllOf) SetPersonalMessageAllowed(v bool)`

SetPersonalMessageAllowed sets PersonalMessageAllowed field to given value.

### HasPersonalMessageAllowed

`func (o *BTUserSummaryInfoAllOf) HasPersonalMessageAllowed() bool`

HasPersonalMessageAllowed returns a boolean if a field has been set.

### GetSource

`func (o *BTUserSummaryInfoAllOf) GetSource() int32`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BTUserSummaryInfoAllOf) GetSourceOk() (*int32, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BTUserSummaryInfoAllOf) SetSource(v int32)`

SetSource sets Source field to given value.

### HasSource

`func (o *BTUserSummaryInfoAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


