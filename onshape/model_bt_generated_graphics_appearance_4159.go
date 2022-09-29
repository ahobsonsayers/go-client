/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6672-3324fec9d980
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTGeneratedGraphicsAppearance4159 struct for BTGeneratedGraphicsAppearance4159
type BTGeneratedGraphicsAppearance4159 struct {
	BtType           *string  `json:"btType,omitempty"`
	Color            []string `json:"color,omitempty"`
	NonTrivial       *bool    `json:"nonTrivial,omitempty"`
	Opacity          *int32   `json:"opacity,omitempty"`
	Reset            *bool    `json:"reset,omitempty"`
	RgbaColor        []string `json:"rgbaColor,omitempty"`
	UsableAppearance *bool    `json:"usableAppearance,omitempty"`
}

// NewBTGeneratedGraphicsAppearance4159 instantiates a new BTGeneratedGraphicsAppearance4159 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTGeneratedGraphicsAppearance4159() *BTGeneratedGraphicsAppearance4159 {
	this := BTGeneratedGraphicsAppearance4159{}
	return &this
}

// NewBTGeneratedGraphicsAppearance4159WithDefaults instantiates a new BTGeneratedGraphicsAppearance4159 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTGeneratedGraphicsAppearance4159WithDefaults() *BTGeneratedGraphicsAppearance4159 {
	this := BTGeneratedGraphicsAppearance4159{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTGeneratedGraphicsAppearance4159) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeneratedGraphicsAppearance4159) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTGeneratedGraphicsAppearance4159) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTGeneratedGraphicsAppearance4159) SetBtType(v string) {
	o.BtType = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *BTGeneratedGraphicsAppearance4159) GetColor() []string {
	if o == nil || o.Color == nil {
		var ret []string
		return ret
	}
	return o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeneratedGraphicsAppearance4159) GetColorOk() ([]string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *BTGeneratedGraphicsAppearance4159) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given []string and assigns it to the Color field.
func (o *BTGeneratedGraphicsAppearance4159) SetColor(v []string) {
	o.Color = v
}

// GetNonTrivial returns the NonTrivial field value if set, zero value otherwise.
func (o *BTGeneratedGraphicsAppearance4159) GetNonTrivial() bool {
	if o == nil || o.NonTrivial == nil {
		var ret bool
		return ret
	}
	return *o.NonTrivial
}

// GetNonTrivialOk returns a tuple with the NonTrivial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeneratedGraphicsAppearance4159) GetNonTrivialOk() (*bool, bool) {
	if o == nil || o.NonTrivial == nil {
		return nil, false
	}
	return o.NonTrivial, true
}

// HasNonTrivial returns a boolean if a field has been set.
func (o *BTGeneratedGraphicsAppearance4159) HasNonTrivial() bool {
	if o != nil && o.NonTrivial != nil {
		return true
	}

	return false
}

// SetNonTrivial gets a reference to the given bool and assigns it to the NonTrivial field.
func (o *BTGeneratedGraphicsAppearance4159) SetNonTrivial(v bool) {
	o.NonTrivial = &v
}

// GetOpacity returns the Opacity field value if set, zero value otherwise.
func (o *BTGeneratedGraphicsAppearance4159) GetOpacity() int32 {
	if o == nil || o.Opacity == nil {
		var ret int32
		return ret
	}
	return *o.Opacity
}

// GetOpacityOk returns a tuple with the Opacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeneratedGraphicsAppearance4159) GetOpacityOk() (*int32, bool) {
	if o == nil || o.Opacity == nil {
		return nil, false
	}
	return o.Opacity, true
}

// HasOpacity returns a boolean if a field has been set.
func (o *BTGeneratedGraphicsAppearance4159) HasOpacity() bool {
	if o != nil && o.Opacity != nil {
		return true
	}

	return false
}

// SetOpacity gets a reference to the given int32 and assigns it to the Opacity field.
func (o *BTGeneratedGraphicsAppearance4159) SetOpacity(v int32) {
	o.Opacity = &v
}

// GetReset returns the Reset field value if set, zero value otherwise.
func (o *BTGeneratedGraphicsAppearance4159) GetReset() bool {
	if o == nil || o.Reset == nil {
		var ret bool
		return ret
	}
	return *o.Reset
}

// GetResetOk returns a tuple with the Reset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeneratedGraphicsAppearance4159) GetResetOk() (*bool, bool) {
	if o == nil || o.Reset == nil {
		return nil, false
	}
	return o.Reset, true
}

// HasReset returns a boolean if a field has been set.
func (o *BTGeneratedGraphicsAppearance4159) HasReset() bool {
	if o != nil && o.Reset != nil {
		return true
	}

	return false
}

// SetReset gets a reference to the given bool and assigns it to the Reset field.
func (o *BTGeneratedGraphicsAppearance4159) SetReset(v bool) {
	o.Reset = &v
}

// GetRgbaColor returns the RgbaColor field value if set, zero value otherwise.
func (o *BTGeneratedGraphicsAppearance4159) GetRgbaColor() []string {
	if o == nil || o.RgbaColor == nil {
		var ret []string
		return ret
	}
	return o.RgbaColor
}

// GetRgbaColorOk returns a tuple with the RgbaColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeneratedGraphicsAppearance4159) GetRgbaColorOk() ([]string, bool) {
	if o == nil || o.RgbaColor == nil {
		return nil, false
	}
	return o.RgbaColor, true
}

// HasRgbaColor returns a boolean if a field has been set.
func (o *BTGeneratedGraphicsAppearance4159) HasRgbaColor() bool {
	if o != nil && o.RgbaColor != nil {
		return true
	}

	return false
}

// SetRgbaColor gets a reference to the given []string and assigns it to the RgbaColor field.
func (o *BTGeneratedGraphicsAppearance4159) SetRgbaColor(v []string) {
	o.RgbaColor = v
}

// GetUsableAppearance returns the UsableAppearance field value if set, zero value otherwise.
func (o *BTGeneratedGraphicsAppearance4159) GetUsableAppearance() bool {
	if o == nil || o.UsableAppearance == nil {
		var ret bool
		return ret
	}
	return *o.UsableAppearance
}

// GetUsableAppearanceOk returns a tuple with the UsableAppearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeneratedGraphicsAppearance4159) GetUsableAppearanceOk() (*bool, bool) {
	if o == nil || o.UsableAppearance == nil {
		return nil, false
	}
	return o.UsableAppearance, true
}

// HasUsableAppearance returns a boolean if a field has been set.
func (o *BTGeneratedGraphicsAppearance4159) HasUsableAppearance() bool {
	if o != nil && o.UsableAppearance != nil {
		return true
	}

	return false
}

// SetUsableAppearance gets a reference to the given bool and assigns it to the UsableAppearance field.
func (o *BTGeneratedGraphicsAppearance4159) SetUsableAppearance(v bool) {
	o.UsableAppearance = &v
}

func (o BTGeneratedGraphicsAppearance4159) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.NonTrivial != nil {
		toSerialize["nonTrivial"] = o.NonTrivial
	}
	if o.Opacity != nil {
		toSerialize["opacity"] = o.Opacity
	}
	if o.Reset != nil {
		toSerialize["reset"] = o.Reset
	}
	if o.RgbaColor != nil {
		toSerialize["rgbaColor"] = o.RgbaColor
	}
	if o.UsableAppearance != nil {
		toSerialize["usableAppearance"] = o.UsableAppearance
	}
	return json.Marshal(toSerialize)
}

type NullableBTGeneratedGraphicsAppearance4159 struct {
	value *BTGeneratedGraphicsAppearance4159
	isSet bool
}

func (v NullableBTGeneratedGraphicsAppearance4159) Get() *BTGeneratedGraphicsAppearance4159 {
	return v.value
}

func (v *NullableBTGeneratedGraphicsAppearance4159) Set(val *BTGeneratedGraphicsAppearance4159) {
	v.value = val
	v.isSet = true
}

func (v NullableBTGeneratedGraphicsAppearance4159) IsSet() bool {
	return v.isSet
}

func (v *NullableBTGeneratedGraphicsAppearance4159) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTGeneratedGraphicsAppearance4159(val *BTGeneratedGraphicsAppearance4159) *NullableBTGeneratedGraphicsAppearance4159 {
	return &NullableBTGeneratedGraphicsAppearance4159{value: val, isSet: true}
}

func (v NullableBTGeneratedGraphicsAppearance4159) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTGeneratedGraphicsAppearance4159) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
