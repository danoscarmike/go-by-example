package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Exported1 struct {
	Letters string
	Numbers []int
}

type Exported2 struct {
	Letters string `json:"letters"`
	Numbers []int  `json:"numbers"`
}

type unexported struct {
	Letters string `json:"letters"`
	numbers []int  `json:"numbers"`
}

func main() {
	boolBytes, _ := json.Marshal(true)
	fmt.Printf("%T\n", boolBytes)
	fmt.Println(string(boolBytes))

	intBytes, _ := json.Marshal(12)
	fmt.Println(string(intBytes))

	floatBytes, _ := json.Marshal(1.234)
	fmt.Println(string(floatBytes))

	stringBytes, _ := json.Marshal("golang")
	fmt.Println(string(stringBytes))

	sliceBytes, _ := json.Marshal([]string{"abc", "def", "ghi"})
	fmt.Println(string(sliceBytes))

	mapBytes, _ := json.Marshal(map[string]int{"abc": 123, "def": 456, "ghi": 789})
	fmt.Println(string(mapBytes))

	exp1 := &Exported1{
		Letters: "abc",
		Numbers: []int{123, 456, 789},
	}
	expStructBytes, _ := json.Marshal(exp1)
	fmt.Println(string(expStructBytes))

	exp2 := &Exported2{
		Letters: "jkl",
		Numbers: []int{101, 112, 131},
	}
	expStructBytes2, _ := json.Marshal(exp2)
	fmt.Println(string(expStructBytes2))

	unexp := &unexported{
		Letters: "stv",
		numbers: []int{415, 161, 718},
	}
	unexpStructBytes, _ := json.Marshal(unexp)
	fmt.Println(string(unexpStructBytes))

	jsonBytes := []byte(`{"nums":123, "strs":["abc","def","ghi","jkl"]}`)

	var jsonData map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &jsonData); err != nil {
		panic(err)
	}

	fmt.Printf("%T\n", jsonData)
	fmt.Println(jsonData)

	strs := jsonData["strs"].([]interface{})
	str1 := strs[0]
	fmt.Printf("%T\n", str1)
	str1Conv := str1.(string)
	fmt.Printf("%T\n", str1Conv)
	fmt.Println(str1, str1Conv)

	str := `{"letters":"abc","numbers":[123, 456, 789]}`
	imp1 := Exported2{}
	json.Unmarshal([]byte(str), &imp1)
	fmt.Println(imp1)
	fmt.Println(imp1.Numbers[1])

	enc := json.NewEncoder(os.Stdout)
	data := map[string]int{"abc": 123, "def": 456}
	enc.Encode(data)

}
