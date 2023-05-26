package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Coffee struct {
	ent.Schema
}

func (Coffee) Fields() []ent.Field {
	return []ent.Field{
		field.Int("sugar").
			NonNegative(),
		field.Int("coffee").
			NonNegative(),
		field.Int("powdered_milk").
			NonNegative(),
		field.Int("coffee_mate").
			NonNegative(),
		field.Int("milk").
			NonNegative(),
		field.Int("water").
			NonNegative(),
		field.Float("rating").
			Range(0.0, 10.0),
	}
}

func (Coffee) Edges() []ent.Edge {
	return nil
}

func (Coffee) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
