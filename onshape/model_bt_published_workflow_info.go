/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6820-1bef41f9cc03
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPublishedWorkflowInfo struct for BTPublishedWorkflowInfo
type BTPublishedWorkflowInfo struct {
	ActiveState *int32  `json:"activeState,omitempty"`
	CompanyId   *string `json:"companyId,omitempty"`
	Description *string `json:"description,omitempty"`
	DocumentId  *string `json:"documentId,omitempty"`
	ElementId   *string `json:"elementId,omitempty"`
	Id          *string `json:"id,omitempty"`
	ImageSrc    *string `json:"imageSrc,omitempty"`
	IsPickable  *bool   `json:"isPickable,omitempty"`
	Json        *string `json:"json,omitempty"`
	Name        *string `json:"name,omitempty"`
	ObjectType  *int32  `json:"objectType,omitempty"`
	OwnerType   *int32  `json:"ownerType,omitempty"`
	VersionId   *string `json:"versionId,omitempty"`
}

// NewBTPublishedWorkflowInfo instantiates a new BTPublishedWorkflowInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPublishedWorkflowInfo() *BTPublishedWorkflowInfo {
	this := BTPublishedWorkflowInfo{}
	return &this
}

// NewBTPublishedWorkflowInfoWithDefaults instantiates a new BTPublishedWorkflowInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPublishedWorkflowInfoWithDefaults() *BTPublishedWorkflowInfo {
	this := BTPublishedWorkflowInfo{}
	return &this
}

// GetActiveState returns the ActiveState field value if set, zero value otherwise.
func (o *BTPublishedWorkflowInfo) GetActiveState() int32 {
	if o == nil || o.ActiveState == nil {
		var ret int32
		return ret
	}
	return *o.ActiveState
}

// GetActiveStateOk returns a tuple with the ActiveState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPublishedWorkflowInfo) GetActiveStateOk() (*int32, bool) {
	if o == nil || o.ActiveState == nil {
		return nil, false
	}
	return o.ActiveState, true
}

// HasActiveState returns a boolean if a field has been set.
func (o *BTPublishedWorkflowInfo) HasActiveState() bool {
	if o != nil && o.ActiveState != nil {
		return true
	}

	return false
}

// SetActiveState gets a reference to the given int32 and assigns it to the ActiveState field.
func (o *BTPublishedWorkflowInfo) SetActiveState(v int32) {
	o.ActiveState = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTPublishedWorkflowInfo) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPublishedWorkflowInfo) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTPublishedWorkflowInfo) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTPublishedWorkflowInfo) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTPublishedWorkflowInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPublishedWorkflowInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTPublishedWorkflowInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTPublishedWorkflowInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTPublishedWorkflowInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPublishedWorkflowInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTPublishedWorkflowInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTPublishedWorkflowInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTPublishedWorkflowInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPublishedWorkflowInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTPublishedWorkflowInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTPublishedWorkflowInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTPublishedWorkflowInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPublishedWorkflowInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTPublishedWorkflowInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTPublishedWorkflowInfo) SetId(v string) {
	o.Id = &v
}

// GetImageSrc returns the ImageSrc field value if set, zero value otherwise.
func (o *BTPublishedWorkflowInfo) GetImageSrc() string {
	if o == nil || o.ImageSrc == nil {
		var ret string
		return ret
	}
	return *o.ImageSrc
}

// GetImageSrcOk returns a tuple with the ImageSrc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPublishedWorkflowInfo) GetImageSrcOk() (*string, bool) {
	if o == nil || o.ImageSrc == nil {
		return nil, false
	}
	return o.ImageSrc, true
}

// HasImageSrc returns a boolean if a field has been set.
func (o *BTPublishedWorkflowInfo) HasImageSrc() bool {
	if o != nil && o.ImageSrc != nil {
		return true
	}

	return false
}

// SetImageSrc gets a reference to the given string and assigns it to the ImageSrc field.
func (o *BTPublishedWorkflowInfo) SetImageSrc(v string) {
	o.ImageSrc = &v
}

// GetIsPickable returns the IsPickable field value if set, zero value otherwise.
func (o *BTPublishedWorkflowInfo) GetIsPickable() bool {
	if o == nil || o.IsPickable == nil {
		var ret bool
		return ret
	}
	return *o.IsPickable
}

// GetIsPickableOk returns a tuple with the IsPickable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPublishedWorkflowInfo) GetIsPickableOk() (*bool, bool) {
	if o == nil || o.IsPickable == nil {
		return nil, false
	}
	return o.IsPickable, true
}

// HasIsPickable returns a boolean if a field has been set.
func (o *BTPublishedWorkflowInfo) HasIsPickable() bool {
	if o != nil && o.IsPickable != nil {
		return true
	}

	return false
}

// SetIsPickable gets a reference to the given bool and assigns it to the IsPickable field.
func (o *BTPublishedWorkflowInfo) SetIsPickable(v bool) {
	o.IsPickable = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *BTPublishedWorkflowInfo) GetJson() string {
	if o == nil || o.Json == nil {
		var ret string
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPublishedWorkflowInfo) GetJsonOk() (*string, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *BTPublishedWorkflowInfo) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given string and assigns it to the Json field.
func (o *BTPublishedWorkflowInfo) SetJson(v string) {
	o.Json = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTPublishedWorkflowInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPublishedWorkflowInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTPublishedWorkflowInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTPublishedWorkflowInfo) SetName(v string) {
	o.Name = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *BTPublishedWorkflowInfo) GetObjectType() int32 {
	if o == nil || o.ObjectType == nil {
		var ret int32
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPublishedWorkflowInfo) GetObjectTypeOk() (*int32, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *BTPublishedWorkflowInfo) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given int32 and assigns it to the ObjectType field.
func (o *BTPublishedWorkflowInfo) SetObjectType(v int32) {
	o.ObjectType = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *BTPublishedWorkflowInfo) GetOwnerType() int32 {
	if o == nil || o.OwnerType == nil {
		var ret int32
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPublishedWorkflowInfo) GetOwnerTypeOk() (*int32, bool) {
	if o == nil || o.OwnerType == nil {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *BTPublishedWorkflowInfo) HasOwnerType() bool {
	if o != nil && o.OwnerType != nil {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given int32 and assigns it to the OwnerType field.
func (o *BTPublishedWorkflowInfo) SetOwnerType(v int32) {
	o.OwnerType = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTPublishedWorkflowInfo) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPublishedWorkflowInfo) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTPublishedWorkflowInfo) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTPublishedWorkflowInfo) SetVersionId(v string) {
	o.VersionId = &v
}

func (o BTPublishedWorkflowInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveState != nil {
		toSerialize["activeState"] = o.ActiveState
	}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ImageSrc != nil {
		toSerialize["imageSrc"] = o.ImageSrc
	}
	if o.IsPickable != nil {
		toSerialize["isPickable"] = o.IsPickable
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.OwnerType != nil {
		toSerialize["ownerType"] = o.OwnerType
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTPublishedWorkflowInfo struct {
	value *BTPublishedWorkflowInfo
	isSet bool
}

func (v NullableBTPublishedWorkflowInfo) Get() *BTPublishedWorkflowInfo {
	return v.value
}

func (v *NullableBTPublishedWorkflowInfo) Set(val *BTPublishedWorkflowInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPublishedWorkflowInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPublishedWorkflowInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPublishedWorkflowInfo(val *BTPublishedWorkflowInfo) *NullableBTPublishedWorkflowInfo {
	return &NullableBTPublishedWorkflowInfo{value: val, isSet: true}
}

func (v NullableBTPublishedWorkflowInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPublishedWorkflowInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
