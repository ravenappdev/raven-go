/*
 * Raven API
 *
 * This is OpenAPI defintion for the APIs of Raven.  You can find out more about Raven at [https://ravenapp.dev/](https://ravenapp.dev/).
 *
 * API version: 1.0.0
 * Contact: api@ravenapp.dev
 */

package raven

type BatchItem struct {
	User *User `json:"user"`
	Data *Data `json:"data,omitempty"`
}

type SendEventBulk struct {
	Event string      `json:"event"`
	Batch []BatchItem `json:"batch,omitempty"`
}
