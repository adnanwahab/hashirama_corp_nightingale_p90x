/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Video
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ListRoomParticipantPublishedTrackResponse struct for ListRoomParticipantPublishedTrackResponse
type ListRoomParticipantPublishedTrackResponse struct {
	PublishedTracks []VideoV1RoomParticipantPublishedTrack `json:"published_tracks,omitempty"`
	Meta            ListCompositionResponseMeta            `json:"meta,omitempty"`
}
