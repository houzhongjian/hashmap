package main

import (
	"log"
	"pandaschool.net/demo/hashmap"
)

func main() {
	mp := hashmap.Make(10)
	mp.Set("name", "张三")
	mp.Set("age", 20)
	log.Println(mp.Get("name").(string))
	log.Println(mp.Get("age").(int))
	mp.Set("name", "李四")
	log.Println(mp.Get("name").(string))
}