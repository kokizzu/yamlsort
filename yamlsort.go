package main
import "os"
import "github.com/goccy/go-yaml"
import "github.com/kokizzu/gotro/L"
import "github.com/kokizzu/gotro/M"
import "fmt"
func main() {
	if len(os.Args) < 2 {
		fmt.Println(`missing argument [fileToBeSorted].yaml`)
		return
	}
	file := L.ReadFile(os.Args[1])
        m := M.SX{}
        _ = yaml.Unmarshal([]byte(file), &m)
        byt, _ := yaml.Marshal(m)
        fmt.Println(string(byt))
}
