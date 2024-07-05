// Code generated by entc, DO NOT EDIT.

package commentindex

import (
	"comment/app/service/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
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
func IDNotIn(ids ...int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
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
func IDGT(id int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ContentID applies equality check predicate on the "content_id" field. It's identical to ContentIDEQ.
func ContentID(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContentID), v))
	})
}

// ObjID applies equality check predicate on the "obj_id" field. It's identical to ObjIDEQ.
func ObjID(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldObjID), v))
	})
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOwnerID), v))
	})
}

// Root applies equality check predicate on the "root" field. It's identical to RootEQ.
func Root(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoot), v))
	})
}

// Parent applies equality check predicate on the "parent" field. It's identical to ParentEQ.
func Parent(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParent), v))
	})
}

// Floor applies equality check predicate on the "floor" field. It's identical to FloorEQ.
func Floor(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFloor), v))
	})
}

// Like applies equality check predicate on the "like" field. It's identical to LikeEQ.
func Like(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLike), v))
	})
}

// Hate applies equality check predicate on the "hate" field. It's identical to HateEQ.
func Hate(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHate), v))
	})
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v int8) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// NextFloor applies equality check predicate on the "next_floor" field. It's identical to NextFloorEQ.
func NextFloor(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNextFloor), v))
	})
}

// Count applies equality check predicate on the "count" field. It's identical to CountEQ.
func Count(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCount), v))
	})
}

// ContentIDEQ applies the EQ predicate on the "content_id" field.
func ContentIDEQ(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContentID), v))
	})
}

// ContentIDNEQ applies the NEQ predicate on the "content_id" field.
func ContentIDNEQ(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContentID), v))
	})
}

// ContentIDIn applies the In predicate on the "content_id" field.
func ContentIDIn(vs ...int64) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContentID), v...))
	})
}

// ContentIDNotIn applies the NotIn predicate on the "content_id" field.
func ContentIDNotIn(vs ...int64) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContentID), v...))
	})
}

// ContentIDGT applies the GT predicate on the "content_id" field.
func ContentIDGT(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContentID), v))
	})
}

// ContentIDGTE applies the GTE predicate on the "content_id" field.
func ContentIDGTE(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContentID), v))
	})
}

// ContentIDLT applies the LT predicate on the "content_id" field.
func ContentIDLT(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContentID), v))
	})
}

// ContentIDLTE applies the LTE predicate on the "content_id" field.
func ContentIDLTE(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContentID), v))
	})
}

// ObjIDEQ applies the EQ predicate on the "obj_id" field.
func ObjIDEQ(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldObjID), v))
	})
}

// ObjIDNEQ applies the NEQ predicate on the "obj_id" field.
func ObjIDNEQ(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldObjID), v))
	})
}

// ObjIDIn applies the In predicate on the "obj_id" field.
func ObjIDIn(vs ...int64) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldObjID), v...))
	})
}

// ObjIDNotIn applies the NotIn predicate on the "obj_id" field.
func ObjIDNotIn(vs ...int64) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldObjID), v...))
	})
}

// ObjIDGT applies the GT predicate on the "obj_id" field.
func ObjIDGT(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldObjID), v))
	})
}

// ObjIDGTE applies the GTE predicate on the "obj_id" field.
func ObjIDGTE(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldObjID), v))
	})
}

// ObjIDLT applies the LT predicate on the "obj_id" field.
func ObjIDLT(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldObjID), v))
	})
}

// ObjIDLTE applies the LTE predicate on the "obj_id" field.
func ObjIDLTE(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldObjID), v))
	})
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOwnerID), v))
	})
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOwnerID), v))
	})
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...int64) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOwnerID), v...))
	})
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...int64) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOwnerID), v...))
	})
}

// OwnerIDGT applies the GT predicate on the "owner_id" field.
func OwnerIDGT(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOwnerID), v))
	})
}

// OwnerIDGTE applies the GTE predicate on the "owner_id" field.
func OwnerIDGTE(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOwnerID), v))
	})
}

// OwnerIDLT applies the LT predicate on the "owner_id" field.
func OwnerIDLT(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOwnerID), v))
	})
}

// OwnerIDLTE applies the LTE predicate on the "owner_id" field.
func OwnerIDLTE(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOwnerID), v))
	})
}

// RootEQ applies the EQ predicate on the "root" field.
func RootEQ(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoot), v))
	})
}

// RootNEQ applies the NEQ predicate on the "root" field.
func RootNEQ(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRoot), v))
	})
}

// RootIn applies the In predicate on the "root" field.
func RootIn(vs ...int64) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRoot), v...))
	})
}

// RootNotIn applies the NotIn predicate on the "root" field.
func RootNotIn(vs ...int64) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRoot), v...))
	})
}

// RootGT applies the GT predicate on the "root" field.
func RootGT(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRoot), v))
	})
}

// RootGTE applies the GTE predicate on the "root" field.
func RootGTE(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRoot), v))
	})
}

// RootLT applies the LT predicate on the "root" field.
func RootLT(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRoot), v))
	})
}

// RootLTE applies the LTE predicate on the "root" field.
func RootLTE(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRoot), v))
	})
}

// ParentEQ applies the EQ predicate on the "parent" field.
func ParentEQ(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParent), v))
	})
}

// ParentNEQ applies the NEQ predicate on the "parent" field.
func ParentNEQ(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldParent), v))
	})
}

// ParentIn applies the In predicate on the "parent" field.
func ParentIn(vs ...int64) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldParent), v...))
	})
}

// ParentNotIn applies the NotIn predicate on the "parent" field.
func ParentNotIn(vs ...int64) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldParent), v...))
	})
}

// ParentGT applies the GT predicate on the "parent" field.
func ParentGT(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldParent), v))
	})
}

// ParentGTE applies the GTE predicate on the "parent" field.
func ParentGTE(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldParent), v))
	})
}

// ParentLT applies the LT predicate on the "parent" field.
func ParentLT(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldParent), v))
	})
}

// ParentLTE applies the LTE predicate on the "parent" field.
func ParentLTE(v int64) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldParent), v))
	})
}

// FloorEQ applies the EQ predicate on the "floor" field.
func FloorEQ(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFloor), v))
	})
}

// FloorNEQ applies the NEQ predicate on the "floor" field.
func FloorNEQ(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFloor), v))
	})
}

// FloorIn applies the In predicate on the "floor" field.
func FloorIn(vs ...int32) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFloor), v...))
	})
}

// FloorNotIn applies the NotIn predicate on the "floor" field.
func FloorNotIn(vs ...int32) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFloor), v...))
	})
}

// FloorGT applies the GT predicate on the "floor" field.
func FloorGT(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFloor), v))
	})
}

// FloorGTE applies the GTE predicate on the "floor" field.
func FloorGTE(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFloor), v))
	})
}

// FloorLT applies the LT predicate on the "floor" field.
func FloorLT(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFloor), v))
	})
}

// FloorLTE applies the LTE predicate on the "floor" field.
func FloorLTE(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFloor), v))
	})
}

// LikeEQ applies the EQ predicate on the "like" field.
func LikeEQ(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLike), v))
	})
}

// LikeNEQ applies the NEQ predicate on the "like" field.
func LikeNEQ(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLike), v))
	})
}

// LikeIn applies the In predicate on the "like" field.
func LikeIn(vs ...int32) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLike), v...))
	})
}

// LikeNotIn applies the NotIn predicate on the "like" field.
func LikeNotIn(vs ...int32) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLike), v...))
	})
}

// LikeGT applies the GT predicate on the "like" field.
func LikeGT(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLike), v))
	})
}

// LikeGTE applies the GTE predicate on the "like" field.
func LikeGTE(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLike), v))
	})
}

// LikeLT applies the LT predicate on the "like" field.
func LikeLT(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLike), v))
	})
}

// LikeLTE applies the LTE predicate on the "like" field.
func LikeLTE(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLike), v))
	})
}

// HateEQ applies the EQ predicate on the "hate" field.
func HateEQ(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHate), v))
	})
}

// HateNEQ applies the NEQ predicate on the "hate" field.
func HateNEQ(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHate), v))
	})
}

// HateIn applies the In predicate on the "hate" field.
func HateIn(vs ...int32) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHate), v...))
	})
}

// HateNotIn applies the NotIn predicate on the "hate" field.
func HateNotIn(vs ...int32) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHate), v...))
	})
}

// HateGT applies the GT predicate on the "hate" field.
func HateGT(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHate), v))
	})
}

// HateGTE applies the GTE predicate on the "hate" field.
func HateGTE(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHate), v))
	})
}

// HateLT applies the LT predicate on the "hate" field.
func HateLT(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHate), v))
	})
}

// HateLTE applies the LTE predicate on the "hate" field.
func HateLTE(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHate), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v int8) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v int8) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...int8) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...int8) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v int8) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldState), v))
	})
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v int8) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldState), v))
	})
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v int8) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldState), v))
	})
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v int8) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldState), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// NextFloorEQ applies the EQ predicate on the "next_floor" field.
func NextFloorEQ(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNextFloor), v))
	})
}

// NextFloorNEQ applies the NEQ predicate on the "next_floor" field.
func NextFloorNEQ(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNextFloor), v))
	})
}

// NextFloorIn applies the In predicate on the "next_floor" field.
func NextFloorIn(vs ...int32) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNextFloor), v...))
	})
}

// NextFloorNotIn applies the NotIn predicate on the "next_floor" field.
func NextFloorNotIn(vs ...int32) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNextFloor), v...))
	})
}

// NextFloorGT applies the GT predicate on the "next_floor" field.
func NextFloorGT(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNextFloor), v))
	})
}

// NextFloorGTE applies the GTE predicate on the "next_floor" field.
func NextFloorGTE(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNextFloor), v))
	})
}

// NextFloorLT applies the LT predicate on the "next_floor" field.
func NextFloorLT(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNextFloor), v))
	})
}

// NextFloorLTE applies the LTE predicate on the "next_floor" field.
func NextFloorLTE(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNextFloor), v))
	})
}

// NextFloorIsNil applies the IsNil predicate on the "next_floor" field.
func NextFloorIsNil() predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldNextFloor)))
	})
}

// NextFloorNotNil applies the NotNil predicate on the "next_floor" field.
func NextFloorNotNil() predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldNextFloor)))
	})
}

// CountEQ applies the EQ predicate on the "count" field.
func CountEQ(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCount), v))
	})
}

// CountNEQ applies the NEQ predicate on the "count" field.
func CountNEQ(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCount), v))
	})
}

// CountIn applies the In predicate on the "count" field.
func CountIn(vs ...int32) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCount), v...))
	})
}

// CountNotIn applies the NotIn predicate on the "count" field.
func CountNotIn(vs ...int32) predicate.CommentIndex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommentIndex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCount), v...))
	})
}

// CountGT applies the GT predicate on the "count" field.
func CountGT(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCount), v))
	})
}

// CountGTE applies the GTE predicate on the "count" field.
func CountGTE(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCount), v))
	})
}

// CountLT applies the LT predicate on the "count" field.
func CountLT(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCount), v))
	})
}

// CountLTE applies the LTE predicate on the "count" field.
func CountLTE(v int32) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCount), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CommentIndex) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CommentIndex) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
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
func Not(p predicate.CommentIndex) predicate.CommentIndex {
	return predicate.CommentIndex(func(s *sql.Selector) {
		p(s.Not())
	})
}
