/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.164.16955-b4ecd192bba6
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSketchImageDisplayData1379 struct for BTSketchImageDisplayData1379
type BTSketchImageDisplayData1379 struct {
	BottomLeftCorner  *BTVector3d389                 `json:"bottomLeftCorner,omitempty"`
	BottomRightCorner *BTVector3d389                 `json:"bottomRightCorner,omitempty"`
	BtType            *string                        `json:"btType,omitempty"`
	Entities          []BTSketchEntityDisplayData354 `json:"entities,omitempty"`
	FeatureId         *string                        `json:"featureId,omitempty"`
	IsOnFlat          *bool                          `json:"isOnFlat,omitempty"`
	Points            []float64                      `json:"points,omitempty"`
	SourceId          *string                        `json:"sourceId,omitempty"`
	TopLeftCorner     *BTVector3d389                 `json:"topLeftCorner,omitempty"`
}

// NewBTSketchImageDisplayData1379 instantiates a new BTSketchImageDisplayData1379 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSketchImageDisplayData1379() *BTSketchImageDisplayData1379 {
	this := BTSketchImageDisplayData1379{}
	return &this
}

// NewBTSketchImageDisplayData1379WithDefaults instantiates a new BTSketchImageDisplayData1379 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSketchImageDisplayData1379WithDefaults() *BTSketchImageDisplayData1379 {
	this := BTSketchImageDisplayData1379{}
	return &this
}

// GetBottomLeftCorner returns the BottomLeftCorner field value if set, zero value otherwise.
func (o *BTSketchImageDisplayData1379) GetBottomLeftCorner() BTVector3d389 {
	if o == nil || o.BottomLeftCorner == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.BottomLeftCorner
}

// GetBottomLeftCornerOk returns a tuple with the BottomLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchImageDisplayData1379) GetBottomLeftCornerOk() (*BTVector3d389, bool) {
	if o == nil || o.BottomLeftCorner == nil {
		return nil, false
	}
	return o.BottomLeftCorner, true
}

// HasBottomLeftCorner returns a boolean if a field has been set.
func (o *BTSketchImageDisplayData1379) HasBottomLeftCorner() bool {
	if o != nil && o.BottomLeftCorner != nil {
		return true
	}

	return false
}

// SetBottomLeftCorner gets a reference to the given BTVector3d389 and assigns it to the BottomLeftCorner field.
func (o *BTSketchImageDisplayData1379) SetBottomLeftCorner(v BTVector3d389) {
	o.BottomLeftCorner = &v
}

// GetBottomRightCorner returns the BottomRightCorner field value if set, zero value otherwise.
func (o *BTSketchImageDisplayData1379) GetBottomRightCorner() BTVector3d389 {
	if o == nil || o.BottomRightCorner == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.BottomRightCorner
}

// GetBottomRightCornerOk returns a tuple with the BottomRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchImageDisplayData1379) GetBottomRightCornerOk() (*BTVector3d389, bool) {
	if o == nil || o.BottomRightCorner == nil {
		return nil, false
	}
	return o.BottomRightCorner, true
}

// HasBottomRightCorner returns a boolean if a field has been set.
func (o *BTSketchImageDisplayData1379) HasBottomRightCorner() bool {
	if o != nil && o.BottomRightCorner != nil {
		return true
	}

	return false
}

// SetBottomRightCorner gets a reference to the given BTVector3d389 and assigns it to the BottomRightCorner field.
func (o *BTSketchImageDisplayData1379) SetBottomRightCorner(v BTVector3d389) {
	o.BottomRightCorner = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSketchImageDisplayData1379) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchImageDisplayData1379) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSketchImageDisplayData1379) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSketchImageDisplayData1379) SetBtType(v string) {
	o.BtType = &v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *BTSketchImageDisplayData1379) GetEntities() []BTSketchEntityDisplayData354 {
	if o == nil || o.Entities == nil {
		var ret []BTSketchEntityDisplayData354
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchImageDisplayData1379) GetEntitiesOk() ([]BTSketchEntityDisplayData354, bool) {
	if o == nil || o.Entities == nil {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *BTSketchImageDisplayData1379) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []BTSketchEntityDisplayData354 and assigns it to the Entities field.
func (o *BTSketchImageDisplayData1379) SetEntities(v []BTSketchEntityDisplayData354) {
	o.Entities = v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTSketchImageDisplayData1379) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchImageDisplayData1379) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTSketchImageDisplayData1379) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTSketchImageDisplayData1379) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetIsOnFlat returns the IsOnFlat field value if set, zero value otherwise.
func (o *BTSketchImageDisplayData1379) GetIsOnFlat() bool {
	if o == nil || o.IsOnFlat == nil {
		var ret bool
		return ret
	}
	return *o.IsOnFlat
}

// GetIsOnFlatOk returns a tuple with the IsOnFlat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchImageDisplayData1379) GetIsOnFlatOk() (*bool, bool) {
	if o == nil || o.IsOnFlat == nil {
		return nil, false
	}
	return o.IsOnFlat, true
}

// HasIsOnFlat returns a boolean if a field has been set.
func (o *BTSketchImageDisplayData1379) HasIsOnFlat() bool {
	if o != nil && o.IsOnFlat != nil {
		return true
	}

	return false
}

// SetIsOnFlat gets a reference to the given bool and assigns it to the IsOnFlat field.
func (o *BTSketchImageDisplayData1379) SetIsOnFlat(v bool) {
	o.IsOnFlat = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *BTSketchImageDisplayData1379) GetPoints() []float64 {
	if o == nil || o.Points == nil {
		var ret []float64
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchImageDisplayData1379) GetPointsOk() ([]float64, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *BTSketchImageDisplayData1379) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given []float64 and assigns it to the Points field.
func (o *BTSketchImageDisplayData1379) SetPoints(v []float64) {
	o.Points = v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *BTSketchImageDisplayData1379) GetSourceId() string {
	if o == nil || o.SourceId == nil {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchImageDisplayData1379) GetSourceIdOk() (*string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *BTSketchImageDisplayData1379) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *BTSketchImageDisplayData1379) SetSourceId(v string) {
	o.SourceId = &v
}

// GetTopLeftCorner returns the TopLeftCorner field value if set, zero value otherwise.
func (o *BTSketchImageDisplayData1379) GetTopLeftCorner() BTVector3d389 {
	if o == nil || o.TopLeftCorner == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.TopLeftCorner
}

// GetTopLeftCornerOk returns a tuple with the TopLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchImageDisplayData1379) GetTopLeftCornerOk() (*BTVector3d389, bool) {
	if o == nil || o.TopLeftCorner == nil {
		return nil, false
	}
	return o.TopLeftCorner, true
}

// HasTopLeftCorner returns a boolean if a field has been set.
func (o *BTSketchImageDisplayData1379) HasTopLeftCorner() bool {
	if o != nil && o.TopLeftCorner != nil {
		return true
	}

	return false
}

// SetTopLeftCorner gets a reference to the given BTVector3d389 and assigns it to the TopLeftCorner field.
func (o *BTSketchImageDisplayData1379) SetTopLeftCorner(v BTVector3d389) {
	o.TopLeftCorner = &v
}

func (o BTSketchImageDisplayData1379) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BottomLeftCorner != nil {
		toSerialize["bottomLeftCorner"] = o.BottomLeftCorner
	}
	if o.BottomRightCorner != nil {
		toSerialize["bottomRightCorner"] = o.BottomRightCorner
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.IsOnFlat != nil {
		toSerialize["isOnFlat"] = o.IsOnFlat
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	if o.SourceId != nil {
		toSerialize["sourceId"] = o.SourceId
	}
	if o.TopLeftCorner != nil {
		toSerialize["topLeftCorner"] = o.TopLeftCorner
	}
	return json.Marshal(toSerialize)
}

type NullableBTSketchImageDisplayData1379 struct {
	value *BTSketchImageDisplayData1379
	isSet bool
}

func (v NullableBTSketchImageDisplayData1379) Get() *BTSketchImageDisplayData1379 {
	return v.value
}

func (v *NullableBTSketchImageDisplayData1379) Set(val *BTSketchImageDisplayData1379) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSketchImageDisplayData1379) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSketchImageDisplayData1379) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSketchImageDisplayData1379(val *BTSketchImageDisplayData1379) *NullableBTSketchImageDisplayData1379 {
	return &NullableBTSketchImageDisplayData1379{value: val, isSet: true}
}

func (v NullableBTSketchImageDisplayData1379) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSketchImageDisplayData1379) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
