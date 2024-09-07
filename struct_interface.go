package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Data is interface for Mydata
type Data interface {
	SetValue(vals map[string]string)
	PrintData()
}

type Mydata struct {
	Name string
	Data []int
}

func (md *Mydata) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	valt := strings.Split(vals["data"], " ")
	vali := []int{}
	for _, i := range valt {
		n, _ := strconv.Atoi(i)
		vali = append(vali, n)
	}
	md.Data = vali
}

func (md *Mydata) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

type Yourdata struct {
	Name string
	Mail string
	Age  int
}

func (md *Yourdata) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	md.Mail = vals["mail"]
	n, _ := strconv.Atoi(vals["age"])
	md.Age = n
}

func (md *Yourdata) PrintData() {
	fmt.Printf("I'm %s (%d) \n", md.Name, md.Age)
	fmt.Printf("mail: %s \n", md.Mail)
}

func main() {
	ob := [2]Data{}
	ob[0] = new(Mydata)
	ob[0].SetValue(map[string]string{
		"name": "Sachiko",
		"data": "55, 66, 77",
	})
	ob[1] = new(Yourdata)
	ob[1].SetValue(map[string]string{
		"name": "Mami",
		"mail": "mami@mume.mo",
		"age":  "34",
	})
	for _, d := range ob {
		d.PrintData()
		fmt.Println()
	}
}
