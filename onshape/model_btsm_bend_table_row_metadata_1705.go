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

// BTSMBendTableRowMetadata1705 struct for BTSMBendTableRowMetadata1705
type BTSMBendTableRowMetadata1705 struct {
	BtType                  *string                        `json:"btType,omitempty"`
	CrossHighlightDataIfAny *BTTableCrossHighlightData1753 `json:"crossHighlightDataIfAny,omitempty"`
	CrossHighlightData      *BTTableCrossHighlightData1753 `json:"crossHighlightData,omitempty"`
	IsJointTypeEditable     *bool                          `json:"isJointTypeEditable,omitempty"`
}

// NewBTSMBendTableRowMetadata1705 instantiates a new BTSMBendTableRowMetadata1705 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSMBendTableRowMetadata1705() *BTSMBendTableRowMetadata1705 {
	this := BTSMBendTableRowMetadata1705{}
	return &this
}

// NewBTSMBendTableRowMetadata1705WithDefaults instantiates a new BTSMBendTableRowMetadata1705 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSMBendTableRowMetadata1705WithDefaults() *BTSMBendTableRowMetadata1705 {
	this := BTSMBendTableRowMetadata1705{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSMBendTableRowMetadata1705) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSMBendTableRowMetadata1705) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSMBendTableRowMetadata1705) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSMBendTableRowMetadata1705) SetBtType(v string) {
	o.BtType = &v
}

// GetCrossHighlightDataIfAny returns the CrossHighlightDataIfAny field value if set, zero value otherwise.
func (o *BTSMBendTableRowMetadata1705) GetCrossHighlightDataIfAny() BTTableCrossHighlightData1753 {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		var ret BTTableCrossHighlightData1753
		return ret
	}
	return *o.CrossHighlightDataIfAny
}

// GetCrossHighlightDataIfAnyOk returns a tuple with the CrossHighlightDataIfAny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSMBendTableRowMetadata1705) GetCrossHighlightDataIfAnyOk() (*BTTableCrossHighlightData1753, bool) {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		return nil, false
	}
	return o.CrossHighlightDataIfAny, true
}

// HasCrossHighlightDataIfAny returns a boolean if a field has been set.
func (o *BTSMBendTableRowMetadata1705) HasCrossHighlightDataIfAny() bool {
	if o != nil && o.CrossHighlightDataIfAny != nil {
		return true
	}

	return false
}

// SetCrossHighlightDataIfAny gets a reference to the given BTTableCrossHighlightData1753 and assigns it to the CrossHighlightDataIfAny field.
func (o *BTSMBendTableRowMetadata1705) SetCrossHighlightDataIfAny(v BTTableCrossHighlightData1753) {
	o.CrossHighlightDataIfAny = &v
}

// GetCrossHighlightData returns the CrossHighlightData field value if set, zero value otherwise.
func (o *BTSMBendTableRowMetadata1705) GetCrossHighlightData() BTTableCrossHighlightData1753 {
	if o == nil || o.CrossHighlightData == nil {
		var ret BTTableCrossHighlightData1753
		return ret
	}
	return *o.CrossHighlightData
}

// GetCrossHighlightDataOk returns a tuple with the CrossHighlightData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSMBendTableRowMetadata1705) GetCrossHighlightDataOk() (*BTTableCrossHighlightData1753, bool) {
	if o == nil || o.CrossHighlightData == nil {
		return nil, false
	}
	return o.CrossHighlightData, true
}

// HasCrossHighlightData returns a boolean if a field has been set.
func (o *BTSMBendTableRowMetadata1705) HasCrossHighlightData() bool {
	if o != nil && o.CrossHighlightData != nil {
		return true
	}

	return false
}

// SetCrossHighlightData gets a reference to the given BTTableCrossHighlightData1753 and assigns it to the CrossHighlightData field.
func (o *BTSMBendTableRowMetadata1705) SetCrossHighlightData(v BTTableCrossHighlightData1753) {
	o.CrossHighlightData = &v
}

// GetIsJointTypeEditable returns the IsJointTypeEditable field value if set, zero value otherwise.
func (o *BTSMBendTableRowMetadata1705) GetIsJointTypeEditable() bool {
	if o == nil || o.IsJointTypeEditable == nil {
		var ret bool
		return ret
	}
	return *o.IsJointTypeEditable
}

// GetIsJointTypeEditableOk returns a tuple with the IsJointTypeEditable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSMBendTableRowMetadata1705) GetIsJointTypeEditableOk() (*bool, bool) {
	if o == nil || o.IsJointTypeEditable == nil {
		return nil, false
	}
	return o.IsJointTypeEditable, true
}

// HasIsJointTypeEditable returns a boolean if a field has been set.
func (o *BTSMBendTableRowMetadata1705) HasIsJointTypeEditable() bool {
	if o != nil && o.IsJointTypeEditable != nil {
		return true
	}

	return false
}

// SetIsJointTypeEditable gets a reference to the given bool and assigns it to the IsJointTypeEditable field.
func (o *BTSMBendTableRowMetadata1705) SetIsJointTypeEditable(v bool) {
	o.IsJointTypeEditable = &v
}

func (o BTSMBendTableRowMetadata1705) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CrossHighlightDataIfAny != nil {
		toSerialize["crossHighlightDataIfAny"] = o.CrossHighlightDataIfAny
	}
	if o.CrossHighlightData != nil {
		toSerialize["crossHighlightData"] = o.CrossHighlightData
	}
	if o.IsJointTypeEditable != nil {
		toSerialize["isJointTypeEditable"] = o.IsJointTypeEditable
	}
	return json.Marshal(toSerialize)
}

type NullableBTSMBendTableRowMetadata1705 struct {
	value *BTSMBendTableRowMetadata1705
	isSet bool
}

func (v NullableBTSMBendTableRowMetadata1705) Get() *BTSMBendTableRowMetadata1705 {
	return v.value
}

func (v *NullableBTSMBendTableRowMetadata1705) Set(val *BTSMBendTableRowMetadata1705) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSMBendTableRowMetadata1705) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSMBendTableRowMetadata1705) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSMBendTableRowMetadata1705(val *BTSMBendTableRowMetadata1705) *NullableBTSMBendTableRowMetadata1705 {
	return &NullableBTSMBendTableRowMetadata1705{value: val, isSet: true}
}

func (v NullableBTSMBendTableRowMetadata1705) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSMBendTableRowMetadata1705) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
