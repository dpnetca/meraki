/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package meraki

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// DataRateHistoryApiService DataRateHistoryApi service
type DataRateHistoryApiService service

type DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest struct {
	ctx context.Context
	ApiService *DataRateHistoryApiService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	resolution *int32
	autoResolution *bool
	clientId *string
	deviceSerial *string
	apTag *string
	band *string
	ssid *int32
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest) T0(t0 string) DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest) T1(t1 string) DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
func (r DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest) Timespan(timespan float32) DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest {
	r.timespan = &timespan
	return r
}

// The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400.
func (r DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest) Resolution(resolution int32) DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest {
	r.resolution = &resolution
	return r
}

// Automatically select a data resolution based on the given timespan; this overrides the value specified by the &#39;resolution&#39; parameter. The default setting is false.
func (r DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest) AutoResolution(autoResolution bool) DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest {
	r.autoResolution = &autoResolution
	return r
}

// Filter results by network client.
func (r DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest) ClientId(clientId string) DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest {
	r.clientId = &clientId
	return r
}

// Filter results by device.
func (r DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest) DeviceSerial(deviceSerial string) DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest {
	r.deviceSerial = &deviceSerial
	return r
}

// Filter results by AP tag.
func (r DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest) ApTag(apTag string) DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest {
	r.apTag = &apTag
	return r
}

// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;).
func (r DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest) Band(band string) DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest {
	r.band = &band
	return r
}

// Filter results by SSID number.
func (r DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest) Ssid(ssid int32) DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest {
	r.ssid = &ssid
	return r
}

func (r DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest) Execute() ([]GetNetworkWirelessDataRateHistory200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessDataRateHistoryExecute(r)
}

/*
GetNetworkWirelessDataRateHistory Return PHY data rates over time for a network, device, or network client

Return PHY data rates over time for a network, device, or network client

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest
*/
func (a *DataRateHistoryApiService) GetNetworkWirelessDataRateHistory(ctx context.Context, networkId string) DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest {
	return DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkWirelessDataRateHistory200ResponseInner
func (a *DataRateHistoryApiService) GetNetworkWirelessDataRateHistoryExecute(r DataRateHistoryApiGetNetworkWirelessDataRateHistoryRequest) ([]GetNetworkWirelessDataRateHistory200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkWirelessDataRateHistory200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataRateHistoryApiService.GetNetworkWirelessDataRateHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/dataRateHistory"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
	}
	if r.resolution != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resolution", r.resolution, "")
	}
	if r.autoResolution != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "autoResolution", r.autoResolution, "")
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientId", r.clientId, "")
	}
	if r.deviceSerial != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deviceSerial", r.deviceSerial, "")
	}
	if r.apTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "apTag", r.apTag, "")
	}
	if r.band != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "band", r.band, "")
	}
	if r.ssid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ssid", r.ssid, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
