/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6820-1bef41f9cc03
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTUserOAuth2SummaryInfo struct for BTUserOAuth2SummaryInfo
type BTUserOAuth2SummaryInfo struct {
	ClientId                  *string               `json:"clientId,omitempty"`
	Company                   *BTCompanySummaryInfo `json:"company,omitempty"`
	CompanyPlan               *bool                 `json:"companyPlan,omitempty"`
	DocumentationName         *string               `json:"documentationName,omitempty"`
	DocumentationNameOverride *string               `json:"documentationNameOverride,omitempty"`
	Email                     *string               `json:"email,omitempty"`
	FirstName                 *string               `json:"firstName,omitempty"`
	GlobalPermissions         *GlobalPermissionInfo `json:"globalPermissions,omitempty"`
	Href                      *string               `json:"href,omitempty"`
	Id                        *string               `json:"id,omitempty"`
	Image                     *string               `json:"image,omitempty"`
	IsGuest                   *bool                 `json:"isGuest,omitempty"`
	IsLight                   *bool                 `json:"isLight,omitempty"`
	LastLoginTime             *JSONTime             `json:"lastLoginTime,omitempty"`
	LastName                  *string               `json:"lastName,omitempty"`
	Name                      *string               `json:"name,omitempty"`
	Oauth2Scopes              *int64                `json:"oauth2Scopes,omitempty"`
	PersonalMessageAllowed    *bool                 `json:"personalMessageAllowed,omitempty"`
	PlanGroup                 *string               `json:"planGroup,omitempty"`
	Role                      *int32                `json:"role,omitempty"`
	Roles                     []string              `json:"roles,omitempty"`
	Source                    *int32                `json:"source,omitempty"`
	State                     *int32                `json:"state,omitempty"`
	ViewRef                   *string               `json:"viewRef,omitempty"`
}

// NewBTUserOAuth2SummaryInfo instantiates a new BTUserOAuth2SummaryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUserOAuth2SummaryInfo() *BTUserOAuth2SummaryInfo {
	this := BTUserOAuth2SummaryInfo{}
	return &this
}

// NewBTUserOAuth2SummaryInfoWithDefaults instantiates a new BTUserOAuth2SummaryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUserOAuth2SummaryInfoWithDefaults() *BTUserOAuth2SummaryInfo {
	this := BTUserOAuth2SummaryInfo{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *BTUserOAuth2SummaryInfo) SetClientId(v string) {
	o.ClientId = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetCompany() BTCompanySummaryInfo {
	if o == nil || o.Company == nil {
		var ret BTCompanySummaryInfo
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetCompanyOk() (*BTCompanySummaryInfo, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given BTCompanySummaryInfo and assigns it to the Company field.
func (o *BTUserOAuth2SummaryInfo) SetCompany(v BTCompanySummaryInfo) {
	o.Company = &v
}

// GetCompanyPlan returns the CompanyPlan field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetCompanyPlan() bool {
	if o == nil || o.CompanyPlan == nil {
		var ret bool
		return ret
	}
	return *o.CompanyPlan
}

// GetCompanyPlanOk returns a tuple with the CompanyPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetCompanyPlanOk() (*bool, bool) {
	if o == nil || o.CompanyPlan == nil {
		return nil, false
	}
	return o.CompanyPlan, true
}

// HasCompanyPlan returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasCompanyPlan() bool {
	if o != nil && o.CompanyPlan != nil {
		return true
	}

	return false
}

// SetCompanyPlan gets a reference to the given bool and assigns it to the CompanyPlan field.
func (o *BTUserOAuth2SummaryInfo) SetCompanyPlan(v bool) {
	o.CompanyPlan = &v
}

// GetDocumentationName returns the DocumentationName field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetDocumentationName() string {
	if o == nil || o.DocumentationName == nil {
		var ret string
		return ret
	}
	return *o.DocumentationName
}

// GetDocumentationNameOk returns a tuple with the DocumentationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetDocumentationNameOk() (*string, bool) {
	if o == nil || o.DocumentationName == nil {
		return nil, false
	}
	return o.DocumentationName, true
}

// HasDocumentationName returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasDocumentationName() bool {
	if o != nil && o.DocumentationName != nil {
		return true
	}

	return false
}

// SetDocumentationName gets a reference to the given string and assigns it to the DocumentationName field.
func (o *BTUserOAuth2SummaryInfo) SetDocumentationName(v string) {
	o.DocumentationName = &v
}

// GetDocumentationNameOverride returns the DocumentationNameOverride field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetDocumentationNameOverride() string {
	if o == nil || o.DocumentationNameOverride == nil {
		var ret string
		return ret
	}
	return *o.DocumentationNameOverride
}

// GetDocumentationNameOverrideOk returns a tuple with the DocumentationNameOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetDocumentationNameOverrideOk() (*string, bool) {
	if o == nil || o.DocumentationNameOverride == nil {
		return nil, false
	}
	return o.DocumentationNameOverride, true
}

// HasDocumentationNameOverride returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasDocumentationNameOverride() bool {
	if o != nil && o.DocumentationNameOverride != nil {
		return true
	}

	return false
}

// SetDocumentationNameOverride gets a reference to the given string and assigns it to the DocumentationNameOverride field.
func (o *BTUserOAuth2SummaryInfo) SetDocumentationNameOverride(v string) {
	o.DocumentationNameOverride = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BTUserOAuth2SummaryInfo) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *BTUserOAuth2SummaryInfo) SetFirstName(v string) {
	o.FirstName = &v
}

// GetGlobalPermissions returns the GlobalPermissions field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetGlobalPermissions() GlobalPermissionInfo {
	if o == nil || o.GlobalPermissions == nil {
		var ret GlobalPermissionInfo
		return ret
	}
	return *o.GlobalPermissions
}

// GetGlobalPermissionsOk returns a tuple with the GlobalPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetGlobalPermissionsOk() (*GlobalPermissionInfo, bool) {
	if o == nil || o.GlobalPermissions == nil {
		return nil, false
	}
	return o.GlobalPermissions, true
}

// HasGlobalPermissions returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasGlobalPermissions() bool {
	if o != nil && o.GlobalPermissions != nil {
		return true
	}

	return false
}

// SetGlobalPermissions gets a reference to the given GlobalPermissionInfo and assigns it to the GlobalPermissions field.
func (o *BTUserOAuth2SummaryInfo) SetGlobalPermissions(v GlobalPermissionInfo) {
	o.GlobalPermissions = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTUserOAuth2SummaryInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTUserOAuth2SummaryInfo) SetId(v string) {
	o.Id = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *BTUserOAuth2SummaryInfo) SetImage(v string) {
	o.Image = &v
}

// GetIsGuest returns the IsGuest field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetIsGuest() bool {
	if o == nil || o.IsGuest == nil {
		var ret bool
		return ret
	}
	return *o.IsGuest
}

// GetIsGuestOk returns a tuple with the IsGuest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetIsGuestOk() (*bool, bool) {
	if o == nil || o.IsGuest == nil {
		return nil, false
	}
	return o.IsGuest, true
}

// HasIsGuest returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasIsGuest() bool {
	if o != nil && o.IsGuest != nil {
		return true
	}

	return false
}

// SetIsGuest gets a reference to the given bool and assigns it to the IsGuest field.
func (o *BTUserOAuth2SummaryInfo) SetIsGuest(v bool) {
	o.IsGuest = &v
}

// GetIsLight returns the IsLight field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetIsLight() bool {
	if o == nil || o.IsLight == nil {
		var ret bool
		return ret
	}
	return *o.IsLight
}

// GetIsLightOk returns a tuple with the IsLight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetIsLightOk() (*bool, bool) {
	if o == nil || o.IsLight == nil {
		return nil, false
	}
	return o.IsLight, true
}

// HasIsLight returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasIsLight() bool {
	if o != nil && o.IsLight != nil {
		return true
	}

	return false
}

// SetIsLight gets a reference to the given bool and assigns it to the IsLight field.
func (o *BTUserOAuth2SummaryInfo) SetIsLight(v bool) {
	o.IsLight = &v
}

// GetLastLoginTime returns the LastLoginTime field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetLastLoginTime() JSONTime {
	if o == nil || o.LastLoginTime == nil {
		var ret JSONTime
		return ret
	}
	return *o.LastLoginTime
}

// GetLastLoginTimeOk returns a tuple with the LastLoginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetLastLoginTimeOk() (*JSONTime, bool) {
	if o == nil || o.LastLoginTime == nil {
		return nil, false
	}
	return o.LastLoginTime, true
}

// HasLastLoginTime returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasLastLoginTime() bool {
	if o != nil && o.LastLoginTime != nil {
		return true
	}

	return false
}

// SetLastLoginTime gets a reference to the given JSONTime and assigns it to the LastLoginTime field.
func (o *BTUserOAuth2SummaryInfo) SetLastLoginTime(v JSONTime) {
	o.LastLoginTime = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *BTUserOAuth2SummaryInfo) SetLastName(v string) {
	o.LastName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTUserOAuth2SummaryInfo) SetName(v string) {
	o.Name = &v
}

// GetOauth2Scopes returns the Oauth2Scopes field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetOauth2Scopes() int64 {
	if o == nil || o.Oauth2Scopes == nil {
		var ret int64
		return ret
	}
	return *o.Oauth2Scopes
}

// GetOauth2ScopesOk returns a tuple with the Oauth2Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetOauth2ScopesOk() (*int64, bool) {
	if o == nil || o.Oauth2Scopes == nil {
		return nil, false
	}
	return o.Oauth2Scopes, true
}

// HasOauth2Scopes returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasOauth2Scopes() bool {
	if o != nil && o.Oauth2Scopes != nil {
		return true
	}

	return false
}

// SetOauth2Scopes gets a reference to the given int64 and assigns it to the Oauth2Scopes field.
func (o *BTUserOAuth2SummaryInfo) SetOauth2Scopes(v int64) {
	o.Oauth2Scopes = &v
}

// GetPersonalMessageAllowed returns the PersonalMessageAllowed field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetPersonalMessageAllowed() bool {
	if o == nil || o.PersonalMessageAllowed == nil {
		var ret bool
		return ret
	}
	return *o.PersonalMessageAllowed
}

// GetPersonalMessageAllowedOk returns a tuple with the PersonalMessageAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetPersonalMessageAllowedOk() (*bool, bool) {
	if o == nil || o.PersonalMessageAllowed == nil {
		return nil, false
	}
	return o.PersonalMessageAllowed, true
}

// HasPersonalMessageAllowed returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasPersonalMessageAllowed() bool {
	if o != nil && o.PersonalMessageAllowed != nil {
		return true
	}

	return false
}

// SetPersonalMessageAllowed gets a reference to the given bool and assigns it to the PersonalMessageAllowed field.
func (o *BTUserOAuth2SummaryInfo) SetPersonalMessageAllowed(v bool) {
	o.PersonalMessageAllowed = &v
}

// GetPlanGroup returns the PlanGroup field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetPlanGroup() string {
	if o == nil || o.PlanGroup == nil {
		var ret string
		return ret
	}
	return *o.PlanGroup
}

// GetPlanGroupOk returns a tuple with the PlanGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetPlanGroupOk() (*string, bool) {
	if o == nil || o.PlanGroup == nil {
		return nil, false
	}
	return o.PlanGroup, true
}

// HasPlanGroup returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasPlanGroup() bool {
	if o != nil && o.PlanGroup != nil {
		return true
	}

	return false
}

// SetPlanGroup gets a reference to the given string and assigns it to the PlanGroup field.
func (o *BTUserOAuth2SummaryInfo) SetPlanGroup(v string) {
	o.PlanGroup = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetRole() int32 {
	if o == nil || o.Role == nil {
		var ret int32
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetRoleOk() (*int32, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given int32 and assigns it to the Role field.
func (o *BTUserOAuth2SummaryInfo) SetRole(v int32) {
	o.Role = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetRolesOk() ([]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *BTUserOAuth2SummaryInfo) SetRoles(v []string) {
	o.Roles = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetSource() int32 {
	if o == nil || o.Source == nil {
		var ret int32
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetSourceOk() (*int32, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given int32 and assigns it to the Source field.
func (o *BTUserOAuth2SummaryInfo) SetSource(v int32) {
	o.Source = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *BTUserOAuth2SummaryInfo) SetState(v int32) {
	o.State = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTUserOAuth2SummaryInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserOAuth2SummaryInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTUserOAuth2SummaryInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTUserOAuth2SummaryInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTUserOAuth2SummaryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.CompanyPlan != nil {
		toSerialize["companyPlan"] = o.CompanyPlan
	}
	if o.DocumentationName != nil {
		toSerialize["documentationName"] = o.DocumentationName
	}
	if o.DocumentationNameOverride != nil {
		toSerialize["documentationNameOverride"] = o.DocumentationNameOverride
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.GlobalPermissions != nil {
		toSerialize["globalPermissions"] = o.GlobalPermissions
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.IsGuest != nil {
		toSerialize["isGuest"] = o.IsGuest
	}
	if o.IsLight != nil {
		toSerialize["isLight"] = o.IsLight
	}
	if o.LastLoginTime != nil {
		toSerialize["lastLoginTime"] = o.LastLoginTime
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Oauth2Scopes != nil {
		toSerialize["oauth2Scopes"] = o.Oauth2Scopes
	}
	if o.PersonalMessageAllowed != nil {
		toSerialize["personalMessageAllowed"] = o.PersonalMessageAllowed
	}
	if o.PlanGroup != nil {
		toSerialize["planGroup"] = o.PlanGroup
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTUserOAuth2SummaryInfo struct {
	value *BTUserOAuth2SummaryInfo
	isSet bool
}

func (v NullableBTUserOAuth2SummaryInfo) Get() *BTUserOAuth2SummaryInfo {
	return v.value
}

func (v *NullableBTUserOAuth2SummaryInfo) Set(val *BTUserOAuth2SummaryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUserOAuth2SummaryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUserOAuth2SummaryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUserOAuth2SummaryInfo(val *BTUserOAuth2SummaryInfo) *NullableBTUserOAuth2SummaryInfo {
	return &NullableBTUserOAuth2SummaryInfo{value: val, isSet: true}
}

func (v NullableBTUserOAuth2SummaryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUserOAuth2SummaryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
