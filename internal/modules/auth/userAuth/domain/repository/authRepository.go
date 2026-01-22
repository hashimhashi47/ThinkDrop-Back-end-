package repository



// → Interface

type AuthRespository interface {
	Insert(model interface{}) error
}
