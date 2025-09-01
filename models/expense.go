package models


type Expense struct {
	ID      	 	int `json:"id"`
	Title   		string `json:"title"`
	Category 		string `json:"category"` 
	Amount  		int `json:"amount"`
	Date     		string	`json:"date"`
}