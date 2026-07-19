package payment

type Transaction struct {
	FromUserID string
	ToUserID   string
	Amount     float64
}