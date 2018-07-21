
package main
import (
        "fmt"
        "./services"  // NEW
)
var appName = "productservice"
func main() {
        fmt.Printf("Starting %v\n", appName)
        services.StartWebServer("6767")           // NEW
}
