/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5944-34bf93ccfb05
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDocumentSelectorInfo struct for BTDocumentSelectorInfo
type BTDocumentSelectorInfo struct {
	Parameters *BTDocumentSelectorParametersInfo `json:"parameters,omitempty"`
	SelectorId *string                           `json:"selectorId,omitempty"`
}

// NewBTDocumentSelectorInfo instantiates a new BTDocumentSelectorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentSelectorInfo() *BTDocumentSelectorInfo {
	this := BTDocumentSelectorInfo{}
	return &this
}

// NewBTDocumentSelectorInfoWithDefaults instantiates a new BTDocumentSelectorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentSelectorInfoWithDefaults() *BTDocumentSelectorInfo {
	this := BTDocumentSelectorInfo{}
	return &this
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTDocumentSelectorInfo) GetParameters() BTDocumentSelectorParametersInfo {
	if o == nil || o.Parameters == nil {
		var ret BTDocumentSelectorParametersInfo
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSelectorInfo) GetParametersOk() (*BTDocumentSelectorParametersInfo, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BTDocumentSelectorInfo) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given BTDocumentSelectorParametersInfo and assigns it to the Parameters field.
func (o *BTDocumentSelectorInfo) SetParameters(v BTDocumentSelectorParametersInfo) {
	o.Parameters = &v
}

// GetSelectorId returns the SelectorId field value if set, zero value otherwise.
func (o *BTDocumentSelectorInfo) GetSelectorId() string {
	if o == nil || o.SelectorId == nil {
		var ret string
		return ret
	}
	return *o.SelectorId
}

// GetSelectorIdOk returns a tuple with the SelectorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSelectorInfo) GetSelectorIdOk() (*string, bool) {
	if o == nil || o.SelectorId == nil {
		return nil, false
	}
	return o.SelectorId, true
}

// HasSelectorId returns a boolean if a field has been set.
func (o *BTDocumentSelectorInfo) HasSelectorId() bool {
	if o != nil && o.SelectorId != nil {
		return true
	}

	return false
}

// SetSelectorId gets a reference to the given string and assigns it to the SelectorId field.
func (o *BTDocumentSelectorInfo) SetSelectorId(v string) {
	o.SelectorId = &v
}

func (o BTDocumentSelectorInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.SelectorId != nil {
		toSerialize["selectorId"] = o.SelectorId
	}
	return json.Marshal(toSerialize)
}

type NullableBTDocumentSelectorInfo struct {
	value *BTDocumentSelectorInfo
	isSet bool
}

func (v NullableBTDocumentSelectorInfo) Get() *BTDocumentSelectorInfo {
	return v.value
}

func (v *NullableBTDocumentSelectorInfo) Set(val *BTDocumentSelectorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentSelectorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentSelectorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentSelectorInfo(val *BTDocumentSelectorInfo) *NullableBTDocumentSelectorInfo {
	return &NullableBTDocumentSelectorInfo{value: val, isSet: true}
}

func (v NullableBTDocumentSelectorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentSelectorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
