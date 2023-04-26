package webmodels

type TestSpending struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
}

type DeleteRequest struct {
	SpendId int `json:"spend_id"`
}

type UpdateRequest struct {
	SpendId int     `json:"spend_id"`
	Name    *string `json:"name"`
	Type    *string `json:"type"`
	Amount  *int    `json:"amount"`
}
