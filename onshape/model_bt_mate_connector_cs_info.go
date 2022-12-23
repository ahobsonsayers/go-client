/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.157.9170-8cb36f16959d
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMateConnectorCSInfo struct for BTMateConnectorCSInfo
type BTMateConnectorCSInfo struct {
	GetxAxis []float64 `json:"getxAxis,omitempty"`
	GetyAxis []float64 `json:"getyAxis,omitempty"`
	GetzAxis []float64 `json:"getzAxis,omitempty"`
	Origin   []float64 `json:"origin,omitempty"`
}

// NewBTMateConnectorCSInfo instantiates a new BTMateConnectorCSInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMateConnectorCSInfo() *BTMateConnectorCSInfo {
	this := BTMateConnectorCSInfo{}
	return &this
}

// NewBTMateConnectorCSInfoWithDefaults instantiates a new BTMateConnectorCSInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMateConnectorCSInfoWithDefaults() *BTMateConnectorCSInfo {
	this := BTMateConnectorCSInfo{}
	return &this
}

// GetGetxAxis returns the GetxAxis field value if set, zero value otherwise.
func (o *BTMateConnectorCSInfo) GetGetxAxis() []float64 {
	if o == nil || o.GetxAxis == nil {
		var ret []float64
		return ret
	}
	return o.GetxAxis
}

// GetGetxAxisOk returns a tuple with the GetxAxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateConnectorCSInfo) GetGetxAxisOk() ([]float64, bool) {
	if o == nil || o.GetxAxis == nil {
		return nil, false
	}
	return o.GetxAxis, true
}

// HasGetxAxis returns a boolean if a field has been set.
func (o *BTMateConnectorCSInfo) HasGetxAxis() bool {
	if o != nil && o.GetxAxis != nil {
		return true
	}

	return false
}

// SetGetxAxis gets a reference to the given []float64 and assigns it to the GetxAxis field.
func (o *BTMateConnectorCSInfo) SetGetxAxis(v []float64) {
	o.GetxAxis = v
}

// GetGetyAxis returns the GetyAxis field value if set, zero value otherwise.
func (o *BTMateConnectorCSInfo) GetGetyAxis() []float64 {
	if o == nil || o.GetyAxis == nil {
		var ret []float64
		return ret
	}
	return o.GetyAxis
}

// GetGetyAxisOk returns a tuple with the GetyAxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateConnectorCSInfo) GetGetyAxisOk() ([]float64, bool) {
	if o == nil || o.GetyAxis == nil {
		return nil, false
	}
	return o.GetyAxis, true
}

// HasGetyAxis returns a boolean if a field has been set.
func (o *BTMateConnectorCSInfo) HasGetyAxis() bool {
	if o != nil && o.GetyAxis != nil {
		return true
	}

	return false
}

// SetGetyAxis gets a reference to the given []float64 and assigns it to the GetyAxis field.
func (o *BTMateConnectorCSInfo) SetGetyAxis(v []float64) {
	o.GetyAxis = v
}

// GetGetzAxis returns the GetzAxis field value if set, zero value otherwise.
func (o *BTMateConnectorCSInfo) GetGetzAxis() []float64 {
	if o == nil || o.GetzAxis == nil {
		var ret []float64
		return ret
	}
	return o.GetzAxis
}

// GetGetzAxisOk returns a tuple with the GetzAxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateConnectorCSInfo) GetGetzAxisOk() ([]float64, bool) {
	if o == nil || o.GetzAxis == nil {
		return nil, false
	}
	return o.GetzAxis, true
}

// HasGetzAxis returns a boolean if a field has been set.
func (o *BTMateConnectorCSInfo) HasGetzAxis() bool {
	if o != nil && o.GetzAxis != nil {
		return true
	}

	return false
}

// SetGetzAxis gets a reference to the given []float64 and assigns it to the GetzAxis field.
func (o *BTMateConnectorCSInfo) SetGetzAxis(v []float64) {
	o.GetzAxis = v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BTMateConnectorCSInfo) GetOrigin() []float64 {
	if o == nil || o.Origin == nil {
		var ret []float64
		return ret
	}
	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateConnectorCSInfo) GetOriginOk() ([]float64, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *BTMateConnectorCSInfo) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given []float64 and assigns it to the Origin field.
func (o *BTMateConnectorCSInfo) SetOrigin(v []float64) {
	o.Origin = v
}

func (o BTMateConnectorCSInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GetxAxis != nil {
		toSerialize["getxAxis"] = o.GetxAxis
	}
	if o.GetyAxis != nil {
		toSerialize["getyAxis"] = o.GetyAxis
	}
	if o.GetzAxis != nil {
		toSerialize["getzAxis"] = o.GetzAxis
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	return json.Marshal(toSerialize)
}

type NullableBTMateConnectorCSInfo struct {
	value *BTMateConnectorCSInfo
	isSet bool
}

func (v NullableBTMateConnectorCSInfo) Get() *BTMateConnectorCSInfo {
	return v.value
}

func (v *NullableBTMateConnectorCSInfo) Set(val *BTMateConnectorCSInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMateConnectorCSInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMateConnectorCSInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMateConnectorCSInfo(val *BTMateConnectorCSInfo) *NullableBTMateConnectorCSInfo {
	return &NullableBTMateConnectorCSInfo{value: val, isSet: true}
}

func (v NullableBTMateConnectorCSInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMateConnectorCSInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
