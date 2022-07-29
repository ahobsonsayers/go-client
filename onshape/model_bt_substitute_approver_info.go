/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5777-46062b6e03f9
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSubstituteApproverInfo struct for BTSubstituteApproverInfo
type BTSubstituteApproverInfo struct {
	CompanyId *string         `json:"companyId,omitempty"`
	Enabled   *bool           `json:"enabled,omitempty"`
	Identity  *BTIdentityInfo `json:"identity,omitempty"`
}

// NewBTSubstituteApproverInfo instantiates a new BTSubstituteApproverInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSubstituteApproverInfo() *BTSubstituteApproverInfo {
	this := BTSubstituteApproverInfo{}
	return &this
}

// NewBTSubstituteApproverInfoWithDefaults instantiates a new BTSubstituteApproverInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSubstituteApproverInfoWithDefaults() *BTSubstituteApproverInfo {
	this := BTSubstituteApproverInfo{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTSubstituteApproverInfo) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSubstituteApproverInfo) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTSubstituteApproverInfo) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTSubstituteApproverInfo) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BTSubstituteApproverInfo) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSubstituteApproverInfo) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BTSubstituteApproverInfo) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BTSubstituteApproverInfo) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *BTSubstituteApproverInfo) GetIdentity() BTIdentityInfo {
	if o == nil || o.Identity == nil {
		var ret BTIdentityInfo
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSubstituteApproverInfo) GetIdentityOk() (*BTIdentityInfo, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *BTSubstituteApproverInfo) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given BTIdentityInfo and assigns it to the Identity field.
func (o *BTSubstituteApproverInfo) SetIdentity(v BTIdentityInfo) {
	o.Identity = &v
}

func (o BTSubstituteApproverInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}
	return json.Marshal(toSerialize)
}

type NullableBTSubstituteApproverInfo struct {
	value *BTSubstituteApproverInfo
	isSet bool
}

func (v NullableBTSubstituteApproverInfo) Get() *BTSubstituteApproverInfo {
	return v.value
}

func (v *NullableBTSubstituteApproverInfo) Set(val *BTSubstituteApproverInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSubstituteApproverInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSubstituteApproverInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSubstituteApproverInfo(val *BTSubstituteApproverInfo) *NullableBTSubstituteApproverInfo {
	return &NullableBTSubstituteApproverInfo{value: val, isSet: true}
}

func (v NullableBTSubstituteApproverInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSubstituteApproverInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
