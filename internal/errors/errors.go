package errors

import (
	"errors"
	"fmt"
	"strings"
)

// Error list.
var (
	ErrInvalidAuthPartner              = errors.New("invalid auth partner")
	ErrInvalidCacheType                = errors.New("invalid cache type")
	ErrInvalidDBFormat                 = errors.New("invalid db address")
	ErrInvalidRequestFormat            = errors.New("invalid request format")
	ErrInvalidRequestData              = errors.New("invalid request data")
	ErrInvalidPaymentGateway           = errors.New("invalid payment gateway")
	ErrInvalidDisbursementGateway      = errors.New("invalid disbursement gateway")
	ErrInvalidPaymentChannel           = errors.New("invalid payment channel")
	ErrInvalidToken                    = errors.New("invalid token")
	ErrInvalidKey                      = errors.New("invalid key")
	ErrAlreadyCancelled                = errors.New("already cancelled")
	ErrAlreadyPaid                     = errors.New("already paid")
	ErrAlreadyDone                     = errors.New("already done")
	ErrInvalidID                       = errors.New("invalid id")
	ErrInvalidExternalID               = errors.New("invalid external id")
	ErrInvalidAmount                   = errors.New("invalid amount")
	ErrInvalidAccountNumber            = errors.New("invalid account number")
	ErrInvalidAccountName              = errors.New("invalid account name")
	ErrInvalidBankCode                 = errors.New("invalid bank code")
	ErrInvalidBankAccount              = errors.New("invalid bank account")
	ErrInvalidSignatureKey             = errors.New("invalid signature key")
	ErrInvalidStatus                   = errors.New("invalid status")
	ErrInvalidMessageType              = errors.New("invalid message type")
	ErrInvalidPartnerID                = errors.New("invalid partner id")
	ErrInvalidPaymentCode              = errors.New("invalid payment code")
	ErrInvalidPaymentMethodStatus      = errors.New("invalid payment method status")
	ErrInvalidDisbursementMethodStatus = errors.New("invalid disbursement method status")
	ErrDuplicateTransactionID          = errors.New("duplicate transaction id")
	ErrExistPartnerPaymentMethod       = errors.New("partner payment method already exist")
	ErrExistPartnerDisbursementMethod  = errors.New("partner disbursement method already exist")
	ErrNotFoundPartner                 = errors.New("partner not found")
	ErrNotFoundPayment                 = errors.New("payment not found")
	ErrNotFoundPaymentMethod           = errors.New("payment method not found")
	ErrNotFoundDisbursement            = errors.New("disbursement not found")
	ErrNotFoundDisbursementMethod      = errors.New("disbursement method not found")
	ErrNotFoundQRCode                  = errors.New("qr code not found")
	ErrNotActivePaymentMethod          = errors.New("payment method is not active")
	ErrMaintenancePaymentMethod        = errors.New("payment method is in maintenance")
	ErrNotActiveDisbursementMethod     = errors.New("disbursement method is not active")
	ErrMaintenanceDisbursementMethod   = errors.New("disbursement method is in maintenance")
	ErrCardNotRegistered               = errors.New("card is not registered")
	ErrInternalDB                      = errors.New("internal database error")
	ErrInternalCache                   = errors.New("internal cache error")
	ErrInternalServer                  = errors.New("internal server error")
	ErrXenditVAInactive                = errors.New("virtual account has been paid or already expired")
	ErrXenditRetailInactive            = errors.New("payment code has been paid or already expired")
	ErrXenditQRInactive                = errors.New("qr code has been paid")
	ErrCCExpired                       = errors.New("card is expired")
	ErrCCDeclined                      = errors.New("card is declined")
	ErrCCInsufficient                  = errors.New("card does not have enough balance")
	ErrCCStolen                        = errors.New("card is marked as stolen")
	ErrCCInactive                      = errors.New("card is inactive")
	ErrCCAmount                        = errors.New("amount is below minimum limit or above maximum limit")
	ErrCCFraud                         = errors.New("card is detected as fraud")
	ErrBlacklistBankAccount            = errors.New("bank account is blacklisted")
	ErrNoInquiryBank                   = errors.New("no inquiry feature for this bank")
	ErrInvalidOVONumber                = errors.New("invalid ovo number")
	ErrVANotFound                      = errors.New("invalid virtual account number")
	ErrBillNotFound                    = errors.New("bill not found")
	ErrInvalidBillerCode               = errors.New("invalid company code")
)

// ErrRequiredField is error for missing field.
func ErrRequiredField(str string) error {
	return fmt.Errorf("required field %s", str)
}

// ErrGTField is error for greater than field.
func ErrGTField(str, value string) error {
	return fmt.Errorf("field %s must be greater than %s", str, value)
}

// ErrGTEField is error for greater than or equal field.
func ErrGTEField(str, value string) error {
	return fmt.Errorf("field %s must be greater than or equal %s", str, value)
}

// ErrLTField is error for lower than field.
func ErrLTField(str, value string) error {
	return fmt.Errorf("field %s must be lower than %s", str, value)
}

// ErrLTEField is error for lower than or equal field.
func ErrLTEField(str, value string) error {
	return fmt.Errorf("field %s must be lower than or equal %s", str, value)
}

// ErrLenField is error for length field.
func ErrLenField(str, value string) error {
	return fmt.Errorf("field %s length must be %s", str, value)
}

// ErrISO3166Alpha2Field is error for ISO 3166-1 alpha-2 field.
func ErrISO3166Alpha2Field(str string) error {
	return fmt.Errorf("field %s must be in ISO 3166-1 alpha-2 format", str)
}

// ErrEmailField is error for email field.
func ErrEmailField(str string) error {
	return fmt.Errorf("field %s must be in email format", str)
}

// ErrURLField is error for url field.
func ErrURLField(str string) error {
	return fmt.Errorf("field %s must be in URL format", str)
}

// ErrInvalidFormatField is error for invalid format field.
func ErrInvalidFormatField(str string) error {
	return fmt.Errorf("invalid format field %s", str)
}

// ErrOneOfField is error for oneof field.
func ErrOneOfField(str, value string) error {
	return fmt.Errorf("field %s must be one of %s", str, strings.Join(strings.Split(value, " "), "/"))
}

// ErrNumericField is error for numeric field.
func ErrNumericField(str string) error {
	return fmt.Errorf("field %s must contain number only", str)
}

// ErrAlphaField is error for alpha field.
func ErrAlphaField(str string) error {
	return fmt.Errorf("field %s must contain letter only", str)
}

// ErrMinPrice is error for min price.
func ErrMinPrice(price float64) error {
	return fmt.Errorf("minimum price is %d", int(price))
}

// ErrMaxPrice is error for max price.
func ErrMaxPrice(price float64) error {
	return fmt.Errorf("maximum price is %d", int(price))
}

// ErrMinAmount is error for min amount.
func ErrMinAmount(amount float64) error {
	return fmt.Errorf("minimum amount is %d", int(amount))
}

// ErrMaxAmount is error for max amount.
func ErrMaxAmount(amount float64) error {
	return fmt.Errorf("maximum amount is %d", int(amount))
}
