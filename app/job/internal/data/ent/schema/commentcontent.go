package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CommentContent holds the schema definition for the CommentContent entity.
type CommentContent struct {
	ent.Schema
}

func (CommentContent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "comment_content",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
	}
}

//	`comment_id` bigint PRIMARY KEY,
//	`obj_id` bigint NOT NULL,
//	`owner_id` bigint NOT NULL,
//	`root` bigint NOT NULL DEFAULT 0,
//	`message` varchar(20) NOT NULL DEFAULT '',
//	`created_at` timestamp NOT NULL DEFAULT (now()),
//	`updated_at` timestamp
//
// Fields of the SensorData.
func (CommentContent) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
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

		field.Int32("floor").
			Comment("该评论的楼层,根评论为0"),

		field.String("message").
			Comment("评论内容").Default(""),

		field.Time("created_at").
			Comment("创建时间").Default(time.Now),

		field.Time("updated_at").
			Comment("修改时间").Optional(),
	}
}

// Edges of the SensorData.
func (CommentContent) Edges() []ent.Edge {
	return nil
}
