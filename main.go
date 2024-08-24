package main

import (
	"errors"
	"fmt"
	"go-simple-orm-crud-postgres/models"
	"gorm.io/gorm"
)

func CreatePayment(db *gorm.DB, payment models.Payment) (int64, error) {
	result := db.Create(&payment)
	if result.RowsAffected == 0 {
		return 0, errors.New("payment not created")
	}
	return result.RowsAffected, nil
}

func SelectPaymentWIthId(db *gorm.DB, id string) (models.Payment, error) {
	var payment models.Payment
	result := db.First(&payment, "id = ?", id)
	if result.RowsAffected == 0 {
		return models.Payment{}, errors.New("payment data not found")
	}
	return payment, nil
}
func main() {
	//TIP Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined or highlighted text
	// to see how GoLand suggests fixing it.
	s := "gopher"
	fmt.Println("Hello and welcome, %s!", s)

	for i := 1; i <= 5; i++ {
		//TIP You can try debugging your code. We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>. To start your debugging session,
		// right-click your code in the editor and select the <b>Debug</b> option.
		fmt.Println("i =", 100/i)
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
