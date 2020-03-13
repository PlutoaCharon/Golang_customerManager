## 客户关系管理系统

> 基于Golang的简单文本界面客户关系管理系统


`view.customerView.go`
- 显示界面
- 接收用户的输入
- 根据用户的输入，调用`customerService`的方法完成客户的管理
```
list 去调用 customerService 的List方法，并显示客户列表
add 方法去调用 customerService 的Add方法, 完成客户添加
delete 方法 调用 customerService 的Delete方法, 完成客户删除
update 方法调用 customerService 的Update方法, 完成客户修改
```

`service.customerService`
- 完成对用户的各种操作
- 对客户的增，删除，修改，显示
- 会声明一个customer的切片
```
List 返回客户信息
NewCustomerService 返回一个customerService实例
Add 将新的客户加入到 customers 
Delete(id int)  删除一个客户
FindById(id int)  返回id号对应的切片的下标
Update(id int, customer model.Customer) 修改客户
```

`model.customer` 
- 表示一个客户
- 客户各种字段
```
customer 表示一个客户信息

type  Customer struct {
Id int
Name string
Gender string...
}

GetInfo 显示该用户的信息
NewCustomer2() 返回客户的方法
```
