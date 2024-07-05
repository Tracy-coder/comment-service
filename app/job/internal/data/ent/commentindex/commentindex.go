// Code generated by entc, DO NOT EDIT.

package commentindex

import (
	"time"
)

const (
	// Label holds the string label denoting the commentindex type in the database.
	Label = "comment_index"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldContentID holds the string denoting the content_id field in the database.
	FieldContentID = "content_id"
	// FieldObjID holds the string denoting the obj_id field in the database.
	FieldObjID = "obj_id"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// FieldRoot holds the string denoting the root field in the database.
	FieldRoot = "root"
	// FieldParent holds the string denoting the parent field in the database.
	FieldParent = "parent"
	// FieldFloor holds the string denoting the floor field in the database.
	FieldFloor = "floor"
	// FieldLike holds the string denoting the like field in the database.
	FieldLike = "like"
	// FieldHate holds the string denoting the hate field in the database.
	FieldHate = "hate"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldNextFloor holds the string denoting the next_floor field in the database.
	FieldNextFloor = "next_floor"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// Table holds the table name of the commentindex in the database.
	Table = "comment_index"
)

// Columns holds all SQL columns for commentindex fields.
var Columns = []string{
	FieldID,
	FieldContentID,
	FieldObjID,
	FieldOwnerID,
	FieldRoot,
	FieldParent,
	FieldFloor,
	FieldLike,
	FieldHate,
	FieldState,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldNextFloor,
	FieldCount,
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
	// DefaultLike holds the default value on creation for the "like" field.
	DefaultLike int32
	// DefaultHate holds the default value on creation for the "hate" field.
	DefaultHate int32
	// DefaultState holds the default value on creation for the "state" field.
	DefaultState int8
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultCount holds the default value on creation for the "count" field.
	DefaultCount int32
)