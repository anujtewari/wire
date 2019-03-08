/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Address
type Address struct {
	// AddressLineOne
	AddressLineOne string `json:"addressLineOne,omitempty"`
	// AddressLineTwo
	AddressLineTwo string `json:"addressLineTwo,omitempty"`
	// AddressLineThree
	AddressLineThree string `json:"addressLineThree,omitempty"`
}
