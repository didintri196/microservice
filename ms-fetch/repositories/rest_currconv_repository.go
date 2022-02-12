package repositories

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ms-fetch/domain/interfaces"
	"ms-fetch/domain/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type RestCurrConvRepository struct {
	HostCurrConv   string
	ApikeyCurrConv string
}

func NewRestCurrConvRepo(hostCurrConv string, apikeyCurrConv string) interfaces.IRestCurrConvRepository {
	return &RestCurrConvRepository{
		HostCurrConv:   hostCurrConv,
		ApikeyCurrConv: apikeyCurrConv,
	}
}

func (repo RestCurrConvRepository) GetIDRtoUSD() (model models.Currency, err error) {

	url := fmt.Sprintf("%s/api/v7/convert?q=IDR_USD&compact=ultra&apiKey=%s", repo.HostCurrConv, repo.ApikeyCurrConv)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return model, err
	}

	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return model, err
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(respBody, &model)
	if err != nil {
		return model, err
	}

	return model, err
}
