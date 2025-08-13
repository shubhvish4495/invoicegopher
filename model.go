package invoicegopher

// Invoice represents the invoice data structure
type Invoice struct {
	InvoiceNo     string
	Date          string
	From          Contact
	BilledTo      Contact
	Items         []InvoiceItem
	Total         float64
	PaymentMethod string
	Note          string
}

// Contact represents contact information
type Contact struct {
	Name    string
	Address string
	Email   string
	LogoURL string
}

// InvoiceItem represents an item in the invoice
type InvoiceItem struct {
	Description string
	Quantity    int
	Price       float64
	Amount      float64
}

// calculateTotal calculates the total amount for all items
func (inv *Invoice) calculateTotal() {
	total := 0.0
	for _, item := range inv.Items {
		item.Amount = float64(item.Quantity) * item.Price
		total += item.Amount
	}
	inv.Total = total
}
