package usecase

import (
	"fmt"
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
	repo := repositories.NewRestEfiseryRepo(uc.Host.Efisery)
	data, err := repo.GetResource()
	if err != nil {
		return presenter, err
	}

	presenter = presenters.NewArrayFilterResourcePresenter().Build(data)
	return presenter, err
}

func (uc ResourceUseCase) ReadResouceUsingCurrUSD() (presenter presenters.ArrayFilterResourceCurrUSDPresenter, err error) {
	repoEfisery := repositories.NewRestEfiseryRepo(uc.Host.Efisery)
	repoCurrConv := repositories.NewRestCurrConvRepo(uc.Host.CurrConv, uc.Apikey.CurrConv)
	repoCacheCurrConv := repositories.NewCacheCurrConvRepo(uc.Cache)

	data, err := repoEfisery.GetResource()
	if err != nil {
		return presenter, err
	}

	//cek cache
	dataUSD, err := repoCacheCurrConv.GetIDRtoUSD()
	if err != nil {
		fmt.Println("Gagal ambil data dari cache")
		// jika cache kosong maka akan rest
		dataUSD, err = repoCurrConv.GetIDRtoUSD()
		if err != nil {
			return presenter, err
		}

		//save ke cache setelah mendapatkan result
		repoCacheCurrConv.SetIDRtoUSD(dataUSD)
	} else {
		fmt.Println("Berhasil ambil dari cache")
	}

	presenter = presenters.NewArrayFilterResourceCurrUSDPresenter().Build(dataUSD, data)
	return presenter, err
}

func (uc ResourceUseCase) ReportWeekResouce() (presenter presenters.ArrayReportResourcePresenter, err error) {
	repo := repositories.NewRestEfiseryRepo(uc.Host.Efisery)
	data, err := repo.GetResource()
	if err != nil {
		return presenter, err
	}

	presenter = presenters.NewArrayReportResourcePresenter().Build(data)
	return presenter, err
}
