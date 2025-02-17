/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.166.18273-3025d52f81b7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTReleasePackageItemParams struct for BTReleasePackageItemParams
type BTReleasePackageItemParams struct {
	Configuration *string                `json:"configuration,omitempty"`
	DocumentId    *string                `json:"documentId,omitempty"`
	ElementId     *string                `json:"elementId,omitempty"`
	Href          *string                `json:"href,omitempty"`
	Id            *string                `json:"id,omitempty"`
	IsIncluded    *bool                  `json:"isIncluded,omitempty"`
	PartId        *string                `json:"partId,omitempty"`
	PartNumber    *string                `json:"partNumber,omitempty"`
	Properties    []BTPropertyValueParam `json:"properties,omitempty"`
	VersionId     *string                `json:"versionId,omitempty"`
	WorkspaceId   *string                `json:"workspaceId,omitempty"`
}

// NewBTReleasePackageItemParams instantiates a new BTReleasePackageItemParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTReleasePackageItemParams() *BTReleasePackageItemParams {
	this := BTReleasePackageItemParams{}
	return &this
}

// NewBTReleasePackageItemParamsWithDefaults instantiates a new BTReleasePackageItemParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTReleasePackageItemParamsWithDefaults() *BTReleasePackageItemParams {
	this := BTReleasePackageItemParams{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BTReleasePackageItemParams) GetConfiguration() string {
	if o == nil || o.Configuration == nil {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemParams) GetConfigurationOk() (*string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BTReleasePackageItemParams) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *BTReleasePackageItemParams) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTReleasePackageItemParams) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemParams) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTReleasePackageItemParams) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTReleasePackageItemParams) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTReleasePackageItemParams) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemParams) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTReleasePackageItemParams) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTReleasePackageItemParams) SetElementId(v string) {
	o.ElementId = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTReleasePackageItemParams) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemParams) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTReleasePackageItemParams) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTReleasePackageItemParams) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTReleasePackageItemParams) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemParams) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTReleasePackageItemParams) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTReleasePackageItemParams) SetId(v string) {
	o.Id = &v
}

// GetIsIncluded returns the IsIncluded field value if set, zero value otherwise.
func (o *BTReleasePackageItemParams) GetIsIncluded() bool {
	if o == nil || o.IsIncluded == nil {
		var ret bool
		return ret
	}
	return *o.IsIncluded
}

// GetIsIncludedOk returns a tuple with the IsIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemParams) GetIsIncludedOk() (*bool, bool) {
	if o == nil || o.IsIncluded == nil {
		return nil, false
	}
	return o.IsIncluded, true
}

// HasIsIncluded returns a boolean if a field has been set.
func (o *BTReleasePackageItemParams) HasIsIncluded() bool {
	if o != nil && o.IsIncluded != nil {
		return true
	}

	return false
}

// SetIsIncluded gets a reference to the given bool and assigns it to the IsIncluded field.
func (o *BTReleasePackageItemParams) SetIsIncluded(v bool) {
	o.IsIncluded = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTReleasePackageItemParams) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemParams) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTReleasePackageItemParams) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTReleasePackageItemParams) SetPartId(v string) {
	o.PartId = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTReleasePackageItemParams) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemParams) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTReleasePackageItemParams) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTReleasePackageItemParams) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *BTReleasePackageItemParams) GetProperties() []BTPropertyValueParam {
	if o == nil || o.Properties == nil {
		var ret []BTPropertyValueParam
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemParams) GetPropertiesOk() ([]BTPropertyValueParam, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *BTReleasePackageItemParams) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []BTPropertyValueParam and assigns it to the Properties field.
func (o *BTReleasePackageItemParams) SetProperties(v []BTPropertyValueParam) {
	o.Properties = v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTReleasePackageItemParams) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemParams) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTReleasePackageItemParams) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTReleasePackageItemParams) SetVersionId(v string) {
	o.VersionId = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *BTReleasePackageItemParams) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemParams) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *BTReleasePackageItemParams) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *BTReleasePackageItemParams) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o BTReleasePackageItemParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsIncluded != nil {
		toSerialize["isIncluded"] = o.IsIncluded
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	if o.WorkspaceId != nil {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableBTReleasePackageItemParams struct {
	value *BTReleasePackageItemParams
	isSet bool
}

func (v NullableBTReleasePackageItemParams) Get() *BTReleasePackageItemParams {
	return v.value
}

func (v *NullableBTReleasePackageItemParams) Set(val *BTReleasePackageItemParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTReleasePackageItemParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTReleasePackageItemParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTReleasePackageItemParams(val *BTReleasePackageItemParams) *NullableBTReleasePackageItemParams {
	return &NullableBTReleasePackageItemParams{value: val, isSet: true}
}

func (v NullableBTReleasePackageItemParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTReleasePackageItemParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
