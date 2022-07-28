/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5750-4f2542599dd4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCurveGeometryInterpolatedSpline116 struct for BTCurveGeometryInterpolatedSpline116
type BTCurveGeometryInterpolatedSpline116 struct {
	BtType              *string               `json:"btType,omitempty"`
	Derivatives         *map[string][]float64 `json:"derivatives,omitempty"`
	EndDerivativeX      *float64              `json:"endDerivativeX,omitempty"`
	EndDerivativeY      *float64              `json:"endDerivativeY,omitempty"`
	EndHandleX          *float64              `json:"endHandleX,omitempty"`
	EndHandleY          *float64              `json:"endHandleY,omitempty"`
	InterpolationPoints []float64             `json:"interpolationPoints,omitempty"`
	IsPeriodic          *bool                 `json:"isPeriodic,omitempty"`
	StartDerivativeX    *float64              `json:"startDerivativeX,omitempty"`
	StartDerivativeY    *float64              `json:"startDerivativeY,omitempty"`
	StartHandleX        *float64              `json:"startHandleX,omitempty"`
	StartHandleY        *float64              `json:"startHandleY,omitempty"`
}

// NewBTCurveGeometryInterpolatedSpline116 instantiates a new BTCurveGeometryInterpolatedSpline116 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCurveGeometryInterpolatedSpline116() *BTCurveGeometryInterpolatedSpline116 {
	this := BTCurveGeometryInterpolatedSpline116{}
	return &this
}

// NewBTCurveGeometryInterpolatedSpline116WithDefaults instantiates a new BTCurveGeometryInterpolatedSpline116 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCurveGeometryInterpolatedSpline116WithDefaults() *BTCurveGeometryInterpolatedSpline116 {
	this := BTCurveGeometryInterpolatedSpline116{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTCurveGeometryInterpolatedSpline116) SetBtType(v string) {
	o.BtType = &v
}

// GetDerivatives returns the Derivatives field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116) GetDerivatives() map[string][]float64 {
	if o == nil || o.Derivatives == nil {
		var ret map[string][]float64
		return ret
	}
	return *o.Derivatives
}

// GetDerivativesOk returns a tuple with the Derivatives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116) GetDerivativesOk() (*map[string][]float64, bool) {
	if o == nil || o.Derivatives == nil {
		return nil, false
	}
	return o.Derivatives, true
}

// HasDerivatives returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116) HasDerivatives() bool {
	if o != nil && o.Derivatives != nil {
		return true
	}

	return false
}

// SetDerivatives gets a reference to the given map[string][]float64 and assigns it to the Derivatives field.
func (o *BTCurveGeometryInterpolatedSpline116) SetDerivatives(v map[string][]float64) {
	o.Derivatives = &v
}

// GetEndDerivativeX returns the EndDerivativeX field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116) GetEndDerivativeX() float64 {
	if o == nil || o.EndDerivativeX == nil {
		var ret float64
		return ret
	}
	return *o.EndDerivativeX
}

// GetEndDerivativeXOk returns a tuple with the EndDerivativeX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116) GetEndDerivativeXOk() (*float64, bool) {
	if o == nil || o.EndDerivativeX == nil {
		return nil, false
	}
	return o.EndDerivativeX, true
}

// HasEndDerivativeX returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116) HasEndDerivativeX() bool {
	if o != nil && o.EndDerivativeX != nil {
		return true
	}

	return false
}

// SetEndDerivativeX gets a reference to the given float64 and assigns it to the EndDerivativeX field.
func (o *BTCurveGeometryInterpolatedSpline116) SetEndDerivativeX(v float64) {
	o.EndDerivativeX = &v
}

// GetEndDerivativeY returns the EndDerivativeY field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116) GetEndDerivativeY() float64 {
	if o == nil || o.EndDerivativeY == nil {
		var ret float64
		return ret
	}
	return *o.EndDerivativeY
}

// GetEndDerivativeYOk returns a tuple with the EndDerivativeY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116) GetEndDerivativeYOk() (*float64, bool) {
	if o == nil || o.EndDerivativeY == nil {
		return nil, false
	}
	return o.EndDerivativeY, true
}

// HasEndDerivativeY returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116) HasEndDerivativeY() bool {
	if o != nil && o.EndDerivativeY != nil {
		return true
	}

	return false
}

// SetEndDerivativeY gets a reference to the given float64 and assigns it to the EndDerivativeY field.
func (o *BTCurveGeometryInterpolatedSpline116) SetEndDerivativeY(v float64) {
	o.EndDerivativeY = &v
}

// GetEndHandleX returns the EndHandleX field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116) GetEndHandleX() float64 {
	if o == nil || o.EndHandleX == nil {
		var ret float64
		return ret
	}
	return *o.EndHandleX
}

// GetEndHandleXOk returns a tuple with the EndHandleX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116) GetEndHandleXOk() (*float64, bool) {
	if o == nil || o.EndHandleX == nil {
		return nil, false
	}
	return o.EndHandleX, true
}

// HasEndHandleX returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116) HasEndHandleX() bool {
	if o != nil && o.EndHandleX != nil {
		return true
	}

	return false
}

// SetEndHandleX gets a reference to the given float64 and assigns it to the EndHandleX field.
func (o *BTCurveGeometryInterpolatedSpline116) SetEndHandleX(v float64) {
	o.EndHandleX = &v
}

// GetEndHandleY returns the EndHandleY field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116) GetEndHandleY() float64 {
	if o == nil || o.EndHandleY == nil {
		var ret float64
		return ret
	}
	return *o.EndHandleY
}

// GetEndHandleYOk returns a tuple with the EndHandleY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116) GetEndHandleYOk() (*float64, bool) {
	if o == nil || o.EndHandleY == nil {
		return nil, false
	}
	return o.EndHandleY, true
}

// HasEndHandleY returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116) HasEndHandleY() bool {
	if o != nil && o.EndHandleY != nil {
		return true
	}

	return false
}

// SetEndHandleY gets a reference to the given float64 and assigns it to the EndHandleY field.
func (o *BTCurveGeometryInterpolatedSpline116) SetEndHandleY(v float64) {
	o.EndHandleY = &v
}

// GetInterpolationPoints returns the InterpolationPoints field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116) GetInterpolationPoints() []float64 {
	if o == nil || o.InterpolationPoints == nil {
		var ret []float64
		return ret
	}
	return o.InterpolationPoints
}

// GetInterpolationPointsOk returns a tuple with the InterpolationPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116) GetInterpolationPointsOk() ([]float64, bool) {
	if o == nil || o.InterpolationPoints == nil {
		return nil, false
	}
	return o.InterpolationPoints, true
}

// HasInterpolationPoints returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116) HasInterpolationPoints() bool {
	if o != nil && o.InterpolationPoints != nil {
		return true
	}

	return false
}

// SetInterpolationPoints gets a reference to the given []float64 and assigns it to the InterpolationPoints field.
func (o *BTCurveGeometryInterpolatedSpline116) SetInterpolationPoints(v []float64) {
	o.InterpolationPoints = v
}

// GetIsPeriodic returns the IsPeriodic field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116) GetIsPeriodic() bool {
	if o == nil || o.IsPeriodic == nil {
		var ret bool
		return ret
	}
	return *o.IsPeriodic
}

// GetIsPeriodicOk returns a tuple with the IsPeriodic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116) GetIsPeriodicOk() (*bool, bool) {
	if o == nil || o.IsPeriodic == nil {
		return nil, false
	}
	return o.IsPeriodic, true
}

// HasIsPeriodic returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116) HasIsPeriodic() bool {
	if o != nil && o.IsPeriodic != nil {
		return true
	}

	return false
}

// SetIsPeriodic gets a reference to the given bool and assigns it to the IsPeriodic field.
func (o *BTCurveGeometryInterpolatedSpline116) SetIsPeriodic(v bool) {
	o.IsPeriodic = &v
}

// GetStartDerivativeX returns the StartDerivativeX field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116) GetStartDerivativeX() float64 {
	if o == nil || o.StartDerivativeX == nil {
		var ret float64
		return ret
	}
	return *o.StartDerivativeX
}

// GetStartDerivativeXOk returns a tuple with the StartDerivativeX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116) GetStartDerivativeXOk() (*float64, bool) {
	if o == nil || o.StartDerivativeX == nil {
		return nil, false
	}
	return o.StartDerivativeX, true
}

// HasStartDerivativeX returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116) HasStartDerivativeX() bool {
	if o != nil && o.StartDerivativeX != nil {
		return true
	}

	return false
}

// SetStartDerivativeX gets a reference to the given float64 and assigns it to the StartDerivativeX field.
func (o *BTCurveGeometryInterpolatedSpline116) SetStartDerivativeX(v float64) {
	o.StartDerivativeX = &v
}

// GetStartDerivativeY returns the StartDerivativeY field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116) GetStartDerivativeY() float64 {
	if o == nil || o.StartDerivativeY == nil {
		var ret float64
		return ret
	}
	return *o.StartDerivativeY
}

// GetStartDerivativeYOk returns a tuple with the StartDerivativeY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116) GetStartDerivativeYOk() (*float64, bool) {
	if o == nil || o.StartDerivativeY == nil {
		return nil, false
	}
	return o.StartDerivativeY, true
}

// HasStartDerivativeY returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116) HasStartDerivativeY() bool {
	if o != nil && o.StartDerivativeY != nil {
		return true
	}

	return false
}

// SetStartDerivativeY gets a reference to the given float64 and assigns it to the StartDerivativeY field.
func (o *BTCurveGeometryInterpolatedSpline116) SetStartDerivativeY(v float64) {
	o.StartDerivativeY = &v
}

// GetStartHandleX returns the StartHandleX field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116) GetStartHandleX() float64 {
	if o == nil || o.StartHandleX == nil {
		var ret float64
		return ret
	}
	return *o.StartHandleX
}

// GetStartHandleXOk returns a tuple with the StartHandleX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116) GetStartHandleXOk() (*float64, bool) {
	if o == nil || o.StartHandleX == nil {
		return nil, false
	}
	return o.StartHandleX, true
}

// HasStartHandleX returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116) HasStartHandleX() bool {
	if o != nil && o.StartHandleX != nil {
		return true
	}

	return false
}

// SetStartHandleX gets a reference to the given float64 and assigns it to the StartHandleX field.
func (o *BTCurveGeometryInterpolatedSpline116) SetStartHandleX(v float64) {
	o.StartHandleX = &v
}

// GetStartHandleY returns the StartHandleY field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116) GetStartHandleY() float64 {
	if o == nil || o.StartHandleY == nil {
		var ret float64
		return ret
	}
	return *o.StartHandleY
}

// GetStartHandleYOk returns a tuple with the StartHandleY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116) GetStartHandleYOk() (*float64, bool) {
	if o == nil || o.StartHandleY == nil {
		return nil, false
	}
	return o.StartHandleY, true
}

// HasStartHandleY returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116) HasStartHandleY() bool {
	if o != nil && o.StartHandleY != nil {
		return true
	}

	return false
}

// SetStartHandleY gets a reference to the given float64 and assigns it to the StartHandleY field.
func (o *BTCurveGeometryInterpolatedSpline116) SetStartHandleY(v float64) {
	o.StartHandleY = &v
}

func (o BTCurveGeometryInterpolatedSpline116) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Derivatives != nil {
		toSerialize["derivatives"] = o.Derivatives
	}
	if o.EndDerivativeX != nil {
		toSerialize["endDerivativeX"] = o.EndDerivativeX
	}
	if o.EndDerivativeY != nil {
		toSerialize["endDerivativeY"] = o.EndDerivativeY
	}
	if o.EndHandleX != nil {
		toSerialize["endHandleX"] = o.EndHandleX
	}
	if o.EndHandleY != nil {
		toSerialize["endHandleY"] = o.EndHandleY
	}
	if o.InterpolationPoints != nil {
		toSerialize["interpolationPoints"] = o.InterpolationPoints
	}
	if o.IsPeriodic != nil {
		toSerialize["isPeriodic"] = o.IsPeriodic
	}
	if o.StartDerivativeX != nil {
		toSerialize["startDerivativeX"] = o.StartDerivativeX
	}
	if o.StartDerivativeY != nil {
		toSerialize["startDerivativeY"] = o.StartDerivativeY
	}
	if o.StartHandleX != nil {
		toSerialize["startHandleX"] = o.StartHandleX
	}
	if o.StartHandleY != nil {
		toSerialize["startHandleY"] = o.StartHandleY
	}
	return json.Marshal(toSerialize)
}

type NullableBTCurveGeometryInterpolatedSpline116 struct {
	value *BTCurveGeometryInterpolatedSpline116
	isSet bool
}

func (v NullableBTCurveGeometryInterpolatedSpline116) Get() *BTCurveGeometryInterpolatedSpline116 {
	return v.value
}

func (v *NullableBTCurveGeometryInterpolatedSpline116) Set(val *BTCurveGeometryInterpolatedSpline116) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCurveGeometryInterpolatedSpline116) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCurveGeometryInterpolatedSpline116) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCurveGeometryInterpolatedSpline116(val *BTCurveGeometryInterpolatedSpline116) *NullableBTCurveGeometryInterpolatedSpline116 {
	return &NullableBTCurveGeometryInterpolatedSpline116{value: val, isSet: true}
}

func (v NullableBTCurveGeometryInterpolatedSpline116) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCurveGeometryInterpolatedSpline116) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
