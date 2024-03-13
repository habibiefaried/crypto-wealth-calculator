package main

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// config.ParseEnv: will parse env var in string value. eg: shell: ${SHELL}
	config.WithOptions(config.ParseEnv)
	// add driver for support yaml content
	config.AddDriver(yaml.Driver)
	err := config.LoadFiles("./config.yml")
	if err != nil {
		panic(err)
	}

	address := strings.ToLower("0xDF1A2934Bf91dF67B3f1972d539a388B73E29dD2")
	url := fmt.Sprintf("https://api.bscscan.com/api?module=account&action=tokentx&address=%v&sort=asc&apikey=%v", address, config.String("bscapi"))
	arrayWealth := make(map[string]string)

	// Perform a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	// Read the body of the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Unmarshal the JSON data into an ApiResponse struct
	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Iterate through the results and print them
	for _, result := range apiResponse.Result {
		key := result.TokenSymbol

		_, exists := arrayWealth[key]
		if !exists {
			arrayWealth[key] = result.ContractAddress
		}
	}

	for k, v := range arrayWealth {
		url = fmt.Sprintf("https://api.bscscan.com/api?module=account&action=tokenbalance&contractaddress=%v&address=%v&tag=latest&apikey=%v", v, address, config.String("bscapi"))
		resp, err = http.Get(url)
		if err != nil {
			fmt.Println("Error fetching data:", err)
			return
		}
		defer resp.Body.Close()

		// Read the body of the response
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		var ApiBalance ApiBalance
		err = json.Unmarshal(body, &ApiBalance)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}

		got, err := ConvertStringToFloat(ApiBalance.Result, 18)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("[%v] = %.5f\n", k, got)
	}

	// BNB Account
	url = fmt.Sprintf("https://api.bscscan.com/api?module=account&action=balance&address=%v&apikey=%v", address, config.String("bscapi"))
	resp, err = http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	// Read the body of the response
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var ApiBalance ApiBalance
	err = json.Unmarshal(body, &ApiBalance)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	got, err := ConvertStringToFloat(ApiBalance.Result, 18)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("[BNB (MAIN ACCOUNT)] = %.5f\n", got)
}
