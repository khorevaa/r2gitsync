// Code generated by entc, DO NOT EDIT.

package storagecommit

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/khorevaa/r2gitsync/internal/services/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v uint) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNumber), v))
	})
}

// ConfigurationVersion applies equality check predicate on the "configuration_version" field. It's identical to ConfigurationVersionEQ.
func ConfigurationVersion(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfigurationVersion), v))
	})
}

// Author applies equality check predicate on the "author" field. It's identical to AuthorEQ.
func Author(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthor), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// Tag applies equality check predicate on the "tag" field. It's identical to TagEQ.
func Tag(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTag), v))
	})
}

// TagDescription applies equality check predicate on the "tag_description" field. It's identical to TagDescriptionEQ.
func TagDescription(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTagDescription), v))
	})
}

// CommitAt applies equality check predicate on the "commit_at" field. It's identical to CommitAtEQ.
func CommitAt(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCommitAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v uint) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNumber), v))
	})
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v uint) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNumber), v))
	})
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...uint) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNumber), v...))
	})
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...uint) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNumber), v...))
	})
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v uint) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNumber), v))
	})
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v uint) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNumber), v))
	})
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v uint) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNumber), v))
	})
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v uint) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNumber), v))
	})
}

// ConfigurationVersionEQ applies the EQ predicate on the "configuration_version" field.
func ConfigurationVersionEQ(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfigurationVersion), v))
	})
}

// ConfigurationVersionNEQ applies the NEQ predicate on the "configuration_version" field.
func ConfigurationVersionNEQ(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldConfigurationVersion), v))
	})
}

// ConfigurationVersionIn applies the In predicate on the "configuration_version" field.
func ConfigurationVersionIn(vs ...string) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldConfigurationVersion), v...))
	})
}

// ConfigurationVersionNotIn applies the NotIn predicate on the "configuration_version" field.
func ConfigurationVersionNotIn(vs ...string) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldConfigurationVersion), v...))
	})
}

// ConfigurationVersionGT applies the GT predicate on the "configuration_version" field.
func ConfigurationVersionGT(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldConfigurationVersion), v))
	})
}

// ConfigurationVersionGTE applies the GTE predicate on the "configuration_version" field.
func ConfigurationVersionGTE(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldConfigurationVersion), v))
	})
}

// ConfigurationVersionLT applies the LT predicate on the "configuration_version" field.
func ConfigurationVersionLT(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldConfigurationVersion), v))
	})
}

// ConfigurationVersionLTE applies the LTE predicate on the "configuration_version" field.
func ConfigurationVersionLTE(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldConfigurationVersion), v))
	})
}

// ConfigurationVersionContains applies the Contains predicate on the "configuration_version" field.
func ConfigurationVersionContains(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldConfigurationVersion), v))
	})
}

// ConfigurationVersionHasPrefix applies the HasPrefix predicate on the "configuration_version" field.
func ConfigurationVersionHasPrefix(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldConfigurationVersion), v))
	})
}

// ConfigurationVersionHasSuffix applies the HasSuffix predicate on the "configuration_version" field.
func ConfigurationVersionHasSuffix(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldConfigurationVersion), v))
	})
}

// ConfigurationVersionEqualFold applies the EqualFold predicate on the "configuration_version" field.
func ConfigurationVersionEqualFold(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldConfigurationVersion), v))
	})
}

// ConfigurationVersionContainsFold applies the ContainsFold predicate on the "configuration_version" field.
func ConfigurationVersionContainsFold(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldConfigurationVersion), v))
	})
}

// AuthorEQ applies the EQ predicate on the "author" field.
func AuthorEQ(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthor), v))
	})
}

// AuthorNEQ applies the NEQ predicate on the "author" field.
func AuthorNEQ(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAuthor), v))
	})
}

// AuthorIn applies the In predicate on the "author" field.
func AuthorIn(vs ...string) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAuthor), v...))
	})
}

// AuthorNotIn applies the NotIn predicate on the "author" field.
func AuthorNotIn(vs ...string) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAuthor), v...))
	})
}

// AuthorGT applies the GT predicate on the "author" field.
func AuthorGT(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAuthor), v))
	})
}

// AuthorGTE applies the GTE predicate on the "author" field.
func AuthorGTE(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAuthor), v))
	})
}

// AuthorLT applies the LT predicate on the "author" field.
func AuthorLT(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAuthor), v))
	})
}

// AuthorLTE applies the LTE predicate on the "author" field.
func AuthorLTE(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAuthor), v))
	})
}

// AuthorContains applies the Contains predicate on the "author" field.
func AuthorContains(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAuthor), v))
	})
}

// AuthorHasPrefix applies the HasPrefix predicate on the "author" field.
func AuthorHasPrefix(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAuthor), v))
	})
}

// AuthorHasSuffix applies the HasSuffix predicate on the "author" field.
func AuthorHasSuffix(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAuthor), v))
	})
}

// AuthorEqualFold applies the EqualFold predicate on the "author" field.
func AuthorEqualFold(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAuthor), v))
	})
}

// AuthorContainsFold applies the ContainsFold predicate on the "author" field.
func AuthorContainsFold(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAuthor), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// TagEQ applies the EQ predicate on the "tag" field.
func TagEQ(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTag), v))
	})
}

// TagNEQ applies the NEQ predicate on the "tag" field.
func TagNEQ(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTag), v))
	})
}

// TagIn applies the In predicate on the "tag" field.
func TagIn(vs ...string) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTag), v...))
	})
}

// TagNotIn applies the NotIn predicate on the "tag" field.
func TagNotIn(vs ...string) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTag), v...))
	})
}

// TagGT applies the GT predicate on the "tag" field.
func TagGT(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTag), v))
	})
}

// TagGTE applies the GTE predicate on the "tag" field.
func TagGTE(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTag), v))
	})
}

// TagLT applies the LT predicate on the "tag" field.
func TagLT(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTag), v))
	})
}

// TagLTE applies the LTE predicate on the "tag" field.
func TagLTE(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTag), v))
	})
}

// TagContains applies the Contains predicate on the "tag" field.
func TagContains(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTag), v))
	})
}

// TagHasPrefix applies the HasPrefix predicate on the "tag" field.
func TagHasPrefix(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTag), v))
	})
}

// TagHasSuffix applies the HasSuffix predicate on the "tag" field.
func TagHasSuffix(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTag), v))
	})
}

// TagEqualFold applies the EqualFold predicate on the "tag" field.
func TagEqualFold(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTag), v))
	})
}

// TagContainsFold applies the ContainsFold predicate on the "tag" field.
func TagContainsFold(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTag), v))
	})
}

// TagDescriptionEQ applies the EQ predicate on the "tag_description" field.
func TagDescriptionEQ(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTagDescription), v))
	})
}

// TagDescriptionNEQ applies the NEQ predicate on the "tag_description" field.
func TagDescriptionNEQ(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTagDescription), v))
	})
}

// TagDescriptionIn applies the In predicate on the "tag_description" field.
func TagDescriptionIn(vs ...string) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTagDescription), v...))
	})
}

// TagDescriptionNotIn applies the NotIn predicate on the "tag_description" field.
func TagDescriptionNotIn(vs ...string) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTagDescription), v...))
	})
}

// TagDescriptionGT applies the GT predicate on the "tag_description" field.
func TagDescriptionGT(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTagDescription), v))
	})
}

// TagDescriptionGTE applies the GTE predicate on the "tag_description" field.
func TagDescriptionGTE(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTagDescription), v))
	})
}

// TagDescriptionLT applies the LT predicate on the "tag_description" field.
func TagDescriptionLT(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTagDescription), v))
	})
}

// TagDescriptionLTE applies the LTE predicate on the "tag_description" field.
func TagDescriptionLTE(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTagDescription), v))
	})
}

// TagDescriptionContains applies the Contains predicate on the "tag_description" field.
func TagDescriptionContains(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTagDescription), v))
	})
}

// TagDescriptionHasPrefix applies the HasPrefix predicate on the "tag_description" field.
func TagDescriptionHasPrefix(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTagDescription), v))
	})
}

// TagDescriptionHasSuffix applies the HasSuffix predicate on the "tag_description" field.
func TagDescriptionHasSuffix(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTagDescription), v))
	})
}

// TagDescriptionEqualFold applies the EqualFold predicate on the "tag_description" field.
func TagDescriptionEqualFold(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTagDescription), v))
	})
}

// TagDescriptionContainsFold applies the ContainsFold predicate on the "tag_description" field.
func TagDescriptionContainsFold(v string) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTagDescription), v))
	})
}

// CommitAtEQ applies the EQ predicate on the "commit_at" field.
func CommitAtEQ(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCommitAt), v))
	})
}

// CommitAtNEQ applies the NEQ predicate on the "commit_at" field.
func CommitAtNEQ(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCommitAt), v))
	})
}

// CommitAtIn applies the In predicate on the "commit_at" field.
func CommitAtIn(vs ...time.Time) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCommitAt), v...))
	})
}

// CommitAtNotIn applies the NotIn predicate on the "commit_at" field.
func CommitAtNotIn(vs ...time.Time) predicate.StorageCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StorageCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCommitAt), v...))
	})
}

// CommitAtGT applies the GT predicate on the "commit_at" field.
func CommitAtGT(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCommitAt), v))
	})
}

// CommitAtGTE applies the GTE predicate on the "commit_at" field.
func CommitAtGTE(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCommitAt), v))
	})
}

// CommitAtLT applies the LT predicate on the "commit_at" field.
func CommitAtLT(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCommitAt), v))
	})
}

// CommitAtLTE applies the LTE predicate on the "commit_at" field.
func CommitAtLTE(v time.Time) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCommitAt), v))
	})
}

// HasStorage applies the HasEdge predicate on the "storage" edge.
func HasStorage() predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StorageTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, StorageTable, StorageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStorageWith applies the HasEdge predicate on the "storage" edge with a given conditions (other predicates).
func HasStorageWith(preds ...predicate.Storage) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StorageInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, StorageTable, StorageColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.StorageCommit) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.StorageCommit) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.StorageCommit) predicate.StorageCommit {
	return predicate.StorageCommit(func(s *sql.Selector) {
		p(s.Not())
	})
}
