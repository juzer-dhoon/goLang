package main

import (
	"fmt"
)

func main() {
	productSno := []int{1, 2, 3, 4, 5, 6}
	productNames := []string{"Pen", "Notebook", "Eraser", "Pencil", "Marker", "More"}
	productPrices := []float64{20, 30, 5, 10, 25, 30}

	var purchaseSno []int
	var purchaseNames []string
	var purchasePrices []float64
	var purchaseQuantities []int
	var purchaseTotals []float64

	var choice int

	for {
		// Show product list
		fmt.Println("\n----Inventory Management System-----")
		fmt.Println("\nAvailable Products:")
		for i := range productSno {
			fmt.Printf("%d. %s - %.2f\n", productSno[i], productNames[i], productPrices[i])
		}
		fmt.Println("6. Exit")

		fmt.Print("Select a product by number (1-6): ")
		fmt.Scan(&choice)

		if choice == 7 {
			break
		}

		if choice < 1 || choice > 6 {
			fmt.Println("Invalid selection. Try again.")
			continue
		}

		var quantity int
		idx := choice - 1
		fmt.Printf("Enter quantity for %s: ", productNames[idx])
		fmt.Scan(&quantity)

		total := productPrices[idx] * float64(quantity)

		// Check if already purchased
		found := false
		for i := range purchaseSno {
			if purchaseSno[i] == productSno[idx] {
				purchaseQuantities[i] += quantity
				purchaseTotals[i] = purchasePrices[i] * float64(purchaseQuantities[i])
				found = true
				break
			}
		}

		if !found {
			purchaseSno = append(purchaseSno, productSno[idx])
			purchaseNames = append(purchaseNames, productNames[idx])
			purchasePrices = append(purchasePrices, productPrices[idx])
			purchaseQuantities = append(purchaseQuantities, quantity)
			purchaseTotals = append(purchaseTotals, total)
		}

		var more string
		fmt.Print("Do you want to add more products? (yes/no): ")
		fmt.Scan(&more)
		if more != "yes" {
			break
		}
	}

	// Final bill
	if len(purchaseSno) > 0 {
		fmt.Println("\n----- BILL -----")
		fmt.Println("Bill No: JD3567  Current Date: May 26, 2026")
		fmt.Printf("%-10s %-10s %-10s %-10s %-10s\n", "Sr.No", "Product", "Price", "Qty", "Total")
		var finalTotal float64
		for i := range purchaseSno {
			fmt.Printf("%-10d %-10s %-9.2f %-10d %-9.2f\n",
				i+1, purchaseNames[i], purchasePrices[i], purchaseQuantities[i], purchaseTotals[i])
			finalTotal += purchaseTotals[i]
		}
		fmt.Printf("\nFinal Total: %.2f\n", finalTotal)
	} else {
		fmt.Println("No items were purchased.")
		fmt.Println("\n----Thank you for using our Inventory Management System-----")
	}
}
