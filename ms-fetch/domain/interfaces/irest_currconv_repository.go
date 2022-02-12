package interfaces

import "ms-fetch/domain/models"

type IRestCurrConvRepository interface {
	GetIDRtoUSD() (models.Currency, error)
}
