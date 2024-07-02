package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func LogJSON(value any) {
	json, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))
}

func SaveJSON(value any) {
	json, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("/home/op/code/graphqlops-go/schema.json", json, 0644)
	if err != nil {
		panic(err)
	}
}
