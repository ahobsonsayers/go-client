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

// BTFolderInfo struct for BTFolderInfo
type BTFolderInfo struct {
	CanMove           *bool                   `json:"canMove,omitempty"`
	CreatedAt         *JSONTime               `json:"createdAt,omitempty"`
	CreatedBy         *BTUserBasicSummaryInfo `json:"createdBy,omitempty"`
	Description       *string                 `json:"description,omitempty"`
	Href              *string                 `json:"href,omitempty"`
	Id                *string                 `json:"id,omitempty"`
	IsContainer       *bool                   `json:"isContainer,omitempty"`
	IsEnterpriseOwned *bool                   `json:"isEnterpriseOwned,omitempty"`
	IsMutable         *bool                   `json:"isMutable,omitempty"`
	JsonType          string                  `json:"jsonType"`
	ModifiedAt        *JSONTime               `json:"modifiedAt,omitempty"`
	ModifiedBy        *BTUserBasicSummaryInfo `json:"modifiedBy,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Owner             *BTOwnerInfo            `json:"owner,omitempty"`
	ProjectId         *string                 `json:"projectId,omitempty"`
	ResourceType      *string                 `json:"resourceType,omitempty"`
	TreeHref          *string                 `json:"treeHref,omitempty"`
	ViewRef           *string                 `json:"viewRef,omitempty"`
	Active            *bool                   `json:"active,omitempty"`
	CanUnshare        *bool                   `json:"canUnshare,omitempty"`
	IsOrphaned        *bool                   `json:"isOrphaned,omitempty"`
	ParentId          *string                 `json:"parentId,omitempty"`
	PermissionSet     []string                `json:"permissionSet,omitempty"`
	Trash             *bool                   `json:"trash,omitempty"`
	TrashedAt         *JSONTime               `json:"trashedAt,omitempty"`
}

// NewBTFolderInfo instantiates a new BTFolderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFolderInfo(jsonType string) *BTFolderInfo {
	this := BTFolderInfo{}
	this.JsonType = jsonType
	return &this
}

// NewBTFolderInfoWithDefaults instantiates a new BTFolderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFolderInfoWithDefaults() *BTFolderInfo {
	this := BTFolderInfo{}
	return &this
}

// GetCanMove returns the CanMove field value if set, zero value otherwise.
func (o *BTFolderInfo) GetCanMove() bool {
	if o == nil || o.CanMove == nil {
		var ret bool
		return ret
	}
	return *o.CanMove
}

// GetCanMoveOk returns a tuple with the CanMove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetCanMoveOk() (*bool, bool) {
	if o == nil || o.CanMove == nil {
		return nil, false
	}
	return o.CanMove, true
}

// HasCanMove returns a boolean if a field has been set.
func (o *BTFolderInfo) HasCanMove() bool {
	if o != nil && o.CanMove != nil {
		return true
	}

	return false
}

// SetCanMove gets a reference to the given bool and assigns it to the CanMove field.
func (o *BTFolderInfo) SetCanMove(v bool) {
	o.CanMove = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BTFolderInfo) GetCreatedAt() JSONTime {
	if o == nil || o.CreatedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetCreatedAtOk() (*JSONTime, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BTFolderInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given JSONTime and assigns it to the CreatedAt field.
func (o *BTFolderInfo) SetCreatedAt(v JSONTime) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BTFolderInfo) GetCreatedBy() BTUserBasicSummaryInfo {
	if o == nil || o.CreatedBy == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BTFolderInfo) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given BTUserBasicSummaryInfo and assigns it to the CreatedBy field.
func (o *BTFolderInfo) SetCreatedBy(v BTUserBasicSummaryInfo) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTFolderInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTFolderInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTFolderInfo) SetDescription(v string) {
	o.Description = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTFolderInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTFolderInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTFolderInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTFolderInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTFolderInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTFolderInfo) SetId(v string) {
	o.Id = &v
}

// GetIsContainer returns the IsContainer field value if set, zero value otherwise.
func (o *BTFolderInfo) GetIsContainer() bool {
	if o == nil || o.IsContainer == nil {
		var ret bool
		return ret
	}
	return *o.IsContainer
}

// GetIsContainerOk returns a tuple with the IsContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetIsContainerOk() (*bool, bool) {
	if o == nil || o.IsContainer == nil {
		return nil, false
	}
	return o.IsContainer, true
}

// HasIsContainer returns a boolean if a field has been set.
func (o *BTFolderInfo) HasIsContainer() bool {
	if o != nil && o.IsContainer != nil {
		return true
	}

	return false
}

// SetIsContainer gets a reference to the given bool and assigns it to the IsContainer field.
func (o *BTFolderInfo) SetIsContainer(v bool) {
	o.IsContainer = &v
}

// GetIsEnterpriseOwned returns the IsEnterpriseOwned field value if set, zero value otherwise.
func (o *BTFolderInfo) GetIsEnterpriseOwned() bool {
	if o == nil || o.IsEnterpriseOwned == nil {
		var ret bool
		return ret
	}
	return *o.IsEnterpriseOwned
}

// GetIsEnterpriseOwnedOk returns a tuple with the IsEnterpriseOwned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetIsEnterpriseOwnedOk() (*bool, bool) {
	if o == nil || o.IsEnterpriseOwned == nil {
		return nil, false
	}
	return o.IsEnterpriseOwned, true
}

// HasIsEnterpriseOwned returns a boolean if a field has been set.
func (o *BTFolderInfo) HasIsEnterpriseOwned() bool {
	if o != nil && o.IsEnterpriseOwned != nil {
		return true
	}

	return false
}

// SetIsEnterpriseOwned gets a reference to the given bool and assigns it to the IsEnterpriseOwned field.
func (o *BTFolderInfo) SetIsEnterpriseOwned(v bool) {
	o.IsEnterpriseOwned = &v
}

// GetIsMutable returns the IsMutable field value if set, zero value otherwise.
func (o *BTFolderInfo) GetIsMutable() bool {
	if o == nil || o.IsMutable == nil {
		var ret bool
		return ret
	}
	return *o.IsMutable
}

// GetIsMutableOk returns a tuple with the IsMutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetIsMutableOk() (*bool, bool) {
	if o == nil || o.IsMutable == nil {
		return nil, false
	}
	return o.IsMutable, true
}

// HasIsMutable returns a boolean if a field has been set.
func (o *BTFolderInfo) HasIsMutable() bool {
	if o != nil && o.IsMutable != nil {
		return true
	}

	return false
}

// SetIsMutable gets a reference to the given bool and assigns it to the IsMutable field.
func (o *BTFolderInfo) SetIsMutable(v bool) {
	o.IsMutable = &v
}

// GetJsonType returns the JsonType field value
func (o *BTFolderInfo) GetJsonType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetJsonTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsonType, true
}

// SetJsonType sets field value
func (o *BTFolderInfo) SetJsonType(v string) {
	o.JsonType = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *BTFolderInfo) GetModifiedAt() JSONTime {
	if o == nil || o.ModifiedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetModifiedAtOk() (*JSONTime, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *BTFolderInfo) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given JSONTime and assigns it to the ModifiedAt field.
func (o *BTFolderInfo) SetModifiedAt(v JSONTime) {
	o.ModifiedAt = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *BTFolderInfo) GetModifiedBy() BTUserBasicSummaryInfo {
	if o == nil || o.ModifiedBy == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *BTFolderInfo) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given BTUserBasicSummaryInfo and assigns it to the ModifiedBy field.
func (o *BTFolderInfo) SetModifiedBy(v BTUserBasicSummaryInfo) {
	o.ModifiedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTFolderInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTFolderInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTFolderInfo) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *BTFolderInfo) GetOwner() BTOwnerInfo {
	if o == nil || o.Owner == nil {
		var ret BTOwnerInfo
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetOwnerOk() (*BTOwnerInfo, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *BTFolderInfo) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given BTOwnerInfo and assigns it to the Owner field.
func (o *BTFolderInfo) SetOwner(v BTOwnerInfo) {
	o.Owner = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BTFolderInfo) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BTFolderInfo) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *BTFolderInfo) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *BTFolderInfo) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *BTFolderInfo) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *BTFolderInfo) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetTreeHref returns the TreeHref field value if set, zero value otherwise.
func (o *BTFolderInfo) GetTreeHref() string {
	if o == nil || o.TreeHref == nil {
		var ret string
		return ret
	}
	return *o.TreeHref
}

// GetTreeHrefOk returns a tuple with the TreeHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetTreeHrefOk() (*string, bool) {
	if o == nil || o.TreeHref == nil {
		return nil, false
	}
	return o.TreeHref, true
}

// HasTreeHref returns a boolean if a field has been set.
func (o *BTFolderInfo) HasTreeHref() bool {
	if o != nil && o.TreeHref != nil {
		return true
	}

	return false
}

// SetTreeHref gets a reference to the given string and assigns it to the TreeHref field.
func (o *BTFolderInfo) SetTreeHref(v string) {
	o.TreeHref = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTFolderInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTFolderInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTFolderInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *BTFolderInfo) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *BTFolderInfo) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *BTFolderInfo) SetActive(v bool) {
	o.Active = &v
}

// GetCanUnshare returns the CanUnshare field value if set, zero value otherwise.
func (o *BTFolderInfo) GetCanUnshare() bool {
	if o == nil || o.CanUnshare == nil {
		var ret bool
		return ret
	}
	return *o.CanUnshare
}

// GetCanUnshareOk returns a tuple with the CanUnshare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetCanUnshareOk() (*bool, bool) {
	if o == nil || o.CanUnshare == nil {
		return nil, false
	}
	return o.CanUnshare, true
}

// HasCanUnshare returns a boolean if a field has been set.
func (o *BTFolderInfo) HasCanUnshare() bool {
	if o != nil && o.CanUnshare != nil {
		return true
	}

	return false
}

// SetCanUnshare gets a reference to the given bool and assigns it to the CanUnshare field.
func (o *BTFolderInfo) SetCanUnshare(v bool) {
	o.CanUnshare = &v
}

// GetIsOrphaned returns the IsOrphaned field value if set, zero value otherwise.
func (o *BTFolderInfo) GetIsOrphaned() bool {
	if o == nil || o.IsOrphaned == nil {
		var ret bool
		return ret
	}
	return *o.IsOrphaned
}

// GetIsOrphanedOk returns a tuple with the IsOrphaned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetIsOrphanedOk() (*bool, bool) {
	if o == nil || o.IsOrphaned == nil {
		return nil, false
	}
	return o.IsOrphaned, true
}

// HasIsOrphaned returns a boolean if a field has been set.
func (o *BTFolderInfo) HasIsOrphaned() bool {
	if o != nil && o.IsOrphaned != nil {
		return true
	}

	return false
}

// SetIsOrphaned gets a reference to the given bool and assigns it to the IsOrphaned field.
func (o *BTFolderInfo) SetIsOrphaned(v bool) {
	o.IsOrphaned = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *BTFolderInfo) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *BTFolderInfo) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *BTFolderInfo) SetParentId(v string) {
	o.ParentId = &v
}

// GetPermissionSet returns the PermissionSet field value if set, zero value otherwise.
func (o *BTFolderInfo) GetPermissionSet() []string {
	if o == nil || o.PermissionSet == nil {
		var ret []string
		return ret
	}
	return o.PermissionSet
}

// GetPermissionSetOk returns a tuple with the PermissionSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetPermissionSetOk() ([]string, bool) {
	if o == nil || o.PermissionSet == nil {
		return nil, false
	}
	return o.PermissionSet, true
}

// HasPermissionSet returns a boolean if a field has been set.
func (o *BTFolderInfo) HasPermissionSet() bool {
	if o != nil && o.PermissionSet != nil {
		return true
	}

	return false
}

// SetPermissionSet gets a reference to the given []string and assigns it to the PermissionSet field.
func (o *BTFolderInfo) SetPermissionSet(v []string) {
	o.PermissionSet = v
}

// GetTrash returns the Trash field value if set, zero value otherwise.
func (o *BTFolderInfo) GetTrash() bool {
	if o == nil || o.Trash == nil {
		var ret bool
		return ret
	}
	return *o.Trash
}

// GetTrashOk returns a tuple with the Trash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetTrashOk() (*bool, bool) {
	if o == nil || o.Trash == nil {
		return nil, false
	}
	return o.Trash, true
}

// HasTrash returns a boolean if a field has been set.
func (o *BTFolderInfo) HasTrash() bool {
	if o != nil && o.Trash != nil {
		return true
	}

	return false
}

// SetTrash gets a reference to the given bool and assigns it to the Trash field.
func (o *BTFolderInfo) SetTrash(v bool) {
	o.Trash = &v
}

// GetTrashedAt returns the TrashedAt field value if set, zero value otherwise.
func (o *BTFolderInfo) GetTrashedAt() JSONTime {
	if o == nil || o.TrashedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.TrashedAt
}

// GetTrashedAtOk returns a tuple with the TrashedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfo) GetTrashedAtOk() (*JSONTime, bool) {
	if o == nil || o.TrashedAt == nil {
		return nil, false
	}
	return o.TrashedAt, true
}

// HasTrashedAt returns a boolean if a field has been set.
func (o *BTFolderInfo) HasTrashedAt() bool {
	if o != nil && o.TrashedAt != nil {
		return true
	}

	return false
}

// SetTrashedAt gets a reference to the given JSONTime and assigns it to the TrashedAt field.
func (o *BTFolderInfo) SetTrashedAt(v JSONTime) {
	o.TrashedAt = &v
}

func (o BTFolderInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanMove != nil {
		toSerialize["canMove"] = o.CanMove
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsContainer != nil {
		toSerialize["isContainer"] = o.IsContainer
	}
	if o.IsEnterpriseOwned != nil {
		toSerialize["isEnterpriseOwned"] = o.IsEnterpriseOwned
	}
	if o.IsMutable != nil {
		toSerialize["isMutable"] = o.IsMutable
	}
	if true {
		toSerialize["jsonType"] = o.JsonType
	}
	if o.ModifiedAt != nil {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if o.ModifiedBy != nil {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}
	if o.TreeHref != nil {
		toSerialize["treeHref"] = o.TreeHref
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.CanUnshare != nil {
		toSerialize["canUnshare"] = o.CanUnshare
	}
	if o.IsOrphaned != nil {
		toSerialize["isOrphaned"] = o.IsOrphaned
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.PermissionSet != nil {
		toSerialize["permissionSet"] = o.PermissionSet
	}
	if o.Trash != nil {
		toSerialize["trash"] = o.Trash
	}
	if o.TrashedAt != nil {
		toSerialize["trashedAt"] = o.TrashedAt
	}
	return json.Marshal(toSerialize)
}

type NullableBTFolderInfo struct {
	value *BTFolderInfo
	isSet bool
}

func (v NullableBTFolderInfo) Get() *BTFolderInfo {
	return v.value
}

func (v *NullableBTFolderInfo) Set(val *BTFolderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFolderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFolderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFolderInfo(val *BTFolderInfo) *NullableBTFolderInfo {
	return &NullableBTFolderInfo{value: val, isSet: true}
}

func (v NullableBTFolderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFolderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
