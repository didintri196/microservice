package usecase

import (
	"ms-fetch/domain/interfaces"
	"ms-fetch/domain/presenters"
	"ms-fetch/repositories"
)

type ResourceUseCase struct {
	*Contract
}

func NewResourceUseCase(ucContract *Contract) interfaces.IResourceUseCase {
	return &ResourceUseCase{ucContract}
}

func (uc ResourceUseCase) ReadResouce() (presenter presenters.ArrayFilterResourcePresenter, err error) {
	repo := repositories.NewEfiseryRepo(uc.Host.Efisery)
	data, err := repo.GetResource()
	if err != nil {
		return presenter, err
	}

	presenter = presenters.NewArrayFilterResourcePresenter().Build(data)

	return presenter, err
}
