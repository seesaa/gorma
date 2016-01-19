package dsl

import (
	"strconv"

	"github.com/bketelsen/gorma"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/design/dsl"
)

// RelationalModel is the DSL that represents a Relational Model
// Examples and more docs here later
func RelationalModel(name string, modeledType *design.UserTypeDefinition, dsl func()) {
	// We can't rely on this being run first, any of the top level DSL could run
	// in any order. The top level DSLs are API, Version, Resource, MediaType and Type.
	// The first one to be called executes InitDesign.
	checkInit()
	if s, ok := relationalStoreDefinition(true); ok {
		if s.RelationalModels == nil {
			s.RelationalModels = make(map[string]*gorma.RelationalModelDefinition)
		}
		models, ok := s.RelationalModels[name]
		if !ok {
			models = &gorma.RelationalModelDefinition{
				Name:             name,
				DefinitionDSL:    dsl,
				Parent:           s,
				ModeledType:      modeledType,
				RelationalFields: make(map[string]*gorma.RelationalFieldDefinition),
			}
		}
		models.PopulateFromModeledType()
		s.RelationalModels[name] = models
	}

}

// TableName creates a TableName() function that returns
// the name of the table to query. Useful for pre-existing
// schemas
func TableName(d string) {
	if r, ok := relationalModelDefinition(false); ok {
		r.TableName = d
	}
}

// Alias overrides the name of the sql store's table or field
func Alias(d string) {
	if r, ok := relationalModelDefinition(false); ok {
		r.Alias = d
	} else if f, ok := relationalFieldDefinition(false); ok {
		f.Alias = d
	}
}

// Cached caches the models for `duration` seconds
func Cached(d string) {
	if r, ok := relationalModelDefinition(false); ok {
		r.Cached = true
		dur, err := strconv.Atoi(d)
		if err != nil {
			dsl.ReportError("Duration %s couldn't be parsed as integer", d)
		}
		r.CacheDuration = dur
	}
}

// NoMedia sets a boolean flag that prevents the generation
// of media helpers
func NoMedia() {
	if r, ok := relationalModelDefinition(false); ok {
		r.NoMedia = true
	}
}

// Roler sets a boolean flag that cause the generation of a
// Role() function that returns the model's Role value
// Requires a field in the model named Role, type String
func Roler() {
	if r, ok := relationalModelDefinition(false); ok {
		r.Roler = true
	}
}

// DynamicTableName sets a boolean flag that cause the generator
// generate function definitions in the database models that specify
// the name of the database table.  Useful when using multiple tables
// with different names but same schema e.g. Users, AdminUsers
func DynamicTableName() {
	if r, ok := relationalModelDefinition(false); ok {
		r.DynamicTableName = true
	}
}

// SQLTag sets the model's struct tag `sql` value
// for indexing and other purposes
func SQLTag(d string) {
	if r, ok := relationalModelDefinition(false); ok {
		r.SQLTag = d
	} else if f, ok := relationalFieldDefinition(false); ok {
		f.SQLTag = d
	}
}
