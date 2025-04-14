package cryptocloud

import (
	"errors"
	"fmt"
)

var (
	ErrRequiredAPIKey = errors.New("api-key is required")
	ErrRequiredAPIUrl = errors.New("api-url is required")
)

type DetailError struct {
	Detail *string `json:"detail,omitempty"`
}

func (e DetailError) Error() string {
	if e.Detail != nil {
		return fmt.Sprintf("detail: %s", *e.Detail)
	}

	return ""
}

type ValidateError struct {
	DetailError
	ValidateError *string `json:"validate_error,omitempty"`
}

func (e ValidateError) Error() string {
	errs := make([]error, 0, 2)

	if e.Detail != nil {
		errs = append(errs, e.DetailError)
	}

	if e.ValidateError != nil {
		errs = append(errs, fmt.Errorf("validate_error: %s", *e.ValidateError))
	}

	return errors.Join(errs...).Error()
}

type CreateInvoiceErrors struct {
	Amount   *string `json:"amount,omitempty"`
	ShopID   *string `json:"shop_id,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

func (e CreateInvoiceErrors) Error() string {
	errs := make([]error, 0, 3)

	if e.Amount != nil {
		errs = append(errs, fmt.Errorf("amount: %s", *e.Amount))
	}

	if e.ShopID != nil {
		errs = append(errs, fmt.Errorf("shop_id: %s", *e.ShopID))
	}

	if e.Currency != nil {
		errs = append(errs, fmt.Errorf("currency: %s", *e.Currency))
	}

	return errors.Join(errs...).Error()
}

type GetInvoicesErrors struct {
	ValidateError
	Start *string `json:"start,omitempty"`
	End   *string `json:"end,omitempty"`
}

func (e GetInvoicesErrors) Error() string {
	errs := make([]error, 0, 3)

	if e.Start != nil {
		errs = append(errs, fmt.Errorf("start: %s", *e.Start))
	}

	if e.End != nil {
		errs = append(errs, fmt.Errorf("end: %s", *e.End))
	}

	if e.Detail != nil {
		errs = append(errs, e.DetailError)
	}

	return errors.Join(errs...).Error()
}
