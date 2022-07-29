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

// BTSMOtherJointTableRowMetadata2640 struct for BTSMOtherJointTableRowMetadata2640
type BTSMOtherJointTableRowMetadata2640 struct {
	CrossHighlightDataIfAny *BTTableCrossHighlightData1753 `json:"crossHighlightDataIfAny,omitempty"`
	BtType                  *string                        `json:"btType,omitempty"`
	CrossHighlightData      *BTTableCrossHighlightData1753 `json:"crossHighlightData,omitempty"`
}

// NewBTSMOtherJointTableRowMetadata2640 instantiates a new BTSMOtherJointTableRowMetadata2640 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSMOtherJointTableRowMetadata2640() *BTSMOtherJointTableRowMetadata2640 {
	this := BTSMOtherJointTableRowMetadata2640{}
	return &this
}

// NewBTSMOtherJointTableRowMetadata2640WithDefaults instantiates a new BTSMOtherJointTableRowMetadata2640 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSMOtherJointTableRowMetadata2640WithDefaults() *BTSMOtherJointTableRowMetadata2640 {
	this := BTSMOtherJointTableRowMetadata2640{}
	return &this
}

// GetCrossHighlightDataIfAny returns the CrossHighlightDataIfAny field value if set, zero value otherwise.
func (o *BTSMOtherJointTableRowMetadata2640) GetCrossHighlightDataIfAny() BTTableCrossHighlightData1753 {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		var ret BTTableCrossHighlightData1753
		return ret
	}
	return *o.CrossHighlightDataIfAny
}

// GetCrossHighlightDataIfAnyOk returns a tuple with the CrossHighlightDataIfAny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSMOtherJointTableRowMetadata2640) GetCrossHighlightDataIfAnyOk() (*BTTableCrossHighlightData1753, bool) {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		return nil, false
	}
	return o.CrossHighlightDataIfAny, true
}

// HasCrossHighlightDataIfAny returns a boolean if a field has been set.
func (o *BTSMOtherJointTableRowMetadata2640) HasCrossHighlightDataIfAny() bool {
	if o != nil && o.CrossHighlightDataIfAny != nil {
		return true
	}

	return false
}

// SetCrossHighlightDataIfAny gets a reference to the given BTTableCrossHighlightData1753 and assigns it to the CrossHighlightDataIfAny field.
func (o *BTSMOtherJointTableRowMetadata2640) SetCrossHighlightDataIfAny(v BTTableCrossHighlightData1753) {
	o.CrossHighlightDataIfAny = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSMOtherJointTableRowMetadata2640) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSMOtherJointTableRowMetadata2640) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSMOtherJointTableRowMetadata2640) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSMOtherJointTableRowMetadata2640) SetBtType(v string) {
	o.BtType = &v
}

// GetCrossHighlightData returns the CrossHighlightData field value if set, zero value otherwise.
func (o *BTSMOtherJointTableRowMetadata2640) GetCrossHighlightData() BTTableCrossHighlightData1753 {
	if o == nil || o.CrossHighlightData == nil {
		var ret BTTableCrossHighlightData1753
		return ret
	}
	return *o.CrossHighlightData
}

// GetCrossHighlightDataOk returns a tuple with the CrossHighlightData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSMOtherJointTableRowMetadata2640) GetCrossHighlightDataOk() (*BTTableCrossHighlightData1753, bool) {
	if o == nil || o.CrossHighlightData == nil {
		return nil, false
	}
	return o.CrossHighlightData, true
}

// HasCrossHighlightData returns a boolean if a field has been set.
func (o *BTSMOtherJointTableRowMetadata2640) HasCrossHighlightData() bool {
	if o != nil && o.CrossHighlightData != nil {
		return true
	}

	return false
}

// SetCrossHighlightData gets a reference to the given BTTableCrossHighlightData1753 and assigns it to the CrossHighlightData field.
func (o *BTSMOtherJointTableRowMetadata2640) SetCrossHighlightData(v BTTableCrossHighlightData1753) {
	o.CrossHighlightData = &v
}

func (o BTSMOtherJointTableRowMetadata2640) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CrossHighlightDataIfAny != nil {
		toSerialize["crossHighlightDataIfAny"] = o.CrossHighlightDataIfAny
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CrossHighlightData != nil {
		toSerialize["crossHighlightData"] = o.CrossHighlightData
	}
	return json.Marshal(toSerialize)
}

type NullableBTSMOtherJointTableRowMetadata2640 struct {
	value *BTSMOtherJointTableRowMetadata2640
	isSet bool
}

func (v NullableBTSMOtherJointTableRowMetadata2640) Get() *BTSMOtherJointTableRowMetadata2640 {
	return v.value
}

func (v *NullableBTSMOtherJointTableRowMetadata2640) Set(val *BTSMOtherJointTableRowMetadata2640) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSMOtherJointTableRowMetadata2640) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSMOtherJointTableRowMetadata2640) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSMOtherJointTableRowMetadata2640(val *BTSMOtherJointTableRowMetadata2640) *NullableBTSMOtherJointTableRowMetadata2640 {
	return &NullableBTSMOtherJointTableRowMetadata2640{value: val, isSet: true}
}

func (v NullableBTSMOtherJointTableRowMetadata2640) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSMOtherJointTableRowMetadata2640) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
