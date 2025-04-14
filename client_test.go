package cryptocloud

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}

type ClientSuite struct {
	suite.Suite

	shopID string
	client *Client
}

func (s *ClientSuite) SetupTest() {
	apiUrl, ok := os.LookupEnv("TEST_API_URL")
	s.Require().True(ok, "TEST_API_URL is not set")

	apiKey, ok := os.LookupEnv("TEST_API_KEY")
	s.Require().True(ok, "TEST_API_KEY is not set")

	shopID, ok := os.LookupEnv("TEST_SHOP_ID")
	s.Require().True(ok, "TEST_SHOP_ID is not set")

	client, err := NewClient(
		WithAPIKey(apiKey),
		WithAPIUrl(apiUrl),
	)
	s.Require().NoError(err)

	s.shopID = shopID
	s.client = client
}

func (s *ClientSuite) TestCreateInvoice() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	result, err := s.client.CreateInvoice(ctx, &InvoiceData{
		ShopID:    s.shopID,
		Amount:    10,
		Currency:  nil,
		OrderID:   nil,
		Email:     nil,
		AddFields: nil,
	})
	s.Require().NoError(err)
	s.NotNil(result)

	err = s.client.CancelInvoice(ctx, result.UUID)
	s.Require().NoError(err)

	result, err = s.client.CreateInvoice(ctx, &InvoiceData{
		ShopID:   s.shopID,
		Amount:   10,
		Currency: EUR.Ptr(),
		OrderID: func() *string {
			orderID := "test-order-id"
			return &orderID
		}(),
		Email:     nil,
		AddFields: nil,
	})
	s.Require().NoError(err)
	s.NotNil(result)

	err = s.client.CancelInvoice(ctx, result.UUID)
	s.Require().NoError(err)

	result, err = s.client.CreateInvoice(ctx, &InvoiceData{
		ShopID:   s.shopID,
		Amount:   100,
		Currency: GBP.Ptr(),
		OrderID: func() *string {
			orderID := fmt.Sprintf("t-%d", time.Now().UnixNano())
			return &orderID
		}(),
		Email: nil,
		AddFields: &InvoiceAddFields{
			TimeToPay: &TimeToPay{
				Hours:   1,
				Minutes: 10,
			},
			EmailToSend:         nil,
			AvailableCurrencies: nil,
			CryptoCurrency:      ETH,
			Period:              PeriodDay,
		},
	})
	s.Require().NoError(err)
	s.NotNil(result)

	err = s.client.CancelInvoice(ctx, result.UUID)
	s.Require().NoError(err)
}

func (s *ClientSuite) TestGetInvoices() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	result, err := s.client.CreateInvoice(ctx, &InvoiceData{
		ShopID: s.shopID,
		Amount: 10,
	})
	s.Require().NoError(err)
	s.NotNil(result)

	now := time.Now()

	invoices, count, err := s.client.GetInvoices(ctx, &GetInvoicesParams{
		Start:  NewDate(now.Year(), now.Month(), now.Day()),
		End:    NewDate(now.Year(), now.Month(), now.Day()),
		Offset: 0,
		Limit:  1,
	})
	s.Require().NoError(err)
	s.Len(*invoices, 1)
	s.Less(1, count)

	err = s.client.CancelInvoice(ctx, result.UUID)
	s.Require().NoError(err)
}
