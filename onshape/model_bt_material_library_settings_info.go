/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7489-3fe01473c812
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMaterialLibrarySettingsInfo struct for BTMaterialLibrarySettingsInfo
type BTMaterialLibrarySettingsInfo struct {
	CompanyLibraries []BTMaterialLibraryMetadataInfo `json:"companyLibraries,omitempty"`
	Libraries        []BTMaterialLibraryMetadataInfo `json:"libraries,omitempty"`
}

// NewBTMaterialLibrarySettingsInfo instantiates a new BTMaterialLibrarySettingsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMaterialLibrarySettingsInfo() *BTMaterialLibrarySettingsInfo {
	this := BTMaterialLibrarySettingsInfo{}
	return &this
}

// NewBTMaterialLibrarySettingsInfoWithDefaults instantiates a new BTMaterialLibrarySettingsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMaterialLibrarySettingsInfoWithDefaults() *BTMaterialLibrarySettingsInfo {
	this := BTMaterialLibrarySettingsInfo{}
	return &this
}

// GetCompanyLibraries returns the CompanyLibraries field value if set, zero value otherwise.
func (o *BTMaterialLibrarySettingsInfo) GetCompanyLibraries() []BTMaterialLibraryMetadataInfo {
	if o == nil || o.CompanyLibraries == nil {
		var ret []BTMaterialLibraryMetadataInfo
		return ret
	}
	return o.CompanyLibraries
}

// GetCompanyLibrariesOk returns a tuple with the CompanyLibraries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialLibrarySettingsInfo) GetCompanyLibrariesOk() ([]BTMaterialLibraryMetadataInfo, bool) {
	if o == nil || o.CompanyLibraries == nil {
		return nil, false
	}
	return o.CompanyLibraries, true
}

// HasCompanyLibraries returns a boolean if a field has been set.
func (o *BTMaterialLibrarySettingsInfo) HasCompanyLibraries() bool {
	if o != nil && o.CompanyLibraries != nil {
		return true
	}

	return false
}

// SetCompanyLibraries gets a reference to the given []BTMaterialLibraryMetadataInfo and assigns it to the CompanyLibraries field.
func (o *BTMaterialLibrarySettingsInfo) SetCompanyLibraries(v []BTMaterialLibraryMetadataInfo) {
	o.CompanyLibraries = v
}

// GetLibraries returns the Libraries field value if set, zero value otherwise.
func (o *BTMaterialLibrarySettingsInfo) GetLibraries() []BTMaterialLibraryMetadataInfo {
	if o == nil || o.Libraries == nil {
		var ret []BTMaterialLibraryMetadataInfo
		return ret
	}
	return o.Libraries
}

// GetLibrariesOk returns a tuple with the Libraries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialLibrarySettingsInfo) GetLibrariesOk() ([]BTMaterialLibraryMetadataInfo, bool) {
	if o == nil || o.Libraries == nil {
		return nil, false
	}
	return o.Libraries, true
}

// HasLibraries returns a boolean if a field has been set.
func (o *BTMaterialLibrarySettingsInfo) HasLibraries() bool {
	if o != nil && o.Libraries != nil {
		return true
	}

	return false
}

// SetLibraries gets a reference to the given []BTMaterialLibraryMetadataInfo and assigns it to the Libraries field.
func (o *BTMaterialLibrarySettingsInfo) SetLibraries(v []BTMaterialLibraryMetadataInfo) {
	o.Libraries = v
}

func (o BTMaterialLibrarySettingsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyLibraries != nil {
		toSerialize["companyLibraries"] = o.CompanyLibraries
	}
	if o.Libraries != nil {
		toSerialize["libraries"] = o.Libraries
	}
	return json.Marshal(toSerialize)
}

type NullableBTMaterialLibrarySettingsInfo struct {
	value *BTMaterialLibrarySettingsInfo
	isSet bool
}

func (v NullableBTMaterialLibrarySettingsInfo) Get() *BTMaterialLibrarySettingsInfo {
	return v.value
}

func (v *NullableBTMaterialLibrarySettingsInfo) Set(val *BTMaterialLibrarySettingsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMaterialLibrarySettingsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMaterialLibrarySettingsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMaterialLibrarySettingsInfo(val *BTMaterialLibrarySettingsInfo) *NullableBTMaterialLibrarySettingsInfo {
	return &NullableBTMaterialLibrarySettingsInfo{value: val, isSet: true}
}

func (v NullableBTMaterialLibrarySettingsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMaterialLibrarySettingsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
