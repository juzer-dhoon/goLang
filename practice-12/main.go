package main

import (
	"fmt"
)

type ProductDetails struct {
	sno   int
	name  string
	price float64
}

type Purchase struct {
	sno      int
	name     string
	price    float64
	quantity int
	total    float64
}

func main() {

	products := []ProductDetails{
		{1, "Pen", 20},
		{2, "Notebook", 30},
		{3, "Eraser", 5},
		{4, "Pencil", 10},
		{5, "Marker", 25},
	}

	var purchases []Purchase
	var choice int

	for {
		// Show product list
		fmt.Println("\n----Inventory Management System-----")
		fmt.Println("\nAvailable Products:")
		for _, p := range products {
			fmt.Printf("%d. %s - %.2f\n", p.sno, p.name, p.price)
		}
		fmt.Println("6. Exit")

		fmt.Print("Select a product by number (1-6): ")
		fmt.Scan(&choice)

		if choice == 6 {
			break // Exit loop
		}

		if choice < 1 || choice > 5 {
			fmt.Println("Invalid selection. Try again.")
			continue
		}

		var quantity int
		selected := products[choice-1]
		fmt.Printf("Enter quantity for %s: ", selected.name)
		fmt.Scan(&quantity)

		total := selected.price * float64(quantity)

		found := false
		for i, item := range purchases {
			if item.sno == selected.sno {
				// Update existing item
				purchases[i].quantity += quantity
				purchases[i].total = purchases[i].price * float64(purchases[i].quantity)
				found = true
				break
			}
		}

		if !found {
			// Add new item
			purchases = append(purchases, Purchase{
				sno:      selected.sno,
				name:     selected.name,
				price:    selected.price,
				quantity: quantity,
				total:    total,
			})
		}

		var more string
		fmt.Print("Do you want to add more products? (yes/no): ")
		fmt.Scan(&more)

		if more != "yes" {
			break
		}
	}

	// Final Bill
	if len(purchases) > 0 {
		fmt.Println("\n----- BILL -----")
		fmt.Println("Bill No: JD3567  Current Date: May 26, 2026")
		fmt.Printf("%-10s %-10s %-10s %-10s %-10s\n", "Sr.No", "Product", "Price", "Qty", "Total")
		var finalTotal float64
		for i, item := range purchases {
			fmt.Printf("%-10d %-10s %-9.2f %-10d %-9.2f\n",
				i+1, item.name, item.price, item.quantity, item.total)
			finalTotal += item.total
		}
		fmt.Printf("\nFinal Total: %.2f\n", finalTotal)
	} else {
		fmt.Println("No items were purchased.")
		fmt.Println("\n----Thank you for using our Inventory Management System-----")
	}
}
