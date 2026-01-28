package authrepository

// → Interface
type AuthRespository interface {
	Insert(model interface{}) error
	FindAnything(model interface{}, Query, Any string) error
}
