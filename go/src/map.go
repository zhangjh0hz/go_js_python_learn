package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m, m2, m3)

	//遍历
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("length = ", len(m))

	fmt.Println("Getting values")
	if cousrName, ok := m["course"]; ok {
		fmt.Println(cousrName, ok)
	} else {
		fmt.Println("Key does not exist")
	}

	if causeName, ok := m["caurse"]; ok {
		fmt.Println(causeName, ok)
	} else {
		fmt.Println("Key does not exist")
	}

	fmt.Println("deleting key")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}

//除了slice，map,function的内建类型，其他都可以作为key,
//map使用哈希表，必须可以比较相等
