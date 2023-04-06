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


// DelegatedApiService DelegatedApi service
type DelegatedApiService service

type DelegatedApiCreateNetworkAppliancePrefixesDelegatedStaticRequest struct {
	ctx context.Context
	ApiService *DelegatedApiService
	networkId string
	createNetworkAppliancePrefixesDelegatedStatic *CreateNetworkAppliancePrefixesDelegatedStaticRequest
}

func (r DelegatedApiCreateNetworkAppliancePrefixesDelegatedStaticRequest) CreateNetworkAppliancePrefixesDelegatedStatic(createNetworkAppliancePrefixesDelegatedStatic CreateNetworkAppliancePrefixesDelegatedStaticRequest) DelegatedApiCreateNetworkAppliancePrefixesDelegatedStaticRequest {
	r.createNetworkAppliancePrefixesDelegatedStatic = &createNetworkAppliancePrefixesDelegatedStatic
	return r
}

func (r DelegatedApiCreateNetworkAppliancePrefixesDelegatedStaticRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateNetworkAppliancePrefixesDelegatedStaticExecute(r)
}

/*
CreateNetworkAppliancePrefixesDelegatedStatic Add a static delegated prefix from a network

Add a static delegated prefix from a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return DelegatedApiCreateNetworkAppliancePrefixesDelegatedStaticRequest
*/
func (a *DelegatedApiService) CreateNetworkAppliancePrefixesDelegatedStatic(ctx context.Context, networkId string) DelegatedApiCreateNetworkAppliancePrefixesDelegatedStaticRequest {
	return DelegatedApiCreateNetworkAppliancePrefixesDelegatedStaticRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DelegatedApiService) CreateNetworkAppliancePrefixesDelegatedStaticExecute(r DelegatedApiCreateNetworkAppliancePrefixesDelegatedStaticRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedApiService.CreateNetworkAppliancePrefixesDelegatedStatic")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/prefixes/delegated/statics"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createNetworkAppliancePrefixesDelegatedStatic == nil {
		return localVarReturnValue, nil, reportError("createNetworkAppliancePrefixesDelegatedStatic is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.createNetworkAppliancePrefixesDelegatedStatic
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

type DelegatedApiDeleteNetworkAppliancePrefixesDelegatedStaticRequest struct {
	ctx context.Context
	ApiService *DelegatedApiService
	networkId string
	staticDelegatedPrefixId string
}

func (r DelegatedApiDeleteNetworkAppliancePrefixesDelegatedStaticRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNetworkAppliancePrefixesDelegatedStaticExecute(r)
}

/*
DeleteNetworkAppliancePrefixesDelegatedStatic Delete a static delegated prefix from a network

Delete a static delegated prefix from a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param staticDelegatedPrefixId
 @return DelegatedApiDeleteNetworkAppliancePrefixesDelegatedStaticRequest
*/
func (a *DelegatedApiService) DeleteNetworkAppliancePrefixesDelegatedStatic(ctx context.Context, networkId string, staticDelegatedPrefixId string) DelegatedApiDeleteNetworkAppliancePrefixesDelegatedStaticRequest {
	return DelegatedApiDeleteNetworkAppliancePrefixesDelegatedStaticRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		staticDelegatedPrefixId: staticDelegatedPrefixId,
	}
}

// Execute executes the request
func (a *DelegatedApiService) DeleteNetworkAppliancePrefixesDelegatedStaticExecute(r DelegatedApiDeleteNetworkAppliancePrefixesDelegatedStaticRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedApiService.DeleteNetworkAppliancePrefixesDelegatedStatic")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"staticDelegatedPrefixId"+"}", url.PathEscape(parameterValueToString(r.staticDelegatedPrefixId, "staticDelegatedPrefixId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type DelegatedApiGetDeviceAppliancePrefixesDelegatedRequest struct {
	ctx context.Context
	ApiService *DelegatedApiService
	serial string
}

func (r DelegatedApiGetDeviceAppliancePrefixesDelegatedRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceAppliancePrefixesDelegatedExecute(r)
}

/*
GetDeviceAppliancePrefixesDelegated Return current delegated IPv6 prefixes on an appliance.

Return current delegated IPv6 prefixes on an appliance.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial
 @return DelegatedApiGetDeviceAppliancePrefixesDelegatedRequest
*/
func (a *DelegatedApiService) GetDeviceAppliancePrefixesDelegated(ctx context.Context, serial string) DelegatedApiGetDeviceAppliancePrefixesDelegatedRequest {
	return DelegatedApiGetDeviceAppliancePrefixesDelegatedRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *DelegatedApiService) GetDeviceAppliancePrefixesDelegatedExecute(r DelegatedApiGetDeviceAppliancePrefixesDelegatedRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedApiService.GetDeviceAppliancePrefixesDelegated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/appliance/prefixes/delegated"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterValueToString(r.serial, "serial")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type DelegatedApiGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest struct {
	ctx context.Context
	ApiService *DelegatedApiService
	serial string
}

func (r DelegatedApiGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceAppliancePrefixesDelegatedVlanAssignmentsExecute(r)
}

/*
GetDeviceAppliancePrefixesDelegatedVlanAssignments Return prefixes assigned to all IPv6 enabled VLANs on an appliance.

Return prefixes assigned to all IPv6 enabled VLANs on an appliance.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial
 @return DelegatedApiGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest
*/
func (a *DelegatedApiService) GetDeviceAppliancePrefixesDelegatedVlanAssignments(ctx context.Context, serial string) DelegatedApiGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest {
	return DelegatedApiGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *DelegatedApiService) GetDeviceAppliancePrefixesDelegatedVlanAssignmentsExecute(r DelegatedApiGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedApiService.GetDeviceAppliancePrefixesDelegatedVlanAssignments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/appliance/prefixes/delegated/vlanAssignments"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterValueToString(r.serial, "serial")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type DelegatedApiGetNetworkAppliancePrefixesDelegatedStaticRequest struct {
	ctx context.Context
	ApiService *DelegatedApiService
	networkId string
	staticDelegatedPrefixId string
}

func (r DelegatedApiGetNetworkAppliancePrefixesDelegatedStaticRequest) Execute() (*GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkAppliancePrefixesDelegatedStaticExecute(r)
}

/*
GetNetworkAppliancePrefixesDelegatedStatic Return a static delegated prefix from a network

Return a static delegated prefix from a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param staticDelegatedPrefixId
 @return DelegatedApiGetNetworkAppliancePrefixesDelegatedStaticRequest
*/
func (a *DelegatedApiService) GetNetworkAppliancePrefixesDelegatedStatic(ctx context.Context, networkId string, staticDelegatedPrefixId string) DelegatedApiGetNetworkAppliancePrefixesDelegatedStaticRequest {
	return DelegatedApiGetNetworkAppliancePrefixesDelegatedStaticRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		staticDelegatedPrefixId: staticDelegatedPrefixId,
	}
}

// Execute executes the request
//  @return GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
func (a *DelegatedApiService) GetNetworkAppliancePrefixesDelegatedStaticExecute(r DelegatedApiGetNetworkAppliancePrefixesDelegatedStaticRequest) (*GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedApiService.GetNetworkAppliancePrefixesDelegatedStatic")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"staticDelegatedPrefixId"+"}", url.PathEscape(parameterValueToString(r.staticDelegatedPrefixId, "staticDelegatedPrefixId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type DelegatedApiGetNetworkAppliancePrefixesDelegatedStaticsRequest struct {
	ctx context.Context
	ApiService *DelegatedApiService
	networkId string
}

func (r DelegatedApiGetNetworkAppliancePrefixesDelegatedStaticsRequest) Execute() ([]GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkAppliancePrefixesDelegatedStaticsExecute(r)
}

/*
GetNetworkAppliancePrefixesDelegatedStatics List static delegated prefixes for a network

List static delegated prefixes for a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return DelegatedApiGetNetworkAppliancePrefixesDelegatedStaticsRequest
*/
func (a *DelegatedApiService) GetNetworkAppliancePrefixesDelegatedStatics(ctx context.Context, networkId string) DelegatedApiGetNetworkAppliancePrefixesDelegatedStaticsRequest {
	return DelegatedApiGetNetworkAppliancePrefixesDelegatedStaticsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
func (a *DelegatedApiService) GetNetworkAppliancePrefixesDelegatedStaticsExecute(r DelegatedApiGetNetworkAppliancePrefixesDelegatedStaticsRequest) ([]GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedApiService.GetNetworkAppliancePrefixesDelegatedStatics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/prefixes/delegated/statics"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type DelegatedApiUpdateNetworkAppliancePrefixesDelegatedStaticRequest struct {
	ctx context.Context
	ApiService *DelegatedApiService
	networkId string
	staticDelegatedPrefixId string
	updateNetworkAppliancePrefixesDelegatedStatic *UpdateNetworkAppliancePrefixesDelegatedStaticRequest
}

func (r DelegatedApiUpdateNetworkAppliancePrefixesDelegatedStaticRequest) UpdateNetworkAppliancePrefixesDelegatedStatic(updateNetworkAppliancePrefixesDelegatedStatic UpdateNetworkAppliancePrefixesDelegatedStaticRequest) DelegatedApiUpdateNetworkAppliancePrefixesDelegatedStaticRequest {
	r.updateNetworkAppliancePrefixesDelegatedStatic = &updateNetworkAppliancePrefixesDelegatedStatic
	return r
}

func (r DelegatedApiUpdateNetworkAppliancePrefixesDelegatedStaticRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateNetworkAppliancePrefixesDelegatedStaticExecute(r)
}

/*
UpdateNetworkAppliancePrefixesDelegatedStatic Update a static delegated prefix from a network

Update a static delegated prefix from a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param staticDelegatedPrefixId
 @return DelegatedApiUpdateNetworkAppliancePrefixesDelegatedStaticRequest
*/
func (a *DelegatedApiService) UpdateNetworkAppliancePrefixesDelegatedStatic(ctx context.Context, networkId string, staticDelegatedPrefixId string) DelegatedApiUpdateNetworkAppliancePrefixesDelegatedStaticRequest {
	return DelegatedApiUpdateNetworkAppliancePrefixesDelegatedStaticRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		staticDelegatedPrefixId: staticDelegatedPrefixId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DelegatedApiService) UpdateNetworkAppliancePrefixesDelegatedStaticExecute(r DelegatedApiUpdateNetworkAppliancePrefixesDelegatedStaticRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedApiService.UpdateNetworkAppliancePrefixesDelegatedStatic")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"staticDelegatedPrefixId"+"}", url.PathEscape(parameterValueToString(r.staticDelegatedPrefixId, "staticDelegatedPrefixId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.updateNetworkAppliancePrefixesDelegatedStatic
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
