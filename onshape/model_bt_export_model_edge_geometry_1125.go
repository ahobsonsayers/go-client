/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6457-6d43038cb4be
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTExportModelEdgeGeometry1125 - struct for BTExportModelEdgeGeometry1125
type BTExportModelEdgeGeometry1125 struct {
	implBTExportModelEdgeGeometry1125 interface{}
}

// BTExportModelArcEdgeGeometry1257AsBTExportModelEdgeGeometry1125 is a convenience function that returns BTExportModelArcEdgeGeometry1257 wrapped in BTExportModelEdgeGeometry1125
func (o *BTExportModelArcEdgeGeometry1257) AsBTExportModelEdgeGeometry1125() *BTExportModelEdgeGeometry1125 {
	return &BTExportModelEdgeGeometry1125{o}
}

// NewBTExportModelEdgeGeometry1125 instantiates a new BTExportModelEdgeGeometry1125 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportModelEdgeGeometry1125() *BTExportModelEdgeGeometry1125 {
	this := BTExportModelEdgeGeometry1125{Newbase_BTExportModelEdgeGeometry1125()}
	return &this
}

// NewBTExportModelEdgeGeometry1125WithDefaults instantiates a new BTExportModelEdgeGeometry1125 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportModelEdgeGeometry1125WithDefaults() *BTExportModelEdgeGeometry1125 {
	this := BTExportModelEdgeGeometry1125{Newbase_BTExportModelEdgeGeometry1125WithDefaults()}
	return &this
}

// GetEndPoint returns the EndPoint field value if set, zero value otherwise.
func (o *BTExportModelEdgeGeometry1125) GetEndPoint() BTVector3d389 {
	type getResult interface {
		GetEndPoint() BTVector3d389
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetEndPoint()
	} else {
		var de BTVector3d389
		return de
	}
}

// GetEndPointOk returns a tuple with the EndPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdgeGeometry1125) GetEndPointOk() (*BTVector3d389, bool) {
	type getResult interface {
		GetEndPointOk() (*BTVector3d389, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetEndPointOk()
	} else {
		return nil, false
	}
}

// HasEndPoint returns a boolean if a field has been set.
func (o *BTExportModelEdgeGeometry1125) HasEndPoint() bool {
	type getResult interface {
		HasEndPoint() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasEndPoint()
	} else {
		return false
	}
}

// SetEndPoint gets a reference to the given BTVector3d389 and assigns it to the EndPoint field.
func (o *BTExportModelEdgeGeometry1125) SetEndPoint(v BTVector3d389) {
	type getResult interface {
		SetEndPoint(v BTVector3d389)
	}

	o.GetActualInstance().(getResult).SetEndPoint(v)
}

// GetEndVector returns the EndVector field value if set, zero value otherwise.
func (o *BTExportModelEdgeGeometry1125) GetEndVector() BTVector3d389 {
	type getResult interface {
		GetEndVector() BTVector3d389
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetEndVector()
	} else {
		var de BTVector3d389
		return de
	}
}

// GetEndVectorOk returns a tuple with the EndVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdgeGeometry1125) GetEndVectorOk() (*BTVector3d389, bool) {
	type getResult interface {
		GetEndVectorOk() (*BTVector3d389, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetEndVectorOk()
	} else {
		return nil, false
	}
}

// HasEndVector returns a boolean if a field has been set.
func (o *BTExportModelEdgeGeometry1125) HasEndVector() bool {
	type getResult interface {
		HasEndVector() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasEndVector()
	} else {
		return false
	}
}

// SetEndVector gets a reference to the given BTVector3d389 and assigns it to the EndVector field.
func (o *BTExportModelEdgeGeometry1125) SetEndVector(v BTVector3d389) {
	type getResult interface {
		SetEndVector(v BTVector3d389)
	}

	o.GetActualInstance().(getResult).SetEndVector(v)
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *BTExportModelEdgeGeometry1125) GetLength() float64 {
	type getResult interface {
		GetLength() float64
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetLength()
	} else {
		var de float64
		return de
	}
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdgeGeometry1125) GetLengthOk() (*float64, bool) {
	type getResult interface {
		GetLengthOk() (*float64, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetLengthOk()
	} else {
		return nil, false
	}
}

// HasLength returns a boolean if a field has been set.
func (o *BTExportModelEdgeGeometry1125) HasLength() bool {
	type getResult interface {
		HasLength() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasLength()
	} else {
		return false
	}
}

// SetLength gets a reference to the given float64 and assigns it to the Length field.
func (o *BTExportModelEdgeGeometry1125) SetLength(v float64) {
	type getResult interface {
		SetLength(v float64)
	}

	o.GetActualInstance().(getResult).SetLength(v)
}

// GetMidPoint returns the MidPoint field value if set, zero value otherwise.
func (o *BTExportModelEdgeGeometry1125) GetMidPoint() BTVector3d389 {
	type getResult interface {
		GetMidPoint() BTVector3d389
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMidPoint()
	} else {
		var de BTVector3d389
		return de
	}
}

// GetMidPointOk returns a tuple with the MidPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdgeGeometry1125) GetMidPointOk() (*BTVector3d389, bool) {
	type getResult interface {
		GetMidPointOk() (*BTVector3d389, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMidPointOk()
	} else {
		return nil, false
	}
}

// HasMidPoint returns a boolean if a field has been set.
func (o *BTExportModelEdgeGeometry1125) HasMidPoint() bool {
	type getResult interface {
		HasMidPoint() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasMidPoint()
	} else {
		return false
	}
}

// SetMidPoint gets a reference to the given BTVector3d389 and assigns it to the MidPoint field.
func (o *BTExportModelEdgeGeometry1125) SetMidPoint(v BTVector3d389) {
	type getResult interface {
		SetMidPoint(v BTVector3d389)
	}

	o.GetActualInstance().(getResult).SetMidPoint(v)
}

// GetQuarterPoint returns the QuarterPoint field value if set, zero value otherwise.
func (o *BTExportModelEdgeGeometry1125) GetQuarterPoint() BTVector3d389 {
	type getResult interface {
		GetQuarterPoint() BTVector3d389
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQuarterPoint()
	} else {
		var de BTVector3d389
		return de
	}
}

// GetQuarterPointOk returns a tuple with the QuarterPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdgeGeometry1125) GetQuarterPointOk() (*BTVector3d389, bool) {
	type getResult interface {
		GetQuarterPointOk() (*BTVector3d389, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQuarterPointOk()
	} else {
		return nil, false
	}
}

// HasQuarterPoint returns a boolean if a field has been set.
func (o *BTExportModelEdgeGeometry1125) HasQuarterPoint() bool {
	type getResult interface {
		HasQuarterPoint() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasQuarterPoint()
	} else {
		return false
	}
}

// SetQuarterPoint gets a reference to the given BTVector3d389 and assigns it to the QuarterPoint field.
func (o *BTExportModelEdgeGeometry1125) SetQuarterPoint(v BTVector3d389) {
	type getResult interface {
		SetQuarterPoint(v BTVector3d389)
	}

	o.GetActualInstance().(getResult).SetQuarterPoint(v)
}

// GetStartPoint returns the StartPoint field value if set, zero value otherwise.
func (o *BTExportModelEdgeGeometry1125) GetStartPoint() BTVector3d389 {
	type getResult interface {
		GetStartPoint() BTVector3d389
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetStartPoint()
	} else {
		var de BTVector3d389
		return de
	}
}

// GetStartPointOk returns a tuple with the StartPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdgeGeometry1125) GetStartPointOk() (*BTVector3d389, bool) {
	type getResult interface {
		GetStartPointOk() (*BTVector3d389, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetStartPointOk()
	} else {
		return nil, false
	}
}

// HasStartPoint returns a boolean if a field has been set.
func (o *BTExportModelEdgeGeometry1125) HasStartPoint() bool {
	type getResult interface {
		HasStartPoint() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasStartPoint()
	} else {
		return false
	}
}

// SetStartPoint gets a reference to the given BTVector3d389 and assigns it to the StartPoint field.
func (o *BTExportModelEdgeGeometry1125) SetStartPoint(v BTVector3d389) {
	type getResult interface {
		SetStartPoint(v BTVector3d389)
	}

	o.GetActualInstance().(getResult).SetStartPoint(v)
}

// GetStartVector returns the StartVector field value if set, zero value otherwise.
func (o *BTExportModelEdgeGeometry1125) GetStartVector() BTVector3d389 {
	type getResult interface {
		GetStartVector() BTVector3d389
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetStartVector()
	} else {
		var de BTVector3d389
		return de
	}
}

// GetStartVectorOk returns a tuple with the StartVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdgeGeometry1125) GetStartVectorOk() (*BTVector3d389, bool) {
	type getResult interface {
		GetStartVectorOk() (*BTVector3d389, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetStartVectorOk()
	} else {
		return nil, false
	}
}

// HasStartVector returns a boolean if a field has been set.
func (o *BTExportModelEdgeGeometry1125) HasStartVector() bool {
	type getResult interface {
		HasStartVector() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasStartVector()
	} else {
		return false
	}
}

// SetStartVector gets a reference to the given BTVector3d389 and assigns it to the StartVector field.
func (o *BTExportModelEdgeGeometry1125) SetStartVector(v BTVector3d389) {
	type getResult interface {
		SetStartVector(v BTVector3d389)
	}

	o.GetActualInstance().(getResult).SetStartVector(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTExportModelEdgeGeometry1125) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTExportModelArcEdgeGeometry-1257'
	if jsonDict["btType"] == "BTExportModelArcEdgeGeometry-1257" {
		// try to unmarshal JSON data into BTExportModelArcEdgeGeometry1257
		var qr *BTExportModelArcEdgeGeometry1257
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTExportModelEdgeGeometry1125 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTExportModelEdgeGeometry1125 = nil
			return fmt.Errorf("Failed to unmarshal BTExportModelEdgeGeometry1125 as BTExportModelArcEdgeGeometry1257: %s", err.Error())
		}
	}

	var qtx *base_BTExportModelEdgeGeometry1125
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTExportModelEdgeGeometry1125 = qtx
		return nil // data stored in dst.base_BTExportModelEdgeGeometry1125, return on the first match
	} else {
		dst.implBTExportModelEdgeGeometry1125 = nil
		return fmt.Errorf("Failed to unmarshal BTExportModelEdgeGeometry1125 as base_BTExportModelEdgeGeometry1125: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTExportModelEdgeGeometry1125) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTExportModelEdgeGeometry1125) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTExportModelEdgeGeometry1125
}

type NullableBTExportModelEdgeGeometry1125 struct {
	value *BTExportModelEdgeGeometry1125
	isSet bool
}

func (v NullableBTExportModelEdgeGeometry1125) Get() *BTExportModelEdgeGeometry1125 {
	return v.value
}

func (v *NullableBTExportModelEdgeGeometry1125) Set(val *BTExportModelEdgeGeometry1125) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportModelEdgeGeometry1125) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportModelEdgeGeometry1125) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportModelEdgeGeometry1125(val *BTExportModelEdgeGeometry1125) *NullableBTExportModelEdgeGeometry1125 {
	return &NullableBTExportModelEdgeGeometry1125{value: val, isSet: true}
}

func (v NullableBTExportModelEdgeGeometry1125) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportModelEdgeGeometry1125) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTExportModelEdgeGeometry1125 struct {
	EndPoint     *BTVector3d389 `json:"endPoint,omitempty"`
	EndVector    *BTVector3d389 `json:"endVector,omitempty"`
	Length       *float64       `json:"length,omitempty"`
	MidPoint     *BTVector3d389 `json:"midPoint,omitempty"`
	QuarterPoint *BTVector3d389 `json:"quarterPoint,omitempty"`
	StartPoint   *BTVector3d389 `json:"startPoint,omitempty"`
	StartVector  *BTVector3d389 `json:"startVector,omitempty"`
}

// Newbase_BTExportModelEdgeGeometry1125 instantiates a new base_BTExportModelEdgeGeometry1125 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTExportModelEdgeGeometry1125() *base_BTExportModelEdgeGeometry1125 {
	this := base_BTExportModelEdgeGeometry1125{}
	return &this
}

// Newbase_BTExportModelEdgeGeometry1125WithDefaults instantiates a new base_BTExportModelEdgeGeometry1125 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTExportModelEdgeGeometry1125WithDefaults() *base_BTExportModelEdgeGeometry1125 {
	this := base_BTExportModelEdgeGeometry1125{}
	return &this
}

// GetEndPoint returns the EndPoint field value if set, zero value otherwise.
func (o *base_BTExportModelEdgeGeometry1125) GetEndPoint() BTVector3d389 {
	if o == nil || o.EndPoint == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.EndPoint
}

// GetEndPointOk returns a tuple with the EndPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTExportModelEdgeGeometry1125) GetEndPointOk() (*BTVector3d389, bool) {
	if o == nil || o.EndPoint == nil {
		return nil, false
	}
	return o.EndPoint, true
}

// HasEndPoint returns a boolean if a field has been set.
func (o *base_BTExportModelEdgeGeometry1125) HasEndPoint() bool {
	if o != nil && o.EndPoint != nil {
		return true
	}

	return false
}

// SetEndPoint gets a reference to the given BTVector3d389 and assigns it to the EndPoint field.
func (o *base_BTExportModelEdgeGeometry1125) SetEndPoint(v BTVector3d389) {
	o.EndPoint = &v
}

// GetEndVector returns the EndVector field value if set, zero value otherwise.
func (o *base_BTExportModelEdgeGeometry1125) GetEndVector() BTVector3d389 {
	if o == nil || o.EndVector == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.EndVector
}

// GetEndVectorOk returns a tuple with the EndVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTExportModelEdgeGeometry1125) GetEndVectorOk() (*BTVector3d389, bool) {
	if o == nil || o.EndVector == nil {
		return nil, false
	}
	return o.EndVector, true
}

// HasEndVector returns a boolean if a field has been set.
func (o *base_BTExportModelEdgeGeometry1125) HasEndVector() bool {
	if o != nil && o.EndVector != nil {
		return true
	}

	return false
}

// SetEndVector gets a reference to the given BTVector3d389 and assigns it to the EndVector field.
func (o *base_BTExportModelEdgeGeometry1125) SetEndVector(v BTVector3d389) {
	o.EndVector = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *base_BTExportModelEdgeGeometry1125) GetLength() float64 {
	if o == nil || o.Length == nil {
		var ret float64
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTExportModelEdgeGeometry1125) GetLengthOk() (*float64, bool) {
	if o == nil || o.Length == nil {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *base_BTExportModelEdgeGeometry1125) HasLength() bool {
	if o != nil && o.Length != nil {
		return true
	}

	return false
}

// SetLength gets a reference to the given float64 and assigns it to the Length field.
func (o *base_BTExportModelEdgeGeometry1125) SetLength(v float64) {
	o.Length = &v
}

// GetMidPoint returns the MidPoint field value if set, zero value otherwise.
func (o *base_BTExportModelEdgeGeometry1125) GetMidPoint() BTVector3d389 {
	if o == nil || o.MidPoint == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.MidPoint
}

// GetMidPointOk returns a tuple with the MidPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTExportModelEdgeGeometry1125) GetMidPointOk() (*BTVector3d389, bool) {
	if o == nil || o.MidPoint == nil {
		return nil, false
	}
	return o.MidPoint, true
}

// HasMidPoint returns a boolean if a field has been set.
func (o *base_BTExportModelEdgeGeometry1125) HasMidPoint() bool {
	if o != nil && o.MidPoint != nil {
		return true
	}

	return false
}

// SetMidPoint gets a reference to the given BTVector3d389 and assigns it to the MidPoint field.
func (o *base_BTExportModelEdgeGeometry1125) SetMidPoint(v BTVector3d389) {
	o.MidPoint = &v
}

// GetQuarterPoint returns the QuarterPoint field value if set, zero value otherwise.
func (o *base_BTExportModelEdgeGeometry1125) GetQuarterPoint() BTVector3d389 {
	if o == nil || o.QuarterPoint == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.QuarterPoint
}

// GetQuarterPointOk returns a tuple with the QuarterPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTExportModelEdgeGeometry1125) GetQuarterPointOk() (*BTVector3d389, bool) {
	if o == nil || o.QuarterPoint == nil {
		return nil, false
	}
	return o.QuarterPoint, true
}

// HasQuarterPoint returns a boolean if a field has been set.
func (o *base_BTExportModelEdgeGeometry1125) HasQuarterPoint() bool {
	if o != nil && o.QuarterPoint != nil {
		return true
	}

	return false
}

// SetQuarterPoint gets a reference to the given BTVector3d389 and assigns it to the QuarterPoint field.
func (o *base_BTExportModelEdgeGeometry1125) SetQuarterPoint(v BTVector3d389) {
	o.QuarterPoint = &v
}

// GetStartPoint returns the StartPoint field value if set, zero value otherwise.
func (o *base_BTExportModelEdgeGeometry1125) GetStartPoint() BTVector3d389 {
	if o == nil || o.StartPoint == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.StartPoint
}

// GetStartPointOk returns a tuple with the StartPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTExportModelEdgeGeometry1125) GetStartPointOk() (*BTVector3d389, bool) {
	if o == nil || o.StartPoint == nil {
		return nil, false
	}
	return o.StartPoint, true
}

// HasStartPoint returns a boolean if a field has been set.
func (o *base_BTExportModelEdgeGeometry1125) HasStartPoint() bool {
	if o != nil && o.StartPoint != nil {
		return true
	}

	return false
}

// SetStartPoint gets a reference to the given BTVector3d389 and assigns it to the StartPoint field.
func (o *base_BTExportModelEdgeGeometry1125) SetStartPoint(v BTVector3d389) {
	o.StartPoint = &v
}

// GetStartVector returns the StartVector field value if set, zero value otherwise.
func (o *base_BTExportModelEdgeGeometry1125) GetStartVector() BTVector3d389 {
	if o == nil || o.StartVector == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.StartVector
}

// GetStartVectorOk returns a tuple with the StartVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTExportModelEdgeGeometry1125) GetStartVectorOk() (*BTVector3d389, bool) {
	if o == nil || o.StartVector == nil {
		return nil, false
	}
	return o.StartVector, true
}

// HasStartVector returns a boolean if a field has been set.
func (o *base_BTExportModelEdgeGeometry1125) HasStartVector() bool {
	if o != nil && o.StartVector != nil {
		return true
	}

	return false
}

// SetStartVector gets a reference to the given BTVector3d389 and assigns it to the StartVector field.
func (o *base_BTExportModelEdgeGeometry1125) SetStartVector(v BTVector3d389) {
	o.StartVector = &v
}

func (o base_BTExportModelEdgeGeometry1125) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EndPoint != nil {
		toSerialize["endPoint"] = o.EndPoint
	}
	if o.EndVector != nil {
		toSerialize["endVector"] = o.EndVector
	}
	if o.Length != nil {
		toSerialize["length"] = o.Length
	}
	if o.MidPoint != nil {
		toSerialize["midPoint"] = o.MidPoint
	}
	if o.QuarterPoint != nil {
		toSerialize["quarterPoint"] = o.QuarterPoint
	}
	if o.StartPoint != nil {
		toSerialize["startPoint"] = o.StartPoint
	}
	if o.StartVector != nil {
		toSerialize["startVector"] = o.StartVector
	}
	return json.Marshal(toSerialize)
}
