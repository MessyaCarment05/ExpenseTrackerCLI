package services

import (
	"expense-tracker-cli/models"
	
)

//save temp expenses
var Expenses[]models.Expense
var id int =1

func AddExpenses(expense models.Expense){
	expense.ID=id
	id++
	Expenses=append(Expenses, expense)
	
}

func ListExpenses()[]models.Expense{
	// ini buat return slice Expenses
	return Expenses 

}