package orm

type DataType interface {
	Value() interface{}
	IsNull() bool
}
