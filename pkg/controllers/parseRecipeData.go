package controllers

import (
	"fmt"
	"strconv"

	"github.com/ohchloeho/recipescraper/pkg/utils"
)

func ProcessJSONData(data map[string]interface{}) {

	// Access recipe name fields
	if name, ok := utils.FindNestedField(data, "headline"); ok {
		fmt.Printf("Recipe Name: %s\n", name)
	}

	if servings, ok := utils.FindNestedField(data, "recipeYield"); ok {
		result := utils.ConvertToString(servings)
		num, err := strconv.Atoi(result[0])
		if err != nil {
			result[0] = "error converting string to int"
		}
		fmt.Printf("Recipe Serves: %v\n", num)
	}

	if ingredients, ok := utils.FindNestedField(data, "recipeIngredient"); ok {
		result := utils.ConvertToString(ingredients)
		// fmt.Printf("Recipe Ingredients: %s\n", result)
		fmt.Println("Ingredients:")
		for _, res := range result {
			fmt.Printf("• %s\n", res)
		}
	}

	if instructions, ok := utils.FindNestedField(data, "recipeInstructions"); ok {
		strings, _ := utils.ConvertInterfaceToSliceOfStrings(instructions)
		count := 1
		fmt.Println("Instructions:")
		for _, str := range strings {
			fmt.Printf("%v. %s\n", count, str)
			count += 1
		}
	}

	if prepTime, ok := utils.FindNestedField(data, "prepTime"); ok {
		fmt.Printf("Prep time: %s\n", prepTime)
	}
	if cookTime, ok := utils.FindNestedField(data, "cookTime"); ok {
		fmt.Printf("Cook time: %s\n", cookTime)
	}
	if totalTime, ok := utils.FindNestedField(data, "totalTime"); ok {
		fmt.Printf("Total time: %s\n", totalTime)
	}
}
