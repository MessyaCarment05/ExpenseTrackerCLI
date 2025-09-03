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
	// untuk print data yang dihapus
	deletedID:=Expenses[index].ID
	deletedTitle:=Expenses[index].Title
	// menggabungkan slice yang sudah mengeluarkan elemen dengan index yang mau dihapus
	Expenses=append(Expenses[:index], Expenses[index+1:]...)
	SaveExpenses()
	fmt.Printf("Successfully delete expense with ID : %d and Description : %s !!!\n", deletedID, deletedTitle)
}

func UpdateExpenses(id int, index int, newTitle string, newCategory string, newAmount int, newDate string){

	updatedID:=Expenses[index].ID
	updatedTitle:=Expenses[index].Title

	Expenses[index].Title=newTitle
	Expenses[index].Category=newCategory
	Expenses[index].Amount=newAmount
	Expenses[index].Date=newDate

	SaveExpenses()
	fmt.Printf("Successfully update expense with ID : %d and Description : %s !!!\n", updatedID, updatedTitle)
}

func ViewByCategory(category string){
	if category=="Food"{
		countFood:=0
		for i:=0 ; i<len(Expenses); i++{
			x:=Expenses[i]
			if(x.Category)=="Food"{
				countFood++
			}
		}
		if countFood==0{
			fmt.Println("No Expenses with Food Category")
			return
		}
		fmt.Println("======================================================================================================")
		fmt.Printf("| %-2s | %-35s | %-20s | %-20s| %-10s |\n", "ID", "Description", "Category", "Amount", "Date")
		fmt.Println("======================================================================================================")
		for i := 0; i < len(Expenses); i++ {
			a := Expenses[i]
			if a.Category=="Food"{
				fmt.Printf("| %-2d | %-35s | %-20s | %-20d| %-10s |\n", a.ID, a.Title, a.Category, a.Amount, a.Date)
			}
			
		}
		fmt.Println("======================================================================================================")
	} else if category=="Transport"{
		countTransport:=0
		for i:=0 ; i<len(Expenses); i++{
			x:=Expenses[i]
			if(x.Category)=="Transport"{
				countTransport++
			}
		}
		if countTransport==0{
			fmt.Println("No Expenses with Transport Category")
			return
		}
		fmt.Println("======================================================================================================")
		fmt.Printf("| %-2s | %-35s | %-20s | %-20s| %-10s |\n", "ID", "Description", "Category", "Amount", "Date")
		fmt.Println("======================================================================================================")
		for i := 0; i < len(Expenses); i++ {
			a := Expenses[i]
			if a.Category=="Transport"{
				fmt.Printf("| %-2d | %-35s | %-20s | %-20d| %-10s |\n", a.ID, a.Title, a.Category, a.Amount, a.Date)
			}
			
		}
		fmt.Println("======================================================================================================")
	} else if category=="Entertainment"{
		countEntertainment:=0
		for i:=0 ; i<len(Expenses); i++{
			x:=Expenses[i]
			if(x.Category)=="Entertainment"{
				countEntertainment++
			}
		}
		if countEntertainment==0{
			fmt.Println("No Expenses with Entertainment Category")
			return
		}
		fmt.Println("======================================================================================================")
		fmt.Printf("| %-2s | %-35s | %-20s | %-20s| %-10s |\n", "ID", "Description", "Category", "Amount", "Date")
		fmt.Println("======================================================================================================")
		for i := 0; i < len(Expenses); i++ {
			a := Expenses[i]
			if a.Category=="Entertainment"{
				fmt.Printf("| %-2d | %-35s | %-20s | %-20d| %-10s |\n", a.ID, a.Title, a.Category, a.Amount, a.Date)
			}
			
		}
		fmt.Println("======================================================================================================")
	} else if category=="Other"{
		countOther:=0
		for i:=0 ; i<len(Expenses); i++{
			x:=Expenses[i]
			if(x.Category)=="Other"{
				countOther++
			}
		}
		if countOther==0{
			fmt.Println("No Expenses with Other Category")
			return
		}
		fmt.Println("======================================================================================================")
		fmt.Printf("| %-2s | %-35s | %-20s | %-20s| %-10s |\n", "ID", "Description", "Category", "Amount", "Date")
		fmt.Println("======================================================================================================")
		for i := 0; i < len(Expenses); i++ {
			a := Expenses[i]
			if a.Category=="Other"{
				fmt.Printf("| %-2d | %-35s | %-20s | %-20d| %-10s |\n", a.ID, a.Title, a.Category, a.Amount, a.Date)
			}
			
		}
		fmt.Println("======================================================================================================")
	}
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