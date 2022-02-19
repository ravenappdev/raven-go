/*
 * Raven API
 *
 * This is OpenAPI defintion for the APIs of Raven.  You can find out more about Raven at [https://ravenapp.dev/](https://ravenapp.dev/).
 *
 * API version: 1.0.0
 * Contact: api@ravenapp.dev
 */

package raven

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/ravenappdev/raven-go/data"
)

// Linger please
var (
	_ context.Context
)

type EventApiService service

/*
EventApiService sends the event in bulk to all the clients specified
This API will send the event in bulk to the clients specified
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param appId app id of raven app
 * @param event the body for the event that has to be triggered

@return SuccessResponse
*/
func (a *EventApiService) SendBulkEvent(ctx context.Context, appId string, event data.SendEventBulk) (data.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue data.Response
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/apps/{app_id}/events/bulk_send"
	localVarPath = strings.Replace(localVarPath, "{"+"app_id"+"}", fmt.Sprintf("%v", appId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &event
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		if localVarHttpResponse.StatusCode == 401 {
			err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, err
			}
			return localVarReturnValue, errors.New(localVarHttpResponse.Status)
		}

		if localVarHttpResponse.StatusCode == 404 {
			err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, err
			}
			return localVarReturnValue, errors.New(localVarHttpResponse.Status)
		}

		return localVarReturnValue, errors.New(localVarHttpResponse.Status)
	}

	return localVarReturnValue, nil
}

/*
EventApiService sends the event to the client specified
This API will send the event to the client specified
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param appId app id of raven app
 * @param event the body for the event that has to be triggered

@return SuccessResponse
*/
func (a *EventApiService) SendEvent(ctx context.Context, appId string, event data.SendEvent) (data.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue data.Response
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/apps/{app_id}/events/send"
	localVarPath = strings.Replace(localVarPath, "{"+"app_id"+"}", fmt.Sprintf("%v", appId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &event
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, err
	}

	if localVarHttpResponse.StatusCode >= 300 {

		if localVarHttpResponse.StatusCode == 401 {
			err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, errors.New(localVarHttpResponse.Status)
			}
			return localVarReturnValue, errors.New(localVarHttpResponse.Status)
		}

		if localVarHttpResponse.StatusCode == 404 {
			err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, errors.New(localVarHttpResponse.Status)
			}
			return localVarReturnValue, errors.New(localVarHttpResponse.Status)
		}

		return localVarReturnValue, errors.New(localVarHttpResponse.Status)
	}
	return localVarReturnValue, nil
}
