package main

import (
	"customerManager/model"
	"customerManager/service"
	"fmt"
)

type customerView struct {
	// 定义必要字段
	key             string // 输入字符
	loop            bool   // 表示是否循环显示菜单
	customerService *service.CustomerService
}

// 显示所有的客户信息
func (cv *customerView) list() {
	customers := cv.customerService.List()

	// 显示
	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Printf("编号\t姓名\t性别\t年龄\t电话\t邮箱\n")
	for i := 0; i < len(customers); i++ {
		fmt.Printf(customers[i].GetInfo())
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
}

// 修改客户
func (c *customerView) update() {
	fmt.Println("请输入待修改的客户编号(-1)退出:")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}
	customer := c.getUserInput()

	if c.customerService.Update(id, customer) {
		fmt.Println("---------------------修改完成---------------------")
	} else {
		fmt.Println("---------------------修改失败，输入的id号不存在----")
	}

}

// 添加用户
func (c *customerView) add() {
	customer := c.getUserInput()
	// id是唯一的，需要系统分配
	//调用
	if c.customerService.Add(customer) {
		fmt.Println("---------------------添加完成---------------------")
	} else {
		fmt.Println("---------------------添加失败---------------------")
	}
}

// 删除用户
func (c *customerView) delete() {
	fmt.Println("---------------------删除客户---------------------")
	fmt.Println("请选择待删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}
	fmt.Println("确认是否删除(Y/N)：")
	//这里同学们可以加入一个循环判断，直到用户输入 y 或者 n,才退出..
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		//调用customerService 的 Delete方法
		if c.customerService.Delete(id) {
			fmt.Println("---------------------删除完成---------------------")
		} else {
			fmt.Println("---------------------删除失败，输入的id号不存在----")
		}
	}
}

// 退出系统
func (c *customerView) exit() {
	fmt.Println("确认是否退出(Y/N):")
	for {
		fmt.Scanln(&c.key)
		if c.key == "Y" || c.key == "N" || c.key == "y" || c.key == "n" {
			break
		}

		fmt.Println("输入有误, 确认是否退出(Y/N)")
	}

	if c.key == "Y" || c.key == "y" {
		c.loop = false
	}
}

// 填入信息
func (c *customerView) getUserInput() model.Customer {
	fmt.Printf("姓名：")
	name := ""
	_, _ = fmt.Scanln(&name)
	fmt.Printf("性别：")
	gender := ""
	_, _ = fmt.Scanln(&gender)
	fmt.Printf("年龄：")
	age := 0
	_, _ = fmt.Scanln(&age)
	fmt.Printf("电话：")
	phone := ""
	_, _ = fmt.Scanln(&phone)
	fmt.Printf("邮箱：")
	email := ""
	_, _ = fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)
	return customer
}

// 显示主菜单
func (cv *customerView) mainMenu() {
	for {
		fmt.Println("----------------------客户信息管理软件----------------------")
		fmt.Println("                  1 添加客户")
		fmt.Println("                  2 修改客户")
		fmt.Println("                  3 删除客户")
		fmt.Println("                  4 客户列表")
		fmt.Println("                  5 退   出")
		fmt.Print("请选择(1-5)：")

		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("你的输入有误,请重新输入")
		}

		if !cv.loop {
			break
		}
	}

	fmt.Println("退出客户关系管理系统")
}
func main() {
	// 显示主菜单
	customerView := customerView{
		key:  "",
		loop: true,
	}
	// 对customerView结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	// 显示主菜单
	customerView.mainMenu()

}
