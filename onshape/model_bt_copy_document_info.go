/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5809-984b9f882afe
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCopyDocumentInfo struct for BTCopyDocumentInfo
type BTCopyDocumentInfo struct {
	NewDocumentId   *string      `json:"newDocumentId,omitempty"`
	NewDocumentName *string      `json:"newDocumentName,omitempty"`
	NewOwner        *BTOwnerInfo `json:"newOwner,omitempty"`
	NewParentId     *string      `json:"newParentId,omitempty"`
	NewProjectId    *string      `json:"newProjectId,omitempty"`
	NewWorkspaceId  *string      `json:"newWorkspaceId,omitempty"`
}

// NewBTCopyDocumentInfo instantiates a new BTCopyDocumentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCopyDocumentInfo() *BTCopyDocumentInfo {
	this := BTCopyDocumentInfo{}
	return &this
}

// NewBTCopyDocumentInfoWithDefaults instantiates a new BTCopyDocumentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCopyDocumentInfoWithDefaults() *BTCopyDocumentInfo {
	this := BTCopyDocumentInfo{}
	return &this
}

// GetNewDocumentId returns the NewDocumentId field value if set, zero value otherwise.
func (o *BTCopyDocumentInfo) GetNewDocumentId() string {
	if o == nil || o.NewDocumentId == nil {
		var ret string
		return ret
	}
	return *o.NewDocumentId
}

// GetNewDocumentIdOk returns a tuple with the NewDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyDocumentInfo) GetNewDocumentIdOk() (*string, bool) {
	if o == nil || o.NewDocumentId == nil {
		return nil, false
	}
	return o.NewDocumentId, true
}

// HasNewDocumentId returns a boolean if a field has been set.
func (o *BTCopyDocumentInfo) HasNewDocumentId() bool {
	if o != nil && o.NewDocumentId != nil {
		return true
	}

	return false
}

// SetNewDocumentId gets a reference to the given string and assigns it to the NewDocumentId field.
func (o *BTCopyDocumentInfo) SetNewDocumentId(v string) {
	o.NewDocumentId = &v
}

// GetNewDocumentName returns the NewDocumentName field value if set, zero value otherwise.
func (o *BTCopyDocumentInfo) GetNewDocumentName() string {
	if o == nil || o.NewDocumentName == nil {
		var ret string
		return ret
	}
	return *o.NewDocumentName
}

// GetNewDocumentNameOk returns a tuple with the NewDocumentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyDocumentInfo) GetNewDocumentNameOk() (*string, bool) {
	if o == nil || o.NewDocumentName == nil {
		return nil, false
	}
	return o.NewDocumentName, true
}

// HasNewDocumentName returns a boolean if a field has been set.
func (o *BTCopyDocumentInfo) HasNewDocumentName() bool {
	if o != nil && o.NewDocumentName != nil {
		return true
	}

	return false
}

// SetNewDocumentName gets a reference to the given string and assigns it to the NewDocumentName field.
func (o *BTCopyDocumentInfo) SetNewDocumentName(v string) {
	o.NewDocumentName = &v
}

// GetNewOwner returns the NewOwner field value if set, zero value otherwise.
func (o *BTCopyDocumentInfo) GetNewOwner() BTOwnerInfo {
	if o == nil || o.NewOwner == nil {
		var ret BTOwnerInfo
		return ret
	}
	return *o.NewOwner
}

// GetNewOwnerOk returns a tuple with the NewOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyDocumentInfo) GetNewOwnerOk() (*BTOwnerInfo, bool) {
	if o == nil || o.NewOwner == nil {
		return nil, false
	}
	return o.NewOwner, true
}

// HasNewOwner returns a boolean if a field has been set.
func (o *BTCopyDocumentInfo) HasNewOwner() bool {
	if o != nil && o.NewOwner != nil {
		return true
	}

	return false
}

// SetNewOwner gets a reference to the given BTOwnerInfo and assigns it to the NewOwner field.
func (o *BTCopyDocumentInfo) SetNewOwner(v BTOwnerInfo) {
	o.NewOwner = &v
}

// GetNewParentId returns the NewParentId field value if set, zero value otherwise.
func (o *BTCopyDocumentInfo) GetNewParentId() string {
	if o == nil || o.NewParentId == nil {
		var ret string
		return ret
	}
	return *o.NewParentId
}

// GetNewParentIdOk returns a tuple with the NewParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyDocumentInfo) GetNewParentIdOk() (*string, bool) {
	if o == nil || o.NewParentId == nil {
		return nil, false
	}
	return o.NewParentId, true
}

// HasNewParentId returns a boolean if a field has been set.
func (o *BTCopyDocumentInfo) HasNewParentId() bool {
	if o != nil && o.NewParentId != nil {
		return true
	}

	return false
}

// SetNewParentId gets a reference to the given string and assigns it to the NewParentId field.
func (o *BTCopyDocumentInfo) SetNewParentId(v string) {
	o.NewParentId = &v
}

// GetNewProjectId returns the NewProjectId field value if set, zero value otherwise.
func (o *BTCopyDocumentInfo) GetNewProjectId() string {
	if o == nil || o.NewProjectId == nil {
		var ret string
		return ret
	}
	return *o.NewProjectId
}

// GetNewProjectIdOk returns a tuple with the NewProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyDocumentInfo) GetNewProjectIdOk() (*string, bool) {
	if o == nil || o.NewProjectId == nil {
		return nil, false
	}
	return o.NewProjectId, true
}

// HasNewProjectId returns a boolean if a field has been set.
func (o *BTCopyDocumentInfo) HasNewProjectId() bool {
	if o != nil && o.NewProjectId != nil {
		return true
	}

	return false
}

// SetNewProjectId gets a reference to the given string and assigns it to the NewProjectId field.
func (o *BTCopyDocumentInfo) SetNewProjectId(v string) {
	o.NewProjectId = &v
}

// GetNewWorkspaceId returns the NewWorkspaceId field value if set, zero value otherwise.
func (o *BTCopyDocumentInfo) GetNewWorkspaceId() string {
	if o == nil || o.NewWorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.NewWorkspaceId
}

// GetNewWorkspaceIdOk returns a tuple with the NewWorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyDocumentInfo) GetNewWorkspaceIdOk() (*string, bool) {
	if o == nil || o.NewWorkspaceId == nil {
		return nil, false
	}
	return o.NewWorkspaceId, true
}

// HasNewWorkspaceId returns a boolean if a field has been set.
func (o *BTCopyDocumentInfo) HasNewWorkspaceId() bool {
	if o != nil && o.NewWorkspaceId != nil {
		return true
	}

	return false
}

// SetNewWorkspaceId gets a reference to the given string and assigns it to the NewWorkspaceId field.
func (o *BTCopyDocumentInfo) SetNewWorkspaceId(v string) {
	o.NewWorkspaceId = &v
}

func (o BTCopyDocumentInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NewDocumentId != nil {
		toSerialize["newDocumentId"] = o.NewDocumentId
	}
	if o.NewDocumentName != nil {
		toSerialize["newDocumentName"] = o.NewDocumentName
	}
	if o.NewOwner != nil {
		toSerialize["newOwner"] = o.NewOwner
	}
	if o.NewParentId != nil {
		toSerialize["newParentId"] = o.NewParentId
	}
	if o.NewProjectId != nil {
		toSerialize["newProjectId"] = o.NewProjectId
	}
	if o.NewWorkspaceId != nil {
		toSerialize["newWorkspaceId"] = o.NewWorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableBTCopyDocumentInfo struct {
	value *BTCopyDocumentInfo
	isSet bool
}

func (v NullableBTCopyDocumentInfo) Get() *BTCopyDocumentInfo {
	return v.value
}

func (v *NullableBTCopyDocumentInfo) Set(val *BTCopyDocumentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCopyDocumentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCopyDocumentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCopyDocumentInfo(val *BTCopyDocumentInfo) *NullableBTCopyDocumentInfo {
	return &NullableBTCopyDocumentInfo{value: val, isSet: true}
}

func (v NullableBTCopyDocumentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCopyDocumentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
