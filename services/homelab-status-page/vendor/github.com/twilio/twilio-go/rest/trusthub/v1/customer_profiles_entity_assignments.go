/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Trusthub
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

// Optional parameters for the method 'CreateCustomerProfileEntityAssignment'
type CreateCustomerProfileEntityAssignmentParams struct {
	// The SID of an object bag that holds information of the different items.
	ObjectSid *string `json:"ObjectSid,omitempty"`
}

func (params *CreateCustomerProfileEntityAssignmentParams) SetObjectSid(ObjectSid string) *CreateCustomerProfileEntityAssignmentParams {
	params.ObjectSid = &ObjectSid
	return params
}

// Create a new Assigned Item.
func (c *ApiService) CreateCustomerProfileEntityAssignment(CustomerProfileSid string, params *CreateCustomerProfileEntityAssignmentParams) (*TrusthubV1CustomerProfileEntityAssignment, error) {
	path := "/v1/CustomerProfiles/{CustomerProfileSid}/EntityAssignments"
	path = strings.Replace(path, "{"+"CustomerProfileSid"+"}", CustomerProfileSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ObjectSid != nil {
		data.Set("ObjectSid", *params.ObjectSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1CustomerProfileEntityAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an Assignment Item Instance.
func (c *ApiService) DeleteCustomerProfileEntityAssignment(CustomerProfileSid string, Sid string) error {
	path := "/v1/CustomerProfiles/{CustomerProfileSid}/EntityAssignments/{Sid}"
	path = strings.Replace(path, "{"+"CustomerProfileSid"+"}", CustomerProfileSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch specific Assigned Item Instance.
func (c *ApiService) FetchCustomerProfileEntityAssignment(CustomerProfileSid string, Sid string) (*TrusthubV1CustomerProfileEntityAssignment, error) {
	path := "/v1/CustomerProfiles/{CustomerProfileSid}/EntityAssignments/{Sid}"
	path = strings.Replace(path, "{"+"CustomerProfileSid"+"}", CustomerProfileSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1CustomerProfileEntityAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCustomerProfileEntityAssignment'
type ListCustomerProfileEntityAssignmentParams struct {
	// A string to filter the results by (EndUserType or SupportingDocumentType) machine-name. This is useful when you want to retrieve the entity-assignment of a specific end-user or supporting document.
	ObjectType *string `json:"ObjectType,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCustomerProfileEntityAssignmentParams) SetObjectType(ObjectType string) *ListCustomerProfileEntityAssignmentParams {
	params.ObjectType = &ObjectType
	return params
}
func (params *ListCustomerProfileEntityAssignmentParams) SetPageSize(PageSize int) *ListCustomerProfileEntityAssignmentParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCustomerProfileEntityAssignmentParams) SetLimit(Limit int) *ListCustomerProfileEntityAssignmentParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of CustomerProfileEntityAssignment records from the API. Request is executed immediately.
func (c *ApiService) PageCustomerProfileEntityAssignment(CustomerProfileSid string, params *ListCustomerProfileEntityAssignmentParams, pageToken, pageNumber string) (*ListCustomerProfileEntityAssignmentResponse, error) {
	path := "/v1/CustomerProfiles/{CustomerProfileSid}/EntityAssignments"

	path = strings.Replace(path, "{"+"CustomerProfileSid"+"}", CustomerProfileSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ObjectType != nil {
		data.Set("ObjectType", *params.ObjectType)
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

	ps := &ListCustomerProfileEntityAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists CustomerProfileEntityAssignment records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCustomerProfileEntityAssignment(CustomerProfileSid string, params *ListCustomerProfileEntityAssignmentParams) ([]TrusthubV1CustomerProfileEntityAssignment, error) {
	response, errors := c.StreamCustomerProfileEntityAssignment(CustomerProfileSid, params)

	records := make([]TrusthubV1CustomerProfileEntityAssignment, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams CustomerProfileEntityAssignment records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCustomerProfileEntityAssignment(CustomerProfileSid string, params *ListCustomerProfileEntityAssignmentParams) (chan TrusthubV1CustomerProfileEntityAssignment, chan error) {
	if params == nil {
		params = &ListCustomerProfileEntityAssignmentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan TrusthubV1CustomerProfileEntityAssignment, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageCustomerProfileEntityAssignment(CustomerProfileSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamCustomerProfileEntityAssignment(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamCustomerProfileEntityAssignment(response *ListCustomerProfileEntityAssignmentResponse, params *ListCustomerProfileEntityAssignmentParams, recordChannel chan TrusthubV1CustomerProfileEntityAssignment, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Results
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListCustomerProfileEntityAssignmentResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListCustomerProfileEntityAssignmentResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListCustomerProfileEntityAssignmentResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCustomerProfileEntityAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
