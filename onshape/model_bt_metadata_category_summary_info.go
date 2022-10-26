/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7070-d84f850e97a2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMetadataCategorySummaryInfo struct for BTMetadataCategorySummaryInfo
type BTMetadataCategorySummaryInfo struct {
	DefaultObjectType *int32  `json:"defaultObjectType,omitempty"`
	Description       *string `json:"description,omitempty"`
	Href              *string `json:"href,omitempty"`
	Id                *string `json:"id,omitempty"`
	Name              *string `json:"name,omitempty"`
	ObjectTypes       []int32 `json:"objectTypes,omitempty"`
	OwnerId           *string `json:"ownerId,omitempty"`
	OwnerType         *int32  `json:"ownerType,omitempty"`
	PublishState      *int32  `json:"publishState,omitempty"`
	ViewRef           *string `json:"viewRef,omitempty"`
}

// NewBTMetadataCategorySummaryInfo instantiates a new BTMetadataCategorySummaryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMetadataCategorySummaryInfo() *BTMetadataCategorySummaryInfo {
	this := BTMetadataCategorySummaryInfo{}
	return &this
}

// NewBTMetadataCategorySummaryInfoWithDefaults instantiates a new BTMetadataCategorySummaryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMetadataCategorySummaryInfoWithDefaults() *BTMetadataCategorySummaryInfo {
	this := BTMetadataCategorySummaryInfo{}
	return &this
}

// GetDefaultObjectType returns the DefaultObjectType field value if set, zero value otherwise.
func (o *BTMetadataCategorySummaryInfo) GetDefaultObjectType() int32 {
	if o == nil || o.DefaultObjectType == nil {
		var ret int32
		return ret
	}
	return *o.DefaultObjectType
}

// GetDefaultObjectTypeOk returns a tuple with the DefaultObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCategorySummaryInfo) GetDefaultObjectTypeOk() (*int32, bool) {
	if o == nil || o.DefaultObjectType == nil {
		return nil, false
	}
	return o.DefaultObjectType, true
}

// HasDefaultObjectType returns a boolean if a field has been set.
func (o *BTMetadataCategorySummaryInfo) HasDefaultObjectType() bool {
	if o != nil && o.DefaultObjectType != nil {
		return true
	}

	return false
}

// SetDefaultObjectType gets a reference to the given int32 and assigns it to the DefaultObjectType field.
func (o *BTMetadataCategorySummaryInfo) SetDefaultObjectType(v int32) {
	o.DefaultObjectType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTMetadataCategorySummaryInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCategorySummaryInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTMetadataCategorySummaryInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTMetadataCategorySummaryInfo) SetDescription(v string) {
	o.Description = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTMetadataCategorySummaryInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCategorySummaryInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTMetadataCategorySummaryInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTMetadataCategorySummaryInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTMetadataCategorySummaryInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCategorySummaryInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTMetadataCategorySummaryInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTMetadataCategorySummaryInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTMetadataCategorySummaryInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCategorySummaryInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTMetadataCategorySummaryInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTMetadataCategorySummaryInfo) SetName(v string) {
	o.Name = &v
}

// GetObjectTypes returns the ObjectTypes field value if set, zero value otherwise.
func (o *BTMetadataCategorySummaryInfo) GetObjectTypes() []int32 {
	if o == nil || o.ObjectTypes == nil {
		var ret []int32
		return ret
	}
	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCategorySummaryInfo) GetObjectTypesOk() ([]int32, bool) {
	if o == nil || o.ObjectTypes == nil {
		return nil, false
	}
	return o.ObjectTypes, true
}

// HasObjectTypes returns a boolean if a field has been set.
func (o *BTMetadataCategorySummaryInfo) HasObjectTypes() bool {
	if o != nil && o.ObjectTypes != nil {
		return true
	}

	return false
}

// SetObjectTypes gets a reference to the given []int32 and assigns it to the ObjectTypes field.
func (o *BTMetadataCategorySummaryInfo) SetObjectTypes(v []int32) {
	o.ObjectTypes = v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *BTMetadataCategorySummaryInfo) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCategorySummaryInfo) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *BTMetadataCategorySummaryInfo) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *BTMetadataCategorySummaryInfo) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *BTMetadataCategorySummaryInfo) GetOwnerType() int32 {
	if o == nil || o.OwnerType == nil {
		var ret int32
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCategorySummaryInfo) GetOwnerTypeOk() (*int32, bool) {
	if o == nil || o.OwnerType == nil {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *BTMetadataCategorySummaryInfo) HasOwnerType() bool {
	if o != nil && o.OwnerType != nil {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given int32 and assigns it to the OwnerType field.
func (o *BTMetadataCategorySummaryInfo) SetOwnerType(v int32) {
	o.OwnerType = &v
}

// GetPublishState returns the PublishState field value if set, zero value otherwise.
func (o *BTMetadataCategorySummaryInfo) GetPublishState() int32 {
	if o == nil || o.PublishState == nil {
		var ret int32
		return ret
	}
	return *o.PublishState
}

// GetPublishStateOk returns a tuple with the PublishState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCategorySummaryInfo) GetPublishStateOk() (*int32, bool) {
	if o == nil || o.PublishState == nil {
		return nil, false
	}
	return o.PublishState, true
}

// HasPublishState returns a boolean if a field has been set.
func (o *BTMetadataCategorySummaryInfo) HasPublishState() bool {
	if o != nil && o.PublishState != nil {
		return true
	}

	return false
}

// SetPublishState gets a reference to the given int32 and assigns it to the PublishState field.
func (o *BTMetadataCategorySummaryInfo) SetPublishState(v int32) {
	o.PublishState = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTMetadataCategorySummaryInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCategorySummaryInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTMetadataCategorySummaryInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTMetadataCategorySummaryInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTMetadataCategorySummaryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultObjectType != nil {
		toSerialize["defaultObjectType"] = o.DefaultObjectType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ObjectTypes != nil {
		toSerialize["objectTypes"] = o.ObjectTypes
	}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.OwnerType != nil {
		toSerialize["ownerType"] = o.OwnerType
	}
	if o.PublishState != nil {
		toSerialize["publishState"] = o.PublishState
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTMetadataCategorySummaryInfo struct {
	value *BTMetadataCategorySummaryInfo
	isSet bool
}

func (v NullableBTMetadataCategorySummaryInfo) Get() *BTMetadataCategorySummaryInfo {
	return v.value
}

func (v *NullableBTMetadataCategorySummaryInfo) Set(val *BTMetadataCategorySummaryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMetadataCategorySummaryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMetadataCategorySummaryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMetadataCategorySummaryInfo(val *BTMetadataCategorySummaryInfo) *NullableBTMetadataCategorySummaryInfo {
	return &NullableBTMetadataCategorySummaryInfo{value: val, isSet: true}
}

func (v NullableBTMetadataCategorySummaryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMetadataCategorySummaryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
