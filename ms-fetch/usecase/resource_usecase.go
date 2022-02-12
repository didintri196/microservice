package usecase

import (
	"ms-fetch/domain/interfaces"
	"ms-fetch/domain/presenters"
)

type ResourceUseCase struct {
	*Contract
}

func NewResourceUseCase(ucContract *Contract) interfaces.IResourceUseCase {
	return &ResourceUseCase{ucContract}
}

func (f ResourceUseCase) ReadResouce() (presenter presenters.ArrayFilterResourcePresenter, err error) {
	//TODO implement me
	panic("implement me")
}
