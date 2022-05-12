// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &BeneficiaryReference{}

// BeneficiaryReference is a reference for the beneficiary
type BeneficiaryReference struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// BeneficiaryReference
	BeneficiaryReference string `json:"beneficiaryReference,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBeneficiaryReference returns a new BeneficiaryReference
func NewBeneficiaryReference(isVariable bool) *BeneficiaryReference {
	br := &BeneficiaryReference{
		tag:              TagBeneficiaryReference,
		isVariableLength: isVariable,
	}
	return br
}

// Parse takes the input string and parses the BeneficiaryReference values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (br *BeneficiaryReference) Parse(record string) (error, int) {
	if utf8.RuneCountInString(record) < 7 {
		return NewTagWrongLengthErr(7, len(record)), 0
	}
	br.tag = record[:6]

	length := 6
	read := 0

	br.BeneficiaryReference, read = br.parseVariableStringField(record[length:], 16)
	length += read

	return nil, length
}

func (br *BeneficiaryReference) UnmarshalJSON(data []byte) error {
	type Alias BeneficiaryReference
	aux := struct {
		*Alias
	}{
		(*Alias)(br),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	br.tag = TagBeneficiaryReference
	return nil
}

// String writes BeneficiaryReference
func (br *BeneficiaryReference) String() string {
	var buf strings.Builder
	buf.Grow(22)
	buf.WriteString(br.tag)
	buf.WriteString(br.BeneficiaryReferenceField())
	return buf.String()
}

// Validate performs WIRE format rule checks on BeneficiaryReference and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (br *BeneficiaryReference) Validate() error {
	if br.tag != TagBeneficiaryReference {
		return fieldError("tag", ErrValidTagForType, br.tag)
	}
	if err := br.isAlphanumeric(br.BeneficiaryReference); err != nil {
		return fieldError("BeneficiaryReference", err, br.BeneficiaryReference)
	}
	return nil
}

// BeneficiaryReferenceField gets a string of the BeneficiaryReference field
func (br *BeneficiaryReference) BeneficiaryReferenceField() string {
	return br.alphaVariableField(br.BeneficiaryReference, 16, br.isVariableLength)
}
