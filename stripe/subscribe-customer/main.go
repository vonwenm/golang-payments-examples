package main

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/sub"
	"os"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_PRIVATE_KEY")

	customers := make(map[string]*stripe.Customer)
	i := customer.List(nil)
	for i.Next() {
		c := i.Customer()
		customers[c.Email] = c
	}

	if john, ok := customers["john.doe@example.com"]; ok {

		// create a subscription
		sub.New(&stripe.SubParams{
			Customer: john.ID,
			Plan:     "Premium Annual",
			Card: &stripe.CardParams{
				Number: "4242424242424242",
				Month:  "12",
				Year:   "2016",
				CVC:    "123",
			},
		})
	}

}
