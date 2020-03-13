package model

import "fmt"

// 声明一个Customer结构体, 表示一个客户信息
type Customer struct {
	Id          int
	Name        string
	Gender      string
	Age         int
	PhoneNumber string
	Email       string
}

// 编写一个工厂模式, 返回一个Customer的实例
func NewCustomer(id int, name string, gender string, age int, phonenumber string, email string) Customer {
	return Customer{
		Id:          id,
		Name:        name,
		Gender:      gender,
		Age:         age,
		PhoneNumber: phonenumber,
		Email:       email,
	}
}

// 新建不带Id的实例方法
func NewCustomer2(name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Name:        name,
		Gender:      gender,
		Age:         age,
		PhoneNumber: phone,
		Email:       email,
	}
}

// 返回用户的信息
func (customer Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\n",
		customer.Id, customer.Name, customer.Gender, customer.Age, customer.PhoneNumber, customer.Email)
	return info
}
