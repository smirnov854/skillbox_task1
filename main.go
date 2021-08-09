package main

import "fmt"

type Category struct {
	id int
	name int
 	path string
}

func (cat Category) setName(name int){
	cat.name = name
}

func (cat Category) getName() int{
	return cat.name
}

type Product struct {
	id int
	categoryId int
	name string
}


func main(){
	firstCategory:= Category{5,11,"tetst"}
	secondCategory:= Category{55,33,"testcat"}
	firstProduct := Product{id: 1,categoryId: 5,name:"first"}
	secondProduct:= Product{id:2, categoryId: 11, name:"second"}

	fmt.Println(firstCategory.path)
	fmt.Println(secondCategory.id)

	fmt.Println(firstProduct)
	fmt.Println(secondProduct)


}
