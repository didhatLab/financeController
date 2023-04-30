package webmodels

type TestSpending struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
	GroupId     *int   `json:"group_id"`
}

type DeleteRequest struct {
	SpendId int  `json:"spend_id"`
	GroupId *int `json:"group_id"`
}

type UpdateRequest struct {
	SpendId     int     `json:"spend_id"`
	Name        *string `json:"name"`
	Type        *string `json:"type"`
	Amount      *int    `json:"amount"`
	Description *string `json:"description"`
	GroupId     *int    `json:"group_id"`
}

type GroupSpendRequest struct {
	GroupId int `json:"group_id"`
}
