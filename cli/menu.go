package cli

import (
	"bufio"
	"expense-tracker-cli/models"
	"expense-tracker-cli/services"
	"expense-tracker-cli/utils"
	"fmt"
	"os"
	"strings"
)

func Run() {
	services.LoadExpenses()
	showMenu()
}

func showMenu() {
	stay := 1
	for stay == 1 {
		fmt.Println("===== Expense Tracker =====")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View All Expenses")
		fmt.Println("3. View Expenses (Filter by Category)")
		fmt.Println("4. Delete Expense")
		fmt.Println("5. Update Expense")
		fmt.Println("6. Expenses Summary")
		fmt.Println("7. Generate Report")
		fmt.Println("8. Exit")
		fmt.Print(">> ")
		var input int
		fmt.Scanf("%d\n", &input)

		if input == 1 {
			inputExpense()
		} else if input == 2 {
			viewAllExpenses()
		} else if input == 3 {
			viewExpensesByCategory()
		} else if input == 4 {
			deleteExpense()
		} else if input ==5 {
			updateExpense()
		} else if input==6{
			expensesSummary()
		} else if input==7{
			generateReport()
		} else if input==8{
			fmt.Println("Thank you !!!")
			stay = 0
		}

	}

}


func inputExpense() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input expense description : ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Input expense category [Food | Transport | Entertainment | Other] : ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)
	for category != "Food" && category != "Transport" && category != "Entertainment" && category != "Other" {
		fmt.Print("Input expense category [Food | Transport | Entertainment | Other] : ")
		category, _ = reader.ReadString('\n')
		category = strings.TrimSpace(category)
	}

	fmt.Print("Input expense amount : ")
	var amount int
	fmt.Scanf("%d\n", &amount)

	fmt.Print("Input expense date (YYYY-MM-DD) : ")
	date, _ := reader.ReadString('\n')
	date = strings.TrimSpace(date)
	for !utils.IsValidDate(date){
		fmt.Print("Input expense date (YYYY-MM-DD) : ")
		date, _ = reader.ReadString('\n')
		date = strings.TrimSpace(date)
	}

	expense := models.Expense{
		Title:    title,
		Category: category,
		Amount:   amount,
		Date:     date,
	}

	services.AddExpenses(expense)
}

func viewAllExpenses() {
	if len(services.ListExpenses()) == 0 {
		fmt.Println("There isn't any Expenses on the List")
		return
	}
	fmt.Println("======================================================================================================")
	fmt.Printf("| %-2s | %-35s | %-20s | %-20s| %-10s |\n", "ID", "Description", "Category", "Amount", "Date")
	fmt.Println("======================================================================================================")
	for i := 0; i < len(services.ListExpenses()); i++ {
		x := services.ListExpenses()[i]
		fmt.Printf("| %-2d | %-35s | %-20s | %-20d| %-10s |\n", x.ID, x.Title, x.Category, x.Amount, x.Date)
	}
	fmt.Println("======================================================================================================")
}

func deleteExpense() {
	viewAllExpenses()
	if len(services.ListExpenses())==0{
		return
	}
	var id int
	stay:=0
	for stay==0 {
		fmt.Print("Input expense ID : ")
		fmt.Scanf("%d\n", &id)
		for i:=0; i<len(services.ListExpenses()); i++{
			x:=services.ListExpenses()[i]
			if(id==x.ID){
				stay=1
				break
			}

		}
		if stay==0{
			fmt.Println("There isn't any ID that matched with existing ID in the Expense List, Please input valid ID !!!")
		}

	}
	services.DeleteExpenses(id)


}
func updateExpense(){
	viewAllExpenses()
	if len(services.ListExpenses())==0{
		return
	}
	var id int
	var index int
	stay:=0
	for stay==0 {
		fmt.Print("Input expense ID : ")
		fmt.Scanf("%d\n", &id)
		for i:=0; i<len(services.ListExpenses()); i++{
			x:=services.ListExpenses()[i]
			if(id==x.ID){
				index=i
				stay=1
				break
			}

		}
		if stay==0{
			fmt.Println("There isn't any ID that matched with existing ID in the Expense List, Please input valid ID !!!")
		}

	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input expense description : ")
	newTitle, _ := reader.ReadString('\n')
	newTitle = strings.TrimSpace(newTitle)

	fmt.Print("Input expense category [Food | Transport | Entertainment | Other] : ")
	newCategory, _ := reader.ReadString('\n')
	newCategory = strings.TrimSpace(newCategory)
	for newCategory != "Food" && newCategory != "Transport" && newCategory != "Entertainment" && newCategory != "Other" {
		fmt.Print("Input expense category [Food | Transport | Entertainment | Other] : ")
		newCategory, _ = reader.ReadString('\n')
		newCategory = strings.TrimSpace(newCategory)
	}

	fmt.Print("Input expense amount : ")
	var newAmount int
	fmt.Scanf("%d\n", &newAmount)

	fmt.Print("Input expense date (YYYY-MM-DD) : ")
	newDate, _ := reader.ReadString('\n')
	newDate = strings.TrimSpace(newDate)
	for !utils.IsValidDate(newDate){
		fmt.Print("Input expense date (YYYY-MM-DD) : ")
		newDate, _ = reader.ReadString('\n')
		newDate = strings.TrimSpace(newDate)
	}

	services.UpdateExpenses(id, index, newTitle,newCategory,newAmount,newDate)

}
func viewExpensesByCategory(){
	stay:=1
	for stay==1{
		if len(services.Expenses)==0{
			fmt.Println("There isn't any Expenses on the List")
			break
		}
		fmt.Println("Choose Expenses Category : ")
		fmt.Println("1. Food")
		fmt.Println("2. Transport")
		fmt.Println("3. Entertainment")
		fmt.Println("4. Other")
		fmt.Println("5. Back")
		fmt.Print(">> ")
		var input int
		fmt.Scanf("%d\n", &input)
		if input==1{
			services.ViewByCategory("Food")
		}else if input==2{
			services.ViewByCategory("Transport")
		}else if input==3{
			services.ViewByCategory("Entertainment")
		}else if input==4{
			services.ViewByCategory("Other")
		}else if input==5{
			stay=0
		}
	}
	
}

func expensesSummary(){
	services.ExpensesSummary()
}
func generateReport() {
	stay:=1
	for stay==1{
		fmt.Println("Choose Expenses Report")
		fmt.Println("1. All Expenses Report")
		fmt.Println("2. Report per Category")
		fmt.Println("3. Back")
		fmt.Print(">> ")
		var input int
		fmt.Scanf("%d\n", &input)
		if input ==1 {
			services.GenerateAllReport()
		}else if input==2{
			
			stay2:=1
			for stay2==1{
				if len(services.Expenses)==0{
					fmt.Println("There isn't any Expenses on the List")
					break
				}
				fmt.Println("Choose Expenses Category : ")
				fmt.Println("1. Food")
				fmt.Println("2. Transport")
				fmt.Println("3. Entertainment")
				fmt.Println("4. Other")
				fmt.Println("5. Back")
				fmt.Print(">> ")
				var input2 int
				fmt.Scanf("%d\n", &input2)
				if input2==1{
					services.GenerateCategoryReport("Food")
				}else if input2==2{
					services.GenerateCategoryReport("Transport")
				}else if input2==3{
					services.GenerateCategoryReport("Entertainment")
				}else if input2==4{
					services.GenerateCategoryReport("Other")
				}else if input2==5{
					stay2=0
				}
			}

		}else if input==3{
			stay=0
		}
	}

}
