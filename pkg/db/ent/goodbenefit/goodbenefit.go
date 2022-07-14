// Code generated by ent, DO NOT EDIT.

package goodbenefit

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the goodbenefit type in the database.
	Label = "good_benefit"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldBenefitAccountID holds the string denoting the benefit_account_id field in the database.
	FieldBenefitAccountID = "benefit_account_id"
	// FieldBenefitIntervalHours holds the string denoting the benefit_interval_hours field in the database.
	FieldBenefitIntervalHours = "benefit_interval_hours"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the goodbenefit in the database.
	Table = "good_benefits"
)

// Columns holds all SQL columns for goodbenefit fields.
var Columns = []string{
	FieldID,
	FieldGoodID,
	FieldBenefitAccountID,
	FieldBenefitIntervalHours,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
