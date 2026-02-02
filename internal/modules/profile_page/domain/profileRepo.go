package domain

type ProfileRepo interface {
	Find(model interface{}, query string, args interface{}, preloads ...string) error
}
