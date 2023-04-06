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


// RendezvousPointsApiService RendezvousPointsApi service
type RendezvousPointsApiService service

type RendezvousPointsApiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest struct {
	ctx context.Context
	ApiService *RendezvousPointsApiService
	networkId string
	createNetworkSwitchRoutingMulticastRendezvousPoint *CreateNetworkSwitchRoutingMulticastRendezvousPointRequest
}

func (r RendezvousPointsApiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest) CreateNetworkSwitchRoutingMulticastRendezvousPoint(createNetworkSwitchRoutingMulticastRendezvousPoint CreateNetworkSwitchRoutingMulticastRendezvousPointRequest) RendezvousPointsApiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest {
	r.createNetworkSwitchRoutingMulticastRendezvousPoint = &createNetworkSwitchRoutingMulticastRendezvousPoint
	return r
}

func (r RendezvousPointsApiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateNetworkSwitchRoutingMulticastRendezvousPointExecute(r)
}

/*
CreateNetworkSwitchRoutingMulticastRendezvousPoint Create a multicast rendezvous point

Create a multicast rendezvous point

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return RendezvousPointsApiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest
*/
func (a *RendezvousPointsApiService) CreateNetworkSwitchRoutingMulticastRendezvousPoint(ctx context.Context, networkId string) RendezvousPointsApiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest {
	return RendezvousPointsApiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RendezvousPointsApiService) CreateNetworkSwitchRoutingMulticastRendezvousPointExecute(r RendezvousPointsApiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RendezvousPointsApiService.CreateNetworkSwitchRoutingMulticastRendezvousPoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/routing/multicast/rendezvousPoints"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createNetworkSwitchRoutingMulticastRendezvousPoint == nil {
		return localVarReturnValue, nil, reportError("createNetworkSwitchRoutingMulticastRendezvousPoint is required and must be specified")
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
	localVarPostBody = r.createNetworkSwitchRoutingMulticastRendezvousPoint
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

type RendezvousPointsApiDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest struct {
	ctx context.Context
	ApiService *RendezvousPointsApiService
	networkId string
	rendezvousPointId string
}

func (r RendezvousPointsApiDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNetworkSwitchRoutingMulticastRendezvousPointExecute(r)
}

/*
DeleteNetworkSwitchRoutingMulticastRendezvousPoint Delete a multicast rendezvous point

Delete a multicast rendezvous point

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param rendezvousPointId
 @return RendezvousPointsApiDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest
*/
func (a *RendezvousPointsApiService) DeleteNetworkSwitchRoutingMulticastRendezvousPoint(ctx context.Context, networkId string, rendezvousPointId string) RendezvousPointsApiDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest {
	return RendezvousPointsApiDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		rendezvousPointId: rendezvousPointId,
	}
}

// Execute executes the request
func (a *RendezvousPointsApiService) DeleteNetworkSwitchRoutingMulticastRendezvousPointExecute(r RendezvousPointsApiDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RendezvousPointsApiService.DeleteNetworkSwitchRoutingMulticastRendezvousPoint")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rendezvousPointId"+"}", url.PathEscape(parameterValueToString(r.rendezvousPointId, "rendezvousPointId")), -1)

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

type RendezvousPointsApiGetNetworkSwitchRoutingMulticastRendezvousPointRequest struct {
	ctx context.Context
	ApiService *RendezvousPointsApiService
	networkId string
	rendezvousPointId string
}

func (r RendezvousPointsApiGetNetworkSwitchRoutingMulticastRendezvousPointRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchRoutingMulticastRendezvousPointExecute(r)
}

/*
GetNetworkSwitchRoutingMulticastRendezvousPoint Return a multicast rendezvous point

Return a multicast rendezvous point

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param rendezvousPointId
 @return RendezvousPointsApiGetNetworkSwitchRoutingMulticastRendezvousPointRequest
*/
func (a *RendezvousPointsApiService) GetNetworkSwitchRoutingMulticastRendezvousPoint(ctx context.Context, networkId string, rendezvousPointId string) RendezvousPointsApiGetNetworkSwitchRoutingMulticastRendezvousPointRequest {
	return RendezvousPointsApiGetNetworkSwitchRoutingMulticastRendezvousPointRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		rendezvousPointId: rendezvousPointId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RendezvousPointsApiService) GetNetworkSwitchRoutingMulticastRendezvousPointExecute(r RendezvousPointsApiGetNetworkSwitchRoutingMulticastRendezvousPointRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RendezvousPointsApiService.GetNetworkSwitchRoutingMulticastRendezvousPoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rendezvousPointId"+"}", url.PathEscape(parameterValueToString(r.rendezvousPointId, "rendezvousPointId")), -1)

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

type RendezvousPointsApiGetNetworkSwitchRoutingMulticastRendezvousPointsRequest struct {
	ctx context.Context
	ApiService *RendezvousPointsApiService
	networkId string
}

func (r RendezvousPointsApiGetNetworkSwitchRoutingMulticastRendezvousPointsRequest) Execute() ([][]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchRoutingMulticastRendezvousPointsExecute(r)
}

/*
GetNetworkSwitchRoutingMulticastRendezvousPoints List multicast rendezvous points

List multicast rendezvous points

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return RendezvousPointsApiGetNetworkSwitchRoutingMulticastRendezvousPointsRequest
*/
func (a *RendezvousPointsApiService) GetNetworkSwitchRoutingMulticastRendezvousPoints(ctx context.Context, networkId string) RendezvousPointsApiGetNetworkSwitchRoutingMulticastRendezvousPointsRequest {
	return RendezvousPointsApiGetNetworkSwitchRoutingMulticastRendezvousPointsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return [][]map[string]interface{}
func (a *RendezvousPointsApiService) GetNetworkSwitchRoutingMulticastRendezvousPointsExecute(r RendezvousPointsApiGetNetworkSwitchRoutingMulticastRendezvousPointsRequest) ([][]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  [][]map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RendezvousPointsApiService.GetNetworkSwitchRoutingMulticastRendezvousPoints")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/routing/multicast/rendezvousPoints"
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

type RendezvousPointsApiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest struct {
	ctx context.Context
	ApiService *RendezvousPointsApiService
	networkId string
	rendezvousPointId string
	updateNetworkSwitchRoutingMulticastRendezvousPoint *UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest
}

func (r RendezvousPointsApiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) UpdateNetworkSwitchRoutingMulticastRendezvousPoint(updateNetworkSwitchRoutingMulticastRendezvousPoint UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) RendezvousPointsApiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest {
	r.updateNetworkSwitchRoutingMulticastRendezvousPoint = &updateNetworkSwitchRoutingMulticastRendezvousPoint
	return r
}

func (r RendezvousPointsApiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateNetworkSwitchRoutingMulticastRendezvousPointExecute(r)
}

/*
UpdateNetworkSwitchRoutingMulticastRendezvousPoint Update a multicast rendezvous point

Update a multicast rendezvous point

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param rendezvousPointId
 @return RendezvousPointsApiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest
*/
func (a *RendezvousPointsApiService) UpdateNetworkSwitchRoutingMulticastRendezvousPoint(ctx context.Context, networkId string, rendezvousPointId string) RendezvousPointsApiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest {
	return RendezvousPointsApiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		rendezvousPointId: rendezvousPointId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RendezvousPointsApiService) UpdateNetworkSwitchRoutingMulticastRendezvousPointExecute(r RendezvousPointsApiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RendezvousPointsApiService.UpdateNetworkSwitchRoutingMulticastRendezvousPoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rendezvousPointId"+"}", url.PathEscape(parameterValueToString(r.rendezvousPointId, "rendezvousPointId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateNetworkSwitchRoutingMulticastRendezvousPoint == nil {
		return localVarReturnValue, nil, reportError("updateNetworkSwitchRoutingMulticastRendezvousPoint is required and must be specified")
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
	localVarPostBody = r.updateNetworkSwitchRoutingMulticastRendezvousPoint
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
