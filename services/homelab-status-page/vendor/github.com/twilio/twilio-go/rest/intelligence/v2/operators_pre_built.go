/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Intelligence
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Fetch a specific Pre-built Operator.
func (c *ApiService) FetchPrebuiltOperator(Sid string) (*IntelligenceV2PrebuiltOperator, error) {
	path := "/v2/Operators/PreBuilt/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IntelligenceV2PrebuiltOperator{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListPrebuiltOperator'
type ListPrebuiltOperatorParams struct {
	// Returns Pre-built Operators with the provided availability type. Possible values: internal, beta, public, retired.
	Availability *string `json:"Availability,omitempty"`
	// Returns Pre-built Operators that support the provided language code.
	LanguageCode *string `json:"LanguageCode,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListPrebuiltOperatorParams) SetAvailability(Availability string) *ListPrebuiltOperatorParams {
	params.Availability = &Availability
	return params
}
func (params *ListPrebuiltOperatorParams) SetLanguageCode(LanguageCode string) *ListPrebuiltOperatorParams {
	params.LanguageCode = &LanguageCode
	return params
}
func (params *ListPrebuiltOperatorParams) SetPageSize(PageSize int) *ListPrebuiltOperatorParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListPrebuiltOperatorParams) SetLimit(Limit int) *ListPrebuiltOperatorParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of PrebuiltOperator records from the API. Request is executed immediately.
func (c *ApiService) PagePrebuiltOperator(params *ListPrebuiltOperatorParams, pageToken, pageNumber string) (*ListPrebuiltOperatorResponse, error) {
	path := "/v2/Operators/PreBuilt"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Availability != nil {
		data.Set("Availability", *params.Availability)
	}
	if params != nil && params.LanguageCode != nil {
		data.Set("LanguageCode", *params.LanguageCode)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListPrebuiltOperatorResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists PrebuiltOperator records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListPrebuiltOperator(params *ListPrebuiltOperatorParams) ([]IntelligenceV2PrebuiltOperator, error) {
	response, errors := c.StreamPrebuiltOperator(params)

	records := make([]IntelligenceV2PrebuiltOperator, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams PrebuiltOperator records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamPrebuiltOperator(params *ListPrebuiltOperatorParams) (chan IntelligenceV2PrebuiltOperator, chan error) {
	if params == nil {
		params = &ListPrebuiltOperatorParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan IntelligenceV2PrebuiltOperator, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PagePrebuiltOperator(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamPrebuiltOperator(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamPrebuiltOperator(response *ListPrebuiltOperatorResponse, params *ListPrebuiltOperatorParams, recordChannel chan IntelligenceV2PrebuiltOperator, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Operators
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListPrebuiltOperatorResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListPrebuiltOperatorResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListPrebuiltOperatorResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListPrebuiltOperatorResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
