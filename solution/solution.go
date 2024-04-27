package main

import "fmt"

func maxProfit(prices []int, transactions int) int {
	if len(prices) <= 1 || transactions == 0 {
		return 0
	}

	profit := 0

	// Jika jumlah transaksi melebihi separuh jumlah hari, kita dapat melakukan transaksi
	// setiap kali harga saham naik.
	if transactions >= len(prices)/2 {
		for i := 1; i < len(prices); i++ {
			if prices[i] > prices[i-1] {
				profit += prices[i] - prices[i-1]
			}
		}
		return profit
	}

	// Jika tidak, kita akan menggunakan pendekatan berbasis transaksi terbatas.
	// Gunakan pendekatan Greedy untuk menemukan titik pembelian dan penjualan.
	buy := make([]int, transactions+1)
	sell := make([]int, transactions+1)

	for i := 0; i <= transactions; i++ {
		buy[i] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := 1; j <= transactions; j++ {
			buy[j] = max(buy[j], sell[j-1]-prices[i])
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}

	return sell[transactions]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{7, 1, 4, 9, 10, 2, 0, 1, 3, 4, 7, 1, 2, 3, 4, 5}

	i := 10000000000000000
	fmt.Println("Maximum profit:", maxProfit(prices, i))
}
