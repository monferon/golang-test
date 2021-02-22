package main

import (
	"fmt"
)

const half = 0.5



//func NewEmployee(name, sex string, age, salary int) employee {
//	return employee{
//		name:   name,
//		age:    age,
//		salary: salary,
//	}
//}

//func (e employee) getInfo() string {
//	return fmt.Sprintf("Сотрудник: %s\nВозраст: %d\nЗарплата: %d\n", e.name, e.age, e.salary)
//}
//
//func (e employee) getYearBirthday() string {
//	return fmt.Sprintf("Код рождения сотрудника - %d\n", 2021-e.age)
//}
//
//func (e employee) setName(name string) {
//	e.name = name
//}
func main() {

	//s := scheduler.NewScheduler()
	//
	//s.Add(context.Background(), func(ctx context.Context){
	//	fmt.Printf("Текущее время: %s\n", time.Now())
	//}, time.Second * 1)


	ms := newMemoryStorage() //указатель на структуру
	ds := newDumbStorage() //указатель на структуру

	spawnEmployees(ms)
	fmt.Println(ms.get(3))

	spawnEmployees(ds)

	//employee1 := NewEmployee("t1", "M", 14, 20000)
	//employee2 := NewEmployee("t2", "M", 20, 30000)
	//employee3 := NewEmployee("t3", "M", 25, 59999)
	//
	//
	//employee2.setName("tratatata" )
	//fmt.Printf("%s\n", employee1.getInfo())
	//fmt.Printf("%s\n", employee2.getInfo())
	//fmt.Printf("%s\n", employee3.getInfo())
	//fmt.Printf("%s\n", employee2.getYearBirthday())

	//var products = [3]string{
	//	"1",
	//	"2",
	//	"3",
	//}
	//for index, value := range products {
	//	fmt.Printf("%d - %s\n", index, value)
	//}
	//for i := 0; i <= 10; i++ {
	//	fmt.Println(i)
	//}
	//fmt.Println(products, len(products))

	//var ages map[string]int
	//
	//fmt.Println("ages nil ->", ages == nil)
	//ages = make(map[string]int)
	//employee := struct {
	//	name   string
	//	sex    string
	//	age    int
	//	salary int
	//}{
	//	name:   "Vasya",
	//	sex:    "M",
	//	age:    12,
	//	salary: 20000,
	//}
	//
	//fmt.Printf("%+v\n", employee)

}

func spawnEmployees(s storage.Storage){
	for i := 1; i <=10; i++ {
		s.Insert(storage.Employee{Id:i})
	}
}
