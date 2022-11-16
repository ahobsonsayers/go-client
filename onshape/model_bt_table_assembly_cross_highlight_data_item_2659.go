/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7489-3fe01473c812
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTTableAssemblyCrossHighlightDataItem2659 struct for BTTableAssemblyCrossHighlightDataItem2659
type BTTableAssemblyCrossHighlightDataItem2659 struct {
	BtType           *string `json:"btType,omitempty"`
	OccurrencePathId *string `json:"occurrencePathId,omitempty"`
}

// NewBTTableAssemblyCrossHighlightDataItem2659 instantiates a new BTTableAssemblyCrossHighlightDataItem2659 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableAssemblyCrossHighlightDataItem2659() *BTTableAssemblyCrossHighlightDataItem2659 {
	this := BTTableAssemblyCrossHighlightDataItem2659{}
	return &this
}

// NewBTTableAssemblyCrossHighlightDataItem2659WithDefaults instantiates a new BTTableAssemblyCrossHighlightDataItem2659 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableAssemblyCrossHighlightDataItem2659WithDefaults() *BTTableAssemblyCrossHighlightDataItem2659 {
	this := BTTableAssemblyCrossHighlightDataItem2659{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTableAssemblyCrossHighlightDataItem2659) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableAssemblyCrossHighlightDataItem2659) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTableAssemblyCrossHighlightDataItem2659) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTableAssemblyCrossHighlightDataItem2659) SetBtType(v string) {
	o.BtType = &v
}

// GetOccurrencePathId returns the OccurrencePathId field value if set, zero value otherwise.
func (o *BTTableAssemblyCrossHighlightDataItem2659) GetOccurrencePathId() string {
	if o == nil || o.OccurrencePathId == nil {
		var ret string
		return ret
	}
	return *o.OccurrencePathId
}

// GetOccurrencePathIdOk returns a tuple with the OccurrencePathId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableAssemblyCrossHighlightDataItem2659) GetOccurrencePathIdOk() (*string, bool) {
	if o == nil || o.OccurrencePathId == nil {
		return nil, false
	}
	return o.OccurrencePathId, true
}

// HasOccurrencePathId returns a boolean if a field has been set.
func (o *BTTableAssemblyCrossHighlightDataItem2659) HasOccurrencePathId() bool {
	if o != nil && o.OccurrencePathId != nil {
		return true
	}

	return false
}

// SetOccurrencePathId gets a reference to the given string and assigns it to the OccurrencePathId field.
func (o *BTTableAssemblyCrossHighlightDataItem2659) SetOccurrencePathId(v string) {
	o.OccurrencePathId = &v
}

func (o BTTableAssemblyCrossHighlightDataItem2659) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.OccurrencePathId != nil {
		toSerialize["occurrencePathId"] = o.OccurrencePathId
	}
	return json.Marshal(toSerialize)
}

type NullableBTTableAssemblyCrossHighlightDataItem2659 struct {
	value *BTTableAssemblyCrossHighlightDataItem2659
	isSet bool
}

func (v NullableBTTableAssemblyCrossHighlightDataItem2659) Get() *BTTableAssemblyCrossHighlightDataItem2659 {
	return v.value
}

func (v *NullableBTTableAssemblyCrossHighlightDataItem2659) Set(val *BTTableAssemblyCrossHighlightDataItem2659) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableAssemblyCrossHighlightDataItem2659) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableAssemblyCrossHighlightDataItem2659) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableAssemblyCrossHighlightDataItem2659(val *BTTableAssemblyCrossHighlightDataItem2659) *NullableBTTableAssemblyCrossHighlightDataItem2659 {
	return &NullableBTTableAssemblyCrossHighlightDataItem2659{value: val, isSet: true}
}

func (v NullableBTTableAssemblyCrossHighlightDataItem2659) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableAssemblyCrossHighlightDataItem2659) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
