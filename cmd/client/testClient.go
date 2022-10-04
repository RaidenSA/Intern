package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {
	// адрес сервиса (как его писать, расскажем в следующем уроке)
	endpoint := "http://localhost:8080/"
	//fmt.Println("Post = 0, get = 1")
	//a:=0
	//fmt.Scan(a)
	//if a ==0{
	// контейнер данных для запроса
	data := url.Values{}
	// приглашение в консоли
	fmt.Println("Введите длинный URL")
	// открываем потоковое чтение из консоли
	reader := bufio.NewReader(os.Stdin)
	// читаем строку из консоли
	long, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	long = strings.TrimSuffix(long, "\n")
	// заполняем контейнер данными
	data.Set("url", long)
	// конструируем HTTP-клиент
	client := &http.Client{}
	// конструируем запрос
	// запрос методом POST должен, кроме заголовков, содержать тело
	// тело должно быть источником потокового чтения io.Reader
	// в большинстве случаев отлично подходит bytes.Buffer
	request, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBufferString(long))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// в заголовках запроса сообщаем, что данные кодированы стандартной URL-схемой
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(long)))
	// отправляем запрос и получаем ответ
	//response, err := client.Do(request)
	response, err := client.Get("http://localhost:8080/3rI4hVUrBH")
	//response, err := client.Get("http://localhost:8080/G0iRMQX4gs")
	//http://ya.ru
	//http://ppkvmeyfa.biz
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// печатаем код ответа
	fmt.Println("Статус-код ", response.Status)
	loc, _ := response.Location()
	fmt.Println("location", loc)
	defer response.Body.Close()
	// читаем поток из тела ответа
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// и печатаем его
	fmt.Println(string(body))
}
