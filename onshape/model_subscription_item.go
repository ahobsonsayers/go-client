/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6797-6824f8df9946
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// SubscriptionItem struct for SubscriptionItem
type SubscriptionItem struct {
	Created  *int64  `json:"created,omitempty"`
	Id       *string `json:"id,omitempty"`
	Object   *string `json:"object,omitempty"`
	Plan     *Plan   `json:"plan,omitempty"`
	Quantity *int32  `json:"quantity,omitempty"`
}

// NewSubscriptionItem instantiates a new SubscriptionItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionItem() *SubscriptionItem {
	this := SubscriptionItem{}
	return &this
}

// NewSubscriptionItemWithDefaults instantiates a new SubscriptionItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionItemWithDefaults() *SubscriptionItem {
	this := SubscriptionItem{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SubscriptionItem) GetCreated() int64 {
	if o == nil || o.Created == nil {
		var ret int64
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionItem) GetCreatedOk() (*int64, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SubscriptionItem) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int64 and assigns it to the Created field.
func (o *SubscriptionItem) SetCreated(v int64) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriptionItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubscriptionItem) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *SubscriptionItem) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionItem) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *SubscriptionItem) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *SubscriptionItem) SetObject(v string) {
	o.Object = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *SubscriptionItem) GetPlan() Plan {
	if o == nil || o.Plan == nil {
		var ret Plan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionItem) GetPlanOk() (*Plan, bool) {
	if o == nil || o.Plan == nil {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *SubscriptionItem) HasPlan() bool {
	if o != nil && o.Plan != nil {
		return true
	}

	return false
}

// SetPlan gets a reference to the given Plan and assigns it to the Plan field.
func (o *SubscriptionItem) SetPlan(v Plan) {
	o.Plan = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *SubscriptionItem) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionItem) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *SubscriptionItem) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *SubscriptionItem) SetQuantity(v int32) {
	o.Quantity = &v
}

func (o SubscriptionItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Plan != nil {
		toSerialize["plan"] = o.Plan
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionItem struct {
	value *SubscriptionItem
	isSet bool
}

func (v NullableSubscriptionItem) Get() *SubscriptionItem {
	return v.value
}

func (v *NullableSubscriptionItem) Set(val *SubscriptionItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionItem(val *SubscriptionItem) *NullableSubscriptionItem {
	return &NullableSubscriptionItem{value: val, isSet: true}
}

func (v NullableSubscriptionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
