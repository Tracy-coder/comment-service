package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CommentSubject holds the schema definition for the CommentSubject entity.
type CommentSubject struct {
	ent.Schema
}

func (CommentSubject) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "comment_subject",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
	}
}

// Fields of the SensorData.
func (CommentSubject) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			SchemaType(map[string]string{
				dialect.MySQL:    "BIGINT",
				dialect.Postgres: "BIGSERIAL",
			}),

		field.Int64("owner_id").
			Comment("作者"),

		field.Int32("count").
			Comment("评论数").Default(0),

		field.Int32("next_floor").
			Comment("下一个楼层数").Default(1),

		field.Int8("state").
			Comment("状态,1为隐藏,0为正常").Default(0),

		field.Time("created_at").
			Comment("创建时间").Default(time.Now),

		field.Time("updated_at").
			Comment("修改时间").Optional(),
	}
}

// Edges of the SensorData.
func (CommentSubject) Edges() []ent.Edge {
	return nil
}
