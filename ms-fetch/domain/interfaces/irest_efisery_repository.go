package interfaces

import "ms-fetch/domain/models"

type IRestEfiseryRepository interface {
	GetResource() ([]models.EfiseryResource, error)
}
