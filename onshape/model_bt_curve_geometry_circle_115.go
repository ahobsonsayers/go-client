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

// BTCurveGeometryCircle115 - struct for BTCurveGeometryCircle115
type BTCurveGeometryCircle115 struct {
	implBTCurveGeometryCircle115 interface{}
}

// BTCurveGeometryEllipse1189AsBTCurveGeometryCircle115 is a convenience function that returns BTCurveGeometryEllipse1189 wrapped in BTCurveGeometryCircle115
func (o *BTCurveGeometryEllipse1189) AsBTCurveGeometryCircle115() *BTCurveGeometryCircle115 {
	return &BTCurveGeometryCircle115{o}
}

// NewBTCurveGeometryCircle115 instantiates a new BTCurveGeometryCircle115 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCurveGeometryCircle115() *BTCurveGeometryCircle115 {
	this := BTCurveGeometryCircle115{Newbase_BTCurveGeometryCircle115()}
	return &this
}

// NewBTCurveGeometryCircle115WithDefaults instantiates a new BTCurveGeometryCircle115 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCurveGeometryCircle115WithDefaults() *BTCurveGeometryCircle115 {
	this := BTCurveGeometryCircle115{Newbase_BTCurveGeometryCircle115WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTCurveGeometryCircle115) GetBtType() string {
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
func (o *BTCurveGeometryCircle115) GetBtTypeOk() (*string, bool) {
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
func (o *BTCurveGeometryCircle115) HasBtType() bool {
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
func (o *BTCurveGeometryCircle115) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetClockwise returns the Clockwise field value if set, zero value otherwise.
func (o *BTCurveGeometryCircle115) GetClockwise() bool {
	type getResult interface {
		GetClockwise() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetClockwise()
	} else {
		var de bool
		return de
	}
}

// GetClockwiseOk returns a tuple with the Clockwise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryCircle115) GetClockwiseOk() (*bool, bool) {
	type getResult interface {
		GetClockwiseOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetClockwiseOk()
	} else {
		return nil, false
	}
}

// HasClockwise returns a boolean if a field has been set.
func (o *BTCurveGeometryCircle115) HasClockwise() bool {
	type getResult interface {
		HasClockwise() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasClockwise()
	} else {
		return false
	}
}

// SetClockwise gets a reference to the given bool and assigns it to the Clockwise field.
func (o *BTCurveGeometryCircle115) SetClockwise(v bool) {
	type getResult interface {
		SetClockwise(v bool)
	}

	o.GetActualInstance().(getResult).SetClockwise(v)
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *BTCurveGeometryCircle115) GetRadius() float64 {
	type getResult interface {
		GetRadius() float64
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRadius()
	} else {
		var de float64
		return de
	}
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryCircle115) GetRadiusOk() (*float64, bool) {
	type getResult interface {
		GetRadiusOk() (*float64, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRadiusOk()
	} else {
		return nil, false
	}
}

// HasRadius returns a boolean if a field has been set.
func (o *BTCurveGeometryCircle115) HasRadius() bool {
	type getResult interface {
		HasRadius() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasRadius()
	} else {
		return false
	}
}

// SetRadius gets a reference to the given float64 and assigns it to the Radius field.
func (o *BTCurveGeometryCircle115) SetRadius(v float64) {
	type getResult interface {
		SetRadius(v float64)
	}

	o.GetActualInstance().(getResult).SetRadius(v)
}

// GetXcenter returns the Xcenter field value if set, zero value otherwise.
func (o *BTCurveGeometryCircle115) GetXcenter() float64 {
	type getResult interface {
		GetXcenter() float64
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetXcenter()
	} else {
		var de float64
		return de
	}
}

// GetXcenterOk returns a tuple with the Xcenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryCircle115) GetXcenterOk() (*float64, bool) {
	type getResult interface {
		GetXcenterOk() (*float64, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetXcenterOk()
	} else {
		return nil, false
	}
}

// HasXcenter returns a boolean if a field has been set.
func (o *BTCurveGeometryCircle115) HasXcenter() bool {
	type getResult interface {
		HasXcenter() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasXcenter()
	} else {
		return false
	}
}

// SetXcenter gets a reference to the given float64 and assigns it to the Xcenter field.
func (o *BTCurveGeometryCircle115) SetXcenter(v float64) {
	type getResult interface {
		SetXcenter(v float64)
	}

	o.GetActualInstance().(getResult).SetXcenter(v)
}

// GetXdir returns the Xdir field value if set, zero value otherwise.
func (o *BTCurveGeometryCircle115) GetXdir() float64 {
	type getResult interface {
		GetXdir() float64
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetXdir()
	} else {
		var de float64
		return de
	}
}

// GetXdirOk returns a tuple with the Xdir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryCircle115) GetXdirOk() (*float64, bool) {
	type getResult interface {
		GetXdirOk() (*float64, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetXdirOk()
	} else {
		return nil, false
	}
}

// HasXdir returns a boolean if a field has been set.
func (o *BTCurveGeometryCircle115) HasXdir() bool {
	type getResult interface {
		HasXdir() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasXdir()
	} else {
		return false
	}
}

// SetXdir gets a reference to the given float64 and assigns it to the Xdir field.
func (o *BTCurveGeometryCircle115) SetXdir(v float64) {
	type getResult interface {
		SetXdir(v float64)
	}

	o.GetActualInstance().(getResult).SetXdir(v)
}

// GetYcenter returns the Ycenter field value if set, zero value otherwise.
func (o *BTCurveGeometryCircle115) GetYcenter() float64 {
	type getResult interface {
		GetYcenter() float64
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetYcenter()
	} else {
		var de float64
		return de
	}
}

// GetYcenterOk returns a tuple with the Ycenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryCircle115) GetYcenterOk() (*float64, bool) {
	type getResult interface {
		GetYcenterOk() (*float64, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetYcenterOk()
	} else {
		return nil, false
	}
}

// HasYcenter returns a boolean if a field has been set.
func (o *BTCurveGeometryCircle115) HasYcenter() bool {
	type getResult interface {
		HasYcenter() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasYcenter()
	} else {
		return false
	}
}

// SetYcenter gets a reference to the given float64 and assigns it to the Ycenter field.
func (o *BTCurveGeometryCircle115) SetYcenter(v float64) {
	type getResult interface {
		SetYcenter(v float64)
	}

	o.GetActualInstance().(getResult).SetYcenter(v)
}

// GetYdir returns the Ydir field value if set, zero value otherwise.
func (o *BTCurveGeometryCircle115) GetYdir() float64 {
	type getResult interface {
		GetYdir() float64
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetYdir()
	} else {
		var de float64
		return de
	}
}

// GetYdirOk returns a tuple with the Ydir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryCircle115) GetYdirOk() (*float64, bool) {
	type getResult interface {
		GetYdirOk() (*float64, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetYdirOk()
	} else {
		return nil, false
	}
}

// HasYdir returns a boolean if a field has been set.
func (o *BTCurveGeometryCircle115) HasYdir() bool {
	type getResult interface {
		HasYdir() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasYdir()
	} else {
		return false
	}
}

// SetYdir gets a reference to the given float64 and assigns it to the Ydir field.
func (o *BTCurveGeometryCircle115) SetYdir(v float64) {
	type getResult interface {
		SetYdir(v float64)
	}

	o.GetActualInstance().(getResult).SetYdir(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTCurveGeometryCircle115) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTCurveGeometryEllipse-1189'
	if jsonDict["btType"] == "BTCurveGeometryEllipse-1189" {
		// try to unmarshal JSON data into BTCurveGeometryEllipse1189
		var qr *BTCurveGeometryEllipse1189
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTCurveGeometryCircle115 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTCurveGeometryCircle115 = nil
			return fmt.Errorf("Failed to unmarshal BTCurveGeometryCircle115 as BTCurveGeometryEllipse1189: %s", err.Error())
		}
	}

	var qtx *base_BTCurveGeometryCircle115
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTCurveGeometryCircle115 = qtx
		return nil // data stored in dst.base_BTCurveGeometryCircle115, return on the first match
	} else {
		dst.implBTCurveGeometryCircle115 = nil
		return fmt.Errorf("Failed to unmarshal BTCurveGeometryCircle115 as base_BTCurveGeometryCircle115: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTCurveGeometryCircle115) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTCurveGeometryCircle115) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTCurveGeometryCircle115
}

type NullableBTCurveGeometryCircle115 struct {
	value *BTCurveGeometryCircle115
	isSet bool
}

func (v NullableBTCurveGeometryCircle115) Get() *BTCurveGeometryCircle115 {
	return v.value
}

func (v *NullableBTCurveGeometryCircle115) Set(val *BTCurveGeometryCircle115) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCurveGeometryCircle115) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCurveGeometryCircle115) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCurveGeometryCircle115(val *BTCurveGeometryCircle115) *NullableBTCurveGeometryCircle115 {
	return &NullableBTCurveGeometryCircle115{value: val, isSet: true}
}

func (v NullableBTCurveGeometryCircle115) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCurveGeometryCircle115) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTCurveGeometryCircle115 struct {
	BtType    *string  `json:"btType,omitempty"`
	Clockwise *bool    `json:"clockwise,omitempty"`
	Radius    *float64 `json:"radius,omitempty"`
	Xcenter   *float64 `json:"xcenter,omitempty"`
	Xdir      *float64 `json:"xdir,omitempty"`
	Ycenter   *float64 `json:"ycenter,omitempty"`
	Ydir      *float64 `json:"ydir,omitempty"`
}

// Newbase_BTCurveGeometryCircle115 instantiates a new base_BTCurveGeometryCircle115 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTCurveGeometryCircle115() *base_BTCurveGeometryCircle115 {
	this := base_BTCurveGeometryCircle115{}
	return &this
}

// Newbase_BTCurveGeometryCircle115WithDefaults instantiates a new base_BTCurveGeometryCircle115 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTCurveGeometryCircle115WithDefaults() *base_BTCurveGeometryCircle115 {
	this := base_BTCurveGeometryCircle115{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTCurveGeometryCircle115) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveGeometryCircle115) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTCurveGeometryCircle115) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTCurveGeometryCircle115) SetBtType(v string) {
	o.BtType = &v
}

// GetClockwise returns the Clockwise field value if set, zero value otherwise.
func (o *base_BTCurveGeometryCircle115) GetClockwise() bool {
	if o == nil || o.Clockwise == nil {
		var ret bool
		return ret
	}
	return *o.Clockwise
}

// GetClockwiseOk returns a tuple with the Clockwise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveGeometryCircle115) GetClockwiseOk() (*bool, bool) {
	if o == nil || o.Clockwise == nil {
		return nil, false
	}
	return o.Clockwise, true
}

// HasClockwise returns a boolean if a field has been set.
func (o *base_BTCurveGeometryCircle115) HasClockwise() bool {
	if o != nil && o.Clockwise != nil {
		return true
	}

	return false
}

// SetClockwise gets a reference to the given bool and assigns it to the Clockwise field.
func (o *base_BTCurveGeometryCircle115) SetClockwise(v bool) {
	o.Clockwise = &v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *base_BTCurveGeometryCircle115) GetRadius() float64 {
	if o == nil || o.Radius == nil {
		var ret float64
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveGeometryCircle115) GetRadiusOk() (*float64, bool) {
	if o == nil || o.Radius == nil {
		return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *base_BTCurveGeometryCircle115) HasRadius() bool {
	if o != nil && o.Radius != nil {
		return true
	}

	return false
}

// SetRadius gets a reference to the given float64 and assigns it to the Radius field.
func (o *base_BTCurveGeometryCircle115) SetRadius(v float64) {
	o.Radius = &v
}

// GetXcenter returns the Xcenter field value if set, zero value otherwise.
func (o *base_BTCurveGeometryCircle115) GetXcenter() float64 {
	if o == nil || o.Xcenter == nil {
		var ret float64
		return ret
	}
	return *o.Xcenter
}

// GetXcenterOk returns a tuple with the Xcenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveGeometryCircle115) GetXcenterOk() (*float64, bool) {
	if o == nil || o.Xcenter == nil {
		return nil, false
	}
	return o.Xcenter, true
}

// HasXcenter returns a boolean if a field has been set.
func (o *base_BTCurveGeometryCircle115) HasXcenter() bool {
	if o != nil && o.Xcenter != nil {
		return true
	}

	return false
}

// SetXcenter gets a reference to the given float64 and assigns it to the Xcenter field.
func (o *base_BTCurveGeometryCircle115) SetXcenter(v float64) {
	o.Xcenter = &v
}

// GetXdir returns the Xdir field value if set, zero value otherwise.
func (o *base_BTCurveGeometryCircle115) GetXdir() float64 {
	if o == nil || o.Xdir == nil {
		var ret float64
		return ret
	}
	return *o.Xdir
}

// GetXdirOk returns a tuple with the Xdir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveGeometryCircle115) GetXdirOk() (*float64, bool) {
	if o == nil || o.Xdir == nil {
		return nil, false
	}
	return o.Xdir, true
}

// HasXdir returns a boolean if a field has been set.
func (o *base_BTCurveGeometryCircle115) HasXdir() bool {
	if o != nil && o.Xdir != nil {
		return true
	}

	return false
}

// SetXdir gets a reference to the given float64 and assigns it to the Xdir field.
func (o *base_BTCurveGeometryCircle115) SetXdir(v float64) {
	o.Xdir = &v
}

// GetYcenter returns the Ycenter field value if set, zero value otherwise.
func (o *base_BTCurveGeometryCircle115) GetYcenter() float64 {
	if o == nil || o.Ycenter == nil {
		var ret float64
		return ret
	}
	return *o.Ycenter
}

// GetYcenterOk returns a tuple with the Ycenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveGeometryCircle115) GetYcenterOk() (*float64, bool) {
	if o == nil || o.Ycenter == nil {
		return nil, false
	}
	return o.Ycenter, true
}

// HasYcenter returns a boolean if a field has been set.
func (o *base_BTCurveGeometryCircle115) HasYcenter() bool {
	if o != nil && o.Ycenter != nil {
		return true
	}

	return false
}

// SetYcenter gets a reference to the given float64 and assigns it to the Ycenter field.
func (o *base_BTCurveGeometryCircle115) SetYcenter(v float64) {
	o.Ycenter = &v
}

// GetYdir returns the Ydir field value if set, zero value otherwise.
func (o *base_BTCurveGeometryCircle115) GetYdir() float64 {
	if o == nil || o.Ydir == nil {
		var ret float64
		return ret
	}
	return *o.Ydir
}

// GetYdirOk returns a tuple with the Ydir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveGeometryCircle115) GetYdirOk() (*float64, bool) {
	if o == nil || o.Ydir == nil {
		return nil, false
	}
	return o.Ydir, true
}

// HasYdir returns a boolean if a field has been set.
func (o *base_BTCurveGeometryCircle115) HasYdir() bool {
	if o != nil && o.Ydir != nil {
		return true
	}

	return false
}

// SetYdir gets a reference to the given float64 and assigns it to the Ydir field.
func (o *base_BTCurveGeometryCircle115) SetYdir(v float64) {
	o.Ydir = &v
}

func (o base_BTCurveGeometryCircle115) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Clockwise != nil {
		toSerialize["clockwise"] = o.Clockwise
	}
	if o.Radius != nil {
		toSerialize["radius"] = o.Radius
	}
	if o.Xcenter != nil {
		toSerialize["xcenter"] = o.Xcenter
	}
	if o.Xdir != nil {
		toSerialize["xdir"] = o.Xdir
	}
	if o.Ycenter != nil {
		toSerialize["ycenter"] = o.Ycenter
	}
	if o.Ydir != nil {
		toSerialize["ydir"] = o.Ydir
	}
	return json.Marshal(toSerialize)
}
