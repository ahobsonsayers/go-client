/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13995-cdd961a1a6ad
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTDocumentSelectorParametersInfo - struct for BTDocumentSelectorParametersInfo
type BTDocumentSelectorParametersInfo struct {
	implBTDocumentSelectorParametersInfo interface{}
}

// BTOtherDocumentSelectorParametersInfoAsBTDocumentSelectorParametersInfo is a convenience function that returns BTOtherDocumentSelectorParametersInfo wrapped in BTDocumentSelectorParametersInfo
func (o *BTOtherDocumentSelectorParametersInfo) AsBTDocumentSelectorParametersInfo() *BTDocumentSelectorParametersInfo {
	return &BTDocumentSelectorParametersInfo{o}
}

// NewBTDocumentSelectorParametersInfo instantiates a new BTDocumentSelectorParametersInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentSelectorParametersInfo(jsonType string) *BTDocumentSelectorParametersInfo {
	this := BTDocumentSelectorParametersInfo{Newbase_BTDocumentSelectorParametersInfo(jsonType)}
	return &this
}

// NewBTDocumentSelectorParametersInfoWithDefaults instantiates a new BTDocumentSelectorParametersInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentSelectorParametersInfoWithDefaults() *BTDocumentSelectorParametersInfo {
	this := BTDocumentSelectorParametersInfo{Newbase_BTDocumentSelectorParametersInfoWithDefaults()}
	return &this
}

// GetJsonType returns the JsonType field value
func (o *BTDocumentSelectorParametersInfo) GetJsonType() string {
	type getResult interface {
		GetJsonType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetJsonType()
	} else {
		var de string
		return de
	}
}

// GetJsonTypeOk returns a tuple with the JsonType field value
// and a boolean to check if the value has been set.
func (o *BTDocumentSelectorParametersInfo) GetJsonTypeOk() (*string, bool) {
	type getResult interface {
		GetJsonTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetJsonTypeOk()
	} else {
		return nil, false
	}
}

// SetJsonType sets field value
func (o *BTDocumentSelectorParametersInfo) SetJsonType(v string) {
	type getResult interface {
		SetJsonType(v string)
	}

	o.GetActualInstance().(getResult).SetJsonType(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTDocumentSelectorParametersInfo) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTOtherDocumentSelectorParametersInfo'
	if jsonDict["jsonType"] == "BTOtherDocumentSelectorParametersInfo" {
		// try to unmarshal JSON data into BTOtherDocumentSelectorParametersInfo
		var qr *BTOtherDocumentSelectorParametersInfo
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTDocumentSelectorParametersInfo = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTDocumentSelectorParametersInfo = nil
			return fmt.Errorf("Failed to unmarshal BTDocumentSelectorParametersInfo as BTOtherDocumentSelectorParametersInfo: %s", err.Error())
		}
	}

	// check if the discriminator value is 'other-documents'
	if jsonDict["jsonType"] == "other-documents" {
		// try to unmarshal JSON data into BTOtherDocumentSelectorParametersInfo
		var qr *BTOtherDocumentSelectorParametersInfo
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTDocumentSelectorParametersInfo = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTDocumentSelectorParametersInfo = nil
			return fmt.Errorf("Failed to unmarshal BTDocumentSelectorParametersInfo as BTOtherDocumentSelectorParametersInfo: %s", err.Error())
		}
	}

	var qtx *base_BTDocumentSelectorParametersInfo
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTDocumentSelectorParametersInfo = qtx
		return nil // data stored in dst.base_BTDocumentSelectorParametersInfo, return on the first match
	} else {
		dst.implBTDocumentSelectorParametersInfo = nil
		return fmt.Errorf("Failed to unmarshal BTDocumentSelectorParametersInfo as base_BTDocumentSelectorParametersInfo: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTDocumentSelectorParametersInfo) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTDocumentSelectorParametersInfo) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTDocumentSelectorParametersInfo
}

type NullableBTDocumentSelectorParametersInfo struct {
	value *BTDocumentSelectorParametersInfo
	isSet bool
}

func (v NullableBTDocumentSelectorParametersInfo) Get() *BTDocumentSelectorParametersInfo {
	return v.value
}

func (v *NullableBTDocumentSelectorParametersInfo) Set(val *BTDocumentSelectorParametersInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentSelectorParametersInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentSelectorParametersInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentSelectorParametersInfo(val *BTDocumentSelectorParametersInfo) *NullableBTDocumentSelectorParametersInfo {
	return &NullableBTDocumentSelectorParametersInfo{value: val, isSet: true}
}

func (v NullableBTDocumentSelectorParametersInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentSelectorParametersInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTDocumentSelectorParametersInfo struct {
	JsonType string `json:"jsonType"`
}

// Newbase_BTDocumentSelectorParametersInfo instantiates a new base_BTDocumentSelectorParametersInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTDocumentSelectorParametersInfo(jsonType string) *base_BTDocumentSelectorParametersInfo {
	this := base_BTDocumentSelectorParametersInfo{}
	this.JsonType = jsonType
	return &this
}

// Newbase_BTDocumentSelectorParametersInfoWithDefaults instantiates a new base_BTDocumentSelectorParametersInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTDocumentSelectorParametersInfoWithDefaults() *base_BTDocumentSelectorParametersInfo {
	this := base_BTDocumentSelectorParametersInfo{}
	return &this
}

// GetJsonType returns the JsonType field value
func (o *base_BTDocumentSelectorParametersInfo) GetJsonType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value
// and a boolean to check if the value has been set.
func (o *base_BTDocumentSelectorParametersInfo) GetJsonTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsonType, true
}

// SetJsonType sets field value
func (o *base_BTDocumentSelectorParametersInfo) SetJsonType(v string) {
	o.JsonType = v
}

func (o base_BTDocumentSelectorParametersInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["jsonType"] = o.JsonType
	}
	return json.Marshal(toSerialize)
}
