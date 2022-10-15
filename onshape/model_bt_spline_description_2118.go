/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.6922-11b87e04d539
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSplineDescription2118 struct for BTSplineDescription2118
type BTSplineDescription2118 struct {
	BtType                    *string        `json:"btType,omitempty"`
	ControlPoints             []float64      `json:"controlPoints,omitempty"`
	Degree                    *int32         `json:"degree,omitempty"`
	Direction                 *BTVector3d389 `json:"direction,omitempty"`
	DirectionOrientedWithFace *BTVector3d389 `json:"directionOrientedWithFace,omitempty"`
	IsPeriodic                *bool          `json:"isPeriodic,omitempty"`
	IsRational                *bool          `json:"isRational,omitempty"`
	Knots                     []float64      `json:"knots,omitempty"`
	Origin                    *BTVector3d389 `json:"origin,omitempty"`
	Type                      *string        `json:"type,omitempty"`
}

// NewBTSplineDescription2118 instantiates a new BTSplineDescription2118 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSplineDescription2118() *BTSplineDescription2118 {
	this := BTSplineDescription2118{}
	return &this
}

// NewBTSplineDescription2118WithDefaults instantiates a new BTSplineDescription2118 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSplineDescription2118WithDefaults() *BTSplineDescription2118 {
	this := BTSplineDescription2118{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSplineDescription2118) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSplineDescription2118) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSplineDescription2118) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSplineDescription2118) SetBtType(v string) {
	o.BtType = &v
}

// GetControlPoints returns the ControlPoints field value if set, zero value otherwise.
func (o *BTSplineDescription2118) GetControlPoints() []float64 {
	if o == nil || o.ControlPoints == nil {
		var ret []float64
		return ret
	}
	return o.ControlPoints
}

// GetControlPointsOk returns a tuple with the ControlPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSplineDescription2118) GetControlPointsOk() ([]float64, bool) {
	if o == nil || o.ControlPoints == nil {
		return nil, false
	}
	return o.ControlPoints, true
}

// HasControlPoints returns a boolean if a field has been set.
func (o *BTSplineDescription2118) HasControlPoints() bool {
	if o != nil && o.ControlPoints != nil {
		return true
	}

	return false
}

// SetControlPoints gets a reference to the given []float64 and assigns it to the ControlPoints field.
func (o *BTSplineDescription2118) SetControlPoints(v []float64) {
	o.ControlPoints = v
}

// GetDegree returns the Degree field value if set, zero value otherwise.
func (o *BTSplineDescription2118) GetDegree() int32 {
	if o == nil || o.Degree == nil {
		var ret int32
		return ret
	}
	return *o.Degree
}

// GetDegreeOk returns a tuple with the Degree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSplineDescription2118) GetDegreeOk() (*int32, bool) {
	if o == nil || o.Degree == nil {
		return nil, false
	}
	return o.Degree, true
}

// HasDegree returns a boolean if a field has been set.
func (o *BTSplineDescription2118) HasDegree() bool {
	if o != nil && o.Degree != nil {
		return true
	}

	return false
}

// SetDegree gets a reference to the given int32 and assigns it to the Degree field.
func (o *BTSplineDescription2118) SetDegree(v int32) {
	o.Degree = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *BTSplineDescription2118) GetDirection() BTVector3d389 {
	if o == nil || o.Direction == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSplineDescription2118) GetDirectionOk() (*BTVector3d389, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *BTSplineDescription2118) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given BTVector3d389 and assigns it to the Direction field.
func (o *BTSplineDescription2118) SetDirection(v BTVector3d389) {
	o.Direction = &v
}

// GetDirectionOrientedWithFace returns the DirectionOrientedWithFace field value if set, zero value otherwise.
func (o *BTSplineDescription2118) GetDirectionOrientedWithFace() BTVector3d389 {
	if o == nil || o.DirectionOrientedWithFace == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.DirectionOrientedWithFace
}

// GetDirectionOrientedWithFaceOk returns a tuple with the DirectionOrientedWithFace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSplineDescription2118) GetDirectionOrientedWithFaceOk() (*BTVector3d389, bool) {
	if o == nil || o.DirectionOrientedWithFace == nil {
		return nil, false
	}
	return o.DirectionOrientedWithFace, true
}

// HasDirectionOrientedWithFace returns a boolean if a field has been set.
func (o *BTSplineDescription2118) HasDirectionOrientedWithFace() bool {
	if o != nil && o.DirectionOrientedWithFace != nil {
		return true
	}

	return false
}

// SetDirectionOrientedWithFace gets a reference to the given BTVector3d389 and assigns it to the DirectionOrientedWithFace field.
func (o *BTSplineDescription2118) SetDirectionOrientedWithFace(v BTVector3d389) {
	o.DirectionOrientedWithFace = &v
}

// GetIsPeriodic returns the IsPeriodic field value if set, zero value otherwise.
func (o *BTSplineDescription2118) GetIsPeriodic() bool {
	if o == nil || o.IsPeriodic == nil {
		var ret bool
		return ret
	}
	return *o.IsPeriodic
}

// GetIsPeriodicOk returns a tuple with the IsPeriodic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSplineDescription2118) GetIsPeriodicOk() (*bool, bool) {
	if o == nil || o.IsPeriodic == nil {
		return nil, false
	}
	return o.IsPeriodic, true
}

// HasIsPeriodic returns a boolean if a field has been set.
func (o *BTSplineDescription2118) HasIsPeriodic() bool {
	if o != nil && o.IsPeriodic != nil {
		return true
	}

	return false
}

// SetIsPeriodic gets a reference to the given bool and assigns it to the IsPeriodic field.
func (o *BTSplineDescription2118) SetIsPeriodic(v bool) {
	o.IsPeriodic = &v
}

// GetIsRational returns the IsRational field value if set, zero value otherwise.
func (o *BTSplineDescription2118) GetIsRational() bool {
	if o == nil || o.IsRational == nil {
		var ret bool
		return ret
	}
	return *o.IsRational
}

// GetIsRationalOk returns a tuple with the IsRational field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSplineDescription2118) GetIsRationalOk() (*bool, bool) {
	if o == nil || o.IsRational == nil {
		return nil, false
	}
	return o.IsRational, true
}

// HasIsRational returns a boolean if a field has been set.
func (o *BTSplineDescription2118) HasIsRational() bool {
	if o != nil && o.IsRational != nil {
		return true
	}

	return false
}

// SetIsRational gets a reference to the given bool and assigns it to the IsRational field.
func (o *BTSplineDescription2118) SetIsRational(v bool) {
	o.IsRational = &v
}

// GetKnots returns the Knots field value if set, zero value otherwise.
func (o *BTSplineDescription2118) GetKnots() []float64 {
	if o == nil || o.Knots == nil {
		var ret []float64
		return ret
	}
	return o.Knots
}

// GetKnotsOk returns a tuple with the Knots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSplineDescription2118) GetKnotsOk() ([]float64, bool) {
	if o == nil || o.Knots == nil {
		return nil, false
	}
	return o.Knots, true
}

// HasKnots returns a boolean if a field has been set.
func (o *BTSplineDescription2118) HasKnots() bool {
	if o != nil && o.Knots != nil {
		return true
	}

	return false
}

// SetKnots gets a reference to the given []float64 and assigns it to the Knots field.
func (o *BTSplineDescription2118) SetKnots(v []float64) {
	o.Knots = v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BTSplineDescription2118) GetOrigin() BTVector3d389 {
	if o == nil || o.Origin == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSplineDescription2118) GetOriginOk() (*BTVector3d389, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *BTSplineDescription2118) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given BTVector3d389 and assigns it to the Origin field.
func (o *BTSplineDescription2118) SetOrigin(v BTVector3d389) {
	o.Origin = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTSplineDescription2118) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSplineDescription2118) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTSplineDescription2118) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTSplineDescription2118) SetType(v string) {
	o.Type = &v
}

func (o BTSplineDescription2118) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ControlPoints != nil {
		toSerialize["controlPoints"] = o.ControlPoints
	}
	if o.Degree != nil {
		toSerialize["degree"] = o.Degree
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.DirectionOrientedWithFace != nil {
		toSerialize["directionOrientedWithFace"] = o.DirectionOrientedWithFace
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
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBTSplineDescription2118 struct {
	value *BTSplineDescription2118
	isSet bool
}

func (v NullableBTSplineDescription2118) Get() *BTSplineDescription2118 {
	return v.value
}

func (v *NullableBTSplineDescription2118) Set(val *BTSplineDescription2118) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSplineDescription2118) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSplineDescription2118) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSplineDescription2118(val *BTSplineDescription2118) *NullableBTSplineDescription2118 {
	return &NullableBTSplineDescription2118{value: val, isSet: true}
}

func (v NullableBTSplineDescription2118) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSplineDescription2118) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
