/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6290-b55936bc8c5a
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAppElementContentEntryInfo struct for BTAppElementContentEntryInfo
type BTAppElementContentEntryInfo struct {
	BaseContent  *string                        `json:"baseContent,omitempty"`
	Deltas       []BTAppElementContentDeltaInfo `json:"deltas,omitempty"`
	SubelementId *string                        `json:"subelementId,omitempty"`
}

// NewBTAppElementContentEntryInfo instantiates a new BTAppElementContentEntryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementContentEntryInfo() *BTAppElementContentEntryInfo {
	this := BTAppElementContentEntryInfo{}
	return &this
}

// NewBTAppElementContentEntryInfoWithDefaults instantiates a new BTAppElementContentEntryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementContentEntryInfoWithDefaults() *BTAppElementContentEntryInfo {
	this := BTAppElementContentEntryInfo{}
	return &this
}

// GetBaseContent returns the BaseContent field value if set, zero value otherwise.
func (o *BTAppElementContentEntryInfo) GetBaseContent() string {
	if o == nil || o.BaseContent == nil {
		var ret string
		return ret
	}
	return *o.BaseContent
}

// GetBaseContentOk returns a tuple with the BaseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementContentEntryInfo) GetBaseContentOk() (*string, bool) {
	if o == nil || o.BaseContent == nil {
		return nil, false
	}
	return o.BaseContent, true
}

// HasBaseContent returns a boolean if a field has been set.
func (o *BTAppElementContentEntryInfo) HasBaseContent() bool {
	if o != nil && o.BaseContent != nil {
		return true
	}

	return false
}

// SetBaseContent gets a reference to the given string and assigns it to the BaseContent field.
func (o *BTAppElementContentEntryInfo) SetBaseContent(v string) {
	o.BaseContent = &v
}

// GetDeltas returns the Deltas field value if set, zero value otherwise.
func (o *BTAppElementContentEntryInfo) GetDeltas() []BTAppElementContentDeltaInfo {
	if o == nil || o.Deltas == nil {
		var ret []BTAppElementContentDeltaInfo
		return ret
	}
	return o.Deltas
}

// GetDeltasOk returns a tuple with the Deltas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementContentEntryInfo) GetDeltasOk() ([]BTAppElementContentDeltaInfo, bool) {
	if o == nil || o.Deltas == nil {
		return nil, false
	}
	return o.Deltas, true
}

// HasDeltas returns a boolean if a field has been set.
func (o *BTAppElementContentEntryInfo) HasDeltas() bool {
	if o != nil && o.Deltas != nil {
		return true
	}

	return false
}

// SetDeltas gets a reference to the given []BTAppElementContentDeltaInfo and assigns it to the Deltas field.
func (o *BTAppElementContentEntryInfo) SetDeltas(v []BTAppElementContentDeltaInfo) {
	o.Deltas = v
}

// GetSubelementId returns the SubelementId field value if set, zero value otherwise.
func (o *BTAppElementContentEntryInfo) GetSubelementId() string {
	if o == nil || o.SubelementId == nil {
		var ret string
		return ret
	}
	return *o.SubelementId
}

// GetSubelementIdOk returns a tuple with the SubelementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementContentEntryInfo) GetSubelementIdOk() (*string, bool) {
	if o == nil || o.SubelementId == nil {
		return nil, false
	}
	return o.SubelementId, true
}

// HasSubelementId returns a boolean if a field has been set.
func (o *BTAppElementContentEntryInfo) HasSubelementId() bool {
	if o != nil && o.SubelementId != nil {
		return true
	}

	return false
}

// SetSubelementId gets a reference to the given string and assigns it to the SubelementId field.
func (o *BTAppElementContentEntryInfo) SetSubelementId(v string) {
	o.SubelementId = &v
}

func (o BTAppElementContentEntryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BaseContent != nil {
		toSerialize["baseContent"] = o.BaseContent
	}
	if o.Deltas != nil {
		toSerialize["deltas"] = o.Deltas
	}
	if o.SubelementId != nil {
		toSerialize["subelementId"] = o.SubelementId
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementContentEntryInfo struct {
	value *BTAppElementContentEntryInfo
	isSet bool
}

func (v NullableBTAppElementContentEntryInfo) Get() *BTAppElementContentEntryInfo {
	return v.value
}

func (v *NullableBTAppElementContentEntryInfo) Set(val *BTAppElementContentEntryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementContentEntryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementContentEntryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementContentEntryInfo(val *BTAppElementContentEntryInfo) *NullableBTAppElementContentEntryInfo {
	return &NullableBTAppElementContentEntryInfo{value: val, isSet: true}
}

func (v NullableBTAppElementContentEntryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementContentEntryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
