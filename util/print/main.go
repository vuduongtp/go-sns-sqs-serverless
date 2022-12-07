package print

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// PrettyPrint print interface as JSON representation
func PrettyPrint(i interface{}) {
	fmt.Println("<<==============")
	fmt.Printf("Type: %s\n", reflect.TypeOf(i))
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
	fmt.Println("==============>>")
}
