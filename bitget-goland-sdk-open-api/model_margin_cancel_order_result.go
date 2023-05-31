/*
Bitget Open API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MarginCancelOrderResult struct for MarginCancelOrderResult
type MarginCancelOrderResult struct {
	ClientOid            *string `json:"clientOid,omitempty"`
	OrderId              *string `json:"orderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginCancelOrderResult MarginCancelOrderResult

// NewMarginCancelOrderResult instantiates a new MarginCancelOrderResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginCancelOrderResult() *MarginCancelOrderResult {
	this := MarginCancelOrderResult{}
	return &this
}

// NewMarginCancelOrderResultWithDefaults instantiates a new MarginCancelOrderResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginCancelOrderResultWithDefaults() *MarginCancelOrderResult {
	this := MarginCancelOrderResult{}
	return &this
}

// GetClientOid returns the ClientOid field value if set, zero value otherwise.
func (o *MarginCancelOrderResult) GetClientOid() string {
	if o == nil || isNil(o.ClientOid) {
		var ret string
		return ret
	}
	return *o.ClientOid
}

// GetClientOidOk returns a tuple with the ClientOid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCancelOrderResult) GetClientOidOk() (*string, bool) {
	if o == nil || isNil(o.ClientOid) {
		return nil, false
	}
	return o.ClientOid, true
}

// HasClientOid returns a boolean if a field has been set.
func (o *MarginCancelOrderResult) HasClientOid() bool {
	if o != nil && !isNil(o.ClientOid) {
		return true
	}

	return false
}

// SetClientOid gets a reference to the given string and assigns it to the ClientOid field.
func (o *MarginCancelOrderResult) SetClientOid(v string) {
	o.ClientOid = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *MarginCancelOrderResult) GetOrderId() string {
	if o == nil || isNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCancelOrderResult) GetOrderIdOk() (*string, bool) {
	if o == nil || isNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *MarginCancelOrderResult) HasOrderId() bool {
	if o != nil && !isNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *MarginCancelOrderResult) SetOrderId(v string) {
	o.OrderId = &v
}

func (o MarginCancelOrderResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClientOid) {
		toSerialize["clientOid"] = o.ClientOid
	}
	if !isNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginCancelOrderResult) UnmarshalJSON(bytes []byte) (err error) {
	varMarginCancelOrderResult := _MarginCancelOrderResult{}

	if err = json.Unmarshal(bytes, &varMarginCancelOrderResult); err == nil {
		*o = MarginCancelOrderResult(varMarginCancelOrderResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "clientOid")
		delete(additionalProperties, "orderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginCancelOrderResult struct {
	value *MarginCancelOrderResult
	isSet bool
}

func (v NullableMarginCancelOrderResult) Get() *MarginCancelOrderResult {
	return v.value
}

func (v *NullableMarginCancelOrderResult) Set(val *MarginCancelOrderResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginCancelOrderResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginCancelOrderResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginCancelOrderResult(val *MarginCancelOrderResult) *NullableMarginCancelOrderResult {
	return &NullableMarginCancelOrderResult{value: val, isSet: true}
}

func (v NullableMarginCancelOrderResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginCancelOrderResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}