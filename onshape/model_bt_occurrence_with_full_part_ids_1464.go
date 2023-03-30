/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13685-0fb99d06bde5
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTOccurrenceWithFullPartIds1464 struct for BTOccurrenceWithFullPartIds1464
type BTOccurrenceWithFullPartIds1464 struct {
	BtType                *string                          `json:"btType,omitempty"`
	FullPathAsString      *string                          `json:"fullPathAsString,omitempty"`
	HeadInstanceId        *string                          `json:"headInstanceId,omitempty"`
	OccurrenceWithoutHead *BTOccurrence74                  `json:"occurrenceWithoutHead,omitempty"`
	OccurrenceWithoutTail *BTOccurrence74                  `json:"occurrenceWithoutTail,omitempty"`
	Parent                *BTOccurrence74                  `json:"parent,omitempty"`
	Path                  []string                         `json:"path,omitempty"`
	RootOccurrence        *bool                            `json:"rootOccurrence,omitempty"`
	TailInstanceId        *string                          `json:"tailInstanceId,omitempty"`
	FullElementId         *BTFullElementIdWithDocument1729 `json:"fullElementId,omitempty"`
	PartIds               []string                         `json:"partIds,omitempty"`
	Transform             *BTBSMatrix386                   `json:"transform,omitempty"`
}

// NewBTOccurrenceWithFullPartIds1464 instantiates a new BTOccurrenceWithFullPartIds1464 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTOccurrenceWithFullPartIds1464() *BTOccurrenceWithFullPartIds1464 {
	this := BTOccurrenceWithFullPartIds1464{}
	return &this
}

// NewBTOccurrenceWithFullPartIds1464WithDefaults instantiates a new BTOccurrenceWithFullPartIds1464 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTOccurrenceWithFullPartIds1464WithDefaults() *BTOccurrenceWithFullPartIds1464 {
	this := BTOccurrenceWithFullPartIds1464{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTOccurrenceWithFullPartIds1464) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceWithFullPartIds1464) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTOccurrenceWithFullPartIds1464) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTOccurrenceWithFullPartIds1464) SetBtType(v string) {
	o.BtType = &v
}

// GetFullPathAsString returns the FullPathAsString field value if set, zero value otherwise.
func (o *BTOccurrenceWithFullPartIds1464) GetFullPathAsString() string {
	if o == nil || o.FullPathAsString == nil {
		var ret string
		return ret
	}
	return *o.FullPathAsString
}

// GetFullPathAsStringOk returns a tuple with the FullPathAsString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceWithFullPartIds1464) GetFullPathAsStringOk() (*string, bool) {
	if o == nil || o.FullPathAsString == nil {
		return nil, false
	}
	return o.FullPathAsString, true
}

// HasFullPathAsString returns a boolean if a field has been set.
func (o *BTOccurrenceWithFullPartIds1464) HasFullPathAsString() bool {
	if o != nil && o.FullPathAsString != nil {
		return true
	}

	return false
}

// SetFullPathAsString gets a reference to the given string and assigns it to the FullPathAsString field.
func (o *BTOccurrenceWithFullPartIds1464) SetFullPathAsString(v string) {
	o.FullPathAsString = &v
}

// GetHeadInstanceId returns the HeadInstanceId field value if set, zero value otherwise.
func (o *BTOccurrenceWithFullPartIds1464) GetHeadInstanceId() string {
	if o == nil || o.HeadInstanceId == nil {
		var ret string
		return ret
	}
	return *o.HeadInstanceId
}

// GetHeadInstanceIdOk returns a tuple with the HeadInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceWithFullPartIds1464) GetHeadInstanceIdOk() (*string, bool) {
	if o == nil || o.HeadInstanceId == nil {
		return nil, false
	}
	return o.HeadInstanceId, true
}

// HasHeadInstanceId returns a boolean if a field has been set.
func (o *BTOccurrenceWithFullPartIds1464) HasHeadInstanceId() bool {
	if o != nil && o.HeadInstanceId != nil {
		return true
	}

	return false
}

// SetHeadInstanceId gets a reference to the given string and assigns it to the HeadInstanceId field.
func (o *BTOccurrenceWithFullPartIds1464) SetHeadInstanceId(v string) {
	o.HeadInstanceId = &v
}

// GetOccurrenceWithoutHead returns the OccurrenceWithoutHead field value if set, zero value otherwise.
func (o *BTOccurrenceWithFullPartIds1464) GetOccurrenceWithoutHead() BTOccurrence74 {
	if o == nil || o.OccurrenceWithoutHead == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.OccurrenceWithoutHead
}

// GetOccurrenceWithoutHeadOk returns a tuple with the OccurrenceWithoutHead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceWithFullPartIds1464) GetOccurrenceWithoutHeadOk() (*BTOccurrence74, bool) {
	if o == nil || o.OccurrenceWithoutHead == nil {
		return nil, false
	}
	return o.OccurrenceWithoutHead, true
}

// HasOccurrenceWithoutHead returns a boolean if a field has been set.
func (o *BTOccurrenceWithFullPartIds1464) HasOccurrenceWithoutHead() bool {
	if o != nil && o.OccurrenceWithoutHead != nil {
		return true
	}

	return false
}

// SetOccurrenceWithoutHead gets a reference to the given BTOccurrence74 and assigns it to the OccurrenceWithoutHead field.
func (o *BTOccurrenceWithFullPartIds1464) SetOccurrenceWithoutHead(v BTOccurrence74) {
	o.OccurrenceWithoutHead = &v
}

// GetOccurrenceWithoutTail returns the OccurrenceWithoutTail field value if set, zero value otherwise.
func (o *BTOccurrenceWithFullPartIds1464) GetOccurrenceWithoutTail() BTOccurrence74 {
	if o == nil || o.OccurrenceWithoutTail == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.OccurrenceWithoutTail
}

// GetOccurrenceWithoutTailOk returns a tuple with the OccurrenceWithoutTail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceWithFullPartIds1464) GetOccurrenceWithoutTailOk() (*BTOccurrence74, bool) {
	if o == nil || o.OccurrenceWithoutTail == nil {
		return nil, false
	}
	return o.OccurrenceWithoutTail, true
}

// HasOccurrenceWithoutTail returns a boolean if a field has been set.
func (o *BTOccurrenceWithFullPartIds1464) HasOccurrenceWithoutTail() bool {
	if o != nil && o.OccurrenceWithoutTail != nil {
		return true
	}

	return false
}

// SetOccurrenceWithoutTail gets a reference to the given BTOccurrence74 and assigns it to the OccurrenceWithoutTail field.
func (o *BTOccurrenceWithFullPartIds1464) SetOccurrenceWithoutTail(v BTOccurrence74) {
	o.OccurrenceWithoutTail = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *BTOccurrenceWithFullPartIds1464) GetParent() BTOccurrence74 {
	if o == nil || o.Parent == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceWithFullPartIds1464) GetParentOk() (*BTOccurrence74, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *BTOccurrenceWithFullPartIds1464) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given BTOccurrence74 and assigns it to the Parent field.
func (o *BTOccurrenceWithFullPartIds1464) SetParent(v BTOccurrence74) {
	o.Parent = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BTOccurrenceWithFullPartIds1464) GetPath() []string {
	if o == nil || o.Path == nil {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceWithFullPartIds1464) GetPathOk() ([]string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BTOccurrenceWithFullPartIds1464) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *BTOccurrenceWithFullPartIds1464) SetPath(v []string) {
	o.Path = v
}

// GetRootOccurrence returns the RootOccurrence field value if set, zero value otherwise.
func (o *BTOccurrenceWithFullPartIds1464) GetRootOccurrence() bool {
	if o == nil || o.RootOccurrence == nil {
		var ret bool
		return ret
	}
	return *o.RootOccurrence
}

// GetRootOccurrenceOk returns a tuple with the RootOccurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceWithFullPartIds1464) GetRootOccurrenceOk() (*bool, bool) {
	if o == nil || o.RootOccurrence == nil {
		return nil, false
	}
	return o.RootOccurrence, true
}

// HasRootOccurrence returns a boolean if a field has been set.
func (o *BTOccurrenceWithFullPartIds1464) HasRootOccurrence() bool {
	if o != nil && o.RootOccurrence != nil {
		return true
	}

	return false
}

// SetRootOccurrence gets a reference to the given bool and assigns it to the RootOccurrence field.
func (o *BTOccurrenceWithFullPartIds1464) SetRootOccurrence(v bool) {
	o.RootOccurrence = &v
}

// GetTailInstanceId returns the TailInstanceId field value if set, zero value otherwise.
func (o *BTOccurrenceWithFullPartIds1464) GetTailInstanceId() string {
	if o == nil || o.TailInstanceId == nil {
		var ret string
		return ret
	}
	return *o.TailInstanceId
}

// GetTailInstanceIdOk returns a tuple with the TailInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceWithFullPartIds1464) GetTailInstanceIdOk() (*string, bool) {
	if o == nil || o.TailInstanceId == nil {
		return nil, false
	}
	return o.TailInstanceId, true
}

// HasTailInstanceId returns a boolean if a field has been set.
func (o *BTOccurrenceWithFullPartIds1464) HasTailInstanceId() bool {
	if o != nil && o.TailInstanceId != nil {
		return true
	}

	return false
}

// SetTailInstanceId gets a reference to the given string and assigns it to the TailInstanceId field.
func (o *BTOccurrenceWithFullPartIds1464) SetTailInstanceId(v string) {
	o.TailInstanceId = &v
}

// GetFullElementId returns the FullElementId field value if set, zero value otherwise.
func (o *BTOccurrenceWithFullPartIds1464) GetFullElementId() BTFullElementIdWithDocument1729 {
	if o == nil || o.FullElementId == nil {
		var ret BTFullElementIdWithDocument1729
		return ret
	}
	return *o.FullElementId
}

// GetFullElementIdOk returns a tuple with the FullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceWithFullPartIds1464) GetFullElementIdOk() (*BTFullElementIdWithDocument1729, bool) {
	if o == nil || o.FullElementId == nil {
		return nil, false
	}
	return o.FullElementId, true
}

// HasFullElementId returns a boolean if a field has been set.
func (o *BTOccurrenceWithFullPartIds1464) HasFullElementId() bool {
	if o != nil && o.FullElementId != nil {
		return true
	}

	return false
}

// SetFullElementId gets a reference to the given BTFullElementIdWithDocument1729 and assigns it to the FullElementId field.
func (o *BTOccurrenceWithFullPartIds1464) SetFullElementId(v BTFullElementIdWithDocument1729) {
	o.FullElementId = &v
}

// GetPartIds returns the PartIds field value if set, zero value otherwise.
func (o *BTOccurrenceWithFullPartIds1464) GetPartIds() []string {
	if o == nil || o.PartIds == nil {
		var ret []string
		return ret
	}
	return o.PartIds
}

// GetPartIdsOk returns a tuple with the PartIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceWithFullPartIds1464) GetPartIdsOk() ([]string, bool) {
	if o == nil || o.PartIds == nil {
		return nil, false
	}
	return o.PartIds, true
}

// HasPartIds returns a boolean if a field has been set.
func (o *BTOccurrenceWithFullPartIds1464) HasPartIds() bool {
	if o != nil && o.PartIds != nil {
		return true
	}

	return false
}

// SetPartIds gets a reference to the given []string and assigns it to the PartIds field.
func (o *BTOccurrenceWithFullPartIds1464) SetPartIds(v []string) {
	o.PartIds = v
}

// GetTransform returns the Transform field value if set, zero value otherwise.
func (o *BTOccurrenceWithFullPartIds1464) GetTransform() BTBSMatrix386 {
	if o == nil || o.Transform == nil {
		var ret BTBSMatrix386
		return ret
	}
	return *o.Transform
}

// GetTransformOk returns a tuple with the Transform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceWithFullPartIds1464) GetTransformOk() (*BTBSMatrix386, bool) {
	if o == nil || o.Transform == nil {
		return nil, false
	}
	return o.Transform, true
}

// HasTransform returns a boolean if a field has been set.
func (o *BTOccurrenceWithFullPartIds1464) HasTransform() bool {
	if o != nil && o.Transform != nil {
		return true
	}

	return false
}

// SetTransform gets a reference to the given BTBSMatrix386 and assigns it to the Transform field.
func (o *BTOccurrenceWithFullPartIds1464) SetTransform(v BTBSMatrix386) {
	o.Transform = &v
}

func (o BTOccurrenceWithFullPartIds1464) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FullPathAsString != nil {
		toSerialize["fullPathAsString"] = o.FullPathAsString
	}
	if o.HeadInstanceId != nil {
		toSerialize["headInstanceId"] = o.HeadInstanceId
	}
	if o.OccurrenceWithoutHead != nil {
		toSerialize["occurrenceWithoutHead"] = o.OccurrenceWithoutHead
	}
	if o.OccurrenceWithoutTail != nil {
		toSerialize["occurrenceWithoutTail"] = o.OccurrenceWithoutTail
	}
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.RootOccurrence != nil {
		toSerialize["rootOccurrence"] = o.RootOccurrence
	}
	if o.TailInstanceId != nil {
		toSerialize["tailInstanceId"] = o.TailInstanceId
	}
	if o.FullElementId != nil {
		toSerialize["fullElementId"] = o.FullElementId
	}
	if o.PartIds != nil {
		toSerialize["partIds"] = o.PartIds
	}
	if o.Transform != nil {
		toSerialize["transform"] = o.Transform
	}
	return json.Marshal(toSerialize)
}

type NullableBTOccurrenceWithFullPartIds1464 struct {
	value *BTOccurrenceWithFullPartIds1464
	isSet bool
}

func (v NullableBTOccurrenceWithFullPartIds1464) Get() *BTOccurrenceWithFullPartIds1464 {
	return v.value
}

func (v *NullableBTOccurrenceWithFullPartIds1464) Set(val *BTOccurrenceWithFullPartIds1464) {
	v.value = val
	v.isSet = true
}

func (v NullableBTOccurrenceWithFullPartIds1464) IsSet() bool {
	return v.isSet
}

func (v *NullableBTOccurrenceWithFullPartIds1464) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTOccurrenceWithFullPartIds1464(val *BTOccurrenceWithFullPartIds1464) *NullableBTOccurrenceWithFullPartIds1464 {
	return &NullableBTOccurrenceWithFullPartIds1464{value: val, isSet: true}
}

func (v NullableBTOccurrenceWithFullPartIds1464) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTOccurrenceWithFullPartIds1464) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
