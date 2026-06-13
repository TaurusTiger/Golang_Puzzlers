package main

import (
	"fmt"
)

type Animal interface {
	// ScientificName 用于获取动物的学名。
	ScientificName() string
	// Category 用于获取动物的基本分类。
	Category() string
}

type Named interface {
	// Name 用于获取名字。
	Name() string
}

type Pet interface {
	Animal
	Named
}

type PetTag struct {
	name  string
	owner string
}

func (pt PetTag) Name() string {
	return pt.name
}

func (pt PetTag) Owner() string {
	return pt.owner
}

type Dog struct {
	PetTag
	scientificName string
}

func (dog Dog) ScientificName() string {
	return dog.scientificName
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	petTag := PetTag{name: "little pig"}

	/*
		## 代码解释

		这是 **Go 语言的类型断言**语法，用于检查 `petTag` 是否实现了 `Named` 接口：

		- `interface{}(petTag)`：将 `petTag` 转换为空接口类型
		- `.(Named)`：尝试断言该值实现了 `Named` 接口
		- `ok`：布尔值，`true` 表示实现了接口，`false` 表示未实现

		这是一种安全的接口实现检查机制，不会在失败时引发 panic。
	*/
	_, ok := interface{}(petTag).(Named)
	fmt.Printf("PetTag implements interface Named: %v\n", ok)
	dog := Dog{
		PetTag:         petTag,
		scientificName: "Labrador Retriever",
	}
	_, ok = interface{}(dog).(Animal)
	fmt.Printf("Dog implements interface Animal: %v\n", ok)
	_, ok = interface{}(dog).(Named)
	fmt.Printf("Dog implements interface Named: %v\n", ok)
	_, ok = interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
}
