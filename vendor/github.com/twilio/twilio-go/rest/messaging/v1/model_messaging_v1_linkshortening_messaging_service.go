/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Messaging
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// MessagingV1LinkshorteningMessagingService struct for MessagingV1LinkshorteningMessagingService
type MessagingV1LinkshorteningMessagingService struct {
	// The unique string identifies the domain resource
	DomainSid *string `json:"domain_sid,omitempty"`
	// The unique string that identifies the messaging service
	MessagingServiceSid *string `json:"messaging_service_sid,omitempty"`
	Url                 *string `json:"url,omitempty"`
}