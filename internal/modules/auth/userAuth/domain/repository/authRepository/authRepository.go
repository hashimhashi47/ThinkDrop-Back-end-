package authrepository



// → Interface

type AuthRespository interface {
	Insert(model interface{}) error
}
