# Invoice Gopher

<img src="logo.png" alt="Invoice Gopher Logo" width="200" height="150">

<br/>

A Go package for generating professional HTML invoices with a modern, responsive design.

## Features

- Generate HTML invoices from Go structs
- Professional, mobile-responsive design using Tailwind CSS
- Print-optimized styling
- Support for company logos
- Customizable invoice items with automatic total calculation
- Template-based generation for easy customization

## Installation

```bash
go get github.com/shubhvish4495/invoicegopher
```

## Usage

### Basic Usage

```go
package main

import (
    "log"
    invoiceGen "github.com/shubhvish4495/invoicegopher"
)

func main() {
    // Generate a sample invoice
    invoice := invoiceGen.GenerateDummyInvoice()
    
    // Generate HTML file
    err := invoiceGen.GenerateInvoice(invoice, "output")
    if err != nil {
        log.Printf("Error generating invoice: %v", err)
        return
    }
}
```

### Custom Invoice

```go
package main

import (
    "time"
    invoiceGen "github.com/shubhvish4495/invoicegopher"
)

func main() {
    invoice := invoiceGen.Invoice{
        InvoiceNo: "INV-001",
        Date:      time.Now().Format("02 January, 2006"),
        From: invoiceGen.Contact{
            Name:    "Your Company",
            Address: "123 Business St., Business City",
            Email:   "contact@yourcompany.com",
            LogoURL: "https://your-logo-url.com/logo.png",
        },
        BilledTo: invoiceGen.Contact{
            Name:    "Client Name",
            Address: "456 Client Ave., Client City",
            Email:   "client@example.com",
        },
        Items: []invoiceGen.InvoiceItem{
            {
                Description: "Web Development",
                Quantity:    1,
                Price:       1500.00,
            },
            {
                Description: "Consulting Hours",
                Quantity:    10,
                Price:       100.00,
            },
        },
        PaymentMethod: "Bank Transfer",
        Note:          "Thank you for your business!",
    }
    
    // Calculate totals
    for i := range invoice.Items {
        invoice.Items[i].Amount = float64(invoice.Items[i].Quantity) * invoice.Items[i].Price
    }
    invoice.CalculateTotal()
    
    // Generate invoice
    err := invoiceGen.GenerateInvoice(invoice, "output")
    if err != nil {
        log.Printf("Error: %v", err)
    }
}
```

## Data Structures

### Invoice
```go
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
```

### Contact
```go
type Contact struct {
    Name    string
    Address string
    Email   string
    LogoURL string
}
```

### InvoiceItem
```go
type InvoiceItem struct {
    Description string
    Quantity    int
    Price       float64
    Amount      float64
}
```

## Template Customization

The invoice template (`invoice_template.html`) uses:
- Tailwind CSS for styling
- Go templates for data binding
- Responsive design that works on mobile and desktop
- Print-optimized CSS for professional printing

To customize the template, modify the `invoice_template.html` file in your project root.

## Output

The generated invoice is saved as an HTML file that can be:
- Opened in any web browser
- Printed directly from the browser
- Converted to PDF using browser print functionality

## Project Structure

```
.
├── go.mod
├── main.go              # Core invoice generation functions
├── model.go             # Data structures and models
├── invoice_template.html # HTML template for invoice
├── sample/
│   └── main.go          # Example usage
└── output/
    └── invoice.html     # Generated invoice output
```

## Requirements

- Go 1.24.4 or later
- Internet connection (for Tailwind CSS CDN and Google Fonts)

## License

This project is open source and available under the MIT License.