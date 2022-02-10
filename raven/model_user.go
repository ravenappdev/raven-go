/*
 * Raven API
 *
 * This is OpenAPI defintion for the APIs of Raven.  You can find out more about Raven at [https://ravenapp.dev/](https://ravenapp.dev/).
 *
 * API version: 1.0.0
 * Contact: api@ravenapp.dev
 */

package raven

type User struct {
	UserId              string   `json:"user_id,omitempty"`
	Mobile              string   `json:"mobile,omitempty"`
	Email               string   `json:"email,omitempty"`
	WhatsappMobile      string   `json:"whatsapp_mobile,omitempty"`
	OnesignalExternalId string   `json:"onesignal_external_id,omitempty"`
	FcmTokens           []string `json:"fcm_tokens,omitempty"`
	IosTokens           []string `json:"ios_tokens,omitempty"`
}
