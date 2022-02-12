package presenters

import (
	"ms-fetch/domain/models"
	"ms-fetch/utils"
)

type ArrayFilterResourcePresenter struct {
	FilterResourcePresenter []FilterResourcePresenter
}

type FilterResourcePresenter struct {
	UUID         string `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	TglParsed    string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}

func NewArrayFilterResourcePresenter() ArrayFilterResourcePresenter {
	return ArrayFilterResourcePresenter{}
}

func (presenter ArrayFilterResourcePresenter) Build(model []models.EfiseryResource) (list ArrayFilterResourcePresenter) {
	for _, row := range model {
		var err error
		var convTgl string

		//percobaan parse tahap 2 tgl valid
		convTgl, err = utils.ParseRFC3339Nano(row.TglParsed)
		//percobaan parse tahap 2 tgl valid
		if err != nil {
			convTgl, err = utils.ParseISO(row.TglParsed)
		}

		if err == nil {
			list.FilterResourcePresenter = append(list.FilterResourcePresenter, FilterResourcePresenter{
				UUID:         row.UUID,
				Komoditas:    row.Komoditas,
				AreaProvinsi: row.AreaProvinsi,
				AreaKota:     row.AreaKota,
				Size:         row.Size,
				Price:        row.Price,
				TglParsed:    convTgl,
				Timestamp:    row.Timestamp,
			})
		}
	}
	return
}
