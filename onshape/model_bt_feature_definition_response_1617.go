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

// BTFeatureDefinitionResponse1617 struct for BTFeatureDefinitionResponse1617
type BTFeatureDefinitionResponse1617 struct {
	BtType                 *string             `json:"btType,omitempty"`
	Feature                *BTMFeature134      `json:"feature,omitempty"`
	FeatureState           *BTFeatureState1688 `json:"featureState,omitempty"`
	LibraryVersion         *int32              `json:"libraryVersion,omitempty"`
	MicroversionSkew       *bool               `json:"microversionSkew,omitempty"`
	RejectMicroversionSkew *bool               `json:"rejectMicroversionSkew,omitempty"`
	SerializationVersion   *string             `json:"serializationVersion,omitempty"`
	SourceMicroversion     *string             `json:"sourceMicroversion,omitempty"`
}

// NewBTFeatureDefinitionResponse1617 instantiates a new BTFeatureDefinitionResponse1617 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFeatureDefinitionResponse1617() *BTFeatureDefinitionResponse1617 {
	this := BTFeatureDefinitionResponse1617{}
	return &this
}

// NewBTFeatureDefinitionResponse1617WithDefaults instantiates a new BTFeatureDefinitionResponse1617 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFeatureDefinitionResponse1617WithDefaults() *BTFeatureDefinitionResponse1617 {
	this := BTFeatureDefinitionResponse1617{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFeatureDefinitionResponse1617) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureDefinitionResponse1617) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFeatureDefinitionResponse1617) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFeatureDefinitionResponse1617) SetBtType(v string) {
	o.BtType = &v
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *BTFeatureDefinitionResponse1617) GetFeature() BTMFeature134 {
	if o == nil || o.Feature == nil {
		var ret BTMFeature134
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureDefinitionResponse1617) GetFeatureOk() (*BTMFeature134, bool) {
	if o == nil || o.Feature == nil {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *BTFeatureDefinitionResponse1617) HasFeature() bool {
	if o != nil && o.Feature != nil {
		return true
	}

	return false
}

// SetFeature gets a reference to the given BTMFeature134 and assigns it to the Feature field.
func (o *BTFeatureDefinitionResponse1617) SetFeature(v BTMFeature134) {
	o.Feature = &v
}

// GetFeatureState returns the FeatureState field value if set, zero value otherwise.
func (o *BTFeatureDefinitionResponse1617) GetFeatureState() BTFeatureState1688 {
	if o == nil || o.FeatureState == nil {
		var ret BTFeatureState1688
		return ret
	}
	return *o.FeatureState
}

// GetFeatureStateOk returns a tuple with the FeatureState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureDefinitionResponse1617) GetFeatureStateOk() (*BTFeatureState1688, bool) {
	if o == nil || o.FeatureState == nil {
		return nil, false
	}
	return o.FeatureState, true
}

// HasFeatureState returns a boolean if a field has been set.
func (o *BTFeatureDefinitionResponse1617) HasFeatureState() bool {
	if o != nil && o.FeatureState != nil {
		return true
	}

	return false
}

// SetFeatureState gets a reference to the given BTFeatureState1688 and assigns it to the FeatureState field.
func (o *BTFeatureDefinitionResponse1617) SetFeatureState(v BTFeatureState1688) {
	o.FeatureState = &v
}

// GetLibraryVersion returns the LibraryVersion field value if set, zero value otherwise.
func (o *BTFeatureDefinitionResponse1617) GetLibraryVersion() int32 {
	if o == nil || o.LibraryVersion == nil {
		var ret int32
		return ret
	}
	return *o.LibraryVersion
}

// GetLibraryVersionOk returns a tuple with the LibraryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureDefinitionResponse1617) GetLibraryVersionOk() (*int32, bool) {
	if o == nil || o.LibraryVersion == nil {
		return nil, false
	}
	return o.LibraryVersion, true
}

// HasLibraryVersion returns a boolean if a field has been set.
func (o *BTFeatureDefinitionResponse1617) HasLibraryVersion() bool {
	if o != nil && o.LibraryVersion != nil {
		return true
	}

	return false
}

// SetLibraryVersion gets a reference to the given int32 and assigns it to the LibraryVersion field.
func (o *BTFeatureDefinitionResponse1617) SetLibraryVersion(v int32) {
	o.LibraryVersion = &v
}

// GetMicroversionSkew returns the MicroversionSkew field value if set, zero value otherwise.
func (o *BTFeatureDefinitionResponse1617) GetMicroversionSkew() bool {
	if o == nil || o.MicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.MicroversionSkew
}

// GetMicroversionSkewOk returns a tuple with the MicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureDefinitionResponse1617) GetMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.MicroversionSkew == nil {
		return nil, false
	}
	return o.MicroversionSkew, true
}

// HasMicroversionSkew returns a boolean if a field has been set.
func (o *BTFeatureDefinitionResponse1617) HasMicroversionSkew() bool {
	if o != nil && o.MicroversionSkew != nil {
		return true
	}

	return false
}

// SetMicroversionSkew gets a reference to the given bool and assigns it to the MicroversionSkew field.
func (o *BTFeatureDefinitionResponse1617) SetMicroversionSkew(v bool) {
	o.MicroversionSkew = &v
}

// GetRejectMicroversionSkew returns the RejectMicroversionSkew field value if set, zero value otherwise.
func (o *BTFeatureDefinitionResponse1617) GetRejectMicroversionSkew() bool {
	if o == nil || o.RejectMicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.RejectMicroversionSkew
}

// GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureDefinitionResponse1617) GetRejectMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.RejectMicroversionSkew == nil {
		return nil, false
	}
	return o.RejectMicroversionSkew, true
}

// HasRejectMicroversionSkew returns a boolean if a field has been set.
func (o *BTFeatureDefinitionResponse1617) HasRejectMicroversionSkew() bool {
	if o != nil && o.RejectMicroversionSkew != nil {
		return true
	}

	return false
}

// SetRejectMicroversionSkew gets a reference to the given bool and assigns it to the RejectMicroversionSkew field.
func (o *BTFeatureDefinitionResponse1617) SetRejectMicroversionSkew(v bool) {
	o.RejectMicroversionSkew = &v
}

// GetSerializationVersion returns the SerializationVersion field value if set, zero value otherwise.
func (o *BTFeatureDefinitionResponse1617) GetSerializationVersion() string {
	if o == nil || o.SerializationVersion == nil {
		var ret string
		return ret
	}
	return *o.SerializationVersion
}

// GetSerializationVersionOk returns a tuple with the SerializationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureDefinitionResponse1617) GetSerializationVersionOk() (*string, bool) {
	if o == nil || o.SerializationVersion == nil {
		return nil, false
	}
	return o.SerializationVersion, true
}

// HasSerializationVersion returns a boolean if a field has been set.
func (o *BTFeatureDefinitionResponse1617) HasSerializationVersion() bool {
	if o != nil && o.SerializationVersion != nil {
		return true
	}

	return false
}

// SetSerializationVersion gets a reference to the given string and assigns it to the SerializationVersion field.
func (o *BTFeatureDefinitionResponse1617) SetSerializationVersion(v string) {
	o.SerializationVersion = &v
}

// GetSourceMicroversion returns the SourceMicroversion field value if set, zero value otherwise.
func (o *BTFeatureDefinitionResponse1617) GetSourceMicroversion() string {
	if o == nil || o.SourceMicroversion == nil {
		var ret string
		return ret
	}
	return *o.SourceMicroversion
}

// GetSourceMicroversionOk returns a tuple with the SourceMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureDefinitionResponse1617) GetSourceMicroversionOk() (*string, bool) {
	if o == nil || o.SourceMicroversion == nil {
		return nil, false
	}
	return o.SourceMicroversion, true
}

// HasSourceMicroversion returns a boolean if a field has been set.
func (o *BTFeatureDefinitionResponse1617) HasSourceMicroversion() bool {
	if o != nil && o.SourceMicroversion != nil {
		return true
	}

	return false
}

// SetSourceMicroversion gets a reference to the given string and assigns it to the SourceMicroversion field.
func (o *BTFeatureDefinitionResponse1617) SetSourceMicroversion(v string) {
	o.SourceMicroversion = &v
}

func (o BTFeatureDefinitionResponse1617) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Feature != nil {
		toSerialize["feature"] = o.Feature
	}
	if o.FeatureState != nil {
		toSerialize["featureState"] = o.FeatureState
	}
	if o.LibraryVersion != nil {
		toSerialize["libraryVersion"] = o.LibraryVersion
	}
	if o.MicroversionSkew != nil {
		toSerialize["microversionSkew"] = o.MicroversionSkew
	}
	if o.RejectMicroversionSkew != nil {
		toSerialize["rejectMicroversionSkew"] = o.RejectMicroversionSkew
	}
	if o.SerializationVersion != nil {
		toSerialize["serializationVersion"] = o.SerializationVersion
	}
	if o.SourceMicroversion != nil {
		toSerialize["sourceMicroversion"] = o.SourceMicroversion
	}
	return json.Marshal(toSerialize)
}

type NullableBTFeatureDefinitionResponse1617 struct {
	value *BTFeatureDefinitionResponse1617
	isSet bool
}

func (v NullableBTFeatureDefinitionResponse1617) Get() *BTFeatureDefinitionResponse1617 {
	return v.value
}

func (v *NullableBTFeatureDefinitionResponse1617) Set(val *BTFeatureDefinitionResponse1617) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFeatureDefinitionResponse1617) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFeatureDefinitionResponse1617) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFeatureDefinitionResponse1617(val *BTFeatureDefinitionResponse1617) *NullableBTFeatureDefinitionResponse1617 {
	return &NullableBTFeatureDefinitionResponse1617{value: val, isSet: true}
}

func (v NullableBTFeatureDefinitionResponse1617) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFeatureDefinitionResponse1617) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
