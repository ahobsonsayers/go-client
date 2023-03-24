/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13452-78145f970001
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAssemblyPartInfo struct for BTAssemblyPartInfo
type BTAssemblyPartInfo struct {
	BodyType             *string                       `json:"bodyType,omitempty"`
	Configuration        *string                       `json:"configuration,omitempty"`
	DocumentId           *string                       `json:"documentId,omitempty"`
	DocumentMicroversion *string                       `json:"documentMicroversion,omitempty"`
	DocumentVersion      *string                       `json:"documentVersion,omitempty"`
	ElementId            *string                       `json:"elementId,omitempty"`
	FullConfiguration    *string                       `json:"fullConfiguration,omitempty"`
	IsStandardContent    *bool                         `json:"isStandardContent,omitempty"`
	MateConnectors       []BTAssemblyMateConnectorInfo `json:"mateConnectors,omitempty"`
	PartId               *string                       `json:"partId,omitempty"`
	PartNumber           *string                       `json:"partNumber,omitempty"`
	Revision             *string                       `json:"revision,omitempty"`
}

// NewBTAssemblyPartInfo instantiates a new BTAssemblyPartInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyPartInfo() *BTAssemblyPartInfo {
	this := BTAssemblyPartInfo{}
	return &this
}

// NewBTAssemblyPartInfoWithDefaults instantiates a new BTAssemblyPartInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyPartInfoWithDefaults() *BTAssemblyPartInfo {
	this := BTAssemblyPartInfo{}
	return &this
}

// GetBodyType returns the BodyType field value if set, zero value otherwise.
func (o *BTAssemblyPartInfo) GetBodyType() string {
	if o == nil || o.BodyType == nil {
		var ret string
		return ret
	}
	return *o.BodyType
}

// GetBodyTypeOk returns a tuple with the BodyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPartInfo) GetBodyTypeOk() (*string, bool) {
	if o == nil || o.BodyType == nil {
		return nil, false
	}
	return o.BodyType, true
}

// HasBodyType returns a boolean if a field has been set.
func (o *BTAssemblyPartInfo) HasBodyType() bool {
	if o != nil && o.BodyType != nil {
		return true
	}

	return false
}

// SetBodyType gets a reference to the given string and assigns it to the BodyType field.
func (o *BTAssemblyPartInfo) SetBodyType(v string) {
	o.BodyType = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BTAssemblyPartInfo) GetConfiguration() string {
	if o == nil || o.Configuration == nil {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPartInfo) GetConfigurationOk() (*string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BTAssemblyPartInfo) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *BTAssemblyPartInfo) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTAssemblyPartInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPartInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTAssemblyPartInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTAssemblyPartInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetDocumentMicroversion returns the DocumentMicroversion field value if set, zero value otherwise.
func (o *BTAssemblyPartInfo) GetDocumentMicroversion() string {
	if o == nil || o.DocumentMicroversion == nil {
		var ret string
		return ret
	}
	return *o.DocumentMicroversion
}

// GetDocumentMicroversionOk returns a tuple with the DocumentMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPartInfo) GetDocumentMicroversionOk() (*string, bool) {
	if o == nil || o.DocumentMicroversion == nil {
		return nil, false
	}
	return o.DocumentMicroversion, true
}

// HasDocumentMicroversion returns a boolean if a field has been set.
func (o *BTAssemblyPartInfo) HasDocumentMicroversion() bool {
	if o != nil && o.DocumentMicroversion != nil {
		return true
	}

	return false
}

// SetDocumentMicroversion gets a reference to the given string and assigns it to the DocumentMicroversion field.
func (o *BTAssemblyPartInfo) SetDocumentMicroversion(v string) {
	o.DocumentMicroversion = &v
}

// GetDocumentVersion returns the DocumentVersion field value if set, zero value otherwise.
func (o *BTAssemblyPartInfo) GetDocumentVersion() string {
	if o == nil || o.DocumentVersion == nil {
		var ret string
		return ret
	}
	return *o.DocumentVersion
}

// GetDocumentVersionOk returns a tuple with the DocumentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPartInfo) GetDocumentVersionOk() (*string, bool) {
	if o == nil || o.DocumentVersion == nil {
		return nil, false
	}
	return o.DocumentVersion, true
}

// HasDocumentVersion returns a boolean if a field has been set.
func (o *BTAssemblyPartInfo) HasDocumentVersion() bool {
	if o != nil && o.DocumentVersion != nil {
		return true
	}

	return false
}

// SetDocumentVersion gets a reference to the given string and assigns it to the DocumentVersion field.
func (o *BTAssemblyPartInfo) SetDocumentVersion(v string) {
	o.DocumentVersion = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTAssemblyPartInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPartInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTAssemblyPartInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTAssemblyPartInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetFullConfiguration returns the FullConfiguration field value if set, zero value otherwise.
func (o *BTAssemblyPartInfo) GetFullConfiguration() string {
	if o == nil || o.FullConfiguration == nil {
		var ret string
		return ret
	}
	return *o.FullConfiguration
}

// GetFullConfigurationOk returns a tuple with the FullConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPartInfo) GetFullConfigurationOk() (*string, bool) {
	if o == nil || o.FullConfiguration == nil {
		return nil, false
	}
	return o.FullConfiguration, true
}

// HasFullConfiguration returns a boolean if a field has been set.
func (o *BTAssemblyPartInfo) HasFullConfiguration() bool {
	if o != nil && o.FullConfiguration != nil {
		return true
	}

	return false
}

// SetFullConfiguration gets a reference to the given string and assigns it to the FullConfiguration field.
func (o *BTAssemblyPartInfo) SetFullConfiguration(v string) {
	o.FullConfiguration = &v
}

// GetIsStandardContent returns the IsStandardContent field value if set, zero value otherwise.
func (o *BTAssemblyPartInfo) GetIsStandardContent() bool {
	if o == nil || o.IsStandardContent == nil {
		var ret bool
		return ret
	}
	return *o.IsStandardContent
}

// GetIsStandardContentOk returns a tuple with the IsStandardContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPartInfo) GetIsStandardContentOk() (*bool, bool) {
	if o == nil || o.IsStandardContent == nil {
		return nil, false
	}
	return o.IsStandardContent, true
}

// HasIsStandardContent returns a boolean if a field has been set.
func (o *BTAssemblyPartInfo) HasIsStandardContent() bool {
	if o != nil && o.IsStandardContent != nil {
		return true
	}

	return false
}

// SetIsStandardContent gets a reference to the given bool and assigns it to the IsStandardContent field.
func (o *BTAssemblyPartInfo) SetIsStandardContent(v bool) {
	o.IsStandardContent = &v
}

// GetMateConnectors returns the MateConnectors field value if set, zero value otherwise.
func (o *BTAssemblyPartInfo) GetMateConnectors() []BTAssemblyMateConnectorInfo {
	if o == nil || o.MateConnectors == nil {
		var ret []BTAssemblyMateConnectorInfo
		return ret
	}
	return o.MateConnectors
}

// GetMateConnectorsOk returns a tuple with the MateConnectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPartInfo) GetMateConnectorsOk() ([]BTAssemblyMateConnectorInfo, bool) {
	if o == nil || o.MateConnectors == nil {
		return nil, false
	}
	return o.MateConnectors, true
}

// HasMateConnectors returns a boolean if a field has been set.
func (o *BTAssemblyPartInfo) HasMateConnectors() bool {
	if o != nil && o.MateConnectors != nil {
		return true
	}

	return false
}

// SetMateConnectors gets a reference to the given []BTAssemblyMateConnectorInfo and assigns it to the MateConnectors field.
func (o *BTAssemblyPartInfo) SetMateConnectors(v []BTAssemblyMateConnectorInfo) {
	o.MateConnectors = v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTAssemblyPartInfo) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPartInfo) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTAssemblyPartInfo) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTAssemblyPartInfo) SetPartId(v string) {
	o.PartId = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTAssemblyPartInfo) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPartInfo) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTAssemblyPartInfo) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTAssemblyPartInfo) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *BTAssemblyPartInfo) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPartInfo) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *BTAssemblyPartInfo) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *BTAssemblyPartInfo) SetRevision(v string) {
	o.Revision = &v
}

func (o BTAssemblyPartInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BodyType != nil {
		toSerialize["bodyType"] = o.BodyType
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.DocumentMicroversion != nil {
		toSerialize["documentMicroversion"] = o.DocumentMicroversion
	}
	if o.DocumentVersion != nil {
		toSerialize["documentVersion"] = o.DocumentVersion
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.FullConfiguration != nil {
		toSerialize["fullConfiguration"] = o.FullConfiguration
	}
	if o.IsStandardContent != nil {
		toSerialize["isStandardContent"] = o.IsStandardContent
	}
	if o.MateConnectors != nil {
		toSerialize["mateConnectors"] = o.MateConnectors
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyPartInfo struct {
	value *BTAssemblyPartInfo
	isSet bool
}

func (v NullableBTAssemblyPartInfo) Get() *BTAssemblyPartInfo {
	return v.value
}

func (v *NullableBTAssemblyPartInfo) Set(val *BTAssemblyPartInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyPartInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyPartInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyPartInfo(val *BTAssemblyPartInfo) *NullableBTAssemblyPartInfo {
	return &NullableBTAssemblyPartInfo{value: val, isSet: true}
}

func (v NullableBTAssemblyPartInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyPartInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
