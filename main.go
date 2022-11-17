package main

import category "github.com/gabrielsc1998/inventory-control-golang/internal/domain/category/entity"

func main() {
	// n, _ := notifications.NewNotification()

	// n.AddError("ctx", "msg")
	// n.AddError("ctx", "1234")

	// print(n.GetErrors("ctx"))

	cat, _ := category.NewCategory("", "", "")

	// print(cat.IsActive())
	n := cat.Notifications()
	print(n.GetErrors())
	// print(n.GetAllErrors())
}
