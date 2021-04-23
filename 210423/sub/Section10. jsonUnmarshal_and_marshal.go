package sub

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string   `json:"name"`      // 가리고 싶을 땐 "-"로 표기
	Age       int      `json:"age"`       // int를 string로 표기하고 싶을 땐 ,와 함께 string을 정의
	Nicknames []string `json:"nicknames"` // omitempty로 해당 값이 없다면 표기하지 않을 수도 있다
}

func (p Person) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + p.Name,
	})
	return v, err
}

func JsonExample() {
	b := []byte(`{"name":"Mike", "age":20, "nicknames":["a","b","c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}

	fmt.Println(p.Name, p.Age, p.Nicknames) // Mike 20 [a b c]

	v, _ := json.Marshal(p)
	fmt.Println(string(v))
	// before `json`setting : {"Name":"Mike","Age":20,"Nicknames":["a","b","c"]}
	// after : {"name":"Mike","age":20,"nicknames":["a","b","c"]}
	// even after custom Marshal func => {"Name":"Mr.Mike"}
}
