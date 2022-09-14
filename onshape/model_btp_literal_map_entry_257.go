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
)

// BTPLiteralMapEntry257 struct for BTPLiteralMapEntry257
type BTPLiteralMapEntry257 struct {
	Atomic              *bool                  `json:"atomic,omitempty"`
	BtType              *string                `json:"btType,omitempty"`
	DocumentationType   *string                `json:"documentationType,omitempty"`
	EndSourceLocation   *int32                 `json:"endSourceLocation,omitempty"`
	Key                 *BTPPropertyAccessor23 `json:"key,omitempty"`
	NodeId              *string                `json:"nodeId,omitempty"`
	ShortDescriptor     *string                `json:"shortDescriptor,omitempty"`
	SpaceAfter          *BTPSpace10            `json:"spaceAfter,omitempty"`
	SpaceBefore         *BTPSpace10            `json:"spaceBefore,omitempty"`
	SpaceDefault        *bool                  `json:"spaceDefault,omitempty"`
	StartSourceLocation *int32                 `json:"startSourceLocation,omitempty"`
	Value               *BTPExpression9        `json:"value,omitempty"`
}

// NewBTPLiteralMapEntry257 instantiates a new BTPLiteralMapEntry257 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPLiteralMapEntry257() *BTPLiteralMapEntry257 {
	this := BTPLiteralMapEntry257{}
	return &this
}

// NewBTPLiteralMapEntry257WithDefaults instantiates a new BTPLiteralMapEntry257 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPLiteralMapEntry257WithDefaults() *BTPLiteralMapEntry257 {
	this := BTPLiteralMapEntry257{}
	return &this
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *BTPLiteralMapEntry257) GetAtomic() bool {
	if o == nil || o.Atomic == nil {
		var ret bool
		return ret
	}
	return *o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLiteralMapEntry257) GetAtomicOk() (*bool, bool) {
	if o == nil || o.Atomic == nil {
		return nil, false
	}
	return o.Atomic, true
}

// HasAtomic returns a boolean if a field has been set.
func (o *BTPLiteralMapEntry257) HasAtomic() bool {
	if o != nil && o.Atomic != nil {
		return true
	}

	return false
}

// SetAtomic gets a reference to the given bool and assigns it to the Atomic field.
func (o *BTPLiteralMapEntry257) SetAtomic(v bool) {
	o.Atomic = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPLiteralMapEntry257) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLiteralMapEntry257) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPLiteralMapEntry257) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPLiteralMapEntry257) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentationType returns the DocumentationType field value if set, zero value otherwise.
func (o *BTPLiteralMapEntry257) GetDocumentationType() string {
	if o == nil || o.DocumentationType == nil {
		var ret string
		return ret
	}
	return *o.DocumentationType
}

// GetDocumentationTypeOk returns a tuple with the DocumentationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLiteralMapEntry257) GetDocumentationTypeOk() (*string, bool) {
	if o == nil || o.DocumentationType == nil {
		return nil, false
	}
	return o.DocumentationType, true
}

// HasDocumentationType returns a boolean if a field has been set.
func (o *BTPLiteralMapEntry257) HasDocumentationType() bool {
	if o != nil && o.DocumentationType != nil {
		return true
	}

	return false
}

// SetDocumentationType gets a reference to the given string and assigns it to the DocumentationType field.
func (o *BTPLiteralMapEntry257) SetDocumentationType(v string) {
	o.DocumentationType = &v
}

// GetEndSourceLocation returns the EndSourceLocation field value if set, zero value otherwise.
func (o *BTPLiteralMapEntry257) GetEndSourceLocation() int32 {
	if o == nil || o.EndSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.EndSourceLocation
}

// GetEndSourceLocationOk returns a tuple with the EndSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLiteralMapEntry257) GetEndSourceLocationOk() (*int32, bool) {
	if o == nil || o.EndSourceLocation == nil {
		return nil, false
	}
	return o.EndSourceLocation, true
}

// HasEndSourceLocation returns a boolean if a field has been set.
func (o *BTPLiteralMapEntry257) HasEndSourceLocation() bool {
	if o != nil && o.EndSourceLocation != nil {
		return true
	}

	return false
}

// SetEndSourceLocation gets a reference to the given int32 and assigns it to the EndSourceLocation field.
func (o *BTPLiteralMapEntry257) SetEndSourceLocation(v int32) {
	o.EndSourceLocation = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *BTPLiteralMapEntry257) GetKey() BTPPropertyAccessor23 {
	if o == nil || o.Key == nil {
		var ret BTPPropertyAccessor23
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLiteralMapEntry257) GetKeyOk() (*BTPPropertyAccessor23, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *BTPLiteralMapEntry257) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given BTPPropertyAccessor23 and assigns it to the Key field.
func (o *BTPLiteralMapEntry257) SetKey(v BTPPropertyAccessor23) {
	o.Key = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPLiteralMapEntry257) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLiteralMapEntry257) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPLiteralMapEntry257) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPLiteralMapEntry257) SetNodeId(v string) {
	o.NodeId = &v
}

// GetShortDescriptor returns the ShortDescriptor field value if set, zero value otherwise.
func (o *BTPLiteralMapEntry257) GetShortDescriptor() string {
	if o == nil || o.ShortDescriptor == nil {
		var ret string
		return ret
	}
	return *o.ShortDescriptor
}

// GetShortDescriptorOk returns a tuple with the ShortDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLiteralMapEntry257) GetShortDescriptorOk() (*string, bool) {
	if o == nil || o.ShortDescriptor == nil {
		return nil, false
	}
	return o.ShortDescriptor, true
}

// HasShortDescriptor returns a boolean if a field has been set.
func (o *BTPLiteralMapEntry257) HasShortDescriptor() bool {
	if o != nil && o.ShortDescriptor != nil {
		return true
	}

	return false
}

// SetShortDescriptor gets a reference to the given string and assigns it to the ShortDescriptor field.
func (o *BTPLiteralMapEntry257) SetShortDescriptor(v string) {
	o.ShortDescriptor = &v
}

// GetSpaceAfter returns the SpaceAfter field value if set, zero value otherwise.
func (o *BTPLiteralMapEntry257) GetSpaceAfter() BTPSpace10 {
	if o == nil || o.SpaceAfter == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfter
}

// GetSpaceAfterOk returns a tuple with the SpaceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLiteralMapEntry257) GetSpaceAfterOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfter == nil {
		return nil, false
	}
	return o.SpaceAfter, true
}

// HasSpaceAfter returns a boolean if a field has been set.
func (o *BTPLiteralMapEntry257) HasSpaceAfter() bool {
	if o != nil && o.SpaceAfter != nil {
		return true
	}

	return false
}

// SetSpaceAfter gets a reference to the given BTPSpace10 and assigns it to the SpaceAfter field.
func (o *BTPLiteralMapEntry257) SetSpaceAfter(v BTPSpace10) {
	o.SpaceAfter = &v
}

// GetSpaceBefore returns the SpaceBefore field value if set, zero value otherwise.
func (o *BTPLiteralMapEntry257) GetSpaceBefore() BTPSpace10 {
	if o == nil || o.SpaceBefore == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBefore
}

// GetSpaceBeforeOk returns a tuple with the SpaceBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLiteralMapEntry257) GetSpaceBeforeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBefore == nil {
		return nil, false
	}
	return o.SpaceBefore, true
}

// HasSpaceBefore returns a boolean if a field has been set.
func (o *BTPLiteralMapEntry257) HasSpaceBefore() bool {
	if o != nil && o.SpaceBefore != nil {
		return true
	}

	return false
}

// SetSpaceBefore gets a reference to the given BTPSpace10 and assigns it to the SpaceBefore field.
func (o *BTPLiteralMapEntry257) SetSpaceBefore(v BTPSpace10) {
	o.SpaceBefore = &v
}

// GetSpaceDefault returns the SpaceDefault field value if set, zero value otherwise.
func (o *BTPLiteralMapEntry257) GetSpaceDefault() bool {
	if o == nil || o.SpaceDefault == nil {
		var ret bool
		return ret
	}
	return *o.SpaceDefault
}

// GetSpaceDefaultOk returns a tuple with the SpaceDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLiteralMapEntry257) GetSpaceDefaultOk() (*bool, bool) {
	if o == nil || o.SpaceDefault == nil {
		return nil, false
	}
	return o.SpaceDefault, true
}

// HasSpaceDefault returns a boolean if a field has been set.
func (o *BTPLiteralMapEntry257) HasSpaceDefault() bool {
	if o != nil && o.SpaceDefault != nil {
		return true
	}

	return false
}

// SetSpaceDefault gets a reference to the given bool and assigns it to the SpaceDefault field.
func (o *BTPLiteralMapEntry257) SetSpaceDefault(v bool) {
	o.SpaceDefault = &v
}

// GetStartSourceLocation returns the StartSourceLocation field value if set, zero value otherwise.
func (o *BTPLiteralMapEntry257) GetStartSourceLocation() int32 {
	if o == nil || o.StartSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.StartSourceLocation
}

// GetStartSourceLocationOk returns a tuple with the StartSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLiteralMapEntry257) GetStartSourceLocationOk() (*int32, bool) {
	if o == nil || o.StartSourceLocation == nil {
		return nil, false
	}
	return o.StartSourceLocation, true
}

// HasStartSourceLocation returns a boolean if a field has been set.
func (o *BTPLiteralMapEntry257) HasStartSourceLocation() bool {
	if o != nil && o.StartSourceLocation != nil {
		return true
	}

	return false
}

// SetStartSourceLocation gets a reference to the given int32 and assigns it to the StartSourceLocation field.
func (o *BTPLiteralMapEntry257) SetStartSourceLocation(v int32) {
	o.StartSourceLocation = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTPLiteralMapEntry257) GetValue() BTPExpression9 {
	if o == nil || o.Value == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLiteralMapEntry257) GetValueOk() (*BTPExpression9, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTPLiteralMapEntry257) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given BTPExpression9 and assigns it to the Value field.
func (o *BTPLiteralMapEntry257) SetValue(v BTPExpression9) {
	o.Value = &v
}

func (o BTPLiteralMapEntry257) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Atomic != nil {
		toSerialize["atomic"] = o.Atomic
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DocumentationType != nil {
		toSerialize["documentationType"] = o.DocumentationType
	}
	if o.EndSourceLocation != nil {
		toSerialize["endSourceLocation"] = o.EndSourceLocation
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ShortDescriptor != nil {
		toSerialize["shortDescriptor"] = o.ShortDescriptor
	}
	if o.SpaceAfter != nil {
		toSerialize["spaceAfter"] = o.SpaceAfter
	}
	if o.SpaceBefore != nil {
		toSerialize["spaceBefore"] = o.SpaceBefore
	}
	if o.SpaceDefault != nil {
		toSerialize["spaceDefault"] = o.SpaceDefault
	}
	if o.StartSourceLocation != nil {
		toSerialize["startSourceLocation"] = o.StartSourceLocation
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTPLiteralMapEntry257 struct {
	value *BTPLiteralMapEntry257
	isSet bool
}

func (v NullableBTPLiteralMapEntry257) Get() *BTPLiteralMapEntry257 {
	return v.value
}

func (v *NullableBTPLiteralMapEntry257) Set(val *BTPLiteralMapEntry257) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPLiteralMapEntry257) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPLiteralMapEntry257) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPLiteralMapEntry257(val *BTPLiteralMapEntry257) *NullableBTPLiteralMapEntry257 {
	return &NullableBTPLiteralMapEntry257{value: val, isSet: true}
}

func (v NullableBTPLiteralMapEntry257) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPLiteralMapEntry257) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
