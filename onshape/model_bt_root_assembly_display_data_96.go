/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6615-4098eb2c62ac
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTRootAssemblyDisplayData96 struct for BTRootAssemblyDisplayData96
type BTRootAssemblyDisplayData96 struct {
	BtType                                 *string                                       `json:"btType,omitempty"`
	DeletedGeometryMateIds                 []string                                      `json:"deletedGeometryMateIds,omitempty"`
	DeletedLoads                           []string                                      `json:"deletedLoads,omitempty"`
	DeletedMateConnectorIds                []string                                      `json:"deletedMateConnectorIds,omitempty"`
	DeletedMateGroupIds                    []string                                      `json:"deletedMateGroupIds,omitempty"`
	DeletedMateIds                         []string                                      `json:"deletedMateIds,omitempty"`
	DeletedOccurrences                     []BTOccurrence74                              `json:"deletedOccurrences,omitempty"`
	ElementId                              *string                                       `json:"elementId,omitempty"`
	FromFullElementId                      *BTFullElementId756                           `json:"fromFullElementId,omitempty"`
	FullElementId                          *BTFullElementId756                           `json:"fullElementId,omitempty"`
	GeometryMates                          []BTGeometryMateDisplayData1050               `json:"geometryMates,omitempty"`
	Incremental                            *bool                                         `json:"incremental,omitempty"`
	InstanceCount                          *int32                                        `json:"instanceCount,omitempty"`
	IsCollapsible                          *bool                                         `json:"isCollapsible,omitempty"`
	IsForInContext                         *bool                                         `json:"isForInContext,omitempty"`
	KeepFromMicroversion                   *bool                                         `json:"keepFromMicroversion,omitempty"`
	Loads                                  []BTLoadDisplayData837                        `json:"loads,omitempty"`
	MateConnectors                         []BTMateConnectorDisplayData94                `json:"mateConnectors,omitempty"`
	MateGroups                             []BTMateGroupDisplayData1990                  `json:"mateGroups,omitempty"`
	Mates                                  []BTMateDisplayData1358                       `json:"mates,omitempty"`
	MicroversionId                         *BTMicroversionId366                          `json:"microversionId,omitempty"`
	MicroversionIdAndConfigurationInterval *BTMicroversionIdAndConfigurationInterval2364 `json:"microversionIdAndConfigurationInterval,omitempty"`
	MicroversionInterval                   *BTMicroversionIdInterval367                  `json:"microversionInterval,omitempty"`
	Occurrences                            []BTOccurrenceDisplayData95                   `json:"occurrences,omitempty"`
	OriginDisplayData                      *BTOriginDisplayData934                       `json:"originDisplayData,omitempty"`
	PartStudioDisplayData                  []BTPartStudioDisplayDataBase2751             `json:"partStudioDisplayData,omitempty"`
	QuickSummary                           *string                                       `json:"quickSummary,omitempty"`
	VersionForRasterization                *BTElementDisplayData326                      `json:"versionForRasterization,omitempty"`
}

// NewBTRootAssemblyDisplayData96 instantiates a new BTRootAssemblyDisplayData96 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTRootAssemblyDisplayData96() *BTRootAssemblyDisplayData96 {
	this := BTRootAssemblyDisplayData96{}
	return &this
}

// NewBTRootAssemblyDisplayData96WithDefaults instantiates a new BTRootAssemblyDisplayData96 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTRootAssemblyDisplayData96WithDefaults() *BTRootAssemblyDisplayData96 {
	this := BTRootAssemblyDisplayData96{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTRootAssemblyDisplayData96) SetBtType(v string) {
	o.BtType = &v
}

// GetDeletedGeometryMateIds returns the DeletedGeometryMateIds field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetDeletedGeometryMateIds() []string {
	if o == nil || o.DeletedGeometryMateIds == nil {
		var ret []string
		return ret
	}
	return o.DeletedGeometryMateIds
}

// GetDeletedGeometryMateIdsOk returns a tuple with the DeletedGeometryMateIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetDeletedGeometryMateIdsOk() ([]string, bool) {
	if o == nil || o.DeletedGeometryMateIds == nil {
		return nil, false
	}
	return o.DeletedGeometryMateIds, true
}

// HasDeletedGeometryMateIds returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasDeletedGeometryMateIds() bool {
	if o != nil && o.DeletedGeometryMateIds != nil {
		return true
	}

	return false
}

// SetDeletedGeometryMateIds gets a reference to the given []string and assigns it to the DeletedGeometryMateIds field.
func (o *BTRootAssemblyDisplayData96) SetDeletedGeometryMateIds(v []string) {
	o.DeletedGeometryMateIds = v
}

// GetDeletedLoads returns the DeletedLoads field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetDeletedLoads() []string {
	if o == nil || o.DeletedLoads == nil {
		var ret []string
		return ret
	}
	return o.DeletedLoads
}

// GetDeletedLoadsOk returns a tuple with the DeletedLoads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetDeletedLoadsOk() ([]string, bool) {
	if o == nil || o.DeletedLoads == nil {
		return nil, false
	}
	return o.DeletedLoads, true
}

// HasDeletedLoads returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasDeletedLoads() bool {
	if o != nil && o.DeletedLoads != nil {
		return true
	}

	return false
}

// SetDeletedLoads gets a reference to the given []string and assigns it to the DeletedLoads field.
func (o *BTRootAssemblyDisplayData96) SetDeletedLoads(v []string) {
	o.DeletedLoads = v
}

// GetDeletedMateConnectorIds returns the DeletedMateConnectorIds field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetDeletedMateConnectorIds() []string {
	if o == nil || o.DeletedMateConnectorIds == nil {
		var ret []string
		return ret
	}
	return o.DeletedMateConnectorIds
}

// GetDeletedMateConnectorIdsOk returns a tuple with the DeletedMateConnectorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetDeletedMateConnectorIdsOk() ([]string, bool) {
	if o == nil || o.DeletedMateConnectorIds == nil {
		return nil, false
	}
	return o.DeletedMateConnectorIds, true
}

// HasDeletedMateConnectorIds returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasDeletedMateConnectorIds() bool {
	if o != nil && o.DeletedMateConnectorIds != nil {
		return true
	}

	return false
}

// SetDeletedMateConnectorIds gets a reference to the given []string and assigns it to the DeletedMateConnectorIds field.
func (o *BTRootAssemblyDisplayData96) SetDeletedMateConnectorIds(v []string) {
	o.DeletedMateConnectorIds = v
}

// GetDeletedMateGroupIds returns the DeletedMateGroupIds field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetDeletedMateGroupIds() []string {
	if o == nil || o.DeletedMateGroupIds == nil {
		var ret []string
		return ret
	}
	return o.DeletedMateGroupIds
}

// GetDeletedMateGroupIdsOk returns a tuple with the DeletedMateGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetDeletedMateGroupIdsOk() ([]string, bool) {
	if o == nil || o.DeletedMateGroupIds == nil {
		return nil, false
	}
	return o.DeletedMateGroupIds, true
}

// HasDeletedMateGroupIds returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasDeletedMateGroupIds() bool {
	if o != nil && o.DeletedMateGroupIds != nil {
		return true
	}

	return false
}

// SetDeletedMateGroupIds gets a reference to the given []string and assigns it to the DeletedMateGroupIds field.
func (o *BTRootAssemblyDisplayData96) SetDeletedMateGroupIds(v []string) {
	o.DeletedMateGroupIds = v
}

// GetDeletedMateIds returns the DeletedMateIds field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetDeletedMateIds() []string {
	if o == nil || o.DeletedMateIds == nil {
		var ret []string
		return ret
	}
	return o.DeletedMateIds
}

// GetDeletedMateIdsOk returns a tuple with the DeletedMateIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetDeletedMateIdsOk() ([]string, bool) {
	if o == nil || o.DeletedMateIds == nil {
		return nil, false
	}
	return o.DeletedMateIds, true
}

// HasDeletedMateIds returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasDeletedMateIds() bool {
	if o != nil && o.DeletedMateIds != nil {
		return true
	}

	return false
}

// SetDeletedMateIds gets a reference to the given []string and assigns it to the DeletedMateIds field.
func (o *BTRootAssemblyDisplayData96) SetDeletedMateIds(v []string) {
	o.DeletedMateIds = v
}

// GetDeletedOccurrences returns the DeletedOccurrences field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetDeletedOccurrences() []BTOccurrence74 {
	if o == nil || o.DeletedOccurrences == nil {
		var ret []BTOccurrence74
		return ret
	}
	return o.DeletedOccurrences
}

// GetDeletedOccurrencesOk returns a tuple with the DeletedOccurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetDeletedOccurrencesOk() ([]BTOccurrence74, bool) {
	if o == nil || o.DeletedOccurrences == nil {
		return nil, false
	}
	return o.DeletedOccurrences, true
}

// HasDeletedOccurrences returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasDeletedOccurrences() bool {
	if o != nil && o.DeletedOccurrences != nil {
		return true
	}

	return false
}

// SetDeletedOccurrences gets a reference to the given []BTOccurrence74 and assigns it to the DeletedOccurrences field.
func (o *BTRootAssemblyDisplayData96) SetDeletedOccurrences(v []BTOccurrence74) {
	o.DeletedOccurrences = v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTRootAssemblyDisplayData96) SetElementId(v string) {
	o.ElementId = &v
}

// GetFromFullElementId returns the FromFullElementId field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetFromFullElementId() BTFullElementId756 {
	if o == nil || o.FromFullElementId == nil {
		var ret BTFullElementId756
		return ret
	}
	return *o.FromFullElementId
}

// GetFromFullElementIdOk returns a tuple with the FromFullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetFromFullElementIdOk() (*BTFullElementId756, bool) {
	if o == nil || o.FromFullElementId == nil {
		return nil, false
	}
	return o.FromFullElementId, true
}

// HasFromFullElementId returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasFromFullElementId() bool {
	if o != nil && o.FromFullElementId != nil {
		return true
	}

	return false
}

// SetFromFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FromFullElementId field.
func (o *BTRootAssemblyDisplayData96) SetFromFullElementId(v BTFullElementId756) {
	o.FromFullElementId = &v
}

// GetFullElementId returns the FullElementId field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetFullElementId() BTFullElementId756 {
	if o == nil || o.FullElementId == nil {
		var ret BTFullElementId756
		return ret
	}
	return *o.FullElementId
}

// GetFullElementIdOk returns a tuple with the FullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetFullElementIdOk() (*BTFullElementId756, bool) {
	if o == nil || o.FullElementId == nil {
		return nil, false
	}
	return o.FullElementId, true
}

// HasFullElementId returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasFullElementId() bool {
	if o != nil && o.FullElementId != nil {
		return true
	}

	return false
}

// SetFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FullElementId field.
func (o *BTRootAssemblyDisplayData96) SetFullElementId(v BTFullElementId756) {
	o.FullElementId = &v
}

// GetGeometryMates returns the GeometryMates field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetGeometryMates() []BTGeometryMateDisplayData1050 {
	if o == nil || o.GeometryMates == nil {
		var ret []BTGeometryMateDisplayData1050
		return ret
	}
	return o.GeometryMates
}

// GetGeometryMatesOk returns a tuple with the GeometryMates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetGeometryMatesOk() ([]BTGeometryMateDisplayData1050, bool) {
	if o == nil || o.GeometryMates == nil {
		return nil, false
	}
	return o.GeometryMates, true
}

// HasGeometryMates returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasGeometryMates() bool {
	if o != nil && o.GeometryMates != nil {
		return true
	}

	return false
}

// SetGeometryMates gets a reference to the given []BTGeometryMateDisplayData1050 and assigns it to the GeometryMates field.
func (o *BTRootAssemblyDisplayData96) SetGeometryMates(v []BTGeometryMateDisplayData1050) {
	o.GeometryMates = v
}

// GetIncremental returns the Incremental field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetIncremental() bool {
	if o == nil || o.Incremental == nil {
		var ret bool
		return ret
	}
	return *o.Incremental
}

// GetIncrementalOk returns a tuple with the Incremental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetIncrementalOk() (*bool, bool) {
	if o == nil || o.Incremental == nil {
		return nil, false
	}
	return o.Incremental, true
}

// HasIncremental returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasIncremental() bool {
	if o != nil && o.Incremental != nil {
		return true
	}

	return false
}

// SetIncremental gets a reference to the given bool and assigns it to the Incremental field.
func (o *BTRootAssemblyDisplayData96) SetIncremental(v bool) {
	o.Incremental = &v
}

// GetInstanceCount returns the InstanceCount field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetInstanceCount() int32 {
	if o == nil || o.InstanceCount == nil {
		var ret int32
		return ret
	}
	return *o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetInstanceCountOk() (*int32, bool) {
	if o == nil || o.InstanceCount == nil {
		return nil, false
	}
	return o.InstanceCount, true
}

// HasInstanceCount returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasInstanceCount() bool {
	if o != nil && o.InstanceCount != nil {
		return true
	}

	return false
}

// SetInstanceCount gets a reference to the given int32 and assigns it to the InstanceCount field.
func (o *BTRootAssemblyDisplayData96) SetInstanceCount(v int32) {
	o.InstanceCount = &v
}

// GetIsCollapsible returns the IsCollapsible field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetIsCollapsible() bool {
	if o == nil || o.IsCollapsible == nil {
		var ret bool
		return ret
	}
	return *o.IsCollapsible
}

// GetIsCollapsibleOk returns a tuple with the IsCollapsible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetIsCollapsibleOk() (*bool, bool) {
	if o == nil || o.IsCollapsible == nil {
		return nil, false
	}
	return o.IsCollapsible, true
}

// HasIsCollapsible returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasIsCollapsible() bool {
	if o != nil && o.IsCollapsible != nil {
		return true
	}

	return false
}

// SetIsCollapsible gets a reference to the given bool and assigns it to the IsCollapsible field.
func (o *BTRootAssemblyDisplayData96) SetIsCollapsible(v bool) {
	o.IsCollapsible = &v
}

// GetIsForInContext returns the IsForInContext field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetIsForInContext() bool {
	if o == nil || o.IsForInContext == nil {
		var ret bool
		return ret
	}
	return *o.IsForInContext
}

// GetIsForInContextOk returns a tuple with the IsForInContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetIsForInContextOk() (*bool, bool) {
	if o == nil || o.IsForInContext == nil {
		return nil, false
	}
	return o.IsForInContext, true
}

// HasIsForInContext returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasIsForInContext() bool {
	if o != nil && o.IsForInContext != nil {
		return true
	}

	return false
}

// SetIsForInContext gets a reference to the given bool and assigns it to the IsForInContext field.
func (o *BTRootAssemblyDisplayData96) SetIsForInContext(v bool) {
	o.IsForInContext = &v
}

// GetKeepFromMicroversion returns the KeepFromMicroversion field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetKeepFromMicroversion() bool {
	if o == nil || o.KeepFromMicroversion == nil {
		var ret bool
		return ret
	}
	return *o.KeepFromMicroversion
}

// GetKeepFromMicroversionOk returns a tuple with the KeepFromMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetKeepFromMicroversionOk() (*bool, bool) {
	if o == nil || o.KeepFromMicroversion == nil {
		return nil, false
	}
	return o.KeepFromMicroversion, true
}

// HasKeepFromMicroversion returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasKeepFromMicroversion() bool {
	if o != nil && o.KeepFromMicroversion != nil {
		return true
	}

	return false
}

// SetKeepFromMicroversion gets a reference to the given bool and assigns it to the KeepFromMicroversion field.
func (o *BTRootAssemblyDisplayData96) SetKeepFromMicroversion(v bool) {
	o.KeepFromMicroversion = &v
}

// GetLoads returns the Loads field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetLoads() []BTLoadDisplayData837 {
	if o == nil || o.Loads == nil {
		var ret []BTLoadDisplayData837
		return ret
	}
	return o.Loads
}

// GetLoadsOk returns a tuple with the Loads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetLoadsOk() ([]BTLoadDisplayData837, bool) {
	if o == nil || o.Loads == nil {
		return nil, false
	}
	return o.Loads, true
}

// HasLoads returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasLoads() bool {
	if o != nil && o.Loads != nil {
		return true
	}

	return false
}

// SetLoads gets a reference to the given []BTLoadDisplayData837 and assigns it to the Loads field.
func (o *BTRootAssemblyDisplayData96) SetLoads(v []BTLoadDisplayData837) {
	o.Loads = v
}

// GetMateConnectors returns the MateConnectors field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetMateConnectors() []BTMateConnectorDisplayData94 {
	if o == nil || o.MateConnectors == nil {
		var ret []BTMateConnectorDisplayData94
		return ret
	}
	return o.MateConnectors
}

// GetMateConnectorsOk returns a tuple with the MateConnectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetMateConnectorsOk() ([]BTMateConnectorDisplayData94, bool) {
	if o == nil || o.MateConnectors == nil {
		return nil, false
	}
	return o.MateConnectors, true
}

// HasMateConnectors returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasMateConnectors() bool {
	if o != nil && o.MateConnectors != nil {
		return true
	}

	return false
}

// SetMateConnectors gets a reference to the given []BTMateConnectorDisplayData94 and assigns it to the MateConnectors field.
func (o *BTRootAssemblyDisplayData96) SetMateConnectors(v []BTMateConnectorDisplayData94) {
	o.MateConnectors = v
}

// GetMateGroups returns the MateGroups field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetMateGroups() []BTMateGroupDisplayData1990 {
	if o == nil || o.MateGroups == nil {
		var ret []BTMateGroupDisplayData1990
		return ret
	}
	return o.MateGroups
}

// GetMateGroupsOk returns a tuple with the MateGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetMateGroupsOk() ([]BTMateGroupDisplayData1990, bool) {
	if o == nil || o.MateGroups == nil {
		return nil, false
	}
	return o.MateGroups, true
}

// HasMateGroups returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasMateGroups() bool {
	if o != nil && o.MateGroups != nil {
		return true
	}

	return false
}

// SetMateGroups gets a reference to the given []BTMateGroupDisplayData1990 and assigns it to the MateGroups field.
func (o *BTRootAssemblyDisplayData96) SetMateGroups(v []BTMateGroupDisplayData1990) {
	o.MateGroups = v
}

// GetMates returns the Mates field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetMates() []BTMateDisplayData1358 {
	if o == nil || o.Mates == nil {
		var ret []BTMateDisplayData1358
		return ret
	}
	return o.Mates
}

// GetMatesOk returns a tuple with the Mates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetMatesOk() ([]BTMateDisplayData1358, bool) {
	if o == nil || o.Mates == nil {
		return nil, false
	}
	return o.Mates, true
}

// HasMates returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasMates() bool {
	if o != nil && o.Mates != nil {
		return true
	}

	return false
}

// SetMates gets a reference to the given []BTMateDisplayData1358 and assigns it to the Mates field.
func (o *BTRootAssemblyDisplayData96) SetMates(v []BTMateDisplayData1358) {
	o.Mates = v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetMicroversionId() BTMicroversionId366 {
	if o == nil || o.MicroversionId == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetMicroversionIdOk() (*BTMicroversionId366, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given BTMicroversionId366 and assigns it to the MicroversionId field.
func (o *BTRootAssemblyDisplayData96) SetMicroversionId(v BTMicroversionId366) {
	o.MicroversionId = &v
}

// GetMicroversionIdAndConfigurationInterval returns the MicroversionIdAndConfigurationInterval field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetMicroversionIdAndConfigurationInterval() BTMicroversionIdAndConfigurationInterval2364 {
	if o == nil || o.MicroversionIdAndConfigurationInterval == nil {
		var ret BTMicroversionIdAndConfigurationInterval2364
		return ret
	}
	return *o.MicroversionIdAndConfigurationInterval
}

// GetMicroversionIdAndConfigurationIntervalOk returns a tuple with the MicroversionIdAndConfigurationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetMicroversionIdAndConfigurationIntervalOk() (*BTMicroversionIdAndConfigurationInterval2364, bool) {
	if o == nil || o.MicroversionIdAndConfigurationInterval == nil {
		return nil, false
	}
	return o.MicroversionIdAndConfigurationInterval, true
}

// HasMicroversionIdAndConfigurationInterval returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasMicroversionIdAndConfigurationInterval() bool {
	if o != nil && o.MicroversionIdAndConfigurationInterval != nil {
		return true
	}

	return false
}

// SetMicroversionIdAndConfigurationInterval gets a reference to the given BTMicroversionIdAndConfigurationInterval2364 and assigns it to the MicroversionIdAndConfigurationInterval field.
func (o *BTRootAssemblyDisplayData96) SetMicroversionIdAndConfigurationInterval(v BTMicroversionIdAndConfigurationInterval2364) {
	o.MicroversionIdAndConfigurationInterval = &v
}

// GetMicroversionInterval returns the MicroversionInterval field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetMicroversionInterval() BTMicroversionIdInterval367 {
	if o == nil || o.MicroversionInterval == nil {
		var ret BTMicroversionIdInterval367
		return ret
	}
	return *o.MicroversionInterval
}

// GetMicroversionIntervalOk returns a tuple with the MicroversionInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetMicroversionIntervalOk() (*BTMicroversionIdInterval367, bool) {
	if o == nil || o.MicroversionInterval == nil {
		return nil, false
	}
	return o.MicroversionInterval, true
}

// HasMicroversionInterval returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasMicroversionInterval() bool {
	if o != nil && o.MicroversionInterval != nil {
		return true
	}

	return false
}

// SetMicroversionInterval gets a reference to the given BTMicroversionIdInterval367 and assigns it to the MicroversionInterval field.
func (o *BTRootAssemblyDisplayData96) SetMicroversionInterval(v BTMicroversionIdInterval367) {
	o.MicroversionInterval = &v
}

// GetOccurrences returns the Occurrences field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetOccurrences() []BTOccurrenceDisplayData95 {
	if o == nil || o.Occurrences == nil {
		var ret []BTOccurrenceDisplayData95
		return ret
	}
	return o.Occurrences
}

// GetOccurrencesOk returns a tuple with the Occurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetOccurrencesOk() ([]BTOccurrenceDisplayData95, bool) {
	if o == nil || o.Occurrences == nil {
		return nil, false
	}
	return o.Occurrences, true
}

// HasOccurrences returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasOccurrences() bool {
	if o != nil && o.Occurrences != nil {
		return true
	}

	return false
}

// SetOccurrences gets a reference to the given []BTOccurrenceDisplayData95 and assigns it to the Occurrences field.
func (o *BTRootAssemblyDisplayData96) SetOccurrences(v []BTOccurrenceDisplayData95) {
	o.Occurrences = v
}

// GetOriginDisplayData returns the OriginDisplayData field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetOriginDisplayData() BTOriginDisplayData934 {
	if o == nil || o.OriginDisplayData == nil {
		var ret BTOriginDisplayData934
		return ret
	}
	return *o.OriginDisplayData
}

// GetOriginDisplayDataOk returns a tuple with the OriginDisplayData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetOriginDisplayDataOk() (*BTOriginDisplayData934, bool) {
	if o == nil || o.OriginDisplayData == nil {
		return nil, false
	}
	return o.OriginDisplayData, true
}

// HasOriginDisplayData returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasOriginDisplayData() bool {
	if o != nil && o.OriginDisplayData != nil {
		return true
	}

	return false
}

// SetOriginDisplayData gets a reference to the given BTOriginDisplayData934 and assigns it to the OriginDisplayData field.
func (o *BTRootAssemblyDisplayData96) SetOriginDisplayData(v BTOriginDisplayData934) {
	o.OriginDisplayData = &v
}

// GetPartStudioDisplayData returns the PartStudioDisplayData field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetPartStudioDisplayData() []BTPartStudioDisplayDataBase2751 {
	if o == nil || o.PartStudioDisplayData == nil {
		var ret []BTPartStudioDisplayDataBase2751
		return ret
	}
	return o.PartStudioDisplayData
}

// GetPartStudioDisplayDataOk returns a tuple with the PartStudioDisplayData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetPartStudioDisplayDataOk() ([]BTPartStudioDisplayDataBase2751, bool) {
	if o == nil || o.PartStudioDisplayData == nil {
		return nil, false
	}
	return o.PartStudioDisplayData, true
}

// HasPartStudioDisplayData returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasPartStudioDisplayData() bool {
	if o != nil && o.PartStudioDisplayData != nil {
		return true
	}

	return false
}

// SetPartStudioDisplayData gets a reference to the given []BTPartStudioDisplayDataBase2751 and assigns it to the PartStudioDisplayData field.
func (o *BTRootAssemblyDisplayData96) SetPartStudioDisplayData(v []BTPartStudioDisplayDataBase2751) {
	o.PartStudioDisplayData = v
}

// GetQuickSummary returns the QuickSummary field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetQuickSummary() string {
	if o == nil || o.QuickSummary == nil {
		var ret string
		return ret
	}
	return *o.QuickSummary
}

// GetQuickSummaryOk returns a tuple with the QuickSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetQuickSummaryOk() (*string, bool) {
	if o == nil || o.QuickSummary == nil {
		return nil, false
	}
	return o.QuickSummary, true
}

// HasQuickSummary returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasQuickSummary() bool {
	if o != nil && o.QuickSummary != nil {
		return true
	}

	return false
}

// SetQuickSummary gets a reference to the given string and assigns it to the QuickSummary field.
func (o *BTRootAssemblyDisplayData96) SetQuickSummary(v string) {
	o.QuickSummary = &v
}

// GetVersionForRasterization returns the VersionForRasterization field value if set, zero value otherwise.
func (o *BTRootAssemblyDisplayData96) GetVersionForRasterization() BTElementDisplayData326 {
	if o == nil || o.VersionForRasterization == nil {
		var ret BTElementDisplayData326
		return ret
	}
	return *o.VersionForRasterization
}

// GetVersionForRasterizationOk returns a tuple with the VersionForRasterization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRootAssemblyDisplayData96) GetVersionForRasterizationOk() (*BTElementDisplayData326, bool) {
	if o == nil || o.VersionForRasterization == nil {
		return nil, false
	}
	return o.VersionForRasterization, true
}

// HasVersionForRasterization returns a boolean if a field has been set.
func (o *BTRootAssemblyDisplayData96) HasVersionForRasterization() bool {
	if o != nil && o.VersionForRasterization != nil {
		return true
	}

	return false
}

// SetVersionForRasterization gets a reference to the given BTElementDisplayData326 and assigns it to the VersionForRasterization field.
func (o *BTRootAssemblyDisplayData96) SetVersionForRasterization(v BTElementDisplayData326) {
	o.VersionForRasterization = &v
}

func (o BTRootAssemblyDisplayData96) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DeletedGeometryMateIds != nil {
		toSerialize["deletedGeometryMateIds"] = o.DeletedGeometryMateIds
	}
	if o.DeletedLoads != nil {
		toSerialize["deletedLoads"] = o.DeletedLoads
	}
	if o.DeletedMateConnectorIds != nil {
		toSerialize["deletedMateConnectorIds"] = o.DeletedMateConnectorIds
	}
	if o.DeletedMateGroupIds != nil {
		toSerialize["deletedMateGroupIds"] = o.DeletedMateGroupIds
	}
	if o.DeletedMateIds != nil {
		toSerialize["deletedMateIds"] = o.DeletedMateIds
	}
	if o.DeletedOccurrences != nil {
		toSerialize["deletedOccurrences"] = o.DeletedOccurrences
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.FromFullElementId != nil {
		toSerialize["fromFullElementId"] = o.FromFullElementId
	}
	if o.FullElementId != nil {
		toSerialize["fullElementId"] = o.FullElementId
	}
	if o.GeometryMates != nil {
		toSerialize["geometryMates"] = o.GeometryMates
	}
	if o.Incremental != nil {
		toSerialize["incremental"] = o.Incremental
	}
	if o.InstanceCount != nil {
		toSerialize["instanceCount"] = o.InstanceCount
	}
	if o.IsCollapsible != nil {
		toSerialize["isCollapsible"] = o.IsCollapsible
	}
	if o.IsForInContext != nil {
		toSerialize["isForInContext"] = o.IsForInContext
	}
	if o.KeepFromMicroversion != nil {
		toSerialize["keepFromMicroversion"] = o.KeepFromMicroversion
	}
	if o.Loads != nil {
		toSerialize["loads"] = o.Loads
	}
	if o.MateConnectors != nil {
		toSerialize["mateConnectors"] = o.MateConnectors
	}
	if o.MateGroups != nil {
		toSerialize["mateGroups"] = o.MateGroups
	}
	if o.Mates != nil {
		toSerialize["mates"] = o.Mates
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	if o.MicroversionIdAndConfigurationInterval != nil {
		toSerialize["microversionIdAndConfigurationInterval"] = o.MicroversionIdAndConfigurationInterval
	}
	if o.MicroversionInterval != nil {
		toSerialize["microversionInterval"] = o.MicroversionInterval
	}
	if o.Occurrences != nil {
		toSerialize["occurrences"] = o.Occurrences
	}
	if o.OriginDisplayData != nil {
		toSerialize["originDisplayData"] = o.OriginDisplayData
	}
	if o.PartStudioDisplayData != nil {
		toSerialize["partStudioDisplayData"] = o.PartStudioDisplayData
	}
	if o.QuickSummary != nil {
		toSerialize["quickSummary"] = o.QuickSummary
	}
	if o.VersionForRasterization != nil {
		toSerialize["versionForRasterization"] = o.VersionForRasterization
	}
	return json.Marshal(toSerialize)
}

type NullableBTRootAssemblyDisplayData96 struct {
	value *BTRootAssemblyDisplayData96
	isSet bool
}

func (v NullableBTRootAssemblyDisplayData96) Get() *BTRootAssemblyDisplayData96 {
	return v.value
}

func (v *NullableBTRootAssemblyDisplayData96) Set(val *BTRootAssemblyDisplayData96) {
	v.value = val
	v.isSet = true
}

func (v NullableBTRootAssemblyDisplayData96) IsSet() bool {
	return v.isSet
}

func (v *NullableBTRootAssemblyDisplayData96) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTRootAssemblyDisplayData96(val *BTRootAssemblyDisplayData96) *NullableBTRootAssemblyDisplayData96 {
	return &NullableBTRootAssemblyDisplayData96{value: val, isSet: true}
}

func (v NullableBTRootAssemblyDisplayData96) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTRootAssemblyDisplayData96) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
