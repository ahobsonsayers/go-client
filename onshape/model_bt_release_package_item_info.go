/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5812-3c52042ec720
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTReleasePackageItemInfo struct for BTReleasePackageItemInfo
type BTReleasePackageItemInfo struct {
	CanExport                         *bool                    `json:"canExport,omitempty"`
	ChangeDetectionStatus             *int32                   `json:"changeDetectionStatus,omitempty"`
	CompanyId                         *string                  `json:"companyId,omitempty"`
	Configuration                     *string                  `json:"configuration,omitempty"`
	ConfigurationKey                  *string                  `json:"configurationKey,omitempty"`
	DiffThumbnailConfigurationKey     *string                  `json:"diffThumbnailConfigurationKey,omitempty"`
	DocumentId                        *string                  `json:"documentId,omitempty"`
	ElementId                         *string                  `json:"elementId,omitempty"`
	ElementType                       *int32                   `json:"elementType,omitempty"`
	Errors                            []BTReleaseItemErrorInfo `json:"errors,omitempty"`
	Href                              *string                  `json:"href,omitempty"`
	Id                                *string                  `json:"id,omitempty"`
	IsRevisionManaged                 *bool                    `json:"isRevisionManaged,omitempty"`
	IsRootItem                        *bool                    `json:"isRootItem,omitempty"`
	IsTranslatable                    *bool                    `json:"isTranslatable,omitempty"`
	MeshState                         *int32                   `json:"meshState,omitempty"`
	MimeType                          *string                  `json:"mimeType,omitempty"`
	Name                              *string                  `json:"name,omitempty"`
	ObsoletionRevisionId              *string                  `json:"obsoletionRevisionId,omitempty"`
	OriginalWorkspaceId               *string                  `json:"originalWorkspaceId,omitempty"`
	PartId                            *string                  `json:"partId,omitempty"`
	PartType                          *string                  `json:"partType,omitempty"`
	Properties                        []BTMetadataPropertyInfo `json:"properties,omitempty"`
	ReferenceIds                      []string                 `json:"referenceIds,omitempty"`
	ReferenceIdsFromOriginalWorkspace []string                 `json:"referenceIdsFromOriginalWorkspace,omitempty"`
	Rpid                              *string                  `json:"rpid,omitempty"`
	SmallThumbnailHref                *string                  `json:"smallThumbnailHref,omitempty"`
	VersionId                         *string                  `json:"versionId,omitempty"`
	ViewRef                           *string                  `json:"viewRef,omitempty"`
	WorkspaceId                       *string                  `json:"workspaceId,omitempty"`
}

// NewBTReleasePackageItemInfo instantiates a new BTReleasePackageItemInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTReleasePackageItemInfo() *BTReleasePackageItemInfo {
	this := BTReleasePackageItemInfo{}
	return &this
}

// NewBTReleasePackageItemInfoWithDefaults instantiates a new BTReleasePackageItemInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTReleasePackageItemInfoWithDefaults() *BTReleasePackageItemInfo {
	this := BTReleasePackageItemInfo{}
	return &this
}

// GetCanExport returns the CanExport field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetCanExport() bool {
	if o == nil || o.CanExport == nil {
		var ret bool
		return ret
	}
	return *o.CanExport
}

// GetCanExportOk returns a tuple with the CanExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetCanExportOk() (*bool, bool) {
	if o == nil || o.CanExport == nil {
		return nil, false
	}
	return o.CanExport, true
}

// HasCanExport returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasCanExport() bool {
	if o != nil && o.CanExport != nil {
		return true
	}

	return false
}

// SetCanExport gets a reference to the given bool and assigns it to the CanExport field.
func (o *BTReleasePackageItemInfo) SetCanExport(v bool) {
	o.CanExport = &v
}

// GetChangeDetectionStatus returns the ChangeDetectionStatus field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetChangeDetectionStatus() int32 {
	if o == nil || o.ChangeDetectionStatus == nil {
		var ret int32
		return ret
	}
	return *o.ChangeDetectionStatus
}

// GetChangeDetectionStatusOk returns a tuple with the ChangeDetectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetChangeDetectionStatusOk() (*int32, bool) {
	if o == nil || o.ChangeDetectionStatus == nil {
		return nil, false
	}
	return o.ChangeDetectionStatus, true
}

// HasChangeDetectionStatus returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasChangeDetectionStatus() bool {
	if o != nil && o.ChangeDetectionStatus != nil {
		return true
	}

	return false
}

// SetChangeDetectionStatus gets a reference to the given int32 and assigns it to the ChangeDetectionStatus field.
func (o *BTReleasePackageItemInfo) SetChangeDetectionStatus(v int32) {
	o.ChangeDetectionStatus = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTReleasePackageItemInfo) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetConfiguration() string {
	if o == nil || o.Configuration == nil {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetConfigurationOk() (*string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *BTReleasePackageItemInfo) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetConfigurationKey returns the ConfigurationKey field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetConfigurationKey() string {
	if o == nil || o.ConfigurationKey == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationKey
}

// GetConfigurationKeyOk returns a tuple with the ConfigurationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetConfigurationKeyOk() (*string, bool) {
	if o == nil || o.ConfigurationKey == nil {
		return nil, false
	}
	return o.ConfigurationKey, true
}

// HasConfigurationKey returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasConfigurationKey() bool {
	if o != nil && o.ConfigurationKey != nil {
		return true
	}

	return false
}

// SetConfigurationKey gets a reference to the given string and assigns it to the ConfigurationKey field.
func (o *BTReleasePackageItemInfo) SetConfigurationKey(v string) {
	o.ConfigurationKey = &v
}

// GetDiffThumbnailConfigurationKey returns the DiffThumbnailConfigurationKey field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetDiffThumbnailConfigurationKey() string {
	if o == nil || o.DiffThumbnailConfigurationKey == nil {
		var ret string
		return ret
	}
	return *o.DiffThumbnailConfigurationKey
}

// GetDiffThumbnailConfigurationKeyOk returns a tuple with the DiffThumbnailConfigurationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetDiffThumbnailConfigurationKeyOk() (*string, bool) {
	if o == nil || o.DiffThumbnailConfigurationKey == nil {
		return nil, false
	}
	return o.DiffThumbnailConfigurationKey, true
}

// HasDiffThumbnailConfigurationKey returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasDiffThumbnailConfigurationKey() bool {
	if o != nil && o.DiffThumbnailConfigurationKey != nil {
		return true
	}

	return false
}

// SetDiffThumbnailConfigurationKey gets a reference to the given string and assigns it to the DiffThumbnailConfigurationKey field.
func (o *BTReleasePackageItemInfo) SetDiffThumbnailConfigurationKey(v string) {
	o.DiffThumbnailConfigurationKey = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTReleasePackageItemInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTReleasePackageItemInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetElementType returns the ElementType field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetElementType() int32 {
	if o == nil || o.ElementType == nil {
		var ret int32
		return ret
	}
	return *o.ElementType
}

// GetElementTypeOk returns a tuple with the ElementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetElementTypeOk() (*int32, bool) {
	if o == nil || o.ElementType == nil {
		return nil, false
	}
	return o.ElementType, true
}

// HasElementType returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasElementType() bool {
	if o != nil && o.ElementType != nil {
		return true
	}

	return false
}

// SetElementType gets a reference to the given int32 and assigns it to the ElementType field.
func (o *BTReleasePackageItemInfo) SetElementType(v int32) {
	o.ElementType = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetErrors() []BTReleaseItemErrorInfo {
	if o == nil || o.Errors == nil {
		var ret []BTReleaseItemErrorInfo
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetErrorsOk() ([]BTReleaseItemErrorInfo, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []BTReleaseItemErrorInfo and assigns it to the Errors field.
func (o *BTReleasePackageItemInfo) SetErrors(v []BTReleaseItemErrorInfo) {
	o.Errors = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTReleasePackageItemInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTReleasePackageItemInfo) SetId(v string) {
	o.Id = &v
}

// GetIsRevisionManaged returns the IsRevisionManaged field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetIsRevisionManaged() bool {
	if o == nil || o.IsRevisionManaged == nil {
		var ret bool
		return ret
	}
	return *o.IsRevisionManaged
}

// GetIsRevisionManagedOk returns a tuple with the IsRevisionManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetIsRevisionManagedOk() (*bool, bool) {
	if o == nil || o.IsRevisionManaged == nil {
		return nil, false
	}
	return o.IsRevisionManaged, true
}

// HasIsRevisionManaged returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasIsRevisionManaged() bool {
	if o != nil && o.IsRevisionManaged != nil {
		return true
	}

	return false
}

// SetIsRevisionManaged gets a reference to the given bool and assigns it to the IsRevisionManaged field.
func (o *BTReleasePackageItemInfo) SetIsRevisionManaged(v bool) {
	o.IsRevisionManaged = &v
}

// GetIsRootItem returns the IsRootItem field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetIsRootItem() bool {
	if o == nil || o.IsRootItem == nil {
		var ret bool
		return ret
	}
	return *o.IsRootItem
}

// GetIsRootItemOk returns a tuple with the IsRootItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetIsRootItemOk() (*bool, bool) {
	if o == nil || o.IsRootItem == nil {
		return nil, false
	}
	return o.IsRootItem, true
}

// HasIsRootItem returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasIsRootItem() bool {
	if o != nil && o.IsRootItem != nil {
		return true
	}

	return false
}

// SetIsRootItem gets a reference to the given bool and assigns it to the IsRootItem field.
func (o *BTReleasePackageItemInfo) SetIsRootItem(v bool) {
	o.IsRootItem = &v
}

// GetIsTranslatable returns the IsTranslatable field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetIsTranslatable() bool {
	if o == nil || o.IsTranslatable == nil {
		var ret bool
		return ret
	}
	return *o.IsTranslatable
}

// GetIsTranslatableOk returns a tuple with the IsTranslatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetIsTranslatableOk() (*bool, bool) {
	if o == nil || o.IsTranslatable == nil {
		return nil, false
	}
	return o.IsTranslatable, true
}

// HasIsTranslatable returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasIsTranslatable() bool {
	if o != nil && o.IsTranslatable != nil {
		return true
	}

	return false
}

// SetIsTranslatable gets a reference to the given bool and assigns it to the IsTranslatable field.
func (o *BTReleasePackageItemInfo) SetIsTranslatable(v bool) {
	o.IsTranslatable = &v
}

// GetMeshState returns the MeshState field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetMeshState() int32 {
	if o == nil || o.MeshState == nil {
		var ret int32
		return ret
	}
	return *o.MeshState
}

// GetMeshStateOk returns a tuple with the MeshState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetMeshStateOk() (*int32, bool) {
	if o == nil || o.MeshState == nil {
		return nil, false
	}
	return o.MeshState, true
}

// HasMeshState returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasMeshState() bool {
	if o != nil && o.MeshState != nil {
		return true
	}

	return false
}

// SetMeshState gets a reference to the given int32 and assigns it to the MeshState field.
func (o *BTReleasePackageItemInfo) SetMeshState(v int32) {
	o.MeshState = &v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetMimeType() string {
	if o == nil || o.MimeType == nil {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetMimeTypeOk() (*string, bool) {
	if o == nil || o.MimeType == nil {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasMimeType() bool {
	if o != nil && o.MimeType != nil {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *BTReleasePackageItemInfo) SetMimeType(v string) {
	o.MimeType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTReleasePackageItemInfo) SetName(v string) {
	o.Name = &v
}

// GetObsoletionRevisionId returns the ObsoletionRevisionId field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetObsoletionRevisionId() string {
	if o == nil || o.ObsoletionRevisionId == nil {
		var ret string
		return ret
	}
	return *o.ObsoletionRevisionId
}

// GetObsoletionRevisionIdOk returns a tuple with the ObsoletionRevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetObsoletionRevisionIdOk() (*string, bool) {
	if o == nil || o.ObsoletionRevisionId == nil {
		return nil, false
	}
	return o.ObsoletionRevisionId, true
}

// HasObsoletionRevisionId returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasObsoletionRevisionId() bool {
	if o != nil && o.ObsoletionRevisionId != nil {
		return true
	}

	return false
}

// SetObsoletionRevisionId gets a reference to the given string and assigns it to the ObsoletionRevisionId field.
func (o *BTReleasePackageItemInfo) SetObsoletionRevisionId(v string) {
	o.ObsoletionRevisionId = &v
}

// GetOriginalWorkspaceId returns the OriginalWorkspaceId field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetOriginalWorkspaceId() string {
	if o == nil || o.OriginalWorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.OriginalWorkspaceId
}

// GetOriginalWorkspaceIdOk returns a tuple with the OriginalWorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetOriginalWorkspaceIdOk() (*string, bool) {
	if o == nil || o.OriginalWorkspaceId == nil {
		return nil, false
	}
	return o.OriginalWorkspaceId, true
}

// HasOriginalWorkspaceId returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasOriginalWorkspaceId() bool {
	if o != nil && o.OriginalWorkspaceId != nil {
		return true
	}

	return false
}

// SetOriginalWorkspaceId gets a reference to the given string and assigns it to the OriginalWorkspaceId field.
func (o *BTReleasePackageItemInfo) SetOriginalWorkspaceId(v string) {
	o.OriginalWorkspaceId = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTReleasePackageItemInfo) SetPartId(v string) {
	o.PartId = &v
}

// GetPartType returns the PartType field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetPartType() string {
	if o == nil || o.PartType == nil {
		var ret string
		return ret
	}
	return *o.PartType
}

// GetPartTypeOk returns a tuple with the PartType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetPartTypeOk() (*string, bool) {
	if o == nil || o.PartType == nil {
		return nil, false
	}
	return o.PartType, true
}

// HasPartType returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasPartType() bool {
	if o != nil && o.PartType != nil {
		return true
	}

	return false
}

// SetPartType gets a reference to the given string and assigns it to the PartType field.
func (o *BTReleasePackageItemInfo) SetPartType(v string) {
	o.PartType = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetProperties() []BTMetadataPropertyInfo {
	if o == nil || o.Properties == nil {
		var ret []BTMetadataPropertyInfo
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetPropertiesOk() ([]BTMetadataPropertyInfo, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []BTMetadataPropertyInfo and assigns it to the Properties field.
func (o *BTReleasePackageItemInfo) SetProperties(v []BTMetadataPropertyInfo) {
	o.Properties = v
}

// GetReferenceIds returns the ReferenceIds field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetReferenceIds() []string {
	if o == nil || o.ReferenceIds == nil {
		var ret []string
		return ret
	}
	return o.ReferenceIds
}

// GetReferenceIdsOk returns a tuple with the ReferenceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetReferenceIdsOk() ([]string, bool) {
	if o == nil || o.ReferenceIds == nil {
		return nil, false
	}
	return o.ReferenceIds, true
}

// HasReferenceIds returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasReferenceIds() bool {
	if o != nil && o.ReferenceIds != nil {
		return true
	}

	return false
}

// SetReferenceIds gets a reference to the given []string and assigns it to the ReferenceIds field.
func (o *BTReleasePackageItemInfo) SetReferenceIds(v []string) {
	o.ReferenceIds = v
}

// GetReferenceIdsFromOriginalWorkspace returns the ReferenceIdsFromOriginalWorkspace field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetReferenceIdsFromOriginalWorkspace() []string {
	if o == nil || o.ReferenceIdsFromOriginalWorkspace == nil {
		var ret []string
		return ret
	}
	return o.ReferenceIdsFromOriginalWorkspace
}

// GetReferenceIdsFromOriginalWorkspaceOk returns a tuple with the ReferenceIdsFromOriginalWorkspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetReferenceIdsFromOriginalWorkspaceOk() ([]string, bool) {
	if o == nil || o.ReferenceIdsFromOriginalWorkspace == nil {
		return nil, false
	}
	return o.ReferenceIdsFromOriginalWorkspace, true
}

// HasReferenceIdsFromOriginalWorkspace returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasReferenceIdsFromOriginalWorkspace() bool {
	if o != nil && o.ReferenceIdsFromOriginalWorkspace != nil {
		return true
	}

	return false
}

// SetReferenceIdsFromOriginalWorkspace gets a reference to the given []string and assigns it to the ReferenceIdsFromOriginalWorkspace field.
func (o *BTReleasePackageItemInfo) SetReferenceIdsFromOriginalWorkspace(v []string) {
	o.ReferenceIdsFromOriginalWorkspace = v
}

// GetRpid returns the Rpid field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetRpid() string {
	if o == nil || o.Rpid == nil {
		var ret string
		return ret
	}
	return *o.Rpid
}

// GetRpidOk returns a tuple with the Rpid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetRpidOk() (*string, bool) {
	if o == nil || o.Rpid == nil {
		return nil, false
	}
	return o.Rpid, true
}

// HasRpid returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasRpid() bool {
	if o != nil && o.Rpid != nil {
		return true
	}

	return false
}

// SetRpid gets a reference to the given string and assigns it to the Rpid field.
func (o *BTReleasePackageItemInfo) SetRpid(v string) {
	o.Rpid = &v
}

// GetSmallThumbnailHref returns the SmallThumbnailHref field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetSmallThumbnailHref() string {
	if o == nil || o.SmallThumbnailHref == nil {
		var ret string
		return ret
	}
	return *o.SmallThumbnailHref
}

// GetSmallThumbnailHrefOk returns a tuple with the SmallThumbnailHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetSmallThumbnailHrefOk() (*string, bool) {
	if o == nil || o.SmallThumbnailHref == nil {
		return nil, false
	}
	return o.SmallThumbnailHref, true
}

// HasSmallThumbnailHref returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasSmallThumbnailHref() bool {
	if o != nil && o.SmallThumbnailHref != nil {
		return true
	}

	return false
}

// SetSmallThumbnailHref gets a reference to the given string and assigns it to the SmallThumbnailHref field.
func (o *BTReleasePackageItemInfo) SetSmallThumbnailHref(v string) {
	o.SmallThumbnailHref = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTReleasePackageItemInfo) SetVersionId(v string) {
	o.VersionId = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTReleasePackageItemInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *BTReleasePackageItemInfo) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageItemInfo) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *BTReleasePackageItemInfo) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *BTReleasePackageItemInfo) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o BTReleasePackageItemInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanExport != nil {
		toSerialize["canExport"] = o.CanExport
	}
	if o.ChangeDetectionStatus != nil {
		toSerialize["changeDetectionStatus"] = o.ChangeDetectionStatus
	}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.ConfigurationKey != nil {
		toSerialize["configurationKey"] = o.ConfigurationKey
	}
	if o.DiffThumbnailConfigurationKey != nil {
		toSerialize["diffThumbnailConfigurationKey"] = o.DiffThumbnailConfigurationKey
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ElementType != nil {
		toSerialize["elementType"] = o.ElementType
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsRevisionManaged != nil {
		toSerialize["isRevisionManaged"] = o.IsRevisionManaged
	}
	if o.IsRootItem != nil {
		toSerialize["isRootItem"] = o.IsRootItem
	}
	if o.IsTranslatable != nil {
		toSerialize["isTranslatable"] = o.IsTranslatable
	}
	if o.MeshState != nil {
		toSerialize["meshState"] = o.MeshState
	}
	if o.MimeType != nil {
		toSerialize["mimeType"] = o.MimeType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ObsoletionRevisionId != nil {
		toSerialize["obsoletionRevisionId"] = o.ObsoletionRevisionId
	}
	if o.OriginalWorkspaceId != nil {
		toSerialize["originalWorkspaceId"] = o.OriginalWorkspaceId
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.PartType != nil {
		toSerialize["partType"] = o.PartType
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.ReferenceIds != nil {
		toSerialize["referenceIds"] = o.ReferenceIds
	}
	if o.ReferenceIdsFromOriginalWorkspace != nil {
		toSerialize["referenceIdsFromOriginalWorkspace"] = o.ReferenceIdsFromOriginalWorkspace
	}
	if o.Rpid != nil {
		toSerialize["rpid"] = o.Rpid
	}
	if o.SmallThumbnailHref != nil {
		toSerialize["smallThumbnailHref"] = o.SmallThumbnailHref
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	if o.WorkspaceId != nil {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableBTReleasePackageItemInfo struct {
	value *BTReleasePackageItemInfo
	isSet bool
}

func (v NullableBTReleasePackageItemInfo) Get() *BTReleasePackageItemInfo {
	return v.value
}

func (v *NullableBTReleasePackageItemInfo) Set(val *BTReleasePackageItemInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTReleasePackageItemInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTReleasePackageItemInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTReleasePackageItemInfo(val *BTReleasePackageItemInfo) *NullableBTReleasePackageItemInfo {
	return &NullableBTReleasePackageItemInfo{value: val, isSet: true}
}

func (v NullableBTReleasePackageItemInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTReleasePackageItemInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
