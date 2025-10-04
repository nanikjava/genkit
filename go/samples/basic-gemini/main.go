// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
	"google.golang.org/genai"
)

type WeatherInput struct {
	Location string `json:"location" jsonschema_description:"Location to get weather for"`
}

func main() {
	ctx := context.Background()

	// Initialize Genkit with the Google AI plugin. When you pass nil for the
	// Config parameter, the Google AI plugin will get the API key from the
	// GEMINI_API_KEY or GOOGLE_API_KEY environment variable, which is the recommended
	// practice.
	g := genkit.Init(ctx, genkit.WithPlugins(&googlegenai.GoogleAI{}))

	// Define a simple flow that generates jokes about a given topic
	genkit.DefineFlow(g, "jokesFlow", func(ctx context.Context, input string) (string, error) {
		resp, err := genkit.Generate(ctx, g,
			ai.WithModelName("googleai/gemini-flash-latest"),
			ai.WithConfig(&genai.GenerateContentConfig{
				Temperature: genai.Ptr[float32](1.0),
				ThinkingConfig: &genai.ThinkingConfig{
					ThinkingBudget: genai.Ptr[int32](0),
				},
			}),
			ai.WithPrompt(`Tell short jokes about %s`, input))
		if err != nil {
			return "", err
		}

		return resp.Text(), nil
	})

	// genkit.DefineTool(
	// 	g, "getWeather", "Gets the current weather in a given location",
	// 	func(ctx *ai.ToolContext, input WeatherInput) (string, error) {
	// 		// Here, we would typically make an API call or database query. For this
	// 		// example, we just return a fixed value.
	// 		log.Printf("Tool 'getWeather' called for location: %s", input.Location)
	// 		return fmt.Sprintf("The current weather in %s is 63°F and sunny.", input.Location), nil
	// 	})

	// PromptWithOutputTypeDotprompt(ctx, g)

	<-ctx.Done()
}

func PromptWithOutputTypeDotprompt(ctx context.Context, g *genkit.Genkit) {
	type countryData struct {
		Name      string `json:"name"`
		Language  string `json:"language"`
		Habitants int    `json:"habitants"`
		GDP       int    `json:"gdp"`
	}
	type countries struct {
		Countries []countryData `json:"countries"`
	}

	prompt := genkit.LoadPrompt(g, "./prompts/hello.prompt", "countries")
	if prompt == nil {
		log.Printf("empty prompt")
		return
	}

	// Call the model.
	resp, err := prompt.Execute(ctx)
	if err != nil {
		log.Fatal(err)
	}
	var c countries
	if err = resp.Output(&c); err != nil {
		log.Fatal(err)
	}

	pretty, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(pretty))
}
