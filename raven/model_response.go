/*
 * Raven API
 *
 * This is OpenAPI defintion for the APIs of Raven.  You can find out more about Raven at [https://ravenapp.dev/](https://ravenapp.dev/).
 *
 * API version: 1.0.0
 * Contact: api@ravenapp.dev
 */

package raven

type Response struct {
	Sucess bool   `json:"success,omitempty"`
	Id     string `json:"id,omitempty"`
	Error  string `json:"error,omitempty"`
}
