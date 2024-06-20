/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/ohchloeho/recipescraper/pkg/controllers"

	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "recipescraper",
	Short: "A CLI tool that scrapes recipe info from URLs",
	Long:  `A minimalistic CLI tool for recipes`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var url string
		fmt.Print("Enter a recipe URL: ")
		fmt.Scanln(&url)

		var wg sync.WaitGroup

		c := colly.NewCollector()

		c.OnHTML("script[type='application/ld+json']", func(e *colly.HTMLElement) {
			wg.Add(1) // Increment the WaitGroup counter for each script tag found

			go func() {
				defer wg.Done() // Decrement the counter when the function completes

				jsonData := e.Text

				if strings.HasPrefix(jsonData, "[") {
					var dataArray []map[string]interface{}
					err := json.Unmarshal([]byte(jsonData), &dataArray)
					if err != nil {
						log.Fatal("Failed to parse JSON:", err)
					}

					controllers.ProcessJSONData(dataArray[0])
				} else {
					var data map[string]interface{}
					err := json.Unmarshal([]byte(jsonData), &data)
					if err != nil {
						log.Fatal("Failed to parse JSON:", err)
					}
					controllers.ProcessJSONData(data)
				}
			}()
		})

		c.Visit(url)

		wg.Wait() // Wait for all tasks to complete
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
