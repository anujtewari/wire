// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// RelatedRemittance is related remittance
type RelatedRemittance struct {
	// tag
	tag string
	// RemittanceIdentification is remittance identification
	RemittanceIdentification string `json:"remittanceIdentification,omitempty"`
	// RemittanceLocationMethod is  remittance location method
	RemittanceLocationMethod string `json:"remittanceLocationMethod,omitempty"`
	// RemittanceLocationElectronicAddress (E-mail or URL address)
	RemittanceLocationElectronicAddress string `json:"remittanceLocationElctronicAddress,omitempty"`
	// RemittanceData is RemittanceData
	RemittanceData RemittanceData `json:"remittanceData,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewRelatedRemittance returns a new RelatedRemittance
func NewRelatedRemittance() RelatedRemittance {
	rr := RelatedRemittance{
		tag: TagRelatedRemittance,
	}
	return rr
}

// Parse takes the input string and parses the RelatedRemittance values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rr *RelatedRemittance) Parse(record string) {
	rr.tag = record[:6]
	rr.RemittanceIdentification = rr.parseStringField(record[6:41])
	rr.RemittanceLocationMethod = rr.parseStringField(record[41:45])
	rr.RemittanceLocationElectronicAddress = rr.parseStringField(record[45:2093])
	rr.RemittanceData.Name = rr.parseStringField(record[2093:2233])
	rr.RemittanceData.AddressType = rr.parseStringField(record[2233:2237])
	rr.RemittanceData.Department = rr.parseStringField(record[2237:2307])
	rr.RemittanceData.SubDepartment = rr.parseStringField(record[2307:2377])
	rr.RemittanceData.StreetName = rr.parseStringField(record[2377:2447])
	rr.RemittanceData.BuildingNumber = rr.parseStringField(record[2447:2463])
	rr.RemittanceData.PostCode = rr.parseStringField(record[2463:2479])
	rr.RemittanceData.TownName = rr.parseStringField(record[2479:2514])
	rr.RemittanceData.CountrySubDivisionState = rr.parseStringField(record[2514:2549])
	rr.RemittanceData.Country = rr.parseStringField(record[2549:2551])
	rr.RemittanceData.AddressLineOne = rr.parseStringField(record[2551:2619])
	rr.RemittanceData.AddressLineTwo = rr.parseStringField(record[2619:2689])
	rr.RemittanceData.AddressLineThree = rr.parseStringField(record[2689:2759])
	rr.RemittanceData.AddressLineFour = rr.parseStringField(record[2759:2829])
	rr.RemittanceData.AddressLineFive = rr.parseStringField(record[2829:2899])
	rr.RemittanceData.AddressLineSix = rr.parseStringField(record[2899:2969])
	rr.RemittanceData.AddressLineSeven = rr.parseStringField(record[2969:3039])
}

// String writes RelatedRemittance
func (rr *RelatedRemittance) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(3035)
	buf.WriteString(rr.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on RelatedRemittance and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (rr *RelatedRemittance) Validate() error {
	if err := rr.fieldInclusion(); err != nil {
		return err
	}

	if err := rr.isAlphanumeric(rr.RemittanceIdentification); err != nil {
		return fieldError("RemittanceIdentification", err, rr.RemittanceIdentification)
	}

	if err := rr.isAlphanumeric(rr.RemittanceLocationElectronicAddress); err != nil {
		return fieldError("RemittanceLocationElectronicAddress", err, rr.RemittanceLocationElectronicAddress)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.Name); err != nil {
		return fieldError("Name", err, rr.RemittanceData.Name)
	}

	if err := rr.isAlphanumeric(rr.RemittanceData.Department); err != nil {
		return fieldError("", err, rr.RemittanceData.Department )
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.SubDepartment); err != nil {
		return fieldError("SubDepartment", err, rr.RemittanceData.SubDepartment)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.StreetName); err != nil {
		return fieldError("StreetName", err, rr.RemittanceData.StreetName)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.BuildingNumber); err != nil {
		return fieldError("BuildingNumber", err, rr.RemittanceData.BuildingNumber)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.PostCode); err != nil {
		return fieldError("PostCode", err, rr.RemittanceData.PostCode)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.TownName); err != nil {
		return fieldError("TownName", err, rr.RemittanceData.TownName)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.CountrySubDivisionState); err != nil {
		return fieldError("CountrySubDivisionState", err, rr.RemittanceData.CountrySubDivisionState)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.Country); err != nil {
		return fieldError("Country", err, rr.RemittanceData.Country)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, rr.RemittanceData.AddressLineOne)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, rr.RemittanceData.AddressLineTwo)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, rr.RemittanceData.AddressLineThree)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineFour); err != nil {
		return fieldError("AddressLineFour", err, rr.RemittanceData.AddressLineFour)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineFive); err != nil {
		return fieldError("AddressLineFive", err, rr.RemittanceData.AddressLineFive)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineSix); err != nil {
		return fieldError("AddressLineSix", err, rr.RemittanceData.AddressLineSix)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineSeven); err != nil {
		return fieldError("AddressLineSeven", err, rr.RemittanceData.AddressLineSeven)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (rr *RelatedRemittance) fieldInclusion() error {
	return nil
}
