/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7180-fb454452a4fd
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// Customer struct for Customer
type Customer struct {
	AccountBalance      *int64                          `json:"accountBalance,omitempty"`
	BusinessVatId       *string                         `json:"businessVatId,omitempty"`
	Cards               *CustomerCardCollection         `json:"cards,omitempty"`
	Created             *int64                          `json:"created,omitempty"`
	Currency            *string                         `json:"currency,omitempty"`
	DefaultCard         *string                         `json:"defaultCard,omitempty"`
	DefaultSource       *string                         `json:"defaultSource,omitempty"`
	DefaultSourceObject *ExternalAccount                `json:"defaultSourceObject,omitempty"`
	Deleted             *bool                           `json:"deleted,omitempty"`
	Delinquent          *bool                           `json:"delinquent,omitempty"`
	Description         *string                         `json:"description,omitempty"`
	Discount            *Discount                       `json:"discount,omitempty"`
	Email               *string                         `json:"email,omitempty"`
	Id                  *string                         `json:"id,omitempty"`
	Livemode            *bool                           `json:"livemode,omitempty"`
	Metadata            *map[string]string              `json:"metadata,omitempty"`
	NextRecurringCharge *NextRecurringCharge            `json:"nextRecurringCharge,omitempty"`
	Object              *string                         `json:"object,omitempty"`
	Shipping            *ShippingDetails                `json:"shipping,omitempty"`
	Sources             *ExternalAccountCollection      `json:"sources,omitempty"`
	Subscription        *Subscription                   `json:"subscription,omitempty"`
	Subscriptions       *CustomerSubscriptionCollection `json:"subscriptions,omitempty"`
	TrialEnd            *int64                          `json:"trialEnd,omitempty"`
}

// NewCustomer instantiates a new Customer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomer() *Customer {
	this := Customer{}
	return &this
}

// NewCustomerWithDefaults instantiates a new Customer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerWithDefaults() *Customer {
	this := Customer{}
	return &this
}

// GetAccountBalance returns the AccountBalance field value if set, zero value otherwise.
func (o *Customer) GetAccountBalance() int64 {
	if o == nil || o.AccountBalance == nil {
		var ret int64
		return ret
	}
	return *o.AccountBalance
}

// GetAccountBalanceOk returns a tuple with the AccountBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetAccountBalanceOk() (*int64, bool) {
	if o == nil || o.AccountBalance == nil {
		return nil, false
	}
	return o.AccountBalance, true
}

// HasAccountBalance returns a boolean if a field has been set.
func (o *Customer) HasAccountBalance() bool {
	if o != nil && o.AccountBalance != nil {
		return true
	}

	return false
}

// SetAccountBalance gets a reference to the given int64 and assigns it to the AccountBalance field.
func (o *Customer) SetAccountBalance(v int64) {
	o.AccountBalance = &v
}

// GetBusinessVatId returns the BusinessVatId field value if set, zero value otherwise.
func (o *Customer) GetBusinessVatId() string {
	if o == nil || o.BusinessVatId == nil {
		var ret string
		return ret
	}
	return *o.BusinessVatId
}

// GetBusinessVatIdOk returns a tuple with the BusinessVatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetBusinessVatIdOk() (*string, bool) {
	if o == nil || o.BusinessVatId == nil {
		return nil, false
	}
	return o.BusinessVatId, true
}

// HasBusinessVatId returns a boolean if a field has been set.
func (o *Customer) HasBusinessVatId() bool {
	if o != nil && o.BusinessVatId != nil {
		return true
	}

	return false
}

// SetBusinessVatId gets a reference to the given string and assigns it to the BusinessVatId field.
func (o *Customer) SetBusinessVatId(v string) {
	o.BusinessVatId = &v
}

// GetCards returns the Cards field value if set, zero value otherwise.
func (o *Customer) GetCards() CustomerCardCollection {
	if o == nil || o.Cards == nil {
		var ret CustomerCardCollection
		return ret
	}
	return *o.Cards
}

// GetCardsOk returns a tuple with the Cards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetCardsOk() (*CustomerCardCollection, bool) {
	if o == nil || o.Cards == nil {
		return nil, false
	}
	return o.Cards, true
}

// HasCards returns a boolean if a field has been set.
func (o *Customer) HasCards() bool {
	if o != nil && o.Cards != nil {
		return true
	}

	return false
}

// SetCards gets a reference to the given CustomerCardCollection and assigns it to the Cards field.
func (o *Customer) SetCards(v CustomerCardCollection) {
	o.Cards = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Customer) GetCreated() int64 {
	if o == nil || o.Created == nil {
		var ret int64
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetCreatedOk() (*int64, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Customer) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int64 and assigns it to the Created field.
func (o *Customer) SetCreated(v int64) {
	o.Created = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Customer) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Customer) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Customer) SetCurrency(v string) {
	o.Currency = &v
}

// GetDefaultCard returns the DefaultCard field value if set, zero value otherwise.
func (o *Customer) GetDefaultCard() string {
	if o == nil || o.DefaultCard == nil {
		var ret string
		return ret
	}
	return *o.DefaultCard
}

// GetDefaultCardOk returns a tuple with the DefaultCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetDefaultCardOk() (*string, bool) {
	if o == nil || o.DefaultCard == nil {
		return nil, false
	}
	return o.DefaultCard, true
}

// HasDefaultCard returns a boolean if a field has been set.
func (o *Customer) HasDefaultCard() bool {
	if o != nil && o.DefaultCard != nil {
		return true
	}

	return false
}

// SetDefaultCard gets a reference to the given string and assigns it to the DefaultCard field.
func (o *Customer) SetDefaultCard(v string) {
	o.DefaultCard = &v
}

// GetDefaultSource returns the DefaultSource field value if set, zero value otherwise.
func (o *Customer) GetDefaultSource() string {
	if o == nil || o.DefaultSource == nil {
		var ret string
		return ret
	}
	return *o.DefaultSource
}

// GetDefaultSourceOk returns a tuple with the DefaultSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetDefaultSourceOk() (*string, bool) {
	if o == nil || o.DefaultSource == nil {
		return nil, false
	}
	return o.DefaultSource, true
}

// HasDefaultSource returns a boolean if a field has been set.
func (o *Customer) HasDefaultSource() bool {
	if o != nil && o.DefaultSource != nil {
		return true
	}

	return false
}

// SetDefaultSource gets a reference to the given string and assigns it to the DefaultSource field.
func (o *Customer) SetDefaultSource(v string) {
	o.DefaultSource = &v
}

// GetDefaultSourceObject returns the DefaultSourceObject field value if set, zero value otherwise.
func (o *Customer) GetDefaultSourceObject() ExternalAccount {
	if o == nil || o.DefaultSourceObject == nil {
		var ret ExternalAccount
		return ret
	}
	return *o.DefaultSourceObject
}

// GetDefaultSourceObjectOk returns a tuple with the DefaultSourceObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetDefaultSourceObjectOk() (*ExternalAccount, bool) {
	if o == nil || o.DefaultSourceObject == nil {
		return nil, false
	}
	return o.DefaultSourceObject, true
}

// HasDefaultSourceObject returns a boolean if a field has been set.
func (o *Customer) HasDefaultSourceObject() bool {
	if o != nil && o.DefaultSourceObject != nil {
		return true
	}

	return false
}

// SetDefaultSourceObject gets a reference to the given ExternalAccount and assigns it to the DefaultSourceObject field.
func (o *Customer) SetDefaultSourceObject(v ExternalAccount) {
	o.DefaultSourceObject = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Customer) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Customer) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *Customer) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetDelinquent returns the Delinquent field value if set, zero value otherwise.
func (o *Customer) GetDelinquent() bool {
	if o == nil || o.Delinquent == nil {
		var ret bool
		return ret
	}
	return *o.Delinquent
}

// GetDelinquentOk returns a tuple with the Delinquent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetDelinquentOk() (*bool, bool) {
	if o == nil || o.Delinquent == nil {
		return nil, false
	}
	return o.Delinquent, true
}

// HasDelinquent returns a boolean if a field has been set.
func (o *Customer) HasDelinquent() bool {
	if o != nil && o.Delinquent != nil {
		return true
	}

	return false
}

// SetDelinquent gets a reference to the given bool and assigns it to the Delinquent field.
func (o *Customer) SetDelinquent(v bool) {
	o.Delinquent = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Customer) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Customer) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Customer) SetDescription(v string) {
	o.Description = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *Customer) GetDiscount() Discount {
	if o == nil || o.Discount == nil {
		var ret Discount
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetDiscountOk() (*Discount, bool) {
	if o == nil || o.Discount == nil {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *Customer) HasDiscount() bool {
	if o != nil && o.Discount != nil {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given Discount and assigns it to the Discount field.
func (o *Customer) SetDiscount(v Discount) {
	o.Discount = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Customer) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Customer) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Customer) SetEmail(v string) {
	o.Email = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Customer) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Customer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Customer) SetId(v string) {
	o.Id = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *Customer) GetLivemode() bool {
	if o == nil || o.Livemode == nil {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetLivemodeOk() (*bool, bool) {
	if o == nil || o.Livemode == nil {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *Customer) HasLivemode() bool {
	if o != nil && o.Livemode != nil {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *Customer) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Customer) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Customer) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *Customer) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetNextRecurringCharge returns the NextRecurringCharge field value if set, zero value otherwise.
func (o *Customer) GetNextRecurringCharge() NextRecurringCharge {
	if o == nil || o.NextRecurringCharge == nil {
		var ret NextRecurringCharge
		return ret
	}
	return *o.NextRecurringCharge
}

// GetNextRecurringChargeOk returns a tuple with the NextRecurringCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetNextRecurringChargeOk() (*NextRecurringCharge, bool) {
	if o == nil || o.NextRecurringCharge == nil {
		return nil, false
	}
	return o.NextRecurringCharge, true
}

// HasNextRecurringCharge returns a boolean if a field has been set.
func (o *Customer) HasNextRecurringCharge() bool {
	if o != nil && o.NextRecurringCharge != nil {
		return true
	}

	return false
}

// SetNextRecurringCharge gets a reference to the given NextRecurringCharge and assigns it to the NextRecurringCharge field.
func (o *Customer) SetNextRecurringCharge(v NextRecurringCharge) {
	o.NextRecurringCharge = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *Customer) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *Customer) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *Customer) SetObject(v string) {
	o.Object = &v
}

// GetShipping returns the Shipping field value if set, zero value otherwise.
func (o *Customer) GetShipping() ShippingDetails {
	if o == nil || o.Shipping == nil {
		var ret ShippingDetails
		return ret
	}
	return *o.Shipping
}

// GetShippingOk returns a tuple with the Shipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetShippingOk() (*ShippingDetails, bool) {
	if o == nil || o.Shipping == nil {
		return nil, false
	}
	return o.Shipping, true
}

// HasShipping returns a boolean if a field has been set.
func (o *Customer) HasShipping() bool {
	if o != nil && o.Shipping != nil {
		return true
	}

	return false
}

// SetShipping gets a reference to the given ShippingDetails and assigns it to the Shipping field.
func (o *Customer) SetShipping(v ShippingDetails) {
	o.Shipping = &v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *Customer) GetSources() ExternalAccountCollection {
	if o == nil || o.Sources == nil {
		var ret ExternalAccountCollection
		return ret
	}
	return *o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetSourcesOk() (*ExternalAccountCollection, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *Customer) HasSources() bool {
	if o != nil && o.Sources != nil {
		return true
	}

	return false
}

// SetSources gets a reference to the given ExternalAccountCollection and assigns it to the Sources field.
func (o *Customer) SetSources(v ExternalAccountCollection) {
	o.Sources = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *Customer) GetSubscription() Subscription {
	if o == nil || o.Subscription == nil {
		var ret Subscription
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetSubscriptionOk() (*Subscription, bool) {
	if o == nil || o.Subscription == nil {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *Customer) HasSubscription() bool {
	if o != nil && o.Subscription != nil {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given Subscription and assigns it to the Subscription field.
func (o *Customer) SetSubscription(v Subscription) {
	o.Subscription = &v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *Customer) GetSubscriptions() CustomerSubscriptionCollection {
	if o == nil || o.Subscriptions == nil {
		var ret CustomerSubscriptionCollection
		return ret
	}
	return *o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetSubscriptionsOk() (*CustomerSubscriptionCollection, bool) {
	if o == nil || o.Subscriptions == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *Customer) HasSubscriptions() bool {
	if o != nil && o.Subscriptions != nil {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given CustomerSubscriptionCollection and assigns it to the Subscriptions field.
func (o *Customer) SetSubscriptions(v CustomerSubscriptionCollection) {
	o.Subscriptions = &v
}

// GetTrialEnd returns the TrialEnd field value if set, zero value otherwise.
func (o *Customer) GetTrialEnd() int64 {
	if o == nil || o.TrialEnd == nil {
		var ret int64
		return ret
	}
	return *o.TrialEnd
}

// GetTrialEndOk returns a tuple with the TrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetTrialEndOk() (*int64, bool) {
	if o == nil || o.TrialEnd == nil {
		return nil, false
	}
	return o.TrialEnd, true
}

// HasTrialEnd returns a boolean if a field has been set.
func (o *Customer) HasTrialEnd() bool {
	if o != nil && o.TrialEnd != nil {
		return true
	}

	return false
}

// SetTrialEnd gets a reference to the given int64 and assigns it to the TrialEnd field.
func (o *Customer) SetTrialEnd(v int64) {
	o.TrialEnd = &v
}

func (o Customer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountBalance != nil {
		toSerialize["accountBalance"] = o.AccountBalance
	}
	if o.BusinessVatId != nil {
		toSerialize["businessVatId"] = o.BusinessVatId
	}
	if o.Cards != nil {
		toSerialize["cards"] = o.Cards
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.DefaultCard != nil {
		toSerialize["defaultCard"] = o.DefaultCard
	}
	if o.DefaultSource != nil {
		toSerialize["defaultSource"] = o.DefaultSource
	}
	if o.DefaultSourceObject != nil {
		toSerialize["defaultSourceObject"] = o.DefaultSourceObject
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.Delinquent != nil {
		toSerialize["delinquent"] = o.Delinquent
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Discount != nil {
		toSerialize["discount"] = o.Discount
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Livemode != nil {
		toSerialize["livemode"] = o.Livemode
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.NextRecurringCharge != nil {
		toSerialize["nextRecurringCharge"] = o.NextRecurringCharge
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Shipping != nil {
		toSerialize["shipping"] = o.Shipping
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	if o.Subscription != nil {
		toSerialize["subscription"] = o.Subscription
	}
	if o.Subscriptions != nil {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if o.TrialEnd != nil {
		toSerialize["trialEnd"] = o.TrialEnd
	}
	return json.Marshal(toSerialize)
}

type NullableCustomer struct {
	value *Customer
	isSet bool
}

func (v NullableCustomer) Get() *Customer {
	return v.value
}

func (v *NullableCustomer) Set(val *Customer) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomer(val *Customer) *NullableCustomer {
	return &NullableCustomer{value: val, isSet: true}
}

func (v NullableCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
