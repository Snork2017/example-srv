package domain

// Entity inside the Domain Layer
// we separate the Entity from its representation in the database
// Its intended purpose is not to mirror the database schema but to encapsulate essential business logic
// Entities do not necessarily replicate the database structure.
// Here we validate entities
type ExampleData struct {
	Field string 
	Field2 string 
}