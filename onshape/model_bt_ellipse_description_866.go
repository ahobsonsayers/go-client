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

// BTEllipseDescription866 struct for BTEllipseDescription866
type BTEllipseDescription866 struct {
	BtType                    *string           `json:"btType,omitempty"`
	Direction                 *BTVector3d389    `json:"direction,omitempty"`
	DirectionOrientedWithFace *BTVector3d389    `json:"directionOrientedWithFace,omitempty"`
	Origin                    *BTVector3d389    `json:"origin,omitempty"`
	Type                      *GBTCurveTypeEnum `json:"type,omitempty"`
	MajorAxis                 *BTVector3d389    `json:"majorAxis,omitempty"`
	MajorRadius               *float64          `json:"majorRadius,omitempty"`
	MinorRadius               *float64          `json:"minorRadius,omitempty"`
	Normal                    *BTVector3d389    `json:"normal,omitempty"`
}

// NewBTEllipseDescription866 instantiates a new BTEllipseDescription866 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTEllipseDescription866() *BTEllipseDescription866 {
	this := BTEllipseDescription866{}
	return &this
}

// NewBTEllipseDescription866WithDefaults instantiates a new BTEllipseDescription866 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTEllipseDescription866WithDefaults() *BTEllipseDescription866 {
	this := BTEllipseDescription866{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTEllipseDescription866) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDescription866) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTEllipseDescription866) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTEllipseDescription866) SetBtType(v string) {
	o.BtType = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *BTEllipseDescription866) GetDirection() BTVector3d389 {
	if o == nil || o.Direction == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDescription866) GetDirectionOk() (*BTVector3d389, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *BTEllipseDescription866) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given BTVector3d389 and assigns it to the Direction field.
func (o *BTEllipseDescription866) SetDirection(v BTVector3d389) {
	o.Direction = &v
}

// GetDirectionOrientedWithFace returns the DirectionOrientedWithFace field value if set, zero value otherwise.
func (o *BTEllipseDescription866) GetDirectionOrientedWithFace() BTVector3d389 {
	if o == nil || o.DirectionOrientedWithFace == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.DirectionOrientedWithFace
}

// GetDirectionOrientedWithFaceOk returns a tuple with the DirectionOrientedWithFace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDescription866) GetDirectionOrientedWithFaceOk() (*BTVector3d389, bool) {
	if o == nil || o.DirectionOrientedWithFace == nil {
		return nil, false
	}
	return o.DirectionOrientedWithFace, true
}

// HasDirectionOrientedWithFace returns a boolean if a field has been set.
func (o *BTEllipseDescription866) HasDirectionOrientedWithFace() bool {
	if o != nil && o.DirectionOrientedWithFace != nil {
		return true
	}

	return false
}

// SetDirectionOrientedWithFace gets a reference to the given BTVector3d389 and assigns it to the DirectionOrientedWithFace field.
func (o *BTEllipseDescription866) SetDirectionOrientedWithFace(v BTVector3d389) {
	o.DirectionOrientedWithFace = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BTEllipseDescription866) GetOrigin() BTVector3d389 {
	if o == nil || o.Origin == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDescription866) GetOriginOk() (*BTVector3d389, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *BTEllipseDescription866) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given BTVector3d389 and assigns it to the Origin field.
func (o *BTEllipseDescription866) SetOrigin(v BTVector3d389) {
	o.Origin = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTEllipseDescription866) GetType() GBTCurveTypeEnum {
	if o == nil || o.Type == nil {
		var ret GBTCurveTypeEnum
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDescription866) GetTypeOk() (*GBTCurveTypeEnum, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTEllipseDescription866) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given GBTCurveTypeEnum and assigns it to the Type field.
func (o *BTEllipseDescription866) SetType(v GBTCurveTypeEnum) {
	o.Type = &v
}

// GetMajorAxis returns the MajorAxis field value if set, zero value otherwise.
func (o *BTEllipseDescription866) GetMajorAxis() BTVector3d389 {
	if o == nil || o.MajorAxis == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.MajorAxis
}

// GetMajorAxisOk returns a tuple with the MajorAxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDescription866) GetMajorAxisOk() (*BTVector3d389, bool) {
	if o == nil || o.MajorAxis == nil {
		return nil, false
	}
	return o.MajorAxis, true
}

// HasMajorAxis returns a boolean if a field has been set.
func (o *BTEllipseDescription866) HasMajorAxis() bool {
	if o != nil && o.MajorAxis != nil {
		return true
	}

	return false
}

// SetMajorAxis gets a reference to the given BTVector3d389 and assigns it to the MajorAxis field.
func (o *BTEllipseDescription866) SetMajorAxis(v BTVector3d389) {
	o.MajorAxis = &v
}

// GetMajorRadius returns the MajorRadius field value if set, zero value otherwise.
func (o *BTEllipseDescription866) GetMajorRadius() float64 {
	if o == nil || o.MajorRadius == nil {
		var ret float64
		return ret
	}
	return *o.MajorRadius
}

// GetMajorRadiusOk returns a tuple with the MajorRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDescription866) GetMajorRadiusOk() (*float64, bool) {
	if o == nil || o.MajorRadius == nil {
		return nil, false
	}
	return o.MajorRadius, true
}

// HasMajorRadius returns a boolean if a field has been set.
func (o *BTEllipseDescription866) HasMajorRadius() bool {
	if o != nil && o.MajorRadius != nil {
		return true
	}

	return false
}

// SetMajorRadius gets a reference to the given float64 and assigns it to the MajorRadius field.
func (o *BTEllipseDescription866) SetMajorRadius(v float64) {
	o.MajorRadius = &v
}

// GetMinorRadius returns the MinorRadius field value if set, zero value otherwise.
func (o *BTEllipseDescription866) GetMinorRadius() float64 {
	if o == nil || o.MinorRadius == nil {
		var ret float64
		return ret
	}
	return *o.MinorRadius
}

// GetMinorRadiusOk returns a tuple with the MinorRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDescription866) GetMinorRadiusOk() (*float64, bool) {
	if o == nil || o.MinorRadius == nil {
		return nil, false
	}
	return o.MinorRadius, true
}

// HasMinorRadius returns a boolean if a field has been set.
func (o *BTEllipseDescription866) HasMinorRadius() bool {
	if o != nil && o.MinorRadius != nil {
		return true
	}

	return false
}

// SetMinorRadius gets a reference to the given float64 and assigns it to the MinorRadius field.
func (o *BTEllipseDescription866) SetMinorRadius(v float64) {
	o.MinorRadius = &v
}

// GetNormal returns the Normal field value if set, zero value otherwise.
func (o *BTEllipseDescription866) GetNormal() BTVector3d389 {
	if o == nil || o.Normal == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Normal
}

// GetNormalOk returns a tuple with the Normal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDescription866) GetNormalOk() (*BTVector3d389, bool) {
	if o == nil || o.Normal == nil {
		return nil, false
	}
	return o.Normal, true
}

// HasNormal returns a boolean if a field has been set.
func (o *BTEllipseDescription866) HasNormal() bool {
	if o != nil && o.Normal != nil {
		return true
	}

	return false
}

// SetNormal gets a reference to the given BTVector3d389 and assigns it to the Normal field.
func (o *BTEllipseDescription866) SetNormal(v BTVector3d389) {
	o.Normal = &v
}

func (o BTEllipseDescription866) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.DirectionOrientedWithFace != nil {
		toSerialize["directionOrientedWithFace"] = o.DirectionOrientedWithFace
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.MajorAxis != nil {
		toSerialize["majorAxis"] = o.MajorAxis
	}
	if o.MajorRadius != nil {
		toSerialize["majorRadius"] = o.MajorRadius
	}
	if o.MinorRadius != nil {
		toSerialize["minorRadius"] = o.MinorRadius
	}
	if o.Normal != nil {
		toSerialize["normal"] = o.Normal
	}
	return json.Marshal(toSerialize)
}

type NullableBTEllipseDescription866 struct {
	value *BTEllipseDescription866
	isSet bool
}

func (v NullableBTEllipseDescription866) Get() *BTEllipseDescription866 {
	return v.value
}

func (v *NullableBTEllipseDescription866) Set(val *BTEllipseDescription866) {
	v.value = val
	v.isSet = true
}

func (v NullableBTEllipseDescription866) IsSet() bool {
	return v.isSet
}

func (v *NullableBTEllipseDescription866) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTEllipseDescription866(val *BTEllipseDescription866) *NullableBTEllipseDescription866 {
	return &NullableBTEllipseDescription866{value: val, isSet: true}
}

func (v NullableBTEllipseDescription866) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTEllipseDescription866) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
