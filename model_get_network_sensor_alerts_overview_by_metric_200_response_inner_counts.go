/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package meraki

import (
	"encoding/json"
)

// checks if the GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts{}

// GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts Counts of sensor alerts over the timespan, by reading metric
type GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts struct {
	// Number of sensor alerts that occurred due to an open door
	Door *int32 `json:"door,omitempty"`
	// Number of sensor alerts that occurred due to humidity readings
	Humidity *int32 `json:"humidity,omitempty"`
	// Number of sensor alerts that occurred due to indoor air quality readings
	IndoorAirQuality *int32 `json:"indoorAirQuality,omitempty"`
	Noise *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise `json:"noise,omitempty"`
	// Number of sensor alerts that occurred due to PM2.5 readings
	Pm25 *int32 `json:"pm25,omitempty"`
	// Number of sensor alerts that occurred due to temperature readings
	Temperature *int32 `json:"temperature,omitempty"`
	// Number of sensor alerts that occurred due to TVOC readings
	Tvoc *int32 `json:"tvoc,omitempty"`
	// Number of sensor alerts that occurred due to the presence of water
	Water *int32 `json:"water,omitempty"`
}

// NewGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts instantiates a new GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts() *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts {
	this := GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts{}
	return &this
}

// NewGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsWithDefaults instantiates a new GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsWithDefaults() *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts {
	this := GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts{}
	return &this
}

// GetDoor returns the Door field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetDoor() int32 {
	if o == nil || IsNil(o.Door) {
		var ret int32
		return ret
	}
	return *o.Door
}

// GetDoorOk returns a tuple with the Door field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetDoorOk() (*int32, bool) {
	if o == nil || IsNil(o.Door) {
		return nil, false
	}
	return o.Door, true
}

// HasDoor returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasDoor() bool {
	if o != nil && !IsNil(o.Door) {
		return true
	}

	return false
}

// SetDoor gets a reference to the given int32 and assigns it to the Door field.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetDoor(v int32) {
	o.Door = &v
}

// GetHumidity returns the Humidity field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetHumidity() int32 {
	if o == nil || IsNil(o.Humidity) {
		var ret int32
		return ret
	}
	return *o.Humidity
}

// GetHumidityOk returns a tuple with the Humidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetHumidityOk() (*int32, bool) {
	if o == nil || IsNil(o.Humidity) {
		return nil, false
	}
	return o.Humidity, true
}

// HasHumidity returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasHumidity() bool {
	if o != nil && !IsNil(o.Humidity) {
		return true
	}

	return false
}

// SetHumidity gets a reference to the given int32 and assigns it to the Humidity field.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetHumidity(v int32) {
	o.Humidity = &v
}

// GetIndoorAirQuality returns the IndoorAirQuality field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetIndoorAirQuality() int32 {
	if o == nil || IsNil(o.IndoorAirQuality) {
		var ret int32
		return ret
	}
	return *o.IndoorAirQuality
}

// GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetIndoorAirQualityOk() (*int32, bool) {
	if o == nil || IsNil(o.IndoorAirQuality) {
		return nil, false
	}
	return o.IndoorAirQuality, true
}

// HasIndoorAirQuality returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasIndoorAirQuality() bool {
	if o != nil && !IsNil(o.IndoorAirQuality) {
		return true
	}

	return false
}

// SetIndoorAirQuality gets a reference to the given int32 and assigns it to the IndoorAirQuality field.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetIndoorAirQuality(v int32) {
	o.IndoorAirQuality = &v
}

// GetNoise returns the Noise field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetNoise() GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise {
	if o == nil || IsNil(o.Noise) {
		var ret GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise
		return ret
	}
	return *o.Noise
}

// GetNoiseOk returns a tuple with the Noise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetNoiseOk() (*GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise, bool) {
	if o == nil || IsNil(o.Noise) {
		return nil, false
	}
	return o.Noise, true
}

// HasNoise returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasNoise() bool {
	if o != nil && !IsNil(o.Noise) {
		return true
	}

	return false
}

// SetNoise gets a reference to the given GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise and assigns it to the Noise field.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetNoise(v GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise) {
	o.Noise = &v
}

// GetPm25 returns the Pm25 field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetPm25() int32 {
	if o == nil || IsNil(o.Pm25) {
		var ret int32
		return ret
	}
	return *o.Pm25
}

// GetPm25Ok returns a tuple with the Pm25 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetPm25Ok() (*int32, bool) {
	if o == nil || IsNil(o.Pm25) {
		return nil, false
	}
	return o.Pm25, true
}

// HasPm25 returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasPm25() bool {
	if o != nil && !IsNil(o.Pm25) {
		return true
	}

	return false
}

// SetPm25 gets a reference to the given int32 and assigns it to the Pm25 field.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetPm25(v int32) {
	o.Pm25 = &v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetTemperature() int32 {
	if o == nil || IsNil(o.Temperature) {
		var ret int32
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetTemperatureOk() (*int32, bool) {
	if o == nil || IsNil(o.Temperature) {
		return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasTemperature() bool {
	if o != nil && !IsNil(o.Temperature) {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given int32 and assigns it to the Temperature field.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetTemperature(v int32) {
	o.Temperature = &v
}

// GetTvoc returns the Tvoc field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetTvoc() int32 {
	if o == nil || IsNil(o.Tvoc) {
		var ret int32
		return ret
	}
	return *o.Tvoc
}

// GetTvocOk returns a tuple with the Tvoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetTvocOk() (*int32, bool) {
	if o == nil || IsNil(o.Tvoc) {
		return nil, false
	}
	return o.Tvoc, true
}

// HasTvoc returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasTvoc() bool {
	if o != nil && !IsNil(o.Tvoc) {
		return true
	}

	return false
}

// SetTvoc gets a reference to the given int32 and assigns it to the Tvoc field.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetTvoc(v int32) {
	o.Tvoc = &v
}

// GetWater returns the Water field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetWater() int32 {
	if o == nil || IsNil(o.Water) {
		var ret int32
		return ret
	}
	return *o.Water
}

// GetWaterOk returns a tuple with the Water field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetWaterOk() (*int32, bool) {
	if o == nil || IsNil(o.Water) {
		return nil, false
	}
	return o.Water, true
}

// HasWater returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasWater() bool {
	if o != nil && !IsNil(o.Water) {
		return true
	}

	return false
}

// SetWater gets a reference to the given int32 and assigns it to the Water field.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetWater(v int32) {
	o.Water = &v
}

func (o GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Door) {
		toSerialize["door"] = o.Door
	}
	if !IsNil(o.Humidity) {
		toSerialize["humidity"] = o.Humidity
	}
	if !IsNil(o.IndoorAirQuality) {
		toSerialize["indoorAirQuality"] = o.IndoorAirQuality
	}
	if !IsNil(o.Noise) {
		toSerialize["noise"] = o.Noise
	}
	if !IsNil(o.Pm25) {
		toSerialize["pm25"] = o.Pm25
	}
	if !IsNil(o.Temperature) {
		toSerialize["temperature"] = o.Temperature
	}
	if !IsNil(o.Tvoc) {
		toSerialize["tvoc"] = o.Tvoc
	}
	if !IsNil(o.Water) {
		toSerialize["water"] = o.Water
	}
	return toSerialize, nil
}

type NullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts struct {
	value *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts
	isSet bool
}

func (v NullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) Get() *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts {
	return v.value
}

func (v *NullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) Set(val *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts(val *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) *NullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts {
	return &NullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts{value: val, isSet: true}
}

func (v NullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


