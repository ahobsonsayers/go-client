/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6820-1bef41f9cc03
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTFSValueWithUnits1817 struct for BTFSValueWithUnits1817
type BTFSValueWithUnits1817 struct {
	BtType      string            `json:"btType"`
	TypeTag     *string           `json:"typeTag,omitempty"`
	UnitToPower *map[string]int32 `json:"unitToPower,omitempty"`
	Value       *float64          `json:"value,omitempty"`
}

// NewBTFSValueWithUnits1817 instantiates a new BTFSValueWithUnits1817 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSValueWithUnits1817(btType string) *BTFSValueWithUnits1817 {
	this := BTFSValueWithUnits1817{}
	this.BtType = btType
	return &this
}

// NewBTFSValueWithUnits1817WithDefaults instantiates a new BTFSValueWithUnits1817 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSValueWithUnits1817WithDefaults() *BTFSValueWithUnits1817 {
	this := BTFSValueWithUnits1817{}
	return &this
}

// GetBtType returns the BtType field value
func (o *BTFSValueWithUnits1817) GetBtType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value
// and a boolean to check if the value has been set.
func (o *BTFSValueWithUnits1817) GetBtTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BtType, true
}

// SetBtType sets field value
func (o *BTFSValueWithUnits1817) SetBtType(v string) {
	o.BtType = v
}

// GetTypeTag returns the TypeTag field value if set, zero value otherwise.
func (o *BTFSValueWithUnits1817) GetTypeTag() string {
	if o == nil || o.TypeTag == nil {
		var ret string
		return ret
	}
	return *o.TypeTag
}

// GetTypeTagOk returns a tuple with the TypeTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueWithUnits1817) GetTypeTagOk() (*string, bool) {
	if o == nil || o.TypeTag == nil {
		return nil, false
	}
	return o.TypeTag, true
}

// HasTypeTag returns a boolean if a field has been set.
func (o *BTFSValueWithUnits1817) HasTypeTag() bool {
	if o != nil && o.TypeTag != nil {
		return true
	}

	return false
}

// SetTypeTag gets a reference to the given string and assigns it to the TypeTag field.
func (o *BTFSValueWithUnits1817) SetTypeTag(v string) {
	o.TypeTag = &v
}

// GetUnitToPower returns the UnitToPower field value if set, zero value otherwise.
func (o *BTFSValueWithUnits1817) GetUnitToPower() map[string]int32 {
	if o == nil || o.UnitToPower == nil {
		var ret map[string]int32
		return ret
	}
	return *o.UnitToPower
}

// GetUnitToPowerOk returns a tuple with the UnitToPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueWithUnits1817) GetUnitToPowerOk() (*map[string]int32, bool) {
	if o == nil || o.UnitToPower == nil {
		return nil, false
	}
	return o.UnitToPower, true
}

// HasUnitToPower returns a boolean if a field has been set.
func (o *BTFSValueWithUnits1817) HasUnitToPower() bool {
	if o != nil && o.UnitToPower != nil {
		return true
	}

	return false
}

// SetUnitToPower gets a reference to the given map[string]int32 and assigns it to the UnitToPower field.
func (o *BTFSValueWithUnits1817) SetUnitToPower(v map[string]int32) {
	o.UnitToPower = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTFSValueWithUnits1817) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueWithUnits1817) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTFSValueWithUnits1817) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *BTFSValueWithUnits1817) SetValue(v float64) {
	o.Value = &v
}

func (o BTFSValueWithUnits1817) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["btType"] = o.BtType
	}
	if o.TypeTag != nil {
		toSerialize["typeTag"] = o.TypeTag
	}
	if o.UnitToPower != nil {
		toSerialize["unitToPower"] = o.UnitToPower
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTFSValueWithUnits1817 struct {
	value *BTFSValueWithUnits1817
	isSet bool
}

func (v NullableBTFSValueWithUnits1817) Get() *BTFSValueWithUnits1817 {
	return v.value
}

func (v *NullableBTFSValueWithUnits1817) Set(val *BTFSValueWithUnits1817) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSValueWithUnits1817) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSValueWithUnits1817) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSValueWithUnits1817(val *BTFSValueWithUnits1817) *NullableBTFSValueWithUnits1817 {
	return &NullableBTFSValueWithUnits1817{value: val, isSet: true}
}

func (v NullableBTFSValueWithUnits1817) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSValueWithUnits1817) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
