/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.166.18273-3025d52f81b7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTBillOfMaterialsItemSourceInfo struct for BTBillOfMaterialsItemSourceInfo
type BTBillOfMaterialsItemSourceInfo struct {
	Configuration                          *string          `json:"configuration,omitempty"`
	DocumentId                             *string          `json:"documentId,omitempty"`
	ElementId                              *string          `json:"elementId,omitempty"`
	FullConfiguration                      *string          `json:"fullConfiguration,omitempty"`
	Href                                   *string          `json:"href,omitempty"`
	IsStandardContent                      *bool            `json:"isStandardContent,omitempty"`
	NonGeometricItemId                     *string          `json:"nonGeometricItemId,omitempty"`
	PartId                                 *string          `json:"partId,omitempty"`
	SourceElementMicroversionId            *string          `json:"sourceElementMicroversionId,omitempty"`
	ThumbnailInfo                          *BTThumbnailInfo `json:"thumbnailInfo,omitempty"`
	VersionMetadataWorkspaceMicroversionId *string          `json:"versionMetadataWorkspaceMicroversionId,omitempty"`
	ViewHref                               *string          `json:"viewHref,omitempty"`
	WvmId                                  *string          `json:"wvmId,omitempty"`
	WvmType                                *string          `json:"wvmType,omitempty"`
}

// NewBTBillOfMaterialsItemSourceInfo instantiates a new BTBillOfMaterialsItemSourceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBillOfMaterialsItemSourceInfo() *BTBillOfMaterialsItemSourceInfo {
	this := BTBillOfMaterialsItemSourceInfo{}
	return &this
}

// NewBTBillOfMaterialsItemSourceInfoWithDefaults instantiates a new BTBillOfMaterialsItemSourceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBillOfMaterialsItemSourceInfoWithDefaults() *BTBillOfMaterialsItemSourceInfo {
	this := BTBillOfMaterialsItemSourceInfo{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BTBillOfMaterialsItemSourceInfo) GetConfiguration() string {
	if o == nil || o.Configuration == nil {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsItemSourceInfo) GetConfigurationOk() (*string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BTBillOfMaterialsItemSourceInfo) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *BTBillOfMaterialsItemSourceInfo) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsItemSourceInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsItemSourceInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsItemSourceInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTBillOfMaterialsItemSourceInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsItemSourceInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsItemSourceInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsItemSourceInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTBillOfMaterialsItemSourceInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetFullConfiguration returns the FullConfiguration field value if set, zero value otherwise.
func (o *BTBillOfMaterialsItemSourceInfo) GetFullConfiguration() string {
	if o == nil || o.FullConfiguration == nil {
		var ret string
		return ret
	}
	return *o.FullConfiguration
}

// GetFullConfigurationOk returns a tuple with the FullConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsItemSourceInfo) GetFullConfigurationOk() (*string, bool) {
	if o == nil || o.FullConfiguration == nil {
		return nil, false
	}
	return o.FullConfiguration, true
}

// HasFullConfiguration returns a boolean if a field has been set.
func (o *BTBillOfMaterialsItemSourceInfo) HasFullConfiguration() bool {
	if o != nil && o.FullConfiguration != nil {
		return true
	}

	return false
}

// SetFullConfiguration gets a reference to the given string and assigns it to the FullConfiguration field.
func (o *BTBillOfMaterialsItemSourceInfo) SetFullConfiguration(v string) {
	o.FullConfiguration = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTBillOfMaterialsItemSourceInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsItemSourceInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTBillOfMaterialsItemSourceInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTBillOfMaterialsItemSourceInfo) SetHref(v string) {
	o.Href = &v
}

// GetIsStandardContent returns the IsStandardContent field value if set, zero value otherwise.
func (o *BTBillOfMaterialsItemSourceInfo) GetIsStandardContent() bool {
	if o == nil || o.IsStandardContent == nil {
		var ret bool
		return ret
	}
	return *o.IsStandardContent
}

// GetIsStandardContentOk returns a tuple with the IsStandardContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsItemSourceInfo) GetIsStandardContentOk() (*bool, bool) {
	if o == nil || o.IsStandardContent == nil {
		return nil, false
	}
	return o.IsStandardContent, true
}

// HasIsStandardContent returns a boolean if a field has been set.
func (o *BTBillOfMaterialsItemSourceInfo) HasIsStandardContent() bool {
	if o != nil && o.IsStandardContent != nil {
		return true
	}

	return false
}

// SetIsStandardContent gets a reference to the given bool and assigns it to the IsStandardContent field.
func (o *BTBillOfMaterialsItemSourceInfo) SetIsStandardContent(v bool) {
	o.IsStandardContent = &v
}

// GetNonGeometricItemId returns the NonGeometricItemId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsItemSourceInfo) GetNonGeometricItemId() string {
	if o == nil || o.NonGeometricItemId == nil {
		var ret string
		return ret
	}
	return *o.NonGeometricItemId
}

// GetNonGeometricItemIdOk returns a tuple with the NonGeometricItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsItemSourceInfo) GetNonGeometricItemIdOk() (*string, bool) {
	if o == nil || o.NonGeometricItemId == nil {
		return nil, false
	}
	return o.NonGeometricItemId, true
}

// HasNonGeometricItemId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsItemSourceInfo) HasNonGeometricItemId() bool {
	if o != nil && o.NonGeometricItemId != nil {
		return true
	}

	return false
}

// SetNonGeometricItemId gets a reference to the given string and assigns it to the NonGeometricItemId field.
func (o *BTBillOfMaterialsItemSourceInfo) SetNonGeometricItemId(v string) {
	o.NonGeometricItemId = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsItemSourceInfo) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsItemSourceInfo) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsItemSourceInfo) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTBillOfMaterialsItemSourceInfo) SetPartId(v string) {
	o.PartId = &v
}

// GetSourceElementMicroversionId returns the SourceElementMicroversionId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsItemSourceInfo) GetSourceElementMicroversionId() string {
	if o == nil || o.SourceElementMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.SourceElementMicroversionId
}

// GetSourceElementMicroversionIdOk returns a tuple with the SourceElementMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsItemSourceInfo) GetSourceElementMicroversionIdOk() (*string, bool) {
	if o == nil || o.SourceElementMicroversionId == nil {
		return nil, false
	}
	return o.SourceElementMicroversionId, true
}

// HasSourceElementMicroversionId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsItemSourceInfo) HasSourceElementMicroversionId() bool {
	if o != nil && o.SourceElementMicroversionId != nil {
		return true
	}

	return false
}

// SetSourceElementMicroversionId gets a reference to the given string and assigns it to the SourceElementMicroversionId field.
func (o *BTBillOfMaterialsItemSourceInfo) SetSourceElementMicroversionId(v string) {
	o.SourceElementMicroversionId = &v
}

// GetThumbnailInfo returns the ThumbnailInfo field value if set, zero value otherwise.
func (o *BTBillOfMaterialsItemSourceInfo) GetThumbnailInfo() BTThumbnailInfo {
	if o == nil || o.ThumbnailInfo == nil {
		var ret BTThumbnailInfo
		return ret
	}
	return *o.ThumbnailInfo
}

// GetThumbnailInfoOk returns a tuple with the ThumbnailInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsItemSourceInfo) GetThumbnailInfoOk() (*BTThumbnailInfo, bool) {
	if o == nil || o.ThumbnailInfo == nil {
		return nil, false
	}
	return o.ThumbnailInfo, true
}

// HasThumbnailInfo returns a boolean if a field has been set.
func (o *BTBillOfMaterialsItemSourceInfo) HasThumbnailInfo() bool {
	if o != nil && o.ThumbnailInfo != nil {
		return true
	}

	return false
}

// SetThumbnailInfo gets a reference to the given BTThumbnailInfo and assigns it to the ThumbnailInfo field.
func (o *BTBillOfMaterialsItemSourceInfo) SetThumbnailInfo(v BTThumbnailInfo) {
	o.ThumbnailInfo = &v
}

// GetVersionMetadataWorkspaceMicroversionId returns the VersionMetadataWorkspaceMicroversionId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsItemSourceInfo) GetVersionMetadataWorkspaceMicroversionId() string {
	if o == nil || o.VersionMetadataWorkspaceMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.VersionMetadataWorkspaceMicroversionId
}

// GetVersionMetadataWorkspaceMicroversionIdOk returns a tuple with the VersionMetadataWorkspaceMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsItemSourceInfo) GetVersionMetadataWorkspaceMicroversionIdOk() (*string, bool) {
	if o == nil || o.VersionMetadataWorkspaceMicroversionId == nil {
		return nil, false
	}
	return o.VersionMetadataWorkspaceMicroversionId, true
}

// HasVersionMetadataWorkspaceMicroversionId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsItemSourceInfo) HasVersionMetadataWorkspaceMicroversionId() bool {
	if o != nil && o.VersionMetadataWorkspaceMicroversionId != nil {
		return true
	}

	return false
}

// SetVersionMetadataWorkspaceMicroversionId gets a reference to the given string and assigns it to the VersionMetadataWorkspaceMicroversionId field.
func (o *BTBillOfMaterialsItemSourceInfo) SetVersionMetadataWorkspaceMicroversionId(v string) {
	o.VersionMetadataWorkspaceMicroversionId = &v
}

// GetViewHref returns the ViewHref field value if set, zero value otherwise.
func (o *BTBillOfMaterialsItemSourceInfo) GetViewHref() string {
	if o == nil || o.ViewHref == nil {
		var ret string
		return ret
	}
	return *o.ViewHref
}

// GetViewHrefOk returns a tuple with the ViewHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsItemSourceInfo) GetViewHrefOk() (*string, bool) {
	if o == nil || o.ViewHref == nil {
		return nil, false
	}
	return o.ViewHref, true
}

// HasViewHref returns a boolean if a field has been set.
func (o *BTBillOfMaterialsItemSourceInfo) HasViewHref() bool {
	if o != nil && o.ViewHref != nil {
		return true
	}

	return false
}

// SetViewHref gets a reference to the given string and assigns it to the ViewHref field.
func (o *BTBillOfMaterialsItemSourceInfo) SetViewHref(v string) {
	o.ViewHref = &v
}

// GetWvmId returns the WvmId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsItemSourceInfo) GetWvmId() string {
	if o == nil || o.WvmId == nil {
		var ret string
		return ret
	}
	return *o.WvmId
}

// GetWvmIdOk returns a tuple with the WvmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsItemSourceInfo) GetWvmIdOk() (*string, bool) {
	if o == nil || o.WvmId == nil {
		return nil, false
	}
	return o.WvmId, true
}

// HasWvmId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsItemSourceInfo) HasWvmId() bool {
	if o != nil && o.WvmId != nil {
		return true
	}

	return false
}

// SetWvmId gets a reference to the given string and assigns it to the WvmId field.
func (o *BTBillOfMaterialsItemSourceInfo) SetWvmId(v string) {
	o.WvmId = &v
}

// GetWvmType returns the WvmType field value if set, zero value otherwise.
func (o *BTBillOfMaterialsItemSourceInfo) GetWvmType() string {
	if o == nil || o.WvmType == nil {
		var ret string
		return ret
	}
	return *o.WvmType
}

// GetWvmTypeOk returns a tuple with the WvmType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsItemSourceInfo) GetWvmTypeOk() (*string, bool) {
	if o == nil || o.WvmType == nil {
		return nil, false
	}
	return o.WvmType, true
}

// HasWvmType returns a boolean if a field has been set.
func (o *BTBillOfMaterialsItemSourceInfo) HasWvmType() bool {
	if o != nil && o.WvmType != nil {
		return true
	}

	return false
}

// SetWvmType gets a reference to the given string and assigns it to the WvmType field.
func (o *BTBillOfMaterialsItemSourceInfo) SetWvmType(v string) {
	o.WvmType = &v
}

func (o BTBillOfMaterialsItemSourceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.FullConfiguration != nil {
		toSerialize["fullConfiguration"] = o.FullConfiguration
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.IsStandardContent != nil {
		toSerialize["isStandardContent"] = o.IsStandardContent
	}
	if o.NonGeometricItemId != nil {
		toSerialize["nonGeometricItemId"] = o.NonGeometricItemId
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.SourceElementMicroversionId != nil {
		toSerialize["sourceElementMicroversionId"] = o.SourceElementMicroversionId
	}
	if o.ThumbnailInfo != nil {
		toSerialize["thumbnailInfo"] = o.ThumbnailInfo
	}
	if o.VersionMetadataWorkspaceMicroversionId != nil {
		toSerialize["versionMetadataWorkspaceMicroversionId"] = o.VersionMetadataWorkspaceMicroversionId
	}
	if o.ViewHref != nil {
		toSerialize["viewHref"] = o.ViewHref
	}
	if o.WvmId != nil {
		toSerialize["wvmId"] = o.WvmId
	}
	if o.WvmType != nil {
		toSerialize["wvmType"] = o.WvmType
	}
	return json.Marshal(toSerialize)
}

type NullableBTBillOfMaterialsItemSourceInfo struct {
	value *BTBillOfMaterialsItemSourceInfo
	isSet bool
}

func (v NullableBTBillOfMaterialsItemSourceInfo) Get() *BTBillOfMaterialsItemSourceInfo {
	return v.value
}

func (v *NullableBTBillOfMaterialsItemSourceInfo) Set(val *BTBillOfMaterialsItemSourceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBillOfMaterialsItemSourceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBillOfMaterialsItemSourceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBillOfMaterialsItemSourceInfo(val *BTBillOfMaterialsItemSourceInfo) *NullableBTBillOfMaterialsItemSourceInfo {
	return &NullableBTBillOfMaterialsItemSourceInfo{value: val, isSet: true}
}

func (v NullableBTBillOfMaterialsItemSourceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBillOfMaterialsItemSourceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
