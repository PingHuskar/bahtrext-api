package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	// "time"
    "io/ioutil"
    // "reflect"
    "encoding/json"
)

const serverPort = 3333


func main() {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s /\n", r.Method)
		})
		server := http.Server{
			Addr:    fmt.Sprintf(":%d", serverPort),
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running http server: %s\n", err)
			}
		}
	}()

	// time.Sleep(100 * time.Millisecond)

    YOUR_INPUT := "1234.56"
    
	requestURL := fmt.Sprintf("http://localhost:%d/br/%v", 3000, YOUR_INPUT)

	// requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	// fmt.Printf("client: got response!\n")
	// fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
    // fmt.Println(resBody)
	// fmt.Printf("client: response body: %s\n", resBody)
    // fmt.Println(reflect.TypeOf(resBody))
    type Dat struct {
        val  string `json:"val"`
        typ string `json:"typ"`
        err string `json:"err"`
        txt string `json:"txt"`
    }

    byt := []byte(resBody)
    var dat Dat
    err = json.Unmarshal(byt, &dat)
    if err != nil {
        panic(err)
    }
    var data map[string]interface{}
	err = json.Unmarshal([]byte(resBody), &data)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	fmt.Printf("Result: %v\n", data["txt"])
}
