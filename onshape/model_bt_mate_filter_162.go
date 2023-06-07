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
	"fmt"
)

// BTMateFilter162 - struct for BTMateFilter162
type BTMateFilter162 struct {
	implBTMateFilter162 interface{}
}

// BTAllowedMateTypeFilter1511AsBTMateFilter162 is a convenience function that returns BTAllowedMateTypeFilter1511 wrapped in BTMateFilter162
func (o *BTAllowedMateTypeFilter1511) AsBTMateFilter162() *BTMateFilter162 {
	return &BTMateFilter162{o}
}

// NewBTMateFilter162 instantiates a new BTMateFilter162 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMateFilter162() *BTMateFilter162 {
	this := BTMateFilter162{Newbase_BTMateFilter162()}
	return &this
}

// NewBTMateFilter162WithDefaults instantiates a new BTMateFilter162 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMateFilter162WithDefaults() *BTMateFilter162 {
	this := BTMateFilter162{Newbase_BTMateFilter162WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMateFilter162) GetBtType() string {
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
func (o *BTMateFilter162) GetBtTypeOk() (*string, bool) {
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
func (o *BTMateFilter162) HasBtType() bool {
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
func (o *BTMateFilter162) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetRequireMateQueryData returns the RequireMateQueryData field value if set, zero value otherwise.
func (o *BTMateFilter162) GetRequireMateQueryData() bool {
	type getResult interface {
		GetRequireMateQueryData() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRequireMateQueryData()
	} else {
		var de bool
		return de
	}
}

// GetRequireMateQueryDataOk returns a tuple with the RequireMateQueryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateFilter162) GetRequireMateQueryDataOk() (*bool, bool) {
	type getResult interface {
		GetRequireMateQueryDataOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRequireMateQueryDataOk()
	} else {
		return nil, false
	}
}

// HasRequireMateQueryData returns a boolean if a field has been set.
func (o *BTMateFilter162) HasRequireMateQueryData() bool {
	type getResult interface {
		HasRequireMateQueryData() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasRequireMateQueryData()
	} else {
		return false
	}
}

// SetRequireMateQueryData gets a reference to the given bool and assigns it to the RequireMateQueryData field.
func (o *BTMateFilter162) SetRequireMateQueryData(v bool) {
	type getResult interface {
		SetRequireMateQueryData(v bool)
	}

	o.GetActualInstance().(getResult).SetRequireMateQueryData(v)
}

// GetTopLevelMateOnly returns the TopLevelMateOnly field value if set, zero value otherwise.
func (o *BTMateFilter162) GetTopLevelMateOnly() bool {
	type getResult interface {
		GetTopLevelMateOnly() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTopLevelMateOnly()
	} else {
		var de bool
		return de
	}
}

// GetTopLevelMateOnlyOk returns a tuple with the TopLevelMateOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateFilter162) GetTopLevelMateOnlyOk() (*bool, bool) {
	type getResult interface {
		GetTopLevelMateOnlyOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTopLevelMateOnlyOk()
	} else {
		return nil, false
	}
}

// HasTopLevelMateOnly returns a boolean if a field has been set.
func (o *BTMateFilter162) HasTopLevelMateOnly() bool {
	type getResult interface {
		HasTopLevelMateOnly() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasTopLevelMateOnly()
	} else {
		return false
	}
}

// SetTopLevelMateOnly gets a reference to the given bool and assigns it to the TopLevelMateOnly field.
func (o *BTMateFilter162) SetTopLevelMateOnly(v bool) {
	type getResult interface {
		SetTopLevelMateOnly(v bool)
	}

	o.GetActualInstance().(getResult).SetTopLevelMateOnly(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTMateFilter162) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTAllowedMateTypeFilter-1511'
	if jsonDict["btType"] == "BTAllowedMateTypeFilter-1511" {
		// try to unmarshal JSON data into BTAllowedMateTypeFilter1511
		var qr *BTAllowedMateTypeFilter1511
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMateFilter162 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMateFilter162 = nil
			return fmt.Errorf("Failed to unmarshal BTMateFilter162 as BTAllowedMateTypeFilter1511: %s", err.Error())
		}
	}

	var qtx *base_BTMateFilter162
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTMateFilter162 = qtx
		return nil // data stored in dst.base_BTMateFilter162, return on the first match
	} else {
		dst.implBTMateFilter162 = nil
		return fmt.Errorf("Failed to unmarshal BTMateFilter162 as base_BTMateFilter162: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTMateFilter162) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTMateFilter162) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTMateFilter162
}

type NullableBTMateFilter162 struct {
	value *BTMateFilter162
	isSet bool
}

func (v NullableBTMateFilter162) Get() *BTMateFilter162 {
	return v.value
}

func (v *NullableBTMateFilter162) Set(val *BTMateFilter162) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMateFilter162) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMateFilter162) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMateFilter162(val *BTMateFilter162) *NullableBTMateFilter162 {
	return &NullableBTMateFilter162{value: val, isSet: true}
}

func (v NullableBTMateFilter162) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMateFilter162) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTMateFilter162 struct {
	BtType               *string `json:"btType,omitempty"`
	RequireMateQueryData *bool   `json:"requireMateQueryData,omitempty"`
	TopLevelMateOnly     *bool   `json:"topLevelMateOnly,omitempty"`
}

// Newbase_BTMateFilter162 instantiates a new base_BTMateFilter162 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTMateFilter162() *base_BTMateFilter162 {
	this := base_BTMateFilter162{}
	return &this
}

// Newbase_BTMateFilter162WithDefaults instantiates a new base_BTMateFilter162 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTMateFilter162WithDefaults() *base_BTMateFilter162 {
	this := base_BTMateFilter162{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTMateFilter162) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMateFilter162) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTMateFilter162) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTMateFilter162) SetBtType(v string) {
	o.BtType = &v
}

// GetRequireMateQueryData returns the RequireMateQueryData field value if set, zero value otherwise.
func (o *base_BTMateFilter162) GetRequireMateQueryData() bool {
	if o == nil || o.RequireMateQueryData == nil {
		var ret bool
		return ret
	}
	return *o.RequireMateQueryData
}

// GetRequireMateQueryDataOk returns a tuple with the RequireMateQueryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMateFilter162) GetRequireMateQueryDataOk() (*bool, bool) {
	if o == nil || o.RequireMateQueryData == nil {
		return nil, false
	}
	return o.RequireMateQueryData, true
}

// HasRequireMateQueryData returns a boolean if a field has been set.
func (o *base_BTMateFilter162) HasRequireMateQueryData() bool {
	if o != nil && o.RequireMateQueryData != nil {
		return true
	}

	return false
}

// SetRequireMateQueryData gets a reference to the given bool and assigns it to the RequireMateQueryData field.
func (o *base_BTMateFilter162) SetRequireMateQueryData(v bool) {
	o.RequireMateQueryData = &v
}

// GetTopLevelMateOnly returns the TopLevelMateOnly field value if set, zero value otherwise.
func (o *base_BTMateFilter162) GetTopLevelMateOnly() bool {
	if o == nil || o.TopLevelMateOnly == nil {
		var ret bool
		return ret
	}
	return *o.TopLevelMateOnly
}

// GetTopLevelMateOnlyOk returns a tuple with the TopLevelMateOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMateFilter162) GetTopLevelMateOnlyOk() (*bool, bool) {
	if o == nil || o.TopLevelMateOnly == nil {
		return nil, false
	}
	return o.TopLevelMateOnly, true
}

// HasTopLevelMateOnly returns a boolean if a field has been set.
func (o *base_BTMateFilter162) HasTopLevelMateOnly() bool {
	if o != nil && o.TopLevelMateOnly != nil {
		return true
	}

	return false
}

// SetTopLevelMateOnly gets a reference to the given bool and assigns it to the TopLevelMateOnly field.
func (o *base_BTMateFilter162) SetTopLevelMateOnly(v bool) {
	o.TopLevelMateOnly = &v
}

func (o base_BTMateFilter162) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.RequireMateQueryData != nil {
		toSerialize["requireMateQueryData"] = o.RequireMateQueryData
	}
	if o.TopLevelMateOnly != nil {
		toSerialize["topLevelMateOnly"] = o.TopLevelMateOnly
	}
	return json.Marshal(toSerialize)
}
