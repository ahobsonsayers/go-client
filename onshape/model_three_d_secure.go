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

// ThreeDSecure struct for ThreeDSecure
type ThreeDSecure struct {
	Amount        *int64  `json:"amount,omitempty"`
	Authenticated *bool   `json:"authenticated,omitempty"`
	Card          *Card   `json:"card,omitempty"`
	Created       *int64  `json:"created,omitempty"`
	Currency      *string `json:"currency,omitempty"`
	Id            *string `json:"id,omitempty"`
	Livemode      *bool   `json:"livemode,omitempty"`
	Object        *string `json:"object,omitempty"`
	RedirectURL   *string `json:"redirectURL,omitempty"`
	Status        *string `json:"status,omitempty"`
}

// NewThreeDSecure instantiates a new ThreeDSecure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDSecure() *ThreeDSecure {
	this := ThreeDSecure{}
	return &this
}

// NewThreeDSecureWithDefaults instantiates a new ThreeDSecure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDSecureWithDefaults() *ThreeDSecure {
	this := ThreeDSecure{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ThreeDSecure) GetAmount() int64 {
	if o == nil || o.Amount == nil {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecure) GetAmountOk() (*int64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ThreeDSecure) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *ThreeDSecure) SetAmount(v int64) {
	o.Amount = &v
}

// GetAuthenticated returns the Authenticated field value if set, zero value otherwise.
func (o *ThreeDSecure) GetAuthenticated() bool {
	if o == nil || o.Authenticated == nil {
		var ret bool
		return ret
	}
	return *o.Authenticated
}

// GetAuthenticatedOk returns a tuple with the Authenticated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecure) GetAuthenticatedOk() (*bool, bool) {
	if o == nil || o.Authenticated == nil {
		return nil, false
	}
	return o.Authenticated, true
}

// HasAuthenticated returns a boolean if a field has been set.
func (o *ThreeDSecure) HasAuthenticated() bool {
	if o != nil && o.Authenticated != nil {
		return true
	}

	return false
}

// SetAuthenticated gets a reference to the given bool and assigns it to the Authenticated field.
func (o *ThreeDSecure) SetAuthenticated(v bool) {
	o.Authenticated = &v
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *ThreeDSecure) GetCard() Card {
	if o == nil || o.Card == nil {
		var ret Card
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecure) GetCardOk() (*Card, bool) {
	if o == nil || o.Card == nil {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *ThreeDSecure) HasCard() bool {
	if o != nil && o.Card != nil {
		return true
	}

	return false
}

// SetCard gets a reference to the given Card and assigns it to the Card field.
func (o *ThreeDSecure) SetCard(v Card) {
	o.Card = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ThreeDSecure) GetCreated() int64 {
	if o == nil || o.Created == nil {
		var ret int64
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecure) GetCreatedOk() (*int64, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ThreeDSecure) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int64 and assigns it to the Created field.
func (o *ThreeDSecure) SetCreated(v int64) {
	o.Created = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ThreeDSecure) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecure) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ThreeDSecure) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ThreeDSecure) SetCurrency(v string) {
	o.Currency = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ThreeDSecure) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecure) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ThreeDSecure) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ThreeDSecure) SetId(v string) {
	o.Id = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *ThreeDSecure) GetLivemode() bool {
	if o == nil || o.Livemode == nil {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecure) GetLivemodeOk() (*bool, bool) {
	if o == nil || o.Livemode == nil {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *ThreeDSecure) HasLivemode() bool {
	if o != nil && o.Livemode != nil {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *ThreeDSecure) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ThreeDSecure) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecure) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ThreeDSecure) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ThreeDSecure) SetObject(v string) {
	o.Object = &v
}

// GetRedirectURL returns the RedirectURL field value if set, zero value otherwise.
func (o *ThreeDSecure) GetRedirectURL() string {
	if o == nil || o.RedirectURL == nil {
		var ret string
		return ret
	}
	return *o.RedirectURL
}

// GetRedirectURLOk returns a tuple with the RedirectURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecure) GetRedirectURLOk() (*string, bool) {
	if o == nil || o.RedirectURL == nil {
		return nil, false
	}
	return o.RedirectURL, true
}

// HasRedirectURL returns a boolean if a field has been set.
func (o *ThreeDSecure) HasRedirectURL() bool {
	if o != nil && o.RedirectURL != nil {
		return true
	}

	return false
}

// SetRedirectURL gets a reference to the given string and assigns it to the RedirectURL field.
func (o *ThreeDSecure) SetRedirectURL(v string) {
	o.RedirectURL = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ThreeDSecure) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecure) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ThreeDSecure) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ThreeDSecure) SetStatus(v string) {
	o.Status = &v
}

func (o ThreeDSecure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Authenticated != nil {
		toSerialize["authenticated"] = o.Authenticated
	}
	if o.Card != nil {
		toSerialize["card"] = o.Card
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Livemode != nil {
		toSerialize["livemode"] = o.Livemode
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.RedirectURL != nil {
		toSerialize["redirectURL"] = o.RedirectURL
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableThreeDSecure struct {
	value *ThreeDSecure
	isSet bool
}

func (v NullableThreeDSecure) Get() *ThreeDSecure {
	return v.value
}

func (v *NullableThreeDSecure) Set(val *ThreeDSecure) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDSecure) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDSecure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDSecure(val *ThreeDSecure) *NullableThreeDSecure {
	return &NullableThreeDSecure{value: val, isSet: true}
}

func (v NullableThreeDSecure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDSecure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
