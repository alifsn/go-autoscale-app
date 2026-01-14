package main

import (
	"fmt"
	"math"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//Simulasi beban CPU
	result := 0.0
	for i := 0; i < 10000000; i++ {
		result += math.Sqrt(float64(i)) * math.Sin(float64(i))
	}
	fmt.Fprintf(w, "Hello from pod %s\nCPU load result: %f\n", os.Getenv("HOSTNAME"), result)
	// fmt.Fprintf(w, "Hello from pod %s\n", os.Getenv("HOSTNAME"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
