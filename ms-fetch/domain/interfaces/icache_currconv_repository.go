package interfaces

import "ms-fetch/domain/models"

type ICacheCurrConvRepository interface {
	GetIDRtoUSD() (models.Currency, error)
	SetIDRtoUSD(models.Currency)
}
