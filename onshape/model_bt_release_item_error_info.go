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

// BTReleaseItemErrorInfo struct for BTReleaseItemErrorInfo
type BTReleaseItemErrorInfo struct {
	ChangeTaskId        *string `json:"changeTaskId,omitempty"`
	DocumentId          *string `json:"documentId,omitempty"`
	Message             *string `json:"message,omitempty"`
	Ordinal             *int32  `json:"ordinal,omitempty"`
	PendingTaskId       *string `json:"pendingTaskId,omitempty"`
	PendingTaskObjectId *string `json:"pendingTaskObjectId,omitempty"`
	PendingTaskType     *string `json:"pendingTaskType,omitempty"`
	ReleaseId           *string `json:"releaseId,omitempty"`
	Severity            *int32  `json:"severity,omitempty"`
	VersionId           *string `json:"versionId,omitempty"`
	WorkspaceId         *string `json:"workspaceId,omitempty"`
}

// NewBTReleaseItemErrorInfo instantiates a new BTReleaseItemErrorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTReleaseItemErrorInfo() *BTReleaseItemErrorInfo {
	this := BTReleaseItemErrorInfo{}
	return &this
}

// NewBTReleaseItemErrorInfoWithDefaults instantiates a new BTReleaseItemErrorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTReleaseItemErrorInfoWithDefaults() *BTReleaseItemErrorInfo {
	this := BTReleaseItemErrorInfo{}
	return &this
}

// GetChangeTaskId returns the ChangeTaskId field value if set, zero value otherwise.
func (o *BTReleaseItemErrorInfo) GetChangeTaskId() string {
	if o == nil || o.ChangeTaskId == nil {
		var ret string
		return ret
	}
	return *o.ChangeTaskId
}

// GetChangeTaskIdOk returns a tuple with the ChangeTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemErrorInfo) GetChangeTaskIdOk() (*string, bool) {
	if o == nil || o.ChangeTaskId == nil {
		return nil, false
	}
	return o.ChangeTaskId, true
}

// HasChangeTaskId returns a boolean if a field has been set.
func (o *BTReleaseItemErrorInfo) HasChangeTaskId() bool {
	if o != nil && o.ChangeTaskId != nil {
		return true
	}

	return false
}

// SetChangeTaskId gets a reference to the given string and assigns it to the ChangeTaskId field.
func (o *BTReleaseItemErrorInfo) SetChangeTaskId(v string) {
	o.ChangeTaskId = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTReleaseItemErrorInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemErrorInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTReleaseItemErrorInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTReleaseItemErrorInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BTReleaseItemErrorInfo) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemErrorInfo) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BTReleaseItemErrorInfo) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BTReleaseItemErrorInfo) SetMessage(v string) {
	o.Message = &v
}

// GetOrdinal returns the Ordinal field value if set, zero value otherwise.
func (o *BTReleaseItemErrorInfo) GetOrdinal() int32 {
	if o == nil || o.Ordinal == nil {
		var ret int32
		return ret
	}
	return *o.Ordinal
}

// GetOrdinalOk returns a tuple with the Ordinal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemErrorInfo) GetOrdinalOk() (*int32, bool) {
	if o == nil || o.Ordinal == nil {
		return nil, false
	}
	return o.Ordinal, true
}

// HasOrdinal returns a boolean if a field has been set.
func (o *BTReleaseItemErrorInfo) HasOrdinal() bool {
	if o != nil && o.Ordinal != nil {
		return true
	}

	return false
}

// SetOrdinal gets a reference to the given int32 and assigns it to the Ordinal field.
func (o *BTReleaseItemErrorInfo) SetOrdinal(v int32) {
	o.Ordinal = &v
}

// GetPendingTaskId returns the PendingTaskId field value if set, zero value otherwise.
func (o *BTReleaseItemErrorInfo) GetPendingTaskId() string {
	if o == nil || o.PendingTaskId == nil {
		var ret string
		return ret
	}
	return *o.PendingTaskId
}

// GetPendingTaskIdOk returns a tuple with the PendingTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemErrorInfo) GetPendingTaskIdOk() (*string, bool) {
	if o == nil || o.PendingTaskId == nil {
		return nil, false
	}
	return o.PendingTaskId, true
}

// HasPendingTaskId returns a boolean if a field has been set.
func (o *BTReleaseItemErrorInfo) HasPendingTaskId() bool {
	if o != nil && o.PendingTaskId != nil {
		return true
	}

	return false
}

// SetPendingTaskId gets a reference to the given string and assigns it to the PendingTaskId field.
func (o *BTReleaseItemErrorInfo) SetPendingTaskId(v string) {
	o.PendingTaskId = &v
}

// GetPendingTaskObjectId returns the PendingTaskObjectId field value if set, zero value otherwise.
func (o *BTReleaseItemErrorInfo) GetPendingTaskObjectId() string {
	if o == nil || o.PendingTaskObjectId == nil {
		var ret string
		return ret
	}
	return *o.PendingTaskObjectId
}

// GetPendingTaskObjectIdOk returns a tuple with the PendingTaskObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemErrorInfo) GetPendingTaskObjectIdOk() (*string, bool) {
	if o == nil || o.PendingTaskObjectId == nil {
		return nil, false
	}
	return o.PendingTaskObjectId, true
}

// HasPendingTaskObjectId returns a boolean if a field has been set.
func (o *BTReleaseItemErrorInfo) HasPendingTaskObjectId() bool {
	if o != nil && o.PendingTaskObjectId != nil {
		return true
	}

	return false
}

// SetPendingTaskObjectId gets a reference to the given string and assigns it to the PendingTaskObjectId field.
func (o *BTReleaseItemErrorInfo) SetPendingTaskObjectId(v string) {
	o.PendingTaskObjectId = &v
}

// GetPendingTaskType returns the PendingTaskType field value if set, zero value otherwise.
func (o *BTReleaseItemErrorInfo) GetPendingTaskType() string {
	if o == nil || o.PendingTaskType == nil {
		var ret string
		return ret
	}
	return *o.PendingTaskType
}

// GetPendingTaskTypeOk returns a tuple with the PendingTaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemErrorInfo) GetPendingTaskTypeOk() (*string, bool) {
	if o == nil || o.PendingTaskType == nil {
		return nil, false
	}
	return o.PendingTaskType, true
}

// HasPendingTaskType returns a boolean if a field has been set.
func (o *BTReleaseItemErrorInfo) HasPendingTaskType() bool {
	if o != nil && o.PendingTaskType != nil {
		return true
	}

	return false
}

// SetPendingTaskType gets a reference to the given string and assigns it to the PendingTaskType field.
func (o *BTReleaseItemErrorInfo) SetPendingTaskType(v string) {
	o.PendingTaskType = &v
}

// GetReleaseId returns the ReleaseId field value if set, zero value otherwise.
func (o *BTReleaseItemErrorInfo) GetReleaseId() string {
	if o == nil || o.ReleaseId == nil {
		var ret string
		return ret
	}
	return *o.ReleaseId
}

// GetReleaseIdOk returns a tuple with the ReleaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemErrorInfo) GetReleaseIdOk() (*string, bool) {
	if o == nil || o.ReleaseId == nil {
		return nil, false
	}
	return o.ReleaseId, true
}

// HasReleaseId returns a boolean if a field has been set.
func (o *BTReleaseItemErrorInfo) HasReleaseId() bool {
	if o != nil && o.ReleaseId != nil {
		return true
	}

	return false
}

// SetReleaseId gets a reference to the given string and assigns it to the ReleaseId field.
func (o *BTReleaseItemErrorInfo) SetReleaseId(v string) {
	o.ReleaseId = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *BTReleaseItemErrorInfo) GetSeverity() int32 {
	if o == nil || o.Severity == nil {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemErrorInfo) GetSeverityOk() (*int32, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *BTReleaseItemErrorInfo) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *BTReleaseItemErrorInfo) SetSeverity(v int32) {
	o.Severity = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTReleaseItemErrorInfo) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemErrorInfo) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTReleaseItemErrorInfo) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTReleaseItemErrorInfo) SetVersionId(v string) {
	o.VersionId = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *BTReleaseItemErrorInfo) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseItemErrorInfo) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *BTReleaseItemErrorInfo) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *BTReleaseItemErrorInfo) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o BTReleaseItemErrorInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangeTaskId != nil {
		toSerialize["changeTaskId"] = o.ChangeTaskId
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Ordinal != nil {
		toSerialize["ordinal"] = o.Ordinal
	}
	if o.PendingTaskId != nil {
		toSerialize["pendingTaskId"] = o.PendingTaskId
	}
	if o.PendingTaskObjectId != nil {
		toSerialize["pendingTaskObjectId"] = o.PendingTaskObjectId
	}
	if o.PendingTaskType != nil {
		toSerialize["pendingTaskType"] = o.PendingTaskType
	}
	if o.ReleaseId != nil {
		toSerialize["releaseId"] = o.ReleaseId
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	if o.WorkspaceId != nil {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableBTReleaseItemErrorInfo struct {
	value *BTReleaseItemErrorInfo
	isSet bool
}

func (v NullableBTReleaseItemErrorInfo) Get() *BTReleaseItemErrorInfo {
	return v.value
}

func (v *NullableBTReleaseItemErrorInfo) Set(val *BTReleaseItemErrorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTReleaseItemErrorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTReleaseItemErrorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTReleaseItemErrorInfo(val *BTReleaseItemErrorInfo) *NullableBTReleaseItemErrorInfo {
	return &NullableBTReleaseItemErrorInfo{value: val, isSet: true}
}

func (v NullableBTReleaseItemErrorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTReleaseItemErrorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
