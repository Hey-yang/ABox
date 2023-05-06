package Test

type toString interface {
	String() string
}

type GetKey[T int32] interface {
	any
	Get() T
}
