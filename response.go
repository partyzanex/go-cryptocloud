package cryptocloud

import (
	"encoding/json"
	"fmt"
)

type response[R any, E error] struct {
	*DetailError
	Status   ResponseStatus  `json:"status"`
	Result   json.RawMessage `json:"result"`
	AllCount int             `json:"all_count,omitempty"`
}

func (r *response[R, E]) getResult() (*R, error) {
	res := new(R)

	if err := json.Unmarshal(r.Result, res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal result: %w", err)
	}

	return res, nil
}

func (r *response[R, E]) getError() error {
	var res E

	if err := json.Unmarshal(r.Result, &res); err != nil {
		return fmt.Errorf("failed to unmarshal error: %w", err)
	}

	return res
}

type ResponseStatus string

const (
	ResponseStatusSuccess ResponseStatus = "success"
	ResponseStatusError   ResponseStatus = "error"
)
