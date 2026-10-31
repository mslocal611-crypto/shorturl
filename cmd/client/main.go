package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
)

func main() {
	endpoint := "http://localhost:8080/"
	links := []string{
		"https://example.com",
		"https://yandex.ru",
		"ftp://path/to/file.txt",
		"https://stackoverflow.com",
	}

	client := resty.New()

	for _, link := range links {
		fmt.Println("--------------------------------")
		fmt.Println("Получение короткой ссылки для ", link+" ...")
		response, err := client.R().
			SetHeader("Content-Type", "application/x-www-form-urlencoded").
			SetBody(
				url.Values{"url": {link}}.Encode(),
			).
			Post(endpoint)
		if err != nil {
			fmt.Println("Ошибка: ", err)
			return
		}
		shortLink := response.String()

		fmt.Println("Статус-код ", response.StatusCode())
		fmt.Println("Короткая ссылка: ", shortLink)

		if response.StatusCode() == http.StatusBadRequest {
			fmt.Println()
			continue
		}

		fmt.Println("Переход по ", shortLink+" ...")
		response, err = client.
			SetRedirectPolicy(resty.NoRedirectPolicy()).
			R().
			Get(shortLink)
		if err != nil && !errors.Is(err, resty.ErrAutoRedirectDisabled) {
			fmt.Println("Ошибка: ", err)
			return
		}
		client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(10))

		fmt.Println("Статус-код ", response.StatusCode())
		fmt.Println("Редирект: ", response.Header().Get("Location"))
		fmt.Println("Тело ответа: ", response.String())

		fmt.Println("--------------------------------")
		fmt.Println()
	}
}
