// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CommentContentColumns holds the columns for the "comment_content" table.
	CommentContentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, SchemaType: map[string]string{"mysql": "BIGINT", "postgres": "BIGSERIAL"}},
		{Name: "obj_id", Type: field.TypeInt64},
		{Name: "owner_id", Type: field.TypeInt64},
		{Name: "root", Type: field.TypeInt64},
		{Name: "floor", Type: field.TypeInt32},
		{Name: "message", Type: field.TypeString, Default: ""},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// CommentContentTable holds the schema information for the "comment_content" table.
	CommentContentTable = &schema.Table{
		Name:       "comment_content",
		Columns:    CommentContentColumns,
		PrimaryKey: []*schema.Column{CommentContentColumns[0]},
	}
	// CommentIndexColumns holds the columns for the "comment_index" table.
	CommentIndexColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, SchemaType: map[string]string{"mysql": "BIGINT", "postgres": "BIGSERIAL"}},
		{Name: "content_id", Type: field.TypeInt64, SchemaType: map[string]string{"mysql": "BIGINT", "postgres": "BIGSERIAL"}},
		{Name: "obj_id", Type: field.TypeInt64},
		{Name: "owner_id", Type: field.TypeInt64},
		{Name: "root", Type: field.TypeInt64},
		{Name: "parent", Type: field.TypeInt64},
		{Name: "floor", Type: field.TypeInt32},
		{Name: "like", Type: field.TypeInt32, Default: 0},
		{Name: "hate", Type: field.TypeInt32, Default: 0},
		{Name: "state", Type: field.TypeInt8, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "next_floor", Type: field.TypeInt32, Nullable: true},
		{Name: "count", Type: field.TypeInt32, Default: 0},
	}
	// CommentIndexTable holds the schema information for the "comment_index" table.
	CommentIndexTable = &schema.Table{
		Name:       "comment_index",
		Columns:    CommentIndexColumns,
		PrimaryKey: []*schema.Column{CommentIndexColumns[0]},
	}
	// CommentSubjectColumns holds the columns for the "comment_subject" table.
	CommentSubjectColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, SchemaType: map[string]string{"mysql": "BIGINT", "postgres": "BIGSERIAL"}},
		{Name: "owner_id", Type: field.TypeInt64},
		{Name: "count", Type: field.TypeInt32, Default: 0},
		{Name: "next_floor", Type: field.TypeInt32, Default: 1},
		{Name: "state", Type: field.TypeInt8, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// CommentSubjectTable holds the schema information for the "comment_subject" table.
	CommentSubjectTable = &schema.Table{
		Name:       "comment_subject",
		Columns:    CommentSubjectColumns,
		PrimaryKey: []*schema.Column{CommentSubjectColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CommentContentTable,
		CommentIndexTable,
		CommentSubjectTable,
	}
)

func init() {
	CommentContentTable.Annotation = &entsql.Annotation{
		Table:     "comment_content",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	CommentIndexTable.Annotation = &entsql.Annotation{
		Table:     "comment_index",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	CommentSubjectTable.Annotation = &entsql.Annotation{
		Table:     "comment_subject",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
}