// Code generated by entc, DO NOT EDIT.

package coinaccountinfo

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the coinaccountinfo type in the database.
	Label = "coin_account_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCoinTypeID holds the string denoting the coin_type_id field in the database.
	FieldCoinTypeID = "coin_type_id"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldGeneratedBy holds the string denoting the generated_by field in the database.
	FieldGeneratedBy = "generated_by"
	// FieldPlatformHoldPrivateKey holds the string denoting the platform_hold_private_key field in the database.
	FieldPlatformHoldPrivateKey = "platform_hold_private_key"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the coinaccountinfo in the database.
	Table = "coin_account_infos"
)

// Columns holds all SQL columns for coinaccountinfo fields.
var Columns = []string{
	FieldID,
	FieldAppID,
	FieldUserID,
	FieldCoinTypeID,
	FieldAddress,
	FieldGeneratedBy,
	FieldPlatformHoldPrivateKey,
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

// GeneratedBy defines the type for the "generated_by" enum field.
type GeneratedBy string

// GeneratedBy values.
const (
	GeneratedByPlatform GeneratedBy = "platform"
	GeneratedByUser     GeneratedBy = "user"
)

func (gb GeneratedBy) String() string {
	return string(gb)
}

// GeneratedByValidator is a validator for the "generated_by" field enum values. It is called by the builders before save.
func GeneratedByValidator(gb GeneratedBy) error {
	switch gb {
	case GeneratedByPlatform, GeneratedByUser:
		return nil
	default:
		return fmt.Errorf("coinaccountinfo: invalid enum value for generated_by field: %q", gb)
	}
}
