package objrepository

import "github.com/FurmanovD/codesamples/interfaces-with-providers/pkg/infra/db/model"

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetObj(id int) (*model.Object, error) {
	return nil, nil
}

func (r *repository) StoreObj(*model.Object) error {
	return nil
}

func (r *repository) UpdateObj(*model.Object) error {
	return nil
}

func (r *repository) ListObj(conditions interface{}) ([]*model.Object, error) {
	return nil, nil
}

func (r *repository) DeleteObj(*model.Object) error {
	return nil
}
