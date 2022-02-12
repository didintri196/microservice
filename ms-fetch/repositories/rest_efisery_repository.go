package repositories

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"ms-fetch/domain/interfaces"
	"ms-fetch/domain/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type RestEfiseryRepository struct {
	HostEfisery string
}

func NewRestEfiseryRepo(hostEfisery string) interfaces.IRestEfiseryRepository {
	return &RestEfiseryRepository{
		HostEfisery: hostEfisery,
	}
}

func (repo RestEfiseryRepository) GetResource() (model []models.EfiseryResource, err error) {

	url := repo.HostEfisery + "/v1/storages/5e1edf521073e315924ceab4/list"

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
