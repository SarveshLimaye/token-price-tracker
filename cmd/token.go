package cmd

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"log"
	"github.com/spf13/cobra"
	"encoding/json"
	"github.com/joho/godotenv"
)



// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Get realtime price of token on uniswap v3",
	Long:  `Paste address  to view the real-time price of the token in $. For example: go run main.go token 0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0`,
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
		  log.Fatal("Error loading .env file")
		}

		if len(args) < 1 {
			fmt.Println("Please enter a token address")
			return
		}

		token := args[0]

		url := "https://deep-index.moralis.io/api/v2.2/erc20/" + token + "/price?chain=eth&include=percent_change&exchange=uniswapv3"
		req, err := http.NewRequest("GET",url,nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Add("Accept", "application/json")
		req.Header.Add("X-API-Key", os.Getenv("MORALIS_API_KEY"))
	  
		res, _ := http.DefaultClient.Do(req)
	  
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
	  
		var result map[string]interface{}
		json.Unmarshal([]byte(body), &result)

	
		
		if result["error"] != nil {
			fmt.Println(result["message"])
			return
		}

		tokenName := result["tokenName"]
		tokenPrice := result["usdPrice"]
		if(tokenPrice == nil){
			fmt.Println("Token not found. Please enter a valid token address")
			return
		}

		fmt.Println("Price of", tokenName, "is $", tokenPrice)
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)
}
