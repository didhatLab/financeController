package webmodels

type TestSpending struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}
