/*
 * Raven API
 *
 * This is OpenAPI defintion for the APIs of Raven.  You can find out more about Raven at [https://ravenapp.dev/](https://ravenapp.dev/).
 *
 * API version: 1.0.0
 * Contact: api@ravenapp.dev
 */

package data

type SendEventRequest struct {
	Event    string         `json:"event"`
	User     *User          `json:"user"`
	Data     *Data          `json:"data,omitempty"`
	Override *EventOverride `json:"override,omitempty"`
}
