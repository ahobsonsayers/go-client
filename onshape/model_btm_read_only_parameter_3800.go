/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6870-e8e79a24dc2c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTMReadOnlyParameter3800 - struct for BTMReadOnlyParameter3800
type BTMReadOnlyParameter3800 struct {
	implBTMReadOnlyParameter3800 interface{}
}

// BTMParameterStringWithTolerances4286AsBTMReadOnlyParameter3800 is a convenience function that returns BTMParameterStringWithTolerances4286 wrapped in BTMReadOnlyParameter3800
func (o *BTMParameterStringWithTolerances4286) AsBTMReadOnlyParameter3800() *BTMReadOnlyParameter3800 {
	return &BTMReadOnlyParameter3800{o}
}

// BTMParameterProgress3232AsBTMReadOnlyParameter3800 is a convenience function that returns BTMParameterProgress3232 wrapped in BTMReadOnlyParameter3800
func (o *BTMParameterProgress3232) AsBTMReadOnlyParameter3800() *BTMReadOnlyParameter3800 {
	return &BTMReadOnlyParameter3800{o}
}

// NewBTMReadOnlyParameter3800 instantiates a new BTMReadOnlyParameter3800 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMReadOnlyParameter3800() *BTMReadOnlyParameter3800 {
	this := BTMReadOnlyParameter3800{Newbase_BTMReadOnlyParameter3800()}
	return &this
}

// NewBTMReadOnlyParameter3800WithDefaults instantiates a new BTMReadOnlyParameter3800 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMReadOnlyParameter3800WithDefaults() *BTMReadOnlyParameter3800 {
	this := BTMReadOnlyParameter3800{Newbase_BTMReadOnlyParameter3800WithDefaults()}
	return &this
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMReadOnlyParameter3800) GetImportMicroversion() string {
	type getResult interface {
		GetImportMicroversion() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversion()
	} else {
		var de string
		return de
	}
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMReadOnlyParameter3800) GetImportMicroversionOk() (*string, bool) {
	type getResult interface {
		GetImportMicroversionOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversionOk()
	} else {
		return nil, false
	}
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMReadOnlyParameter3800) HasImportMicroversion() bool {
	type getResult interface {
		HasImportMicroversion() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasImportMicroversion()
	} else {
		return false
	}
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMReadOnlyParameter3800) SetImportMicroversion(v string) {
	type getResult interface {
		SetImportMicroversion(v string)
	}

	o.GetActualInstance().(getResult).SetImportMicroversion(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMReadOnlyParameter3800) GetNodeId() string {
	type getResult interface {
		GetNodeId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeId()
	} else {
		var de string
		return de
	}
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMReadOnlyParameter3800) GetNodeIdOk() (*string, bool) {
	type getResult interface {
		GetNodeIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeIdOk()
	} else {
		return nil, false
	}
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMReadOnlyParameter3800) HasNodeId() bool {
	type getResult interface {
		HasNodeId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasNodeId()
	} else {
		return false
	}
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMReadOnlyParameter3800) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMReadOnlyParameter3800) GetParameterId() string {
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
func (o *BTMReadOnlyParameter3800) GetParameterIdOk() (*string, bool) {
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
func (o *BTMReadOnlyParameter3800) HasParameterId() bool {
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
func (o *BTMReadOnlyParameter3800) SetParameterId(v string) {
	type getResult interface {
		SetParameterId(v string)
	}

	o.GetActualInstance().(getResult).SetParameterId(v)
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMReadOnlyParameter3800) GetBtType() string {
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
func (o *BTMReadOnlyParameter3800) GetBtTypeOk() (*string, bool) {
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
func (o *BTMReadOnlyParameter3800) HasBtType() bool {
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
func (o *BTMReadOnlyParameter3800) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTMReadOnlyParameter3800) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTMParameterProgress-3232'
	if jsonDict["btType"] == "BTMParameterProgress-3232" {
		// try to unmarshal JSON data into BTMParameterProgress3232
		var qr *BTMParameterProgress3232
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMReadOnlyParameter3800 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMReadOnlyParameter3800 = nil
			return fmt.Errorf("Failed to unmarshal BTMReadOnlyParameter3800 as BTMParameterProgress3232: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterStringWithTolerances-4286'
	if jsonDict["btType"] == "BTMParameterStringWithTolerances-4286" {
		// try to unmarshal JSON data into BTMParameterStringWithTolerances4286
		var qr *BTMParameterStringWithTolerances4286
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMReadOnlyParameter3800 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMReadOnlyParameter3800 = nil
			return fmt.Errorf("Failed to unmarshal BTMReadOnlyParameter3800 as BTMParameterStringWithTolerances4286: %s", err.Error())
		}
	}

	var qtx *base_BTMReadOnlyParameter3800
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTMReadOnlyParameter3800 = qtx
		return nil // data stored in dst.base_BTMReadOnlyParameter3800, return on the first match
	} else {
		dst.implBTMReadOnlyParameter3800 = nil
		return fmt.Errorf("Failed to unmarshal BTMReadOnlyParameter3800 as base_BTMReadOnlyParameter3800: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTMReadOnlyParameter3800) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTMReadOnlyParameter3800) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTMReadOnlyParameter3800
}

type NullableBTMReadOnlyParameter3800 struct {
	value *BTMReadOnlyParameter3800
	isSet bool
}

func (v NullableBTMReadOnlyParameter3800) Get() *BTMReadOnlyParameter3800 {
	return v.value
}

func (v *NullableBTMReadOnlyParameter3800) Set(val *BTMReadOnlyParameter3800) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMReadOnlyParameter3800) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMReadOnlyParameter3800) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMReadOnlyParameter3800(val *BTMReadOnlyParameter3800) *NullableBTMReadOnlyParameter3800 {
	return &NullableBTMReadOnlyParameter3800{value: val, isSet: true}
}

func (v NullableBTMReadOnlyParameter3800) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMReadOnlyParameter3800) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTMReadOnlyParameter3800 struct {
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	NodeId             *string `json:"nodeId,omitempty"`
	ParameterId        *string `json:"parameterId,omitempty"`
	BtType             *string `json:"btType,omitempty"`
}

// Newbase_BTMReadOnlyParameter3800 instantiates a new base_BTMReadOnlyParameter3800 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTMReadOnlyParameter3800() *base_BTMReadOnlyParameter3800 {
	this := base_BTMReadOnlyParameter3800{}
	return &this
}

// Newbase_BTMReadOnlyParameter3800WithDefaults instantiates a new base_BTMReadOnlyParameter3800 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTMReadOnlyParameter3800WithDefaults() *base_BTMReadOnlyParameter3800 {
	this := base_BTMReadOnlyParameter3800{}
	return &this
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *base_BTMReadOnlyParameter3800) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMReadOnlyParameter3800) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *base_BTMReadOnlyParameter3800) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *base_BTMReadOnlyParameter3800) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTMReadOnlyParameter3800) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMReadOnlyParameter3800) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTMReadOnlyParameter3800) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTMReadOnlyParameter3800) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *base_BTMReadOnlyParameter3800) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMReadOnlyParameter3800) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *base_BTMReadOnlyParameter3800) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *base_BTMReadOnlyParameter3800) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTMReadOnlyParameter3800) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMReadOnlyParameter3800) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTMReadOnlyParameter3800) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTMReadOnlyParameter3800) SetBtType(v string) {
	o.BtType = &v
}

func (o base_BTMReadOnlyParameter3800) MarshalJSON() ([]byte, error) {
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
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}
