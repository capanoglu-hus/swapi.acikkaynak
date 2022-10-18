package goapi

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func Client() {
	client := resty.New() // yeni bir resty client oluşturuldu
	_ = client // boş atama yapıldı
	resp, err := client.R(). // get ile sorgu atıldı resp ile gelen sonuclar alındı err hata mesajları için
		EnableTrace().
		Get("https://swapi.dev/api/people/1/")

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err) //hata mesajları için
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time()) // atılan sorguda gecen zaman
	fmt.Println("  Received At:", resp.ReceivedAt()) // sorgunun alındıgı tarih
	fmt.Println("  Body       :\n", resp)  //istenilen kısım
	fmt.Println()
}
