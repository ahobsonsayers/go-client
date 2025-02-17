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

// BTAppElementReferenceParams struct for BTAppElementReferenceParams
type BTAppElementReferenceParams struct {
	HasDocumentMicroversions *bool    `json:"hasDocumentMicroversions,omitempty"`
	IdTag                    *string  `json:"idTag,omitempty"`
	IdTagMicroversionId      *string  `json:"idTagMicroversionId,omitempty"`
	IsLocked                 *bool    `json:"isLocked,omitempty"`
	IsSketchOnly             *bool    `json:"isSketchOnly,omitempty"`
	ParentChangeId           *string  `json:"parentChangeId,omitempty"`
	ParentViewId             *string  `json:"parentViewId,omitempty"`
	PartNumber               *string  `json:"partNumber,omitempty"`
	PureSketch               *bool    `json:"pureSketch,omitempty"`
	ReferenceType            *int32   `json:"referenceType,omitempty"`
	ReturnError              *bool    `json:"returnError,omitempty"`
	Revision                 *string  `json:"revision,omitempty"`
	SketchIds                []string `json:"sketchIds,omitempty"`
	TargetConfiguration      *string  `json:"targetConfiguration,omitempty"`
	TargetDocumentId         *string  `json:"targetDocumentId,omitempty"`
	TargetElementId          *string  `json:"targetElementId,omitempty"`
	TargetMicroversionId     *string  `json:"targetMicroversionId,omitempty"`
	TargetVersionId          *string  `json:"targetVersionId,omitempty"`
	TrackNewVersions         *bool    `json:"trackNewVersions,omitempty"`
	TransactionId            *string  `json:"transactionId,omitempty"`
	UpdateSketchInfo         *bool    `json:"updateSketchInfo,omitempty"`
}

// NewBTAppElementReferenceParams instantiates a new BTAppElementReferenceParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementReferenceParams() *BTAppElementReferenceParams {
	this := BTAppElementReferenceParams{}
	return &this
}

// NewBTAppElementReferenceParamsWithDefaults instantiates a new BTAppElementReferenceParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementReferenceParamsWithDefaults() *BTAppElementReferenceParams {
	this := BTAppElementReferenceParams{}
	return &this
}

// GetHasDocumentMicroversions returns the HasDocumentMicroversions field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetHasDocumentMicroversions() bool {
	if o == nil || o.HasDocumentMicroversions == nil {
		var ret bool
		return ret
	}
	return *o.HasDocumentMicroversions
}

// GetHasDocumentMicroversionsOk returns a tuple with the HasDocumentMicroversions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetHasDocumentMicroversionsOk() (*bool, bool) {
	if o == nil || o.HasDocumentMicroversions == nil {
		return nil, false
	}
	return o.HasDocumentMicroversions, true
}

// HasHasDocumentMicroversions returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasHasDocumentMicroversions() bool {
	if o != nil && o.HasDocumentMicroversions != nil {
		return true
	}

	return false
}

// SetHasDocumentMicroversions gets a reference to the given bool and assigns it to the HasDocumentMicroversions field.
func (o *BTAppElementReferenceParams) SetHasDocumentMicroversions(v bool) {
	o.HasDocumentMicroversions = &v
}

// GetIdTag returns the IdTag field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetIdTag() string {
	if o == nil || o.IdTag == nil {
		var ret string
		return ret
	}
	return *o.IdTag
}

// GetIdTagOk returns a tuple with the IdTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetIdTagOk() (*string, bool) {
	if o == nil || o.IdTag == nil {
		return nil, false
	}
	return o.IdTag, true
}

// HasIdTag returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasIdTag() bool {
	if o != nil && o.IdTag != nil {
		return true
	}

	return false
}

// SetIdTag gets a reference to the given string and assigns it to the IdTag field.
func (o *BTAppElementReferenceParams) SetIdTag(v string) {
	o.IdTag = &v
}

// GetIdTagMicroversionId returns the IdTagMicroversionId field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetIdTagMicroversionId() string {
	if o == nil || o.IdTagMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.IdTagMicroversionId
}

// GetIdTagMicroversionIdOk returns a tuple with the IdTagMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetIdTagMicroversionIdOk() (*string, bool) {
	if o == nil || o.IdTagMicroversionId == nil {
		return nil, false
	}
	return o.IdTagMicroversionId, true
}

// HasIdTagMicroversionId returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasIdTagMicroversionId() bool {
	if o != nil && o.IdTagMicroversionId != nil {
		return true
	}

	return false
}

// SetIdTagMicroversionId gets a reference to the given string and assigns it to the IdTagMicroversionId field.
func (o *BTAppElementReferenceParams) SetIdTagMicroversionId(v string) {
	o.IdTagMicroversionId = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetIsLocked() bool {
	if o == nil || o.IsLocked == nil {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetIsLockedOk() (*bool, bool) {
	if o == nil || o.IsLocked == nil {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasIsLocked() bool {
	if o != nil && o.IsLocked != nil {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *BTAppElementReferenceParams) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetIsSketchOnly returns the IsSketchOnly field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetIsSketchOnly() bool {
	if o == nil || o.IsSketchOnly == nil {
		var ret bool
		return ret
	}
	return *o.IsSketchOnly
}

// GetIsSketchOnlyOk returns a tuple with the IsSketchOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetIsSketchOnlyOk() (*bool, bool) {
	if o == nil || o.IsSketchOnly == nil {
		return nil, false
	}
	return o.IsSketchOnly, true
}

// HasIsSketchOnly returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasIsSketchOnly() bool {
	if o != nil && o.IsSketchOnly != nil {
		return true
	}

	return false
}

// SetIsSketchOnly gets a reference to the given bool and assigns it to the IsSketchOnly field.
func (o *BTAppElementReferenceParams) SetIsSketchOnly(v bool) {
	o.IsSketchOnly = &v
}

// GetParentChangeId returns the ParentChangeId field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetParentChangeId() string {
	if o == nil || o.ParentChangeId == nil {
		var ret string
		return ret
	}
	return *o.ParentChangeId
}

// GetParentChangeIdOk returns a tuple with the ParentChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetParentChangeIdOk() (*string, bool) {
	if o == nil || o.ParentChangeId == nil {
		return nil, false
	}
	return o.ParentChangeId, true
}

// HasParentChangeId returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasParentChangeId() bool {
	if o != nil && o.ParentChangeId != nil {
		return true
	}

	return false
}

// SetParentChangeId gets a reference to the given string and assigns it to the ParentChangeId field.
func (o *BTAppElementReferenceParams) SetParentChangeId(v string) {
	o.ParentChangeId = &v
}

// GetParentViewId returns the ParentViewId field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetParentViewId() string {
	if o == nil || o.ParentViewId == nil {
		var ret string
		return ret
	}
	return *o.ParentViewId
}

// GetParentViewIdOk returns a tuple with the ParentViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetParentViewIdOk() (*string, bool) {
	if o == nil || o.ParentViewId == nil {
		return nil, false
	}
	return o.ParentViewId, true
}

// HasParentViewId returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasParentViewId() bool {
	if o != nil && o.ParentViewId != nil {
		return true
	}

	return false
}

// SetParentViewId gets a reference to the given string and assigns it to the ParentViewId field.
func (o *BTAppElementReferenceParams) SetParentViewId(v string) {
	o.ParentViewId = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTAppElementReferenceParams) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPureSketch returns the PureSketch field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetPureSketch() bool {
	if o == nil || o.PureSketch == nil {
		var ret bool
		return ret
	}
	return *o.PureSketch
}

// GetPureSketchOk returns a tuple with the PureSketch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetPureSketchOk() (*bool, bool) {
	if o == nil || o.PureSketch == nil {
		return nil, false
	}
	return o.PureSketch, true
}

// HasPureSketch returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasPureSketch() bool {
	if o != nil && o.PureSketch != nil {
		return true
	}

	return false
}

// SetPureSketch gets a reference to the given bool and assigns it to the PureSketch field.
func (o *BTAppElementReferenceParams) SetPureSketch(v bool) {
	o.PureSketch = &v
}

// GetReferenceType returns the ReferenceType field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetReferenceType() int32 {
	if o == nil || o.ReferenceType == nil {
		var ret int32
		return ret
	}
	return *o.ReferenceType
}

// GetReferenceTypeOk returns a tuple with the ReferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetReferenceTypeOk() (*int32, bool) {
	if o == nil || o.ReferenceType == nil {
		return nil, false
	}
	return o.ReferenceType, true
}

// HasReferenceType returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasReferenceType() bool {
	if o != nil && o.ReferenceType != nil {
		return true
	}

	return false
}

// SetReferenceType gets a reference to the given int32 and assigns it to the ReferenceType field.
func (o *BTAppElementReferenceParams) SetReferenceType(v int32) {
	o.ReferenceType = &v
}

// GetReturnError returns the ReturnError field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetReturnError() bool {
	if o == nil || o.ReturnError == nil {
		var ret bool
		return ret
	}
	return *o.ReturnError
}

// GetReturnErrorOk returns a tuple with the ReturnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetReturnErrorOk() (*bool, bool) {
	if o == nil || o.ReturnError == nil {
		return nil, false
	}
	return o.ReturnError, true
}

// HasReturnError returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasReturnError() bool {
	if o != nil && o.ReturnError != nil {
		return true
	}

	return false
}

// SetReturnError gets a reference to the given bool and assigns it to the ReturnError field.
func (o *BTAppElementReferenceParams) SetReturnError(v bool) {
	o.ReturnError = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *BTAppElementReferenceParams) SetRevision(v string) {
	o.Revision = &v
}

// GetSketchIds returns the SketchIds field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetSketchIds() []string {
	if o == nil || o.SketchIds == nil {
		var ret []string
		return ret
	}
	return o.SketchIds
}

// GetSketchIdsOk returns a tuple with the SketchIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetSketchIdsOk() ([]string, bool) {
	if o == nil || o.SketchIds == nil {
		return nil, false
	}
	return o.SketchIds, true
}

// HasSketchIds returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasSketchIds() bool {
	if o != nil && o.SketchIds != nil {
		return true
	}

	return false
}

// SetSketchIds gets a reference to the given []string and assigns it to the SketchIds field.
func (o *BTAppElementReferenceParams) SetSketchIds(v []string) {
	o.SketchIds = v
}

// GetTargetConfiguration returns the TargetConfiguration field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetTargetConfiguration() string {
	if o == nil || o.TargetConfiguration == nil {
		var ret string
		return ret
	}
	return *o.TargetConfiguration
}

// GetTargetConfigurationOk returns a tuple with the TargetConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetTargetConfigurationOk() (*string, bool) {
	if o == nil || o.TargetConfiguration == nil {
		return nil, false
	}
	return o.TargetConfiguration, true
}

// HasTargetConfiguration returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasTargetConfiguration() bool {
	if o != nil && o.TargetConfiguration != nil {
		return true
	}

	return false
}

// SetTargetConfiguration gets a reference to the given string and assigns it to the TargetConfiguration field.
func (o *BTAppElementReferenceParams) SetTargetConfiguration(v string) {
	o.TargetConfiguration = &v
}

// GetTargetDocumentId returns the TargetDocumentId field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetTargetDocumentId() string {
	if o == nil || o.TargetDocumentId == nil {
		var ret string
		return ret
	}
	return *o.TargetDocumentId
}

// GetTargetDocumentIdOk returns a tuple with the TargetDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetTargetDocumentIdOk() (*string, bool) {
	if o == nil || o.TargetDocumentId == nil {
		return nil, false
	}
	return o.TargetDocumentId, true
}

// HasTargetDocumentId returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasTargetDocumentId() bool {
	if o != nil && o.TargetDocumentId != nil {
		return true
	}

	return false
}

// SetTargetDocumentId gets a reference to the given string and assigns it to the TargetDocumentId field.
func (o *BTAppElementReferenceParams) SetTargetDocumentId(v string) {
	o.TargetDocumentId = &v
}

// GetTargetElementId returns the TargetElementId field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetTargetElementId() string {
	if o == nil || o.TargetElementId == nil {
		var ret string
		return ret
	}
	return *o.TargetElementId
}

// GetTargetElementIdOk returns a tuple with the TargetElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetTargetElementIdOk() (*string, bool) {
	if o == nil || o.TargetElementId == nil {
		return nil, false
	}
	return o.TargetElementId, true
}

// HasTargetElementId returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasTargetElementId() bool {
	if o != nil && o.TargetElementId != nil {
		return true
	}

	return false
}

// SetTargetElementId gets a reference to the given string and assigns it to the TargetElementId field.
func (o *BTAppElementReferenceParams) SetTargetElementId(v string) {
	o.TargetElementId = &v
}

// GetTargetMicroversionId returns the TargetMicroversionId field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetTargetMicroversionId() string {
	if o == nil || o.TargetMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.TargetMicroversionId
}

// GetTargetMicroversionIdOk returns a tuple with the TargetMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetTargetMicroversionIdOk() (*string, bool) {
	if o == nil || o.TargetMicroversionId == nil {
		return nil, false
	}
	return o.TargetMicroversionId, true
}

// HasTargetMicroversionId returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasTargetMicroversionId() bool {
	if o != nil && o.TargetMicroversionId != nil {
		return true
	}

	return false
}

// SetTargetMicroversionId gets a reference to the given string and assigns it to the TargetMicroversionId field.
func (o *BTAppElementReferenceParams) SetTargetMicroversionId(v string) {
	o.TargetMicroversionId = &v
}

// GetTargetVersionId returns the TargetVersionId field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetTargetVersionId() string {
	if o == nil || o.TargetVersionId == nil {
		var ret string
		return ret
	}
	return *o.TargetVersionId
}

// GetTargetVersionIdOk returns a tuple with the TargetVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetTargetVersionIdOk() (*string, bool) {
	if o == nil || o.TargetVersionId == nil {
		return nil, false
	}
	return o.TargetVersionId, true
}

// HasTargetVersionId returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasTargetVersionId() bool {
	if o != nil && o.TargetVersionId != nil {
		return true
	}

	return false
}

// SetTargetVersionId gets a reference to the given string and assigns it to the TargetVersionId field.
func (o *BTAppElementReferenceParams) SetTargetVersionId(v string) {
	o.TargetVersionId = &v
}

// GetTrackNewVersions returns the TrackNewVersions field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetTrackNewVersions() bool {
	if o == nil || o.TrackNewVersions == nil {
		var ret bool
		return ret
	}
	return *o.TrackNewVersions
}

// GetTrackNewVersionsOk returns a tuple with the TrackNewVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetTrackNewVersionsOk() (*bool, bool) {
	if o == nil || o.TrackNewVersions == nil {
		return nil, false
	}
	return o.TrackNewVersions, true
}

// HasTrackNewVersions returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasTrackNewVersions() bool {
	if o != nil && o.TrackNewVersions != nil {
		return true
	}

	return false
}

// SetTrackNewVersions gets a reference to the given bool and assigns it to the TrackNewVersions field.
func (o *BTAppElementReferenceParams) SetTrackNewVersions(v bool) {
	o.TrackNewVersions = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetTransactionId() string {
	if o == nil || o.TransactionId == nil {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetTransactionIdOk() (*string, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasTransactionId() bool {
	if o != nil && o.TransactionId != nil {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *BTAppElementReferenceParams) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetUpdateSketchInfo returns the UpdateSketchInfo field value if set, zero value otherwise.
func (o *BTAppElementReferenceParams) GetUpdateSketchInfo() bool {
	if o == nil || o.UpdateSketchInfo == nil {
		var ret bool
		return ret
	}
	return *o.UpdateSketchInfo
}

// GetUpdateSketchInfoOk returns a tuple with the UpdateSketchInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceParams) GetUpdateSketchInfoOk() (*bool, bool) {
	if o == nil || o.UpdateSketchInfo == nil {
		return nil, false
	}
	return o.UpdateSketchInfo, true
}

// HasUpdateSketchInfo returns a boolean if a field has been set.
func (o *BTAppElementReferenceParams) HasUpdateSketchInfo() bool {
	if o != nil && o.UpdateSketchInfo != nil {
		return true
	}

	return false
}

// SetUpdateSketchInfo gets a reference to the given bool and assigns it to the UpdateSketchInfo field.
func (o *BTAppElementReferenceParams) SetUpdateSketchInfo(v bool) {
	o.UpdateSketchInfo = &v
}

func (o BTAppElementReferenceParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HasDocumentMicroversions != nil {
		toSerialize["hasDocumentMicroversions"] = o.HasDocumentMicroversions
	}
	if o.IdTag != nil {
		toSerialize["idTag"] = o.IdTag
	}
	if o.IdTagMicroversionId != nil {
		toSerialize["idTagMicroversionId"] = o.IdTagMicroversionId
	}
	if o.IsLocked != nil {
		toSerialize["isLocked"] = o.IsLocked
	}
	if o.IsSketchOnly != nil {
		toSerialize["isSketchOnly"] = o.IsSketchOnly
	}
	if o.ParentChangeId != nil {
		toSerialize["parentChangeId"] = o.ParentChangeId
	}
	if o.ParentViewId != nil {
		toSerialize["parentViewId"] = o.ParentViewId
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.PureSketch != nil {
		toSerialize["pureSketch"] = o.PureSketch
	}
	if o.ReferenceType != nil {
		toSerialize["referenceType"] = o.ReferenceType
	}
	if o.ReturnError != nil {
		toSerialize["returnError"] = o.ReturnError
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.SketchIds != nil {
		toSerialize["sketchIds"] = o.SketchIds
	}
	if o.TargetConfiguration != nil {
		toSerialize["targetConfiguration"] = o.TargetConfiguration
	}
	if o.TargetDocumentId != nil {
		toSerialize["targetDocumentId"] = o.TargetDocumentId
	}
	if o.TargetElementId != nil {
		toSerialize["targetElementId"] = o.TargetElementId
	}
	if o.TargetMicroversionId != nil {
		toSerialize["targetMicroversionId"] = o.TargetMicroversionId
	}
	if o.TargetVersionId != nil {
		toSerialize["targetVersionId"] = o.TargetVersionId
	}
	if o.TrackNewVersions != nil {
		toSerialize["trackNewVersions"] = o.TrackNewVersions
	}
	if o.TransactionId != nil {
		toSerialize["transactionId"] = o.TransactionId
	}
	if o.UpdateSketchInfo != nil {
		toSerialize["updateSketchInfo"] = o.UpdateSketchInfo
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementReferenceParams struct {
	value *BTAppElementReferenceParams
	isSet bool
}

func (v NullableBTAppElementReferenceParams) Get() *BTAppElementReferenceParams {
	return v.value
}

func (v *NullableBTAppElementReferenceParams) Set(val *BTAppElementReferenceParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementReferenceParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementReferenceParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementReferenceParams(val *BTAppElementReferenceParams) *NullableBTAppElementReferenceParams {
	return &NullableBTAppElementReferenceParams{value: val, isSet: true}
}

func (v NullableBTAppElementReferenceParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementReferenceParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
