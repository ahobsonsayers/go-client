/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13452-78145f970001
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTViewDataInfo struct for BTViewDataInfo
type BTViewDataInfo struct {
	Angle          *float64  `json:"angle,omitempty"`
	CameraViewport []float64 `json:"cameraViewport,omitempty"`
	IsPerspective  *bool     `json:"isPerspective,omitempty"`
	ViewMatrix     []float64 `json:"viewMatrix,omitempty"`
}

// NewBTViewDataInfo instantiates a new BTViewDataInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTViewDataInfo() *BTViewDataInfo {
	this := BTViewDataInfo{}
	return &this
}

// NewBTViewDataInfoWithDefaults instantiates a new BTViewDataInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTViewDataInfoWithDefaults() *BTViewDataInfo {
	this := BTViewDataInfo{}
	return &this
}

// GetAngle returns the Angle field value if set, zero value otherwise.
func (o *BTViewDataInfo) GetAngle() float64 {
	if o == nil || o.Angle == nil {
		var ret float64
		return ret
	}
	return *o.Angle
}

// GetAngleOk returns a tuple with the Angle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTViewDataInfo) GetAngleOk() (*float64, bool) {
	if o == nil || o.Angle == nil {
		return nil, false
	}
	return o.Angle, true
}

// HasAngle returns a boolean if a field has been set.
func (o *BTViewDataInfo) HasAngle() bool {
	if o != nil && o.Angle != nil {
		return true
	}

	return false
}

// SetAngle gets a reference to the given float64 and assigns it to the Angle field.
func (o *BTViewDataInfo) SetAngle(v float64) {
	o.Angle = &v
}

// GetCameraViewport returns the CameraViewport field value if set, zero value otherwise.
func (o *BTViewDataInfo) GetCameraViewport() []float64 {
	if o == nil || o.CameraViewport == nil {
		var ret []float64
		return ret
	}
	return o.CameraViewport
}

// GetCameraViewportOk returns a tuple with the CameraViewport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTViewDataInfo) GetCameraViewportOk() ([]float64, bool) {
	if o == nil || o.CameraViewport == nil {
		return nil, false
	}
	return o.CameraViewport, true
}

// HasCameraViewport returns a boolean if a field has been set.
func (o *BTViewDataInfo) HasCameraViewport() bool {
	if o != nil && o.CameraViewport != nil {
		return true
	}

	return false
}

// SetCameraViewport gets a reference to the given []float64 and assigns it to the CameraViewport field.
func (o *BTViewDataInfo) SetCameraViewport(v []float64) {
	o.CameraViewport = v
}

// GetIsPerspective returns the IsPerspective field value if set, zero value otherwise.
func (o *BTViewDataInfo) GetIsPerspective() bool {
	if o == nil || o.IsPerspective == nil {
		var ret bool
		return ret
	}
	return *o.IsPerspective
}

// GetIsPerspectiveOk returns a tuple with the IsPerspective field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTViewDataInfo) GetIsPerspectiveOk() (*bool, bool) {
	if o == nil || o.IsPerspective == nil {
		return nil, false
	}
	return o.IsPerspective, true
}

// HasIsPerspective returns a boolean if a field has been set.
func (o *BTViewDataInfo) HasIsPerspective() bool {
	if o != nil && o.IsPerspective != nil {
		return true
	}

	return false
}

// SetIsPerspective gets a reference to the given bool and assigns it to the IsPerspective field.
func (o *BTViewDataInfo) SetIsPerspective(v bool) {
	o.IsPerspective = &v
}

// GetViewMatrix returns the ViewMatrix field value if set, zero value otherwise.
func (o *BTViewDataInfo) GetViewMatrix() []float64 {
	if o == nil || o.ViewMatrix == nil {
		var ret []float64
		return ret
	}
	return o.ViewMatrix
}

// GetViewMatrixOk returns a tuple with the ViewMatrix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTViewDataInfo) GetViewMatrixOk() ([]float64, bool) {
	if o == nil || o.ViewMatrix == nil {
		return nil, false
	}
	return o.ViewMatrix, true
}

// HasViewMatrix returns a boolean if a field has been set.
func (o *BTViewDataInfo) HasViewMatrix() bool {
	if o != nil && o.ViewMatrix != nil {
		return true
	}

	return false
}

// SetViewMatrix gets a reference to the given []float64 and assigns it to the ViewMatrix field.
func (o *BTViewDataInfo) SetViewMatrix(v []float64) {
	o.ViewMatrix = v
}

func (o BTViewDataInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Angle != nil {
		toSerialize["angle"] = o.Angle
	}
	if o.CameraViewport != nil {
		toSerialize["cameraViewport"] = o.CameraViewport
	}
	if o.IsPerspective != nil {
		toSerialize["isPerspective"] = o.IsPerspective
	}
	if o.ViewMatrix != nil {
		toSerialize["viewMatrix"] = o.ViewMatrix
	}
	return json.Marshal(toSerialize)
}

type NullableBTViewDataInfo struct {
	value *BTViewDataInfo
	isSet bool
}

func (v NullableBTViewDataInfo) Get() *BTViewDataInfo {
	return v.value
}

func (v *NullableBTViewDataInfo) Set(val *BTViewDataInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTViewDataInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTViewDataInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTViewDataInfo(val *BTViewDataInfo) *NullableBTViewDataInfo {
	return &NullableBTViewDataInfo{value: val, isSet: true}
}

func (v NullableBTViewDataInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTViewDataInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
