/*
SLOs

OpenAPI schema for SLOs endpoints

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slo

import (
	"encoding/json"
)

// checks if the TimeWindowCalendarAligned type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeWindowCalendarAligned{}

// TimeWindowCalendarAligned Defines properties for calendar aligned time window
type TimeWindowCalendarAligned struct {
	// the duration formatted as {duration}{unit}
	Duration string                            `json:"duration"`
	Calendar TimeWindowCalendarAlignedCalendar `json:"calendar"`
}

// NewTimeWindowCalendarAligned instantiates a new TimeWindowCalendarAligned object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeWindowCalendarAligned(duration string, calendar TimeWindowCalendarAlignedCalendar) *TimeWindowCalendarAligned {
	this := TimeWindowCalendarAligned{}
	this.Duration = duration
	this.Calendar = calendar
	return &this
}

// NewTimeWindowCalendarAlignedWithDefaults instantiates a new TimeWindowCalendarAligned object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeWindowCalendarAlignedWithDefaults() *TimeWindowCalendarAligned {
	this := TimeWindowCalendarAligned{}
	return &this
}

// GetDuration returns the Duration field value
func (o *TimeWindowCalendarAligned) GetDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *TimeWindowCalendarAligned) GetDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *TimeWindowCalendarAligned) SetDuration(v string) {
	o.Duration = v
}

// GetCalendar returns the Calendar field value
func (o *TimeWindowCalendarAligned) GetCalendar() TimeWindowCalendarAlignedCalendar {
	if o == nil {
		var ret TimeWindowCalendarAlignedCalendar
		return ret
	}

	return o.Calendar
}

// GetCalendarOk returns a tuple with the Calendar field value
// and a boolean to check if the value has been set.
func (o *TimeWindowCalendarAligned) GetCalendarOk() (*TimeWindowCalendarAlignedCalendar, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Calendar, true
}

// SetCalendar sets field value
func (o *TimeWindowCalendarAligned) SetCalendar(v TimeWindowCalendarAlignedCalendar) {
	o.Calendar = v
}

func (o TimeWindowCalendarAligned) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeWindowCalendarAligned) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["duration"] = o.Duration
	toSerialize["calendar"] = o.Calendar
	return toSerialize, nil
}

type NullableTimeWindowCalendarAligned struct {
	value *TimeWindowCalendarAligned
	isSet bool
}

func (v NullableTimeWindowCalendarAligned) Get() *TimeWindowCalendarAligned {
	return v.value
}

func (v *NullableTimeWindowCalendarAligned) Set(val *TimeWindowCalendarAligned) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeWindowCalendarAligned) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeWindowCalendarAligned) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeWindowCalendarAligned(val *TimeWindowCalendarAligned) *NullableTimeWindowCalendarAligned {
	return &NullableTimeWindowCalendarAligned{value: val, isSet: true}
}

func (v NullableTimeWindowCalendarAligned) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeWindowCalendarAligned) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}