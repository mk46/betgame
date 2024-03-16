/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Sync
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
	"strings"
)

// Optional parameters for the method 'CreateStreamMessage'
type CreateStreamMessageParams struct {
	// A JSON string that represents an arbitrary, schema-less object that makes up the Stream Message body. Can be up to 4 KiB in length.
	Data *interface{} `json:"Data,omitempty"`
}

func (params *CreateStreamMessageParams) SetData(Data interface{}) *CreateStreamMessageParams {
	params.Data = &Data
	return params
}

// Create a new Stream Message.
func (c *ApiService) CreateStreamMessage(ServiceSid string, StreamSid string, params *CreateStreamMessageParams) (*SyncV1StreamMessage, error) {
	path := "/v1/Services/{ServiceSid}/Streams/{StreamSid}/Messages"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"StreamSid"+"}", StreamSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Data != nil {
		v, err := json.Marshal(params.Data)

		if err != nil {
			return nil, err
		}

		data.Set("Data", string(v))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1StreamMessage{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}