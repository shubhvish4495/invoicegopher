package invoicegopher

import (
	"reflect"
	"testing"
	"time"
)

func TestGenerateDummyInvoice(t *testing.T) {
	expectedInvoice := GenerateDummyInvoice()

	actualInvoice := Invoice{
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
	for i := range actualInvoice.Items {
		actualInvoice.Items[i].Amount = float64(actualInvoice.Items[i].Quantity) * actualInvoice.Items[i].Price
	}
	actualInvoice.calculateTotal()

	if !reflect.DeepEqual(expectedInvoice, actualInvoice) {
		t.Fail()
	}
}

func TestGenerateInvoice(t *testing.T) {
	d, err := GenerateInvoice(GenerateDummyInvoice())
	if err != nil || d == nil {
		t.Fail()
	}
}
