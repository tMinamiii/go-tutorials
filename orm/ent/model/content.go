package ent_model

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Content struct {
	ent.Schema
}

func (Content) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id"),
		field.String("content"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Content) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("user_id"),
	}
}
