package invoicegopher

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
)

// GenerateDummyInvoice creates a sample invoice
// should be used only for testing purposes
func GenerateDummyInvoice() Invoice {
	invoice := Invoice{
		InvoiceNo: "000001",
		Date:      time.Now().Format("02 January, 2006"),
		From: Contact{
			Name:    "Olivia Wilson",
			Address: "123 Anywhere St., Any City",
			Email:   "hello@reallygreatsite.com",
			LogoURL: "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png",
		},
		BilledTo: Contact{
			Name:    "Studio Shodwe",
			Address: "123 Anywhere St., Any City",
			Email:   "hello@reallygreatsite.com",
		},
		Items: []InvoiceItem{
			{
				Description: "Logo",
				Quantity:    1,
				Price:       500.00,
			},
			{
				Description: "Banner (2x6m)",
				Quantity:    2,
				Price:       45.00,
			},
			{
				Description: "Poster (1x2m)",
				Quantity:    3,
				Price:       55.00,
			},
		},
		PaymentMethod: "Cash",
		Note:          "Thank you for choosing us!",
	}

	// Calculate amounts for each item and total
	for i := range invoice.Items {
		invoice.Items[i].Amount = float64(invoice.Items[i].Quantity) * invoice.Items[i].Price
	}
	invoice.calculateTotal()

	return invoice
}

// GenerateInvoice generates an invoice from a template
// invoice is the invoice to generate
// returns the invoice as a byte array and an error if any
func GenerateInvoice(invoice Invoice) ([]byte, error) {

	// Parse the HTML template
	tmpl, err := template.ParseFiles("invoice_template.html")
	if err != nil {
		return nil, fmt.Errorf("error parsing template: %v", err)
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, invoice)
	if err != nil {
		return nil, fmt.Errorf("error executing template: %v", err)

	}

	return buf.Bytes(), nil
}
