/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.157.9170-8cb36f16959d
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDocumentSearchParams struct for BTDocumentSearchParams
type BTDocumentSearchParams struct {
	DocumentFilter *int32  `json:"documentFilter,omitempty"`
	FoundIn        *string `json:"foundIn,omitempty"`
	Limit          *int32  `json:"limit,omitempty"`
	LuceneSyntax   *bool   `json:"luceneSyntax,omitempty"`
	Offset         *int32  `json:"offset,omitempty"`
	OwnerId        *string `json:"ownerId,omitempty"`
	ParentId       *string `json:"parentId,omitempty"`
	RawQuery       *string `json:"rawQuery,omitempty"`
	SortColumn     *string `json:"sortColumn,omitempty"`
	SortOrder      *string `json:"sortOrder,omitempty"`
	Type           *string `json:"type,omitempty"`
	When           *string `json:"when,omitempty"`
}

// NewBTDocumentSearchParams instantiates a new BTDocumentSearchParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentSearchParams() *BTDocumentSearchParams {
	this := BTDocumentSearchParams{}
	return &this
}

// NewBTDocumentSearchParamsWithDefaults instantiates a new BTDocumentSearchParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentSearchParamsWithDefaults() *BTDocumentSearchParams {
	this := BTDocumentSearchParams{}
	return &this
}

// GetDocumentFilter returns the DocumentFilter field value if set, zero value otherwise.
func (o *BTDocumentSearchParams) GetDocumentFilter() int32 {
	if o == nil || o.DocumentFilter == nil {
		var ret int32
		return ret
	}
	return *o.DocumentFilter
}

// GetDocumentFilterOk returns a tuple with the DocumentFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchParams) GetDocumentFilterOk() (*int32, bool) {
	if o == nil || o.DocumentFilter == nil {
		return nil, false
	}
	return o.DocumentFilter, true
}

// HasDocumentFilter returns a boolean if a field has been set.
func (o *BTDocumentSearchParams) HasDocumentFilter() bool {
	if o != nil && o.DocumentFilter != nil {
		return true
	}

	return false
}

// SetDocumentFilter gets a reference to the given int32 and assigns it to the DocumentFilter field.
func (o *BTDocumentSearchParams) SetDocumentFilter(v int32) {
	o.DocumentFilter = &v
}

// GetFoundIn returns the FoundIn field value if set, zero value otherwise.
func (o *BTDocumentSearchParams) GetFoundIn() string {
	if o == nil || o.FoundIn == nil {
		var ret string
		return ret
	}
	return *o.FoundIn
}

// GetFoundInOk returns a tuple with the FoundIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchParams) GetFoundInOk() (*string, bool) {
	if o == nil || o.FoundIn == nil {
		return nil, false
	}
	return o.FoundIn, true
}

// HasFoundIn returns a boolean if a field has been set.
func (o *BTDocumentSearchParams) HasFoundIn() bool {
	if o != nil && o.FoundIn != nil {
		return true
	}

	return false
}

// SetFoundIn gets a reference to the given string and assigns it to the FoundIn field.
func (o *BTDocumentSearchParams) SetFoundIn(v string) {
	o.FoundIn = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *BTDocumentSearchParams) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchParams) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *BTDocumentSearchParams) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *BTDocumentSearchParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetLuceneSyntax returns the LuceneSyntax field value if set, zero value otherwise.
func (o *BTDocumentSearchParams) GetLuceneSyntax() bool {
	if o == nil || o.LuceneSyntax == nil {
		var ret bool
		return ret
	}
	return *o.LuceneSyntax
}

// GetLuceneSyntaxOk returns a tuple with the LuceneSyntax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchParams) GetLuceneSyntaxOk() (*bool, bool) {
	if o == nil || o.LuceneSyntax == nil {
		return nil, false
	}
	return o.LuceneSyntax, true
}

// HasLuceneSyntax returns a boolean if a field has been set.
func (o *BTDocumentSearchParams) HasLuceneSyntax() bool {
	if o != nil && o.LuceneSyntax != nil {
		return true
	}

	return false
}

// SetLuceneSyntax gets a reference to the given bool and assigns it to the LuceneSyntax field.
func (o *BTDocumentSearchParams) SetLuceneSyntax(v bool) {
	o.LuceneSyntax = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *BTDocumentSearchParams) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchParams) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *BTDocumentSearchParams) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *BTDocumentSearchParams) SetOffset(v int32) {
	o.Offset = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *BTDocumentSearchParams) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchParams) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *BTDocumentSearchParams) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *BTDocumentSearchParams) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *BTDocumentSearchParams) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchParams) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *BTDocumentSearchParams) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *BTDocumentSearchParams) SetParentId(v string) {
	o.ParentId = &v
}

// GetRawQuery returns the RawQuery field value if set, zero value otherwise.
func (o *BTDocumentSearchParams) GetRawQuery() string {
	if o == nil || o.RawQuery == nil {
		var ret string
		return ret
	}
	return *o.RawQuery
}

// GetRawQueryOk returns a tuple with the RawQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchParams) GetRawQueryOk() (*string, bool) {
	if o == nil || o.RawQuery == nil {
		return nil, false
	}
	return o.RawQuery, true
}

// HasRawQuery returns a boolean if a field has been set.
func (o *BTDocumentSearchParams) HasRawQuery() bool {
	if o != nil && o.RawQuery != nil {
		return true
	}

	return false
}

// SetRawQuery gets a reference to the given string and assigns it to the RawQuery field.
func (o *BTDocumentSearchParams) SetRawQuery(v string) {
	o.RawQuery = &v
}

// GetSortColumn returns the SortColumn field value if set, zero value otherwise.
func (o *BTDocumentSearchParams) GetSortColumn() string {
	if o == nil || o.SortColumn == nil {
		var ret string
		return ret
	}
	return *o.SortColumn
}

// GetSortColumnOk returns a tuple with the SortColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchParams) GetSortColumnOk() (*string, bool) {
	if o == nil || o.SortColumn == nil {
		return nil, false
	}
	return o.SortColumn, true
}

// HasSortColumn returns a boolean if a field has been set.
func (o *BTDocumentSearchParams) HasSortColumn() bool {
	if o != nil && o.SortColumn != nil {
		return true
	}

	return false
}

// SetSortColumn gets a reference to the given string and assigns it to the SortColumn field.
func (o *BTDocumentSearchParams) SetSortColumn(v string) {
	o.SortColumn = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *BTDocumentSearchParams) GetSortOrder() string {
	if o == nil || o.SortOrder == nil {
		var ret string
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchParams) GetSortOrderOk() (*string, bool) {
	if o == nil || o.SortOrder == nil {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *BTDocumentSearchParams) HasSortOrder() bool {
	if o != nil && o.SortOrder != nil {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given string and assigns it to the SortOrder field.
func (o *BTDocumentSearchParams) SetSortOrder(v string) {
	o.SortOrder = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTDocumentSearchParams) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchParams) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTDocumentSearchParams) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTDocumentSearchParams) SetType(v string) {
	o.Type = &v
}

// GetWhen returns the When field value if set, zero value otherwise.
func (o *BTDocumentSearchParams) GetWhen() string {
	if o == nil || o.When == nil {
		var ret string
		return ret
	}
	return *o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSearchParams) GetWhenOk() (*string, bool) {
	if o == nil || o.When == nil {
		return nil, false
	}
	return o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *BTDocumentSearchParams) HasWhen() bool {
	if o != nil && o.When != nil {
		return true
	}

	return false
}

// SetWhen gets a reference to the given string and assigns it to the When field.
func (o *BTDocumentSearchParams) SetWhen(v string) {
	o.When = &v
}

func (o BTDocumentSearchParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentFilter != nil {
		toSerialize["documentFilter"] = o.DocumentFilter
	}
	if o.FoundIn != nil {
		toSerialize["foundIn"] = o.FoundIn
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.LuceneSyntax != nil {
		toSerialize["luceneSyntax"] = o.LuceneSyntax
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.RawQuery != nil {
		toSerialize["rawQuery"] = o.RawQuery
	}
	if o.SortColumn != nil {
		toSerialize["sortColumn"] = o.SortColumn
	}
	if o.SortOrder != nil {
		toSerialize["sortOrder"] = o.SortOrder
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.When != nil {
		toSerialize["when"] = o.When
	}
	return json.Marshal(toSerialize)
}

type NullableBTDocumentSearchParams struct {
	value *BTDocumentSearchParams
	isSet bool
}

func (v NullableBTDocumentSearchParams) Get() *BTDocumentSearchParams {
	return v.value
}

func (v *NullableBTDocumentSearchParams) Set(val *BTDocumentSearchParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentSearchParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentSearchParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentSearchParams(val *BTDocumentSearchParams) *NullableBTDocumentSearchParams {
	return &NullableBTDocumentSearchParams{value: val, isSet: true}
}

func (v NullableBTDocumentSearchParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentSearchParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
