// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CreditEntriesColumns holds the columns for the "credit_entries" table.
	CreditEntriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "namespace", Type: field.TypeString},
		{Name: "subject", Type: field.TypeString},
		{Name: "entry_type", Type: field.TypeEnum, Enums: []string{"GRANT", "VOID_GRANT", "RESET"}},
		{Name: "type", Type: field.TypeEnum, Nullable: true, Enums: []string{"USAGE"}},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true},
		{Name: "priority", Type: field.TypeUint8, Default: 1},
		{Name: "effective_at", Type: field.TypeTime},
		{Name: "expiration_period_duration", Type: field.TypeEnum, Nullable: true, Enums: []string{"HOUR", "DAY", "WEEK", "MONTH", "YEAR"}},
		{Name: "expiration_period_count", Type: field.TypeUint8, Nullable: true},
		{Name: "rollover_type", Type: field.TypeEnum, Nullable: true, Enums: []string{"ORIGINAL_AMOUNT", "REMAINING_AMOUNT"}},
		{Name: "rollover_max_amount", Type: field.TypeFloat64, Nullable: true},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "parent_id", Type: field.TypeString, Unique: true, Nullable: true, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "feature_id", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "char(26)"}},
	}
	// CreditEntriesTable holds the schema information for the "credit_entries" table.
	CreditEntriesTable = &schema.Table{
		Name:       "credit_entries",
		Columns:    CreditEntriesColumns,
		PrimaryKey: []*schema.Column{CreditEntriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "credit_entries_credit_entries_children",
				Columns:    []*schema.Column{CreditEntriesColumns[15]},
				RefColumns: []*schema.Column{CreditEntriesColumns[0]},
				OnDelete:   schema.Restrict,
			},
			{
				Symbol:     "credit_entries_features_credit_grants",
				Columns:    []*schema.Column{CreditEntriesColumns[16]},
				RefColumns: []*schema.Column{FeaturesColumns[0]},
				OnDelete:   schema.Restrict,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "creditentry_namespace_subject",
				Unique:  false,
				Columns: []*schema.Column{CreditEntriesColumns[3], CreditEntriesColumns[4]},
			},
		},
	}
	// FeaturesColumns holds the columns for the "features" table.
	FeaturesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "namespace", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "meter_slug", Type: field.TypeString},
		{Name: "meter_group_by_filters", Type: field.TypeJSON, Nullable: true},
		{Name: "archived", Type: field.TypeBool, Default: false},
	}
	// FeaturesTable holds the schema information for the "features" table.
	FeaturesTable = &schema.Table{
		Name:       "features",
		Columns:    FeaturesColumns,
		PrimaryKey: []*schema.Column{FeaturesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "feature_namespace_id",
				Unique:  false,
				Columns: []*schema.Column{FeaturesColumns[3], FeaturesColumns[0]},
			},
		},
	}
	// LedgersColumns holds the columns for the "ledgers" table.
	LedgersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "namespace", Type: field.TypeString},
		{Name: "subject", Type: field.TypeString},
		{Name: "highwatermark", Type: field.TypeTime},
	}
	// LedgersTable holds the schema information for the "ledgers" table.
	LedgersTable = &schema.Table{
		Name:       "ledgers",
		Columns:    LedgersColumns,
		PrimaryKey: []*schema.Column{LedgersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "ledger_namespace_subject",
				Unique:  true,
				Columns: []*schema.Column{LedgersColumns[3], LedgersColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CreditEntriesTable,
		FeaturesTable,
		LedgersTable,
	}
)

func init() {
	CreditEntriesTable.ForeignKeys[0].RefTable = CreditEntriesTable
	CreditEntriesTable.ForeignKeys[1].RefTable = FeaturesTable
}