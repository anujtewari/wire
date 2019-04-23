/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type MessageDisposition struct {
	// FormatVersion 30 
	FormatVersion string `json:"formatVersion,omitempty"`
	TestProductionCode TestProductionCodeEnum `json:"testProductionCode,omitempty"`
	// MessageDuplicationCode  * ` ` - Original Message * `R` - Retrieval of an original message * `P` - Resend 
	MessageDuplicationCode string `json:"messageDuplicationCode,omitempty"`
}