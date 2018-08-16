
package main
import (
        "fmt"
        "./services"  // NEW
        zipkin "github.com/openzipkin/zipkin-go"
	      zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	      logreporter "github.com/openzipkin/zipkin-go/reporter/log"
)
var appName = "productservice"
func main() {
        fmt.Printf("Starting %v\n", appName)
        services.StartWebServer("6767")           // NEW
}
