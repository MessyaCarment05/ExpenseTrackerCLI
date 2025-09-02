package services

import (
	"expense-tracker-cli/models"
	"expense-tracker-cli/utils"
	"fmt"
	"path/filepath"
)

//save temp expenses
var Expenses[]models.Expense
var LastID int
var dataFile=filepath.Join("data", "expenses.json")

func AddExpenses(expense models.Expense){
	expense.ID=LastID
	LastID++
	Expenses=append(Expenses, expense)
	fmt.Println("Successfully add new expense !!!")
	SaveExpenses()
	
}

func ListExpenses()[]models.Expense{
	// ini buat return slice Expenses
	return Expenses 

}

func DeleteExpenses(id int){
	var index int
	for i:=0; i<len(Expenses); i++{
			x:=Expenses[i]
			if(id==x.ID){
				index=i
				break
			}
	}
	
	deletedID:=Expenses[index].ID
	deletedTitle:=Expenses[index].Title
	// menggabungkan slice yang sudah mengeluarkan elemen dengan index yang mau dihapus
	Expenses=append(Expenses[:index], Expenses[index+1:]...)
	SaveExpenses()
	fmt.Printf("Successfully delete expense with ID : %d and Description : %s !!!\n", deletedID, deletedTitle)
}

func LoadExpenses(){
	err:=utils.LoadJSON(dataFile, &Expenses)
	if err!=nil{
		fmt.Println("Load Data Failed: ", err)
		Expenses=[]models.Expense{}
		LastID=1
		return
	}
	LastID=1
	for _, expense:=range Expenses{
		if expense.ID>=LastID{
			LastID=expense.ID+1

		}
	}
	

}

func SaveExpenses(){
	err:=utils.SaveJSON(dataFile, Expenses)
	if err!=nil{
		fmt.Println("Failed save expense : ", err)
	}
}