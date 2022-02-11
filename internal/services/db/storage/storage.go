// Code generated by entc, DO NOT EDIT.

package storage

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the storage type in the database.
	Label = "storage"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldConnectionString holds the string denoting the connection_string field in the database.
	FieldConnectionString = "connection_string"
	// FieldDevelop holds the string denoting the develop field in the database.
	FieldDevelop = "develop"
	// FieldExtension holds the string denoting the extension field in the database.
	FieldExtension = "extension"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the storage in the database.
	Table = "storages"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "storages"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_storages"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "storages"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "storage_parent"
)

// Columns holds all SQL columns for storage fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldConnectionString,
	FieldDevelop,
	FieldExtension,
	FieldType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "storages"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"project_storages",
	"storage_parent",
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeHTTP Type = "http"
	TypeFile Type = "file"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeHTTP, TypeFile:
		return nil
	default:
		return fmt.Errorf("storage: invalid enum value for type field: %q", _type)
	}
}
