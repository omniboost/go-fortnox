package fortnox_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	fortnox "github.com/omniboost/go-fortnox"
)

func TestPostInvoice(t *testing.T) {
	req := client.NewPostInvoiceRequest()
	invoice := req.RequestBody().Invoice
	invoice.CurrencyRate = fortnox.Number(1.0)
	invoice.CurrencyUnit = 1.0
	invoice.CustomerName = "Leon Bogaert"
	invoice.CustomerNumber = "1204"
	invoice.InvoiceType = fortnox.InvoiceTypeInvoice
	rows := fortnox.InvoiceRows{
		{
			AccountNumber:     1249,
			Price:             12.0,
			DeliveredQuantity: "1.0",
		},
	}
	invoice.InvoiceRows = append(invoice.InvoiceRows, rows...)
	req.RequestBody().Invoice = invoice
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
