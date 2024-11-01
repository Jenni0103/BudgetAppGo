package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Budget struct {
	totalBudget   float64
	expenses      map[string]float64
	totalExpenses float64
}

// NewBudget membuat instance Budget baru
func NewBudget(total float64) *Budget {
	return &Budget{
		totalBudget:   total,
		expenses:      make(map[string]float64),
		totalExpenses: 0,
	}
}

// AddExpense menambah pengeluaran ke dalam budget
func (b *Budget) AddExpense(name string, amount float64) {
	b.expenses[name] = amount
	b.totalExpenses += amount
}

// RemainingBudget menghitung sisa budget
func (b *Budget) RemainingBudget() float64 {
	return b.totalBudget - b.totalExpenses
}

// DisplayExpenses menampilkan daftar pengeluaran dan sisa budget
func (b *Budget) DisplayExpenses() {
	fmt.Println("Pengeluaran:")
	for name, amount := range b.expenses {
		fmt.Printf("- %s: Rp%.2f\n", name, amount)
	}
	fmt.Printf("Total Pengeluaran: Rp%.2f\n", b.totalExpenses)
	fmt.Printf("Sisa Budget: Rp%.2f\n", b.RemainingBudget())
}

// parseAmount mengambil input string dengan format koma lalu mengubah ke float
func parseAmount(input string) (float64, error) {
	// Menghapus tanda koma dan spasi tambahan
	cleanedInput := strings.ReplaceAll(input, ",", "")
	return strconv.ParseFloat(cleanedInput, 64)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Meminta user memasukkan total budget
	fmt.Println("Selamat datang di aplikasi pengelolaan budget!")
	fmt.Println("Silakan masukkan total budget dengan format angka yang benar (contoh: 1,000,000):")
	fmt.Print("Masukkan total budget: Rp")
	totalBudgetInput, _ := reader.ReadString('\n')
	totalBudgetInput = strings.TrimSpace(totalBudgetInput)
	totalBudget, err := parseAmount(totalBudgetInput)
	if err != nil {
		fmt.Println("Input tidak valid. Pastikan formatnya benar (misalnya, 1,000,000).")
		return
	}

	budget := NewBudget(totalBudget)

	for {
		fmt.Println("\nPilih opsi:")
		fmt.Println("1. Tambah Pengeluaran")
		fmt.Println("2. Lihat Pengeluaran dan Sisa Budget")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih opsi (1/2/3): ")

		optionInput, _ := reader.ReadString('\n')
		optionInput = strings.TrimSpace(optionInput)
		option, err := strconv.Atoi(optionInput)
		if err != nil {
			fmt.Println("Pilihan tidak valid, masukkan angka 1, 2, atau 3.")
			continue
		}

		switch option {
		case 1:
			// Tambah Pengeluaran
			fmt.Print("Nama Pengeluaran: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Println("Jumlah Pengeluaran (contoh: 50,000):")
			fmt.Print("Rp")
			amountInput, _ := reader.ReadString('\n')
			amountInput = strings.TrimSpace(amountInput)
			amount, err := parseAmount(amountInput)
			if err != nil {
				fmt.Println("Input tidak valid. Pastikan formatnya benar (misalnya, 50,000).")
				continue
			}

			budget.AddExpense(name, amount)
			fmt.Println("Pengeluaran berhasil ditambahkan!")

		case 2:
			// Lihat Pengeluaran dan Sisa Budget
			budget.DisplayExpenses()

		case 3:
			fmt.Println("Terima kasih telah menggunakan aplikasi budget.")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
