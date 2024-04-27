## Table of Contents
- [Overview](#overview)
- [Penjelasan solution.go](#penjelasan-solutiongo)
- [Penjelasan altSolution.go](#penjelasan-altsolutiongo)


## Overview
**Problem Statement:** <br> 
Sebagai seorang IT Consultant, Anda ditugaskan untuk menyelesaikan masalah dalam melakukan pembelian dan penjualan stok dalam beberapa hari untuk mendapatkan keuntungan maksimum. Anda diberikan sebuah array angka integer positif yang mewakili harga stok pada berbagai hari. Selain itu, Anda juga diberikan sebuah integer \( i \), yang merupakan jumlah transaksi yang diizinkan. Setiap transaksi terdiri dari pembelian stok pada suatu hari dan penjualannya di kemudian hari. Tugas Anda adalah menulis sebuah fungsi yang menghasilkan keuntungan maksimum yang dapat diperoleh dengan menggunakan \( i \) transaksi tersebut.


## Penjelasan solution.go
Pada fungsi `maxProfit`:

```go
func maxProfit(prices []int, transactions int) int {
```
Ini adalah deklarasi fungsi `maxProfit` yang mengambil dua parameter: `prices`, yang merupakan slice dari harga saham, dan `transactions`, yang merupakan jumlah transaksi maksimum yang dapat dilakukan.

```go
	if len(prices) <= 1 || transactions == 0 {
		return 0
	}
```
Kondisi pertama adalah untuk mengecek apakah panjang `prices` kurang dari atau sama dengan 1, atau jika `transactions` sama dengan 0. Jika salah satu kondisi tersebut terpenuhi, fungsi akan mengembalikan nilai 0, karena tidak mungkin mendapatkan profit dalam kasus ini.

```go
	profit := 0
```
Variabel `profit` diinisialisasi dengan nilai 0. Ini akan digunakan untuk menyimpan profit maksimum yang akan dihitung dalam fungsi.

```go
	if transactions >= len(prices)/2 {
		for i := 1; i < len(prices); i++ {
			if prices[i] > prices[i-1] {
				profit += prices[i] - prices[i-1]
			}
		}
		return profit
	}
```
Ini adalah kasus khusus ketika jumlah transaksi yang diminta oleh pengguna melebihi setengah panjang `prices`. Dalam hal ini, kita dapat memanfaatkan setiap kenaikan harga untuk mendapatkan profit. Oleh karena itu, kita cukup melakukan iterasi melalui harga dan menambahkan selisih harga jika harga saat ini lebih besar dari harga sebelumnya. Akhirnya, fungsi mengembalikan nilai `profit` yang telah dihitung.

```go
	buy := make([]int, transactions+1)
	sell := make([]int, transactions+1)
```
Kita membuat dua slice, `buy` dan `sell`, dengan panjang `transactions + 1`. `buy[i]` akan menyimpan profit maksimum setelah melakukan `i` pembelian, sedangkan `sell[i]` akan menyimpan profit maksimum setelah melakukan `i` penjualan.

```go
	for i := 0; i <= transactions; i++ {
		buy[i] = -prices[0]
	}
```
Pada awalnya, semua nilai dalam slice `buy` diinisialisasi dengan nilai negatif dari harga saham pertama. Ini menandakan bahwa kita tidak melakukan pembelian apa pun, tetapi hanya mempertimbangkan kemungkinan pembelian di masa depan.

```go
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= transactions; j++ {
			buy[j] = max(buy[j], sell[j-1]-prices[i])
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}
```
Di sini, kita melakukan iterasi melalui semua harga saham. Untuk setiap harga, kita iterasi melalui semua jumlah transaksi yang mungkin. Dalam setiap iterasi, kita perbarui nilai `buy[j]` dan `sell[j]` menggunakan rumus yang diberikan dalam algoritma Dynamic Programming untuk menghitung profit maksimum setelah melakukan `j` transaksi pada titik tertentu.

```go
	return sell[transactions]
```
Akhirnya, kita mengembalikan nilai `sell[transactions]`, yang merupakan profit maksimum setelah melakukan `transactions` transaksi. Ini adalah hasil akhir dari algoritma yang diimplementasikan.

## Penjelasan altSolution.go
Pada fungsi `maxProfit`:

```go
func maxProfit(prices []int, transactions int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	if transactions >= n/2 {
		profit := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				profit += prices[i] - prices[i-1]
			}
		}
		return profit
	}
```
Fungsi `maxProfit` ini mengambil sebuah slice `int` yang mewakili harga-harga dan sebuah integer `transactions` yang mewakili jumlah transaksi yang diizinkan. Fungsi ini menghitung profit maksimum yang dapat diperoleh dengan transaksi dan harga-harga yang diberikan.

Jika hanya ada satu atau nol harga, maka akan mengembalikan 0, karena tidak ada profit yang bisa didapatkan.

Jika jumlah transaksi yang diizinkan lebih dari atau sama dengan setengah dari jumlah harga, maka kita bisa melakukan transaksi sebanyak yang diinginkan. Jadi, fungsi tersebut menghitung profit secara langsung dengan mengiterasi harga-harga dan menambahkan selisih antara harga-harga yang berturut-turut di mana harga tersebut meningkat.

```go
	dp := make([][]int, transactions+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}
```
Di sini, kita menginisialisasi slice 2D `dp` untuk menyimpan nilai-nilai pemrograman dinamis. Ini memiliki `transactions+1` baris dan `n` kolom.

```go
	for i := 1; i <= transactions; i++ {
		maxDiff := -prices[0]
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i][j-1], prices[j]+maxDiff)
			maxDiff = max(maxDiff, dp[i-1][j]-prices[j])
		}
	}
```
Ini adalah bagian pemrograman dinamis dari fungsi tersebut. Ini mengiterasi transaksi dan harga-harga, menghitung profit maksimum yang dapat diperoleh pada setiap transaksi dan setiap harga.

```go
	return dp[transactions][n-1]
```
Akhirnya, fungsi tersebut mengembalikan profit maksimum yang dapat diperoleh setelah `transactions` transaksi.

```go
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
Ini adalah fungsi pembantu `max` yang mengembalikan nilai maksimum dari dua bilangan bulat `a` dan `b`.

```go
func main() {
	prices := []int{7, 1, 4, 9, 10, 2, 0, 1, 3, 4, 7, 1, 2, 3, 4, 5}

	transactions := 3
	fmt.Println("Maximum profit:", maxProfit(prices, transactions)) 
}
```
Di fungsi `main`, sebuah slice `prices` yang berisi beberapa harga contoh dibuat, dan jumlah transaksi yang diizinkan diatur menjadi 3. Kemudian, dipanggil fungsi `maxProfit` dengan nilai-nilai tersebut dan mencetak profit maksimum yang diperoleh.