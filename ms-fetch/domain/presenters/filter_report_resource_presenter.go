package presenters

import (
	"fmt"
	"ms-fetch/domain/models"
	"ms-fetch/utils"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
)

type ArrayReportResourcePresenter struct {
	ReportResourcePresenter []ReportResourcePresenter
}

type ReportResourcePresenter struct {
	AreaProvinsi   string    `json:"area_provinsi"`
	Tahun          int       `json:"tahun"`
	Bulan          int       `json:"bulan"`
	MingguKe       int       `json:"minggu_ke"`
	SizeStatistic  Statistic `json:"size_statistic"`
	PriceStatistic Statistic `json:"price_statistic"`
}

type Statistic struct {
	Max    float64 `json:"max"`
	Min    float64 `json:"min"`
	Median float64 `json:"median"`
	Avg    float64 `json:"avg"`
}

func NewArrayReportResourcePresenter() ArrayReportResourcePresenter {
	return ArrayReportResourcePresenter{}
}

func (presenter ArrayReportResourcePresenter) Build(model []models.EfiseryResource) (hasil ArrayReportResourcePresenter) {
	var list ArrayFilterResourcePresenter
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

	var tmp = map[string][]FilterResourcePresenter{}

	for _, rowResource := range list.FilterResourcePresenter {
		tglConv := utils.ParseDatetoTime(rowResource.TglParsed)
		formatTgl := fmt.Sprintf("%d-%d-%d", tglConv.Year(), int(tglConv.Month()), int(tglConv.Weekday()))
		keyGroup := fmt.Sprintf("%s_%s", rowResource.AreaProvinsi, formatTgl)
		tmp[keyGroup] = append(tmp[keyGroup], rowResource)
	}

	for k := range tmp {
		splitKey := strings.Split(k, "_")
		splitDate := strings.Split(splitKey[1], "-")
		tahun, _ := strconv.Atoi(splitDate[0])
		bulan, _ := strconv.Atoi(splitDate[1])
		mingguKe, _ := strconv.Atoi(splitDate[2])

		sizeList := []float64{}
		priceList := []float64{}
		for _, rowData := range tmp[k] {
			size, _ := strconv.ParseFloat(rowData.Size, 64)
			sizeList = append(sizeList, size)
			price, _ := strconv.ParseFloat(rowData.Price, 64)
			priceList = append(priceList, price)
		}

		valueMedianSize, _ := stats.Median(sizeList)
		valueMeanSize, _ := stats.Mean(sizeList)
		valueMaxSize, _ := stats.Max(sizeList)
		valueMinSize, _ := stats.Min(sizeList)

		valueMedianPrice, _ := stats.Median(priceList)
		valueMeanPrice, _ := stats.Mean(priceList)
		valueMaxPrice, _ := stats.Max(priceList)
		valueMinPrice, _ := stats.Min(priceList)

		hasil.ReportResourcePresenter = append(hasil.ReportResourcePresenter, ReportResourcePresenter{
			Tahun:        tahun,
			Bulan:        bulan,
			MingguKe:     mingguKe,
			AreaProvinsi: splitKey[0],
			SizeStatistic: Statistic{
				Max:    valueMaxSize,
				Min:    valueMinSize,
				Median: valueMedianSize,
				Avg:    valueMeanSize,
			},
			PriceStatistic: Statistic{
				Max:    valueMaxPrice,
				Min:    valueMinPrice,
				Median: valueMedianPrice,
				Avg:    valueMeanPrice,
			},
		})
	}
	return hasil
}
