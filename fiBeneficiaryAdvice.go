// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// FIBeneficiaryAdvice is the financial institution beneficiary advice
type FIBeneficiaryAdvice struct {
	// tag
	tag string
	// Advice
	Advice Advice `json:"advice,omitEmpty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIBeneficiaryAdvice returns a new FIBeneficiaryAdvice
func NewFIBeneficiaryAdvice() FIBeneficiaryAdvice {
	fiba := FIBeneficiaryAdvice{
		tag: TagFIBeneficiaryAdvice,
	}
	return fiba
}

// Parse takes the input string and parses the FIBeneficiaryAdvice values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fiba *FIBeneficiaryAdvice) Parse(record string) {
	fiba.tag = record[:6]
	fiba.Advice.AdviceCode = fiba.parseStringField(record[6:9])
	fiba.Advice.LineOne = fiba.parseStringField(record[9:35])
	fiba.Advice.LineTwo = fiba.parseStringField(record[35:68])
	fiba.Advice.LineThree = fiba.parseStringField(record[68:101])
	fiba.Advice.LineFour = fiba.parseStringField(record[101:134])
	fiba.Advice.LineFive = fiba.parseStringField(record[134:167])
	fiba.Advice.LineSix = fiba.parseStringField(record[167:200])
}

// String writes FIBeneficiaryAdvice
func (fiba *FIBeneficiaryAdvice) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(200)
	buf.WriteString(fiba.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on FIBeneficiaryAdvice and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fiba *FIBeneficiaryAdvice) Validate() error {
	if err := fiba.fieldInclusion(); err != nil {
		return err
	}
	if err := fiba.isAdviceCode(fiba.Advice.AdviceCode); err != nil {
		return fieldError("AdviceCode", err, fiba.Advice.AdviceCode)
	}
	if err:= fiba.isAlphanumeric(fiba.Advice.LineOne); err!= nil {
		return fieldError("LineOne", err, fiba.Advice.LineOne)
	}
	if err:= fiba.isAlphanumeric(fiba.Advice.LineTwo); err!= nil {
		return fieldError("LineTwo", err, fiba.Advice.LineTwo)
	}
	if err:= fiba.isAlphanumeric(fiba.Advice.LineThree); err!= nil {
		return fieldError("LineThree", err, fiba.Advice.LineThree)
	}
	if err:= fiba.isAlphanumeric(fiba.Advice.LineFour); err!= nil {
		return fieldError("LineFour", err, fiba.Advice.LineFour)
	}
	if err:= fiba.isAlphanumeric(fiba.Advice.LineFive); err!= nil {
		return fieldError("LineFive", err, fiba.Advice.LineFive)
	}
	if err:= fiba.isAlphanumeric(fiba.Advice.LineSix); err!= nil {
		return fieldError("LineSix", err, fiba.Advice.LineSix)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (fiba *FIBeneficiaryAdvice) fieldInclusion() error {
	return nil
}
