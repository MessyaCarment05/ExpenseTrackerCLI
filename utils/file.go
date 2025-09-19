package utils

import (
	"encoding/json"
	"os"
)

// Save struct expenses ke JSON
func SaveJSON(filename string, dataList interface{})error{
	file,err:=json.MarshalIndent(dataList, "", "  ")
	if err!=nil{
		return err
	}
	return os.WriteFile(filename, file, 0644)
}

// Membaca file json ke dalam variabel target->Expenses
func LoadJSON(filename string, target interface{}) error{
	dataList,err:=os.ReadFile(filename)
	if err!=nil{
		return err
	}

	// kalau file expenses.json masih kosong langsung return nil
	if len(dataList)==0{
		return nil
	}
	return json.Unmarshal(dataList, target)
}

