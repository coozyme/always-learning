package valueobject

type LineOfProduct struct {
	Quantity int
	Product  Product
}

func NewLineProduct(product Product, quantity int) *LineOfProduct {
	return &LineOfProduct{
		Product:  product,
		Quantity: quantity,
	}
}
