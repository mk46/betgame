/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Pricing
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"

	"github.com/twilio/twilio-go/client"
)

// PricingV1MessagingMessagingCountryInstanceInboundSmsPrices struct for PricingV1MessagingMessagingCountryInstanceInboundSmsPrices
type PricingV1MessagingMessagingCountryInstanceInboundSmsPrices struct {
	BasePrice    float32 `json:"base_price,omitempty"`
	CurrentPrice float32 `json:"current_price,omitempty"`
	NumberType   string  `json:"number_type,omitempty"`
}

func (response *PricingV1MessagingMessagingCountryInstanceInboundSmsPrices) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		BasePrice    interface{} `json:"base_price"`
		CurrentPrice interface{} `json:"current_price"`
		NumberType   string      `json:"number_type"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = PricingV1MessagingMessagingCountryInstanceInboundSmsPrices{
		NumberType: raw.NumberType,
	}

	responseBasePrice, err := client.UnmarshalFloat32(&raw.BasePrice)
	if err != nil {
		return err
	}
	response.BasePrice = *responseBasePrice

	responseCurrentPrice, err := client.UnmarshalFloat32(&raw.CurrentPrice)
	if err != nil {
		return err
	}
	response.CurrentPrice = *responseCurrentPrice

	return
}
