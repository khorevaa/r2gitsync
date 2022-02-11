// Code generated by entc, DO NOT EDIT.

package project

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the project type in the database.
	Label = "project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeStorages holds the string denoting the storages edge name in mutations.
	EdgeStorages = "storages"
	// EdgeMasterStorage holds the string denoting the master_storage edge name in mutations.
	EdgeMasterStorage = "master_storage"
	// EdgeDevelopStorage holds the string denoting the develop_storage edge name in mutations.
	EdgeDevelopStorage = "develop_storage"
	// Table holds the table name of the project in the database.
	Table = "projects"
	// StoragesTable is the table that holds the storages relation/edge.
	StoragesTable = "storages"
	// StoragesInverseTable is the table name for the Storage entity.
	// It exists in this package in order to avoid circular dependency with the "storage" package.
	StoragesInverseTable = "storages"
	// StoragesColumn is the table column denoting the storages relation/edge.
	StoragesColumn = "project_storages"
	// MasterStorageTable is the table that holds the master_storage relation/edge.
	MasterStorageTable = "projects"
	// MasterStorageInverseTable is the table name for the Storage entity.
	// It exists in this package in order to avoid circular dependency with the "storage" package.
	MasterStorageInverseTable = "storages"
	// MasterStorageColumn is the table column denoting the master_storage relation/edge.
	MasterStorageColumn = "project_master_storage"
	// DevelopStorageTable is the table that holds the develop_storage relation/edge.
	DevelopStorageTable = "projects"
	// DevelopStorageInverseTable is the table name for the Storage entity.
	// It exists in this package in order to avoid circular dependency with the "storage" package.
	DevelopStorageInverseTable = "storages"
	// DevelopStorageColumn is the table column denoting the develop_storage relation/edge.
	DevelopStorageColumn = "project_develop_storage"
)

// Columns holds all SQL columns for project fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldCode,
	FieldName,
	FieldDescription,
	FieldType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "projects"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"project_master_storage",
	"project_develop_storage",
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
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeConfiguration Type = "configuration"
	TypeExtension     Type = "extension"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeConfiguration, TypeExtension:
		return nil
	default:
		return fmt.Errorf("project: invalid enum value for type field: %q", _type)
	}
}
