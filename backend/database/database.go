package database

// Datastore interface is exactly what it says on the tin.
//
// The API separates what can be stored from how it is stored.
type Datastore interface {
	AddTask(task Task) error
	Close()
}
