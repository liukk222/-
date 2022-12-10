package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string
	Age   int
	Phone string
}
type ContactBook struct {
	Contacts []Contact
}

func (cb *ContactBook) AddContact(name string, age int, phone string) {
	// 创建新的联系人
	c := Contact{
		Name:  name,
		Age:   age,
		Phone: phone,
	}

	// 将新的联系人添加到通讯录中
	cb.Contacts = append(cb.Contacts, c)
}

func (cb *ContactBook) FindContact(name string) *Contact {
	// 遍历通讯录中的所有联系人
	for _, c := range cb.Contacts {
		if c.Name == name {
			// 返回找到的联系人
			return &c

		} // 未找到联系人

	}
	return nil
}
func (cb *ContactBook) DeleteContact(name string) {
	// 遍历通讯录中的所有联系人
	for i, c := range cb.Contacts {
		if c.Name == name {
			// 删除找到的联系人
			cb.Contacts = append(cb.Contacts[:i], cb.Contacts[i+1:]...)
			return
		}
	}
}

func (cb *ContactBook) SaveToFile(filePath string) error {
	// 打开文件
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	// 创建编码器
	encoder := json.NewEncoder(file)

	// 将通讯录的内容编码为 JSON 并写入文件
	return encoder.Encode(cb)
}
func (cb *ContactBook) LoadFromFile(filePath string) error {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	// 将文件中的内容解码为通讯录
	return decoder.Decode(cb)
}
func (cb *ContactBook) PrintContacts() {
	// 遍历通讯录中的所有联系人
	for _, c := range cb.Contacts {
		// 打印联系人的信息
		fmt.Printf("Name: %s\nAge: %d\nPhone: %s\n\n", c.Name, c.Age, c.Phone)
	}
}

func main() {
	cb := ContactBook{}
	// 添加联系人
	cb.AddContact("Tom", 25, "123-456-7890")
	cb.AddContact("Jerry", 30, "234-567-8901")
	cb.AddContact("Bob", 35, "345-678-9012")

	// 查找联系人
	c := cb.FindContact("Jerry")
	if c != nil {
		fmt.Printf("Found contact: %+v\n", c)
	} else {
		fmt.Println("Contact not found.")
	}

	// 删除联系人
	cb.DeleteContact("Tom")

	// 保存通讯录到文件
	cb.SaveToFile("test.txt")

	// 重新创建通讯录
	cb = ContactBook{}

	// 从文件中读取通讯录
	cb.LoadFromFile("test.txt")

	// 打印通讯录中的所有联系人
	cb.PrintContacts()

}
