/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6358-16771368e3d1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTInstanceFolder3627 struct for BTInstanceFolder3627
type BTInstanceFolder3627 struct {
	BtType                      *string                               `json:"btType,omitempty"`
	ImportMicroversion          *string                               `json:"importMicroversion,omitempty"`
	NodeId                      *string                               `json:"nodeId,omitempty"`
	AssemblyInstance            *bool                                 `json:"assemblyInstance,omitempty"`
	AssemblyPattern             *bool                                 `json:"assemblyPattern,omitempty"`
	AssemblyReplicate           *bool                                 `json:"assemblyReplicate,omitempty"`
	ClonedInstance              *bool                                 `json:"clonedInstance,omitempty"`
	CustomData                  *map[string]BTReferenceCustomData1551 `json:"customData,omitempty"`
	InstanceFolder              *bool                                 `json:"instanceFolder,omitempty"`
	InstanceName                *string                               `json:"instanceName,omitempty"`
	IsFlattenedPart             *bool                                 `json:"isFlattenedPart,omitempty"`
	Locked                      *bool                                 `json:"locked,omitempty"`
	MicroversionId              *BTMicroversionId366                  `json:"microversionId,omitempty"`
	ParametricInstance          *bool                                 `json:"parametricInstance,omitempty"`
	PartInstance                *bool                                 `json:"partInstance,omitempty"`
	Releasable                  *bool                                 `json:"releasable,omitempty"`
	RevisionCustomData          *BTRevisionCustomData2090             `json:"revisionCustomData,omitempty"`
	StandardContent             *bool                                 `json:"standardContent,omitempty"`
	StandardContentParametersId *string                               `json:"standardContentParametersId,omitempty"`
	Suppressed                  *bool                                 `json:"suppressed,omitempty"`
	SuppressedFieldIndex        *int32                                `json:"suppressedFieldIndex,omitempty"`
	SuppressionConfigured       *bool                                 `json:"suppressionConfigured,omitempty"`
	SuppressionState            *BTMSuppressionState1924              `json:"suppressionState,omitempty"`
	ValidRevisionReference      *bool                                 `json:"validRevisionReference,omitempty"`
	Version                     *int32                                `json:"version,omitempty"`
	FolderEnd                   *bool                                 `json:"folderEnd,omitempty"`
	FolderStart                 *bool                                 `json:"folderStart,omitempty"`
	Name                        *string                               `json:"name,omitempty"`
	Parameters                  []BTMParameter1                       `json:"parameters,omitempty"`
	StartNodeId                 *string                               `json:"startNodeId,omitempty"`
	StartNodeIdRaw              *BTObjectId                           `json:"startNodeIdRaw,omitempty"`
}

// NewBTInstanceFolder3627 instantiates a new BTInstanceFolder3627 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTInstanceFolder3627() *BTInstanceFolder3627 {
	this := BTInstanceFolder3627{}
	return &this
}

// NewBTInstanceFolder3627WithDefaults instantiates a new BTInstanceFolder3627 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTInstanceFolder3627WithDefaults() *BTInstanceFolder3627 {
	this := BTInstanceFolder3627{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTInstanceFolder3627) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTInstanceFolder3627) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTInstanceFolder3627) SetNodeId(v string) {
	o.NodeId = &v
}

// GetAssemblyInstance returns the AssemblyInstance field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetAssemblyInstance() bool {
	if o == nil || o.AssemblyInstance == nil {
		var ret bool
		return ret
	}
	return *o.AssemblyInstance
}

// GetAssemblyInstanceOk returns a tuple with the AssemblyInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetAssemblyInstanceOk() (*bool, bool) {
	if o == nil || o.AssemblyInstance == nil {
		return nil, false
	}
	return o.AssemblyInstance, true
}

// HasAssemblyInstance returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasAssemblyInstance() bool {
	if o != nil && o.AssemblyInstance != nil {
		return true
	}

	return false
}

// SetAssemblyInstance gets a reference to the given bool and assigns it to the AssemblyInstance field.
func (o *BTInstanceFolder3627) SetAssemblyInstance(v bool) {
	o.AssemblyInstance = &v
}

// GetAssemblyPattern returns the AssemblyPattern field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetAssemblyPattern() bool {
	if o == nil || o.AssemblyPattern == nil {
		var ret bool
		return ret
	}
	return *o.AssemblyPattern
}

// GetAssemblyPatternOk returns a tuple with the AssemblyPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetAssemblyPatternOk() (*bool, bool) {
	if o == nil || o.AssemblyPattern == nil {
		return nil, false
	}
	return o.AssemblyPattern, true
}

// HasAssemblyPattern returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasAssemblyPattern() bool {
	if o != nil && o.AssemblyPattern != nil {
		return true
	}

	return false
}

// SetAssemblyPattern gets a reference to the given bool and assigns it to the AssemblyPattern field.
func (o *BTInstanceFolder3627) SetAssemblyPattern(v bool) {
	o.AssemblyPattern = &v
}

// GetAssemblyReplicate returns the AssemblyReplicate field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetAssemblyReplicate() bool {
	if o == nil || o.AssemblyReplicate == nil {
		var ret bool
		return ret
	}
	return *o.AssemblyReplicate
}

// GetAssemblyReplicateOk returns a tuple with the AssemblyReplicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetAssemblyReplicateOk() (*bool, bool) {
	if o == nil || o.AssemblyReplicate == nil {
		return nil, false
	}
	return o.AssemblyReplicate, true
}

// HasAssemblyReplicate returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasAssemblyReplicate() bool {
	if o != nil && o.AssemblyReplicate != nil {
		return true
	}

	return false
}

// SetAssemblyReplicate gets a reference to the given bool and assigns it to the AssemblyReplicate field.
func (o *BTInstanceFolder3627) SetAssemblyReplicate(v bool) {
	o.AssemblyReplicate = &v
}

// GetClonedInstance returns the ClonedInstance field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetClonedInstance() bool {
	if o == nil || o.ClonedInstance == nil {
		var ret bool
		return ret
	}
	return *o.ClonedInstance
}

// GetClonedInstanceOk returns a tuple with the ClonedInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetClonedInstanceOk() (*bool, bool) {
	if o == nil || o.ClonedInstance == nil {
		return nil, false
	}
	return o.ClonedInstance, true
}

// HasClonedInstance returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasClonedInstance() bool {
	if o != nil && o.ClonedInstance != nil {
		return true
	}

	return false
}

// SetClonedInstance gets a reference to the given bool and assigns it to the ClonedInstance field.
func (o *BTInstanceFolder3627) SetClonedInstance(v bool) {
	o.ClonedInstance = &v
}

// GetCustomData returns the CustomData field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetCustomData() map[string]BTReferenceCustomData1551 {
	if o == nil || o.CustomData == nil {
		var ret map[string]BTReferenceCustomData1551
		return ret
	}
	return *o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetCustomDataOk() (*map[string]BTReferenceCustomData1551, bool) {
	if o == nil || o.CustomData == nil {
		return nil, false
	}
	return o.CustomData, true
}

// HasCustomData returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasCustomData() bool {
	if o != nil && o.CustomData != nil {
		return true
	}

	return false
}

// SetCustomData gets a reference to the given map[string]BTReferenceCustomData1551 and assigns it to the CustomData field.
func (o *BTInstanceFolder3627) SetCustomData(v map[string]BTReferenceCustomData1551) {
	o.CustomData = &v
}

// GetInstanceFolder returns the InstanceFolder field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetInstanceFolder() bool {
	if o == nil || o.InstanceFolder == nil {
		var ret bool
		return ret
	}
	return *o.InstanceFolder
}

// GetInstanceFolderOk returns a tuple with the InstanceFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetInstanceFolderOk() (*bool, bool) {
	if o == nil || o.InstanceFolder == nil {
		return nil, false
	}
	return o.InstanceFolder, true
}

// HasInstanceFolder returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasInstanceFolder() bool {
	if o != nil && o.InstanceFolder != nil {
		return true
	}

	return false
}

// SetInstanceFolder gets a reference to the given bool and assigns it to the InstanceFolder field.
func (o *BTInstanceFolder3627) SetInstanceFolder(v bool) {
	o.InstanceFolder = &v
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetInstanceName() string {
	if o == nil || o.InstanceName == nil {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetInstanceNameOk() (*string, bool) {
	if o == nil || o.InstanceName == nil {
		return nil, false
	}
	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasInstanceName() bool {
	if o != nil && o.InstanceName != nil {
		return true
	}

	return false
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *BTInstanceFolder3627) SetInstanceName(v string) {
	o.InstanceName = &v
}

// GetIsFlattenedPart returns the IsFlattenedPart field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetIsFlattenedPart() bool {
	if o == nil || o.IsFlattenedPart == nil {
		var ret bool
		return ret
	}
	return *o.IsFlattenedPart
}

// GetIsFlattenedPartOk returns a tuple with the IsFlattenedPart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetIsFlattenedPartOk() (*bool, bool) {
	if o == nil || o.IsFlattenedPart == nil {
		return nil, false
	}
	return o.IsFlattenedPart, true
}

// HasIsFlattenedPart returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasIsFlattenedPart() bool {
	if o != nil && o.IsFlattenedPart != nil {
		return true
	}

	return false
}

// SetIsFlattenedPart gets a reference to the given bool and assigns it to the IsFlattenedPart field.
func (o *BTInstanceFolder3627) SetIsFlattenedPart(v bool) {
	o.IsFlattenedPart = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *BTInstanceFolder3627) SetLocked(v bool) {
	o.Locked = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetMicroversionId() BTMicroversionId366 {
	if o == nil || o.MicroversionId == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetMicroversionIdOk() (*BTMicroversionId366, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given BTMicroversionId366 and assigns it to the MicroversionId field.
func (o *BTInstanceFolder3627) SetMicroversionId(v BTMicroversionId366) {
	o.MicroversionId = &v
}

// GetParametricInstance returns the ParametricInstance field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetParametricInstance() bool {
	if o == nil || o.ParametricInstance == nil {
		var ret bool
		return ret
	}
	return *o.ParametricInstance
}

// GetParametricInstanceOk returns a tuple with the ParametricInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetParametricInstanceOk() (*bool, bool) {
	if o == nil || o.ParametricInstance == nil {
		return nil, false
	}
	return o.ParametricInstance, true
}

// HasParametricInstance returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasParametricInstance() bool {
	if o != nil && o.ParametricInstance != nil {
		return true
	}

	return false
}

// SetParametricInstance gets a reference to the given bool and assigns it to the ParametricInstance field.
func (o *BTInstanceFolder3627) SetParametricInstance(v bool) {
	o.ParametricInstance = &v
}

// GetPartInstance returns the PartInstance field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetPartInstance() bool {
	if o == nil || o.PartInstance == nil {
		var ret bool
		return ret
	}
	return *o.PartInstance
}

// GetPartInstanceOk returns a tuple with the PartInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetPartInstanceOk() (*bool, bool) {
	if o == nil || o.PartInstance == nil {
		return nil, false
	}
	return o.PartInstance, true
}

// HasPartInstance returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasPartInstance() bool {
	if o != nil && o.PartInstance != nil {
		return true
	}

	return false
}

// SetPartInstance gets a reference to the given bool and assigns it to the PartInstance field.
func (o *BTInstanceFolder3627) SetPartInstance(v bool) {
	o.PartInstance = &v
}

// GetReleasable returns the Releasable field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetReleasable() bool {
	if o == nil || o.Releasable == nil {
		var ret bool
		return ret
	}
	return *o.Releasable
}

// GetReleasableOk returns a tuple with the Releasable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetReleasableOk() (*bool, bool) {
	if o == nil || o.Releasable == nil {
		return nil, false
	}
	return o.Releasable, true
}

// HasReleasable returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasReleasable() bool {
	if o != nil && o.Releasable != nil {
		return true
	}

	return false
}

// SetReleasable gets a reference to the given bool and assigns it to the Releasable field.
func (o *BTInstanceFolder3627) SetReleasable(v bool) {
	o.Releasable = &v
}

// GetRevisionCustomData returns the RevisionCustomData field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetRevisionCustomData() BTRevisionCustomData2090 {
	if o == nil || o.RevisionCustomData == nil {
		var ret BTRevisionCustomData2090
		return ret
	}
	return *o.RevisionCustomData
}

// GetRevisionCustomDataOk returns a tuple with the RevisionCustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetRevisionCustomDataOk() (*BTRevisionCustomData2090, bool) {
	if o == nil || o.RevisionCustomData == nil {
		return nil, false
	}
	return o.RevisionCustomData, true
}

// HasRevisionCustomData returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasRevisionCustomData() bool {
	if o != nil && o.RevisionCustomData != nil {
		return true
	}

	return false
}

// SetRevisionCustomData gets a reference to the given BTRevisionCustomData2090 and assigns it to the RevisionCustomData field.
func (o *BTInstanceFolder3627) SetRevisionCustomData(v BTRevisionCustomData2090) {
	o.RevisionCustomData = &v
}

// GetStandardContent returns the StandardContent field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetStandardContent() bool {
	if o == nil || o.StandardContent == nil {
		var ret bool
		return ret
	}
	return *o.StandardContent
}

// GetStandardContentOk returns a tuple with the StandardContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetStandardContentOk() (*bool, bool) {
	if o == nil || o.StandardContent == nil {
		return nil, false
	}
	return o.StandardContent, true
}

// HasStandardContent returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasStandardContent() bool {
	if o != nil && o.StandardContent != nil {
		return true
	}

	return false
}

// SetStandardContent gets a reference to the given bool and assigns it to the StandardContent field.
func (o *BTInstanceFolder3627) SetStandardContent(v bool) {
	o.StandardContent = &v
}

// GetStandardContentParametersId returns the StandardContentParametersId field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetStandardContentParametersId() string {
	if o == nil || o.StandardContentParametersId == nil {
		var ret string
		return ret
	}
	return *o.StandardContentParametersId
}

// GetStandardContentParametersIdOk returns a tuple with the StandardContentParametersId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetStandardContentParametersIdOk() (*string, bool) {
	if o == nil || o.StandardContentParametersId == nil {
		return nil, false
	}
	return o.StandardContentParametersId, true
}

// HasStandardContentParametersId returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasStandardContentParametersId() bool {
	if o != nil && o.StandardContentParametersId != nil {
		return true
	}

	return false
}

// SetStandardContentParametersId gets a reference to the given string and assigns it to the StandardContentParametersId field.
func (o *BTInstanceFolder3627) SetStandardContentParametersId(v string) {
	o.StandardContentParametersId = &v
}

// GetSuppressed returns the Suppressed field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetSuppressed() bool {
	if o == nil || o.Suppressed == nil {
		var ret bool
		return ret
	}
	return *o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetSuppressedOk() (*bool, bool) {
	if o == nil || o.Suppressed == nil {
		return nil, false
	}
	return o.Suppressed, true
}

// HasSuppressed returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasSuppressed() bool {
	if o != nil && o.Suppressed != nil {
		return true
	}

	return false
}

// SetSuppressed gets a reference to the given bool and assigns it to the Suppressed field.
func (o *BTInstanceFolder3627) SetSuppressed(v bool) {
	o.Suppressed = &v
}

// GetSuppressedFieldIndex returns the SuppressedFieldIndex field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetSuppressedFieldIndex() int32 {
	if o == nil || o.SuppressedFieldIndex == nil {
		var ret int32
		return ret
	}
	return *o.SuppressedFieldIndex
}

// GetSuppressedFieldIndexOk returns a tuple with the SuppressedFieldIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetSuppressedFieldIndexOk() (*int32, bool) {
	if o == nil || o.SuppressedFieldIndex == nil {
		return nil, false
	}
	return o.SuppressedFieldIndex, true
}

// HasSuppressedFieldIndex returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasSuppressedFieldIndex() bool {
	if o != nil && o.SuppressedFieldIndex != nil {
		return true
	}

	return false
}

// SetSuppressedFieldIndex gets a reference to the given int32 and assigns it to the SuppressedFieldIndex field.
func (o *BTInstanceFolder3627) SetSuppressedFieldIndex(v int32) {
	o.SuppressedFieldIndex = &v
}

// GetSuppressionConfigured returns the SuppressionConfigured field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetSuppressionConfigured() bool {
	if o == nil || o.SuppressionConfigured == nil {
		var ret bool
		return ret
	}
	return *o.SuppressionConfigured
}

// GetSuppressionConfiguredOk returns a tuple with the SuppressionConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetSuppressionConfiguredOk() (*bool, bool) {
	if o == nil || o.SuppressionConfigured == nil {
		return nil, false
	}
	return o.SuppressionConfigured, true
}

// HasSuppressionConfigured returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasSuppressionConfigured() bool {
	if o != nil && o.SuppressionConfigured != nil {
		return true
	}

	return false
}

// SetSuppressionConfigured gets a reference to the given bool and assigns it to the SuppressionConfigured field.
func (o *BTInstanceFolder3627) SetSuppressionConfigured(v bool) {
	o.SuppressionConfigured = &v
}

// GetSuppressionState returns the SuppressionState field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetSuppressionState() BTMSuppressionState1924 {
	if o == nil || o.SuppressionState == nil {
		var ret BTMSuppressionState1924
		return ret
	}
	return *o.SuppressionState
}

// GetSuppressionStateOk returns a tuple with the SuppressionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetSuppressionStateOk() (*BTMSuppressionState1924, bool) {
	if o == nil || o.SuppressionState == nil {
		return nil, false
	}
	return o.SuppressionState, true
}

// HasSuppressionState returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasSuppressionState() bool {
	if o != nil && o.SuppressionState != nil {
		return true
	}

	return false
}

// SetSuppressionState gets a reference to the given BTMSuppressionState1924 and assigns it to the SuppressionState field.
func (o *BTInstanceFolder3627) SetSuppressionState(v BTMSuppressionState1924) {
	o.SuppressionState = &v
}

// GetValidRevisionReference returns the ValidRevisionReference field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetValidRevisionReference() bool {
	if o == nil || o.ValidRevisionReference == nil {
		var ret bool
		return ret
	}
	return *o.ValidRevisionReference
}

// GetValidRevisionReferenceOk returns a tuple with the ValidRevisionReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetValidRevisionReferenceOk() (*bool, bool) {
	if o == nil || o.ValidRevisionReference == nil {
		return nil, false
	}
	return o.ValidRevisionReference, true
}

// HasValidRevisionReference returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasValidRevisionReference() bool {
	if o != nil && o.ValidRevisionReference != nil {
		return true
	}

	return false
}

// SetValidRevisionReference gets a reference to the given bool and assigns it to the ValidRevisionReference field.
func (o *BTInstanceFolder3627) SetValidRevisionReference(v bool) {
	o.ValidRevisionReference = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *BTInstanceFolder3627) SetVersion(v int32) {
	o.Version = &v
}

// GetFolderEnd returns the FolderEnd field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetFolderEnd() bool {
	if o == nil || o.FolderEnd == nil {
		var ret bool
		return ret
	}
	return *o.FolderEnd
}

// GetFolderEndOk returns a tuple with the FolderEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetFolderEndOk() (*bool, bool) {
	if o == nil || o.FolderEnd == nil {
		return nil, false
	}
	return o.FolderEnd, true
}

// HasFolderEnd returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasFolderEnd() bool {
	if o != nil && o.FolderEnd != nil {
		return true
	}

	return false
}

// SetFolderEnd gets a reference to the given bool and assigns it to the FolderEnd field.
func (o *BTInstanceFolder3627) SetFolderEnd(v bool) {
	o.FolderEnd = &v
}

// GetFolderStart returns the FolderStart field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetFolderStart() bool {
	if o == nil || o.FolderStart == nil {
		var ret bool
		return ret
	}
	return *o.FolderStart
}

// GetFolderStartOk returns a tuple with the FolderStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetFolderStartOk() (*bool, bool) {
	if o == nil || o.FolderStart == nil {
		return nil, false
	}
	return o.FolderStart, true
}

// HasFolderStart returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasFolderStart() bool {
	if o != nil && o.FolderStart != nil {
		return true
	}

	return false
}

// SetFolderStart gets a reference to the given bool and assigns it to the FolderStart field.
func (o *BTInstanceFolder3627) SetFolderStart(v bool) {
	o.FolderStart = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTInstanceFolder3627) SetName(v string) {
	o.Name = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetParameters() []BTMParameter1 {
	if o == nil || o.Parameters == nil {
		var ret []BTMParameter1
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetParametersOk() ([]BTMParameter1, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []BTMParameter1 and assigns it to the Parameters field.
func (o *BTInstanceFolder3627) SetParameters(v []BTMParameter1) {
	o.Parameters = v
}

// GetStartNodeId returns the StartNodeId field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetStartNodeId() string {
	if o == nil || o.StartNodeId == nil {
		var ret string
		return ret
	}
	return *o.StartNodeId
}

// GetStartNodeIdOk returns a tuple with the StartNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetStartNodeIdOk() (*string, bool) {
	if o == nil || o.StartNodeId == nil {
		return nil, false
	}
	return o.StartNodeId, true
}

// HasStartNodeId returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasStartNodeId() bool {
	if o != nil && o.StartNodeId != nil {
		return true
	}

	return false
}

// SetStartNodeId gets a reference to the given string and assigns it to the StartNodeId field.
func (o *BTInstanceFolder3627) SetStartNodeId(v string) {
	o.StartNodeId = &v
}

// GetStartNodeIdRaw returns the StartNodeIdRaw field value if set, zero value otherwise.
func (o *BTInstanceFolder3627) GetStartNodeIdRaw() BTObjectId {
	if o == nil || o.StartNodeIdRaw == nil {
		var ret BTObjectId
		return ret
	}
	return *o.StartNodeIdRaw
}

// GetStartNodeIdRawOk returns a tuple with the StartNodeIdRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceFolder3627) GetStartNodeIdRawOk() (*BTObjectId, bool) {
	if o == nil || o.StartNodeIdRaw == nil {
		return nil, false
	}
	return o.StartNodeIdRaw, true
}

// HasStartNodeIdRaw returns a boolean if a field has been set.
func (o *BTInstanceFolder3627) HasStartNodeIdRaw() bool {
	if o != nil && o.StartNodeIdRaw != nil {
		return true
	}

	return false
}

// SetStartNodeIdRaw gets a reference to the given BTObjectId and assigns it to the StartNodeIdRaw field.
func (o *BTInstanceFolder3627) SetStartNodeIdRaw(v BTObjectId) {
	o.StartNodeIdRaw = &v
}

func (o BTInstanceFolder3627) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.AssemblyInstance != nil {
		toSerialize["assemblyInstance"] = o.AssemblyInstance
	}
	if o.AssemblyPattern != nil {
		toSerialize["assemblyPattern"] = o.AssemblyPattern
	}
	if o.AssemblyReplicate != nil {
		toSerialize["assemblyReplicate"] = o.AssemblyReplicate
	}
	if o.ClonedInstance != nil {
		toSerialize["clonedInstance"] = o.ClonedInstance
	}
	if o.CustomData != nil {
		toSerialize["customData"] = o.CustomData
	}
	if o.InstanceFolder != nil {
		toSerialize["instanceFolder"] = o.InstanceFolder
	}
	if o.InstanceName != nil {
		toSerialize["instanceName"] = o.InstanceName
	}
	if o.IsFlattenedPart != nil {
		toSerialize["isFlattenedPart"] = o.IsFlattenedPart
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	if o.ParametricInstance != nil {
		toSerialize["parametricInstance"] = o.ParametricInstance
	}
	if o.PartInstance != nil {
		toSerialize["partInstance"] = o.PartInstance
	}
	if o.Releasable != nil {
		toSerialize["releasable"] = o.Releasable
	}
	if o.RevisionCustomData != nil {
		toSerialize["revisionCustomData"] = o.RevisionCustomData
	}
	if o.StandardContent != nil {
		toSerialize["standardContent"] = o.StandardContent
	}
	if o.StandardContentParametersId != nil {
		toSerialize["standardContentParametersId"] = o.StandardContentParametersId
	}
	if o.Suppressed != nil {
		toSerialize["suppressed"] = o.Suppressed
	}
	if o.SuppressedFieldIndex != nil {
		toSerialize["suppressedFieldIndex"] = o.SuppressedFieldIndex
	}
	if o.SuppressionConfigured != nil {
		toSerialize["suppressionConfigured"] = o.SuppressionConfigured
	}
	if o.SuppressionState != nil {
		toSerialize["suppressionState"] = o.SuppressionState
	}
	if o.ValidRevisionReference != nil {
		toSerialize["validRevisionReference"] = o.ValidRevisionReference
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.FolderEnd != nil {
		toSerialize["folderEnd"] = o.FolderEnd
	}
	if o.FolderStart != nil {
		toSerialize["folderStart"] = o.FolderStart
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.StartNodeId != nil {
		toSerialize["startNodeId"] = o.StartNodeId
	}
	if o.StartNodeIdRaw != nil {
		toSerialize["startNodeIdRaw"] = o.StartNodeIdRaw
	}
	return json.Marshal(toSerialize)
}

type NullableBTInstanceFolder3627 struct {
	value *BTInstanceFolder3627
	isSet bool
}

func (v NullableBTInstanceFolder3627) Get() *BTInstanceFolder3627 {
	return v.value
}

func (v *NullableBTInstanceFolder3627) Set(val *BTInstanceFolder3627) {
	v.value = val
	v.isSet = true
}

func (v NullableBTInstanceFolder3627) IsSet() bool {
	return v.isSet
}

func (v *NullableBTInstanceFolder3627) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTInstanceFolder3627(val *BTInstanceFolder3627) *NullableBTInstanceFolder3627 {
	return &NullableBTInstanceFolder3627{value: val, isSet: true}
}

func (v NullableBTInstanceFolder3627) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTInstanceFolder3627) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
