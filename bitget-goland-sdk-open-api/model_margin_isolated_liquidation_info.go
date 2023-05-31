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

// MarginIsolatedLiquidationInfo struct for MarginIsolatedLiquidationInfo
type MarginIsolatedLiquidationInfo struct {
	Ctime                *string `json:"ctime,omitempty"`
	LiqEndTime           *string `json:"liqEndTime,omitempty"`
	LiqFee               *string `json:"liqFee,omitempty"`
	LiqId                *string `json:"liqId,omitempty"`
	LiqRisk              *string `json:"liqRisk,omitempty"`
	LiqStartTime         *string `json:"liqStartTime,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	TotalAssets          *string `json:"totalAssets,omitempty"`
	TotalDebt            *string `json:"totalDebt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginIsolatedLiquidationInfo MarginIsolatedLiquidationInfo

// NewMarginIsolatedLiquidationInfo instantiates a new MarginIsolatedLiquidationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginIsolatedLiquidationInfo() *MarginIsolatedLiquidationInfo {
	this := MarginIsolatedLiquidationInfo{}
	return &this
}

// NewMarginIsolatedLiquidationInfoWithDefaults instantiates a new MarginIsolatedLiquidationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginIsolatedLiquidationInfoWithDefaults() *MarginIsolatedLiquidationInfo {
	this := MarginIsolatedLiquidationInfo{}
	return &this
}

// GetCtime returns the Ctime field value if set, zero value otherwise.
func (o *MarginIsolatedLiquidationInfo) GetCtime() string {
	if o == nil || isNil(o.Ctime) {
		var ret string
		return ret
	}
	return *o.Ctime
}

// GetCtimeOk returns a tuple with the Ctime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLiquidationInfo) GetCtimeOk() (*string, bool) {
	if o == nil || isNil(o.Ctime) {
		return nil, false
	}
	return o.Ctime, true
}

// HasCtime returns a boolean if a field has been set.
func (o *MarginIsolatedLiquidationInfo) HasCtime() bool {
	if o != nil && !isNil(o.Ctime) {
		return true
	}

	return false
}

// SetCtime gets a reference to the given string and assigns it to the Ctime field.
func (o *MarginIsolatedLiquidationInfo) SetCtime(v string) {
	o.Ctime = &v
}

// GetLiqEndTime returns the LiqEndTime field value if set, zero value otherwise.
func (o *MarginIsolatedLiquidationInfo) GetLiqEndTime() string {
	if o == nil || isNil(o.LiqEndTime) {
		var ret string
		return ret
	}
	return *o.LiqEndTime
}

// GetLiqEndTimeOk returns a tuple with the LiqEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLiquidationInfo) GetLiqEndTimeOk() (*string, bool) {
	if o == nil || isNil(o.LiqEndTime) {
		return nil, false
	}
	return o.LiqEndTime, true
}

// HasLiqEndTime returns a boolean if a field has been set.
func (o *MarginIsolatedLiquidationInfo) HasLiqEndTime() bool {
	if o != nil && !isNil(o.LiqEndTime) {
		return true
	}

	return false
}

// SetLiqEndTime gets a reference to the given string and assigns it to the LiqEndTime field.
func (o *MarginIsolatedLiquidationInfo) SetLiqEndTime(v string) {
	o.LiqEndTime = &v
}

// GetLiqFee returns the LiqFee field value if set, zero value otherwise.
func (o *MarginIsolatedLiquidationInfo) GetLiqFee() string {
	if o == nil || isNil(o.LiqFee) {
		var ret string
		return ret
	}
	return *o.LiqFee
}

// GetLiqFeeOk returns a tuple with the LiqFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLiquidationInfo) GetLiqFeeOk() (*string, bool) {
	if o == nil || isNil(o.LiqFee) {
		return nil, false
	}
	return o.LiqFee, true
}

// HasLiqFee returns a boolean if a field has been set.
func (o *MarginIsolatedLiquidationInfo) HasLiqFee() bool {
	if o != nil && !isNil(o.LiqFee) {
		return true
	}

	return false
}

// SetLiqFee gets a reference to the given string and assigns it to the LiqFee field.
func (o *MarginIsolatedLiquidationInfo) SetLiqFee(v string) {
	o.LiqFee = &v
}

// GetLiqId returns the LiqId field value if set, zero value otherwise.
func (o *MarginIsolatedLiquidationInfo) GetLiqId() string {
	if o == nil || isNil(o.LiqId) {
		var ret string
		return ret
	}
	return *o.LiqId
}

// GetLiqIdOk returns a tuple with the LiqId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLiquidationInfo) GetLiqIdOk() (*string, bool) {
	if o == nil || isNil(o.LiqId) {
		return nil, false
	}
	return o.LiqId, true
}

// HasLiqId returns a boolean if a field has been set.
func (o *MarginIsolatedLiquidationInfo) HasLiqId() bool {
	if o != nil && !isNil(o.LiqId) {
		return true
	}

	return false
}

// SetLiqId gets a reference to the given string and assigns it to the LiqId field.
func (o *MarginIsolatedLiquidationInfo) SetLiqId(v string) {
	o.LiqId = &v
}

// GetLiqRisk returns the LiqRisk field value if set, zero value otherwise.
func (o *MarginIsolatedLiquidationInfo) GetLiqRisk() string {
	if o == nil || isNil(o.LiqRisk) {
		var ret string
		return ret
	}
	return *o.LiqRisk
}

// GetLiqRiskOk returns a tuple with the LiqRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLiquidationInfo) GetLiqRiskOk() (*string, bool) {
	if o == nil || isNil(o.LiqRisk) {
		return nil, false
	}
	return o.LiqRisk, true
}

// HasLiqRisk returns a boolean if a field has been set.
func (o *MarginIsolatedLiquidationInfo) HasLiqRisk() bool {
	if o != nil && !isNil(o.LiqRisk) {
		return true
	}

	return false
}

// SetLiqRisk gets a reference to the given string and assigns it to the LiqRisk field.
func (o *MarginIsolatedLiquidationInfo) SetLiqRisk(v string) {
	o.LiqRisk = &v
}

// GetLiqStartTime returns the LiqStartTime field value if set, zero value otherwise.
func (o *MarginIsolatedLiquidationInfo) GetLiqStartTime() string {
	if o == nil || isNil(o.LiqStartTime) {
		var ret string
		return ret
	}
	return *o.LiqStartTime
}

// GetLiqStartTimeOk returns a tuple with the LiqStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLiquidationInfo) GetLiqStartTimeOk() (*string, bool) {
	if o == nil || isNil(o.LiqStartTime) {
		return nil, false
	}
	return o.LiqStartTime, true
}

// HasLiqStartTime returns a boolean if a field has been set.
func (o *MarginIsolatedLiquidationInfo) HasLiqStartTime() bool {
	if o != nil && !isNil(o.LiqStartTime) {
		return true
	}

	return false
}

// SetLiqStartTime gets a reference to the given string and assigns it to the LiqStartTime field.
func (o *MarginIsolatedLiquidationInfo) SetLiqStartTime(v string) {
	o.LiqStartTime = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MarginIsolatedLiquidationInfo) GetSymbol() string {
	if o == nil || isNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLiquidationInfo) GetSymbolOk() (*string, bool) {
	if o == nil || isNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MarginIsolatedLiquidationInfo) HasSymbol() bool {
	if o != nil && !isNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MarginIsolatedLiquidationInfo) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTotalAssets returns the TotalAssets field value if set, zero value otherwise.
func (o *MarginIsolatedLiquidationInfo) GetTotalAssets() string {
	if o == nil || isNil(o.TotalAssets) {
		var ret string
		return ret
	}
	return *o.TotalAssets
}

// GetTotalAssetsOk returns a tuple with the TotalAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLiquidationInfo) GetTotalAssetsOk() (*string, bool) {
	if o == nil || isNil(o.TotalAssets) {
		return nil, false
	}
	return o.TotalAssets, true
}

// HasTotalAssets returns a boolean if a field has been set.
func (o *MarginIsolatedLiquidationInfo) HasTotalAssets() bool {
	if o != nil && !isNil(o.TotalAssets) {
		return true
	}

	return false
}

// SetTotalAssets gets a reference to the given string and assigns it to the TotalAssets field.
func (o *MarginIsolatedLiquidationInfo) SetTotalAssets(v string) {
	o.TotalAssets = &v
}

// GetTotalDebt returns the TotalDebt field value if set, zero value otherwise.
func (o *MarginIsolatedLiquidationInfo) GetTotalDebt() string {
	if o == nil || isNil(o.TotalDebt) {
		var ret string
		return ret
	}
	return *o.TotalDebt
}

// GetTotalDebtOk returns a tuple with the TotalDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedLiquidationInfo) GetTotalDebtOk() (*string, bool) {
	if o == nil || isNil(o.TotalDebt) {
		return nil, false
	}
	return o.TotalDebt, true
}

// HasTotalDebt returns a boolean if a field has been set.
func (o *MarginIsolatedLiquidationInfo) HasTotalDebt() bool {
	if o != nil && !isNil(o.TotalDebt) {
		return true
	}

	return false
}

// SetTotalDebt gets a reference to the given string and assigns it to the TotalDebt field.
func (o *MarginIsolatedLiquidationInfo) SetTotalDebt(v string) {
	o.TotalDebt = &v
}

func (o MarginIsolatedLiquidationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ctime) {
		toSerialize["ctime"] = o.Ctime
	}
	if !isNil(o.LiqEndTime) {
		toSerialize["liqEndTime"] = o.LiqEndTime
	}
	if !isNil(o.LiqFee) {
		toSerialize["liqFee"] = o.LiqFee
	}
	if !isNil(o.LiqId) {
		toSerialize["liqId"] = o.LiqId
	}
	if !isNil(o.LiqRisk) {
		toSerialize["liqRisk"] = o.LiqRisk
	}
	if !isNil(o.LiqStartTime) {
		toSerialize["liqStartTime"] = o.LiqStartTime
	}
	if !isNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !isNil(o.TotalAssets) {
		toSerialize["totalAssets"] = o.TotalAssets
	}
	if !isNil(o.TotalDebt) {
		toSerialize["totalDebt"] = o.TotalDebt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginIsolatedLiquidationInfo) UnmarshalJSON(bytes []byte) (err error) {
	varMarginIsolatedLiquidationInfo := _MarginIsolatedLiquidationInfo{}

	if err = json.Unmarshal(bytes, &varMarginIsolatedLiquidationInfo); err == nil {
		*o = MarginIsolatedLiquidationInfo(varMarginIsolatedLiquidationInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ctime")
		delete(additionalProperties, "liqEndTime")
		delete(additionalProperties, "liqFee")
		delete(additionalProperties, "liqId")
		delete(additionalProperties, "liqRisk")
		delete(additionalProperties, "liqStartTime")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "totalAssets")
		delete(additionalProperties, "totalDebt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginIsolatedLiquidationInfo struct {
	value *MarginIsolatedLiquidationInfo
	isSet bool
}

func (v NullableMarginIsolatedLiquidationInfo) Get() *MarginIsolatedLiquidationInfo {
	return v.value
}

func (v *NullableMarginIsolatedLiquidationInfo) Set(val *MarginIsolatedLiquidationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginIsolatedLiquidationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginIsolatedLiquidationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginIsolatedLiquidationInfo(val *MarginIsolatedLiquidationInfo) *NullableMarginIsolatedLiquidationInfo {
	return &NullableMarginIsolatedLiquidationInfo{value: val, isSet: true}
}

func (v NullableMarginIsolatedLiquidationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginIsolatedLiquidationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}