package cryptocloud

import (
	"context"
	"fmt"
	"net/http"

	"resty.dev/v3"
)

type Client struct {
	client *resty.Client
}

const (
	pathCreateInvoice      = "/v2/invoice/create"
	pathCancelInvoice      = "/v2/invoice/merchant/canceled"
	pathGetInvoices        = "/v2/invoice/merchant/list"
	pathGetBalance         = "/v2/merchant/wallet/balance/all"
	pathGetStatistics      = "/v2/invoice/merchant/statistics"
	pathCreateStaticWallet = "/v2/invoice/static/create"
)

func NewClient(options ...Option) (*Client, error) {
	cfg, err := newConfig(options)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: cfg.client,
	}, nil
}

func sendRequest[P, R any, E error](r *resty.Request, method, path string, payload P) (*R, error) {
	result := new(response[R, E])

	_, err := r.
		SetBody(payload).
		SetResult(result).
		Execute(method, path)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	switch result.Status {
	case ResponseStatusSuccess:
		return result.getResult()
	case ResponseStatusError:
		return nil, result.getError()
	default:
		return nil, fmt.Errorf("unknown response status: %s", result.Status)
	}
}

func (c *Client) CreateInvoice(ctx context.Context, invoiceData *InvoiceData) (*InvoiceResult, error) {
	return sendRequest[*InvoiceData, InvoiceResult, CreateInvoiceErrors](
		c.client.R().SetContext(ctx),
		http.MethodPost,
		pathCreateInvoice,
		invoiceData,
	)
}

func (c *Client) CancelInvoice(ctx context.Context, invoiceID string) error {
	_, err := sendRequest[cancelInvoiceRequest, any, ValidateError](
		c.client.R().SetContext(ctx),
		http.MethodPost,
		pathCancelInvoice,
		cancelInvoiceRequest{
			UUID: invoiceID,
		},
	)

	return err
}

func (c *Client) GetInvoices(ctx context.Context, params *GetInvoicesParams) (*[]*Invoice, int, error) {
	result := new(response[[]*Invoice, GetInvoicesErrors])

	_, err := c.client.R().SetContext(ctx).
		SetBody(params).
		SetResult(result).
		Execute(http.MethodPost, pathGetInvoices)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute request: %w", err)
	}

	switch result.Status {
	case ResponseStatusSuccess:
		var res *[]*Invoice

		res, err = result.getResult()

		return res, result.AllCount, err
	case ResponseStatusError:
		return nil, 0, result.getError()
	default:
		return nil, 0, fmt.Errorf("unknown response status: %s", result.Status)
	}
}

func (c *Client) GetInvoiceInfo(ctx context.Context, invoiceID string) (*Invoice, error) {
	return sendRequest[getInvoiceInfoRequest, Invoice, ValidateError](
		c.client.R().SetContext(ctx),
		http.MethodPost,
		pathCreateInvoice,
		getInvoiceInfoRequest{
			UUID: invoiceID,
		},
	)
}

func (c *Client) GetBalance(ctx context.Context) (*[]*Balance, error) {
	return sendRequest[any, []*Balance, ValidateError](
		c.client.R().SetContext(ctx),
		http.MethodPost,
		pathGetBalance,
		nil,
	)
}

func (c *Client) GetStatistics(ctx context.Context, start, end Date) (*Statistics, error) {
	return sendRequest[GetStatisticsRequest, Statistics, DetailError](
		c.client.R().SetContext(ctx),
		http.MethodPost,
		pathGetStatistics,
		GetStatisticsRequest{
			Start: start,
			End:   end,
		},
	)
}

func (c *Client) GreateStaticWallet(ctx context.Context, shopID, currency, orderID string) (*Wallet, error) {
	return sendRequest[CreateStaticWalletRequest, Wallet, ValidateError](
		c.client.R().SetContext(ctx),
		http.MethodPost,
		pathCreateStaticWallet,
		CreateStaticWalletRequest{
			ShopID:   shopID,
			Currency: currency,
			Identity: orderID,
		},
	)
}
