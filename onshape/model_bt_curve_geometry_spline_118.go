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

// BTCurveGeometrySpline118 - struct for BTCurveGeometrySpline118
type BTCurveGeometrySpline118 struct {
	implBTCurveGeometrySpline118 interface{}
}

// BTCurveGeometryControlPointSpline2197AsBTCurveGeometrySpline118 is a convenience function that returns BTCurveGeometryControlPointSpline2197 wrapped in BTCurveGeometrySpline118
func (o *BTCurveGeometryControlPointSpline2197) AsBTCurveGeometrySpline118() *BTCurveGeometrySpline118 {
	return &BTCurveGeometrySpline118{o}
}

// NewBTCurveGeometrySpline118 instantiates a new BTCurveGeometrySpline118 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCurveGeometrySpline118() *BTCurveGeometrySpline118 {
	this := BTCurveGeometrySpline118{Newbase_BTCurveGeometrySpline118()}
	return &this
}

// NewBTCurveGeometrySpline118WithDefaults instantiates a new BTCurveGeometrySpline118 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCurveGeometrySpline118WithDefaults() *BTCurveGeometrySpline118 {
	this := BTCurveGeometrySpline118{Newbase_BTCurveGeometrySpline118WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTCurveGeometrySpline118) GetBtType() string {
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
func (o *BTCurveGeometrySpline118) GetBtTypeOk() (*string, bool) {
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
func (o *BTCurveGeometrySpline118) HasBtType() bool {
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
func (o *BTCurveGeometrySpline118) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetControlPointCount returns the ControlPointCount field value if set, zero value otherwise.
func (o *BTCurveGeometrySpline118) GetControlPointCount() int32 {
	type getResult interface {
		GetControlPointCount() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetControlPointCount()
	} else {
		var de int32
		return de
	}
}

// GetControlPointCountOk returns a tuple with the ControlPointCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometrySpline118) GetControlPointCountOk() (*int32, bool) {
	type getResult interface {
		GetControlPointCountOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetControlPointCountOk()
	} else {
		return nil, false
	}
}

// HasControlPointCount returns a boolean if a field has been set.
func (o *BTCurveGeometrySpline118) HasControlPointCount() bool {
	type getResult interface {
		HasControlPointCount() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasControlPointCount()
	} else {
		return false
	}
}

// SetControlPointCount gets a reference to the given int32 and assigns it to the ControlPointCount field.
func (o *BTCurveGeometrySpline118) SetControlPointCount(v int32) {
	type getResult interface {
		SetControlPointCount(v int32)
	}

	o.GetActualInstance().(getResult).SetControlPointCount(v)
}

// GetControlPoints returns the ControlPoints field value if set, zero value otherwise.
func (o *BTCurveGeometrySpline118) GetControlPoints() []float64 {
	type getResult interface {
		GetControlPoints() []float64
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetControlPoints()
	} else {
		var de []float64
		return de
	}
}

// GetControlPointsOk returns a tuple with the ControlPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometrySpline118) GetControlPointsOk() ([]float64, bool) {
	type getResult interface {
		GetControlPointsOk() ([]float64, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetControlPointsOk()
	} else {
		return nil, false
	}
}

// HasControlPoints returns a boolean if a field has been set.
func (o *BTCurveGeometrySpline118) HasControlPoints() bool {
	type getResult interface {
		HasControlPoints() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasControlPoints()
	} else {
		return false
	}
}

// SetControlPoints gets a reference to the given []float64 and assigns it to the ControlPoints field.
func (o *BTCurveGeometrySpline118) SetControlPoints(v []float64) {
	type getResult interface {
		SetControlPoints(v []float64)
	}

	o.GetActualInstance().(getResult).SetControlPoints(v)
}

// GetDegree returns the Degree field value if set, zero value otherwise.
func (o *BTCurveGeometrySpline118) GetDegree() int32 {
	type getResult interface {
		GetDegree() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDegree()
	} else {
		var de int32
		return de
	}
}

// GetDegreeOk returns a tuple with the Degree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometrySpline118) GetDegreeOk() (*int32, bool) {
	type getResult interface {
		GetDegreeOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDegreeOk()
	} else {
		return nil, false
	}
}

// HasDegree returns a boolean if a field has been set.
func (o *BTCurveGeometrySpline118) HasDegree() bool {
	type getResult interface {
		HasDegree() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDegree()
	} else {
		return false
	}
}

// SetDegree gets a reference to the given int32 and assigns it to the Degree field.
func (o *BTCurveGeometrySpline118) SetDegree(v int32) {
	type getResult interface {
		SetDegree(v int32)
	}

	o.GetActualInstance().(getResult).SetDegree(v)
}

// GetIsPeriodic returns the IsPeriodic field value if set, zero value otherwise.
func (o *BTCurveGeometrySpline118) GetIsPeriodic() bool {
	type getResult interface {
		GetIsPeriodic() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIsPeriodic()
	} else {
		var de bool
		return de
	}
}

// GetIsPeriodicOk returns a tuple with the IsPeriodic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometrySpline118) GetIsPeriodicOk() (*bool, bool) {
	type getResult interface {
		GetIsPeriodicOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIsPeriodicOk()
	} else {
		return nil, false
	}
}

// HasIsPeriodic returns a boolean if a field has been set.
func (o *BTCurveGeometrySpline118) HasIsPeriodic() bool {
	type getResult interface {
		HasIsPeriodic() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasIsPeriodic()
	} else {
		return false
	}
}

// SetIsPeriodic gets a reference to the given bool and assigns it to the IsPeriodic field.
func (o *BTCurveGeometrySpline118) SetIsPeriodic(v bool) {
	type getResult interface {
		SetIsPeriodic(v bool)
	}

	o.GetActualInstance().(getResult).SetIsPeriodic(v)
}

// GetIsRational returns the IsRational field value if set, zero value otherwise.
func (o *BTCurveGeometrySpline118) GetIsRational() bool {
	type getResult interface {
		GetIsRational() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIsRational()
	} else {
		var de bool
		return de
	}
}

// GetIsRationalOk returns a tuple with the IsRational field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometrySpline118) GetIsRationalOk() (*bool, bool) {
	type getResult interface {
		GetIsRationalOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIsRationalOk()
	} else {
		return nil, false
	}
}

// HasIsRational returns a boolean if a field has been set.
func (o *BTCurveGeometrySpline118) HasIsRational() bool {
	type getResult interface {
		HasIsRational() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasIsRational()
	} else {
		return false
	}
}

// SetIsRational gets a reference to the given bool and assigns it to the IsRational field.
func (o *BTCurveGeometrySpline118) SetIsRational(v bool) {
	type getResult interface {
		SetIsRational(v bool)
	}

	o.GetActualInstance().(getResult).SetIsRational(v)
}

// GetKnots returns the Knots field value if set, zero value otherwise.
func (o *BTCurveGeometrySpline118) GetKnots() []float64 {
	type getResult interface {
		GetKnots() []float64
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetKnots()
	} else {
		var de []float64
		return de
	}
}

// GetKnotsOk returns a tuple with the Knots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometrySpline118) GetKnotsOk() ([]float64, bool) {
	type getResult interface {
		GetKnotsOk() ([]float64, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetKnotsOk()
	} else {
		return nil, false
	}
}

// HasKnots returns a boolean if a field has been set.
func (o *BTCurveGeometrySpline118) HasKnots() bool {
	type getResult interface {
		HasKnots() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasKnots()
	} else {
		return false
	}
}

// SetKnots gets a reference to the given []float64 and assigns it to the Knots field.
func (o *BTCurveGeometrySpline118) SetKnots(v []float64) {
	type getResult interface {
		SetKnots(v []float64)
	}

	o.GetActualInstance().(getResult).SetKnots(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTCurveGeometrySpline118) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTCurveGeometryControlPointSpline-2197'
	if jsonDict["btType"] == "BTCurveGeometryControlPointSpline-2197" {
		// try to unmarshal JSON data into BTCurveGeometryControlPointSpline2197
		var qr *BTCurveGeometryControlPointSpline2197
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTCurveGeometrySpline118 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTCurveGeometrySpline118 = nil
			return fmt.Errorf("Failed to unmarshal BTCurveGeometrySpline118 as BTCurveGeometryControlPointSpline2197: %s", err.Error())
		}
	}

	var qtx *base_BTCurveGeometrySpline118
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTCurveGeometrySpline118 = qtx
		return nil // data stored in dst.base_BTCurveGeometrySpline118, return on the first match
	} else {
		dst.implBTCurveGeometrySpline118 = nil
		return fmt.Errorf("Failed to unmarshal BTCurveGeometrySpline118 as base_BTCurveGeometrySpline118: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTCurveGeometrySpline118) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTCurveGeometrySpline118) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTCurveGeometrySpline118
}

type NullableBTCurveGeometrySpline118 struct {
	value *BTCurveGeometrySpline118
	isSet bool
}

func (v NullableBTCurveGeometrySpline118) Get() *BTCurveGeometrySpline118 {
	return v.value
}

func (v *NullableBTCurveGeometrySpline118) Set(val *BTCurveGeometrySpline118) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCurveGeometrySpline118) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCurveGeometrySpline118) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCurveGeometrySpline118(val *BTCurveGeometrySpline118) *NullableBTCurveGeometrySpline118 {
	return &NullableBTCurveGeometrySpline118{value: val, isSet: true}
}

func (v NullableBTCurveGeometrySpline118) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCurveGeometrySpline118) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTCurveGeometrySpline118 struct {
	BtType            *string   `json:"btType,omitempty"`
	ControlPointCount *int32    `json:"controlPointCount,omitempty"`
	ControlPoints     []float64 `json:"controlPoints,omitempty"`
	Degree            *int32    `json:"degree,omitempty"`
	IsPeriodic        *bool     `json:"isPeriodic,omitempty"`
	IsRational        *bool     `json:"isRational,omitempty"`
	Knots             []float64 `json:"knots,omitempty"`
}

// Newbase_BTCurveGeometrySpline118 instantiates a new base_BTCurveGeometrySpline118 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTCurveGeometrySpline118() *base_BTCurveGeometrySpline118 {
	this := base_BTCurveGeometrySpline118{}
	return &this
}

// Newbase_BTCurveGeometrySpline118WithDefaults instantiates a new base_BTCurveGeometrySpline118 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTCurveGeometrySpline118WithDefaults() *base_BTCurveGeometrySpline118 {
	this := base_BTCurveGeometrySpline118{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTCurveGeometrySpline118) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveGeometrySpline118) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTCurveGeometrySpline118) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTCurveGeometrySpline118) SetBtType(v string) {
	o.BtType = &v
}

// GetControlPointCount returns the ControlPointCount field value if set, zero value otherwise.
func (o *base_BTCurveGeometrySpline118) GetControlPointCount() int32 {
	if o == nil || o.ControlPointCount == nil {
		var ret int32
		return ret
	}
	return *o.ControlPointCount
}

// GetControlPointCountOk returns a tuple with the ControlPointCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveGeometrySpline118) GetControlPointCountOk() (*int32, bool) {
	if o == nil || o.ControlPointCount == nil {
		return nil, false
	}
	return o.ControlPointCount, true
}

// HasControlPointCount returns a boolean if a field has been set.
func (o *base_BTCurveGeometrySpline118) HasControlPointCount() bool {
	if o != nil && o.ControlPointCount != nil {
		return true
	}

	return false
}

// SetControlPointCount gets a reference to the given int32 and assigns it to the ControlPointCount field.
func (o *base_BTCurveGeometrySpline118) SetControlPointCount(v int32) {
	o.ControlPointCount = &v
}

// GetControlPoints returns the ControlPoints field value if set, zero value otherwise.
func (o *base_BTCurveGeometrySpline118) GetControlPoints() []float64 {
	if o == nil || o.ControlPoints == nil {
		var ret []float64
		return ret
	}
	return o.ControlPoints
}

// GetControlPointsOk returns a tuple with the ControlPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveGeometrySpline118) GetControlPointsOk() ([]float64, bool) {
	if o == nil || o.ControlPoints == nil {
		return nil, false
	}
	return o.ControlPoints, true
}

// HasControlPoints returns a boolean if a field has been set.
func (o *base_BTCurveGeometrySpline118) HasControlPoints() bool {
	if o != nil && o.ControlPoints != nil {
		return true
	}

	return false
}

// SetControlPoints gets a reference to the given []float64 and assigns it to the ControlPoints field.
func (o *base_BTCurveGeometrySpline118) SetControlPoints(v []float64) {
	o.ControlPoints = v
}

// GetDegree returns the Degree field value if set, zero value otherwise.
func (o *base_BTCurveGeometrySpline118) GetDegree() int32 {
	if o == nil || o.Degree == nil {
		var ret int32
		return ret
	}
	return *o.Degree
}

// GetDegreeOk returns a tuple with the Degree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveGeometrySpline118) GetDegreeOk() (*int32, bool) {
	if o == nil || o.Degree == nil {
		return nil, false
	}
	return o.Degree, true
}

// HasDegree returns a boolean if a field has been set.
func (o *base_BTCurveGeometrySpline118) HasDegree() bool {
	if o != nil && o.Degree != nil {
		return true
	}

	return false
}

// SetDegree gets a reference to the given int32 and assigns it to the Degree field.
func (o *base_BTCurveGeometrySpline118) SetDegree(v int32) {
	o.Degree = &v
}

// GetIsPeriodic returns the IsPeriodic field value if set, zero value otherwise.
func (o *base_BTCurveGeometrySpline118) GetIsPeriodic() bool {
	if o == nil || o.IsPeriodic == nil {
		var ret bool
		return ret
	}
	return *o.IsPeriodic
}

// GetIsPeriodicOk returns a tuple with the IsPeriodic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveGeometrySpline118) GetIsPeriodicOk() (*bool, bool) {
	if o == nil || o.IsPeriodic == nil {
		return nil, false
	}
	return o.IsPeriodic, true
}

// HasIsPeriodic returns a boolean if a field has been set.
func (o *base_BTCurveGeometrySpline118) HasIsPeriodic() bool {
	if o != nil && o.IsPeriodic != nil {
		return true
	}

	return false
}

// SetIsPeriodic gets a reference to the given bool and assigns it to the IsPeriodic field.
func (o *base_BTCurveGeometrySpline118) SetIsPeriodic(v bool) {
	o.IsPeriodic = &v
}

// GetIsRational returns the IsRational field value if set, zero value otherwise.
func (o *base_BTCurveGeometrySpline118) GetIsRational() bool {
	if o == nil || o.IsRational == nil {
		var ret bool
		return ret
	}
	return *o.IsRational
}

// GetIsRationalOk returns a tuple with the IsRational field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveGeometrySpline118) GetIsRationalOk() (*bool, bool) {
	if o == nil || o.IsRational == nil {
		return nil, false
	}
	return o.IsRational, true
}

// HasIsRational returns a boolean if a field has been set.
func (o *base_BTCurveGeometrySpline118) HasIsRational() bool {
	if o != nil && o.IsRational != nil {
		return true
	}

	return false
}

// SetIsRational gets a reference to the given bool and assigns it to the IsRational field.
func (o *base_BTCurveGeometrySpline118) SetIsRational(v bool) {
	o.IsRational = &v
}

// GetKnots returns the Knots field value if set, zero value otherwise.
func (o *base_BTCurveGeometrySpline118) GetKnots() []float64 {
	if o == nil || o.Knots == nil {
		var ret []float64
		return ret
	}
	return o.Knots
}

// GetKnotsOk returns a tuple with the Knots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveGeometrySpline118) GetKnotsOk() ([]float64, bool) {
	if o == nil || o.Knots == nil {
		return nil, false
	}
	return o.Knots, true
}

// HasKnots returns a boolean if a field has been set.
func (o *base_BTCurveGeometrySpline118) HasKnots() bool {
	if o != nil && o.Knots != nil {
		return true
	}

	return false
}

// SetKnots gets a reference to the given []float64 and assigns it to the Knots field.
func (o *base_BTCurveGeometrySpline118) SetKnots(v []float64) {
	o.Knots = v
}

func (o base_BTCurveGeometrySpline118) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ControlPointCount != nil {
		toSerialize["controlPointCount"] = o.ControlPointCount
	}
	if o.ControlPoints != nil {
		toSerialize["controlPoints"] = o.ControlPoints
	}
	if o.Degree != nil {
		toSerialize["degree"] = o.Degree
	}
	if o.IsPeriodic != nil {
		toSerialize["isPeriodic"] = o.IsPeriodic
	}
	if o.IsRational != nil {
		toSerialize["isRational"] = o.IsRational
	}
	if o.Knots != nil {
		toSerialize["knots"] = o.Knots
	}
	return json.Marshal(toSerialize)
}
