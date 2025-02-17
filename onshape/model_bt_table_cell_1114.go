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

// BTTableCell1114 - struct for BTTableCell1114
type BTTableCell1114 struct {
	implBTTableCell1114 interface{}
}

// BTTableCellParameterWithValue2122AsBTTableCell1114 is a convenience function that returns BTTableCellParameterWithValue2122 wrapped in BTTableCell1114
func (o *BTTableCellParameterWithValue2122) AsBTTableCell1114() *BTTableCell1114 {
	return &BTTableCell1114{o}
}

// BTTableCellPropertyParameter2983AsBTTableCell1114 is a convenience function that returns BTTableCellPropertyParameter2983 wrapped in BTTableCell1114
func (o *BTTableCellPropertyParameter2983) AsBTTableCell1114() *BTTableCell1114 {
	return &BTTableCell1114{o}
}

// BTTableTestCellString2112AsBTTableCell1114 is a convenience function that returns BTTableTestCellString2112 wrapped in BTTableCell1114
func (o *BTTableTestCellString2112) AsBTTableCell1114() *BTTableCell1114 {
	return &BTTableCell1114{o}
}

// BTTableCellParameter2399AsBTTableCell1114 is a convenience function that returns BTTableCellParameter2399 wrapped in BTTableCell1114
func (o *BTTableCellParameter2399) AsBTTableCell1114() *BTTableCell1114 {
	return &BTTableCell1114{o}
}

// BTTableTestCellDouble2509AsBTTableCell1114 is a convenience function that returns BTTableTestCellDouble2509 wrapped in BTTableCell1114
func (o *BTTableTestCellDouble2509) AsBTTableCell1114() *BTTableCell1114 {
	return &BTTableCell1114{o}
}

// NewBTTableCell1114 instantiates a new BTTableCell1114 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableCell1114() *BTTableCell1114 {
	this := BTTableCell1114{Newbase_BTTableCell1114()}
	return &this
}

// NewBTTableCell1114WithDefaults instantiates a new BTTableCell1114 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableCell1114WithDefaults() *BTTableCell1114 {
	this := BTTableCell1114{Newbase_BTTableCell1114WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTableCell1114) GetBtType() string {
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
func (o *BTTableCell1114) GetBtTypeOk() (*string, bool) {
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
func (o *BTTableCell1114) HasBtType() bool {
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
func (o *BTTableCell1114) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetIsEverVisible returns the IsEverVisible field value if set, zero value otherwise.
func (o *BTTableCell1114) GetIsEverVisible() bool {
	type getResult interface {
		GetIsEverVisible() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIsEverVisible()
	} else {
		var de bool
		return de
	}
}

// GetIsEverVisibleOk returns a tuple with the IsEverVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCell1114) GetIsEverVisibleOk() (*bool, bool) {
	type getResult interface {
		GetIsEverVisibleOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIsEverVisibleOk()
	} else {
		return nil, false
	}
}

// HasIsEverVisible returns a boolean if a field has been set.
func (o *BTTableCell1114) HasIsEverVisible() bool {
	type getResult interface {
		HasIsEverVisible() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasIsEverVisible()
	} else {
		return false
	}
}

// SetIsEverVisible gets a reference to the given bool and assigns it to the IsEverVisible field.
func (o *BTTableCell1114) SetIsEverVisible(v bool) {
	type getResult interface {
		SetIsEverVisible(v bool)
	}

	o.GetActualInstance().(getResult).SetIsEverVisible(v)
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *BTTableCell1114) GetIsReadOnly() bool {
	type getResult interface {
		GetIsReadOnly() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIsReadOnly()
	} else {
		var de bool
		return de
	}
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCell1114) GetIsReadOnlyOk() (*bool, bool) {
	type getResult interface {
		GetIsReadOnlyOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIsReadOnlyOk()
	} else {
		return nil, false
	}
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *BTTableCell1114) HasIsReadOnly() bool {
	type getResult interface {
		HasIsReadOnly() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasIsReadOnly()
	} else {
		return false
	}
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *BTTableCell1114) SetIsReadOnly(v bool) {
	type getResult interface {
		SetIsReadOnly(v bool)
	}

	o.GetActualInstance().(getResult).SetIsReadOnly(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTTableCell1114) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTTableCellParameter-2399'
	if jsonDict["btType"] == "BTTableCellParameter-2399" {
		// try to unmarshal JSON data into BTTableCellParameter2399
		var qr *BTTableCellParameter2399
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableCell1114 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableCell1114 = nil
			return fmt.Errorf("Failed to unmarshal BTTableCell1114 as BTTableCellParameter2399: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTTableCellParameterWithValue-2122'
	if jsonDict["btType"] == "BTTableCellParameterWithValue-2122" {
		// try to unmarshal JSON data into BTTableCellParameterWithValue2122
		var qr *BTTableCellParameterWithValue2122
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableCell1114 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableCell1114 = nil
			return fmt.Errorf("Failed to unmarshal BTTableCell1114 as BTTableCellParameterWithValue2122: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTTableCellPropertyParameter-2983'
	if jsonDict["btType"] == "BTTableCellPropertyParameter-2983" {
		// try to unmarshal JSON data into BTTableCellPropertyParameter2983
		var qr *BTTableCellPropertyParameter2983
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableCell1114 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableCell1114 = nil
			return fmt.Errorf("Failed to unmarshal BTTableCell1114 as BTTableCellPropertyParameter2983: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTTableTestCellDouble-2509'
	if jsonDict["btType"] == "BTTableTestCellDouble-2509" {
		// try to unmarshal JSON data into BTTableTestCellDouble2509
		var qr *BTTableTestCellDouble2509
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableCell1114 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableCell1114 = nil
			return fmt.Errorf("Failed to unmarshal BTTableCell1114 as BTTableTestCellDouble2509: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTTableTestCellString-2112'
	if jsonDict["btType"] == "BTTableTestCellString-2112" {
		// try to unmarshal JSON data into BTTableTestCellString2112
		var qr *BTTableTestCellString2112
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableCell1114 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableCell1114 = nil
			return fmt.Errorf("Failed to unmarshal BTTableCell1114 as BTTableTestCellString2112: %s", err.Error())
		}
	}

	var qtx *base_BTTableCell1114
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTTableCell1114 = qtx
		return nil // data stored in dst.base_BTTableCell1114, return on the first match
	} else {
		dst.implBTTableCell1114 = nil
		return fmt.Errorf("Failed to unmarshal BTTableCell1114 as base_BTTableCell1114: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTTableCell1114) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTTableCell1114) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTTableCell1114
}

type NullableBTTableCell1114 struct {
	value *BTTableCell1114
	isSet bool
}

func (v NullableBTTableCell1114) Get() *BTTableCell1114 {
	return v.value
}

func (v *NullableBTTableCell1114) Set(val *BTTableCell1114) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableCell1114) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableCell1114) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableCell1114(val *BTTableCell1114) *NullableBTTableCell1114 {
	return &NullableBTTableCell1114{value: val, isSet: true}
}

func (v NullableBTTableCell1114) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableCell1114) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTTableCell1114 struct {
	BtType        *string `json:"btType,omitempty"`
	IsEverVisible *bool   `json:"isEverVisible,omitempty"`
	IsReadOnly    *bool   `json:"isReadOnly,omitempty"`
}

// Newbase_BTTableCell1114 instantiates a new base_BTTableCell1114 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTTableCell1114() *base_BTTableCell1114 {
	this := base_BTTableCell1114{}
	return &this
}

// Newbase_BTTableCell1114WithDefaults instantiates a new base_BTTableCell1114 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTTableCell1114WithDefaults() *base_BTTableCell1114 {
	this := base_BTTableCell1114{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTTableCell1114) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableCell1114) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTTableCell1114) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTTableCell1114) SetBtType(v string) {
	o.BtType = &v
}

// GetIsEverVisible returns the IsEverVisible field value if set, zero value otherwise.
func (o *base_BTTableCell1114) GetIsEverVisible() bool {
	if o == nil || o.IsEverVisible == nil {
		var ret bool
		return ret
	}
	return *o.IsEverVisible
}

// GetIsEverVisibleOk returns a tuple with the IsEverVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableCell1114) GetIsEverVisibleOk() (*bool, bool) {
	if o == nil || o.IsEverVisible == nil {
		return nil, false
	}
	return o.IsEverVisible, true
}

// HasIsEverVisible returns a boolean if a field has been set.
func (o *base_BTTableCell1114) HasIsEverVisible() bool {
	if o != nil && o.IsEverVisible != nil {
		return true
	}

	return false
}

// SetIsEverVisible gets a reference to the given bool and assigns it to the IsEverVisible field.
func (o *base_BTTableCell1114) SetIsEverVisible(v bool) {
	o.IsEverVisible = &v
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *base_BTTableCell1114) GetIsReadOnly() bool {
	if o == nil || o.IsReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableCell1114) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil || o.IsReadOnly == nil {
		return nil, false
	}
	return o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *base_BTTableCell1114) HasIsReadOnly() bool {
	if o != nil && o.IsReadOnly != nil {
		return true
	}

	return false
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *base_BTTableCell1114) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

func (o base_BTTableCell1114) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsEverVisible != nil {
		toSerialize["isEverVisible"] = o.IsEverVisible
	}
	if o.IsReadOnly != nil {
		toSerialize["isReadOnly"] = o.IsReadOnly
	}
	return json.Marshal(toSerialize)
}
