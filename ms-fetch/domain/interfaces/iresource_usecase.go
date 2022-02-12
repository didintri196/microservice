package interfaces

import "ms-fetch/domain/presenters"

type IResourceUseCase interface {
	ReadResouce() (presenter presenters.ArrayFilterResourcePresenter, err error)
}
