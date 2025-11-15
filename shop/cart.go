package shop

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Item хранит сведения о строке заказа.
type Item struct {
	Name     string
	Price    float64
	Quantity float64
}

// Cart собирает товары и считает итог.
type Cart struct {
	Items []Item
}

// NewCart создаёт пустую корзину.
func NewCart() *Cart {
	return &Cart{}
}

// Load читает строки из reader и добавляет товары в корзину.
func (c *Cart) Load(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	lineNo := 0

	for scanner.Scan() {
		lineNo++
		if err := c.AddLine(scanner.Text(), lineNo); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// AddLine обрабатывает одну строку и добавляет товар, если строка валидна.
func (c *Cart) AddLine(raw string, lineNo int) error {
	text := strings.TrimSpace(raw)
	if text == "" || strings.HasPrefix(text, "#") {
		return nil
	}

	parts := strings.Fields(text)
	if len(parts) < 2 {
		return fmt.Errorf("строка %d: ожидаются цена и количество", lineNo)
	}

	priceField := parts[len(parts)-2]
	amountField := parts[len(parts)-1]

	price, err := strconv.ParseFloat(priceField, 64)
	if err != nil {
		return fmt.Errorf("строка %d: неверная цена %q: %w", lineNo, priceField, err)
	}

	quantity, err := strconv.ParseFloat(amountField, 64)
	if err != nil {
		return fmt.Errorf("строка %d: неверное количество %q: %w", lineNo, amountField, err)
	}

	name := strings.Join(parts[:len(parts)-2], " ")
	c.Items = append(c.Items, Item{Name: name, Price: price, Quantity: quantity})
	return nil
}

// Total возвращает сумму всех товаров.
func (c Cart) Total() float64 {
	var sum float64
	for _, item := range c.Items {
		sum += item.Price * item.Quantity
	}
	return sum
}
