/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.164.16955-b4ecd192bba6
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMConfiguredValueByBoolean1501 struct for BTMConfiguredValueByBoolean1501
type BTMConfiguredValueByBoolean1501 struct {
	BtType                   *string        `json:"btType,omitempty"`
	ConfigurationValueString *string        `json:"configurationValueString,omitempty"`
	ImportMicroversion       *string        `json:"importMicroversion,omitempty"`
	NodeId                   *string        `json:"nodeId,omitempty"`
	Value                    *BTMParameter1 `json:"value,omitempty"`
	BooleanValue             *bool          `json:"booleanValue,omitempty"`
}

// NewBTMConfiguredValueByBoolean1501 instantiates a new BTMConfiguredValueByBoolean1501 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMConfiguredValueByBoolean1501() *BTMConfiguredValueByBoolean1501 {
	this := BTMConfiguredValueByBoolean1501{}
	return &this
}

// NewBTMConfiguredValueByBoolean1501WithDefaults instantiates a new BTMConfiguredValueByBoolean1501 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMConfiguredValueByBoolean1501WithDefaults() *BTMConfiguredValueByBoolean1501 {
	this := BTMConfiguredValueByBoolean1501{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMConfiguredValueByBoolean1501) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValueByBoolean1501) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMConfiguredValueByBoolean1501) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMConfiguredValueByBoolean1501) SetBtType(v string) {
	o.BtType = &v
}

// GetConfigurationValueString returns the ConfigurationValueString field value if set, zero value otherwise.
func (o *BTMConfiguredValueByBoolean1501) GetConfigurationValueString() string {
	if o == nil || o.ConfigurationValueString == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationValueString
}

// GetConfigurationValueStringOk returns a tuple with the ConfigurationValueString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValueByBoolean1501) GetConfigurationValueStringOk() (*string, bool) {
	if o == nil || o.ConfigurationValueString == nil {
		return nil, false
	}
	return o.ConfigurationValueString, true
}

// HasConfigurationValueString returns a boolean if a field has been set.
func (o *BTMConfiguredValueByBoolean1501) HasConfigurationValueString() bool {
	if o != nil && o.ConfigurationValueString != nil {
		return true
	}

	return false
}

// SetConfigurationValueString gets a reference to the given string and assigns it to the ConfigurationValueString field.
func (o *BTMConfiguredValueByBoolean1501) SetConfigurationValueString(v string) {
	o.ConfigurationValueString = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMConfiguredValueByBoolean1501) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValueByBoolean1501) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMConfiguredValueByBoolean1501) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMConfiguredValueByBoolean1501) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMConfiguredValueByBoolean1501) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValueByBoolean1501) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMConfiguredValueByBoolean1501) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMConfiguredValueByBoolean1501) SetNodeId(v string) {
	o.NodeId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTMConfiguredValueByBoolean1501) GetValue() BTMParameter1 {
	if o == nil || o.Value == nil {
		var ret BTMParameter1
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValueByBoolean1501) GetValueOk() (*BTMParameter1, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTMConfiguredValueByBoolean1501) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given BTMParameter1 and assigns it to the Value field.
func (o *BTMConfiguredValueByBoolean1501) SetValue(v BTMParameter1) {
	o.Value = &v
}

// GetBooleanValue returns the BooleanValue field value if set, zero value otherwise.
func (o *BTMConfiguredValueByBoolean1501) GetBooleanValue() bool {
	if o == nil || o.BooleanValue == nil {
		var ret bool
		return ret
	}
	return *o.BooleanValue
}

// GetBooleanValueOk returns a tuple with the BooleanValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValueByBoolean1501) GetBooleanValueOk() (*bool, bool) {
	if o == nil || o.BooleanValue == nil {
		return nil, false
	}
	return o.BooleanValue, true
}

// HasBooleanValue returns a boolean if a field has been set.
func (o *BTMConfiguredValueByBoolean1501) HasBooleanValue() bool {
	if o != nil && o.BooleanValue != nil {
		return true
	}

	return false
}

// SetBooleanValue gets a reference to the given bool and assigns it to the BooleanValue field.
func (o *BTMConfiguredValueByBoolean1501) SetBooleanValue(v bool) {
	o.BooleanValue = &v
}

func (o BTMConfiguredValueByBoolean1501) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ConfigurationValueString != nil {
		toSerialize["configurationValueString"] = o.ConfigurationValueString
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.BooleanValue != nil {
		toSerialize["booleanValue"] = o.BooleanValue
	}
	return json.Marshal(toSerialize)
}

type NullableBTMConfiguredValueByBoolean1501 struct {
	value *BTMConfiguredValueByBoolean1501
	isSet bool
}

func (v NullableBTMConfiguredValueByBoolean1501) Get() *BTMConfiguredValueByBoolean1501 {
	return v.value
}

func (v *NullableBTMConfiguredValueByBoolean1501) Set(val *BTMConfiguredValueByBoolean1501) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMConfiguredValueByBoolean1501) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMConfiguredValueByBoolean1501) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMConfiguredValueByBoolean1501(val *BTMConfiguredValueByBoolean1501) *NullableBTMConfiguredValueByBoolean1501 {
	return &NullableBTMConfiguredValueByBoolean1501{value: val, isSet: true}
}

func (v NullableBTMConfiguredValueByBoolean1501) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMConfiguredValueByBoolean1501) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
