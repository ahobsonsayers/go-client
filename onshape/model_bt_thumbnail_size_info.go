/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.164.16485-2fb60926fdbd
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTThumbnailSizeInfo struct for BTThumbnailSizeInfo
type BTThumbnailSizeInfo struct {
	Href            *string `json:"href,omitempty"`
	MediaType       *string `json:"mediaType,omitempty"`
	RenderMode      *string `json:"renderMode,omitempty"`
	SheetName       *string `json:"sheetName,omitempty"`
	Size            *string `json:"size,omitempty"`
	UniqueId        *string `json:"uniqueId,omitempty"`
	ViewOrientation *string `json:"viewOrientation,omitempty"`
}

// NewBTThumbnailSizeInfo instantiates a new BTThumbnailSizeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTThumbnailSizeInfo() *BTThumbnailSizeInfo {
	this := BTThumbnailSizeInfo{}
	return &this
}

// NewBTThumbnailSizeInfoWithDefaults instantiates a new BTThumbnailSizeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTThumbnailSizeInfoWithDefaults() *BTThumbnailSizeInfo {
	this := BTThumbnailSizeInfo{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTThumbnailSizeInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTThumbnailSizeInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTThumbnailSizeInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTThumbnailSizeInfo) SetHref(v string) {
	o.Href = &v
}

// GetMediaType returns the MediaType field value if set, zero value otherwise.
func (o *BTThumbnailSizeInfo) GetMediaType() string {
	if o == nil || o.MediaType == nil {
		var ret string
		return ret
	}
	return *o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTThumbnailSizeInfo) GetMediaTypeOk() (*string, bool) {
	if o == nil || o.MediaType == nil {
		return nil, false
	}
	return o.MediaType, true
}

// HasMediaType returns a boolean if a field has been set.
func (o *BTThumbnailSizeInfo) HasMediaType() bool {
	if o != nil && o.MediaType != nil {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given string and assigns it to the MediaType field.
func (o *BTThumbnailSizeInfo) SetMediaType(v string) {
	o.MediaType = &v
}

// GetRenderMode returns the RenderMode field value if set, zero value otherwise.
func (o *BTThumbnailSizeInfo) GetRenderMode() string {
	if o == nil || o.RenderMode == nil {
		var ret string
		return ret
	}
	return *o.RenderMode
}

// GetRenderModeOk returns a tuple with the RenderMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTThumbnailSizeInfo) GetRenderModeOk() (*string, bool) {
	if o == nil || o.RenderMode == nil {
		return nil, false
	}
	return o.RenderMode, true
}

// HasRenderMode returns a boolean if a field has been set.
func (o *BTThumbnailSizeInfo) HasRenderMode() bool {
	if o != nil && o.RenderMode != nil {
		return true
	}

	return false
}

// SetRenderMode gets a reference to the given string and assigns it to the RenderMode field.
func (o *BTThumbnailSizeInfo) SetRenderMode(v string) {
	o.RenderMode = &v
}

// GetSheetName returns the SheetName field value if set, zero value otherwise.
func (o *BTThumbnailSizeInfo) GetSheetName() string {
	if o == nil || o.SheetName == nil {
		var ret string
		return ret
	}
	return *o.SheetName
}

// GetSheetNameOk returns a tuple with the SheetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTThumbnailSizeInfo) GetSheetNameOk() (*string, bool) {
	if o == nil || o.SheetName == nil {
		return nil, false
	}
	return o.SheetName, true
}

// HasSheetName returns a boolean if a field has been set.
func (o *BTThumbnailSizeInfo) HasSheetName() bool {
	if o != nil && o.SheetName != nil {
		return true
	}

	return false
}

// SetSheetName gets a reference to the given string and assigns it to the SheetName field.
func (o *BTThumbnailSizeInfo) SetSheetName(v string) {
	o.SheetName = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *BTThumbnailSizeInfo) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTThumbnailSizeInfo) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *BTThumbnailSizeInfo) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *BTThumbnailSizeInfo) SetSize(v string) {
	o.Size = &v
}

// GetUniqueId returns the UniqueId field value if set, zero value otherwise.
func (o *BTThumbnailSizeInfo) GetUniqueId() string {
	if o == nil || o.UniqueId == nil {
		var ret string
		return ret
	}
	return *o.UniqueId
}

// GetUniqueIdOk returns a tuple with the UniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTThumbnailSizeInfo) GetUniqueIdOk() (*string, bool) {
	if o == nil || o.UniqueId == nil {
		return nil, false
	}
	return o.UniqueId, true
}

// HasUniqueId returns a boolean if a field has been set.
func (o *BTThumbnailSizeInfo) HasUniqueId() bool {
	if o != nil && o.UniqueId != nil {
		return true
	}

	return false
}

// SetUniqueId gets a reference to the given string and assigns it to the UniqueId field.
func (o *BTThumbnailSizeInfo) SetUniqueId(v string) {
	o.UniqueId = &v
}

// GetViewOrientation returns the ViewOrientation field value if set, zero value otherwise.
func (o *BTThumbnailSizeInfo) GetViewOrientation() string {
	if o == nil || o.ViewOrientation == nil {
		var ret string
		return ret
	}
	return *o.ViewOrientation
}

// GetViewOrientationOk returns a tuple with the ViewOrientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTThumbnailSizeInfo) GetViewOrientationOk() (*string, bool) {
	if o == nil || o.ViewOrientation == nil {
		return nil, false
	}
	return o.ViewOrientation, true
}

// HasViewOrientation returns a boolean if a field has been set.
func (o *BTThumbnailSizeInfo) HasViewOrientation() bool {
	if o != nil && o.ViewOrientation != nil {
		return true
	}

	return false
}

// SetViewOrientation gets a reference to the given string and assigns it to the ViewOrientation field.
func (o *BTThumbnailSizeInfo) SetViewOrientation(v string) {
	o.ViewOrientation = &v
}

func (o BTThumbnailSizeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.MediaType != nil {
		toSerialize["mediaType"] = o.MediaType
	}
	if o.RenderMode != nil {
		toSerialize["renderMode"] = o.RenderMode
	}
	if o.SheetName != nil {
		toSerialize["sheetName"] = o.SheetName
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.UniqueId != nil {
		toSerialize["uniqueId"] = o.UniqueId
	}
	if o.ViewOrientation != nil {
		toSerialize["viewOrientation"] = o.ViewOrientation
	}
	return json.Marshal(toSerialize)
}

type NullableBTThumbnailSizeInfo struct {
	value *BTThumbnailSizeInfo
	isSet bool
}

func (v NullableBTThumbnailSizeInfo) Get() *BTThumbnailSizeInfo {
	return v.value
}

func (v *NullableBTThumbnailSizeInfo) Set(val *BTThumbnailSizeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTThumbnailSizeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTThumbnailSizeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTThumbnailSizeInfo(val *BTThumbnailSizeInfo) *NullableBTThumbnailSizeInfo {
	return &NullableBTThumbnailSizeInfo{value: val, isSet: true}
}

func (v NullableBTThumbnailSizeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTThumbnailSizeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
