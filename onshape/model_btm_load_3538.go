/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5884-a8034368dd2e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMLoad3538 struct for BTMLoad3538
type BTMLoad3538 struct {
	FeatureId                              *string                                   `json:"featureId,omitempty"`
	FeatureType                            *string                                   `json:"featureType,omitempty"`
	ImportMicroversion                     *string                                   `json:"importMicroversion,omitempty"`
	Name                                   *string                                   `json:"name,omitempty"`
	Namespace                              *string                                   `json:"namespace,omitempty"`
	NodeId                                 *string                                   `json:"nodeId,omitempty"`
	Parameters                             []BTMParameter1                           `json:"parameters,omitempty"`
	ReturnAfterSubfeatures                 *bool                                     `json:"returnAfterSubfeatures,omitempty"`
	SubFeatures                            []BTMFeature134                           `json:"subFeatures,omitempty"`
	Suppressed                             *bool                                     `json:"suppressed,omitempty"`
	SuppressionConfigured                  *bool                                     `json:"suppressionConfigured,omitempty"`
	VariableStudioReference                *bool                                     `json:"variableStudioReference,omitempty"`
	AuxiliaryTreeFeature                   *bool                                     `json:"auxiliaryTreeFeature,omitempty"`
	BtType                                 *string                                   `json:"btType,omitempty"`
	FeatureFolder                          *bool                                     `json:"featureFolder,omitempty"`
	FeatureListFieldIndex                  *int32                                    `json:"featureListFieldIndex,omitempty"`
	FieldIndexForOwnedMateConnectors       *int32                                    `json:"fieldIndexForOwnedMateConnectors,omitempty"`
	OccurrenceQueriesFromAllConfigurations []BTMIndividualQueryWithOccurrenceBase904 `json:"occurrenceQueriesFromAllConfigurations,omitempty"`
	ParametricInstanceFeature              *bool                                     `json:"parametricInstanceFeature,omitempty"`
	Version                                *int32                                    `json:"version,omitempty"`
	DefinedByComponents                    *bool                                     `json:"definedByComponents,omitempty"`
	LoadComponentParameterIds              *map[string]string                        `json:"loadComponentParameterIds,omitempty"`
	LoadRegionParameterId                  *string                                   `json:"loadRegionParameterId,omitempty"`
	LoadType                               *string                                   `json:"loadType,omitempty"`
	MagnitudeParameterId                   *string                                   `json:"magnitudeParameterId,omitempty"`
	MagnitudeQuantityType                  *string                                   `json:"magnitudeQuantityType,omitempty"`
	StructuralLoad                         *bool                                     `json:"structuralLoad,omitempty"`
	SuppressedInSimulations                *map[string]int32                         `json:"suppressedInSimulations,omitempty"`
}

// NewBTMLoad3538 instantiates a new BTMLoad3538 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMLoad3538() *BTMLoad3538 {
	this := BTMLoad3538{}
	return &this
}

// NewBTMLoad3538WithDefaults instantiates a new BTMLoad3538 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMLoad3538WithDefaults() *BTMLoad3538 {
	this := BTMLoad3538{}
	return &this
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTMLoad3538) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTMLoad3538) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTMLoad3538) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise.
func (o *BTMLoad3538) GetFeatureType() string {
	if o == nil || o.FeatureType == nil {
		var ret string
		return ret
	}
	return *o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetFeatureTypeOk() (*string, bool) {
	if o == nil || o.FeatureType == nil {
		return nil, false
	}
	return o.FeatureType, true
}

// HasFeatureType returns a boolean if a field has been set.
func (o *BTMLoad3538) HasFeatureType() bool {
	if o != nil && o.FeatureType != nil {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given string and assigns it to the FeatureType field.
func (o *BTMLoad3538) SetFeatureType(v string) {
	o.FeatureType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMLoad3538) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMLoad3538) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMLoad3538) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTMLoad3538) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTMLoad3538) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTMLoad3538) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMLoad3538) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMLoad3538) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMLoad3538) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMLoad3538) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMLoad3538) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMLoad3538) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTMLoad3538) GetParameters() []BTMParameter1 {
	if o == nil || o.Parameters == nil {
		var ret []BTMParameter1
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetParametersOk() ([]BTMParameter1, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BTMLoad3538) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []BTMParameter1 and assigns it to the Parameters field.
func (o *BTMLoad3538) SetParameters(v []BTMParameter1) {
	o.Parameters = v
}

// GetReturnAfterSubfeatures returns the ReturnAfterSubfeatures field value if set, zero value otherwise.
func (o *BTMLoad3538) GetReturnAfterSubfeatures() bool {
	if o == nil || o.ReturnAfterSubfeatures == nil {
		var ret bool
		return ret
	}
	return *o.ReturnAfterSubfeatures
}

// GetReturnAfterSubfeaturesOk returns a tuple with the ReturnAfterSubfeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetReturnAfterSubfeaturesOk() (*bool, bool) {
	if o == nil || o.ReturnAfterSubfeatures == nil {
		return nil, false
	}
	return o.ReturnAfterSubfeatures, true
}

// HasReturnAfterSubfeatures returns a boolean if a field has been set.
func (o *BTMLoad3538) HasReturnAfterSubfeatures() bool {
	if o != nil && o.ReturnAfterSubfeatures != nil {
		return true
	}

	return false
}

// SetReturnAfterSubfeatures gets a reference to the given bool and assigns it to the ReturnAfterSubfeatures field.
func (o *BTMLoad3538) SetReturnAfterSubfeatures(v bool) {
	o.ReturnAfterSubfeatures = &v
}

// GetSubFeatures returns the SubFeatures field value if set, zero value otherwise.
func (o *BTMLoad3538) GetSubFeatures() []BTMFeature134 {
	if o == nil || o.SubFeatures == nil {
		var ret []BTMFeature134
		return ret
	}
	return o.SubFeatures
}

// GetSubFeaturesOk returns a tuple with the SubFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetSubFeaturesOk() ([]BTMFeature134, bool) {
	if o == nil || o.SubFeatures == nil {
		return nil, false
	}
	return o.SubFeatures, true
}

// HasSubFeatures returns a boolean if a field has been set.
func (o *BTMLoad3538) HasSubFeatures() bool {
	if o != nil && o.SubFeatures != nil {
		return true
	}

	return false
}

// SetSubFeatures gets a reference to the given []BTMFeature134 and assigns it to the SubFeatures field.
func (o *BTMLoad3538) SetSubFeatures(v []BTMFeature134) {
	o.SubFeatures = v
}

// GetSuppressed returns the Suppressed field value if set, zero value otherwise.
func (o *BTMLoad3538) GetSuppressed() bool {
	if o == nil || o.Suppressed == nil {
		var ret bool
		return ret
	}
	return *o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetSuppressedOk() (*bool, bool) {
	if o == nil || o.Suppressed == nil {
		return nil, false
	}
	return o.Suppressed, true
}

// HasSuppressed returns a boolean if a field has been set.
func (o *BTMLoad3538) HasSuppressed() bool {
	if o != nil && o.Suppressed != nil {
		return true
	}

	return false
}

// SetSuppressed gets a reference to the given bool and assigns it to the Suppressed field.
func (o *BTMLoad3538) SetSuppressed(v bool) {
	o.Suppressed = &v
}

// GetSuppressionConfigured returns the SuppressionConfigured field value if set, zero value otherwise.
func (o *BTMLoad3538) GetSuppressionConfigured() bool {
	if o == nil || o.SuppressionConfigured == nil {
		var ret bool
		return ret
	}
	return *o.SuppressionConfigured
}

// GetSuppressionConfiguredOk returns a tuple with the SuppressionConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetSuppressionConfiguredOk() (*bool, bool) {
	if o == nil || o.SuppressionConfigured == nil {
		return nil, false
	}
	return o.SuppressionConfigured, true
}

// HasSuppressionConfigured returns a boolean if a field has been set.
func (o *BTMLoad3538) HasSuppressionConfigured() bool {
	if o != nil && o.SuppressionConfigured != nil {
		return true
	}

	return false
}

// SetSuppressionConfigured gets a reference to the given bool and assigns it to the SuppressionConfigured field.
func (o *BTMLoad3538) SetSuppressionConfigured(v bool) {
	o.SuppressionConfigured = &v
}

// GetVariableStudioReference returns the VariableStudioReference field value if set, zero value otherwise.
func (o *BTMLoad3538) GetVariableStudioReference() bool {
	if o == nil || o.VariableStudioReference == nil {
		var ret bool
		return ret
	}
	return *o.VariableStudioReference
}

// GetVariableStudioReferenceOk returns a tuple with the VariableStudioReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetVariableStudioReferenceOk() (*bool, bool) {
	if o == nil || o.VariableStudioReference == nil {
		return nil, false
	}
	return o.VariableStudioReference, true
}

// HasVariableStudioReference returns a boolean if a field has been set.
func (o *BTMLoad3538) HasVariableStudioReference() bool {
	if o != nil && o.VariableStudioReference != nil {
		return true
	}

	return false
}

// SetVariableStudioReference gets a reference to the given bool and assigns it to the VariableStudioReference field.
func (o *BTMLoad3538) SetVariableStudioReference(v bool) {
	o.VariableStudioReference = &v
}

// GetAuxiliaryTreeFeature returns the AuxiliaryTreeFeature field value if set, zero value otherwise.
func (o *BTMLoad3538) GetAuxiliaryTreeFeature() bool {
	if o == nil || o.AuxiliaryTreeFeature == nil {
		var ret bool
		return ret
	}
	return *o.AuxiliaryTreeFeature
}

// GetAuxiliaryTreeFeatureOk returns a tuple with the AuxiliaryTreeFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetAuxiliaryTreeFeatureOk() (*bool, bool) {
	if o == nil || o.AuxiliaryTreeFeature == nil {
		return nil, false
	}
	return o.AuxiliaryTreeFeature, true
}

// HasAuxiliaryTreeFeature returns a boolean if a field has been set.
func (o *BTMLoad3538) HasAuxiliaryTreeFeature() bool {
	if o != nil && o.AuxiliaryTreeFeature != nil {
		return true
	}

	return false
}

// SetAuxiliaryTreeFeature gets a reference to the given bool and assigns it to the AuxiliaryTreeFeature field.
func (o *BTMLoad3538) SetAuxiliaryTreeFeature(v bool) {
	o.AuxiliaryTreeFeature = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMLoad3538) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMLoad3538) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMLoad3538) SetBtType(v string) {
	o.BtType = &v
}

// GetFeatureFolder returns the FeatureFolder field value if set, zero value otherwise.
func (o *BTMLoad3538) GetFeatureFolder() bool {
	if o == nil || o.FeatureFolder == nil {
		var ret bool
		return ret
	}
	return *o.FeatureFolder
}

// GetFeatureFolderOk returns a tuple with the FeatureFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetFeatureFolderOk() (*bool, bool) {
	if o == nil || o.FeatureFolder == nil {
		return nil, false
	}
	return o.FeatureFolder, true
}

// HasFeatureFolder returns a boolean if a field has been set.
func (o *BTMLoad3538) HasFeatureFolder() bool {
	if o != nil && o.FeatureFolder != nil {
		return true
	}

	return false
}

// SetFeatureFolder gets a reference to the given bool and assigns it to the FeatureFolder field.
func (o *BTMLoad3538) SetFeatureFolder(v bool) {
	o.FeatureFolder = &v
}

// GetFeatureListFieldIndex returns the FeatureListFieldIndex field value if set, zero value otherwise.
func (o *BTMLoad3538) GetFeatureListFieldIndex() int32 {
	if o == nil || o.FeatureListFieldIndex == nil {
		var ret int32
		return ret
	}
	return *o.FeatureListFieldIndex
}

// GetFeatureListFieldIndexOk returns a tuple with the FeatureListFieldIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetFeatureListFieldIndexOk() (*int32, bool) {
	if o == nil || o.FeatureListFieldIndex == nil {
		return nil, false
	}
	return o.FeatureListFieldIndex, true
}

// HasFeatureListFieldIndex returns a boolean if a field has been set.
func (o *BTMLoad3538) HasFeatureListFieldIndex() bool {
	if o != nil && o.FeatureListFieldIndex != nil {
		return true
	}

	return false
}

// SetFeatureListFieldIndex gets a reference to the given int32 and assigns it to the FeatureListFieldIndex field.
func (o *BTMLoad3538) SetFeatureListFieldIndex(v int32) {
	o.FeatureListFieldIndex = &v
}

// GetFieldIndexForOwnedMateConnectors returns the FieldIndexForOwnedMateConnectors field value if set, zero value otherwise.
func (o *BTMLoad3538) GetFieldIndexForOwnedMateConnectors() int32 {
	if o == nil || o.FieldIndexForOwnedMateConnectors == nil {
		var ret int32
		return ret
	}
	return *o.FieldIndexForOwnedMateConnectors
}

// GetFieldIndexForOwnedMateConnectorsOk returns a tuple with the FieldIndexForOwnedMateConnectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetFieldIndexForOwnedMateConnectorsOk() (*int32, bool) {
	if o == nil || o.FieldIndexForOwnedMateConnectors == nil {
		return nil, false
	}
	return o.FieldIndexForOwnedMateConnectors, true
}

// HasFieldIndexForOwnedMateConnectors returns a boolean if a field has been set.
func (o *BTMLoad3538) HasFieldIndexForOwnedMateConnectors() bool {
	if o != nil && o.FieldIndexForOwnedMateConnectors != nil {
		return true
	}

	return false
}

// SetFieldIndexForOwnedMateConnectors gets a reference to the given int32 and assigns it to the FieldIndexForOwnedMateConnectors field.
func (o *BTMLoad3538) SetFieldIndexForOwnedMateConnectors(v int32) {
	o.FieldIndexForOwnedMateConnectors = &v
}

// GetOccurrenceQueriesFromAllConfigurations returns the OccurrenceQueriesFromAllConfigurations field value if set, zero value otherwise.
func (o *BTMLoad3538) GetOccurrenceQueriesFromAllConfigurations() []BTMIndividualQueryWithOccurrenceBase904 {
	if o == nil || o.OccurrenceQueriesFromAllConfigurations == nil {
		var ret []BTMIndividualQueryWithOccurrenceBase904
		return ret
	}
	return o.OccurrenceQueriesFromAllConfigurations
}

// GetOccurrenceQueriesFromAllConfigurationsOk returns a tuple with the OccurrenceQueriesFromAllConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetOccurrenceQueriesFromAllConfigurationsOk() ([]BTMIndividualQueryWithOccurrenceBase904, bool) {
	if o == nil || o.OccurrenceQueriesFromAllConfigurations == nil {
		return nil, false
	}
	return o.OccurrenceQueriesFromAllConfigurations, true
}

// HasOccurrenceQueriesFromAllConfigurations returns a boolean if a field has been set.
func (o *BTMLoad3538) HasOccurrenceQueriesFromAllConfigurations() bool {
	if o != nil && o.OccurrenceQueriesFromAllConfigurations != nil {
		return true
	}

	return false
}

// SetOccurrenceQueriesFromAllConfigurations gets a reference to the given []BTMIndividualQueryWithOccurrenceBase904 and assigns it to the OccurrenceQueriesFromAllConfigurations field.
func (o *BTMLoad3538) SetOccurrenceQueriesFromAllConfigurations(v []BTMIndividualQueryWithOccurrenceBase904) {
	o.OccurrenceQueriesFromAllConfigurations = v
}

// GetParametricInstanceFeature returns the ParametricInstanceFeature field value if set, zero value otherwise.
func (o *BTMLoad3538) GetParametricInstanceFeature() bool {
	if o == nil || o.ParametricInstanceFeature == nil {
		var ret bool
		return ret
	}
	return *o.ParametricInstanceFeature
}

// GetParametricInstanceFeatureOk returns a tuple with the ParametricInstanceFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetParametricInstanceFeatureOk() (*bool, bool) {
	if o == nil || o.ParametricInstanceFeature == nil {
		return nil, false
	}
	return o.ParametricInstanceFeature, true
}

// HasParametricInstanceFeature returns a boolean if a field has been set.
func (o *BTMLoad3538) HasParametricInstanceFeature() bool {
	if o != nil && o.ParametricInstanceFeature != nil {
		return true
	}

	return false
}

// SetParametricInstanceFeature gets a reference to the given bool and assigns it to the ParametricInstanceFeature field.
func (o *BTMLoad3538) SetParametricInstanceFeature(v bool) {
	o.ParametricInstanceFeature = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BTMLoad3538) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BTMLoad3538) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *BTMLoad3538) SetVersion(v int32) {
	o.Version = &v
}

// GetDefinedByComponents returns the DefinedByComponents field value if set, zero value otherwise.
func (o *BTMLoad3538) GetDefinedByComponents() bool {
	if o == nil || o.DefinedByComponents == nil {
		var ret bool
		return ret
	}
	return *o.DefinedByComponents
}

// GetDefinedByComponentsOk returns a tuple with the DefinedByComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetDefinedByComponentsOk() (*bool, bool) {
	if o == nil || o.DefinedByComponents == nil {
		return nil, false
	}
	return o.DefinedByComponents, true
}

// HasDefinedByComponents returns a boolean if a field has been set.
func (o *BTMLoad3538) HasDefinedByComponents() bool {
	if o != nil && o.DefinedByComponents != nil {
		return true
	}

	return false
}

// SetDefinedByComponents gets a reference to the given bool and assigns it to the DefinedByComponents field.
func (o *BTMLoad3538) SetDefinedByComponents(v bool) {
	o.DefinedByComponents = &v
}

// GetLoadComponentParameterIds returns the LoadComponentParameterIds field value if set, zero value otherwise.
func (o *BTMLoad3538) GetLoadComponentParameterIds() map[string]string {
	if o == nil || o.LoadComponentParameterIds == nil {
		var ret map[string]string
		return ret
	}
	return *o.LoadComponentParameterIds
}

// GetLoadComponentParameterIdsOk returns a tuple with the LoadComponentParameterIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetLoadComponentParameterIdsOk() (*map[string]string, bool) {
	if o == nil || o.LoadComponentParameterIds == nil {
		return nil, false
	}
	return o.LoadComponentParameterIds, true
}

// HasLoadComponentParameterIds returns a boolean if a field has been set.
func (o *BTMLoad3538) HasLoadComponentParameterIds() bool {
	if o != nil && o.LoadComponentParameterIds != nil {
		return true
	}

	return false
}

// SetLoadComponentParameterIds gets a reference to the given map[string]string and assigns it to the LoadComponentParameterIds field.
func (o *BTMLoad3538) SetLoadComponentParameterIds(v map[string]string) {
	o.LoadComponentParameterIds = &v
}

// GetLoadRegionParameterId returns the LoadRegionParameterId field value if set, zero value otherwise.
func (o *BTMLoad3538) GetLoadRegionParameterId() string {
	if o == nil || o.LoadRegionParameterId == nil {
		var ret string
		return ret
	}
	return *o.LoadRegionParameterId
}

// GetLoadRegionParameterIdOk returns a tuple with the LoadRegionParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetLoadRegionParameterIdOk() (*string, bool) {
	if o == nil || o.LoadRegionParameterId == nil {
		return nil, false
	}
	return o.LoadRegionParameterId, true
}

// HasLoadRegionParameterId returns a boolean if a field has been set.
func (o *BTMLoad3538) HasLoadRegionParameterId() bool {
	if o != nil && o.LoadRegionParameterId != nil {
		return true
	}

	return false
}

// SetLoadRegionParameterId gets a reference to the given string and assigns it to the LoadRegionParameterId field.
func (o *BTMLoad3538) SetLoadRegionParameterId(v string) {
	o.LoadRegionParameterId = &v
}

// GetLoadType returns the LoadType field value if set, zero value otherwise.
func (o *BTMLoad3538) GetLoadType() string {
	if o == nil || o.LoadType == nil {
		var ret string
		return ret
	}
	return *o.LoadType
}

// GetLoadTypeOk returns a tuple with the LoadType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetLoadTypeOk() (*string, bool) {
	if o == nil || o.LoadType == nil {
		return nil, false
	}
	return o.LoadType, true
}

// HasLoadType returns a boolean if a field has been set.
func (o *BTMLoad3538) HasLoadType() bool {
	if o != nil && o.LoadType != nil {
		return true
	}

	return false
}

// SetLoadType gets a reference to the given string and assigns it to the LoadType field.
func (o *BTMLoad3538) SetLoadType(v string) {
	o.LoadType = &v
}

// GetMagnitudeParameterId returns the MagnitudeParameterId field value if set, zero value otherwise.
func (o *BTMLoad3538) GetMagnitudeParameterId() string {
	if o == nil || o.MagnitudeParameterId == nil {
		var ret string
		return ret
	}
	return *o.MagnitudeParameterId
}

// GetMagnitudeParameterIdOk returns a tuple with the MagnitudeParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetMagnitudeParameterIdOk() (*string, bool) {
	if o == nil || o.MagnitudeParameterId == nil {
		return nil, false
	}
	return o.MagnitudeParameterId, true
}

// HasMagnitudeParameterId returns a boolean if a field has been set.
func (o *BTMLoad3538) HasMagnitudeParameterId() bool {
	if o != nil && o.MagnitudeParameterId != nil {
		return true
	}

	return false
}

// SetMagnitudeParameterId gets a reference to the given string and assigns it to the MagnitudeParameterId field.
func (o *BTMLoad3538) SetMagnitudeParameterId(v string) {
	o.MagnitudeParameterId = &v
}

// GetMagnitudeQuantityType returns the MagnitudeQuantityType field value if set, zero value otherwise.
func (o *BTMLoad3538) GetMagnitudeQuantityType() string {
	if o == nil || o.MagnitudeQuantityType == nil {
		var ret string
		return ret
	}
	return *o.MagnitudeQuantityType
}

// GetMagnitudeQuantityTypeOk returns a tuple with the MagnitudeQuantityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetMagnitudeQuantityTypeOk() (*string, bool) {
	if o == nil || o.MagnitudeQuantityType == nil {
		return nil, false
	}
	return o.MagnitudeQuantityType, true
}

// HasMagnitudeQuantityType returns a boolean if a field has been set.
func (o *BTMLoad3538) HasMagnitudeQuantityType() bool {
	if o != nil && o.MagnitudeQuantityType != nil {
		return true
	}

	return false
}

// SetMagnitudeQuantityType gets a reference to the given string and assigns it to the MagnitudeQuantityType field.
func (o *BTMLoad3538) SetMagnitudeQuantityType(v string) {
	o.MagnitudeQuantityType = &v
}

// GetStructuralLoad returns the StructuralLoad field value if set, zero value otherwise.
func (o *BTMLoad3538) GetStructuralLoad() bool {
	if o == nil || o.StructuralLoad == nil {
		var ret bool
		return ret
	}
	return *o.StructuralLoad
}

// GetStructuralLoadOk returns a tuple with the StructuralLoad field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetStructuralLoadOk() (*bool, bool) {
	if o == nil || o.StructuralLoad == nil {
		return nil, false
	}
	return o.StructuralLoad, true
}

// HasStructuralLoad returns a boolean if a field has been set.
func (o *BTMLoad3538) HasStructuralLoad() bool {
	if o != nil && o.StructuralLoad != nil {
		return true
	}

	return false
}

// SetStructuralLoad gets a reference to the given bool and assigns it to the StructuralLoad field.
func (o *BTMLoad3538) SetStructuralLoad(v bool) {
	o.StructuralLoad = &v
}

// GetSuppressedInSimulations returns the SuppressedInSimulations field value if set, zero value otherwise.
func (o *BTMLoad3538) GetSuppressedInSimulations() map[string]int32 {
	if o == nil || o.SuppressedInSimulations == nil {
		var ret map[string]int32
		return ret
	}
	return *o.SuppressedInSimulations
}

// GetSuppressedInSimulationsOk returns a tuple with the SuppressedInSimulations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMLoad3538) GetSuppressedInSimulationsOk() (*map[string]int32, bool) {
	if o == nil || o.SuppressedInSimulations == nil {
		return nil, false
	}
	return o.SuppressedInSimulations, true
}

// HasSuppressedInSimulations returns a boolean if a field has been set.
func (o *BTMLoad3538) HasSuppressedInSimulations() bool {
	if o != nil && o.SuppressedInSimulations != nil {
		return true
	}

	return false
}

// SetSuppressedInSimulations gets a reference to the given map[string]int32 and assigns it to the SuppressedInSimulations field.
func (o *BTMLoad3538) SetSuppressedInSimulations(v map[string]int32) {
	o.SuppressedInSimulations = &v
}

func (o BTMLoad3538) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.FeatureType != nil {
		toSerialize["featureType"] = o.FeatureType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.ReturnAfterSubfeatures != nil {
		toSerialize["returnAfterSubfeatures"] = o.ReturnAfterSubfeatures
	}
	if o.SubFeatures != nil {
		toSerialize["subFeatures"] = o.SubFeatures
	}
	if o.Suppressed != nil {
		toSerialize["suppressed"] = o.Suppressed
	}
	if o.SuppressionConfigured != nil {
		toSerialize["suppressionConfigured"] = o.SuppressionConfigured
	}
	if o.VariableStudioReference != nil {
		toSerialize["variableStudioReference"] = o.VariableStudioReference
	}
	if o.AuxiliaryTreeFeature != nil {
		toSerialize["auxiliaryTreeFeature"] = o.AuxiliaryTreeFeature
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FeatureFolder != nil {
		toSerialize["featureFolder"] = o.FeatureFolder
	}
	if o.FeatureListFieldIndex != nil {
		toSerialize["featureListFieldIndex"] = o.FeatureListFieldIndex
	}
	if o.FieldIndexForOwnedMateConnectors != nil {
		toSerialize["fieldIndexForOwnedMateConnectors"] = o.FieldIndexForOwnedMateConnectors
	}
	if o.OccurrenceQueriesFromAllConfigurations != nil {
		toSerialize["occurrenceQueriesFromAllConfigurations"] = o.OccurrenceQueriesFromAllConfigurations
	}
	if o.ParametricInstanceFeature != nil {
		toSerialize["parametricInstanceFeature"] = o.ParametricInstanceFeature
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.DefinedByComponents != nil {
		toSerialize["definedByComponents"] = o.DefinedByComponents
	}
	if o.LoadComponentParameterIds != nil {
		toSerialize["loadComponentParameterIds"] = o.LoadComponentParameterIds
	}
	if o.LoadRegionParameterId != nil {
		toSerialize["loadRegionParameterId"] = o.LoadRegionParameterId
	}
	if o.LoadType != nil {
		toSerialize["loadType"] = o.LoadType
	}
	if o.MagnitudeParameterId != nil {
		toSerialize["magnitudeParameterId"] = o.MagnitudeParameterId
	}
	if o.MagnitudeQuantityType != nil {
		toSerialize["magnitudeQuantityType"] = o.MagnitudeQuantityType
	}
	if o.StructuralLoad != nil {
		toSerialize["structuralLoad"] = o.StructuralLoad
	}
	if o.SuppressedInSimulations != nil {
		toSerialize["suppressedInSimulations"] = o.SuppressedInSimulations
	}
	return json.Marshal(toSerialize)
}

type NullableBTMLoad3538 struct {
	value *BTMLoad3538
	isSet bool
}

func (v NullableBTMLoad3538) Get() *BTMLoad3538 {
	return v.value
}

func (v *NullableBTMLoad3538) Set(val *BTMLoad3538) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMLoad3538) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMLoad3538) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMLoad3538(val *BTMLoad3538) *NullableBTMLoad3538 {
	return &NullableBTMLoad3538{value: val, isSet: true}
}

func (v NullableBTMLoad3538) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMLoad3538) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
