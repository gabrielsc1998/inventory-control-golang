package entity

import (
	"github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/domain/notifications"
	vo_id "github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/domain/value-objects/id"
	vo_name "github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/domain/value-objects/name"
)

type Category struct {
	id            *vo_id.ID
	name          *vo_name.Name
	description   string
	status        bool
	notifications *notifications.Notifications
}

func NewCategory(id string, name string, description string) (*Category, error) {

	category := &Category{
		status:        true,
		description:   description,
		notifications: notifications.NewNotification(),
	}

	category.generateId(id)
	category.generateName(name)

	return category, nil
}

func (c *Category) generateId(id string) {
	categoryId, idError := vo_id.NewId(id)
	if idError != nil {
		c.notifications.AddError("id", idError.Error())
	} else {
		c.id = categoryId
	}
}

func (c *Category) generateName(name string) {
	categoryName, nameError := vo_name.NewName(name)
	if nameError != nil {
		c.notifications.AddError("name", nameError.Error())
	} else {
		c.name = categoryName
	}
}

func (c *Category) Id() string {
	return c.id.Value()
}

func (c *Category) Name() string {
	return c.name.Value()
}

func (c *Category) Activate() {
	c.status = true
}

func (c *Category) Deactivate() {
	c.status = false
}

func (c *Category) IsActive() bool {
	return c.status
}

func (c *Category) ChangeName(name string) error {
	newName, error := vo_name.NewName(name)

	if error != nil {
		return error
	}

	c.name = newName

	return nil
}

func (c *Category) Notifications() *notifications.Notifications {
	return c.notifications
}
