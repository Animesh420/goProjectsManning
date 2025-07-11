package main

import (
	"flag"
	"fmt"
	"learngo-pockets/moneyconverter/ecbank"
	"learngo-pockets/moneyconverter/money"
	"os"
)

func main() {

	from := flag.String("from", "", "source currency, required")
	to := flag.String("to", "EUR", "target currency")

	flag.Parse()

	value := flag.Arg(0)

	if value == "" {
		_, _ = fmt.Fprintln(
			os.Stderr, "missing amount to convert")
		flag.Usage()
		os.Exit(1)
	}

	// fmt.Println(*from, *to, value)

	fromCurrency, err := money.ParseCurrency(*from)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse source currency %q: %s.\n", *from, err.Error())
		os.Exit(1)
	}

	toCurrency, err := money.ParseCurrency(*to)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse source currency %q: %s.\n", *to, err.Error())
		os.Exit(1)
	}

	quantity, err := money.ParseDecimal(value)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse value %q: %s.\n", value, err.Error())
		os.Exit(1)
	}

	amount, err := money.NewAmount(quantity, fromCurrency)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	rates := ecbank.Client{}

	convertedAmount, err := money.Convert(amount, toCurrency, rates)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to convert %s to %s: %s.\n", amount, toCurrency, err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s = %s\n", amount, convertedAmount)
}
