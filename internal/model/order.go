package model

type Order struct {
	ID int `json:"id"`
	CustomerName string `json:"customer_name"`
	Amount float64 `json:"amount"`
}

//json tags ensure correct Api response format