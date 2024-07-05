// Code generated by entc, DO NOT EDIT.

package ent

import (
	"comment/app/service/internal/data/ent/commentcontent"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// CommentContent is the model entity for the CommentContent schema.
type CommentContent struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// ObjID holds the value of the "obj_id" field.
	// 作品id
	ObjID int64 `json:"obj_id,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	// 作者
	OwnerID int64 `json:"owner_id,omitempty"`
	// Root holds the value of the "root" field.
	// 根评论,为0表示该评论即为根评论
	Root int64 `json:"root,omitempty"`
	// Floor holds the value of the "floor" field.
	// 该评论的楼层,根评论为0
	Floor int32 `json:"floor,omitempty"`
	// Message holds the value of the "message" field.
	// 评论内容
	Message string `json:"message,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	// 创建时间
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	// 修改时间
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CommentContent) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case commentcontent.FieldID, commentcontent.FieldObjID, commentcontent.FieldOwnerID, commentcontent.FieldRoot, commentcontent.FieldFloor:
			values[i] = new(sql.NullInt64)
		case commentcontent.FieldMessage:
			values[i] = new(sql.NullString)
		case commentcontent.FieldCreatedAt, commentcontent.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CommentContent", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CommentContent fields.
func (cc *CommentContent) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case commentcontent.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cc.ID = int64(value.Int64)
		case commentcontent.FieldObjID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field obj_id", values[i])
			} else if value.Valid {
				cc.ObjID = value.Int64
			}
		case commentcontent.FieldOwnerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				cc.OwnerID = value.Int64
			}
		case commentcontent.FieldRoot:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field root", values[i])
			} else if value.Valid {
				cc.Root = value.Int64
			}
		case commentcontent.FieldFloor:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field floor", values[i])
			} else if value.Valid {
				cc.Floor = int32(value.Int64)
			}
		case commentcontent.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				cc.Message = value.String
			}
		case commentcontent.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cc.CreatedAt = value.Time
			}
		case commentcontent.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cc.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CommentContent.
// Note that you need to call CommentContent.Unwrap() before calling this method if this CommentContent
// was returned from a transaction, and the transaction was committed or rolled back.
func (cc *CommentContent) Update() *CommentContentUpdateOne {
	return (&CommentContentClient{config: cc.config}).UpdateOne(cc)
}

// Unwrap unwraps the CommentContent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cc *CommentContent) Unwrap() *CommentContent {
	tx, ok := cc.config.driver.(*txDriver)
	if !ok {
		panic("ent: CommentContent is not a transactional entity")
	}
	cc.config.driver = tx.drv
	return cc
}

// String implements the fmt.Stringer.
func (cc *CommentContent) String() string {
	var builder strings.Builder
	builder.WriteString("CommentContent(")
	builder.WriteString(fmt.Sprintf("id=%v", cc.ID))
	builder.WriteString(", obj_id=")
	builder.WriteString(fmt.Sprintf("%v", cc.ObjID))
	builder.WriteString(", owner_id=")
	builder.WriteString(fmt.Sprintf("%v", cc.OwnerID))
	builder.WriteString(", root=")
	builder.WriteString(fmt.Sprintf("%v", cc.Root))
	builder.WriteString(", floor=")
	builder.WriteString(fmt.Sprintf("%v", cc.Floor))
	builder.WriteString(", message=")
	builder.WriteString(cc.Message)
	builder.WriteString(", created_at=")
	builder.WriteString(cc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(cc.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// CommentContents is a parsable slice of CommentContent.
type CommentContents []*CommentContent

func (cc CommentContents) config(cfg config) {
	for _i := range cc {
		cc[_i].config = cfg
	}
}
