/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.6979-8ce9514d51cf
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTReleasePackageInfo struct for BTReleasePackageInfo
type BTReleasePackageInfo struct {
	ChangeOrderId       *string                    `json:"changeOrderId,omitempty"`
	ColumnNames         *map[string]string         `json:"columnNames,omitempty"`
	Comments            []BTCommentInfo            `json:"comments,omitempty"`
	CompanyId           *string                    `json:"companyId,omitempty"`
	CreatedAt           *JSONTime                  `json:"createdAt,omitempty"`
	CreatedBy           *BTUserBasicSummaryInfo    `json:"createdBy,omitempty"`
	Description         *string                    `json:"description,omitempty"`
	Detailed            *bool                      `json:"detailed,omitempty"`
	DocumentId          *string                    `json:"documentId,omitempty"`
	Href                *string                    `json:"href,omitempty"`
	Id                  *string                    `json:"id,omitempty"`
	IsObsoletion        *bool                      `json:"isObsoletion,omitempty"`
	Items               []BTReleasePackageItemInfo `json:"items,omitempty"`
	LinkedVersionIds    []string                   `json:"linkedVersionIds,omitempty"`
	ModifiedAt          *JSONTime                  `json:"modifiedAt,omitempty"`
	ModifiedBy          *BTUserBasicSummaryInfo    `json:"modifiedBy,omitempty"`
	Name                *string                    `json:"name,omitempty"`
	OriginalWorkspaceId *string                    `json:"originalWorkspaceId,omitempty"`
	PackageThumbnail    *string                    `json:"packageThumbnail,omitempty"`
	ParentComments      []BTReleaseCommentListInfo `json:"parentComments,omitempty"`
	ParentPackages      []string                   `json:"parentPackages,omitempty"`
	Properties          []BTWorkflowPropertyInfo   `json:"properties,omitempty"`
	RevisionRuleId      *string                    `json:"revisionRuleId,omitempty"`
	VersionId           *string                    `json:"versionId,omitempty"`
	ViewRef             *string                    `json:"viewRef,omitempty"`
	Workflow            *BTWorkflowSnapshotInfo    `json:"workflow,omitempty"`
	WorkflowError       *string                    `json:"workflowError,omitempty"`
	WorkflowId          *BTPublishedWorkflowId     `json:"workflowId,omitempty"`
	WorkspaceId         *string                    `json:"workspaceId,omitempty"`
}

// NewBTReleasePackageInfo instantiates a new BTReleasePackageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTReleasePackageInfo() *BTReleasePackageInfo {
	this := BTReleasePackageInfo{}
	return &this
}

// NewBTReleasePackageInfoWithDefaults instantiates a new BTReleasePackageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTReleasePackageInfoWithDefaults() *BTReleasePackageInfo {
	this := BTReleasePackageInfo{}
	return &this
}

// GetChangeOrderId returns the ChangeOrderId field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetChangeOrderId() string {
	if o == nil || o.ChangeOrderId == nil {
		var ret string
		return ret
	}
	return *o.ChangeOrderId
}

// GetChangeOrderIdOk returns a tuple with the ChangeOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetChangeOrderIdOk() (*string, bool) {
	if o == nil || o.ChangeOrderId == nil {
		return nil, false
	}
	return o.ChangeOrderId, true
}

// HasChangeOrderId returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasChangeOrderId() bool {
	if o != nil && o.ChangeOrderId != nil {
		return true
	}

	return false
}

// SetChangeOrderId gets a reference to the given string and assigns it to the ChangeOrderId field.
func (o *BTReleasePackageInfo) SetChangeOrderId(v string) {
	o.ChangeOrderId = &v
}

// GetColumnNames returns the ColumnNames field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetColumnNames() map[string]string {
	if o == nil || o.ColumnNames == nil {
		var ret map[string]string
		return ret
	}
	return *o.ColumnNames
}

// GetColumnNamesOk returns a tuple with the ColumnNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetColumnNamesOk() (*map[string]string, bool) {
	if o == nil || o.ColumnNames == nil {
		return nil, false
	}
	return o.ColumnNames, true
}

// HasColumnNames returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasColumnNames() bool {
	if o != nil && o.ColumnNames != nil {
		return true
	}

	return false
}

// SetColumnNames gets a reference to the given map[string]string and assigns it to the ColumnNames field.
func (o *BTReleasePackageInfo) SetColumnNames(v map[string]string) {
	o.ColumnNames = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetComments() []BTCommentInfo {
	if o == nil || o.Comments == nil {
		var ret []BTCommentInfo
		return ret
	}
	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetCommentsOk() ([]BTCommentInfo, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given []BTCommentInfo and assigns it to the Comments field.
func (o *BTReleasePackageInfo) SetComments(v []BTCommentInfo) {
	o.Comments = v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTReleasePackageInfo) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetCreatedAt() JSONTime {
	if o == nil || o.CreatedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetCreatedAtOk() (*JSONTime, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given JSONTime and assigns it to the CreatedAt field.
func (o *BTReleasePackageInfo) SetCreatedAt(v JSONTime) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetCreatedBy() BTUserBasicSummaryInfo {
	if o == nil || o.CreatedBy == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given BTUserBasicSummaryInfo and assigns it to the CreatedBy field.
func (o *BTReleasePackageInfo) SetCreatedBy(v BTUserBasicSummaryInfo) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTReleasePackageInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDetailed returns the Detailed field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetDetailed() bool {
	if o == nil || o.Detailed == nil {
		var ret bool
		return ret
	}
	return *o.Detailed
}

// GetDetailedOk returns a tuple with the Detailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetDetailedOk() (*bool, bool) {
	if o == nil || o.Detailed == nil {
		return nil, false
	}
	return o.Detailed, true
}

// HasDetailed returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasDetailed() bool {
	if o != nil && o.Detailed != nil {
		return true
	}

	return false
}

// SetDetailed gets a reference to the given bool and assigns it to the Detailed field.
func (o *BTReleasePackageInfo) SetDetailed(v bool) {
	o.Detailed = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTReleasePackageInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTReleasePackageInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTReleasePackageInfo) SetId(v string) {
	o.Id = &v
}

// GetIsObsoletion returns the IsObsoletion field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetIsObsoletion() bool {
	if o == nil || o.IsObsoletion == nil {
		var ret bool
		return ret
	}
	return *o.IsObsoletion
}

// GetIsObsoletionOk returns a tuple with the IsObsoletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetIsObsoletionOk() (*bool, bool) {
	if o == nil || o.IsObsoletion == nil {
		return nil, false
	}
	return o.IsObsoletion, true
}

// HasIsObsoletion returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasIsObsoletion() bool {
	if o != nil && o.IsObsoletion != nil {
		return true
	}

	return false
}

// SetIsObsoletion gets a reference to the given bool and assigns it to the IsObsoletion field.
func (o *BTReleasePackageInfo) SetIsObsoletion(v bool) {
	o.IsObsoletion = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetItems() []BTReleasePackageItemInfo {
	if o == nil || o.Items == nil {
		var ret []BTReleasePackageItemInfo
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetItemsOk() ([]BTReleasePackageItemInfo, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BTReleasePackageItemInfo and assigns it to the Items field.
func (o *BTReleasePackageInfo) SetItems(v []BTReleasePackageItemInfo) {
	o.Items = v
}

// GetLinkedVersionIds returns the LinkedVersionIds field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetLinkedVersionIds() []string {
	if o == nil || o.LinkedVersionIds == nil {
		var ret []string
		return ret
	}
	return o.LinkedVersionIds
}

// GetLinkedVersionIdsOk returns a tuple with the LinkedVersionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetLinkedVersionIdsOk() ([]string, bool) {
	if o == nil || o.LinkedVersionIds == nil {
		return nil, false
	}
	return o.LinkedVersionIds, true
}

// HasLinkedVersionIds returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasLinkedVersionIds() bool {
	if o != nil && o.LinkedVersionIds != nil {
		return true
	}

	return false
}

// SetLinkedVersionIds gets a reference to the given []string and assigns it to the LinkedVersionIds field.
func (o *BTReleasePackageInfo) SetLinkedVersionIds(v []string) {
	o.LinkedVersionIds = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetModifiedAt() JSONTime {
	if o == nil || o.ModifiedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetModifiedAtOk() (*JSONTime, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given JSONTime and assigns it to the ModifiedAt field.
func (o *BTReleasePackageInfo) SetModifiedAt(v JSONTime) {
	o.ModifiedAt = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetModifiedBy() BTUserBasicSummaryInfo {
	if o == nil || o.ModifiedBy == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given BTUserBasicSummaryInfo and assigns it to the ModifiedBy field.
func (o *BTReleasePackageInfo) SetModifiedBy(v BTUserBasicSummaryInfo) {
	o.ModifiedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTReleasePackageInfo) SetName(v string) {
	o.Name = &v
}

// GetOriginalWorkspaceId returns the OriginalWorkspaceId field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetOriginalWorkspaceId() string {
	if o == nil || o.OriginalWorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.OriginalWorkspaceId
}

// GetOriginalWorkspaceIdOk returns a tuple with the OriginalWorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetOriginalWorkspaceIdOk() (*string, bool) {
	if o == nil || o.OriginalWorkspaceId == nil {
		return nil, false
	}
	return o.OriginalWorkspaceId, true
}

// HasOriginalWorkspaceId returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasOriginalWorkspaceId() bool {
	if o != nil && o.OriginalWorkspaceId != nil {
		return true
	}

	return false
}

// SetOriginalWorkspaceId gets a reference to the given string and assigns it to the OriginalWorkspaceId field.
func (o *BTReleasePackageInfo) SetOriginalWorkspaceId(v string) {
	o.OriginalWorkspaceId = &v
}

// GetPackageThumbnail returns the PackageThumbnail field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetPackageThumbnail() string {
	if o == nil || o.PackageThumbnail == nil {
		var ret string
		return ret
	}
	return *o.PackageThumbnail
}

// GetPackageThumbnailOk returns a tuple with the PackageThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetPackageThumbnailOk() (*string, bool) {
	if o == nil || o.PackageThumbnail == nil {
		return nil, false
	}
	return o.PackageThumbnail, true
}

// HasPackageThumbnail returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasPackageThumbnail() bool {
	if o != nil && o.PackageThumbnail != nil {
		return true
	}

	return false
}

// SetPackageThumbnail gets a reference to the given string and assigns it to the PackageThumbnail field.
func (o *BTReleasePackageInfo) SetPackageThumbnail(v string) {
	o.PackageThumbnail = &v
}

// GetParentComments returns the ParentComments field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetParentComments() []BTReleaseCommentListInfo {
	if o == nil || o.ParentComments == nil {
		var ret []BTReleaseCommentListInfo
		return ret
	}
	return o.ParentComments
}

// GetParentCommentsOk returns a tuple with the ParentComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetParentCommentsOk() ([]BTReleaseCommentListInfo, bool) {
	if o == nil || o.ParentComments == nil {
		return nil, false
	}
	return o.ParentComments, true
}

// HasParentComments returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasParentComments() bool {
	if o != nil && o.ParentComments != nil {
		return true
	}

	return false
}

// SetParentComments gets a reference to the given []BTReleaseCommentListInfo and assigns it to the ParentComments field.
func (o *BTReleasePackageInfo) SetParentComments(v []BTReleaseCommentListInfo) {
	o.ParentComments = v
}

// GetParentPackages returns the ParentPackages field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetParentPackages() []string {
	if o == nil || o.ParentPackages == nil {
		var ret []string
		return ret
	}
	return o.ParentPackages
}

// GetParentPackagesOk returns a tuple with the ParentPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetParentPackagesOk() ([]string, bool) {
	if o == nil || o.ParentPackages == nil {
		return nil, false
	}
	return o.ParentPackages, true
}

// HasParentPackages returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasParentPackages() bool {
	if o != nil && o.ParentPackages != nil {
		return true
	}

	return false
}

// SetParentPackages gets a reference to the given []string and assigns it to the ParentPackages field.
func (o *BTReleasePackageInfo) SetParentPackages(v []string) {
	o.ParentPackages = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetProperties() []BTWorkflowPropertyInfo {
	if o == nil || o.Properties == nil {
		var ret []BTWorkflowPropertyInfo
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetPropertiesOk() ([]BTWorkflowPropertyInfo, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []BTWorkflowPropertyInfo and assigns it to the Properties field.
func (o *BTReleasePackageInfo) SetProperties(v []BTWorkflowPropertyInfo) {
	o.Properties = v
}

// GetRevisionRuleId returns the RevisionRuleId field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetRevisionRuleId() string {
	if o == nil || o.RevisionRuleId == nil {
		var ret string
		return ret
	}
	return *o.RevisionRuleId
}

// GetRevisionRuleIdOk returns a tuple with the RevisionRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetRevisionRuleIdOk() (*string, bool) {
	if o == nil || o.RevisionRuleId == nil {
		return nil, false
	}
	return o.RevisionRuleId, true
}

// HasRevisionRuleId returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasRevisionRuleId() bool {
	if o != nil && o.RevisionRuleId != nil {
		return true
	}

	return false
}

// SetRevisionRuleId gets a reference to the given string and assigns it to the RevisionRuleId field.
func (o *BTReleasePackageInfo) SetRevisionRuleId(v string) {
	o.RevisionRuleId = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTReleasePackageInfo) SetVersionId(v string) {
	o.VersionId = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTReleasePackageInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetWorkflow() BTWorkflowSnapshotInfo {
	if o == nil || o.Workflow == nil {
		var ret BTWorkflowSnapshotInfo
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetWorkflowOk() (*BTWorkflowSnapshotInfo, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasWorkflow() bool {
	if o != nil && o.Workflow != nil {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given BTWorkflowSnapshotInfo and assigns it to the Workflow field.
func (o *BTReleasePackageInfo) SetWorkflow(v BTWorkflowSnapshotInfo) {
	o.Workflow = &v
}

// GetWorkflowError returns the WorkflowError field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetWorkflowError() string {
	if o == nil || o.WorkflowError == nil {
		var ret string
		return ret
	}
	return *o.WorkflowError
}

// GetWorkflowErrorOk returns a tuple with the WorkflowError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetWorkflowErrorOk() (*string, bool) {
	if o == nil || o.WorkflowError == nil {
		return nil, false
	}
	return o.WorkflowError, true
}

// HasWorkflowError returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasWorkflowError() bool {
	if o != nil && o.WorkflowError != nil {
		return true
	}

	return false
}

// SetWorkflowError gets a reference to the given string and assigns it to the WorkflowError field.
func (o *BTReleasePackageInfo) SetWorkflowError(v string) {
	o.WorkflowError = &v
}

// GetWorkflowId returns the WorkflowId field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetWorkflowId() BTPublishedWorkflowId {
	if o == nil || o.WorkflowId == nil {
		var ret BTPublishedWorkflowId
		return ret
	}
	return *o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetWorkflowIdOk() (*BTPublishedWorkflowId, bool) {
	if o == nil || o.WorkflowId == nil {
		return nil, false
	}
	return o.WorkflowId, true
}

// HasWorkflowId returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasWorkflowId() bool {
	if o != nil && o.WorkflowId != nil {
		return true
	}

	return false
}

// SetWorkflowId gets a reference to the given BTPublishedWorkflowId and assigns it to the WorkflowId field.
func (o *BTReleasePackageInfo) SetWorkflowId(v BTPublishedWorkflowId) {
	o.WorkflowId = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *BTReleasePackageInfo) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageInfo) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *BTReleasePackageInfo) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *BTReleasePackageInfo) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o BTReleasePackageInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangeOrderId != nil {
		toSerialize["changeOrderId"] = o.ChangeOrderId
	}
	if o.ColumnNames != nil {
		toSerialize["columnNames"] = o.ColumnNames
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
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
	if o.Detailed != nil {
		toSerialize["detailed"] = o.Detailed
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsObsoletion != nil {
		toSerialize["isObsoletion"] = o.IsObsoletion
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.LinkedVersionIds != nil {
		toSerialize["linkedVersionIds"] = o.LinkedVersionIds
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
	if o.OriginalWorkspaceId != nil {
		toSerialize["originalWorkspaceId"] = o.OriginalWorkspaceId
	}
	if o.PackageThumbnail != nil {
		toSerialize["packageThumbnail"] = o.PackageThumbnail
	}
	if o.ParentComments != nil {
		toSerialize["parentComments"] = o.ParentComments
	}
	if o.ParentPackages != nil {
		toSerialize["parentPackages"] = o.ParentPackages
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.RevisionRuleId != nil {
		toSerialize["revisionRuleId"] = o.RevisionRuleId
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	if o.Workflow != nil {
		toSerialize["workflow"] = o.Workflow
	}
	if o.WorkflowError != nil {
		toSerialize["workflowError"] = o.WorkflowError
	}
	if o.WorkflowId != nil {
		toSerialize["workflowId"] = o.WorkflowId
	}
	if o.WorkspaceId != nil {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableBTReleasePackageInfo struct {
	value *BTReleasePackageInfo
	isSet bool
}

func (v NullableBTReleasePackageInfo) Get() *BTReleasePackageInfo {
	return v.value
}

func (v *NullableBTReleasePackageInfo) Set(val *BTReleasePackageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTReleasePackageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTReleasePackageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTReleasePackageInfo(val *BTReleasePackageInfo) *NullableBTReleasePackageInfo {
	return &NullableBTReleasePackageInfo{value: val, isSet: true}
}

func (v NullableBTReleasePackageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTReleasePackageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
