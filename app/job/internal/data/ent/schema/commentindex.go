package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CommentIndex holds the schema definition for the CommentIndex entity.
type CommentIndex struct {
	ent.Schema
}

func (CommentIndex) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "comment_index",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
	}
}

// Fields of the SensorData.
func (CommentIndex) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			SchemaType(map[string]string{
				dialect.MySQL:    "BIGINT",
				dialect.Postgres: "BIGSERIAL",
			}),
		field.Int64("content_id").
			SchemaType(map[string]string{
				dialect.MySQL:    "BIGINT",
				dialect.Postgres: "BIGSERIAL",
			}),

		field.Int64("obj_id").
			Comment("作品id"),

		field.Int64("owner_id").
			Comment("作者"),

		field.Int64("root").
			Comment("根评论,为0表示该评论即为根评论"),
		field.Int64("parent").
			Comment("回复的是哪条评论,为0表示该评论即为根评论"),

		field.Int32("floor").
			Comment("该评论的楼层,根评论为0"),

		field.Int32("like").
			Comment("like").Default(0),

		field.Int32("hate").
			Comment("hate").Default(0),

		field.Int8("state").
			Comment("状态,1为隐藏,0为正常").Default(0),

		field.Time("created_at").
			Comment("创建时间").Default(time.Now),

		field.Time("updated_at").
			Comment("修改时间").Optional(),
		field.Int32("next_floor").Comment("主评论楼层才有,标识子评论的下一楼层").Optional(),
		field.Int32("count").
			Comment("评论数").Default(0),
	}
}

// Edges of the SensorData.
func (CommentIndex) Edges() []ent.Edge {
	return nil
}
