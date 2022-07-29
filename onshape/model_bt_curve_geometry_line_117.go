/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5777-46062b6e03f9
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCurveGeometryLine117 struct for BTCurveGeometryLine117
type BTCurveGeometryLine117 struct {
	BtType *string  `json:"btType,omitempty"`
	DirX   *float64 `json:"dirX,omitempty"`
	DirY   *float64 `json:"dirY,omitempty"`
	PntX   *float64 `json:"pntX,omitempty"`
	PntY   *float64 `json:"pntY,omitempty"`
}

// NewBTCurveGeometryLine117 instantiates a new BTCurveGeometryLine117 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCurveGeometryLine117() *BTCurveGeometryLine117 {
	this := BTCurveGeometryLine117{}
	return &this
}

// NewBTCurveGeometryLine117WithDefaults instantiates a new BTCurveGeometryLine117 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCurveGeometryLine117WithDefaults() *BTCurveGeometryLine117 {
	this := BTCurveGeometryLine117{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTCurveGeometryLine117) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryLine117) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTCurveGeometryLine117) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTCurveGeometryLine117) SetBtType(v string) {
	o.BtType = &v
}

// GetDirX returns the DirX field value if set, zero value otherwise.
func (o *BTCurveGeometryLine117) GetDirX() float64 {
	if o == nil || o.DirX == nil {
		var ret float64
		return ret
	}
	return *o.DirX
}

// GetDirXOk returns a tuple with the DirX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryLine117) GetDirXOk() (*float64, bool) {
	if o == nil || o.DirX == nil {
		return nil, false
	}
	return o.DirX, true
}

// HasDirX returns a boolean if a field has been set.
func (o *BTCurveGeometryLine117) HasDirX() bool {
	if o != nil && o.DirX != nil {
		return true
	}

	return false
}

// SetDirX gets a reference to the given float64 and assigns it to the DirX field.
func (o *BTCurveGeometryLine117) SetDirX(v float64) {
	o.DirX = &v
}

// GetDirY returns the DirY field value if set, zero value otherwise.
func (o *BTCurveGeometryLine117) GetDirY() float64 {
	if o == nil || o.DirY == nil {
		var ret float64
		return ret
	}
	return *o.DirY
}

// GetDirYOk returns a tuple with the DirY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryLine117) GetDirYOk() (*float64, bool) {
	if o == nil || o.DirY == nil {
		return nil, false
	}
	return o.DirY, true
}

// HasDirY returns a boolean if a field has been set.
func (o *BTCurveGeometryLine117) HasDirY() bool {
	if o != nil && o.DirY != nil {
		return true
	}

	return false
}

// SetDirY gets a reference to the given float64 and assigns it to the DirY field.
func (o *BTCurveGeometryLine117) SetDirY(v float64) {
	o.DirY = &v
}

// GetPntX returns the PntX field value if set, zero value otherwise.
func (o *BTCurveGeometryLine117) GetPntX() float64 {
	if o == nil || o.PntX == nil {
		var ret float64
		return ret
	}
	return *o.PntX
}

// GetPntXOk returns a tuple with the PntX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryLine117) GetPntXOk() (*float64, bool) {
	if o == nil || o.PntX == nil {
		return nil, false
	}
	return o.PntX, true
}

// HasPntX returns a boolean if a field has been set.
func (o *BTCurveGeometryLine117) HasPntX() bool {
	if o != nil && o.PntX != nil {
		return true
	}

	return false
}

// SetPntX gets a reference to the given float64 and assigns it to the PntX field.
func (o *BTCurveGeometryLine117) SetPntX(v float64) {
	o.PntX = &v
}

// GetPntY returns the PntY field value if set, zero value otherwise.
func (o *BTCurveGeometryLine117) GetPntY() float64 {
	if o == nil || o.PntY == nil {
		var ret float64
		return ret
	}
	return *o.PntY
}

// GetPntYOk returns a tuple with the PntY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryLine117) GetPntYOk() (*float64, bool) {
	if o == nil || o.PntY == nil {
		return nil, false
	}
	return o.PntY, true
}

// HasPntY returns a boolean if a field has been set.
func (o *BTCurveGeometryLine117) HasPntY() bool {
	if o != nil && o.PntY != nil {
		return true
	}

	return false
}

// SetPntY gets a reference to the given float64 and assigns it to the PntY field.
func (o *BTCurveGeometryLine117) SetPntY(v float64) {
	o.PntY = &v
}

func (o BTCurveGeometryLine117) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DirX != nil {
		toSerialize["dirX"] = o.DirX
	}
	if o.DirY != nil {
		toSerialize["dirY"] = o.DirY
	}
	if o.PntX != nil {
		toSerialize["pntX"] = o.PntX
	}
	if o.PntY != nil {
		toSerialize["pntY"] = o.PntY
	}
	return json.Marshal(toSerialize)
}

type NullableBTCurveGeometryLine117 struct {
	value *BTCurveGeometryLine117
	isSet bool
}

func (v NullableBTCurveGeometryLine117) Get() *BTCurveGeometryLine117 {
	return v.value
}

func (v *NullableBTCurveGeometryLine117) Set(val *BTCurveGeometryLine117) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCurveGeometryLine117) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCurveGeometryLine117) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCurveGeometryLine117(val *BTCurveGeometryLine117) *NullableBTCurveGeometryLine117 {
	return &NullableBTCurveGeometryLine117{value: val, isSet: true}
}

func (v NullableBTCurveGeometryLine117) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCurveGeometryLine117) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
