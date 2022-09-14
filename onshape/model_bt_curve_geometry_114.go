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

// BTCurveGeometry114 - struct for BTCurveGeometry114
type BTCurveGeometry114 struct {
	implBTCurveGeometry114 interface{}
}

// BTCurveGeometryCircle115AsBTCurveGeometry114 is a convenience function that returns BTCurveGeometryCircle115 wrapped in BTCurveGeometry114
func (o *BTCurveGeometryCircle115) AsBTCurveGeometry114() *BTCurveGeometry114 {
	return &BTCurveGeometry114{o}
}

// BTCurveGeometryEllipse1189AsBTCurveGeometry114 is a convenience function that returns BTCurveGeometryEllipse1189 wrapped in BTCurveGeometry114
func (o *BTCurveGeometryEllipse1189) AsBTCurveGeometry114() *BTCurveGeometry114 {
	return &BTCurveGeometry114{o}
}

// BTCurveGeometryLine117AsBTCurveGeometry114 is a convenience function that returns BTCurveGeometryLine117 wrapped in BTCurveGeometry114
func (o *BTCurveGeometryLine117) AsBTCurveGeometry114() *BTCurveGeometry114 {
	return &BTCurveGeometry114{o}
}

// BTCurveGeometryControlPointSpline2197AsBTCurveGeometry114 is a convenience function that returns BTCurveGeometryControlPointSpline2197 wrapped in BTCurveGeometry114
func (o *BTCurveGeometryControlPointSpline2197) AsBTCurveGeometry114() *BTCurveGeometry114 {
	return &BTCurveGeometry114{o}
}

// BTCurveGeometrySpline118AsBTCurveGeometry114 is a convenience function that returns BTCurveGeometrySpline118 wrapped in BTCurveGeometry114
func (o *BTCurveGeometrySpline118) AsBTCurveGeometry114() *BTCurveGeometry114 {
	return &BTCurveGeometry114{o}
}

// BTCurveGeometryConic2284AsBTCurveGeometry114 is a convenience function that returns BTCurveGeometryConic2284 wrapped in BTCurveGeometry114
func (o *BTCurveGeometryConic2284) AsBTCurveGeometry114() *BTCurveGeometry114 {
	return &BTCurveGeometry114{o}
}

// BTCurveGeometryInterpolatedSpline116AsBTCurveGeometry114 is a convenience function that returns BTCurveGeometryInterpolatedSpline116 wrapped in BTCurveGeometry114
func (o *BTCurveGeometryInterpolatedSpline116) AsBTCurveGeometry114() *BTCurveGeometry114 {
	return &BTCurveGeometry114{o}
}

// NewBTCurveGeometry114 instantiates a new BTCurveGeometry114 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCurveGeometry114() *BTCurveGeometry114 {
	this := BTCurveGeometry114{Newbase_BTCurveGeometry114()}
	return &this
}

// NewBTCurveGeometry114WithDefaults instantiates a new BTCurveGeometry114 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCurveGeometry114WithDefaults() *BTCurveGeometry114 {
	this := BTCurveGeometry114{Newbase_BTCurveGeometry114WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTCurveGeometry114) GetBtType() string {
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
func (o *BTCurveGeometry114) GetBtTypeOk() (*string, bool) {
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
func (o *BTCurveGeometry114) HasBtType() bool {
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
func (o *BTCurveGeometry114) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTCurveGeometry114) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTCurveGeometryCircle-115'
	if jsonDict["btType"] == "BTCurveGeometryCircle-115" {
		// try to unmarshal JSON data into BTCurveGeometryCircle115
		var qr *BTCurveGeometryCircle115
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTCurveGeometry114 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTCurveGeometry114 = nil
			return fmt.Errorf("Failed to unmarshal BTCurveGeometry114 as BTCurveGeometryCircle115: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTCurveGeometryConic-2284'
	if jsonDict["btType"] == "BTCurveGeometryConic-2284" {
		// try to unmarshal JSON data into BTCurveGeometryConic2284
		var qr *BTCurveGeometryConic2284
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTCurveGeometry114 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTCurveGeometry114 = nil
			return fmt.Errorf("Failed to unmarshal BTCurveGeometry114 as BTCurveGeometryConic2284: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTCurveGeometryControlPointSpline-2197'
	if jsonDict["btType"] == "BTCurveGeometryControlPointSpline-2197" {
		// try to unmarshal JSON data into BTCurveGeometryControlPointSpline2197
		var qr *BTCurveGeometryControlPointSpline2197
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTCurveGeometry114 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTCurveGeometry114 = nil
			return fmt.Errorf("Failed to unmarshal BTCurveGeometry114 as BTCurveGeometryControlPointSpline2197: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTCurveGeometryEllipse-1189'
	if jsonDict["btType"] == "BTCurveGeometryEllipse-1189" {
		// try to unmarshal JSON data into BTCurveGeometryEllipse1189
		var qr *BTCurveGeometryEllipse1189
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTCurveGeometry114 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTCurveGeometry114 = nil
			return fmt.Errorf("Failed to unmarshal BTCurveGeometry114 as BTCurveGeometryEllipse1189: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTCurveGeometryInterpolatedSpline-116'
	if jsonDict["btType"] == "BTCurveGeometryInterpolatedSpline-116" {
		// try to unmarshal JSON data into BTCurveGeometryInterpolatedSpline116
		var qr *BTCurveGeometryInterpolatedSpline116
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTCurveGeometry114 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTCurveGeometry114 = nil
			return fmt.Errorf("Failed to unmarshal BTCurveGeometry114 as BTCurveGeometryInterpolatedSpline116: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTCurveGeometryLine-117'
	if jsonDict["btType"] == "BTCurveGeometryLine-117" {
		// try to unmarshal JSON data into BTCurveGeometryLine117
		var qr *BTCurveGeometryLine117
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTCurveGeometry114 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTCurveGeometry114 = nil
			return fmt.Errorf("Failed to unmarshal BTCurveGeometry114 as BTCurveGeometryLine117: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTCurveGeometrySpline-118'
	if jsonDict["btType"] == "BTCurveGeometrySpline-118" {
		// try to unmarshal JSON data into BTCurveGeometrySpline118
		var qr *BTCurveGeometrySpline118
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTCurveGeometry114 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTCurveGeometry114 = nil
			return fmt.Errorf("Failed to unmarshal BTCurveGeometry114 as BTCurveGeometrySpline118: %s", err.Error())
		}
	}

	var qtx *base_BTCurveGeometry114
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTCurveGeometry114 = qtx
		return nil // data stored in dst.base_BTCurveGeometry114, return on the first match
	} else {
		dst.implBTCurveGeometry114 = nil
		return fmt.Errorf("Failed to unmarshal BTCurveGeometry114 as base_BTCurveGeometry114: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTCurveGeometry114) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTCurveGeometry114) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTCurveGeometry114
}

type NullableBTCurveGeometry114 struct {
	value *BTCurveGeometry114
	isSet bool
}

func (v NullableBTCurveGeometry114) Get() *BTCurveGeometry114 {
	return v.value
}

func (v *NullableBTCurveGeometry114) Set(val *BTCurveGeometry114) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCurveGeometry114) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCurveGeometry114) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCurveGeometry114(val *BTCurveGeometry114) *NullableBTCurveGeometry114 {
	return &NullableBTCurveGeometry114{value: val, isSet: true}
}

func (v NullableBTCurveGeometry114) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCurveGeometry114) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTCurveGeometry114 struct {
	BtType *string `json:"btType,omitempty"`
}

// Newbase_BTCurveGeometry114 instantiates a new base_BTCurveGeometry114 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTCurveGeometry114() *base_BTCurveGeometry114 {
	this := base_BTCurveGeometry114{}
	return &this
}

// Newbase_BTCurveGeometry114WithDefaults instantiates a new base_BTCurveGeometry114 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTCurveGeometry114WithDefaults() *base_BTCurveGeometry114 {
	this := base_BTCurveGeometry114{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTCurveGeometry114) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveGeometry114) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTCurveGeometry114) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTCurveGeometry114) SetBtType(v string) {
	o.BtType = &v
}

func (o base_BTCurveGeometry114) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}
