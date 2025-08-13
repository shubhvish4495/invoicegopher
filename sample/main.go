package main

import (
	"log"
	"os"

	invoiceGen "github.com/shubhvish4495/invoicegopher"
)

func main() {
	bytes, err := invoiceGen.GenerateInvoice(invoiceGen.GenerateDummyInvoice())
	if err != nil {
		log.Printf("error while creating invoice %v", err)
		return
	}

	f, err := os.Create("output/temp.html")
	if err != nil {
		log.Printf("error while creating file %v", err)
		return
	}

	_, err = f.Write(bytes)
	if err != nil {
		log.Printf("error while writing to file %v", err)
		return
	}
}
