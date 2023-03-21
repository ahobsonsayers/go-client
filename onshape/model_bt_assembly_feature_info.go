/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13200-ff216a970a02
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAssemblyFeatureInfo List of Assembly features including those are created by replicates.
type BTAssemblyFeatureInfo struct {
	FeatureData *BTAssemblyFeatureDataInfo `json:"featureData,omitempty"`
	FeatureType *string                    `json:"featureType,omitempty"`
	Id          *string                    `json:"id,omitempty"`
	Suppressed  *bool                      `json:"suppressed,omitempty"`
}

// NewBTAssemblyFeatureInfo instantiates a new BTAssemblyFeatureInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyFeatureInfo() *BTAssemblyFeatureInfo {
	this := BTAssemblyFeatureInfo{}
	return &this
}

// NewBTAssemblyFeatureInfoWithDefaults instantiates a new BTAssemblyFeatureInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyFeatureInfoWithDefaults() *BTAssemblyFeatureInfo {
	this := BTAssemblyFeatureInfo{}
	return &this
}

// GetFeatureData returns the FeatureData field value if set, zero value otherwise.
func (o *BTAssemblyFeatureInfo) GetFeatureData() BTAssemblyFeatureDataInfo {
	if o == nil || o.FeatureData == nil {
		var ret BTAssemblyFeatureDataInfo
		return ret
	}
	return *o.FeatureData
}

// GetFeatureDataOk returns a tuple with the FeatureData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureInfo) GetFeatureDataOk() (*BTAssemblyFeatureDataInfo, bool) {
	if o == nil || o.FeatureData == nil {
		return nil, false
	}
	return o.FeatureData, true
}

// HasFeatureData returns a boolean if a field has been set.
func (o *BTAssemblyFeatureInfo) HasFeatureData() bool {
	if o != nil && o.FeatureData != nil {
		return true
	}

	return false
}

// SetFeatureData gets a reference to the given BTAssemblyFeatureDataInfo and assigns it to the FeatureData field.
func (o *BTAssemblyFeatureInfo) SetFeatureData(v BTAssemblyFeatureDataInfo) {
	o.FeatureData = &v
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise.
func (o *BTAssemblyFeatureInfo) GetFeatureType() string {
	if o == nil || o.FeatureType == nil {
		var ret string
		return ret
	}
	return *o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureInfo) GetFeatureTypeOk() (*string, bool) {
	if o == nil || o.FeatureType == nil {
		return nil, false
	}
	return o.FeatureType, true
}

// HasFeatureType returns a boolean if a field has been set.
func (o *BTAssemblyFeatureInfo) HasFeatureType() bool {
	if o != nil && o.FeatureType != nil {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given string and assigns it to the FeatureType field.
func (o *BTAssemblyFeatureInfo) SetFeatureType(v string) {
	o.FeatureType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTAssemblyFeatureInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTAssemblyFeatureInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTAssemblyFeatureInfo) SetId(v string) {
	o.Id = &v
}

// GetSuppressed returns the Suppressed field value if set, zero value otherwise.
func (o *BTAssemblyFeatureInfo) GetSuppressed() bool {
	if o == nil || o.Suppressed == nil {
		var ret bool
		return ret
	}
	return *o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureInfo) GetSuppressedOk() (*bool, bool) {
	if o == nil || o.Suppressed == nil {
		return nil, false
	}
	return o.Suppressed, true
}

// HasSuppressed returns a boolean if a field has been set.
func (o *BTAssemblyFeatureInfo) HasSuppressed() bool {
	if o != nil && o.Suppressed != nil {
		return true
	}

	return false
}

// SetSuppressed gets a reference to the given bool and assigns it to the Suppressed field.
func (o *BTAssemblyFeatureInfo) SetSuppressed(v bool) {
	o.Suppressed = &v
}

func (o BTAssemblyFeatureInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FeatureData != nil {
		toSerialize["featureData"] = o.FeatureData
	}
	if o.FeatureType != nil {
		toSerialize["featureType"] = o.FeatureType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Suppressed != nil {
		toSerialize["suppressed"] = o.Suppressed
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyFeatureInfo struct {
	value *BTAssemblyFeatureInfo
	isSet bool
}

func (v NullableBTAssemblyFeatureInfo) Get() *BTAssemblyFeatureInfo {
	return v.value
}

func (v *NullableBTAssemblyFeatureInfo) Set(val *BTAssemblyFeatureInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyFeatureInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyFeatureInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyFeatureInfo(val *BTAssemblyFeatureInfo) *NullableBTAssemblyFeatureInfo {
	return &NullableBTAssemblyFeatureInfo{value: val, isSet: true}
}

func (v NullableBTAssemblyFeatureInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyFeatureInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
