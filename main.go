package main 
import (
	"fmt"
	"net/http"
	"time"
	"os"
	"io/ioutil"
)

// TODO: Network test 
func main() {
     benchmark(os.Args[1]) 
}

func benchmark(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Connection error: %f", time.Since(start).Seconds())
	        return
	}
        num, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	result := ((float64(len(num)) * 8) / time.Since(start).Seconds()) / 1000000
	fmt.Printf("Speed: %f MBPS\n", result)
}
