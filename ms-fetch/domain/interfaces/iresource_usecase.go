package interfaces

import "ms-fetch/domain/presenters"

type IResourceUseCase interface {
	ReadResouce() (presenter presenters.ArrayFilterResourcePresenter, err error)

	ReadResouceUsingCurrUSD() (presenter presenters.ArrayFilterResourceCurrUSDPresenter, err error)

	ReportWeekResouce() (presenter presenters.ArrayReportResourcePresenter, err error)
}
