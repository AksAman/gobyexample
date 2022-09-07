package jsonExamples

import (
	"encoding/json"
	"fmt"
	"os"
)

type FruitsResponse struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func Run() {
	runStructExamples()
	runArbitraryDecodingExamples()
	runStreamingJSONExample()
}

func runStreamingJSONExample() {

	encoder := json.NewEncoder(os.Stdout)

	data := map[string]int{"apples": 5, "mangoes": 2}
	encoder.Encode(data)
}

func runStructExamples() {
	fruits := []FruitsResponse{}

	fruits = append(fruits, FruitsResponse{
		Page: 1,
		Fruits: []string{
			"apple",
			"oranges",
			"pears",
		},
	})

	fruits = append(fruits, FruitsResponse{
		Page: 2,
		Fruits: []string{
			"peaches",
			"mangos",
		},
	})

	// serialization or marshalling
	serializedData, err := json.MarshalIndent(fruits, "", "    ")

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("serializedData: %v\n", string(serializedData))

	// de-serialization or unmarshalling
	serializedBytes := []byte(`[{"page":1,"fruits":["apple","oranges","pears"]},{"page":2,"fruits":["peaches","mangos"]}]`)
	var deserializedFruits []FruitsResponse
	_ = json.Unmarshal(serializedBytes, &deserializedFruits)
	fmt.Printf("deserializedFruits: %#v\n", deserializedFruits)
}

func runArbitraryDecodingExamples() {
	fmt.Println("\n----- arbitary")
	serializedBytes := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var data map[string]interface{}

	_ = json.Unmarshal(serializedBytes, &data)
	fmt.Printf("data: %v\n", data)

	num := data["num"].(float64)
	fmt.Printf("num: %v, type: %T\n", num, num)

}
