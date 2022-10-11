/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6859-c8a2bd7d8338
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDocumentSearchHitInfo struct for BTDocumentSearchHitInfo
type BTDocumentSearchHitInfo struct {
	DocumentId             *string                           `json:"documentId,omitempty"`
	ElementName            *string                           `json:"elementName,omitempty"`
	HighlightedFields      *map[string][]string              `json:"highlightedFields,omitempty"`
	HitId                  *string                           `json:"hitId,omitempty"`
	Name                   *string                           `json:"name,omitempty"`
	ProjectId              *string                           `json:"projectId,omitempty"`
	SourceMap              map[string]map[string]interface{} `json:"sourceMap,omitempty"`
	Type                   *string                           `json:"type,omitempty"`
	VersionOrWorkspaceName *string                           `json:"versionOrWorkspaceName,omitempty"`
}

// NewBTDocumentSearchHitInfo instantiates a new BTDocumentSearchHitInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentSearchHitInfo() *BTDocumentSearchHitInfo {
	this := BTDocumentSearchHitInfo{}
	return &this
}

// NewBTDocumentSearchHitInfoWithDefaults instantiates a new BTDocumentSearchHitInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentSearchHitInfoWithDefaults() *BTDocumentSearchHitInfo {
	this := BTDocumentSearchHitInfo{}
	return &this
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTDocumentSearchHitInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchHitInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTDocumentSearchHitInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTDocumentSearchHitInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementName returns the ElementName field value if set, zero value otherwise.
func (o *BTDocumentSearchHitInfo) GetElementName() string {
	if o == nil || o.ElementName == nil {
		var ret string
		return ret
	}
	return *o.ElementName
}

// GetElementNameOk returns a tuple with the ElementName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchHitInfo) GetElementNameOk() (*string, bool) {
	if o == nil || o.ElementName == nil {
		return nil, false
	}
	return o.ElementName, true
}

// HasElementName returns a boolean if a field has been set.
func (o *BTDocumentSearchHitInfo) HasElementName() bool {
	if o != nil && o.ElementName != nil {
		return true
	}

	return false
}

// SetElementName gets a reference to the given string and assigns it to the ElementName field.
func (o *BTDocumentSearchHitInfo) SetElementName(v string) {
	o.ElementName = &v
}

// GetHighlightedFields returns the HighlightedFields field value if set, zero value otherwise.
func (o *BTDocumentSearchHitInfo) GetHighlightedFields() map[string][]string {
	if o == nil || o.HighlightedFields == nil {
		var ret map[string][]string
		return ret
	}
	return *o.HighlightedFields
}

// GetHighlightedFieldsOk returns a tuple with the HighlightedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchHitInfo) GetHighlightedFieldsOk() (*map[string][]string, bool) {
	if o == nil || o.HighlightedFields == nil {
		return nil, false
	}
	return o.HighlightedFields, true
}

// HasHighlightedFields returns a boolean if a field has been set.
func (o *BTDocumentSearchHitInfo) HasHighlightedFields() bool {
	if o != nil && o.HighlightedFields != nil {
		return true
	}

	return false
}

// SetHighlightedFields gets a reference to the given map[string][]string and assigns it to the HighlightedFields field.
func (o *BTDocumentSearchHitInfo) SetHighlightedFields(v map[string][]string) {
	o.HighlightedFields = &v
}

// GetHitId returns the HitId field value if set, zero value otherwise.
func (o *BTDocumentSearchHitInfo) GetHitId() string {
	if o == nil || o.HitId == nil {
		var ret string
		return ret
	}
	return *o.HitId
}

// GetHitIdOk returns a tuple with the HitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchHitInfo) GetHitIdOk() (*string, bool) {
	if o == nil || o.HitId == nil {
		return nil, false
	}
	return o.HitId, true
}

// HasHitId returns a boolean if a field has been set.
func (o *BTDocumentSearchHitInfo) HasHitId() bool {
	if o != nil && o.HitId != nil {
		return true
	}

	return false
}

// SetHitId gets a reference to the given string and assigns it to the HitId field.
func (o *BTDocumentSearchHitInfo) SetHitId(v string) {
	o.HitId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTDocumentSearchHitInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchHitInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTDocumentSearchHitInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTDocumentSearchHitInfo) SetName(v string) {
	o.Name = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BTDocumentSearchHitInfo) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchHitInfo) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BTDocumentSearchHitInfo) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *BTDocumentSearchHitInfo) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetSourceMap returns the SourceMap field value if set, zero value otherwise.
func (o *BTDocumentSearchHitInfo) GetSourceMap() map[string]map[string]interface{} {
	if o == nil || o.SourceMap == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.SourceMap
}

// GetSourceMapOk returns a tuple with the SourceMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchHitInfo) GetSourceMapOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.SourceMap == nil {
		return nil, false
	}
	return o.SourceMap, true
}

// HasSourceMap returns a boolean if a field has been set.
func (o *BTDocumentSearchHitInfo) HasSourceMap() bool {
	if o != nil && o.SourceMap != nil {
		return true
	}

	return false
}

// SetSourceMap gets a reference to the given map[string]map[string]interface{} and assigns it to the SourceMap field.
func (o *BTDocumentSearchHitInfo) SetSourceMap(v map[string]map[string]interface{}) {
	o.SourceMap = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTDocumentSearchHitInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchHitInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTDocumentSearchHitInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTDocumentSearchHitInfo) SetType(v string) {
	o.Type = &v
}

// GetVersionOrWorkspaceName returns the VersionOrWorkspaceName field value if set, zero value otherwise.
func (o *BTDocumentSearchHitInfo) GetVersionOrWorkspaceName() string {
	if o == nil || o.VersionOrWorkspaceName == nil {
		var ret string
		return ret
	}
	return *o.VersionOrWorkspaceName
}

// GetVersionOrWorkspaceNameOk returns a tuple with the VersionOrWorkspaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchHitInfo) GetVersionOrWorkspaceNameOk() (*string, bool) {
	if o == nil || o.VersionOrWorkspaceName == nil {
		return nil, false
	}
	return o.VersionOrWorkspaceName, true
}

// HasVersionOrWorkspaceName returns a boolean if a field has been set.
func (o *BTDocumentSearchHitInfo) HasVersionOrWorkspaceName() bool {
	if o != nil && o.VersionOrWorkspaceName != nil {
		return true
	}

	return false
}

// SetVersionOrWorkspaceName gets a reference to the given string and assigns it to the VersionOrWorkspaceName field.
func (o *BTDocumentSearchHitInfo) SetVersionOrWorkspaceName(v string) {
	o.VersionOrWorkspaceName = &v
}

func (o BTDocumentSearchHitInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementName != nil {
		toSerialize["elementName"] = o.ElementName
	}
	if o.HighlightedFields != nil {
		toSerialize["highlightedFields"] = o.HighlightedFields
	}
	if o.HitId != nil {
		toSerialize["hitId"] = o.HitId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.SourceMap != nil {
		toSerialize["sourceMap"] = o.SourceMap
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.VersionOrWorkspaceName != nil {
		toSerialize["versionOrWorkspaceName"] = o.VersionOrWorkspaceName
	}
	return json.Marshal(toSerialize)
}

type NullableBTDocumentSearchHitInfo struct {
	value *BTDocumentSearchHitInfo
	isSet bool
}

func (v NullableBTDocumentSearchHitInfo) Get() *BTDocumentSearchHitInfo {
	return v.value
}

func (v *NullableBTDocumentSearchHitInfo) Set(val *BTDocumentSearchHitInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentSearchHitInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentSearchHitInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentSearchHitInfo(val *BTDocumentSearchHitInfo) *NullableBTDocumentSearchHitInfo {
	return &NullableBTDocumentSearchHitInfo{value: val, isSet: true}
}

func (v NullableBTDocumentSearchHitInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentSearchHitInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
