/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// IdentificationCode : Identification Code: * `B` - SWIFT Bank Identifier Code (BIC) * `C` - CHIPS Participant * `D` - Demand Deposit Account (DDA) Number * `F` - Fed Routing Number * `T` - SWIFT BIC or Bank Entity Identifier (BEI) and Account Number * `U` - CHIPS Identifier * `1` - Passport Number * `2` - Tax Identification Number * `3` - Driver’s License Number * `4` - Alien Registration Number * `5` - Corporate Identification * `9` - Other Identification 
type IdentificationCode string

// List of IdentificationCode
const (
	B IdentificationCode = "B"
	C IdentificationCode = "C"
	D IdentificationCode = "D"
	F IdentificationCode = "F"
	T IdentificationCode = "T"
	U IdentificationCode = "U"
	_1 IdentificationCode = "1"
	_2 IdentificationCode = "2"
	_3 IdentificationCode = "3"
	_4 IdentificationCode = "4"
	_5 IdentificationCode = "5"
	_9 IdentificationCode = "9"
)