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

// BTAssemblyMateConnectorInfo struct for BTAssemblyMateConnectorInfo
type BTAssemblyMateConnectorInfo struct {
	FeatureId       *string                `json:"featureId,omitempty"`
	MateConnectorCS *BTMateConnectorCSInfo `json:"mateConnectorCS,omitempty"`
}

// NewBTAssemblyMateConnectorInfo instantiates a new BTAssemblyMateConnectorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyMateConnectorInfo() *BTAssemblyMateConnectorInfo {
	this := BTAssemblyMateConnectorInfo{}
	return &this
}

// NewBTAssemblyMateConnectorInfoWithDefaults instantiates a new BTAssemblyMateConnectorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyMateConnectorInfoWithDefaults() *BTAssemblyMateConnectorInfo {
	this := BTAssemblyMateConnectorInfo{}
	return &this
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTAssemblyMateConnectorInfo) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyMateConnectorInfo) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTAssemblyMateConnectorInfo) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTAssemblyMateConnectorInfo) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetMateConnectorCS returns the MateConnectorCS field value if set, zero value otherwise.
func (o *BTAssemblyMateConnectorInfo) GetMateConnectorCS() BTMateConnectorCSInfo {
	if o == nil || o.MateConnectorCS == nil {
		var ret BTMateConnectorCSInfo
		return ret
	}
	return *o.MateConnectorCS
}

// GetMateConnectorCSOk returns a tuple with the MateConnectorCS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyMateConnectorInfo) GetMateConnectorCSOk() (*BTMateConnectorCSInfo, bool) {
	if o == nil || o.MateConnectorCS == nil {
		return nil, false
	}
	return o.MateConnectorCS, true
}

// HasMateConnectorCS returns a boolean if a field has been set.
func (o *BTAssemblyMateConnectorInfo) HasMateConnectorCS() bool {
	if o != nil && o.MateConnectorCS != nil {
		return true
	}

	return false
}

// SetMateConnectorCS gets a reference to the given BTMateConnectorCSInfo and assigns it to the MateConnectorCS field.
func (o *BTAssemblyMateConnectorInfo) SetMateConnectorCS(v BTMateConnectorCSInfo) {
	o.MateConnectorCS = &v
}

func (o BTAssemblyMateConnectorInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.MateConnectorCS != nil {
		toSerialize["mateConnectorCS"] = o.MateConnectorCS
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyMateConnectorInfo struct {
	value *BTAssemblyMateConnectorInfo
	isSet bool
}

func (v NullableBTAssemblyMateConnectorInfo) Get() *BTAssemblyMateConnectorInfo {
	return v.value
}

func (v *NullableBTAssemblyMateConnectorInfo) Set(val *BTAssemblyMateConnectorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyMateConnectorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyMateConnectorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyMateConnectorInfo(val *BTAssemblyMateConnectorInfo) *NullableBTAssemblyMateConnectorInfo {
	return &NullableBTAssemblyMateConnectorInfo{value: val, isSet: true}
}

func (v NullableBTAssemblyMateConnectorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyMateConnectorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
