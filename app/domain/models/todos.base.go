// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

// Todo is an object representing the database table.
type Todo struct {
	ID     uint
	Text   string
	Done   bool
	UserID uint
}