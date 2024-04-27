# Stock Share Transactions
## Table of Contents
- [Overview](#overview)
- [Penjelasan solution.go](#penjelasan-solutiongo)

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
