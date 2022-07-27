/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5726-0a2a2c07c0aa
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTOwner3114 struct for BTOwner3114
type BTOwner3114 struct {
	BtType           *string `json:"btType,omitempty"`
	CompanyId        *bool   `json:"companyId,omitempty"`
	OwnerId          *string `json:"ownerId,omitempty"`
	OwnerType        *string `json:"ownerType,omitempty"`
	OwnerTypeOrdinal *int32  `json:"ownerTypeOrdinal,omitempty"`
	UserId           *bool   `json:"userId,omitempty"`
}

// NewBTOwner3114 instantiates a new BTOwner3114 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTOwner3114() *BTOwner3114 {
	this := BTOwner3114{}
	return &this
}

// NewBTOwner3114WithDefaults instantiates a new BTOwner3114 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTOwner3114WithDefaults() *BTOwner3114 {
	this := BTOwner3114{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTOwner3114) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOwner3114) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTOwner3114) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTOwner3114) SetBtType(v string) {
	o.BtType = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTOwner3114) GetCompanyId() bool {
	if o == nil || o.CompanyId == nil {
		var ret bool
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOwner3114) GetCompanyIdOk() (*bool, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTOwner3114) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given bool and assigns it to the CompanyId field.
func (o *BTOwner3114) SetCompanyId(v bool) {
	o.CompanyId = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *BTOwner3114) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOwner3114) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *BTOwner3114) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *BTOwner3114) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *BTOwner3114) GetOwnerType() string {
	if o == nil || o.OwnerType == nil {
		var ret string
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOwner3114) GetOwnerTypeOk() (*string, bool) {
	if o == nil || o.OwnerType == nil {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *BTOwner3114) HasOwnerType() bool {
	if o != nil && o.OwnerType != nil {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given string and assigns it to the OwnerType field.
func (o *BTOwner3114) SetOwnerType(v string) {
	o.OwnerType = &v
}

// GetOwnerTypeOrdinal returns the OwnerTypeOrdinal field value if set, zero value otherwise.
func (o *BTOwner3114) GetOwnerTypeOrdinal() int32 {
	if o == nil || o.OwnerTypeOrdinal == nil {
		var ret int32
		return ret
	}
	return *o.OwnerTypeOrdinal
}

// GetOwnerTypeOrdinalOk returns a tuple with the OwnerTypeOrdinal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOwner3114) GetOwnerTypeOrdinalOk() (*int32, bool) {
	if o == nil || o.OwnerTypeOrdinal == nil {
		return nil, false
	}
	return o.OwnerTypeOrdinal, true
}

// HasOwnerTypeOrdinal returns a boolean if a field has been set.
func (o *BTOwner3114) HasOwnerTypeOrdinal() bool {
	if o != nil && o.OwnerTypeOrdinal != nil {
		return true
	}

	return false
}

// SetOwnerTypeOrdinal gets a reference to the given int32 and assigns it to the OwnerTypeOrdinal field.
func (o *BTOwner3114) SetOwnerTypeOrdinal(v int32) {
	o.OwnerTypeOrdinal = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BTOwner3114) GetUserId() bool {
	if o == nil || o.UserId == nil {
		var ret bool
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOwner3114) GetUserIdOk() (*bool, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BTOwner3114) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given bool and assigns it to the UserId field.
func (o *BTOwner3114) SetUserId(v bool) {
	o.UserId = &v
}

func (o BTOwner3114) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.OwnerType != nil {
		toSerialize["ownerType"] = o.OwnerType
	}
	if o.OwnerTypeOrdinal != nil {
		toSerialize["ownerTypeOrdinal"] = o.OwnerTypeOrdinal
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableBTOwner3114 struct {
	value *BTOwner3114
	isSet bool
}

func (v NullableBTOwner3114) Get() *BTOwner3114 {
	return v.value
}

func (v *NullableBTOwner3114) Set(val *BTOwner3114) {
	v.value = val
	v.isSet = true
}

func (v NullableBTOwner3114) IsSet() bool {
	return v.isSet
}

func (v *NullableBTOwner3114) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTOwner3114(val *BTOwner3114) *NullableBTOwner3114 {
	return &NullableBTOwner3114{value: val, isSet: true}
}

func (v NullableBTOwner3114) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTOwner3114) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
