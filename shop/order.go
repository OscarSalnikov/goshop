package shop

import "io"

// OrderCost вычисляет стоимость заказа, используя Cart.
func OrderCost(reader io.Reader) (float64, error) {
	cart := NewCart()
	if err := cart.Load(reader); err != nil {
		return 0, err
	}
	return cart.Total(), nil
}
