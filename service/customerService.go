package service

import (
	"customerManager/model"
)

// 该CustomerService, 完成对Customer的操作, 包括增删改查
type CustomerService struct {
	customers   []model.Customer
	customerNum int // 声明一个字段, 表示当前切片含有多少个客户
}

// 初始化
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "小明", "男", 20, "12345678", "xiaoming@qq.com")
	customerService.customers = append(customerService.customers, customer) // 使用切片将数据添加
	return customerService
}

// 显示客户列表, 返回客户切片
func (c *CustomerService) List() []model.Customer {
	return c.customers
}

// 添加客户
func (c *CustomerService) Add(customer model.Customer) bool {
	c.customerNum++
	customer.Id = c.customerNum
	c.customers = append(c.customers, customer)
	return true
}

//根据id删除客户
func (c *CustomerService) Delete(id int) bool {
	index := c.FindById(id)
	//如果index == -1, 说明没有这个客户
	if index == -1 {
		return false
	}
	// 从切片中删除一个元素
	c.customers = append(c.customers[:index], c.customers[index+1:]...)
	return true
}

//根据id查找客户在切片中对应下标,如果没有该客户，返回-1
func (c *CustomerService) FindById(id int) int {
	index := -1
	//遍历this.customers 寻找Id号
	for i := 0; i < len(c.customers); i++ {
		if c.customers[i].Id == id {
			index = i
		}
	}
	return index
}

// 修改客户
func (c *CustomerService) Update(id int, customer model.Customer) bool {
	index := c.FindById(id)
	//如果index == -1, 说明没有这个客户
	if index == -1 {
		return false
	}
	// 修改客户
	c.customers[index-1] = customer
	return true
}
