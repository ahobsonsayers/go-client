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
	"fmt"
)

// BTMConfigurationParameter819 - struct for BTMConfigurationParameter819
type BTMConfigurationParameter819 struct {
	implBTMConfigurationParameter819 interface{}
}

// BTMConfigurationParameterBoolean2550AsBTMConfigurationParameter819 is a convenience function that returns BTMConfigurationParameterBoolean2550 wrapped in BTMConfigurationParameter819
func (o *BTMConfigurationParameterBoolean2550) AsBTMConfigurationParameter819() *BTMConfigurationParameter819 {
	return &BTMConfigurationParameter819{o}
}

// BTMConfigurationParameterQuantity1826AsBTMConfigurationParameter819 is a convenience function that returns BTMConfigurationParameterQuantity1826 wrapped in BTMConfigurationParameter819
func (o *BTMConfigurationParameterQuantity1826) AsBTMConfigurationParameter819() *BTMConfigurationParameter819 {
	return &BTMConfigurationParameter819{o}
}

// BTMConfigurationParameterString872AsBTMConfigurationParameter819 is a convenience function that returns BTMConfigurationParameterString872 wrapped in BTMConfigurationParameter819
func (o *BTMConfigurationParameterString872) AsBTMConfigurationParameter819() *BTMConfigurationParameter819 {
	return &BTMConfigurationParameter819{o}
}

// BTMConfigurationParameterEnum105AsBTMConfigurationParameter819 is a convenience function that returns BTMConfigurationParameterEnum105 wrapped in BTMConfigurationParameter819
func (o *BTMConfigurationParameterEnum105) AsBTMConfigurationParameter819() *BTMConfigurationParameter819 {
	return &BTMConfigurationParameter819{o}
}

// NewBTMConfigurationParameter819 instantiates a new BTMConfigurationParameter819 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMConfigurationParameter819() *BTMConfigurationParameter819 {
	this := BTMConfigurationParameter819{Newbase_BTMConfigurationParameter819()}
	return &this
}

// NewBTMConfigurationParameter819WithDefaults instantiates a new BTMConfigurationParameter819 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMConfigurationParameter819WithDefaults() *BTMConfigurationParameter819 {
	this := BTMConfigurationParameter819{Newbase_BTMConfigurationParameter819WithDefaults()}
	return &this
}

// GetGeneratedParameterId returns the GeneratedParameterId field value if set, zero value otherwise.
func (o *BTMConfigurationParameter819) GetGeneratedParameterId() BTTreeNode20 {
	type getResult interface {
		GetGeneratedParameterId() BTTreeNode20
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetGeneratedParameterId()
	} else {
		var de BTTreeNode20
		return de
	}
}

// GetGeneratedParameterIdOk returns a tuple with the GeneratedParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameter819) GetGeneratedParameterIdOk() (*BTTreeNode20, bool) {
	type getResult interface {
		GetGeneratedParameterIdOk() (*BTTreeNode20, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetGeneratedParameterIdOk()
	} else {
		return nil, false
	}
}

// HasGeneratedParameterId returns a boolean if a field has been set.
func (o *BTMConfigurationParameter819) HasGeneratedParameterId() bool {
	type getResult interface {
		HasGeneratedParameterId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasGeneratedParameterId()
	} else {
		return false
	}
}

// SetGeneratedParameterId gets a reference to the given BTTreeNode20 and assigns it to the GeneratedParameterId field.
func (o *BTMConfigurationParameter819) SetGeneratedParameterId(v BTTreeNode20) {
	type getResult interface {
		SetGeneratedParameterId(v BTTreeNode20)
	}

	o.GetActualInstance().(getResult).SetGeneratedParameterId(v)
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMConfigurationParameter819) GetImportMicroversion() string {
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
func (o *BTMConfigurationParameter819) GetImportMicroversionOk() (*string, bool) {
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
func (o *BTMConfigurationParameter819) HasImportMicroversion() bool {
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
func (o *BTMConfigurationParameter819) SetImportMicroversion(v string) {
	type getResult interface {
		SetImportMicroversion(v string)
	}

	o.GetActualInstance().(getResult).SetImportMicroversion(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMConfigurationParameter819) GetNodeId() string {
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
func (o *BTMConfigurationParameter819) GetNodeIdOk() (*string, bool) {
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
func (o *BTMConfigurationParameter819) HasNodeId() bool {
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
func (o *BTMConfigurationParameter819) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMConfigurationParameter819) GetParameterId() string {
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
func (o *BTMConfigurationParameter819) GetParameterIdOk() (*string, bool) {
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
func (o *BTMConfigurationParameter819) HasParameterId() bool {
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
func (o *BTMConfigurationParameter819) SetParameterId(v string) {
	type getResult interface {
		SetParameterId(v string)
	}

	o.GetActualInstance().(getResult).SetParameterId(v)
}

// GetParameterName returns the ParameterName field value if set, zero value otherwise.
func (o *BTMConfigurationParameter819) GetParameterName() string {
	type getResult interface {
		GetParameterName() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParameterName()
	} else {
		var de string
		return de
	}
}

// GetParameterNameOk returns a tuple with the ParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameter819) GetParameterNameOk() (*string, bool) {
	type getResult interface {
		GetParameterNameOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParameterNameOk()
	} else {
		return nil, false
	}
}

// HasParameterName returns a boolean if a field has been set.
func (o *BTMConfigurationParameter819) HasParameterName() bool {
	type getResult interface {
		HasParameterName() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasParameterName()
	} else {
		return false
	}
}

// SetParameterName gets a reference to the given string and assigns it to the ParameterName field.
func (o *BTMConfigurationParameter819) SetParameterName(v string) {
	type getResult interface {
		SetParameterName(v string)
	}

	o.GetActualInstance().(getResult).SetParameterName(v)
}

// GetParameterType returns the ParameterType field value if set, zero value otherwise.
func (o *BTMConfigurationParameter819) GetParameterType() string {
	type getResult interface {
		GetParameterType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParameterType()
	} else {
		var de string
		return de
	}
}

// GetParameterTypeOk returns a tuple with the ParameterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameter819) GetParameterTypeOk() (*string, bool) {
	type getResult interface {
		GetParameterTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParameterTypeOk()
	} else {
		return nil, false
	}
}

// HasParameterType returns a boolean if a field has been set.
func (o *BTMConfigurationParameter819) HasParameterType() bool {
	type getResult interface {
		HasParameterType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasParameterType()
	} else {
		return false
	}
}

// SetParameterType gets a reference to the given string and assigns it to the ParameterType field.
func (o *BTMConfigurationParameter819) SetParameterType(v string) {
	type getResult interface {
		SetParameterType(v string)
	}

	o.GetActualInstance().(getResult).SetParameterType(v)
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *BTMConfigurationParameter819) GetValid() bool {
	type getResult interface {
		GetValid() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetValid()
	} else {
		var de bool
		return de
	}
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameter819) GetValidOk() (*bool, bool) {
	type getResult interface {
		GetValidOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetValidOk()
	} else {
		return nil, false
	}
}

// HasValid returns a boolean if a field has been set.
func (o *BTMConfigurationParameter819) HasValid() bool {
	type getResult interface {
		HasValid() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasValid()
	} else {
		return false
	}
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *BTMConfigurationParameter819) SetValid(v bool) {
	type getResult interface {
		SetValid(v bool)
	}

	o.GetActualInstance().(getResult).SetValid(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTMConfigurationParameter819) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTMConfigurationParameterBoolean-2550'
	if jsonDict["btType"] == "BTMConfigurationParameterBoolean-2550" {
		// try to unmarshal JSON data into BTMConfigurationParameterBoolean2550
		var qr *BTMConfigurationParameterBoolean2550
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMConfigurationParameter819 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMConfigurationParameter819 = nil
			return fmt.Errorf("Failed to unmarshal BTMConfigurationParameter819 as BTMConfigurationParameterBoolean2550: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMConfigurationParameterEnum-105'
	if jsonDict["btType"] == "BTMConfigurationParameterEnum-105" {
		// try to unmarshal JSON data into BTMConfigurationParameterEnum105
		var qr *BTMConfigurationParameterEnum105
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMConfigurationParameter819 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMConfigurationParameter819 = nil
			return fmt.Errorf("Failed to unmarshal BTMConfigurationParameter819 as BTMConfigurationParameterEnum105: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMConfigurationParameterQuantity-1826'
	if jsonDict["btType"] == "BTMConfigurationParameterQuantity-1826" {
		// try to unmarshal JSON data into BTMConfigurationParameterQuantity1826
		var qr *BTMConfigurationParameterQuantity1826
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMConfigurationParameter819 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMConfigurationParameter819 = nil
			return fmt.Errorf("Failed to unmarshal BTMConfigurationParameter819 as BTMConfigurationParameterQuantity1826: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMConfigurationParameterString-872'
	if jsonDict["btType"] == "BTMConfigurationParameterString-872" {
		// try to unmarshal JSON data into BTMConfigurationParameterString872
		var qr *BTMConfigurationParameterString872
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMConfigurationParameter819 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMConfigurationParameter819 = nil
			return fmt.Errorf("Failed to unmarshal BTMConfigurationParameter819 as BTMConfigurationParameterString872: %s", err.Error())
		}
	}

	var qtx *base_BTMConfigurationParameter819
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTMConfigurationParameter819 = qtx
		return nil // data stored in dst.base_BTMConfigurationParameter819, return on the first match
	} else {
		dst.implBTMConfigurationParameter819 = nil
		return fmt.Errorf("Failed to unmarshal BTMConfigurationParameter819 as base_BTMConfigurationParameter819: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTMConfigurationParameter819) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTMConfigurationParameter819) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTMConfigurationParameter819
}

type NullableBTMConfigurationParameter819 struct {
	value *BTMConfigurationParameter819
	isSet bool
}

func (v NullableBTMConfigurationParameter819) Get() *BTMConfigurationParameter819 {
	return v.value
}

func (v *NullableBTMConfigurationParameter819) Set(val *BTMConfigurationParameter819) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMConfigurationParameter819) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMConfigurationParameter819) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMConfigurationParameter819(val *BTMConfigurationParameter819) *NullableBTMConfigurationParameter819 {
	return &NullableBTMConfigurationParameter819{value: val, isSet: true}
}

func (v NullableBTMConfigurationParameter819) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMConfigurationParameter819) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTMConfigurationParameter819 struct {
	GeneratedParameterId *BTTreeNode20 `json:"generatedParameterId,omitempty"`
	ImportMicroversion   *string       `json:"importMicroversion,omitempty"`
	NodeId               *string       `json:"nodeId,omitempty"`
	ParameterId          *string       `json:"parameterId,omitempty"`
	ParameterName        *string       `json:"parameterName,omitempty"`
	ParameterType        *string       `json:"parameterType,omitempty"`
	Valid                *bool         `json:"valid,omitempty"`
}

// Newbase_BTMConfigurationParameter819 instantiates a new base_BTMConfigurationParameter819 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTMConfigurationParameter819() *base_BTMConfigurationParameter819 {
	this := base_BTMConfigurationParameter819{}
	return &this
}

// Newbase_BTMConfigurationParameter819WithDefaults instantiates a new base_BTMConfigurationParameter819 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTMConfigurationParameter819WithDefaults() *base_BTMConfigurationParameter819 {
	this := base_BTMConfigurationParameter819{}
	return &this
}

// GetGeneratedParameterId returns the GeneratedParameterId field value if set, zero value otherwise.
func (o *base_BTMConfigurationParameter819) GetGeneratedParameterId() BTTreeNode20 {
	if o == nil || o.GeneratedParameterId == nil {
		var ret BTTreeNode20
		return ret
	}
	return *o.GeneratedParameterId
}

// GetGeneratedParameterIdOk returns a tuple with the GeneratedParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMConfigurationParameter819) GetGeneratedParameterIdOk() (*BTTreeNode20, bool) {
	if o == nil || o.GeneratedParameterId == nil {
		return nil, false
	}
	return o.GeneratedParameterId, true
}

// HasGeneratedParameterId returns a boolean if a field has been set.
func (o *base_BTMConfigurationParameter819) HasGeneratedParameterId() bool {
	if o != nil && o.GeneratedParameterId != nil {
		return true
	}

	return false
}

// SetGeneratedParameterId gets a reference to the given BTTreeNode20 and assigns it to the GeneratedParameterId field.
func (o *base_BTMConfigurationParameter819) SetGeneratedParameterId(v BTTreeNode20) {
	o.GeneratedParameterId = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *base_BTMConfigurationParameter819) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMConfigurationParameter819) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *base_BTMConfigurationParameter819) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *base_BTMConfigurationParameter819) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTMConfigurationParameter819) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMConfigurationParameter819) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTMConfigurationParameter819) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTMConfigurationParameter819) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *base_BTMConfigurationParameter819) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMConfigurationParameter819) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *base_BTMConfigurationParameter819) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *base_BTMConfigurationParameter819) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetParameterName returns the ParameterName field value if set, zero value otherwise.
func (o *base_BTMConfigurationParameter819) GetParameterName() string {
	if o == nil || o.ParameterName == nil {
		var ret string
		return ret
	}
	return *o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMConfigurationParameter819) GetParameterNameOk() (*string, bool) {
	if o == nil || o.ParameterName == nil {
		return nil, false
	}
	return o.ParameterName, true
}

// HasParameterName returns a boolean if a field has been set.
func (o *base_BTMConfigurationParameter819) HasParameterName() bool {
	if o != nil && o.ParameterName != nil {
		return true
	}

	return false
}

// SetParameterName gets a reference to the given string and assigns it to the ParameterName field.
func (o *base_BTMConfigurationParameter819) SetParameterName(v string) {
	o.ParameterName = &v
}

// GetParameterType returns the ParameterType field value if set, zero value otherwise.
func (o *base_BTMConfigurationParameter819) GetParameterType() string {
	if o == nil || o.ParameterType == nil {
		var ret string
		return ret
	}
	return *o.ParameterType
}

// GetParameterTypeOk returns a tuple with the ParameterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMConfigurationParameter819) GetParameterTypeOk() (*string, bool) {
	if o == nil || o.ParameterType == nil {
		return nil, false
	}
	return o.ParameterType, true
}

// HasParameterType returns a boolean if a field has been set.
func (o *base_BTMConfigurationParameter819) HasParameterType() bool {
	if o != nil && o.ParameterType != nil {
		return true
	}

	return false
}

// SetParameterType gets a reference to the given string and assigns it to the ParameterType field.
func (o *base_BTMConfigurationParameter819) SetParameterType(v string) {
	o.ParameterType = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *base_BTMConfigurationParameter819) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMConfigurationParameter819) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *base_BTMConfigurationParameter819) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *base_BTMConfigurationParameter819) SetValid(v bool) {
	o.Valid = &v
}

func (o base_BTMConfigurationParameter819) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GeneratedParameterId != nil {
		toSerialize["generatedParameterId"] = o.GeneratedParameterId
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.ParameterName != nil {
		toSerialize["parameterName"] = o.ParameterName
	}
	if o.ParameterType != nil {
		toSerialize["parameterType"] = o.ParameterType
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}
	return json.Marshal(toSerialize)
}
