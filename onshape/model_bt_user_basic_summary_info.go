/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6290-b55936bc8c5a
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTUserBasicSummaryInfo - struct for BTUserBasicSummaryInfo
type BTUserBasicSummaryInfo struct {
	implBTUserBasicSummaryInfo interface{}
}

// BTUserSummaryInfoAsBTUserBasicSummaryInfo is a convenience function that returns BTUserSummaryInfo wrapped in BTUserBasicSummaryInfo
func (o *BTUserSummaryInfo) AsBTUserBasicSummaryInfo() *BTUserBasicSummaryInfo {
	return &BTUserBasicSummaryInfo{o}
}

// NewBTUserBasicSummaryInfo instantiates a new BTUserBasicSummaryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUserBasicSummaryInfo() *BTUserBasicSummaryInfo {
	this := BTUserBasicSummaryInfo{Newbase_BTUserBasicSummaryInfo()}
	return &this
}

// NewBTUserBasicSummaryInfoWithDefaults instantiates a new BTUserBasicSummaryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUserBasicSummaryInfoWithDefaults() *BTUserBasicSummaryInfo {
	this := BTUserBasicSummaryInfo{Newbase_BTUserBasicSummaryInfoWithDefaults()}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTUserBasicSummaryInfo) GetHref() string {
	type getResult interface {
		GetHref() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetHref()
	} else {
		var de string
		return de
	}
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserBasicSummaryInfo) GetHrefOk() (*string, bool) {
	type getResult interface {
		GetHrefOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetHrefOk()
	} else {
		return nil, false
	}
}

// HasHref returns a boolean if a field has been set.
func (o *BTUserBasicSummaryInfo) HasHref() bool {
	type getResult interface {
		HasHref() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasHref()
	} else {
		return false
	}
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTUserBasicSummaryInfo) SetHref(v string) {
	type getResult interface {
		SetHref(v string)
	}

	o.GetActualInstance().(getResult).SetHref(v)
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTUserBasicSummaryInfo) GetId() string {
	type getResult interface {
		GetId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetId()
	} else {
		var de string
		return de
	}
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserBasicSummaryInfo) GetIdOk() (*string, bool) {
	type getResult interface {
		GetIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIdOk()
	} else {
		return nil, false
	}
}

// HasId returns a boolean if a field has been set.
func (o *BTUserBasicSummaryInfo) HasId() bool {
	type getResult interface {
		HasId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasId()
	} else {
		return false
	}
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTUserBasicSummaryInfo) SetId(v string) {
	type getResult interface {
		SetId(v string)
	}

	o.GetActualInstance().(getResult).SetId(v)
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BTUserBasicSummaryInfo) GetImage() string {
	type getResult interface {
		GetImage() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImage()
	} else {
		var de string
		return de
	}
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserBasicSummaryInfo) GetImageOk() (*string, bool) {
	type getResult interface {
		GetImageOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImageOk()
	} else {
		return nil, false
	}
}

// HasImage returns a boolean if a field has been set.
func (o *BTUserBasicSummaryInfo) HasImage() bool {
	type getResult interface {
		HasImage() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasImage()
	} else {
		return false
	}
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *BTUserBasicSummaryInfo) SetImage(v string) {
	type getResult interface {
		SetImage(v string)
	}

	o.GetActualInstance().(getResult).SetImage(v)
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTUserBasicSummaryInfo) GetName() string {
	type getResult interface {
		GetName() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetName()
	} else {
		var de string
		return de
	}
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserBasicSummaryInfo) GetNameOk() (*string, bool) {
	type getResult interface {
		GetNameOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNameOk()
	} else {
		return nil, false
	}
}

// HasName returns a boolean if a field has been set.
func (o *BTUserBasicSummaryInfo) HasName() bool {
	type getResult interface {
		HasName() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasName()
	} else {
		return false
	}
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTUserBasicSummaryInfo) SetName(v string) {
	type getResult interface {
		SetName(v string)
	}

	o.GetActualInstance().(getResult).SetName(v)
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BTUserBasicSummaryInfo) GetState() int32 {
	type getResult interface {
		GetState() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetState()
	} else {
		var de int32
		return de
	}
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserBasicSummaryInfo) GetStateOk() (*int32, bool) {
	type getResult interface {
		GetStateOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetStateOk()
	} else {
		return nil, false
	}
}

// HasState returns a boolean if a field has been set.
func (o *BTUserBasicSummaryInfo) HasState() bool {
	type getResult interface {
		HasState() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasState()
	} else {
		return false
	}
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *BTUserBasicSummaryInfo) SetState(v int32) {
	type getResult interface {
		SetState(v int32)
	}

	o.GetActualInstance().(getResult).SetState(v)
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTUserBasicSummaryInfo) GetViewRef() string {
	type getResult interface {
		GetViewRef() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetViewRef()
	} else {
		var de string
		return de
	}
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserBasicSummaryInfo) GetViewRefOk() (*string, bool) {
	type getResult interface {
		GetViewRefOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetViewRefOk()
	} else {
		return nil, false
	}
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTUserBasicSummaryInfo) HasViewRef() bool {
	type getResult interface {
		HasViewRef() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasViewRef()
	} else {
		return false
	}
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTUserBasicSummaryInfo) SetViewRef(v string) {
	type getResult interface {
		SetViewRef(v string)
	}

	o.GetActualInstance().(getResult).SetViewRef(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTUserBasicSummaryInfo) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTUserSummaryInfo'
	if jsonDict["jsonType"] == "BTUserSummaryInfo" {
		// try to unmarshal JSON data into BTUserSummaryInfo
		var qr *BTUserSummaryInfo
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTUserBasicSummaryInfo = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTUserBasicSummaryInfo = nil
			return fmt.Errorf("Failed to unmarshal BTUserBasicSummaryInfo as BTUserSummaryInfo: %s", err.Error())
		}
	}

	// check if the discriminator value is 'user-summary'
	if jsonDict["jsonType"] == "user-summary" {
		// try to unmarshal JSON data into BTUserSummaryInfo
		var qr *BTUserSummaryInfo
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTUserBasicSummaryInfo = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTUserBasicSummaryInfo = nil
			return fmt.Errorf("Failed to unmarshal BTUserBasicSummaryInfo as BTUserSummaryInfo: %s", err.Error())
		}
	}

	var qtx *base_BTUserBasicSummaryInfo
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTUserBasicSummaryInfo = qtx
		return nil // data stored in dst.base_BTUserBasicSummaryInfo, return on the first match
	} else {
		dst.implBTUserBasicSummaryInfo = nil
		return fmt.Errorf("Failed to unmarshal BTUserBasicSummaryInfo as base_BTUserBasicSummaryInfo: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTUserBasicSummaryInfo) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTUserBasicSummaryInfo) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTUserBasicSummaryInfo
}

type NullableBTUserBasicSummaryInfo struct {
	value *BTUserBasicSummaryInfo
	isSet bool
}

func (v NullableBTUserBasicSummaryInfo) Get() *BTUserBasicSummaryInfo {
	return v.value
}

func (v *NullableBTUserBasicSummaryInfo) Set(val *BTUserBasicSummaryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUserBasicSummaryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUserBasicSummaryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUserBasicSummaryInfo(val *BTUserBasicSummaryInfo) *NullableBTUserBasicSummaryInfo {
	return &NullableBTUserBasicSummaryInfo{value: val, isSet: true}
}

func (v NullableBTUserBasicSummaryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUserBasicSummaryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTUserBasicSummaryInfo struct {
	Href    *string `json:"href,omitempty"`
	Id      *string `json:"id,omitempty"`
	Image   *string `json:"image,omitempty"`
	Name    *string `json:"name,omitempty"`
	State   *int32  `json:"state,omitempty"`
	ViewRef *string `json:"viewRef,omitempty"`
}

// Newbase_BTUserBasicSummaryInfo instantiates a new base_BTUserBasicSummaryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTUserBasicSummaryInfo() *base_BTUserBasicSummaryInfo {
	this := base_BTUserBasicSummaryInfo{}
	return &this
}

// Newbase_BTUserBasicSummaryInfoWithDefaults instantiates a new base_BTUserBasicSummaryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTUserBasicSummaryInfoWithDefaults() *base_BTUserBasicSummaryInfo {
	this := base_BTUserBasicSummaryInfo{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *base_BTUserBasicSummaryInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTUserBasicSummaryInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *base_BTUserBasicSummaryInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *base_BTUserBasicSummaryInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *base_BTUserBasicSummaryInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTUserBasicSummaryInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *base_BTUserBasicSummaryInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *base_BTUserBasicSummaryInfo) SetId(v string) {
	o.Id = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *base_BTUserBasicSummaryInfo) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTUserBasicSummaryInfo) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *base_BTUserBasicSummaryInfo) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *base_BTUserBasicSummaryInfo) SetImage(v string) {
	o.Image = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *base_BTUserBasicSummaryInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTUserBasicSummaryInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *base_BTUserBasicSummaryInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *base_BTUserBasicSummaryInfo) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *base_BTUserBasicSummaryInfo) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTUserBasicSummaryInfo) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *base_BTUserBasicSummaryInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *base_BTUserBasicSummaryInfo) SetState(v int32) {
	o.State = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *base_BTUserBasicSummaryInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTUserBasicSummaryInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *base_BTUserBasicSummaryInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *base_BTUserBasicSummaryInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o base_BTUserBasicSummaryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}
