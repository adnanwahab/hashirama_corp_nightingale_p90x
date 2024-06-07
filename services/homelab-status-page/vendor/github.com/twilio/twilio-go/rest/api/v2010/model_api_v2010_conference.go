/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ApiV2010Conference struct for ApiV2010Conference
type ApiV2010Conference struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created this Conference resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The date and time in UTC that this resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *string `json:"date_created,omitempty"`
	// The date and time in UTC that this resource was last updated, specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *string `json:"date_updated,omitempty"`
	// The API version used to create this conference.
	ApiVersion *string `json:"api_version,omitempty"`
	// A string that you assigned to describe this conference room. Maximum length is 128 characters.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// A string that represents the Twilio Region where the conference audio was mixed. May be `us1`, `ie1`,  `de1`, `sg1`, `br1`, `au1`, and `jp1`. Basic conference audio will always be mixed in `us1`. Global Conference audio will be mixed nearest to the majority of participants.
	Region *string `json:"region,omitempty"`
	// The unique, Twilio-provided string used to identify this Conference resource.
	Sid    *string `json:"sid,omitempty"`
	Status *string `json:"status,omitempty"`
	// The URI of this resource, relative to `https://api.twilio.com`.
	Uri *string `json:"uri,omitempty"`
	// A list of related resources identified by their URIs relative to `https://api.twilio.com`.
	SubresourceUris       *map[string]interface{} `json:"subresource_uris,omitempty"`
	ReasonConferenceEnded *string                 `json:"reason_conference_ended,omitempty"`
	// The call SID that caused the conference to end.
	CallSidEndingConference *string `json:"call_sid_ending_conference,omitempty"`
}
