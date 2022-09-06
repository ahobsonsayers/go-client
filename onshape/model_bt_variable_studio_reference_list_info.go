/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6344-db89a80956dd
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTVariableStudioReferenceListInfo struct for BTVariableStudioReferenceListInfo
type BTVariableStudioReferenceListInfo struct {
	// List of variable studio references
	References []BTVariableStudioReferenceInfo `json:"references,omitempty"`
}

// NewBTVariableStudioReferenceListInfo instantiates a new BTVariableStudioReferenceListInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTVariableStudioReferenceListInfo() *BTVariableStudioReferenceListInfo {
	this := BTVariableStudioReferenceListInfo{}
	return &this
}

// NewBTVariableStudioReferenceListInfoWithDefaults instantiates a new BTVariableStudioReferenceListInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTVariableStudioReferenceListInfoWithDefaults() *BTVariableStudioReferenceListInfo {
	this := BTVariableStudioReferenceListInfo{}
	return &this
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *BTVariableStudioReferenceListInfo) GetReferences() []BTVariableStudioReferenceInfo {
	if o == nil || o.References == nil {
		var ret []BTVariableStudioReferenceInfo
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVariableStudioReferenceListInfo) GetReferencesOk() ([]BTVariableStudioReferenceInfo, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *BTVariableStudioReferenceListInfo) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []BTVariableStudioReferenceInfo and assigns it to the References field.
func (o *BTVariableStudioReferenceListInfo) SetReferences(v []BTVariableStudioReferenceInfo) {
	o.References = v
}

func (o BTVariableStudioReferenceListInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	return json.Marshal(toSerialize)
}

type NullableBTVariableStudioReferenceListInfo struct {
	value *BTVariableStudioReferenceListInfo
	isSet bool
}

func (v NullableBTVariableStudioReferenceListInfo) Get() *BTVariableStudioReferenceListInfo {
	return v.value
}

func (v *NullableBTVariableStudioReferenceListInfo) Set(val *BTVariableStudioReferenceListInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTVariableStudioReferenceListInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTVariableStudioReferenceListInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTVariableStudioReferenceListInfo(val *BTVariableStudioReferenceListInfo) *NullableBTVariableStudioReferenceListInfo {
	return &NullableBTVariableStudioReferenceListInfo{value: val, isSet: true}
}

func (v NullableBTVariableStudioReferenceListInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTVariableStudioReferenceListInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
