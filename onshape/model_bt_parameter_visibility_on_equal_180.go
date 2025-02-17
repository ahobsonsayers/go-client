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
	"fmt"
)

// BTParameterVisibilityOnEqual180 - struct for BTParameterVisibilityOnEqual180
type BTParameterVisibilityOnEqual180 struct {
	implBTParameterVisibilityOnEqual180 interface{}
}

// BTParameterVisibilityOnMateDOFType2114AsBTParameterVisibilityOnEqual180 is a convenience function that returns BTParameterVisibilityOnMateDOFType2114 wrapped in BTParameterVisibilityOnEqual180
func (o *BTParameterVisibilityOnMateDOFType2114) AsBTParameterVisibilityOnEqual180() *BTParameterVisibilityOnEqual180 {
	return &BTParameterVisibilityOnEqual180{o}
}

// NewBTParameterVisibilityOnEqual180 instantiates a new BTParameterVisibilityOnEqual180 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterVisibilityOnEqual180() *BTParameterVisibilityOnEqual180 {
	this := BTParameterVisibilityOnEqual180{Newbase_BTParameterVisibilityOnEqual180()}
	return &this
}

// NewBTParameterVisibilityOnEqual180WithDefaults instantiates a new BTParameterVisibilityOnEqual180 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterVisibilityOnEqual180WithDefaults() *BTParameterVisibilityOnEqual180 {
	this := BTParameterVisibilityOnEqual180{Newbase_BTParameterVisibilityOnEqual180WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterVisibilityOnEqual180) GetBtType() string {
	type getResult interface {
		GetBtType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtType()
	} else {
		var de string
		return de
	}
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterVisibilityOnEqual180) GetBtTypeOk() (*string, bool) {
	type getResult interface {
		GetBtTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtTypeOk()
	} else {
		return nil, false
	}
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterVisibilityOnEqual180) HasBtType() bool {
	type getResult interface {
		HasBtType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasBtType()
	} else {
		return false
	}
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterVisibilityOnEqual180) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetInArray returns the InArray field value if set, zero value otherwise.
func (o *BTParameterVisibilityOnEqual180) GetInArray() bool {
	type getResult interface {
		GetInArray() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetInArray()
	} else {
		var de bool
		return de
	}
}

// GetInArrayOk returns a tuple with the InArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterVisibilityOnEqual180) GetInArrayOk() (*bool, bool) {
	type getResult interface {
		GetInArrayOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetInArrayOk()
	} else {
		return nil, false
	}
}

// HasInArray returns a boolean if a field has been set.
func (o *BTParameterVisibilityOnEqual180) HasInArray() bool {
	type getResult interface {
		HasInArray() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasInArray()
	} else {
		return false
	}
}

// SetInArray gets a reference to the given bool and assigns it to the InArray field.
func (o *BTParameterVisibilityOnEqual180) SetInArray(v bool) {
	type getResult interface {
		SetInArray(v bool)
	}

	o.GetActualInstance().(getResult).SetInArray(v)
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTParameterVisibilityOnEqual180) GetParameterId() string {
	type getResult interface {
		GetParameterId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParameterId()
	} else {
		var de string
		return de
	}
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterVisibilityOnEqual180) GetParameterIdOk() (*string, bool) {
	type getResult interface {
		GetParameterIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParameterIdOk()
	} else {
		return nil, false
	}
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTParameterVisibilityOnEqual180) HasParameterId() bool {
	type getResult interface {
		HasParameterId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasParameterId()
	} else {
		return false
	}
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTParameterVisibilityOnEqual180) SetParameterId(v string) {
	type getResult interface {
		SetParameterId(v string)
	}

	o.GetActualInstance().(getResult).SetParameterId(v)
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTParameterVisibilityOnEqual180) GetValue() string {
	type getResult interface {
		GetValue() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetValue()
	} else {
		var de string
		return de
	}
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterVisibilityOnEqual180) GetValueOk() (*string, bool) {
	type getResult interface {
		GetValueOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetValueOk()
	} else {
		return nil, false
	}
}

// HasValue returns a boolean if a field has been set.
func (o *BTParameterVisibilityOnEqual180) HasValue() bool {
	type getResult interface {
		HasValue() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasValue()
	} else {
		return false
	}
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *BTParameterVisibilityOnEqual180) SetValue(v string) {
	type getResult interface {
		SetValue(v string)
	}

	o.GetActualInstance().(getResult).SetValue(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTParameterVisibilityOnEqual180) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTParameterVisibilityOnMateDOFType-2114'
	if jsonDict["btType"] == "BTParameterVisibilityOnMateDOFType-2114" {
		// try to unmarshal JSON data into BTParameterVisibilityOnMateDOFType2114
		var qr *BTParameterVisibilityOnMateDOFType2114
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTParameterVisibilityOnEqual180 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTParameterVisibilityOnEqual180 = nil
			return fmt.Errorf("Failed to unmarshal BTParameterVisibilityOnEqual180 as BTParameterVisibilityOnMateDOFType2114: %s", err.Error())
		}
	}

	var qtx *base_BTParameterVisibilityOnEqual180
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTParameterVisibilityOnEqual180 = qtx
		return nil // data stored in dst.base_BTParameterVisibilityOnEqual180, return on the first match
	} else {
		dst.implBTParameterVisibilityOnEqual180 = nil
		return fmt.Errorf("Failed to unmarshal BTParameterVisibilityOnEqual180 as base_BTParameterVisibilityOnEqual180: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTParameterVisibilityOnEqual180) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTParameterVisibilityOnEqual180) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTParameterVisibilityOnEqual180
}

type NullableBTParameterVisibilityOnEqual180 struct {
	value *BTParameterVisibilityOnEqual180
	isSet bool
}

func (v NullableBTParameterVisibilityOnEqual180) Get() *BTParameterVisibilityOnEqual180 {
	return v.value
}

func (v *NullableBTParameterVisibilityOnEqual180) Set(val *BTParameterVisibilityOnEqual180) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterVisibilityOnEqual180) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterVisibilityOnEqual180) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterVisibilityOnEqual180(val *BTParameterVisibilityOnEqual180) *NullableBTParameterVisibilityOnEqual180 {
	return &NullableBTParameterVisibilityOnEqual180{value: val, isSet: true}
}

func (v NullableBTParameterVisibilityOnEqual180) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterVisibilityOnEqual180) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTParameterVisibilityOnEqual180 struct {
	BtType      *string `json:"btType,omitempty"`
	InArray     *bool   `json:"inArray,omitempty"`
	ParameterId *string `json:"parameterId,omitempty"`
	Value       *string `json:"value,omitempty"`
}

// Newbase_BTParameterVisibilityOnEqual180 instantiates a new base_BTParameterVisibilityOnEqual180 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTParameterVisibilityOnEqual180() *base_BTParameterVisibilityOnEqual180 {
	this := base_BTParameterVisibilityOnEqual180{}
	return &this
}

// Newbase_BTParameterVisibilityOnEqual180WithDefaults instantiates a new base_BTParameterVisibilityOnEqual180 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTParameterVisibilityOnEqual180WithDefaults() *base_BTParameterVisibilityOnEqual180 {
	this := base_BTParameterVisibilityOnEqual180{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTParameterVisibilityOnEqual180) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTParameterVisibilityOnEqual180) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTParameterVisibilityOnEqual180) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTParameterVisibilityOnEqual180) SetBtType(v string) {
	o.BtType = &v
}

// GetInArray returns the InArray field value if set, zero value otherwise.
func (o *base_BTParameterVisibilityOnEqual180) GetInArray() bool {
	if o == nil || o.InArray == nil {
		var ret bool
		return ret
	}
	return *o.InArray
}

// GetInArrayOk returns a tuple with the InArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTParameterVisibilityOnEqual180) GetInArrayOk() (*bool, bool) {
	if o == nil || o.InArray == nil {
		return nil, false
	}
	return o.InArray, true
}

// HasInArray returns a boolean if a field has been set.
func (o *base_BTParameterVisibilityOnEqual180) HasInArray() bool {
	if o != nil && o.InArray != nil {
		return true
	}

	return false
}

// SetInArray gets a reference to the given bool and assigns it to the InArray field.
func (o *base_BTParameterVisibilityOnEqual180) SetInArray(v bool) {
	o.InArray = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *base_BTParameterVisibilityOnEqual180) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTParameterVisibilityOnEqual180) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *base_BTParameterVisibilityOnEqual180) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *base_BTParameterVisibilityOnEqual180) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *base_BTParameterVisibilityOnEqual180) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTParameterVisibilityOnEqual180) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *base_BTParameterVisibilityOnEqual180) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *base_BTParameterVisibilityOnEqual180) SetValue(v string) {
	o.Value = &v
}

func (o base_BTParameterVisibilityOnEqual180) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.InArray != nil {
		toSerialize["inArray"] = o.InArray
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}
