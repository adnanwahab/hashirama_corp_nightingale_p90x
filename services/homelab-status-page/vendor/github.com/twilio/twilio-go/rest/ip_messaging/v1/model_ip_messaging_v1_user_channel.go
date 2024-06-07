/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Ip_messaging
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// IpMessagingV1UserChannel struct for IpMessagingV1UserChannel
type IpMessagingV1UserChannel struct {
	AccountSid               *string                 `json:"account_sid,omitempty"`
	ServiceSid               *string                 `json:"service_sid,omitempty"`
	ChannelSid               *string                 `json:"channel_sid,omitempty"`
	MemberSid                *string                 `json:"member_sid,omitempty"`
	Status                   *string                 `json:"status,omitempty"`
	LastConsumedMessageIndex *int                    `json:"last_consumed_message_index,omitempty"`
	UnreadMessagesCount      *int                    `json:"unread_messages_count,omitempty"`
	Links                    *map[string]interface{} `json:"links,omitempty"`
}
