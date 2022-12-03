/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.156.8083-f61e3e2b5294
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTTableBaseRowMetadata3181 - struct for BTTableBaseRowMetadata3181
type BTTableBaseRowMetadata3181 struct {
	implBTTableBaseRowMetadata3181 interface{}
}

// BTSMOtherJointTableRowMetadata2640AsBTTableBaseRowMetadata3181 is a convenience function that returns BTSMOtherJointTableRowMetadata2640 wrapped in BTTableBaseRowMetadata3181
func (o *BTSMOtherJointTableRowMetadata2640) AsBTTableBaseRowMetadata3181() *BTTableBaseRowMetadata3181 {
	return &BTTableBaseRowMetadata3181{o}
}

// BTSMBendTableRowMetadata1705AsBTTableBaseRowMetadata3181 is a convenience function that returns BTSMBendTableRowMetadata1705 wrapped in BTTableBaseRowMetadata3181
func (o *BTSMBendTableRowMetadata1705) AsBTTableBaseRowMetadata3181() *BTTableBaseRowMetadata3181 {
	return &BTTableBaseRowMetadata3181{o}
}

// BTBaseSMJointTableRowMetadata2232AsBTTableBaseRowMetadata3181 is a convenience function that returns BTBaseSMJointTableRowMetadata2232 wrapped in BTTableBaseRowMetadata3181
func (o *BTBaseSMJointTableRowMetadata2232) AsBTTableBaseRowMetadata3181() *BTTableBaseRowMetadata3181 {
	return &BTTableBaseRowMetadata3181{o}
}

// BTVariableTableRowMetadata3912AsBTTableBaseRowMetadata3181 is a convenience function that returns BTVariableTableRowMetadata3912 wrapped in BTTableBaseRowMetadata3181
func (o *BTVariableTableRowMetadata3912) AsBTTableBaseRowMetadata3181() *BTTableBaseRowMetadata3181 {
	return &BTTableBaseRowMetadata3181{o}
}

// BTBillOfMaterialsTableRowMetadata1300AsBTTableBaseRowMetadata3181 is a convenience function that returns BTBillOfMaterialsTableRowMetadata1300 wrapped in BTTableBaseRowMetadata3181
func (o *BTBillOfMaterialsTableRowMetadata1300) AsBTTableBaseRowMetadata3181() *BTTableBaseRowMetadata3181 {
	return &BTTableBaseRowMetadata3181{o}
}

// BTFSTableRowMetadata2262AsBTTableBaseRowMetadata3181 is a convenience function that returns BTFSTableRowMetadata2262 wrapped in BTTableBaseRowMetadata3181
func (o *BTFSTableRowMetadata2262) AsBTTableBaseRowMetadata3181() *BTTableBaseRowMetadata3181 {
	return &BTTableBaseRowMetadata3181{o}
}

// NewBTTableBaseRowMetadata3181 instantiates a new BTTableBaseRowMetadata3181 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableBaseRowMetadata3181() *BTTableBaseRowMetadata3181 {
	this := BTTableBaseRowMetadata3181{Newbase_BTTableBaseRowMetadata3181()}
	return &this
}

// NewBTTableBaseRowMetadata3181WithDefaults instantiates a new BTTableBaseRowMetadata3181 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableBaseRowMetadata3181WithDefaults() *BTTableBaseRowMetadata3181 {
	this := BTTableBaseRowMetadata3181{Newbase_BTTableBaseRowMetadata3181WithDefaults()}
	return &this
}

// GetCrossHighlightDataIfAny returns the CrossHighlightDataIfAny field value if set, zero value otherwise.
func (o *BTTableBaseRowMetadata3181) GetCrossHighlightDataIfAny() BTTableBaseCrossHighlightData2609 {
	type getResult interface {
		GetCrossHighlightDataIfAny() BTTableBaseCrossHighlightData2609
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetCrossHighlightDataIfAny()
	} else {
		var de BTTableBaseCrossHighlightData2609
		return de
	}
}

// GetCrossHighlightDataIfAnyOk returns a tuple with the CrossHighlightDataIfAny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableBaseRowMetadata3181) GetCrossHighlightDataIfAnyOk() (*BTTableBaseCrossHighlightData2609, bool) {
	type getResult interface {
		GetCrossHighlightDataIfAnyOk() (*BTTableBaseCrossHighlightData2609, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetCrossHighlightDataIfAnyOk()
	} else {
		return nil, false
	}
}

// HasCrossHighlightDataIfAny returns a boolean if a field has been set.
func (o *BTTableBaseRowMetadata3181) HasCrossHighlightDataIfAny() bool {
	type getResult interface {
		HasCrossHighlightDataIfAny() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasCrossHighlightDataIfAny()
	} else {
		return false
	}
}

// SetCrossHighlightDataIfAny gets a reference to the given BTTableBaseCrossHighlightData2609 and assigns it to the CrossHighlightDataIfAny field.
func (o *BTTableBaseRowMetadata3181) SetCrossHighlightDataIfAny(v BTTableBaseCrossHighlightData2609) {
	type getResult interface {
		SetCrossHighlightDataIfAny(v BTTableBaseCrossHighlightData2609)
	}

	o.GetActualInstance().(getResult).SetCrossHighlightDataIfAny(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTTableBaseRowMetadata3181) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTBaseSMJointTableRowMetadata-2232'
	if jsonDict["btType"] == "BTBaseSMJointTableRowMetadata-2232" {
		// try to unmarshal JSON data into BTBaseSMJointTableRowMetadata2232
		var qr *BTBaseSMJointTableRowMetadata2232
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableBaseRowMetadata3181 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableBaseRowMetadata3181 = nil
			return fmt.Errorf("Failed to unmarshal BTTableBaseRowMetadata3181 as BTBaseSMJointTableRowMetadata2232: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTBillOfMaterialsTableRowMetadata-1300'
	if jsonDict["btType"] == "BTBillOfMaterialsTableRowMetadata-1300" {
		// try to unmarshal JSON data into BTBillOfMaterialsTableRowMetadata1300
		var qr *BTBillOfMaterialsTableRowMetadata1300
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableBaseRowMetadata3181 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableBaseRowMetadata3181 = nil
			return fmt.Errorf("Failed to unmarshal BTTableBaseRowMetadata3181 as BTBillOfMaterialsTableRowMetadata1300: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFSTableRowMetadata-2262'
	if jsonDict["btType"] == "BTFSTableRowMetadata-2262" {
		// try to unmarshal JSON data into BTFSTableRowMetadata2262
		var qr *BTFSTableRowMetadata2262
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableBaseRowMetadata3181 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableBaseRowMetadata3181 = nil
			return fmt.Errorf("Failed to unmarshal BTTableBaseRowMetadata3181 as BTFSTableRowMetadata2262: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSMBendTableRowMetadata-1705'
	if jsonDict["btType"] == "BTSMBendTableRowMetadata-1705" {
		// try to unmarshal JSON data into BTSMBendTableRowMetadata1705
		var qr *BTSMBendTableRowMetadata1705
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableBaseRowMetadata3181 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableBaseRowMetadata3181 = nil
			return fmt.Errorf("Failed to unmarshal BTTableBaseRowMetadata3181 as BTSMBendTableRowMetadata1705: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSMOtherJointTableRowMetadata-2640'
	if jsonDict["btType"] == "BTSMOtherJointTableRowMetadata-2640" {
		// try to unmarshal JSON data into BTSMOtherJointTableRowMetadata2640
		var qr *BTSMOtherJointTableRowMetadata2640
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableBaseRowMetadata3181 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableBaseRowMetadata3181 = nil
			return fmt.Errorf("Failed to unmarshal BTTableBaseRowMetadata3181 as BTSMOtherJointTableRowMetadata2640: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTVariableTableRowMetadata-3912'
	if jsonDict["btType"] == "BTVariableTableRowMetadata-3912" {
		// try to unmarshal JSON data into BTVariableTableRowMetadata3912
		var qr *BTVariableTableRowMetadata3912
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableBaseRowMetadata3181 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableBaseRowMetadata3181 = nil
			return fmt.Errorf("Failed to unmarshal BTTableBaseRowMetadata3181 as BTVariableTableRowMetadata3912: %s", err.Error())
		}
	}

	var qtx *base_BTTableBaseRowMetadata3181
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTTableBaseRowMetadata3181 = qtx
		return nil // data stored in dst.base_BTTableBaseRowMetadata3181, return on the first match
	} else {
		dst.implBTTableBaseRowMetadata3181 = nil
		return fmt.Errorf("Failed to unmarshal BTTableBaseRowMetadata3181 as base_BTTableBaseRowMetadata3181: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTTableBaseRowMetadata3181) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTTableBaseRowMetadata3181) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTTableBaseRowMetadata3181
}

type NullableBTTableBaseRowMetadata3181 struct {
	value *BTTableBaseRowMetadata3181
	isSet bool
}

func (v NullableBTTableBaseRowMetadata3181) Get() *BTTableBaseRowMetadata3181 {
	return v.value
}

func (v *NullableBTTableBaseRowMetadata3181) Set(val *BTTableBaseRowMetadata3181) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableBaseRowMetadata3181) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableBaseRowMetadata3181) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableBaseRowMetadata3181(val *BTTableBaseRowMetadata3181) *NullableBTTableBaseRowMetadata3181 {
	return &NullableBTTableBaseRowMetadata3181{value: val, isSet: true}
}

func (v NullableBTTableBaseRowMetadata3181) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableBaseRowMetadata3181) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTTableBaseRowMetadata3181 struct {
	CrossHighlightDataIfAny *BTTableBaseCrossHighlightData2609 `json:"crossHighlightDataIfAny,omitempty"`
}

// Newbase_BTTableBaseRowMetadata3181 instantiates a new base_BTTableBaseRowMetadata3181 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTTableBaseRowMetadata3181() *base_BTTableBaseRowMetadata3181 {
	this := base_BTTableBaseRowMetadata3181{}
	return &this
}

// Newbase_BTTableBaseRowMetadata3181WithDefaults instantiates a new base_BTTableBaseRowMetadata3181 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTTableBaseRowMetadata3181WithDefaults() *base_BTTableBaseRowMetadata3181 {
	this := base_BTTableBaseRowMetadata3181{}
	return &this
}

// GetCrossHighlightDataIfAny returns the CrossHighlightDataIfAny field value if set, zero value otherwise.
func (o *base_BTTableBaseRowMetadata3181) GetCrossHighlightDataIfAny() BTTableBaseCrossHighlightData2609 {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		var ret BTTableBaseCrossHighlightData2609
		return ret
	}
	return *o.CrossHighlightDataIfAny
}

// GetCrossHighlightDataIfAnyOk returns a tuple with the CrossHighlightDataIfAny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableBaseRowMetadata3181) GetCrossHighlightDataIfAnyOk() (*BTTableBaseCrossHighlightData2609, bool) {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		return nil, false
	}
	return o.CrossHighlightDataIfAny, true
}

// HasCrossHighlightDataIfAny returns a boolean if a field has been set.
func (o *base_BTTableBaseRowMetadata3181) HasCrossHighlightDataIfAny() bool {
	if o != nil && o.CrossHighlightDataIfAny != nil {
		return true
	}

	return false
}

// SetCrossHighlightDataIfAny gets a reference to the given BTTableBaseCrossHighlightData2609 and assigns it to the CrossHighlightDataIfAny field.
func (o *base_BTTableBaseRowMetadata3181) SetCrossHighlightDataIfAny(v BTTableBaseCrossHighlightData2609) {
	o.CrossHighlightDataIfAny = &v
}

func (o base_BTTableBaseRowMetadata3181) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CrossHighlightDataIfAny != nil {
		toSerialize["crossHighlightDataIfAny"] = o.CrossHighlightDataIfAny
	}
	return json.Marshal(toSerialize)
}
