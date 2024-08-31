package main

import (
    "context"
    "log"
    "os"

    "github.com/google/generative-ai-go/genai"
    "google.golang.org/api/option"
)

func main() {
    log.Println("Starting Gemini App")

    // Create a context
    ctx := context.Background()

    // Initialize the client with the API key
    apiKey := os.Getenv("GOOGLE_API_KEY")
    if apiKey == "" {
        log.Fatal("GOOGLE_API_KEY environment variable is not set")
    }

    client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
    if err != nil {
        log.Fatalf("Failed to create Generative AI client: %v", err)
    }

    // Define the generative model to use
    model := client.GenerativeModel("gemini-1.5-flash")

    // Dummy image data placeholder; replace with actual image data
    imgData := []byte{}

    // Generate content using the model
    resp, err := model.GenerateContent(
        ctx,
        genai.Text("What's in this photo?"),
        genai.ImageData("jpeg", imgData),
    )
    if err != nil {
        log.Fatalf("Failed to generate content: %v", err)
    }

    // Log the response (example logging, adjust according to actual response structure)
    log.Printf("Generated content: %+v\n", resp)
}

