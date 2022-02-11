// Code generated by entc, DO NOT EDIT.

package pluginversion

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the pluginversion type in the database.
	Label = "plugin_version"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldBroken holds the string denoting the broken field in the database.
	FieldBroken = "broken"
	// EdgePlugin holds the string denoting the plugin edge name in mutations.
	EdgePlugin = "plugin"
	// Table holds the table name of the pluginversion in the database.
	Table = "plugin_versions"
	// PluginTable is the table that holds the plugin relation/edge.
	PluginTable = "plugin_versions"
	// PluginInverseTable is the table name for the Plugin entity.
	// It exists in this package in order to avoid circular dependency with the "plugin" package.
	PluginInverseTable = "plugins"
	// PluginColumn is the table column denoting the plugin relation/edge.
	PluginColumn = "plugin_version_plugin"
)

// Columns holds all SQL columns for pluginversion fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldVersion,
	FieldDescription,
	FieldBroken,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "plugin_versions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"plugin_version_plugin",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
