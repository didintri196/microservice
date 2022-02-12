package presenters

import (
	"fmt"
	"ms-fetch/domain/models"
	"ms-fetch/utils"
	"strconv"

	"github.com/leekchan/accounting"
)

type ArrayFilterResourceCurrUSDPresenter struct {
	FilterResourceCurrUSDPresenter []FilterResourceCurrUSDPresenter
}

type FilterResourceCurrUSDPresenter struct {
	UUID         string `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	PriceUSD     string `json:"price_usd"`
	TglParsed    string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}

func NewArrayFilterResourceCurrUSDPresenter() ArrayFilterResourceCurrUSDPresenter {
	return ArrayFilterResourceCurrUSDPresenter{}
}

func (presenter ArrayFilterResourceCurrUSDPresenter) Build(currUSD models.Currency, model []models.EfiseryResource) (list ArrayFilterResourceCurrUSDPresenter) {
	acUSD := accounting.Accounting{Symbol: "$", Precision: 2}
	for _, row := range model {
		var err error
		var convTgl string

		//percobaan parse tahap 2 tgl valid
		convTgl, err = utils.ParseRFC3339Nano(row.TglParsed)
		//percobaan parse tahap 2 tgl valid
		if err != nil {
			convTgl, err = utils.ParseISO(row.TglParsed)
		}

		currIDR, _ := strconv.Atoi(row.Price)
		currUSD, _ := strconv.ParseFloat(fmt.Sprintf("%.6f", currUSD.IdrUsd), 64)
		convUSD := currUSD * float64(currIDR)

		if err == nil {
			list.FilterResourceCurrUSDPresenter = append(list.FilterResourceCurrUSDPresenter, FilterResourceCurrUSDPresenter{
				UUID:         row.UUID,
				Komoditas:    row.Komoditas,
				AreaProvinsi: row.AreaProvinsi,
				AreaKota:     row.AreaKota,
				Size:         row.Size,
				Price:        row.Price,
				TglParsed:    convTgl,
				Timestamp:    row.Timestamp,
				PriceUSD:     acUSD.FormatMoney(convUSD),
			})
		}
	}
	return
}
