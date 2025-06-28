package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	//Only exported fields will be encoded/decoded in JSON. Fields must start with capital letters to be exported.
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolb, _ := json.Marshal(true)
	fmt.Println(string(bolb))

	intb, _ := json.Marshal(1)
	fmt.Println(string(intb))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapJson := map[string]int{"apple": 1, "banana": 2}
	mapB, _ := json.Marshal(mapJson)
	fmt.Println(string(mapB))

	//	The JSON package can automatically encode your custom data types. It will only include exported fields in the encoded output and will by default use those names as the JSON keys.

	//The JSON package can automatically encode your custom data types. It will only include exported fields in the encoded output and will by default use those names as the JSON keys.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}

	res1B, _ := json.Marshal(res1D)
	fmt.Println("Response 1:", string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	//You can use tags on struct field declarations to customize the encoded JSON key names. Check the definition of response2 above to see an example of such tags.
	fmt.Println(string(res2B))

	//Now let’s look at decoding JSON data into Go values. Here’s an example for a generic data structure.
	byt := []byte(`{"page":1,"fruits":["apple","peach","pear"]}`)

	//We need to provide a variable where the JSON package can put the decoded data. This map[string]interface{} will hold a map of strings to arbitrary data types.
	var dat map[string]any

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	//	In order to use the values in the decoded map, we’ll need to convert them to their appropriate type. For example here we convert the value in num to the expected float64 type.
	num := dat["page"].(float64)
	fmt.Println(num)

	//Accessing nested data requires a series of conversions.
	strs := dat["fruits"].([]interface{})
	str1 := strs[0]
	fmt.Println(str1)

	//We can also decode JSON into custom data types. This has the advantages of adding additional type-safety to our programs and eliminating the need for type assertions when accessing the decoded data.
	str := `{"page":1,"fruits":["apple","pears"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	//In the examples above we always used bytes and strings as intermediates between the data and JSON representation on standard out. We can also stream JSON encodings directly to os.Writers like os.Stdout or even HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 1, "banana": 2}
	enc.Encode(d)

	//	Streaming reads from os.Readers like os.Stdin or HTTP request bodies is done with json.Decoder.
	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)

}
