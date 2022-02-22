/*
 * Raven API
 *
 * This is OpenAPI defintion for the APIs of Raven.  You can find out more about Raven at [https://ravenapp.dev/](https://ravenapp.dev/).
 *
 * API version: 1.0.0
 * Contact: api@ravenapp.dev
 */

package data

type Attachments struct {
	Filename string `json:"filename,omitempty"`
	Content  string `json:"content,omitempty"`
	Url      string `json:"url,omitempty"`
}

type EmailOverride struct {
	From        *EmailRecipient  `json:"from,omitempty"`
	Cc          []EmailRecipient `json:"cc,omitempty"`
	Bcc         []EmailRecipient `json:"bcc,omitempty"`
	Attachments *Attachments     `json:"attachments,omitempty"`
	ScheduledAt string           `json:"scheduled_at,omitempty"`
}

type ProviderOverride struct {
	Payload     interface{}              `json:"payload,omitempty"`
	FormParams  []ProviderOverrideParams `json:"form_params,omitempty"`
	QueryParams []ProviderOverrideParams `json:"query_params,omitempty"`
	Config      interface{}              `json:"config,omitempty"`
}

type ProviderOverrideParams struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type PushOverride struct {
	ScheduledAt string `json:"scheduled_at,omitempty"`
}

type SlackOverride struct {
	ScheduledAt string `json:"scheduled_at,omitempty"`
}

type SmsOverride struct {
	Sender      string `json:"sender,omitempty"`
	ScheduledAt string `json:"scheduled_at,omitempty"`
}

type VoiceOverride struct {
	ScheduledAt string `json:"scheduled_at,omitempty"`
}

type WebhookOverride struct {
	ScheduledAt string `json:"scheduled_at,omitempty"`
}

type WhatsappOverride struct {
	ScheduledAt string `json:"scheduled_at,omitempty"`
}

type EventOverride struct {
	Email     *EmailOverride    `json:"email,omitempty"`
	Sms       *SmsOverride      `json:"sms,omitempty"`
	Whatsapp  *WhatsappOverride `json:"whatsapp,omitempty"`
	Push      *PushOverride     `json:"push,omitempty"`
	Webhook   *WebhookOverride  `json:"webhook,omitempty"`
	Voice     *VoiceOverride    `json:"voice,omitempty"`
	Slack     *SlackOverride    `json:"slack,omitempty"`
	Providers *ProviderOverride `json:"providers,omitempty"`
}
