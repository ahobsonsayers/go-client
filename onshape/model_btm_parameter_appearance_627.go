/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.6979-8ce9514d51cf
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMParameterAppearance627 struct for BTMParameterAppearance627
type BTMParameterAppearance627 struct {
	ImportMicroversion *string                   `json:"importMicroversion,omitempty"`
	NodeId             *string                   `json:"nodeId,omitempty"`
	ParameterId        *string                   `json:"parameterId,omitempty"`
	Appearance         *BTGraphicsAppearance1152 `json:"appearance,omitempty"`
	BtType             *string                   `json:"btType,omitempty"`
}

// NewBTMParameterAppearance627 instantiates a new BTMParameterAppearance627 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterAppearance627() *BTMParameterAppearance627 {
	this := BTMParameterAppearance627{}
	return &this
}

// NewBTMParameterAppearance627WithDefaults instantiates a new BTMParameterAppearance627 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterAppearance627WithDefaults() *BTMParameterAppearance627 {
	this := BTMParameterAppearance627{}
	return &this
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMParameterAppearance627) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterAppearance627) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMParameterAppearance627) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMParameterAppearance627) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMParameterAppearance627) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterAppearance627) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMParameterAppearance627) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMParameterAppearance627) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMParameterAppearance627) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterAppearance627) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMParameterAppearance627) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMParameterAppearance627) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetAppearance returns the Appearance field value if set, zero value otherwise.
func (o *BTMParameterAppearance627) GetAppearance() BTGraphicsAppearance1152 {
	if o == nil || o.Appearance == nil {
		var ret BTGraphicsAppearance1152
		return ret
	}
	return *o.Appearance
}

// GetAppearanceOk returns a tuple with the Appearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterAppearance627) GetAppearanceOk() (*BTGraphicsAppearance1152, bool) {
	if o == nil || o.Appearance == nil {
		return nil, false
	}
	return o.Appearance, true
}

// HasAppearance returns a boolean if a field has been set.
func (o *BTMParameterAppearance627) HasAppearance() bool {
	if o != nil && o.Appearance != nil {
		return true
	}

	return false
}

// SetAppearance gets a reference to the given BTGraphicsAppearance1152 and assigns it to the Appearance field.
func (o *BTMParameterAppearance627) SetAppearance(v BTGraphicsAppearance1152) {
	o.Appearance = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterAppearance627) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterAppearance627) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterAppearance627) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterAppearance627) SetBtType(v string) {
	o.BtType = &v
}

func (o BTMParameterAppearance627) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.Appearance != nil {
		toSerialize["appearance"] = o.Appearance
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameterAppearance627 struct {
	value *BTMParameterAppearance627
	isSet bool
}

func (v NullableBTMParameterAppearance627) Get() *BTMParameterAppearance627 {
	return v.value
}

func (v *NullableBTMParameterAppearance627) Set(val *BTMParameterAppearance627) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterAppearance627) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterAppearance627) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterAppearance627(val *BTMParameterAppearance627) *NullableBTMParameterAppearance627 {
	return &NullableBTMParameterAppearance627{value: val, isSet: true}
}

func (v NullableBTMParameterAppearance627) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterAppearance627) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
