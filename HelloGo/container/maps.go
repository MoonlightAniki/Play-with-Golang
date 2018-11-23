package main

import "fmt"

func main() {

	fmt.Println("================================ Creating map ================================")

	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "immoc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 = make(map[string]int) // m3 == nil

	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m3)

	fmt.Println("================================ Traversing map ================================")
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k, _ := range m {
		fmt.Println(k)
	}
	for _, v := range m {
		fmt.Println(v)
	}

	fmt.Println("================================ Geting values ================================")
	name := m["name"]
	course := m["course"]
	site := m["site"]
	quality := m["quality"]
	fmt.Printf("name:%s, course:%s, site:%s, quality:%s\n", name, course, site, quality)

	fmt.Printf("cause:%s\n", m["cause"])
	fmt.Println(m["cause"] == "")

	cause, ok := m["cause"]
	fmt.Println(cause, ok)

	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	if name, ok := m["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("================================ Deleting value ================================")

	name, ok = m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
