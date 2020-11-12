package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// 任何类型 都实现了 interface{}

	// 在编译期不能确定输入数据的类型，
	// 但是该类型肯定实现了interface{}，
	// 所以需要逐层断言，直到最后

	// 不同的数据结构的断言方式不同
	// map:

	// slice:

	// array:

	parseNestMapByJSON()

	// 在编译期就确定了输入数据的类型，所以类型断言时候应该一步到位，直接断言为该类型
	parseNestMapByMap()

	// 所以最主要的就是：编译期能否确定数据类型

}

type srcM struct {
	Sm interface{} `json:"sm"`
}

func parseNestMapByJSON() {

	srcMapStr := `{"sm": {
		"dir1": {
			"dir2": {
				"dir3": {
					"filePath": "main.go"
				}
			}
		}
	}}`

	srcMapRes := srcM{}
	err := json.Unmarshal([]byte(srcMapStr), &srcMapRes)
	if err != nil {
		fmt.Println(err)
		return
	}

	parseNestMap(srcMapRes)
}

func parseNestMapByMap() {

	srcMap := map[string]map[string]map[string]map[string]string{
		"dir1": {
			"dir2": {
				"dir3": {
					"filePath": "main.go",
				},
			},
		},
	}

	parseNestMap(srcMap)
}

func parseNestMap(m interface{}) {

	srcSlice := []string{
		"dir1", "dir2", "dir3", "filePath",
	}

	var (
		res string
	)

	for _, v := range srcSlice {

		m1, ok := m.(map[string]interface{})
		if ok {
			m2 := m1[v]
			m = m2
			fmt.Println("need handle...")
			fmt.Println(m)
		}

		s1, ok := m.(string)
		if ok {
			res = s1
			fmt.Println("end...")
		}

	}

	fmt.Println("result: ", res)
}
