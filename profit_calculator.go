package main

import (
	"fmt"
	"os"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	var input int

	for {
		fmt.Println("Welcome to Profit Calculator")
		fmt.Println("1. Earning Before Taxes Calculator")
		fmt.Println("2. Profit Calculator")
		fmt.Println("3. Ratio Calculator")
		fmt.Println("4. Exit")
		fmt.Println("Choose your menu")
		fmt.Scan(&input)

		switch input {
		case 1:
			fmt.Print("Enter the revenue: ")
			fmt.Scan(&revenue)
			fmt.Print("Enter the expenses: ")
			fmt.Scan(&expenses)
			EBT := revenue - expenses
			fmt.Printf("Earning Before Taxes: %.2f\n", EBT)
		case 2:
			fmt.Print("Enter the revenue: ")
			fmt.Scan(&revenue)
			fmt.Print("Enter the expenses: ")
			fmt.Scan(&expenses)
			fmt.Print("Enter the tax rate: ")
			fmt.Scan(&taxRate)
			profit := (revenue - expenses) * (1 - (taxRate / 100))
			fmt.Printf("Profit: %.2f\n", profit)
		case 3:
			fmt.Print("Enter the revenue: ")
			fmt.Scan(&revenue)
			fmt.Print("Enter the expenses: ")
			fmt.Scan(&expenses)
			fmt.Print("Enter the tax rate: ")
			fmt.Scan(&taxRate)
			profit := (revenue - expenses) * (1 - (taxRate / 100))
			EBT := revenue - expenses
			ratio := EBT / profit
			fmt.Printf("Ratio: %.2f\n", ratio)
		case 4:
			fmt.Println("Exiting")
			os.Exit(0)
		default:
			fmt.Println("Invalid input. Exiting.")
			os.Exit(1)
		}
	}

}
