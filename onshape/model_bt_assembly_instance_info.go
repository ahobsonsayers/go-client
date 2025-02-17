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

// BTAssemblyInstanceInfo List of instances including those created by patterns and replicates.
type BTAssemblyInstanceInfo struct {
	Configuration        *string                   `json:"configuration,omitempty"`
	DocumentId           *string                   `json:"documentId,omitempty"`
	DocumentMicroversion *string                   `json:"documentMicroversion,omitempty"`
	DocumentVersion      *string                   `json:"documentVersion,omitempty"`
	ElementId            *string                   `json:"elementId,omitempty"`
	FeatureId            *string                   `json:"featureId,omitempty"`
	FullConfiguration    *string                   `json:"fullConfiguration,omitempty"`
	Id                   *string                   `json:"id,omitempty"`
	IsStandardContent    *bool                     `json:"isStandardContent,omitempty"`
	Name                 *string                   `json:"name,omitempty"`
	PartId               *string                   `json:"partId,omitempty"`
	PartNumber           *string                   `json:"partNumber,omitempty"`
	Revision             *string                   `json:"revision,omitempty"`
	Status               *BTAssemblyInstanceStatus `json:"status,omitempty"`
	Suppressed           *bool                     `json:"suppressed,omitempty"`
	Type                 *BTAssemblyInstanceType   `json:"type,omitempty"`
}

// NewBTAssemblyInstanceInfo instantiates a new BTAssemblyInstanceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyInstanceInfo() *BTAssemblyInstanceInfo {
	this := BTAssemblyInstanceInfo{}
	return &this
}

// NewBTAssemblyInstanceInfoWithDefaults instantiates a new BTAssemblyInstanceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyInstanceInfoWithDefaults() *BTAssemblyInstanceInfo {
	this := BTAssemblyInstanceInfo{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BTAssemblyInstanceInfo) GetConfiguration() string {
	if o == nil || o.Configuration == nil {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceInfo) GetConfigurationOk() (*string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BTAssemblyInstanceInfo) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *BTAssemblyInstanceInfo) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTAssemblyInstanceInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTAssemblyInstanceInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTAssemblyInstanceInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetDocumentMicroversion returns the DocumentMicroversion field value if set, zero value otherwise.
func (o *BTAssemblyInstanceInfo) GetDocumentMicroversion() string {
	if o == nil || o.DocumentMicroversion == nil {
		var ret string
		return ret
	}
	return *o.DocumentMicroversion
}

// GetDocumentMicroversionOk returns a tuple with the DocumentMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceInfo) GetDocumentMicroversionOk() (*string, bool) {
	if o == nil || o.DocumentMicroversion == nil {
		return nil, false
	}
	return o.DocumentMicroversion, true
}

// HasDocumentMicroversion returns a boolean if a field has been set.
func (o *BTAssemblyInstanceInfo) HasDocumentMicroversion() bool {
	if o != nil && o.DocumentMicroversion != nil {
		return true
	}

	return false
}

// SetDocumentMicroversion gets a reference to the given string and assigns it to the DocumentMicroversion field.
func (o *BTAssemblyInstanceInfo) SetDocumentMicroversion(v string) {
	o.DocumentMicroversion = &v
}

// GetDocumentVersion returns the DocumentVersion field value if set, zero value otherwise.
func (o *BTAssemblyInstanceInfo) GetDocumentVersion() string {
	if o == nil || o.DocumentVersion == nil {
		var ret string
		return ret
	}
	return *o.DocumentVersion
}

// GetDocumentVersionOk returns a tuple with the DocumentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceInfo) GetDocumentVersionOk() (*string, bool) {
	if o == nil || o.DocumentVersion == nil {
		return nil, false
	}
	return o.DocumentVersion, true
}

// HasDocumentVersion returns a boolean if a field has been set.
func (o *BTAssemblyInstanceInfo) HasDocumentVersion() bool {
	if o != nil && o.DocumentVersion != nil {
		return true
	}

	return false
}

// SetDocumentVersion gets a reference to the given string and assigns it to the DocumentVersion field.
func (o *BTAssemblyInstanceInfo) SetDocumentVersion(v string) {
	o.DocumentVersion = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTAssemblyInstanceInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTAssemblyInstanceInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTAssemblyInstanceInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTAssemblyInstanceInfo) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceInfo) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTAssemblyInstanceInfo) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTAssemblyInstanceInfo) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetFullConfiguration returns the FullConfiguration field value if set, zero value otherwise.
func (o *BTAssemblyInstanceInfo) GetFullConfiguration() string {
	if o == nil || o.FullConfiguration == nil {
		var ret string
		return ret
	}
	return *o.FullConfiguration
}

// GetFullConfigurationOk returns a tuple with the FullConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceInfo) GetFullConfigurationOk() (*string, bool) {
	if o == nil || o.FullConfiguration == nil {
		return nil, false
	}
	return o.FullConfiguration, true
}

// HasFullConfiguration returns a boolean if a field has been set.
func (o *BTAssemblyInstanceInfo) HasFullConfiguration() bool {
	if o != nil && o.FullConfiguration != nil {
		return true
	}

	return false
}

// SetFullConfiguration gets a reference to the given string and assigns it to the FullConfiguration field.
func (o *BTAssemblyInstanceInfo) SetFullConfiguration(v string) {
	o.FullConfiguration = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTAssemblyInstanceInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTAssemblyInstanceInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTAssemblyInstanceInfo) SetId(v string) {
	o.Id = &v
}

// GetIsStandardContent returns the IsStandardContent field value if set, zero value otherwise.
func (o *BTAssemblyInstanceInfo) GetIsStandardContent() bool {
	if o == nil || o.IsStandardContent == nil {
		var ret bool
		return ret
	}
	return *o.IsStandardContent
}

// GetIsStandardContentOk returns a tuple with the IsStandardContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceInfo) GetIsStandardContentOk() (*bool, bool) {
	if o == nil || o.IsStandardContent == nil {
		return nil, false
	}
	return o.IsStandardContent, true
}

// HasIsStandardContent returns a boolean if a field has been set.
func (o *BTAssemblyInstanceInfo) HasIsStandardContent() bool {
	if o != nil && o.IsStandardContent != nil {
		return true
	}

	return false
}

// SetIsStandardContent gets a reference to the given bool and assigns it to the IsStandardContent field.
func (o *BTAssemblyInstanceInfo) SetIsStandardContent(v bool) {
	o.IsStandardContent = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTAssemblyInstanceInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTAssemblyInstanceInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTAssemblyInstanceInfo) SetName(v string) {
	o.Name = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTAssemblyInstanceInfo) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceInfo) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTAssemblyInstanceInfo) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTAssemblyInstanceInfo) SetPartId(v string) {
	o.PartId = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTAssemblyInstanceInfo) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceInfo) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTAssemblyInstanceInfo) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTAssemblyInstanceInfo) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *BTAssemblyInstanceInfo) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceInfo) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *BTAssemblyInstanceInfo) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *BTAssemblyInstanceInfo) SetRevision(v string) {
	o.Revision = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BTAssemblyInstanceInfo) GetStatus() BTAssemblyInstanceStatus {
	if o == nil || o.Status == nil {
		var ret BTAssemblyInstanceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceInfo) GetStatusOk() (*BTAssemblyInstanceStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BTAssemblyInstanceInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BTAssemblyInstanceStatus and assigns it to the Status field.
func (o *BTAssemblyInstanceInfo) SetStatus(v BTAssemblyInstanceStatus) {
	o.Status = &v
}

// GetSuppressed returns the Suppressed field value if set, zero value otherwise.
func (o *BTAssemblyInstanceInfo) GetSuppressed() bool {
	if o == nil || o.Suppressed == nil {
		var ret bool
		return ret
	}
	return *o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceInfo) GetSuppressedOk() (*bool, bool) {
	if o == nil || o.Suppressed == nil {
		return nil, false
	}
	return o.Suppressed, true
}

// HasSuppressed returns a boolean if a field has been set.
func (o *BTAssemblyInstanceInfo) HasSuppressed() bool {
	if o != nil && o.Suppressed != nil {
		return true
	}

	return false
}

// SetSuppressed gets a reference to the given bool and assigns it to the Suppressed field.
func (o *BTAssemblyInstanceInfo) SetSuppressed(v bool) {
	o.Suppressed = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTAssemblyInstanceInfo) GetType() BTAssemblyInstanceType {
	if o == nil || o.Type == nil {
		var ret BTAssemblyInstanceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceInfo) GetTypeOk() (*BTAssemblyInstanceType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTAssemblyInstanceInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given BTAssemblyInstanceType and assigns it to the Type field.
func (o *BTAssemblyInstanceInfo) SetType(v BTAssemblyInstanceType) {
	o.Type = &v
}

func (o BTAssemblyInstanceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.FullConfiguration != nil {
		toSerialize["fullConfiguration"] = o.FullConfiguration
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsStandardContent != nil {
		toSerialize["isStandardContent"] = o.IsStandardContent
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Suppressed != nil {
		toSerialize["suppressed"] = o.Suppressed
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyInstanceInfo struct {
	value *BTAssemblyInstanceInfo
	isSet bool
}

func (v NullableBTAssemblyInstanceInfo) Get() *BTAssemblyInstanceInfo {
	return v.value
}

func (v *NullableBTAssemblyInstanceInfo) Set(val *BTAssemblyInstanceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyInstanceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyInstanceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyInstanceInfo(val *BTAssemblyInstanceInfo) *NullableBTAssemblyInstanceInfo {
	return &NullableBTAssemblyInstanceInfo{value: val, isSet: true}
}

func (v NullableBTAssemblyInstanceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyInstanceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
