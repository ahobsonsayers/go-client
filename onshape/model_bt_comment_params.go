/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.156.7203-6065bde47225
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCommentParams struct for BTCommentParams
type BTCommentParams struct {
	AssemblyFeature   *string           `json:"assemblyFeature,omitempty"`
	Assignee          *string           `json:"assignee,omitempty"`
	DocumentId        *string           `json:"documentId,omitempty"`
	ElementFeature    *string           `json:"elementFeature,omitempty"`
	ElementId         *string           `json:"elementId,omitempty"`
	ElementOccurrence *string           `json:"elementOccurrence,omitempty"`
	ElementQuery      *string           `json:"elementQuery,omitempty"`
	Id                *string           `json:"id,omitempty"`
	Message           *string           `json:"message,omitempty"`
	ObjectId          *string           `json:"objectId,omitempty"`
	ObjectType        *int32            `json:"objectType,omitempty"`
	ParentId          *string           `json:"parentId,omitempty"`
	VersionId         *string           `json:"versionId,omitempty"`
	ViewData          *BTViewDataParams `json:"viewData,omitempty"`
	WorkspaceId       *string           `json:"workspaceId,omitempty"`
}

// NewBTCommentParams instantiates a new BTCommentParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCommentParams() *BTCommentParams {
	this := BTCommentParams{}
	return &this
}

// NewBTCommentParamsWithDefaults instantiates a new BTCommentParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCommentParamsWithDefaults() *BTCommentParams {
	this := BTCommentParams{}
	return &this
}

// GetAssemblyFeature returns the AssemblyFeature field value if set, zero value otherwise.
func (o *BTCommentParams) GetAssemblyFeature() string {
	if o == nil || o.AssemblyFeature == nil {
		var ret string
		return ret
	}
	return *o.AssemblyFeature
}

// GetAssemblyFeatureOk returns a tuple with the AssemblyFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentParams) GetAssemblyFeatureOk() (*string, bool) {
	if o == nil || o.AssemblyFeature == nil {
		return nil, false
	}
	return o.AssemblyFeature, true
}

// HasAssemblyFeature returns a boolean if a field has been set.
func (o *BTCommentParams) HasAssemblyFeature() bool {
	if o != nil && o.AssemblyFeature != nil {
		return true
	}

	return false
}

// SetAssemblyFeature gets a reference to the given string and assigns it to the AssemblyFeature field.
func (o *BTCommentParams) SetAssemblyFeature(v string) {
	o.AssemblyFeature = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *BTCommentParams) GetAssignee() string {
	if o == nil || o.Assignee == nil {
		var ret string
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentParams) GetAssigneeOk() (*string, bool) {
	if o == nil || o.Assignee == nil {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *BTCommentParams) HasAssignee() bool {
	if o != nil && o.Assignee != nil {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given string and assigns it to the Assignee field.
func (o *BTCommentParams) SetAssignee(v string) {
	o.Assignee = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTCommentParams) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentParams) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTCommentParams) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTCommentParams) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementFeature returns the ElementFeature field value if set, zero value otherwise.
func (o *BTCommentParams) GetElementFeature() string {
	if o == nil || o.ElementFeature == nil {
		var ret string
		return ret
	}
	return *o.ElementFeature
}

// GetElementFeatureOk returns a tuple with the ElementFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentParams) GetElementFeatureOk() (*string, bool) {
	if o == nil || o.ElementFeature == nil {
		return nil, false
	}
	return o.ElementFeature, true
}

// HasElementFeature returns a boolean if a field has been set.
func (o *BTCommentParams) HasElementFeature() bool {
	if o != nil && o.ElementFeature != nil {
		return true
	}

	return false
}

// SetElementFeature gets a reference to the given string and assigns it to the ElementFeature field.
func (o *BTCommentParams) SetElementFeature(v string) {
	o.ElementFeature = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTCommentParams) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentParams) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTCommentParams) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTCommentParams) SetElementId(v string) {
	o.ElementId = &v
}

// GetElementOccurrence returns the ElementOccurrence field value if set, zero value otherwise.
func (o *BTCommentParams) GetElementOccurrence() string {
	if o == nil || o.ElementOccurrence == nil {
		var ret string
		return ret
	}
	return *o.ElementOccurrence
}

// GetElementOccurrenceOk returns a tuple with the ElementOccurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentParams) GetElementOccurrenceOk() (*string, bool) {
	if o == nil || o.ElementOccurrence == nil {
		return nil, false
	}
	return o.ElementOccurrence, true
}

// HasElementOccurrence returns a boolean if a field has been set.
func (o *BTCommentParams) HasElementOccurrence() bool {
	if o != nil && o.ElementOccurrence != nil {
		return true
	}

	return false
}

// SetElementOccurrence gets a reference to the given string and assigns it to the ElementOccurrence field.
func (o *BTCommentParams) SetElementOccurrence(v string) {
	o.ElementOccurrence = &v
}

// GetElementQuery returns the ElementQuery field value if set, zero value otherwise.
func (o *BTCommentParams) GetElementQuery() string {
	if o == nil || o.ElementQuery == nil {
		var ret string
		return ret
	}
	return *o.ElementQuery
}

// GetElementQueryOk returns a tuple with the ElementQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentParams) GetElementQueryOk() (*string, bool) {
	if o == nil || o.ElementQuery == nil {
		return nil, false
	}
	return o.ElementQuery, true
}

// HasElementQuery returns a boolean if a field has been set.
func (o *BTCommentParams) HasElementQuery() bool {
	if o != nil && o.ElementQuery != nil {
		return true
	}

	return false
}

// SetElementQuery gets a reference to the given string and assigns it to the ElementQuery field.
func (o *BTCommentParams) SetElementQuery(v string) {
	o.ElementQuery = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTCommentParams) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentParams) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTCommentParams) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTCommentParams) SetId(v string) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BTCommentParams) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentParams) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BTCommentParams) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BTCommentParams) SetMessage(v string) {
	o.Message = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *BTCommentParams) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentParams) GetObjectIdOk() (*string, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *BTCommentParams) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *BTCommentParams) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *BTCommentParams) GetObjectType() int32 {
	if o == nil || o.ObjectType == nil {
		var ret int32
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentParams) GetObjectTypeOk() (*int32, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *BTCommentParams) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given int32 and assigns it to the ObjectType field.
func (o *BTCommentParams) SetObjectType(v int32) {
	o.ObjectType = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *BTCommentParams) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentParams) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *BTCommentParams) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *BTCommentParams) SetParentId(v string) {
	o.ParentId = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTCommentParams) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentParams) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTCommentParams) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTCommentParams) SetVersionId(v string) {
	o.VersionId = &v
}

// GetViewData returns the ViewData field value if set, zero value otherwise.
func (o *BTCommentParams) GetViewData() BTViewDataParams {
	if o == nil || o.ViewData == nil {
		var ret BTViewDataParams
		return ret
	}
	return *o.ViewData
}

// GetViewDataOk returns a tuple with the ViewData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentParams) GetViewDataOk() (*BTViewDataParams, bool) {
	if o == nil || o.ViewData == nil {
		return nil, false
	}
	return o.ViewData, true
}

// HasViewData returns a boolean if a field has been set.
func (o *BTCommentParams) HasViewData() bool {
	if o != nil && o.ViewData != nil {
		return true
	}

	return false
}

// SetViewData gets a reference to the given BTViewDataParams and assigns it to the ViewData field.
func (o *BTCommentParams) SetViewData(v BTViewDataParams) {
	o.ViewData = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *BTCommentParams) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommentParams) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *BTCommentParams) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *BTCommentParams) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o BTCommentParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssemblyFeature != nil {
		toSerialize["assemblyFeature"] = o.AssemblyFeature
	}
	if o.Assignee != nil {
		toSerialize["assignee"] = o.Assignee
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementFeature != nil {
		toSerialize["elementFeature"] = o.ElementFeature
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ElementOccurrence != nil {
		toSerialize["elementOccurrence"] = o.ElementOccurrence
	}
	if o.ElementQuery != nil {
		toSerialize["elementQuery"] = o.ElementQuery
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ObjectId != nil {
		toSerialize["objectId"] = o.ObjectId
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	if o.ViewData != nil {
		toSerialize["viewData"] = o.ViewData
	}
	if o.WorkspaceId != nil {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableBTCommentParams struct {
	value *BTCommentParams
	isSet bool
}

func (v NullableBTCommentParams) Get() *BTCommentParams {
	return v.value
}

func (v *NullableBTCommentParams) Set(val *BTCommentParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCommentParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCommentParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCommentParams(val *BTCommentParams) *NullableBTCommentParams {
	return &NullableBTCommentParams{value: val, isSet: true}
}

func (v NullableBTCommentParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCommentParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
