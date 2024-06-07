/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
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

// FlexV1PluginVersionArchive struct for FlexV1PluginVersionArchive
type FlexV1PluginVersionArchive struct {
	// The unique string that we created to identify the Flex Plugin Version resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the Flex Plugin resource this Flex Plugin Version belongs to.
	PluginSid *string `json:"plugin_sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Flex Plugin Version resource and owns this resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique version of this Flex Plugin Version.
	Version *string `json:"version,omitempty"`
	// The URL of where the Flex Plugin Version JavaScript bundle is hosted on.
	PluginUrl *string `json:"plugin_url,omitempty"`
	// A changelog that describes the changes this Flex Plugin Version brings.
	Changelog *string `json:"changelog,omitempty"`
	// Whether to inject credentials while accessing this Plugin Version. The default value is false.
	Private *bool `json:"private,omitempty"`
	// Whether the Flex Plugin Version is archived. The default value is false.
	Archived *bool `json:"archived,omitempty"`
	// The date and time in GMT when the Flex Plugin Version was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The absolute URL of the Flex Plugin Version resource.
	Url *string `json:"url,omitempty"`
}
