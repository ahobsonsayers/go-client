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
)

// BTCurveGeometryControlPointSpline2197 struct for BTCurveGeometryControlPointSpline2197
type BTCurveGeometryControlPointSpline2197 struct {
	BtType            *string   `json:"btType,omitempty"`
	ControlPointCount *int32    `json:"controlPointCount,omitempty"`
	ControlPoints     []float64 `json:"controlPoints,omitempty"`
	Degree            *int32    `json:"degree,omitempty"`
	IsPeriodic        *bool     `json:"isPeriodic,omitempty"`
	IsRational        *bool     `json:"isRational,omitempty"`
	Knots             []float64 `json:"knots,omitempty"`
	IsBezier          *bool     `json:"isBezier,omitempty"`
}

// NewBTCurveGeometryControlPointSpline2197 instantiates a new BTCurveGeometryControlPointSpline2197 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCurveGeometryControlPointSpline2197() *BTCurveGeometryControlPointSpline2197 {
	this := BTCurveGeometryControlPointSpline2197{}
	return &this
}

// NewBTCurveGeometryControlPointSpline2197WithDefaults instantiates a new BTCurveGeometryControlPointSpline2197 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCurveGeometryControlPointSpline2197WithDefaults() *BTCurveGeometryControlPointSpline2197 {
	this := BTCurveGeometryControlPointSpline2197{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTCurveGeometryControlPointSpline2197) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryControlPointSpline2197) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTCurveGeometryControlPointSpline2197) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTCurveGeometryControlPointSpline2197) SetBtType(v string) {
	o.BtType = &v
}

// GetControlPointCount returns the ControlPointCount field value if set, zero value otherwise.
func (o *BTCurveGeometryControlPointSpline2197) GetControlPointCount() int32 {
	if o == nil || o.ControlPointCount == nil {
		var ret int32
		return ret
	}
	return *o.ControlPointCount
}

// GetControlPointCountOk returns a tuple with the ControlPointCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryControlPointSpline2197) GetControlPointCountOk() (*int32, bool) {
	if o == nil || o.ControlPointCount == nil {
		return nil, false
	}
	return o.ControlPointCount, true
}

// HasControlPointCount returns a boolean if a field has been set.
func (o *BTCurveGeometryControlPointSpline2197) HasControlPointCount() bool {
	if o != nil && o.ControlPointCount != nil {
		return true
	}

	return false
}

// SetControlPointCount gets a reference to the given int32 and assigns it to the ControlPointCount field.
func (o *BTCurveGeometryControlPointSpline2197) SetControlPointCount(v int32) {
	o.ControlPointCount = &v
}

// GetControlPoints returns the ControlPoints field value if set, zero value otherwise.
func (o *BTCurveGeometryControlPointSpline2197) GetControlPoints() []float64 {
	if o == nil || o.ControlPoints == nil {
		var ret []float64
		return ret
	}
	return o.ControlPoints
}

// GetControlPointsOk returns a tuple with the ControlPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryControlPointSpline2197) GetControlPointsOk() ([]float64, bool) {
	if o == nil || o.ControlPoints == nil {
		return nil, false
	}
	return o.ControlPoints, true
}

// HasControlPoints returns a boolean if a field has been set.
func (o *BTCurveGeometryControlPointSpline2197) HasControlPoints() bool {
	if o != nil && o.ControlPoints != nil {
		return true
	}

	return false
}

// SetControlPoints gets a reference to the given []float64 and assigns it to the ControlPoints field.
func (o *BTCurveGeometryControlPointSpline2197) SetControlPoints(v []float64) {
	o.ControlPoints = v
}

// GetDegree returns the Degree field value if set, zero value otherwise.
func (o *BTCurveGeometryControlPointSpline2197) GetDegree() int32 {
	if o == nil || o.Degree == nil {
		var ret int32
		return ret
	}
	return *o.Degree
}

// GetDegreeOk returns a tuple with the Degree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryControlPointSpline2197) GetDegreeOk() (*int32, bool) {
	if o == nil || o.Degree == nil {
		return nil, false
	}
	return o.Degree, true
}

// HasDegree returns a boolean if a field has been set.
func (o *BTCurveGeometryControlPointSpline2197) HasDegree() bool {
	if o != nil && o.Degree != nil {
		return true
	}

	return false
}

// SetDegree gets a reference to the given int32 and assigns it to the Degree field.
func (o *BTCurveGeometryControlPointSpline2197) SetDegree(v int32) {
	o.Degree = &v
}

// GetIsPeriodic returns the IsPeriodic field value if set, zero value otherwise.
func (o *BTCurveGeometryControlPointSpline2197) GetIsPeriodic() bool {
	if o == nil || o.IsPeriodic == nil {
		var ret bool
		return ret
	}
	return *o.IsPeriodic
}

// GetIsPeriodicOk returns a tuple with the IsPeriodic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryControlPointSpline2197) GetIsPeriodicOk() (*bool, bool) {
	if o == nil || o.IsPeriodic == nil {
		return nil, false
	}
	return o.IsPeriodic, true
}

// HasIsPeriodic returns a boolean if a field has been set.
func (o *BTCurveGeometryControlPointSpline2197) HasIsPeriodic() bool {
	if o != nil && o.IsPeriodic != nil {
		return true
	}

	return false
}

// SetIsPeriodic gets a reference to the given bool and assigns it to the IsPeriodic field.
func (o *BTCurveGeometryControlPointSpline2197) SetIsPeriodic(v bool) {
	o.IsPeriodic = &v
}

// GetIsRational returns the IsRational field value if set, zero value otherwise.
func (o *BTCurveGeometryControlPointSpline2197) GetIsRational() bool {
	if o == nil || o.IsRational == nil {
		var ret bool
		return ret
	}
	return *o.IsRational
}

// GetIsRationalOk returns a tuple with the IsRational field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryControlPointSpline2197) GetIsRationalOk() (*bool, bool) {
	if o == nil || o.IsRational == nil {
		return nil, false
	}
	return o.IsRational, true
}

// HasIsRational returns a boolean if a field has been set.
func (o *BTCurveGeometryControlPointSpline2197) HasIsRational() bool {
	if o != nil && o.IsRational != nil {
		return true
	}

	return false
}

// SetIsRational gets a reference to the given bool and assigns it to the IsRational field.
func (o *BTCurveGeometryControlPointSpline2197) SetIsRational(v bool) {
	o.IsRational = &v
}

// GetKnots returns the Knots field value if set, zero value otherwise.
func (o *BTCurveGeometryControlPointSpline2197) GetKnots() []float64 {
	if o == nil || o.Knots == nil {
		var ret []float64
		return ret
	}
	return o.Knots
}

// GetKnotsOk returns a tuple with the Knots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryControlPointSpline2197) GetKnotsOk() ([]float64, bool) {
	if o == nil || o.Knots == nil {
		return nil, false
	}
	return o.Knots, true
}

// HasKnots returns a boolean if a field has been set.
func (o *BTCurveGeometryControlPointSpline2197) HasKnots() bool {
	if o != nil && o.Knots != nil {
		return true
	}

	return false
}

// SetKnots gets a reference to the given []float64 and assigns it to the Knots field.
func (o *BTCurveGeometryControlPointSpline2197) SetKnots(v []float64) {
	o.Knots = v
}

// GetIsBezier returns the IsBezier field value if set, zero value otherwise.
func (o *BTCurveGeometryControlPointSpline2197) GetIsBezier() bool {
	if o == nil || o.IsBezier == nil {
		var ret bool
		return ret
	}
	return *o.IsBezier
}

// GetIsBezierOk returns a tuple with the IsBezier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryControlPointSpline2197) GetIsBezierOk() (*bool, bool) {
	if o == nil || o.IsBezier == nil {
		return nil, false
	}
	return o.IsBezier, true
}

// HasIsBezier returns a boolean if a field has been set.
func (o *BTCurveGeometryControlPointSpline2197) HasIsBezier() bool {
	if o != nil && o.IsBezier != nil {
		return true
	}

	return false
}

// SetIsBezier gets a reference to the given bool and assigns it to the IsBezier field.
func (o *BTCurveGeometryControlPointSpline2197) SetIsBezier(v bool) {
	o.IsBezier = &v
}

func (o BTCurveGeometryControlPointSpline2197) MarshalJSON() ([]byte, error) {
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
	if o.IsBezier != nil {
		toSerialize["isBezier"] = o.IsBezier
	}
	return json.Marshal(toSerialize)
}

type NullableBTCurveGeometryControlPointSpline2197 struct {
	value *BTCurveGeometryControlPointSpline2197
	isSet bool
}

func (v NullableBTCurveGeometryControlPointSpline2197) Get() *BTCurveGeometryControlPointSpline2197 {
	return v.value
}

func (v *NullableBTCurveGeometryControlPointSpline2197) Set(val *BTCurveGeometryControlPointSpline2197) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCurveGeometryControlPointSpline2197) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCurveGeometryControlPointSpline2197) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCurveGeometryControlPointSpline2197(val *BTCurveGeometryControlPointSpline2197) *NullableBTCurveGeometryControlPointSpline2197 {
	return &NullableBTCurveGeometryControlPointSpline2197{value: val, isSet: true}
}

func (v NullableBTCurveGeometryControlPointSpline2197) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCurveGeometryControlPointSpline2197) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
