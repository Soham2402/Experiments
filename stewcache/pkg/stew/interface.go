package stew


type BasicChunkTypes interface {
	int | string | float64 | float32
}

type BasicChunkHandler[T BasicChunkTypes] interface {
    Get(Key string, Value T) T
    Delete(Key string)
    Set(Key string, Value T)
}