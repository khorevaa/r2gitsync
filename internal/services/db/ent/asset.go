// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/khorevaa/r2gitsync/internal/services/db/ent/asset"
)

// Asset is the model entity for the Asset schema.
type Asset struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// FileName holds the value of the "file_name" field.
	FileName string `json:"file_name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Asset) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case asset.FieldFileName:
			values[i] = new(sql.NullString)
		case asset.FieldCreatedAt, asset.FieldUpdatedAt, asset.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case asset.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Asset", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Asset fields.
func (a *Asset) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case asset.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case asset.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case asset.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case asset.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = new(time.Time)
				*a.DeletedAt = value.Time
			}
		case asset.FieldFileName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_name", values[i])
			} else if value.Valid {
				a.FileName = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Asset.
// Note that you need to call Asset.Unwrap() before calling this method if this Asset
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Asset) Update() *AssetUpdateOne {
	return (&AssetClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Asset entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Asset) Unwrap() *Asset {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Asset is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Asset) String() string {
	var builder strings.Builder
	builder.WriteString("Asset(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	if v := a.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", file_name=")
	builder.WriteString(a.FileName)
	builder.WriteByte(')')
	return builder.String()
}

// Assets is a parsable slice of Asset.
type Assets []*Asset

func (a Assets) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
