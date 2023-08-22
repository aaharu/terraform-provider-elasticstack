/*
SLOs

OpenAPI schema for SLOs endpoints

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slo

import (
	"encoding/json"
)

// checks if the IndicatorPropertiesCustomMetricParamsGoodMetricsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndicatorPropertiesCustomMetricParamsGoodMetricsInner{}

// IndicatorPropertiesCustomMetricParamsGoodMetricsInner struct for IndicatorPropertiesCustomMetricParamsGoodMetricsInner
type IndicatorPropertiesCustomMetricParamsGoodMetricsInner struct {
	// The name of the metric. Only valid options are A-Z
	Name string `json:"name"`
	// The aggregation type of the metric. Only valid option is \"sum\"
	Aggregation string `json:"aggregation"`
	// The field of the metric.
	Field string `json:"field"`
	// The filter to apply to the metric.
	Filter *string `json:"filter,omitempty"`
}

// NewIndicatorPropertiesCustomMetricParamsGoodMetricsInner instantiates a new IndicatorPropertiesCustomMetricParamsGoodMetricsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndicatorPropertiesCustomMetricParamsGoodMetricsInner(name string, aggregation string, field string) *IndicatorPropertiesCustomMetricParamsGoodMetricsInner {
	this := IndicatorPropertiesCustomMetricParamsGoodMetricsInner{}
	this.Name = name
	this.Aggregation = aggregation
	this.Field = field
	return &this
}

// NewIndicatorPropertiesCustomMetricParamsGoodMetricsInnerWithDefaults instantiates a new IndicatorPropertiesCustomMetricParamsGoodMetricsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndicatorPropertiesCustomMetricParamsGoodMetricsInnerWithDefaults() *IndicatorPropertiesCustomMetricParamsGoodMetricsInner {
	this := IndicatorPropertiesCustomMetricParamsGoodMetricsInner{}
	return &this
}

// GetName returns the Name field value
func (o *IndicatorPropertiesCustomMetricParamsGoodMetricsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IndicatorPropertiesCustomMetricParamsGoodMetricsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IndicatorPropertiesCustomMetricParamsGoodMetricsInner) SetName(v string) {
	o.Name = v
}

// GetAggregation returns the Aggregation field value
func (o *IndicatorPropertiesCustomMetricParamsGoodMetricsInner) GetAggregation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *IndicatorPropertiesCustomMetricParamsGoodMetricsInner) GetAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value
func (o *IndicatorPropertiesCustomMetricParamsGoodMetricsInner) SetAggregation(v string) {
	o.Aggregation = v
}

// GetField returns the Field field value
func (o *IndicatorPropertiesCustomMetricParamsGoodMetricsInner) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *IndicatorPropertiesCustomMetricParamsGoodMetricsInner) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *IndicatorPropertiesCustomMetricParamsGoodMetricsInner) SetField(v string) {
	o.Field = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *IndicatorPropertiesCustomMetricParamsGoodMetricsInner) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorPropertiesCustomMetricParamsGoodMetricsInner) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *IndicatorPropertiesCustomMetricParamsGoodMetricsInner) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *IndicatorPropertiesCustomMetricParamsGoodMetricsInner) SetFilter(v string) {
	o.Filter = &v
}

func (o IndicatorPropertiesCustomMetricParamsGoodMetricsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndicatorPropertiesCustomMetricParamsGoodMetricsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["aggregation"] = o.Aggregation
	toSerialize["field"] = o.Field
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	return toSerialize, nil
}

type NullableIndicatorPropertiesCustomMetricParamsGoodMetricsInner struct {
	value *IndicatorPropertiesCustomMetricParamsGoodMetricsInner
	isSet bool
}

func (v NullableIndicatorPropertiesCustomMetricParamsGoodMetricsInner) Get() *IndicatorPropertiesCustomMetricParamsGoodMetricsInner {
	return v.value
}

func (v *NullableIndicatorPropertiesCustomMetricParamsGoodMetricsInner) Set(val *IndicatorPropertiesCustomMetricParamsGoodMetricsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicatorPropertiesCustomMetricParamsGoodMetricsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicatorPropertiesCustomMetricParamsGoodMetricsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicatorPropertiesCustomMetricParamsGoodMetricsInner(val *IndicatorPropertiesCustomMetricParamsGoodMetricsInner) *NullableIndicatorPropertiesCustomMetricParamsGoodMetricsInner {
	return &NullableIndicatorPropertiesCustomMetricParamsGoodMetricsInner{value: val, isSet: true}
}

func (v NullableIndicatorPropertiesCustomMetricParamsGoodMetricsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicatorPropertiesCustomMetricParamsGoodMetricsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}