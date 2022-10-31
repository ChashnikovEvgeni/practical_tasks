package main

import (
	"encoding/json"
	"fmt"
)

// Реализовать паттерн «адаптер» на любом примере.

// Интерфейс который нужно реализовать
type DataSaver interface {
	GetByteData() []byte
}

// структура с данными пользователя
type PersonData struct {
	Id      uint64
	Name    string
	Age     uint8
	Balance float64
}

// получение json из данной структуры
func (p *PersonData) GetByteData() []byte {
	byteData, _ := json.Marshal(p)

	return byteData
}

// Условное хранилище данных
type DataStorage struct {
	Data []byte
}

// метод сохранения данных в хранилище
func (d *DataStorage) SaveData(p DataSaver) {
	d.Data = p.GetByteData()
}

// метод для просмотра данных, находящихся в хранилище
func (d *DataStorage) ShowData() {
	var pd PersonData
	err := json.Unmarshal(d.Data, &pd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pd)
}

// структура данных пользователя в ином сервисе
// Данна структура не реализует методы интерфейса DataSaver
// Значит данную структуру нельзя передать в метод  хранилиза SaveData(p DataSaver)
type Account struct {
	Id        uint64
	FirstName string
	LastName  string
	Age       uint8
	Address   string
	region    string
	Balance   float64
}

// получение данных из json файла
func (a *Account) GetAccountData(jsonb []byte) {
	err := json.Unmarshal(jsonb, &a)
	if err != nil {
		fmt.Println(err)
	}
}

// Адапетр реализующий интерфейс DataSaver, а значит способный взаимодействовать с хранилищем
// SaveData(p DataSaver) принимает объект вида, который реализиует интерфейс DataSaver
type DataAdapter struct {
	*Account
}

// реализация метода GetByteData() интерфейса DataSaver
func (d *DataAdapter) GetByteData() []byte {
	var pd PersonData
	pd.Id = d.Id
	pd.Name = d.FirstName
	pd.Age = d.Age
	pd.Balance = d.Balance
	byteData, _ := json.Marshal(pd)
	return byteData

}

func Task_21_solution() {
	// инициализация структур
	datastrorage := &DataStorage{}

	account := &Account{Id: 1, FirstName: "Ivan", LastName: "Ivanov", Age: 19, Address: "Moscow", region: "Russia", Balance: 700.25}

	dataAdapter := &DataAdapter{Account: account}

	// Вызов методов хранилища
	datastrorage.SaveData(dataAdapter)
	datastrorage.ShowData()
	/*
			output:
		{1 Ivan 19 700.25}

	*/

}
