package permission

import (
	"fmt"

	"gorm.io/gorm"
)

// TableName ...
const TableName = "permissions"

// Permission struct
type Permission struct {
	ID   int
	Name string
}

// Permissions using for many permissions
type Permissions []Permission

// BeforeCreate ..
func (permission *Permission) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("beforeCreate Called")
	return nil
}

// AfterCreate ..
func (permission *Permission) AfterCreate(tx *gorm.DB) error {
	fmt.Println("afterCreate Called")
	return nil
}
