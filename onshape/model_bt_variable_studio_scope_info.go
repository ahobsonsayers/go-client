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

// BTVariableStudioScopeInfo struct for BTVariableStudioScopeInfo
type BTVariableStudioScopeInfo struct {
	// Whether variable studio is automatically inserted into part studios and assemblies in the workspace
	IsAutomaticallyInserted bool `json:"isAutomaticallyInserted"`
}

// NewBTVariableStudioScopeInfo instantiates a new BTVariableStudioScopeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTVariableStudioScopeInfo(isAutomaticallyInserted bool) *BTVariableStudioScopeInfo {
	this := BTVariableStudioScopeInfo{}
	this.IsAutomaticallyInserted = isAutomaticallyInserted
	return &this
}

// NewBTVariableStudioScopeInfoWithDefaults instantiates a new BTVariableStudioScopeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTVariableStudioScopeInfoWithDefaults() *BTVariableStudioScopeInfo {
	this := BTVariableStudioScopeInfo{}
	return &this
}

// GetIsAutomaticallyInserted returns the IsAutomaticallyInserted field value
func (o *BTVariableStudioScopeInfo) GetIsAutomaticallyInserted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAutomaticallyInserted
}

// GetIsAutomaticallyInsertedOk returns a tuple with the IsAutomaticallyInserted field value
// and a boolean to check if the value has been set.
func (o *BTVariableStudioScopeInfo) GetIsAutomaticallyInsertedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAutomaticallyInserted, true
}

// SetIsAutomaticallyInserted sets field value
func (o *BTVariableStudioScopeInfo) SetIsAutomaticallyInserted(v bool) {
	o.IsAutomaticallyInserted = v
}

func (o BTVariableStudioScopeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isAutomaticallyInserted"] = o.IsAutomaticallyInserted
	}
	return json.Marshal(toSerialize)
}

type NullableBTVariableStudioScopeInfo struct {
	value *BTVariableStudioScopeInfo
	isSet bool
}

func (v NullableBTVariableStudioScopeInfo) Get() *BTVariableStudioScopeInfo {
	return v.value
}

func (v *NullableBTVariableStudioScopeInfo) Set(val *BTVariableStudioScopeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTVariableStudioScopeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTVariableStudioScopeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTVariableStudioScopeInfo(val *BTVariableStudioScopeInfo) *NullableBTVariableStudioScopeInfo {
	return &NullableBTVariableStudioScopeInfo{value: val, isSet: true}
}

func (v NullableBTVariableStudioScopeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTVariableStudioScopeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
