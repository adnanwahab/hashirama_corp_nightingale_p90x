/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Routes
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// RoutesV2PhoneNumber struct for RoutesV2PhoneNumber
type RoutesV2PhoneNumber struct {
	// The phone number in E.164 format
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
	// A 34 character string that uniquely identifies the Inbound Processing Region assignments for this phone number.
	Sid *string `json:"sid,omitempty"`
	// The unique SID identifier of the Account.
	AccountSid *string `json:"account_sid,omitempty"`
	// A human readable description of the Inbound Processing Region assignments for this phone number, up to 64 characters.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The Inbound Processing Region used for this phone number for voice.
	VoiceRegion *string `json:"voice_region,omitempty"`
	// The date that this phone number was assigned an Inbound Processing Region, given in ISO 8601 format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that the Inbound Processing Region was updated for this phone number, given in ISO 8601 format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
}