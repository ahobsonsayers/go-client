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

// BTCompanyInfo struct for BTCompanyInfo
type BTCompanyInfo struct {
	Address           *BTAddressInfo `json:"address,omitempty"`
	Admin             *bool          `json:"admin,omitempty"`
	Description       *string        `json:"description,omitempty"`
	DomainPrefix      *string        `json:"domainPrefix,omitempty"`
	EnterpriseBaseUrl *string        `json:"enterpriseBaseUrl,omitempty"`
	EnterpriseSubtype *int32         `json:"enterpriseSubtype,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id    *string `json:"id,omitempty"`
	Image *string `json:"image,omitempty"`
	// Name of the resource.
	Name                    *string         `json:"name,omitempty"`
	NoPublicDocuments       *bool           `json:"noPublicDocuments,omitempty"`
	OwnerId                 *string         `json:"ownerId,omitempty"`
	Purchase                *BTPurchaseInfo `json:"purchase,omitempty"`
	SecondaryDomainPrefixes []string        `json:"secondaryDomainPrefixes,omitempty"`
	State                   *int32          `json:"state,omitempty"`
	Type                    *int32          `json:"type,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef *string `json:"viewRef,omitempty"`
}

// NewBTCompanyInfo instantiates a new BTCompanyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCompanyInfo() *BTCompanyInfo {
	this := BTCompanyInfo{}
	return &this
}

// NewBTCompanyInfoWithDefaults instantiates a new BTCompanyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCompanyInfoWithDefaults() *BTCompanyInfo {
	this := BTCompanyInfo{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *BTCompanyInfo) GetAddress() BTAddressInfo {
	if o == nil || o.Address == nil {
		var ret BTAddressInfo
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyInfo) GetAddressOk() (*BTAddressInfo, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *BTCompanyInfo) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given BTAddressInfo and assigns it to the Address field.
func (o *BTCompanyInfo) SetAddress(v BTAddressInfo) {
	o.Address = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *BTCompanyInfo) GetAdmin() bool {
	if o == nil || o.Admin == nil {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyInfo) GetAdminOk() (*bool, bool) {
	if o == nil || o.Admin == nil {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *BTCompanyInfo) HasAdmin() bool {
	if o != nil && o.Admin != nil {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *BTCompanyInfo) SetAdmin(v bool) {
	o.Admin = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTCompanyInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTCompanyInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTCompanyInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDomainPrefix returns the DomainPrefix field value if set, zero value otherwise.
func (o *BTCompanyInfo) GetDomainPrefix() string {
	if o == nil || o.DomainPrefix == nil {
		var ret string
		return ret
	}
	return *o.DomainPrefix
}

// GetDomainPrefixOk returns a tuple with the DomainPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyInfo) GetDomainPrefixOk() (*string, bool) {
	if o == nil || o.DomainPrefix == nil {
		return nil, false
	}
	return o.DomainPrefix, true
}

// HasDomainPrefix returns a boolean if a field has been set.
func (o *BTCompanyInfo) HasDomainPrefix() bool {
	if o != nil && o.DomainPrefix != nil {
		return true
	}

	return false
}

// SetDomainPrefix gets a reference to the given string and assigns it to the DomainPrefix field.
func (o *BTCompanyInfo) SetDomainPrefix(v string) {
	o.DomainPrefix = &v
}

// GetEnterpriseBaseUrl returns the EnterpriseBaseUrl field value if set, zero value otherwise.
func (o *BTCompanyInfo) GetEnterpriseBaseUrl() string {
	if o == nil || o.EnterpriseBaseUrl == nil {
		var ret string
		return ret
	}
	return *o.EnterpriseBaseUrl
}

// GetEnterpriseBaseUrlOk returns a tuple with the EnterpriseBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyInfo) GetEnterpriseBaseUrlOk() (*string, bool) {
	if o == nil || o.EnterpriseBaseUrl == nil {
		return nil, false
	}
	return o.EnterpriseBaseUrl, true
}

// HasEnterpriseBaseUrl returns a boolean if a field has been set.
func (o *BTCompanyInfo) HasEnterpriseBaseUrl() bool {
	if o != nil && o.EnterpriseBaseUrl != nil {
		return true
	}

	return false
}

// SetEnterpriseBaseUrl gets a reference to the given string and assigns it to the EnterpriseBaseUrl field.
func (o *BTCompanyInfo) SetEnterpriseBaseUrl(v string) {
	o.EnterpriseBaseUrl = &v
}

// GetEnterpriseSubtype returns the EnterpriseSubtype field value if set, zero value otherwise.
func (o *BTCompanyInfo) GetEnterpriseSubtype() int32 {
	if o == nil || o.EnterpriseSubtype == nil {
		var ret int32
		return ret
	}
	return *o.EnterpriseSubtype
}

// GetEnterpriseSubtypeOk returns a tuple with the EnterpriseSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyInfo) GetEnterpriseSubtypeOk() (*int32, bool) {
	if o == nil || o.EnterpriseSubtype == nil {
		return nil, false
	}
	return o.EnterpriseSubtype, true
}

// HasEnterpriseSubtype returns a boolean if a field has been set.
func (o *BTCompanyInfo) HasEnterpriseSubtype() bool {
	if o != nil && o.EnterpriseSubtype != nil {
		return true
	}

	return false
}

// SetEnterpriseSubtype gets a reference to the given int32 and assigns it to the EnterpriseSubtype field.
func (o *BTCompanyInfo) SetEnterpriseSubtype(v int32) {
	o.EnterpriseSubtype = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTCompanyInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTCompanyInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTCompanyInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTCompanyInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTCompanyInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTCompanyInfo) SetId(v string) {
	o.Id = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BTCompanyInfo) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyInfo) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *BTCompanyInfo) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *BTCompanyInfo) SetImage(v string) {
	o.Image = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTCompanyInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTCompanyInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTCompanyInfo) SetName(v string) {
	o.Name = &v
}

// GetNoPublicDocuments returns the NoPublicDocuments field value if set, zero value otherwise.
func (o *BTCompanyInfo) GetNoPublicDocuments() bool {
	if o == nil || o.NoPublicDocuments == nil {
		var ret bool
		return ret
	}
	return *o.NoPublicDocuments
}

// GetNoPublicDocumentsOk returns a tuple with the NoPublicDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyInfo) GetNoPublicDocumentsOk() (*bool, bool) {
	if o == nil || o.NoPublicDocuments == nil {
		return nil, false
	}
	return o.NoPublicDocuments, true
}

// HasNoPublicDocuments returns a boolean if a field has been set.
func (o *BTCompanyInfo) HasNoPublicDocuments() bool {
	if o != nil && o.NoPublicDocuments != nil {
		return true
	}

	return false
}

// SetNoPublicDocuments gets a reference to the given bool and assigns it to the NoPublicDocuments field.
func (o *BTCompanyInfo) SetNoPublicDocuments(v bool) {
	o.NoPublicDocuments = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *BTCompanyInfo) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyInfo) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *BTCompanyInfo) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *BTCompanyInfo) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetPurchase returns the Purchase field value if set, zero value otherwise.
func (o *BTCompanyInfo) GetPurchase() BTPurchaseInfo {
	if o == nil || o.Purchase == nil {
		var ret BTPurchaseInfo
		return ret
	}
	return *o.Purchase
}

// GetPurchaseOk returns a tuple with the Purchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyInfo) GetPurchaseOk() (*BTPurchaseInfo, bool) {
	if o == nil || o.Purchase == nil {
		return nil, false
	}
	return o.Purchase, true
}

// HasPurchase returns a boolean if a field has been set.
func (o *BTCompanyInfo) HasPurchase() bool {
	if o != nil && o.Purchase != nil {
		return true
	}

	return false
}

// SetPurchase gets a reference to the given BTPurchaseInfo and assigns it to the Purchase field.
func (o *BTCompanyInfo) SetPurchase(v BTPurchaseInfo) {
	o.Purchase = &v
}

// GetSecondaryDomainPrefixes returns the SecondaryDomainPrefixes field value if set, zero value otherwise.
func (o *BTCompanyInfo) GetSecondaryDomainPrefixes() []string {
	if o == nil || o.SecondaryDomainPrefixes == nil {
		var ret []string
		return ret
	}
	return o.SecondaryDomainPrefixes
}

// GetSecondaryDomainPrefixesOk returns a tuple with the SecondaryDomainPrefixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyInfo) GetSecondaryDomainPrefixesOk() ([]string, bool) {
	if o == nil || o.SecondaryDomainPrefixes == nil {
		return nil, false
	}
	return o.SecondaryDomainPrefixes, true
}

// HasSecondaryDomainPrefixes returns a boolean if a field has been set.
func (o *BTCompanyInfo) HasSecondaryDomainPrefixes() bool {
	if o != nil && o.SecondaryDomainPrefixes != nil {
		return true
	}

	return false
}

// SetSecondaryDomainPrefixes gets a reference to the given []string and assigns it to the SecondaryDomainPrefixes field.
func (o *BTCompanyInfo) SetSecondaryDomainPrefixes(v []string) {
	o.SecondaryDomainPrefixes = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BTCompanyInfo) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyInfo) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BTCompanyInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *BTCompanyInfo) SetState(v int32) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTCompanyInfo) GetType() int32 {
	if o == nil || o.Type == nil {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyInfo) GetTypeOk() (*int32, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTCompanyInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *BTCompanyInfo) SetType(v int32) {
	o.Type = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTCompanyInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTCompanyInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTCompanyInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTCompanyInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Admin != nil {
		toSerialize["admin"] = o.Admin
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DomainPrefix != nil {
		toSerialize["domainPrefix"] = o.DomainPrefix
	}
	if o.EnterpriseBaseUrl != nil {
		toSerialize["enterpriseBaseUrl"] = o.EnterpriseBaseUrl
	}
	if o.EnterpriseSubtype != nil {
		toSerialize["enterpriseSubtype"] = o.EnterpriseSubtype
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NoPublicDocuments != nil {
		toSerialize["noPublicDocuments"] = o.NoPublicDocuments
	}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.Purchase != nil {
		toSerialize["purchase"] = o.Purchase
	}
	if o.SecondaryDomainPrefixes != nil {
		toSerialize["secondaryDomainPrefixes"] = o.SecondaryDomainPrefixes
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTCompanyInfo struct {
	value *BTCompanyInfo
	isSet bool
}

func (v NullableBTCompanyInfo) Get() *BTCompanyInfo {
	return v.value
}

func (v *NullableBTCompanyInfo) Set(val *BTCompanyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCompanyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCompanyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCompanyInfo(val *BTCompanyInfo) *NullableBTCompanyInfo {
	return &NullableBTCompanyInfo{value: val, isSet: true}
}

func (v NullableBTCompanyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCompanyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
