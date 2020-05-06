import (
      "encoding/json"
      "fmt"
)

// Pretty Print struct, map, array and slice
func PrettyPrint(v interface{}) (err error) {
      b, err := json.MarshalIndent(v, "", "  ")
      if err == nil {
              fmt.Println(string(b))
      }
      return
}
