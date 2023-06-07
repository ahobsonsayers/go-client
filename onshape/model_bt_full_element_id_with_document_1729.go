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

// BTFullElementIdWithDocument1729 - struct for BTFullElementIdWithDocument1729
type BTFullElementIdWithDocument1729 struct {
	implBTFullElementIdWithDocument1729 interface{}
}

// BTFullElementIdAndPartId643AsBTFullElementIdWithDocument1729 is a convenience function that returns BTFullElementIdAndPartId643 wrapped in BTFullElementIdWithDocument1729
func (o *BTFullElementIdAndPartId643) AsBTFullElementIdWithDocument1729() *BTFullElementIdWithDocument1729 {
	return &BTFullElementIdWithDocument1729{o}
}

// NewBTFullElementIdWithDocument1729 instantiates a new BTFullElementIdWithDocument1729 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFullElementIdWithDocument1729() *BTFullElementIdWithDocument1729 {
	this := BTFullElementIdWithDocument1729{Newbase_BTFullElementIdWithDocument1729()}
	return &this
}

// NewBTFullElementIdWithDocument1729WithDefaults instantiates a new BTFullElementIdWithDocument1729 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFullElementIdWithDocument1729WithDefaults() *BTFullElementIdWithDocument1729 {
	this := BTFullElementIdWithDocument1729{Newbase_BTFullElementIdWithDocument1729WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFullElementIdWithDocument1729) GetBtType() string {
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
func (o *BTFullElementIdWithDocument1729) GetBtTypeOk() (*string, bool) {
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
func (o *BTFullElementIdWithDocument1729) HasBtType() bool {
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
func (o *BTFullElementIdWithDocument1729) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetConfigured returns the Configured field value if set, zero value otherwise.
func (o *BTFullElementIdWithDocument1729) GetConfigured() bool {
	type getResult interface {
		GetConfigured() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetConfigured()
	} else {
		var de bool
		return de
	}
}

// GetConfiguredOk returns a tuple with the Configured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementIdWithDocument1729) GetConfiguredOk() (*bool, bool) {
	type getResult interface {
		GetConfiguredOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetConfiguredOk()
	} else {
		return nil, false
	}
}

// HasConfigured returns a boolean if a field has been set.
func (o *BTFullElementIdWithDocument1729) HasConfigured() bool {
	type getResult interface {
		HasConfigured() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasConfigured()
	} else {
		return false
	}
}

// SetConfigured gets a reference to the given bool and assigns it to the Configured field.
func (o *BTFullElementIdWithDocument1729) SetConfigured(v bool) {
	type getResult interface {
		SetConfigured(v bool)
	}

	o.GetActualInstance().(getResult).SetConfigured(v)
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTFullElementIdWithDocument1729) GetDocumentId() string {
	type getResult interface {
		GetDocumentId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDocumentId()
	} else {
		var de string
		return de
	}
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementIdWithDocument1729) GetDocumentIdOk() (*string, bool) {
	type getResult interface {
		GetDocumentIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDocumentIdOk()
	} else {
		return nil, false
	}
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTFullElementIdWithDocument1729) HasDocumentId() bool {
	type getResult interface {
		HasDocumentId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDocumentId()
	} else {
		return false
	}
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTFullElementIdWithDocument1729) SetDocumentId(v string) {
	type getResult interface {
		SetDocumentId(v string)
	}

	o.GetActualInstance().(getResult).SetDocumentId(v)
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTFullElementIdWithDocument1729) GetElementId() string {
	type getResult interface {
		GetElementId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetElementId()
	} else {
		var de string
		return de
	}
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementIdWithDocument1729) GetElementIdOk() (*string, bool) {
	type getResult interface {
		GetElementIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetElementIdOk()
	} else {
		return nil, false
	}
}

// HasElementId returns a boolean if a field has been set.
func (o *BTFullElementIdWithDocument1729) HasElementId() bool {
	type getResult interface {
		HasElementId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasElementId()
	} else {
		return false
	}
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTFullElementIdWithDocument1729) SetElementId(v string) {
	type getResult interface {
		SetElementId(v string)
	}

	o.GetActualInstance().(getResult).SetElementId(v)
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTFullElementIdWithDocument1729) GetMicroversionId() BTMicroversionId366 {
	type getResult interface {
		GetMicroversionId() BTMicroversionId366
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMicroversionId()
	} else {
		var de BTMicroversionId366
		return de
	}
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementIdWithDocument1729) GetMicroversionIdOk() (*BTMicroversionId366, bool) {
	type getResult interface {
		GetMicroversionIdOk() (*BTMicroversionId366, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMicroversionIdOk()
	} else {
		return nil, false
	}
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTFullElementIdWithDocument1729) HasMicroversionId() bool {
	type getResult interface {
		HasMicroversionId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasMicroversionId()
	} else {
		return false
	}
}

// SetMicroversionId gets a reference to the given BTMicroversionId366 and assigns it to the MicroversionId field.
func (o *BTFullElementIdWithDocument1729) SetMicroversionId(v BTMicroversionId366) {
	type getResult interface {
		SetMicroversionId(v BTMicroversionId366)
	}

	o.GetActualInstance().(getResult).SetMicroversionId(v)
}

// GetMicroversionIdAndConfiguration returns the MicroversionIdAndConfiguration field value if set, zero value otherwise.
func (o *BTFullElementIdWithDocument1729) GetMicroversionIdAndConfiguration() BTMicroversionIdAndConfiguration2338 {
	type getResult interface {
		GetMicroversionIdAndConfiguration() BTMicroversionIdAndConfiguration2338
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMicroversionIdAndConfiguration()
	} else {
		var de BTMicroversionIdAndConfiguration2338
		return de
	}
}

// GetMicroversionIdAndConfigurationOk returns a tuple with the MicroversionIdAndConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementIdWithDocument1729) GetMicroversionIdAndConfigurationOk() (*BTMicroversionIdAndConfiguration2338, bool) {
	type getResult interface {
		GetMicroversionIdAndConfigurationOk() (*BTMicroversionIdAndConfiguration2338, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMicroversionIdAndConfigurationOk()
	} else {
		return nil, false
	}
}

// HasMicroversionIdAndConfiguration returns a boolean if a field has been set.
func (o *BTFullElementIdWithDocument1729) HasMicroversionIdAndConfiguration() bool {
	type getResult interface {
		HasMicroversionIdAndConfiguration() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasMicroversionIdAndConfiguration()
	} else {
		return false
	}
}

// SetMicroversionIdAndConfiguration gets a reference to the given BTMicroversionIdAndConfiguration2338 and assigns it to the MicroversionIdAndConfiguration field.
func (o *BTFullElementIdWithDocument1729) SetMicroversionIdAndConfiguration(v BTMicroversionIdAndConfiguration2338) {
	type getResult interface {
		SetMicroversionIdAndConfiguration(v BTMicroversionIdAndConfiguration2338)
	}

	o.GetActualInstance().(getResult).SetMicroversionIdAndConfiguration(v)
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *BTFullElementIdWithDocument1729) GetTarget() BTMicroversionIdAndConfiguration2338 {
	type getResult interface {
		GetTarget() BTMicroversionIdAndConfiguration2338
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTarget()
	} else {
		var de BTMicroversionIdAndConfiguration2338
		return de
	}
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementIdWithDocument1729) GetTargetOk() (*BTMicroversionIdAndConfiguration2338, bool) {
	type getResult interface {
		GetTargetOk() (*BTMicroversionIdAndConfiguration2338, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTargetOk()
	} else {
		return nil, false
	}
}

// HasTarget returns a boolean if a field has been set.
func (o *BTFullElementIdWithDocument1729) HasTarget() bool {
	type getResult interface {
		HasTarget() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasTarget()
	} else {
		return false
	}
}

// SetTarget gets a reference to the given BTMicroversionIdAndConfiguration2338 and assigns it to the Target field.
func (o *BTFullElementIdWithDocument1729) SetTarget(v BTMicroversionIdAndConfiguration2338) {
	type getResult interface {
		SetTarget(v BTMicroversionIdAndConfiguration2338)
	}

	o.GetActualInstance().(getResult).SetTarget(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTFullElementIdWithDocument1729) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTFullElementIdAndPartId-643'
	if jsonDict["btType"] == "BTFullElementIdAndPartId-643" {
		// try to unmarshal JSON data into BTFullElementIdAndPartId643
		var qr *BTFullElementIdAndPartId643
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFullElementIdWithDocument1729 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFullElementIdWithDocument1729 = nil
			return fmt.Errorf("Failed to unmarshal BTFullElementIdWithDocument1729 as BTFullElementIdAndPartId643: %s", err.Error())
		}
	}

	var qtx *base_BTFullElementIdWithDocument1729
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTFullElementIdWithDocument1729 = qtx
		return nil // data stored in dst.base_BTFullElementIdWithDocument1729, return on the first match
	} else {
		dst.implBTFullElementIdWithDocument1729 = nil
		return fmt.Errorf("Failed to unmarshal BTFullElementIdWithDocument1729 as base_BTFullElementIdWithDocument1729: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTFullElementIdWithDocument1729) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTFullElementIdWithDocument1729) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTFullElementIdWithDocument1729
}

type NullableBTFullElementIdWithDocument1729 struct {
	value *BTFullElementIdWithDocument1729
	isSet bool
}

func (v NullableBTFullElementIdWithDocument1729) Get() *BTFullElementIdWithDocument1729 {
	return v.value
}

func (v *NullableBTFullElementIdWithDocument1729) Set(val *BTFullElementIdWithDocument1729) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFullElementIdWithDocument1729) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFullElementIdWithDocument1729) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFullElementIdWithDocument1729(val *BTFullElementIdWithDocument1729) *NullableBTFullElementIdWithDocument1729 {
	return &NullableBTFullElementIdWithDocument1729{value: val, isSet: true}
}

func (v NullableBTFullElementIdWithDocument1729) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFullElementIdWithDocument1729) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTFullElementIdWithDocument1729 struct {
	BtType                         *string                               `json:"btType,omitempty"`
	Configured                     *bool                                 `json:"configured,omitempty"`
	DocumentId                     *string                               `json:"documentId,omitempty"`
	ElementId                      *string                               `json:"elementId,omitempty"`
	MicroversionId                 *BTMicroversionId366                  `json:"microversionId,omitempty"`
	MicroversionIdAndConfiguration *BTMicroversionIdAndConfiguration2338 `json:"microversionIdAndConfiguration,omitempty"`
	Target                         *BTMicroversionIdAndConfiguration2338 `json:"target,omitempty"`
}

// Newbase_BTFullElementIdWithDocument1729 instantiates a new base_BTFullElementIdWithDocument1729 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTFullElementIdWithDocument1729() *base_BTFullElementIdWithDocument1729 {
	this := base_BTFullElementIdWithDocument1729{}
	return &this
}

// Newbase_BTFullElementIdWithDocument1729WithDefaults instantiates a new base_BTFullElementIdWithDocument1729 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTFullElementIdWithDocument1729WithDefaults() *base_BTFullElementIdWithDocument1729 {
	this := base_BTFullElementIdWithDocument1729{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTFullElementIdWithDocument1729) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFullElementIdWithDocument1729) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTFullElementIdWithDocument1729) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTFullElementIdWithDocument1729) SetBtType(v string) {
	o.BtType = &v
}

// GetConfigured returns the Configured field value if set, zero value otherwise.
func (o *base_BTFullElementIdWithDocument1729) GetConfigured() bool {
	if o == nil || o.Configured == nil {
		var ret bool
		return ret
	}
	return *o.Configured
}

// GetConfiguredOk returns a tuple with the Configured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFullElementIdWithDocument1729) GetConfiguredOk() (*bool, bool) {
	if o == nil || o.Configured == nil {
		return nil, false
	}
	return o.Configured, true
}

// HasConfigured returns a boolean if a field has been set.
func (o *base_BTFullElementIdWithDocument1729) HasConfigured() bool {
	if o != nil && o.Configured != nil {
		return true
	}

	return false
}

// SetConfigured gets a reference to the given bool and assigns it to the Configured field.
func (o *base_BTFullElementIdWithDocument1729) SetConfigured(v bool) {
	o.Configured = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *base_BTFullElementIdWithDocument1729) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFullElementIdWithDocument1729) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *base_BTFullElementIdWithDocument1729) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *base_BTFullElementIdWithDocument1729) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *base_BTFullElementIdWithDocument1729) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFullElementIdWithDocument1729) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *base_BTFullElementIdWithDocument1729) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *base_BTFullElementIdWithDocument1729) SetElementId(v string) {
	o.ElementId = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *base_BTFullElementIdWithDocument1729) GetMicroversionId() BTMicroversionId366 {
	if o == nil || o.MicroversionId == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFullElementIdWithDocument1729) GetMicroversionIdOk() (*BTMicroversionId366, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *base_BTFullElementIdWithDocument1729) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given BTMicroversionId366 and assigns it to the MicroversionId field.
func (o *base_BTFullElementIdWithDocument1729) SetMicroversionId(v BTMicroversionId366) {
	o.MicroversionId = &v
}

// GetMicroversionIdAndConfiguration returns the MicroversionIdAndConfiguration field value if set, zero value otherwise.
func (o *base_BTFullElementIdWithDocument1729) GetMicroversionIdAndConfiguration() BTMicroversionIdAndConfiguration2338 {
	if o == nil || o.MicroversionIdAndConfiguration == nil {
		var ret BTMicroversionIdAndConfiguration2338
		return ret
	}
	return *o.MicroversionIdAndConfiguration
}

// GetMicroversionIdAndConfigurationOk returns a tuple with the MicroversionIdAndConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFullElementIdWithDocument1729) GetMicroversionIdAndConfigurationOk() (*BTMicroversionIdAndConfiguration2338, bool) {
	if o == nil || o.MicroversionIdAndConfiguration == nil {
		return nil, false
	}
	return o.MicroversionIdAndConfiguration, true
}

// HasMicroversionIdAndConfiguration returns a boolean if a field has been set.
func (o *base_BTFullElementIdWithDocument1729) HasMicroversionIdAndConfiguration() bool {
	if o != nil && o.MicroversionIdAndConfiguration != nil {
		return true
	}

	return false
}

// SetMicroversionIdAndConfiguration gets a reference to the given BTMicroversionIdAndConfiguration2338 and assigns it to the MicroversionIdAndConfiguration field.
func (o *base_BTFullElementIdWithDocument1729) SetMicroversionIdAndConfiguration(v BTMicroversionIdAndConfiguration2338) {
	o.MicroversionIdAndConfiguration = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *base_BTFullElementIdWithDocument1729) GetTarget() BTMicroversionIdAndConfiguration2338 {
	if o == nil || o.Target == nil {
		var ret BTMicroversionIdAndConfiguration2338
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFullElementIdWithDocument1729) GetTargetOk() (*BTMicroversionIdAndConfiguration2338, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *base_BTFullElementIdWithDocument1729) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given BTMicroversionIdAndConfiguration2338 and assigns it to the Target field.
func (o *base_BTFullElementIdWithDocument1729) SetTarget(v BTMicroversionIdAndConfiguration2338) {
	o.Target = &v
}

func (o base_BTFullElementIdWithDocument1729) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Configured != nil {
		toSerialize["configured"] = o.Configured
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	if o.MicroversionIdAndConfiguration != nil {
		toSerialize["microversionIdAndConfiguration"] = o.MicroversionIdAndConfiguration
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}
