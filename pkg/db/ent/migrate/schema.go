// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CoinAccountInfosColumns holds the columns for the "coin_account_infos" table.
	CoinAccountInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "coin_type_id", Type: field.TypeUUID},
		{Name: "address", Type: field.TypeString},
		{Name: "generated_by", Type: field.TypeEnum, Enums: []string{"platform", "user"}},
		{Name: "used_for", Type: field.TypeEnum, Enums: []string{"benefit", "offline", "user", "paying", "withdraw"}},
		{Name: "idle", Type: field.TypeBool},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// CoinAccountInfosTable holds the schema information for the "coin_account_infos" table.
	CoinAccountInfosTable = &schema.Table{
		Name:       "coin_account_infos",
		Columns:    CoinAccountInfosColumns,
		PrimaryKey: []*schema.Column{CoinAccountInfosColumns[0]},
	}
	// CoinAccountTransactionsColumns holds the columns for the "coin_account_transactions" table.
	CoinAccountTransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "from_address_id", Type: field.TypeUUID},
		{Name: "to_address_id", Type: field.TypeUUID},
		{Name: "coin_type_id", Type: field.TypeUUID},
		{Name: "amount", Type: field.TypeUint64},
		{Name: "message", Type: field.TypeString},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"wait", "paying", "successful", "fail"}},
		{Name: "chain_transaction_id", Type: field.TypeString},
		{Name: "platform_transaction_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// CoinAccountTransactionsTable holds the schema information for the "coin_account_transactions" table.
	CoinAccountTransactionsTable = &schema.Table{
		Name:       "coin_account_transactions",
		Columns:    CoinAccountTransactionsColumns,
		PrimaryKey: []*schema.Column{CoinAccountTransactionsColumns[0]},
	}
	// PlatformBenefitsColumns holds the columns for the "platform_benefits" table.
	PlatformBenefitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "benefit_account_id", Type: field.TypeUUID},
		{Name: "amount", Type: field.TypeUint64},
		{Name: "last_benefit_timestamp", Type: field.TypeUint32, Unique: true},
		{Name: "chain_transaction_id", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// PlatformBenefitsTable holds the schema information for the "platform_benefits" table.
	PlatformBenefitsTable = &schema.Table{
		Name:       "platform_benefits",
		Columns:    PlatformBenefitsColumns,
		PrimaryKey: []*schema.Column{PlatformBenefitsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "platformbenefit_good_id_chain_transaction_id",
				Unique:  true,
				Columns: []*schema.Column{PlatformBenefitsColumns[1], PlatformBenefitsColumns[5]},
			},
		},
	}
	// PlatformSettingsColumns holds the columns for the "platform_settings" table.
	PlatformSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "good_id", Type: field.TypeUUID, Unique: true},
		{Name: "benefit_account_id", Type: field.TypeUUID},
		{Name: "platform_offline_account_id", Type: field.TypeUUID},
		{Name: "user_online_account_id", Type: field.TypeUUID},
		{Name: "user_offline_account_id", Type: field.TypeUUID},
		{Name: "benefit_interval_hours", Type: field.TypeInt32},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// PlatformSettingsTable holds the schema information for the "platform_settings" table.
	PlatformSettingsTable = &schema.Table{
		Name:       "platform_settings",
		Columns:    PlatformSettingsColumns,
		PrimaryKey: []*schema.Column{PlatformSettingsColumns[0]},
	}
	// UserBenefitsColumns holds the columns for the "user_benefits" table.
	UserBenefitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "order_id", Type: field.TypeUUID},
		{Name: "amount", Type: field.TypeUint64},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// UserBenefitsTable holds the schema information for the "user_benefits" table.
	UserBenefitsTable = &schema.Table{
		Name:       "user_benefits",
		Columns:    UserBenefitsColumns,
		PrimaryKey: []*schema.Column{UserBenefitsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CoinAccountInfosTable,
		CoinAccountTransactionsTable,
		PlatformBenefitsTable,
		PlatformSettingsTable,
		UserBenefitsTable,
	}
)

func init() {
}
