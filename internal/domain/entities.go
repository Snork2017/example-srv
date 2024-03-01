package domain

// Entity inside the Domain Layer
// we separate the Entity from its representation in the database
// Its purpose to encapsulate essential business logic models.
// Entities do not necessarily replicate the database structure.
// Here we validate entities
type ExampleData struct {
	Field  string
	Field2 string
}
