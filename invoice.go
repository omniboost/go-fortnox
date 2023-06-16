package fortnox

import (
	"github.com/omniboost/go-fortnox/utils"
	"gopkg.in/guregu/null.v3"
)

type Invoices []Invoice

type Invoice struct {
	URL                       utils.URL         `json:"@url,omitempty"`                   // Direct url to the record.
	UrlTaxReductionList       string            `json:"@UrlTaxReductionList,omitempty"`   // Direct url to the tax reduction for the invoice. This is visible even if no tax reduction exists.
	AccountingMethod          string            `json:"AccountingMethod"`                 // Accounting Method. Can be ACCRUAL orCASH
	AdministrationFee         float64           `json:"AdministrationFee"`                // The invoice administration fee.
	AdministrationFeeVAT      float64           `json:"AdministrationFeeVAT,omitempty"`   // VAT of the invoice administration fee.
	Address1                  string            `json:"Address1"`                         // Invoice address 1.
	Address2                  string            `json:"Address2"`                         // Invoice address 2.
	Balance                   float64           `json:"Balance,omitempty"`                // Balance of the invoice.
	BasisTaxReduction         float64           `json:"BasisTaxReduction,omitempty"`      // Basis of tax reduction.
	Booked                    null.Bool         `json:"Booked"`                           // If the invoice is bookkept. This value can be changed by using the action “bookkeep”.
	Cancelled                 null.Bool         `json:"Cancelled,omitempty"`              // If the invoice is cancelled. This value can be changed by using the action “cancel”.
	Credit                    string            `json:"Credit,omitempty"`                 // If the invoice is a credit invoice.
	CreditInvoiceReference    string            `json:"CreditInvoiceReference,omitempty"` // Reference to the credit invoice, if one exits. The reference must be a document number for an existing credit invoice.
	City                      string            `json:"City,omitempty"`                   // City for the invoice address.
	Comments                  string            `json:"Comments"`                         // Comments of the invoice
	ContractReference         string               `json:"ContractReference,omitempty"`      // Reference to the contract, if one exists.
	ContributionPercent       float64           `json:"ContributionPercent,omitempty"`    // Invoice contribution in percent.
	ContributionValue         float64           `json:"ContributionValue,omitempty"`      // Invoice contribution in amount.
	Country                   string            `json:"Country"`                          // Country for the invoice address. Must be a name of an existing country.
	CostCenter                string            `json:"CostCenter,omitempty"`             // Code of the cost center. The code must be of an existing cost center.
	Currency                  string            `json:"Currency"`                         // Code of the currency. The code must be of an existing currency.
	CurrencyRate              Number            `json:"CurrencyRate"`                     // Currency rate used for the invoice.
	CurrencyUnit              float64           `json:"CurrencyUnit"`                     // Currency unit used for the invoice.
	CustomerName              string            `json:"CustomerName"`                     // Name of the customer.
	CustomerNumber            string            `json:"CustomerNumber"`                   // Customer number of the customer. The customer number must be of an existing customer.
	DeliveryAddress1          string            `json:"DeliveryAddress1"`                 // Invoice delivery address 1.
	DeliveryAddress2          string            `json:"DeliveryAddress2"`                 // Invoice delivery address 2.
	DeliveryCity              string            `json:"DeliveryCity"`                     // City for the invoice delivery address.
	DeliveryCountry           string            `json:"DeliveryCountry"`                  // Country for the invoice delivery address. Must be a name of an existing country.
	DeliveryDate              Date              `json:"DeliveryDate"`                     // Date of delivery. Must be a valid date according to our date format.
	DeliveryName              string            `json:"DeliveryName"`                     // Name of the recipient of the delivery
	DeliveryZipCode           string            `json:"DeliveryZipCode"`                  // ZipCode for the invoice delivery address.
	DocumentNumber            string            `json:"DocumentNumber"`                   // The invoice number. If no document number is provided, the next number in the series will be used.
	DueDate                   Date              `json:"DueDate"`                          // Due date of the invoice. Must be a valid date according to our date format.
	EDIInformation            *EDIInformation   `json:"EDIInformation,omitempty"`         // The properties for this object is listed in the table for “EDI Information”.
	EmailInformation          *EmailInformation `json:"EmailInformation,omitempty"`       // The properties for this object is listed in the table for “Email Information”.
	EUQuarterlyReport         bool              `json:"EUQuarterlyReport"`                // EU Quarterly Report On / Off
	ExternalInvoiceReference1 string            `json:"ExternalInvoiceReference1"`        // External invoice reference 1.
	ExternalInvoiceReference2 string            `json:"ExternalInvoiceReference2"`        // External invoice reference 1.
	FinalPayDate              Date              `json:"FinalPayDate,omitempty"`           // The date when the invoice became fully paid. Only available if the invoice is fully paid.
	Freight                   float64           `json:"Freight"`                          // Freight cost of the invoice.
	FreightVAT                float64           `json:"FreightVAT,omitempty"`             // VAT of the freight cost.
	Gross                     float64           `json:"Gross,omitempty"`                  // Gross value of the invoice
	HouseWork                 null.Bool         `json:"HouseWork,omitempty"`              // If there is any row of the invoice marked “house work”.
	InvoiceDate               Date              `json:"InvoiceDate"`                      // Invoice date. Must be a valid date according to our date format.
	InvoicePeriodStart        Date              `json:"InvoicePeriodStart,omitempty"`     // Start date of the invoice period, only applicable for contract invoices.
	InvoicePeriodEnd          Date              `json:"InvoicePeriodEnd,omitempty"`       // End date of the invoice period, only applicable for contract invoices.
	InvoiceReference          string            `json:"InvoiceReference,omitempty"`
	InvoiceRows               InvoiceRows       `json:"InvoiceRows"`                  // The properties for the object in this array is listed in the table “Invoice Rows”.
	InvoiceType               InvoiceType       `json:"InvoiceType"`                  // The type of invoice. Can be INVOICE AGREEMENTINVOICE INTRESTINVOICE SUMMARYINVOICE or CASHINVOICE.
	Labels                    Labels            `json:"Labels"`                       // The properties for the object in this array is listed in the table “Labels”
	Language                  string            `json:"Language"`                     // Language code. Can be SV or EN.
	LastRemindDate            Date              `json:"LastRemindDate,omitempty"`     // Date of last reminder.
	Net                       float64           `json:"Net,omitempty"`                // Net amount
	NotCompleted              bool              `json:"NotCompleted"`                 // If the invoice is set as not completed.
	NoxFinans                 null.Bool         `json:"NoxFinans,omitempty"`          // If the invoice is managed by NoxFinans
	OCR                       string            `json:"OCR"`                          // OCR number of the invoice.
	OfferReference            string            `json:"OfferReference,omitempty"`     // Reference to the offer, if one exists.
	OrderReference            string            `json:"OrderReference,omitempty"`     // Reference to the order, if one exists.
	OrganisationNumber        string            `json:"OrganisationNumber,omitempty"` // Organisation number of the customer for the invoice.
	OurReference              string            `json:"OurReference"`                 // Our reference.
	PaymentWay                PaymentWay        `json:"PaymentWay,omitempty"`         // CASH, CARD, AG
	Phone1                    string            `json:"Phone1"`                       // Phone number 1 of the customer for the invoice.
	Phone2                    string            `json:"Phone2"`                       // Phone number 2 of the customer for the invoice.
	PriceList                 string            `json:"PriceList"`                    // Code of the price list. The code must be of an existing price list.
	PrintTemplate             string            `json:"PrintTemplate"`                // Print template of the invoice. Must be an existing print template.
	Project                   string            `json:"Project"`                      // Code of the project. The code must be of an existing project.
	WarehouseReady            null.Bool         `json:"WarehouseReady,omitempty"`
	OutboundDate              Date              `json:"OutboundDate"`
	Remarks                   string            `json:"Remarks"`                // Remarks of the invoice. This is the invoice text shown on the invoice.
	Reminders                 int               `json:"Reminders,omitempty"`    // Number of reminders sent to the customer.
	RoundOff                  float64           `json:"RoundOff,omitempty"`     // Round off amount for the invoice.
	Sent                      null.Bool         `json:"Sent,omitempty"`         // If the document is printed or sent in any way.
	TaxReduction              int               `json:"TaxReduction,omitempty"` // The amount of tax reduction.
	TermsOfDelivery           string            `json:"TermsOfDelivery"`        // Code of the terms of delivery. The code must be of an existing terms of delivery.
	TermsOfPayment            string            `json:"TermsOfPayment"`         // Code of the terms of payment. The code must be of an existing terms of payment.
	Total                     float64           `json:"Total,omitempty"`        // The total amount of the invoice.
	TotalVAT                  float64           `json:"TotalVAT,omitempty"`     // The total VAT amount of the invoice.
	TotalToPay                float64           `json:"TotalToPay,omitempty"`
	VATIncluded               bool              `json:"VATIncluded"`             // If the price of the invoice is including VAT.
	VoucherNumber             int               `json:"VoucherNumber,omitempty"` // Voucher number for the invoice. This is created when the invoice is bookkept.
	VoucherSeries             string            `json:"VoucherSeries,omitempty"` // Voucher series for the invoice. This is created when the invoice is bookkept.
	VoucherYear               int               `json:"VoucherYear,omitempty"`   // Voucher year for the invoice. This is created when the invoice is bookkept.
	WayOfDelivery             string            `json:"WayOfDelivery,omitempty"` // Code of the way of delivery. The code must be of an existing way of delivery.
	YourOrderNumber           string            `json:"YourOrderNumber"`         // Your order number.
	YourReference             string            `json:"YourReference"`           // Your reference.
	ZipCode                   string            `json:"ZipCode"`                 // Zip code of the invoice.
}

type EDIInformation struct {
	EDIGlobalLocationNumber         string    `json:"EDIGlobalLocationNumber"`         // Invoice address GLN for EDI
	EDIGlobalLocationNumberDelivery string    `json:"EDIGlobalLocationNumberDelivery"` // Delivery address GLN for EDI
	EDIInvoiceExtra1                string    `json:"EDIInvoiceExtra1"`                // Extra EDI Information
	EDIInvoiceExtra2                string    `json:"EDIInvoiceExtra2"`                // Extra EDI Information
	EDIOurElectronicReference       string    `json:"EDIOurElectronicReference"`       // Our electronic reference for EDI
	EDIYourElectronicReference      string    `json:"EDIYourElectronicReference"`      // Your electronic reference for EDI
	EDIStatus                       EDIStatus `json:"EDIStatus,omitempty"`             // Status of the send process of the invoice. Can have the following codes:
}

type EDIStatus string

var (
	EDIStatusSentToCrediflow                 EDIStatus = "1"
	EDIStatusCheckedByCrediflow              EDIStatus = "2"
	EDIStatusDeliveredElectronically         EDIStatus = "3"
	EDIStatusDeliveredToPrintingService      EDIStatus = "4"
	EDIStatusDeliveredToPostenByPrintCompany EDIStatus = "5"
	EDIStatusDeclinedByCrediflow             EDIStatus = "6"
	EDIStatusDeclinedByReceiver              EDIStatus = "7"
)

type EmailInformation struct {
	EmailAddressFrom string `json:"EmailAddressFrom"` // Reply to adress. Must be a valid e-mail address
	EmailAddressTo   string `json:"EmailAddressTo"`   // Customer e-mail address. Must be a valid e-mail address.
	EmailAddressCC   string `json:"EmailAddressCC"`   // Customer e-mail address – Copy. Must be a valid e-mail address.
	EmailAddressBCC  string `json:"EmailAddressBCC"`  // Customer e-mail address – Blind carbon copy. Must be a valid e-mail address.
	EmailSubject     string `json:"EmailSubject"`     // Subject of the e-mail
	// The variable {no} = document number. The variable {name} =  customer name
	EmailBody string `json:"EmailBody"` // Body of the e-mail.
}

type InvoiceRows []InvoiceRow

type InvoiceRow struct {
	AccountNumber       int     `json:"AccountNumber"`                 // Account number. If not provided the predefined account will be used.
	ArticleNumber       string  `json:"ArticleNumber"`                 // Article number.
	ContributionPercent string  `json:"ContributionPercent,omitempty"` // Contribution Percent.
	ContributionValue   string  `json:"ContributionValue,omitempty"`   // Contribution Value.
	CostCenter          string  `json:"CostCenter,omitempty"`          // Code of the cost center for the row. The code must be of an existing cost center.
	DeliveredQuantity   string  `json:"DeliveredQuantity"`             // Delivered quantity.
	Description         string  `json:"Description"`                   // Row description.
	Discount            float64 `json:"Discount"`                      // Discount amount.
	// 12 digits (for amount) / 5 digits (for percent)
	DiscountType           string `json:"DiscountType"`           // The type of discount used for the row. Can be AMOUNT or PERCENT.
	HouseWork              bool   `json:"HouseWork"`              // If the row is housework
	HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"` // Hours to be reported if the quantity of the row should not be used as hours.
	HouseWorkType          string `json:"HouseWorkType"`          // The type of house work. Can be
	// CONSTRUCTIONELECTRICITYGLASSMETALWORK
	// GROUNDDRAINAGEWORKMASONRY
	// PAINTINGWALLPAPERINGHVACMAJORAPPLIANCEREPAIR
	// MOVINGSERVICESITSERVICESCLEANINGTEXTILECLOTHINGSNOWPLOWING
	// GARDENINGBABYSITTINGOTHERCAREOTHERCOSTS or empty.
	Price             float64 `json:"Price"`                       // Price per unit
	PriceExcludingVAT float64 `json:"PriceExcludingVAT,omitempty"` // Price per unit excluding VAT (regardless of value of VATIncluded flag)
	Project           string  `json:"Project,omitempty"`           // Code of the project for the row. The code must be of an existing project.
	StockPointCode    string  `json:"StockPointCode,omitempty"`
	Total             float64 `json:"Total,omitempty"`             // Total amount for the row.
	TotalExcludingVAT float64 `json:"TotalExcludingVAT,omitempty"` // Total amount for the row excluding VAT (regardless of value of VATIncluded flag)
	Unit              string  `json:"Unit"`                        // Code of the unit for the row. The code must be of an existing unit.
	VAT               int     `json:"VAT"`                         // VAT percentage of the row. The percentage needs to be of an existing VAT percentage.
}

type InvoiceType string

var (
	InvoiceTypeInvoice          InvoiceType = "INVOICE"
	InvoiceTypeAgreementInvoice InvoiceType = "AGREEMENTINVOICE"
	InvoiceTypeIntrestInvoice   InvoiceType = "INTRESTINVOICE"
	InvoiceTypeSummaryInvoice   InvoiceType = "SUMMARYINVOICE"
	InvoiceTypeCashInvoice      InvoiceType = "CASHINVOICE"
)

type Labels []Label

type Label struct {
	ID string `json:"Id"` // Id of the label.
}

type PaymentWay string

var (
	PaymentWayCash PaymentWay = "CASH" // Payment Way is set to Cash payment
	PaymentWayCard PaymentWay = "Card" // Payment Way is set to Card.
	PaymentWayAG   PaymentWay = "AG"   // Payment Way is set to Direct debit
)
