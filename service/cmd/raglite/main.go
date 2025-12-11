package main

import (
	"fmt"
	"log"
	"os"
    "net/http"
    "encoding/json"
    "github.com/iannsp/cafestore/internal/raglite"
)

var chat raglite.Chat
var uipath string
var datapath string
func main() {
	geminiApiKey := os.Getenv("GEMINI_API_KEY")
    uipath = os.Getenv("UI_PATH")
    datapath = os.Getenv("DATA_PATH")

	if geminiApiKey == "" {
		log.Fatal("Please set the GEMINI_API_KEY environment variable")
	}

	if uipath == "" {
		log.Fatal("Please set the UI_PATH environment variable")
	}

	if datapath == "" {
		log.Fatal("Please set the DATA_PATH environment variable")
	}

    hs := raglite.NewHttpServer("8080")
    chat = raglite.NewChat(geminiApiKey)
    chat.Prompt(loadPrompt(datapath))
    chat.AttachRoutes(&hs)
    hs.AttachRoutes("/", index)
    hs.AttachRoutes("/api/chat", handleMessage)

    err := hs.ListenAndServe()
    if err != nil{
        fmt.Println("Rag Lite Running at http://localhost:8080")
	    log.Fatal(err)
    }
}

// Helper function to read the JSON file
func loadPromptInstruction(filename string) string {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		// If file doesn't exist, log a warning and return empty string (or a default)
		log.Printf("Warning: Could not read %s: %v. Using default behavior.", filename, err)
		return ""
	}
    return string(fileData)
}


// Helper function to build prompt
func loadPrompt(datapath string) string{
    prompt := loadPromptInstruction(datapath+"prompt.txt")
    prompt = prompt + `\n ## 1. Árvore de Categorias \n` + loadPromptInstruction(datapath+"categorias.json")
    prompt = prompt + `\n ## 2. Catalogo\n` + loadPromptInstruction(datapath+"catalog_cafe_small.json")
    return prompt
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1. Parse Request
	var req raglite.ChatRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Message == "" {
		http.Error(w, "Mensagem Invalida.", http.StatusBadRequest)
        return
	}

	if len(req.Message) > 100 {
        log.Println(len(req.Message))
		http.Error(w, "Mensagem muito longa.", http.StatusBadRequest)
        return
	}


	w.Header().Set("Content-Type", "application/json")
    response, err := chat.ProcessMessage(req.Message)
    if err != nil {
        jsonError(w, err.Error())
        return
    }
	json.NewEncoder(w).Encode(response)
}

func index(w http.ResponseWriter, r *http.Request) {
    index := uipath + "raglite.html"
    http.ServeFile(w, r, index)
}

func jsonError(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(raglite.ChatResponse{Error: msg})
}


