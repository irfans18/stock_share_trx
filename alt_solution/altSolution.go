package main

import (
	"fmt"
)

// Fungsi maxProfit mengambil slice harga dan jumlah transaksi yang dapat dilakukan.
// Ini mengembalikan maksimum profit yang dapat diperoleh dari transaksi tersebut.
func maxProfit(prices []int, transactions int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	
	// Jika jumlah transaksi yang diizinkan lebih besar atau sama dengan setengah jumlah harga,
	// kita bisa memanfaatkan setiap kenaikan harga, sehingga profit adalah total dari setiap
	// kenaikan harga.
	if transactions >= n/2 {
		profit := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				profit += prices[i] - prices[i-1]
			}
		}
		return profit
	}

	// Jika jumlah transaksi lebih kecil daripada setengah jumlah harga, kita menggunakan pendekatan
	// dinamis untuk menghitung profit maksimum yang dapat diperoleh.
	dp := make([][]int, transactions+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 1; i <= transactions; i++ {
		maxDiff := -prices[0]
		for j := 1; j < n; j++ {
			// dp[i][j] menyimpan profit maksimum pada transaksi i dan pada hari j.
			// Ini dihitung dengan membandingkan profit pada hari j-1 dengan harga pada hari j ditambah
			// dengan nilai maksimum perbedaan (maxDiff) antara profit dari transaksi sebelumnya dan harga pada hari j.
			dp[i][j] = max(dp[i][j-1], prices[j]+maxDiff)
			// maxDiff diupdate dengan membandingkan nilai sebelumnya dengan perbedaan
			// antara profit dari transaksi sebelumnya pada hari j dan harga pada hari j.
			maxDiff = max(maxDiff, dp[i-1][j]-prices[j])
		}
	}

	// Profit maksimum terletak pada dp[transactions][n-1], yaitu profit maksimum dengan jumlah
	// transaksi yang diizinkan dan pada hari terakhir.
	return dp[transactions][n-1]
}

// Fungsi max mengambil dua integer dan mengembalikan nilai yang lebih besar di antara keduanya.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{7, 1, 4, 9, 10, 2, 0, 1, 3, 4, 7, 1, 2, 3, 4, 5}

	transactions := 3
	fmt.Println("Maximum profit:", maxProfit(prices, transactions)) 
}
