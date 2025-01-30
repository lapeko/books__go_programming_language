package main

import (
	"fmt"
	"reflect"
)

type Player struct {
	Awards []string `json:"wrds_for_example"`
	Name   string
}

func (p *Player) GetName() string {
	return p.Name
}

func (p Player) SetName(name string) {
	p.Name = name
}

func main() {
	var p Player
	v := reflect.ValueOf(&p)
	t := v.Type()
	for i := 0; i < v.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
	}

	fmt.Println("// also fine")
	var p2 Player
	t2 := reflect.TypeOf(&p2)
	for i := 0; i < t2.NumMethod(); i++ {
		fmt.Println(t2.Method(i).Name)
	}

	vAwards := v.Elem().FieldByName("Awards")
	if vAwards.CanSet() {
		vAwards.Set(reflect.Append(vAwards, reflect.ValueOf("qwe")))
	} else {
		fmt.Println("awards field is not addressable")
	}
	tAwards, _ := t.Elem().FieldByName("Awards")
	fmt.Println(tAwards.Tag.Get("json"))

}
