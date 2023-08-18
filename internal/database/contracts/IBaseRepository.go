package contracts

type IBaseRepository interface {
	ReadAll(Response interface{}) error
	ReadAllBy(Response interface{}, column string, value interface{}) error
	ReadBy(model interface{}, column string, value interface{}) error
	Create(model interface{}) error
	UpdateBy(model interface{}, UpdateObject interface{}, column string, value interface{}) error
	DeleteBy(model interface{}, column string, value interface{}) error
}
