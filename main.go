package main

import (
	"fmt"
	"intropia/rg"
	"net/http"
	"strconv"
	"time"
)

func main() {

	// we can host here a web server and use the random string as a json response
	// creating web server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// read options from query string if provided
		options := rg.Options{}

		// Read query parameters
		queryParams := r.URL.Query()

		// set default options
		options.Length = 16
		options.Alpha = true
		options.Numeric = true
		options.Special = false
		options.UseUpper = false
		options.UseLower = true
		// set options from query string
		if lengthStr := queryParams.Get("length"); lengthStr != "" {
			// Parse and set the "length" parameter
			length, err := strconv.Atoi(lengthStr)
			if err == nil {
				options.Length = length
			}
		}

		if alphaStr := queryParams.Get("alpha"); alphaStr != "" {
			// Parse and set the "alpha" parameter
			options.Alpha, _ = strconv.ParseBool(alphaStr)
		}

		if numericStr := queryParams.Get("numeric"); numericStr != "" {
			// Parse and set the "numeric" parameter
			options.Numeric, _ = strconv.ParseBool(numericStr)
		}

		if specialStr := queryParams.Get("special"); specialStr != "" {
			// Parse and set the "special" parameter
			options.Special, _ = strconv.ParseBool(specialStr)
		}

		if useUpperStr := queryParams.Get("useUpper"); useUpperStr != "" {
			// Parse and set the "useUpper" parameter
			options.UseUpper, _ = strconv.ParseBool(useUpperStr)
		}

		if useLowerStr := queryParams.Get("useLower"); useLowerStr != "" {
			// Parse and set the "useLower" parameter
			options.UseLower, _ = strconv.ParseBool(useLowerStr)
		}

		// set header to json
		w.Header().Set("Content-Type", "application/json")
		// measure the time
		start := time.Now()
		randomString, errGen := rg.GenerateRandomString(options)
		elapsed := time.Since(start)
		fmt.Printf("Time elapsed: %s \n", elapsed)

		if errGen != nil {
			// return error
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, errGen.Error())))
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}
		}

		fmt.Printf("Random string: %s \n", randomString)
		// write the response
		_, err := w.Write([]byte(fmt.Sprintf(`{"random_string": "%s"}`, randomString)))
		if err != nil {
			// return error
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}

		}
	})
	// print starting message
	fmt.Println("Starting server on port 8080")
	// start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
}
