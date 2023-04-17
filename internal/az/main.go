package az

type ICloud interface {
	GetList() (string, error)
	GetItem(string) (string, error)
}
