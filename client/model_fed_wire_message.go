/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// FedWireMessage
type FedWireMessage struct {
	// FedWireMessage ID
	Id string `json:"id,omitempty"`
	MessageDisposition MessageDisposition `json:"messageDisposition,omitempty"`
	ReceiptTimeStamp ReceiptTimeStamp `json:"receiptTimeStamp,omitempty"`
	OutputMessageAccountabilityData OutputMessageAccountabilityData `json:"outputMessageAccountabilityData,omitempty"`
	ErrorWire ErrorWire `json:"errorWire,omitempty"`
	SenderSupplied SenderSupplied `json:"senderSupplied"`
	TypeSubType TypeSubType `json:"typeSubType"`
	InputMessageAccountabilityData InputMessageAccountabilityData `json:"inputMessageAccountabilityData"`
	Amount Amount `json:"amount"`
	SenderDepositoryInstitution SenderDepositoryInstitution `json:"senderDepositoryInstitution"`
	ReceiverDepositoryInstitution ReceiverDepositoryInstitution `json:"receiverDepositoryInstitution"`
	BusinessFunctionCode BusinessFunctionCode `json:"businessFunctionCode"`
	SenderReference SenderReference `json:"senderReference,omitempty"`
	PreviousMessageIdentifier PreviousMessageIdentifier `json:"previousMessageIdentifier,omitempty"`
	LocalInstrument LocalInstrument `json:"localInstrument,omitempty"`
	PaymentNotification PaymentNotification `json:"paymentNotification,omitempty"`
	Charges Charges `json:"charges,omitempty"`
	InstructedAmount InstructedAmount `json:"instructedAmount,omitempty"`
	ExchangeRate ExchangeRate `json:"exchangeRate,omitempty"`
	BeneficiaryIntermediaryFI FinancialInstitution `json:"beneficiaryIntermediaryFI,omitempty"`
	BeneficiaryFI FinancialInstitution `json:"beneficiaryFI,omitempty"`
	Beneficiary Personal `json:"beneficiary,omitempty"`
	BeneficiaryReference BeneficiaryReference `json:"beneficiaryReference,omitempty"`
	AccountDebitedDrawdown AccountDebitedDrawdown `json:"accountDebitedDrawdown,omitempty"`
	Originator Personal `json:"originator,omitempty"`
	OriginatorOptionF OriginatorOptionF `json:"originatorOptionF,omitempty"`
	OriginatorFI FinancialInstitution `json:"originatorFI,omitempty"`
	InstructingFI FinancialInstitution `json:"instructingFI,omitempty"`
	AccountCreditedDrawdown AccountCreditedDrawdown `json:"accountCreditedDrawdown,omitempty"`
	OriginatorToBeneficiary OriginatorToBeneficiary `json:"originatorToBeneficiary,omitempty"`
	FiReceiverFI FiToFi `json:"fiReceiverFI,omitempty"`
	FiDrawdownDebitAccountAdvice Advice `json:"fiDrawdownDebitAccountAdvice,omitempty"`
	FiIntermediaryFI FiToFi `json:"fiIntermediaryFI,omitempty"`
	FiIntermediaryFIAdvice Advice `json:"fiIntermediaryFIAdvice,omitempty"`
	FiBeneficiaryFI FiToFi `json:"fiBeneficiaryFI,omitempty"`
	FiBeneficiaryFIAdvice Advice `json:"fiBeneficiaryFIAdvice,omitempty"`
	FiBeneficiary FiToFi `json:"fiBeneficiary,omitempty"`
	FiBeneficiaryAdvice Advice `json:"fiBeneficiaryAdvice,omitempty"`
	FiPaymentMethodToBeneficiary FiPaymentMethodToBeneficiary `json:"fiPaymentMethodToBeneficiary,omitempty"`
	FiAdditionalFIToFI AdditionalFiToFi `json:"fiAdditionalFIToFI,omitempty"`
	CurrencyInstructedAmount CurrencyInstructedAmount `json:"currencyInstructedAmount,omitempty"`
	OrderingCustomer CoverPayment `json:"orderingCustomer,omitempty"`
	OrderingInstitution CoverPayment `json:"orderingInstitution,omitempty"`
	IntermediaryInstitution CoverPayment `json:"intermediaryInstitution,omitempty"`
	InstitutionAccount CoverPayment `json:"institutionAccount,omitempty"`
	BeneficiaryCustomer CoverPayment `json:"beneficiaryCustomer,omitempty"`
	Remittance CoverPayment `json:"remittance,omitempty"`
	SenderToReceiver CoverPayment `json:"senderToReceiver,omitempty"`
	UnstructuredAddenda UnstructuredAddenda `json:"unstructuredAddenda,omitempty"`
	RelatedRemittance RelatedRemittance `json:"relatedRemittance,omitempty"`
	RemittanceOriginator RemittanceOriginator `json:"remittanceOriginator,omitempty"`
	RemittanceBeneficiary RemittanceBeneficiary `json:"remittanceBeneficiary,omitempty"`
	PrimaryRemittanceDocument PrimaryRemittanceDocument `json:"primaryRemittanceDocument,omitempty"`
	ActualAmountPaid RemittanceAmount `json:"actualAmountPaid,omitempty"`
	GrossAmountRemittanceDocument RemittanceAmount `json:"grossAmountRemittanceDocument,omitempty"`
	AmountNegotiatedDiscount RemittanceAmount `json:"amountNegotiatedDiscount,omitempty"`
	Adjustment Adjustment `json:"adjustment,omitempty"`
	DateRemittanceDocument DateRemittanceDocument `json:"dateRemittanceDocument,omitempty"`
	SecondaryRemittanceDocument SecondaryRemittanceDocument `json:"secondaryRemittanceDocument,omitempty"`
	RemittanceFreeText RemittanceFreeText `json:"remittanceFreeText,omitempty"`
	ServiceMessage ServiceMessage `json:"serviceMessage,omitempty"`
}