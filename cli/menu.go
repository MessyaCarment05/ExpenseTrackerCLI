package cli

import (
	"bufio"
	"expense-tracker-cli/models"
	"expense-tracker-cli/services"
	"fmt"
	"os"
	"strings"
)

func Run() {
	showMenu()
}

func showMenu() {
	stay:=1
	for stay==1{
		fmt.Println("===== Expense Tracker =====")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View All Expenses")
		fmt.Println("3. Delete Expense")
		fmt.Println("4. Generate Report")
		fmt.Println("5. Exit")
		fmt.Print(">> ")
		var input int
		fmt.Scanf("%d\n", &input)
		
		if input==1 {
			inputExpense()
		}else if input==2 {
			viewAllExpenses()
		}else if input==3 {
			deleteExpense()
		}else if input==4 {
			generateReport()
		}else{
			fmt.Println("Thank you !!!")
			stay=0
		}

	}

}

func inputExpense(){
	reader:=bufio.NewReader(os.Stdin)
	fmt.Print("Input expense description : ")
	title, _:=reader.ReadString('\n')
	title=strings.TrimSpace(title)

	fmt.Print("Input expense category [Food | Transport | Entertainment | Other] : ")
	category, _:=reader.ReadString('\n')
	category=strings.TrimSpace(category)
	for category!="Food"&&category!="Transport"&&category!="Entertainment"&&category!="Other"{
		fmt.Print("Input expense category [Food | Transport | Entertainment | Other] : ")
		category, _=reader.ReadString('\n')
		category=strings.TrimSpace(category)
	}

	fmt.Print("Input expense amount : ")
	var amount int
	fmt.Scanf("%d\n", &amount)

	fmt.Print("Input expense date (YYYY-MM-DD) : ")
	date,_:=reader.ReadString('\n')
	date=strings.TrimSpace(date)
	
	expense:=models.Expense{
		Title : title,
		Category : category,
		Amount : amount,
		Date:date,

	}

	services.AddExpenses(expense)
	// fmt.Println("Jumlah expense:", len(services.ListExpenses()))

	






}


func viewAllExpenses(){
	if len(services.ListExpenses())==0{
		fmt.Println("No Expenses List")
		return
	}
	fmt.Println("============================================================")
	for i:=0; i<len(services.ListExpenses()); i++{
		x:=services.ListExpenses()[i]
		fmt.Printf("| %d | %-35s | %-10s | %d| %-5s |\n",x.ID, x.Title, x.Category, x.Amount, x.Date )
	}
	fmt.Println("============================================================")
}

func deleteExpense(){
	fmt.Println("test3")

}
func  generateReport(){
	fmt.Println("test4")

}