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

// BTMateGroupDisplayData1990 struct for BTMateGroupDisplayData1990
type BTMateGroupDisplayData1990 struct {
	BtType           *string                          `json:"btType,omitempty"`
	Hidden           *bool                            `json:"hidden,omitempty"`
	IsDerivedFeature *bool                            `json:"isDerivedFeature,omitempty"`
	NodeId           *string                          `json:"nodeId,omitempty"`
	OccurrenceIds    []string                         `json:"occurrenceIds,omitempty"`
	OwnerOccurrence  *BTOccurrence74                  `json:"ownerOccurrence,omitempty"`
	Status           *GBTAssemblyFeatureDisplayStatus `json:"status,omitempty"`
}

// NewBTMateGroupDisplayData1990 instantiates a new BTMateGroupDisplayData1990 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMateGroupDisplayData1990() *BTMateGroupDisplayData1990 {
	this := BTMateGroupDisplayData1990{}
	return &this
}

// NewBTMateGroupDisplayData1990WithDefaults instantiates a new BTMateGroupDisplayData1990 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMateGroupDisplayData1990WithDefaults() *BTMateGroupDisplayData1990 {
	this := BTMateGroupDisplayData1990{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMateGroupDisplayData1990) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateGroupDisplayData1990) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMateGroupDisplayData1990) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMateGroupDisplayData1990) SetBtType(v string) {
	o.BtType = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *BTMateGroupDisplayData1990) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateGroupDisplayData1990) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *BTMateGroupDisplayData1990) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *BTMateGroupDisplayData1990) SetHidden(v bool) {
	o.Hidden = &v
}

// GetIsDerivedFeature returns the IsDerivedFeature field value if set, zero value otherwise.
func (o *BTMateGroupDisplayData1990) GetIsDerivedFeature() bool {
	if o == nil || o.IsDerivedFeature == nil {
		var ret bool
		return ret
	}
	return *o.IsDerivedFeature
}

// GetIsDerivedFeatureOk returns a tuple with the IsDerivedFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateGroupDisplayData1990) GetIsDerivedFeatureOk() (*bool, bool) {
	if o == nil || o.IsDerivedFeature == nil {
		return nil, false
	}
	return o.IsDerivedFeature, true
}

// HasIsDerivedFeature returns a boolean if a field has been set.
func (o *BTMateGroupDisplayData1990) HasIsDerivedFeature() bool {
	if o != nil && o.IsDerivedFeature != nil {
		return true
	}

	return false
}

// SetIsDerivedFeature gets a reference to the given bool and assigns it to the IsDerivedFeature field.
func (o *BTMateGroupDisplayData1990) SetIsDerivedFeature(v bool) {
	o.IsDerivedFeature = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMateGroupDisplayData1990) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateGroupDisplayData1990) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMateGroupDisplayData1990) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMateGroupDisplayData1990) SetNodeId(v string) {
	o.NodeId = &v
}

// GetOccurrenceIds returns the OccurrenceIds field value if set, zero value otherwise.
func (o *BTMateGroupDisplayData1990) GetOccurrenceIds() []string {
	if o == nil || o.OccurrenceIds == nil {
		var ret []string
		return ret
	}
	return o.OccurrenceIds
}

// GetOccurrenceIdsOk returns a tuple with the OccurrenceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateGroupDisplayData1990) GetOccurrenceIdsOk() ([]string, bool) {
	if o == nil || o.OccurrenceIds == nil {
		return nil, false
	}
	return o.OccurrenceIds, true
}

// HasOccurrenceIds returns a boolean if a field has been set.
func (o *BTMateGroupDisplayData1990) HasOccurrenceIds() bool {
	if o != nil && o.OccurrenceIds != nil {
		return true
	}

	return false
}

// SetOccurrenceIds gets a reference to the given []string and assigns it to the OccurrenceIds field.
func (o *BTMateGroupDisplayData1990) SetOccurrenceIds(v []string) {
	o.OccurrenceIds = v
}

// GetOwnerOccurrence returns the OwnerOccurrence field value if set, zero value otherwise.
func (o *BTMateGroupDisplayData1990) GetOwnerOccurrence() BTOccurrence74 {
	if o == nil || o.OwnerOccurrence == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.OwnerOccurrence
}

// GetOwnerOccurrenceOk returns a tuple with the OwnerOccurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateGroupDisplayData1990) GetOwnerOccurrenceOk() (*BTOccurrence74, bool) {
	if o == nil || o.OwnerOccurrence == nil {
		return nil, false
	}
	return o.OwnerOccurrence, true
}

// HasOwnerOccurrence returns a boolean if a field has been set.
func (o *BTMateGroupDisplayData1990) HasOwnerOccurrence() bool {
	if o != nil && o.OwnerOccurrence != nil {
		return true
	}

	return false
}

// SetOwnerOccurrence gets a reference to the given BTOccurrence74 and assigns it to the OwnerOccurrence field.
func (o *BTMateGroupDisplayData1990) SetOwnerOccurrence(v BTOccurrence74) {
	o.OwnerOccurrence = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BTMateGroupDisplayData1990) GetStatus() GBTAssemblyFeatureDisplayStatus {
	if o == nil || o.Status == nil {
		var ret GBTAssemblyFeatureDisplayStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateGroupDisplayData1990) GetStatusOk() (*GBTAssemblyFeatureDisplayStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BTMateGroupDisplayData1990) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GBTAssemblyFeatureDisplayStatus and assigns it to the Status field.
func (o *BTMateGroupDisplayData1990) SetStatus(v GBTAssemblyFeatureDisplayStatus) {
	o.Status = &v
}

func (o BTMateGroupDisplayData1990) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	if o.IsDerivedFeature != nil {
		toSerialize["isDerivedFeature"] = o.IsDerivedFeature
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.OccurrenceIds != nil {
		toSerialize["occurrenceIds"] = o.OccurrenceIds
	}
	if o.OwnerOccurrence != nil {
		toSerialize["ownerOccurrence"] = o.OwnerOccurrence
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableBTMateGroupDisplayData1990 struct {
	value *BTMateGroupDisplayData1990
	isSet bool
}

func (v NullableBTMateGroupDisplayData1990) Get() *BTMateGroupDisplayData1990 {
	return v.value
}

func (v *NullableBTMateGroupDisplayData1990) Set(val *BTMateGroupDisplayData1990) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMateGroupDisplayData1990) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMateGroupDisplayData1990) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMateGroupDisplayData1990(val *BTMateGroupDisplayData1990) *NullableBTMateGroupDisplayData1990 {
	return &NullableBTMateGroupDisplayData1990{value: val, isSet: true}
}

func (v NullableBTMateGroupDisplayData1990) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMateGroupDisplayData1990) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
