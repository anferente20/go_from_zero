package gobases3

const (
	Small  = "S"
	Medium = "M"
	Large  = "L"
)

type product interface {
	GetPrice() float32
}

type SmallProduct struct {
	price float32
}

func (s SmallProduct) GetPrice() float32 {
	return s.price
}

type MediumProduct struct {
	price float32
}

func (m MediumProduct) GetPrice() float32 {
	return m.price * 1.03
}

type LargeProduct struct {
	price float32
}

func (l LargeProduct) GetPrice() float32 {
	return (l.price * 1.06) + float32(2500)
}

func NewProduct(productType string, price float32) product {
	switch productType {
	case Medium:
		return MediumProduct{price: price}
	case Large:
		return LargeProduct{price: price}
	default:
		return SmallProduct{price: price}
	}
}
