package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	brasilCh := make(chan string)
	viacepCh := make(chan string)

	go func() {
		getBrasilAPI(brasilCh)
	}()
	go func() {
		getViaCepAPI(viacepCh)
	}()

	select {
	case viaCepResp := <-viacepCh:
		fmt.Printf("Via CEP response: %s", viaCepResp)
	case brasilAPIResp := <-brasilCh:
		fmt.Printf("Brasil API response: %s", brasilAPIResp)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout.")
	}
}

func getBrasilAPI(brasilCh chan<- string) {
	brasilAPI, err := http.Get("https://brasilapi.com.br/api/cep/v1/59600-786")
	if err != nil {
		panic(err)
	}
	defer brasilAPI.Body.Close()
	body, err := io.ReadAll(brasilAPI.Body)
	if err != nil {
		panic(err)
	}
	brasilCh <- string(body)
}

func getViaCepAPI(viacepCh chan<- string) {
	viaCepAPI, err := http.Get("https://viacep.com.br/ws/59600-786/json/")
	if err != nil {
		panic(err)
	}
	defer viaCepAPI.Body.Close()
	body, err := io.ReadAll(viaCepAPI.Body)
	if err != nil {
		panic(err)
	}
	viacepCh <- string(body)
}
