package main

import (
	"fmt"
	"os"

	"github.com/OscarSalnikov/goshop/shop"
)

func main() {
	cart := shop.NewCart()

	if len(os.Args) == 1 {
		fmt.Println("Введите строки заказа (Ctrl+D, на Windows Ctrl+Z+Enter, завершит ввод):")
		if err := cart.Load(os.Stdin); err != nil {
			fail(err)
		}
	} else {
		for _, path := range os.Args[1:] {
			if err := loadFile(cart, path); err != nil {
				fail(err)
			}
		}
	}

	fmt.Printf("Итоговая сумма: %.2f\n", cart.Total())
}

func loadFile(cart *shop.Cart, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return cart.Load(file)
}

func fail(err error) {
	fmt.Fprintf(os.Stderr, "ошибка: %v\n", err)
	os.Exit(1)
}
