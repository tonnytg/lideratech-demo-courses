package main

import (
    "context"
    "database/sql"
    "log"
    "os"

    "github.com/google/generative-ai-go/genai"
    _ "github.com/mattn/go-sqlite3"
    "google.golang.org/api/option"
)

// ModelConfig represents a configuration for a specific AI model.
type ModelConfig struct {
    ModelName string
    APIKey    string
}

// Initialize database
func initDB() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "./prompts.db")
    if err != nil {
        return nil, err
    }

    createTableQuery := `
    CREATE TABLE IF NOT EXISTS prompts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        model_name TEXT,
        prompt TEXT,
        response TEXT
    );`

    _, err = db.Exec(createTableQuery)
    if err != nil {
        return nil, err
    }

    return db, nil
}

// Save prompt and response to database
func savePrompt(db *sql.DB, modelName, prompt, response string) error {
    insertQuery := `INSERT INTO prompts (model_name, prompt, response) VALUES (?, ?, ?)`
    _, err := db.Exec(insertQuery, modelName, prompt, response)
    return err
}

// Retrieve prompts by model name
func getPromptsByModel(db *sql.DB, modelName string) ([]map[string]string, error) {
    rows, err := db.Query("SELECT id, prompt, response FROM prompts WHERE model_name = ?", modelName)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var prompts []map[string]string
    for rows.Next() {
        var id int
        var prompt, response string
        err = rows.Scan(&id, &prompt, &response)
        if err != nil {
            return nil, err
        }
        prompts = append(prompts, map[string]string{"id": string(id), "prompt": prompt, "response": response})
    }

    return prompts, nil
}

func main() {
    log.Println("Starting Multi-Model Gemini App")

    // Initialize the database
    db, err := initDB()
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    defer db.Close()

    // Define model configurations
    models := []ModelConfig{
        {"gemini-1.5-flash", os.Getenv("GOOGLE_API_KEY_GEMINI")},
        {"another-model", os.Getenv("GOOGLE_API_KEY_ANOTHER_MODEL")},
    }

    // Choose the model to use (for demonstration purposes, we'll use the first one)
    selectedModel := models[0]

    // Create a context
    ctx := context.Background()

    // Initialize the client for the selected model
    client, err := genai.NewClient(ctx, option.WithAPIKey(selectedModel.APIKey))
    if err != nil {
        log.Fatalf("Failed to create Generative AI client: %v", err)
    }

    // Define the generative model to use
    model := client.GenerativeModel(selectedModel.ModelName)

    // Dummy image data placeholder; replace with actual image data
    imgData := []byte{}

    // Define the prompt
    prompt := "What's in this photo?"

    // Generate content using the model
    resp, err := model.GenerateContent(
        ctx,
        genai.Text(prompt),
        genai.ImageData("jpeg", imgData),
    )
    if err != nil {
        log.Fatalf("Failed to generate content: %v", err)
    }

    // Assuming resp contains a string; adjust according to actual response structure
    responseText := "Response received from model" // Placeholder; replace with actual response extraction

    // Save the prompt and response to the database
    err = savePrompt(db, selectedModel.ModelName, prompt, responseText)
    if err != nil {
        log.Fatalf("Failed to save prompt: %v", err)
    }

    log.Printf("Prompt saved successfully for model %s", selectedModel.ModelName)

    // Example: Retrieve all prompts for the selected model
    prompts, err := getPromptsByModel(db, selectedModel.ModelName)
    if err != nil {
        log.Fatalf("Failed to retrieve prompts: %v", err)
    }

    // Display retrieved prompts
    for _, p := range prompts {
        log.Printf("ID: %s, Prompt: %s, Response: %s", p["id"], p["prompt"], p["response"])
    }
}

