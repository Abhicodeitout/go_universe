package main

import (
	"fmt"
	"strings"
)

func main() {
	// Define a map of possible customer queries and corresponding responses
	queries := map[string]string{
		"hello":      "Hi there! How can I help you today?",
		"help":       "Sure, what do you need help with?",
		"order":      "Can you provide your order number?",
		"tracking":   "Let me check on the status of your order for you.",
		"complaint":  "I'm sorry to hear that. Let me escalate this to our customer service team for you.",
		"cancel":     "To cancel your order, please contact our customer service team at 1-800-123-4567.",
		"thanks":     "You're welcome! Is there anything else I can assist you with?",
		"goodbye":    "Goodbye! Have a great day!",
		"default":    "I'm sorry, I didn't understand your query. Could you please try again?",
	}

	fmt.Println("Welcome to the customer service chatbot!")

	for {
		fmt.Print("Customer query: ")
		var query string
		fmt.Scanln(&query)

		// Convert query to lowercase and remove whitespace
		query = strings.ToLower(strings.TrimSpace(query))

		// Check if query matches a known query in the map
		response, ok := queries[query]
		if !ok {
			// If query does not match a known query, use the default response
			response = queries["default"]
		}

		// Print the bot's response
		fmt.Println("Bot response:", response)

		// If the customer says goodbye, exit the loop
		if query == "goodbye" {
			break
		}
	}
}
