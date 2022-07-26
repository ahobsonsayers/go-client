/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5700-1c3471f20a39
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTInsertablesListResponse struct for BTInsertablesListResponse
type BTInsertablesListResponse struct {
	CanSaveVersion          *bool                     `json:"canSaveVersion,omitempty"`
	ChangesSinceVersionSave *int32                    `json:"changesSinceVersionSave,omitempty"`
	Configuration           *map[string]BTFSValue1888 `json:"configuration,omitempty"`
	ConfigurationKey        *string                   `json:"configurationKey,omitempty"`
	HasMultipleVersions     *bool                     `json:"hasMultipleVersions,omitempty"`
	Href                    *string                   `json:"href,omitempty"`
	Items                   []BTInsertableInfo        `json:"items,omitempty"`
	Next                    *string                   `json:"next,omitempty"`
	Previous                *string                   `json:"previous,omitempty"`
	UpdatedThumbnailUri     *string                   `json:"updatedThumbnailUri,omitempty"`
}

// NewBTInsertablesListResponse instantiates a new BTInsertablesListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTInsertablesListResponse() *BTInsertablesListResponse {
	this := BTInsertablesListResponse{}
	return &this
}

// NewBTInsertablesListResponseWithDefaults instantiates a new BTInsertablesListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTInsertablesListResponseWithDefaults() *BTInsertablesListResponse {
	this := BTInsertablesListResponse{}
	return &this
}

// GetCanSaveVersion returns the CanSaveVersion field value if set, zero value otherwise.
func (o *BTInsertablesListResponse) GetCanSaveVersion() bool {
	if o == nil || o.CanSaveVersion == nil {
		var ret bool
		return ret
	}
	return *o.CanSaveVersion
}

// GetCanSaveVersionOk returns a tuple with the CanSaveVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablesListResponse) GetCanSaveVersionOk() (*bool, bool) {
	if o == nil || o.CanSaveVersion == nil {
		return nil, false
	}
	return o.CanSaveVersion, true
}

// HasCanSaveVersion returns a boolean if a field has been set.
func (o *BTInsertablesListResponse) HasCanSaveVersion() bool {
	if o != nil && o.CanSaveVersion != nil {
		return true
	}

	return false
}

// SetCanSaveVersion gets a reference to the given bool and assigns it to the CanSaveVersion field.
func (o *BTInsertablesListResponse) SetCanSaveVersion(v bool) {
	o.CanSaveVersion = &v
}

// GetChangesSinceVersionSave returns the ChangesSinceVersionSave field value if set, zero value otherwise.
func (o *BTInsertablesListResponse) GetChangesSinceVersionSave() int32 {
	if o == nil || o.ChangesSinceVersionSave == nil {
		var ret int32
		return ret
	}
	return *o.ChangesSinceVersionSave
}

// GetChangesSinceVersionSaveOk returns a tuple with the ChangesSinceVersionSave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablesListResponse) GetChangesSinceVersionSaveOk() (*int32, bool) {
	if o == nil || o.ChangesSinceVersionSave == nil {
		return nil, false
	}
	return o.ChangesSinceVersionSave, true
}

// HasChangesSinceVersionSave returns a boolean if a field has been set.
func (o *BTInsertablesListResponse) HasChangesSinceVersionSave() bool {
	if o != nil && o.ChangesSinceVersionSave != nil {
		return true
	}

	return false
}

// SetChangesSinceVersionSave gets a reference to the given int32 and assigns it to the ChangesSinceVersionSave field.
func (o *BTInsertablesListResponse) SetChangesSinceVersionSave(v int32) {
	o.ChangesSinceVersionSave = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BTInsertablesListResponse) GetConfiguration() map[string]BTFSValue1888 {
	if o == nil || o.Configuration == nil {
		var ret map[string]BTFSValue1888
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablesListResponse) GetConfigurationOk() (*map[string]BTFSValue1888, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BTInsertablesListResponse) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]BTFSValue1888 and assigns it to the Configuration field.
func (o *BTInsertablesListResponse) SetConfiguration(v map[string]BTFSValue1888) {
	o.Configuration = &v
}

// GetConfigurationKey returns the ConfigurationKey field value if set, zero value otherwise.
func (o *BTInsertablesListResponse) GetConfigurationKey() string {
	if o == nil || o.ConfigurationKey == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationKey
}

// GetConfigurationKeyOk returns a tuple with the ConfigurationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablesListResponse) GetConfigurationKeyOk() (*string, bool) {
	if o == nil || o.ConfigurationKey == nil {
		return nil, false
	}
	return o.ConfigurationKey, true
}

// HasConfigurationKey returns a boolean if a field has been set.
func (o *BTInsertablesListResponse) HasConfigurationKey() bool {
	if o != nil && o.ConfigurationKey != nil {
		return true
	}

	return false
}

// SetConfigurationKey gets a reference to the given string and assigns it to the ConfigurationKey field.
func (o *BTInsertablesListResponse) SetConfigurationKey(v string) {
	o.ConfigurationKey = &v
}

// GetHasMultipleVersions returns the HasMultipleVersions field value if set, zero value otherwise.
func (o *BTInsertablesListResponse) GetHasMultipleVersions() bool {
	if o == nil || o.HasMultipleVersions == nil {
		var ret bool
		return ret
	}
	return *o.HasMultipleVersions
}

// GetHasMultipleVersionsOk returns a tuple with the HasMultipleVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablesListResponse) GetHasMultipleVersionsOk() (*bool, bool) {
	if o == nil || o.HasMultipleVersions == nil {
		return nil, false
	}
	return o.HasMultipleVersions, true
}

// HasHasMultipleVersions returns a boolean if a field has been set.
func (o *BTInsertablesListResponse) HasHasMultipleVersions() bool {
	if o != nil && o.HasMultipleVersions != nil {
		return true
	}

	return false
}

// SetHasMultipleVersions gets a reference to the given bool and assigns it to the HasMultipleVersions field.
func (o *BTInsertablesListResponse) SetHasMultipleVersions(v bool) {
	o.HasMultipleVersions = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTInsertablesListResponse) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablesListResponse) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTInsertablesListResponse) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTInsertablesListResponse) SetHref(v string) {
	o.Href = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BTInsertablesListResponse) GetItems() []BTInsertableInfo {
	if o == nil || o.Items == nil {
		var ret []BTInsertableInfo
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablesListResponse) GetItemsOk() ([]BTInsertableInfo, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BTInsertablesListResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BTInsertableInfo and assigns it to the Items field.
func (o *BTInsertablesListResponse) SetItems(v []BTInsertableInfo) {
	o.Items = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *BTInsertablesListResponse) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablesListResponse) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *BTInsertablesListResponse) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *BTInsertablesListResponse) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *BTInsertablesListResponse) GetPrevious() string {
	if o == nil || o.Previous == nil {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablesListResponse) GetPreviousOk() (*string, bool) {
	if o == nil || o.Previous == nil {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *BTInsertablesListResponse) HasPrevious() bool {
	if o != nil && o.Previous != nil {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *BTInsertablesListResponse) SetPrevious(v string) {
	o.Previous = &v
}

// GetUpdatedThumbnailUri returns the UpdatedThumbnailUri field value if set, zero value otherwise.
func (o *BTInsertablesListResponse) GetUpdatedThumbnailUri() string {
	if o == nil || o.UpdatedThumbnailUri == nil {
		var ret string
		return ret
	}
	return *o.UpdatedThumbnailUri
}

// GetUpdatedThumbnailUriOk returns a tuple with the UpdatedThumbnailUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInsertablesListResponse) GetUpdatedThumbnailUriOk() (*string, bool) {
	if o == nil || o.UpdatedThumbnailUri == nil {
		return nil, false
	}
	return o.UpdatedThumbnailUri, true
}

// HasUpdatedThumbnailUri returns a boolean if a field has been set.
func (o *BTInsertablesListResponse) HasUpdatedThumbnailUri() bool {
	if o != nil && o.UpdatedThumbnailUri != nil {
		return true
	}

	return false
}

// SetUpdatedThumbnailUri gets a reference to the given string and assigns it to the UpdatedThumbnailUri field.
func (o *BTInsertablesListResponse) SetUpdatedThumbnailUri(v string) {
	o.UpdatedThumbnailUri = &v
}

func (o BTInsertablesListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanSaveVersion != nil {
		toSerialize["canSaveVersion"] = o.CanSaveVersion
	}
	if o.ChangesSinceVersionSave != nil {
		toSerialize["changesSinceVersionSave"] = o.ChangesSinceVersionSave
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.ConfigurationKey != nil {
		toSerialize["configurationKey"] = o.ConfigurationKey
	}
	if o.HasMultipleVersions != nil {
		toSerialize["hasMultipleVersions"] = o.HasMultipleVersions
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	if o.UpdatedThumbnailUri != nil {
		toSerialize["updatedThumbnailUri"] = o.UpdatedThumbnailUri
	}
	return json.Marshal(toSerialize)
}

type NullableBTInsertablesListResponse struct {
	value *BTInsertablesListResponse
	isSet bool
}

func (v NullableBTInsertablesListResponse) Get() *BTInsertablesListResponse {
	return v.value
}

func (v *NullableBTInsertablesListResponse) Set(val *BTInsertablesListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBTInsertablesListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBTInsertablesListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTInsertablesListResponse(val *BTInsertablesListResponse) *NullableBTInsertablesListResponse {
	return &NullableBTInsertablesListResponse{value: val, isSet: true}
}

func (v NullableBTInsertablesListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTInsertablesListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
