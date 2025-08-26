package stew

// BasicChunkTypes is a type constraint that defines the basic types that can be used
// with the BasicChunkHandler interface. It includes the following types:
//   - int
//   - string
//   - float64
//   - float32
type BasicChunkTypes interface {
	int | string | float64 | float32
}

// BasicChunkHandler is a generic interface that defines the basic operations
// for handling chunks of data with a specific type T that satisfies BasicChunkTypes.
//
// The interface provides methods for getting, setting, and deleting key-value pairs
// where the value is of type T.
type BasicChunkHandler[T BasicChunkTypes] interface {
	// Get retrieves the value associated with the given key.
	// It returns the value if found, otherwise returns the zero value of type T.
	Get(Key string, Value T) T

	// Delete removes the key-value pair associated with the given key.
	// If the key does not exist, it does nothing.
	Delete(Key string)

	// Set associates the given value with the specified key.
	// If the key already exists, its value will be updated.
	Set(Key string, Value T)
}
