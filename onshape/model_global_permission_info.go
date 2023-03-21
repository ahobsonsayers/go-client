/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13200-ff216a970a02
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// GlobalPermissionInfo struct for GlobalPermissionInfo
type GlobalPermissionInfo struct {
	AccessReports                   *bool `json:"accessReports,omitempty"`
	AdminEnterprise                 *bool `json:"adminEnterprise,omitempty"`
	AllowAppStoreAccess             *bool `json:"allowAppStoreAccess,omitempty"`
	AllowPublicDocumentsAccess      *bool `json:"allowPublicDocumentsAccess,omitempty"`
	ApproveReleases                 *bool `json:"approveReleases,omitempty"`
	CreateChangeOrders              *bool `json:"createChangeOrders,omitempty"`
	CreateChangeRequests            *bool `json:"createChangeRequests,omitempty"`
	CreateDocumentsInRoot           *bool `json:"createDocumentsInRoot,omitempty"`
	CreateProject                   *bool `json:"createProject,omitempty"`
	CreateReleases                  *bool `json:"createReleases,omitempty"`
	CreateTasks                     *bool `json:"createTasks,omitempty"`
	DeletePermanently               *bool `json:"deletePermanently,omitempty"`
	ManageGuestUsers                *bool `json:"manageGuestUsers,omitempty"`
	ManageNonGeometricItems         *bool `json:"manageNonGeometricItems,omitempty"`
	ManageRbac                      *bool `json:"manageRbac,omitempty"`
	ManageUsers                     *bool `json:"manageUsers,omitempty"`
	ManageWorkflows                 *bool `json:"manageWorkflows,omitempty"`
	ShareForAnonymousAccess         *bool `json:"shareForAnonymousAccess,omitempty"`
	TransferDocumentsFromEnterprise *bool `json:"transferDocumentsFromEnterprise,omitempty"`
	ViewChangeOrders                *bool `json:"viewChangeOrders,omitempty"`
	ViewChangeRequests              *bool `json:"viewChangeRequests,omitempty"`
}

// NewGlobalPermissionInfo instantiates a new GlobalPermissionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalPermissionInfo() *GlobalPermissionInfo {
	this := GlobalPermissionInfo{}
	return &this
}

// NewGlobalPermissionInfoWithDefaults instantiates a new GlobalPermissionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalPermissionInfoWithDefaults() *GlobalPermissionInfo {
	this := GlobalPermissionInfo{}
	return &this
}

// GetAccessReports returns the AccessReports field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetAccessReports() bool {
	if o == nil || o.AccessReports == nil {
		var ret bool
		return ret
	}
	return *o.AccessReports
}

// GetAccessReportsOk returns a tuple with the AccessReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetAccessReportsOk() (*bool, bool) {
	if o == nil || o.AccessReports == nil {
		return nil, false
	}
	return o.AccessReports, true
}

// HasAccessReports returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasAccessReports() bool {
	if o != nil && o.AccessReports != nil {
		return true
	}

	return false
}

// SetAccessReports gets a reference to the given bool and assigns it to the AccessReports field.
func (o *GlobalPermissionInfo) SetAccessReports(v bool) {
	o.AccessReports = &v
}

// GetAdminEnterprise returns the AdminEnterprise field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetAdminEnterprise() bool {
	if o == nil || o.AdminEnterprise == nil {
		var ret bool
		return ret
	}
	return *o.AdminEnterprise
}

// GetAdminEnterpriseOk returns a tuple with the AdminEnterprise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetAdminEnterpriseOk() (*bool, bool) {
	if o == nil || o.AdminEnterprise == nil {
		return nil, false
	}
	return o.AdminEnterprise, true
}

// HasAdminEnterprise returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasAdminEnterprise() bool {
	if o != nil && o.AdminEnterprise != nil {
		return true
	}

	return false
}

// SetAdminEnterprise gets a reference to the given bool and assigns it to the AdminEnterprise field.
func (o *GlobalPermissionInfo) SetAdminEnterprise(v bool) {
	o.AdminEnterprise = &v
}

// GetAllowAppStoreAccess returns the AllowAppStoreAccess field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetAllowAppStoreAccess() bool {
	if o == nil || o.AllowAppStoreAccess == nil {
		var ret bool
		return ret
	}
	return *o.AllowAppStoreAccess
}

// GetAllowAppStoreAccessOk returns a tuple with the AllowAppStoreAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetAllowAppStoreAccessOk() (*bool, bool) {
	if o == nil || o.AllowAppStoreAccess == nil {
		return nil, false
	}
	return o.AllowAppStoreAccess, true
}

// HasAllowAppStoreAccess returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasAllowAppStoreAccess() bool {
	if o != nil && o.AllowAppStoreAccess != nil {
		return true
	}

	return false
}

// SetAllowAppStoreAccess gets a reference to the given bool and assigns it to the AllowAppStoreAccess field.
func (o *GlobalPermissionInfo) SetAllowAppStoreAccess(v bool) {
	o.AllowAppStoreAccess = &v
}

// GetAllowPublicDocumentsAccess returns the AllowPublicDocumentsAccess field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetAllowPublicDocumentsAccess() bool {
	if o == nil || o.AllowPublicDocumentsAccess == nil {
		var ret bool
		return ret
	}
	return *o.AllowPublicDocumentsAccess
}

// GetAllowPublicDocumentsAccessOk returns a tuple with the AllowPublicDocumentsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetAllowPublicDocumentsAccessOk() (*bool, bool) {
	if o == nil || o.AllowPublicDocumentsAccess == nil {
		return nil, false
	}
	return o.AllowPublicDocumentsAccess, true
}

// HasAllowPublicDocumentsAccess returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasAllowPublicDocumentsAccess() bool {
	if o != nil && o.AllowPublicDocumentsAccess != nil {
		return true
	}

	return false
}

// SetAllowPublicDocumentsAccess gets a reference to the given bool and assigns it to the AllowPublicDocumentsAccess field.
func (o *GlobalPermissionInfo) SetAllowPublicDocumentsAccess(v bool) {
	o.AllowPublicDocumentsAccess = &v
}

// GetApproveReleases returns the ApproveReleases field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetApproveReleases() bool {
	if o == nil || o.ApproveReleases == nil {
		var ret bool
		return ret
	}
	return *o.ApproveReleases
}

// GetApproveReleasesOk returns a tuple with the ApproveReleases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetApproveReleasesOk() (*bool, bool) {
	if o == nil || o.ApproveReleases == nil {
		return nil, false
	}
	return o.ApproveReleases, true
}

// HasApproveReleases returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasApproveReleases() bool {
	if o != nil && o.ApproveReleases != nil {
		return true
	}

	return false
}

// SetApproveReleases gets a reference to the given bool and assigns it to the ApproveReleases field.
func (o *GlobalPermissionInfo) SetApproveReleases(v bool) {
	o.ApproveReleases = &v
}

// GetCreateChangeOrders returns the CreateChangeOrders field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetCreateChangeOrders() bool {
	if o == nil || o.CreateChangeOrders == nil {
		var ret bool
		return ret
	}
	return *o.CreateChangeOrders
}

// GetCreateChangeOrdersOk returns a tuple with the CreateChangeOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetCreateChangeOrdersOk() (*bool, bool) {
	if o == nil || o.CreateChangeOrders == nil {
		return nil, false
	}
	return o.CreateChangeOrders, true
}

// HasCreateChangeOrders returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasCreateChangeOrders() bool {
	if o != nil && o.CreateChangeOrders != nil {
		return true
	}

	return false
}

// SetCreateChangeOrders gets a reference to the given bool and assigns it to the CreateChangeOrders field.
func (o *GlobalPermissionInfo) SetCreateChangeOrders(v bool) {
	o.CreateChangeOrders = &v
}

// GetCreateChangeRequests returns the CreateChangeRequests field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetCreateChangeRequests() bool {
	if o == nil || o.CreateChangeRequests == nil {
		var ret bool
		return ret
	}
	return *o.CreateChangeRequests
}

// GetCreateChangeRequestsOk returns a tuple with the CreateChangeRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetCreateChangeRequestsOk() (*bool, bool) {
	if o == nil || o.CreateChangeRequests == nil {
		return nil, false
	}
	return o.CreateChangeRequests, true
}

// HasCreateChangeRequests returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasCreateChangeRequests() bool {
	if o != nil && o.CreateChangeRequests != nil {
		return true
	}

	return false
}

// SetCreateChangeRequests gets a reference to the given bool and assigns it to the CreateChangeRequests field.
func (o *GlobalPermissionInfo) SetCreateChangeRequests(v bool) {
	o.CreateChangeRequests = &v
}

// GetCreateDocumentsInRoot returns the CreateDocumentsInRoot field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetCreateDocumentsInRoot() bool {
	if o == nil || o.CreateDocumentsInRoot == nil {
		var ret bool
		return ret
	}
	return *o.CreateDocumentsInRoot
}

// GetCreateDocumentsInRootOk returns a tuple with the CreateDocumentsInRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetCreateDocumentsInRootOk() (*bool, bool) {
	if o == nil || o.CreateDocumentsInRoot == nil {
		return nil, false
	}
	return o.CreateDocumentsInRoot, true
}

// HasCreateDocumentsInRoot returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasCreateDocumentsInRoot() bool {
	if o != nil && o.CreateDocumentsInRoot != nil {
		return true
	}

	return false
}

// SetCreateDocumentsInRoot gets a reference to the given bool and assigns it to the CreateDocumentsInRoot field.
func (o *GlobalPermissionInfo) SetCreateDocumentsInRoot(v bool) {
	o.CreateDocumentsInRoot = &v
}

// GetCreateProject returns the CreateProject field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetCreateProject() bool {
	if o == nil || o.CreateProject == nil {
		var ret bool
		return ret
	}
	return *o.CreateProject
}

// GetCreateProjectOk returns a tuple with the CreateProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetCreateProjectOk() (*bool, bool) {
	if o == nil || o.CreateProject == nil {
		return nil, false
	}
	return o.CreateProject, true
}

// HasCreateProject returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasCreateProject() bool {
	if o != nil && o.CreateProject != nil {
		return true
	}

	return false
}

// SetCreateProject gets a reference to the given bool and assigns it to the CreateProject field.
func (o *GlobalPermissionInfo) SetCreateProject(v bool) {
	o.CreateProject = &v
}

// GetCreateReleases returns the CreateReleases field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetCreateReleases() bool {
	if o == nil || o.CreateReleases == nil {
		var ret bool
		return ret
	}
	return *o.CreateReleases
}

// GetCreateReleasesOk returns a tuple with the CreateReleases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetCreateReleasesOk() (*bool, bool) {
	if o == nil || o.CreateReleases == nil {
		return nil, false
	}
	return o.CreateReleases, true
}

// HasCreateReleases returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasCreateReleases() bool {
	if o != nil && o.CreateReleases != nil {
		return true
	}

	return false
}

// SetCreateReleases gets a reference to the given bool and assigns it to the CreateReleases field.
func (o *GlobalPermissionInfo) SetCreateReleases(v bool) {
	o.CreateReleases = &v
}

// GetCreateTasks returns the CreateTasks field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetCreateTasks() bool {
	if o == nil || o.CreateTasks == nil {
		var ret bool
		return ret
	}
	return *o.CreateTasks
}

// GetCreateTasksOk returns a tuple with the CreateTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetCreateTasksOk() (*bool, bool) {
	if o == nil || o.CreateTasks == nil {
		return nil, false
	}
	return o.CreateTasks, true
}

// HasCreateTasks returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasCreateTasks() bool {
	if o != nil && o.CreateTasks != nil {
		return true
	}

	return false
}

// SetCreateTasks gets a reference to the given bool and assigns it to the CreateTasks field.
func (o *GlobalPermissionInfo) SetCreateTasks(v bool) {
	o.CreateTasks = &v
}

// GetDeletePermanently returns the DeletePermanently field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetDeletePermanently() bool {
	if o == nil || o.DeletePermanently == nil {
		var ret bool
		return ret
	}
	return *o.DeletePermanently
}

// GetDeletePermanentlyOk returns a tuple with the DeletePermanently field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetDeletePermanentlyOk() (*bool, bool) {
	if o == nil || o.DeletePermanently == nil {
		return nil, false
	}
	return o.DeletePermanently, true
}

// HasDeletePermanently returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasDeletePermanently() bool {
	if o != nil && o.DeletePermanently != nil {
		return true
	}

	return false
}

// SetDeletePermanently gets a reference to the given bool and assigns it to the DeletePermanently field.
func (o *GlobalPermissionInfo) SetDeletePermanently(v bool) {
	o.DeletePermanently = &v
}

// GetManageGuestUsers returns the ManageGuestUsers field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetManageGuestUsers() bool {
	if o == nil || o.ManageGuestUsers == nil {
		var ret bool
		return ret
	}
	return *o.ManageGuestUsers
}

// GetManageGuestUsersOk returns a tuple with the ManageGuestUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetManageGuestUsersOk() (*bool, bool) {
	if o == nil || o.ManageGuestUsers == nil {
		return nil, false
	}
	return o.ManageGuestUsers, true
}

// HasManageGuestUsers returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasManageGuestUsers() bool {
	if o != nil && o.ManageGuestUsers != nil {
		return true
	}

	return false
}

// SetManageGuestUsers gets a reference to the given bool and assigns it to the ManageGuestUsers field.
func (o *GlobalPermissionInfo) SetManageGuestUsers(v bool) {
	o.ManageGuestUsers = &v
}

// GetManageNonGeometricItems returns the ManageNonGeometricItems field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetManageNonGeometricItems() bool {
	if o == nil || o.ManageNonGeometricItems == nil {
		var ret bool
		return ret
	}
	return *o.ManageNonGeometricItems
}

// GetManageNonGeometricItemsOk returns a tuple with the ManageNonGeometricItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetManageNonGeometricItemsOk() (*bool, bool) {
	if o == nil || o.ManageNonGeometricItems == nil {
		return nil, false
	}
	return o.ManageNonGeometricItems, true
}

// HasManageNonGeometricItems returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasManageNonGeometricItems() bool {
	if o != nil && o.ManageNonGeometricItems != nil {
		return true
	}

	return false
}

// SetManageNonGeometricItems gets a reference to the given bool and assigns it to the ManageNonGeometricItems field.
func (o *GlobalPermissionInfo) SetManageNonGeometricItems(v bool) {
	o.ManageNonGeometricItems = &v
}

// GetManageRbac returns the ManageRbac field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetManageRbac() bool {
	if o == nil || o.ManageRbac == nil {
		var ret bool
		return ret
	}
	return *o.ManageRbac
}

// GetManageRbacOk returns a tuple with the ManageRbac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetManageRbacOk() (*bool, bool) {
	if o == nil || o.ManageRbac == nil {
		return nil, false
	}
	return o.ManageRbac, true
}

// HasManageRbac returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasManageRbac() bool {
	if o != nil && o.ManageRbac != nil {
		return true
	}

	return false
}

// SetManageRbac gets a reference to the given bool and assigns it to the ManageRbac field.
func (o *GlobalPermissionInfo) SetManageRbac(v bool) {
	o.ManageRbac = &v
}

// GetManageUsers returns the ManageUsers field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetManageUsers() bool {
	if o == nil || o.ManageUsers == nil {
		var ret bool
		return ret
	}
	return *o.ManageUsers
}

// GetManageUsersOk returns a tuple with the ManageUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetManageUsersOk() (*bool, bool) {
	if o == nil || o.ManageUsers == nil {
		return nil, false
	}
	return o.ManageUsers, true
}

// HasManageUsers returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasManageUsers() bool {
	if o != nil && o.ManageUsers != nil {
		return true
	}

	return false
}

// SetManageUsers gets a reference to the given bool and assigns it to the ManageUsers field.
func (o *GlobalPermissionInfo) SetManageUsers(v bool) {
	o.ManageUsers = &v
}

// GetManageWorkflows returns the ManageWorkflows field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetManageWorkflows() bool {
	if o == nil || o.ManageWorkflows == nil {
		var ret bool
		return ret
	}
	return *o.ManageWorkflows
}

// GetManageWorkflowsOk returns a tuple with the ManageWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetManageWorkflowsOk() (*bool, bool) {
	if o == nil || o.ManageWorkflows == nil {
		return nil, false
	}
	return o.ManageWorkflows, true
}

// HasManageWorkflows returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasManageWorkflows() bool {
	if o != nil && o.ManageWorkflows != nil {
		return true
	}

	return false
}

// SetManageWorkflows gets a reference to the given bool and assigns it to the ManageWorkflows field.
func (o *GlobalPermissionInfo) SetManageWorkflows(v bool) {
	o.ManageWorkflows = &v
}

// GetShareForAnonymousAccess returns the ShareForAnonymousAccess field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetShareForAnonymousAccess() bool {
	if o == nil || o.ShareForAnonymousAccess == nil {
		var ret bool
		return ret
	}
	return *o.ShareForAnonymousAccess
}

// GetShareForAnonymousAccessOk returns a tuple with the ShareForAnonymousAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetShareForAnonymousAccessOk() (*bool, bool) {
	if o == nil || o.ShareForAnonymousAccess == nil {
		return nil, false
	}
	return o.ShareForAnonymousAccess, true
}

// HasShareForAnonymousAccess returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasShareForAnonymousAccess() bool {
	if o != nil && o.ShareForAnonymousAccess != nil {
		return true
	}

	return false
}

// SetShareForAnonymousAccess gets a reference to the given bool and assigns it to the ShareForAnonymousAccess field.
func (o *GlobalPermissionInfo) SetShareForAnonymousAccess(v bool) {
	o.ShareForAnonymousAccess = &v
}

// GetTransferDocumentsFromEnterprise returns the TransferDocumentsFromEnterprise field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetTransferDocumentsFromEnterprise() bool {
	if o == nil || o.TransferDocumentsFromEnterprise == nil {
		var ret bool
		return ret
	}
	return *o.TransferDocumentsFromEnterprise
}

// GetTransferDocumentsFromEnterpriseOk returns a tuple with the TransferDocumentsFromEnterprise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetTransferDocumentsFromEnterpriseOk() (*bool, bool) {
	if o == nil || o.TransferDocumentsFromEnterprise == nil {
		return nil, false
	}
	return o.TransferDocumentsFromEnterprise, true
}

// HasTransferDocumentsFromEnterprise returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasTransferDocumentsFromEnterprise() bool {
	if o != nil && o.TransferDocumentsFromEnterprise != nil {
		return true
	}

	return false
}

// SetTransferDocumentsFromEnterprise gets a reference to the given bool and assigns it to the TransferDocumentsFromEnterprise field.
func (o *GlobalPermissionInfo) SetTransferDocumentsFromEnterprise(v bool) {
	o.TransferDocumentsFromEnterprise = &v
}

// GetViewChangeOrders returns the ViewChangeOrders field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetViewChangeOrders() bool {
	if o == nil || o.ViewChangeOrders == nil {
		var ret bool
		return ret
	}
	return *o.ViewChangeOrders
}

// GetViewChangeOrdersOk returns a tuple with the ViewChangeOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetViewChangeOrdersOk() (*bool, bool) {
	if o == nil || o.ViewChangeOrders == nil {
		return nil, false
	}
	return o.ViewChangeOrders, true
}

// HasViewChangeOrders returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasViewChangeOrders() bool {
	if o != nil && o.ViewChangeOrders != nil {
		return true
	}

	return false
}

// SetViewChangeOrders gets a reference to the given bool and assigns it to the ViewChangeOrders field.
func (o *GlobalPermissionInfo) SetViewChangeOrders(v bool) {
	o.ViewChangeOrders = &v
}

// GetViewChangeRequests returns the ViewChangeRequests field value if set, zero value otherwise.
func (o *GlobalPermissionInfo) GetViewChangeRequests() bool {
	if o == nil || o.ViewChangeRequests == nil {
		var ret bool
		return ret
	}
	return *o.ViewChangeRequests
}

// GetViewChangeRequestsOk returns a tuple with the ViewChangeRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionInfo) GetViewChangeRequestsOk() (*bool, bool) {
	if o == nil || o.ViewChangeRequests == nil {
		return nil, false
	}
	return o.ViewChangeRequests, true
}

// HasViewChangeRequests returns a boolean if a field has been set.
func (o *GlobalPermissionInfo) HasViewChangeRequests() bool {
	if o != nil && o.ViewChangeRequests != nil {
		return true
	}

	return false
}

// SetViewChangeRequests gets a reference to the given bool and assigns it to the ViewChangeRequests field.
func (o *GlobalPermissionInfo) SetViewChangeRequests(v bool) {
	o.ViewChangeRequests = &v
}

func (o GlobalPermissionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessReports != nil {
		toSerialize["accessReports"] = o.AccessReports
	}
	if o.AdminEnterprise != nil {
		toSerialize["adminEnterprise"] = o.AdminEnterprise
	}
	if o.AllowAppStoreAccess != nil {
		toSerialize["allowAppStoreAccess"] = o.AllowAppStoreAccess
	}
	if o.AllowPublicDocumentsAccess != nil {
		toSerialize["allowPublicDocumentsAccess"] = o.AllowPublicDocumentsAccess
	}
	if o.ApproveReleases != nil {
		toSerialize["approveReleases"] = o.ApproveReleases
	}
	if o.CreateChangeOrders != nil {
		toSerialize["createChangeOrders"] = o.CreateChangeOrders
	}
	if o.CreateChangeRequests != nil {
		toSerialize["createChangeRequests"] = o.CreateChangeRequests
	}
	if o.CreateDocumentsInRoot != nil {
		toSerialize["createDocumentsInRoot"] = o.CreateDocumentsInRoot
	}
	if o.CreateProject != nil {
		toSerialize["createProject"] = o.CreateProject
	}
	if o.CreateReleases != nil {
		toSerialize["createReleases"] = o.CreateReleases
	}
	if o.CreateTasks != nil {
		toSerialize["createTasks"] = o.CreateTasks
	}
	if o.DeletePermanently != nil {
		toSerialize["deletePermanently"] = o.DeletePermanently
	}
	if o.ManageGuestUsers != nil {
		toSerialize["manageGuestUsers"] = o.ManageGuestUsers
	}
	if o.ManageNonGeometricItems != nil {
		toSerialize["manageNonGeometricItems"] = o.ManageNonGeometricItems
	}
	if o.ManageRbac != nil {
		toSerialize["manageRbac"] = o.ManageRbac
	}
	if o.ManageUsers != nil {
		toSerialize["manageUsers"] = o.ManageUsers
	}
	if o.ManageWorkflows != nil {
		toSerialize["manageWorkflows"] = o.ManageWorkflows
	}
	if o.ShareForAnonymousAccess != nil {
		toSerialize["shareForAnonymousAccess"] = o.ShareForAnonymousAccess
	}
	if o.TransferDocumentsFromEnterprise != nil {
		toSerialize["transferDocumentsFromEnterprise"] = o.TransferDocumentsFromEnterprise
	}
	if o.ViewChangeOrders != nil {
		toSerialize["viewChangeOrders"] = o.ViewChangeOrders
	}
	if o.ViewChangeRequests != nil {
		toSerialize["viewChangeRequests"] = o.ViewChangeRequests
	}
	return json.Marshal(toSerialize)
}

type NullableGlobalPermissionInfo struct {
	value *GlobalPermissionInfo
	isSet bool
}

func (v NullableGlobalPermissionInfo) Get() *GlobalPermissionInfo {
	return v.value
}

func (v *NullableGlobalPermissionInfo) Set(val *GlobalPermissionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalPermissionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalPermissionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalPermissionInfo(val *GlobalPermissionInfo) *NullableGlobalPermissionInfo {
	return &NullableGlobalPermissionInfo{value: val, isSet: true}
}

func (v NullableGlobalPermissionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalPermissionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
