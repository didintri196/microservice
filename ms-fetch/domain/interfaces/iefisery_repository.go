package interfaces

import "ms-fetch/domain/models"

type IEfiseryRepository interface {
	GetResource() ([]models.EfiseryResource, error)
}
