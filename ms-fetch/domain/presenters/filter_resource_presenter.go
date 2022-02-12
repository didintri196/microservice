package presenters

import (
	"ms-fetch/domain/models"
	"time"
)

type ArrayFilterResourcePresenter struct {
	FilterResourcePresenter []FilterResourcePresenter
}

type FilterResourcePresenter struct {
	UUID         string    `json:"uuid"`
	Komoditas    string    `json:"komoditas"`
	AreaProvinsi string    `json:"area_provinsi"`
	AreaKota     string    `json:"area_kota"`
	Size         string    `json:"size"`
	Price        string    `json:"price"`
	TglParsed    time.Time `json:"tgl_parsed"`
	Timestamp    string    `json:"timestamp"`
}

func NewArrayFilterResourcePresenter() ArrayFilterResourcePresenter {
	return ArrayFilterResourcePresenter{}
}

func (presenter ArrayFilterResourcePresenter) Build(model []models.EfiseryResource) (list ArrayFilterResourcePresenter) {
	for _, row := range model {
		list.FilterResourcePresenter = append(list.FilterResourcePresenter, FilterResourcePresenter{
			UUID:         row.UUID,
			Komoditas:    row.Komoditas,
			AreaProvinsi: row.AreaProvinsi,
			AreaKota:     row.AreaKota,
			Size:         row.Size,
			Price:        row.Price,
			TglParsed:    row.TglParsed,
			Timestamp:    row.Timestamp,
		})
	}
	return
}
