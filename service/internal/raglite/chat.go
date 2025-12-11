package raglite

import(
    "context"
	"github.com/google/generative-ai-go/genai"
    "google.golang.org/api/option"
)
type Chat struct{
    geminiApiKey string
    prompt string
    session *genai.ChatSession
}
// ChatRequest receives data from frontend

type ChatRequest struct {
	Message string `json:"message"`
}

// ChatResponse sends data to frontend
type ChatResponse struct {
	Reply string `json:"reply"`
	Error string `json:"error,omitempty"`
}

// SystemConfig maps the JSON file content
type SystemConfig struct {
	Instruction string `json:"instruction"`
}

func NewChat(geminiApiKey string) Chat{
    return Chat{
        geminiApiKey : geminiApiKey,
    }
}

func (c *Chat) AttachRoutes(hs *HttpServer){

}
func (c *Chat)Prompt(prompt string){
    c.prompt = prompt
}

func (c *Chat) startSession()  error{
	ctx := context.Background()
    client, err := genai.NewClient(ctx, option.WithAPIKey(c.geminiApiKey))
    if err != nil {
        return  err
    }
    model := client.GenerativeModel("gemini-flash-latest")
    c.session = model.StartChat()
    c.session.SendMessage(ctx, genai.Text(c.prompt))

    return nil
}
func (c *Chat) ProcessMessage(message string) (ChatResponse, error){
    var chatResponse ChatResponse
	ctx := context.Background()

    if c.session == nil {
        err := c.startSession()
        if err != nil {
            return chatResponse, err
        }
    }

	resp, err := c.session.SendMessage(ctx, genai.Text(message))
	if err != nil {
        return chatResponse, err
	}

	responseText := ""
	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		for _, part := range resp.Candidates[0].Content.Parts {
			if txt, ok := part.(genai.Text); ok {
				responseText += string(txt)
			}
		}
	}
    chatResponse.Reply = responseText
    return chatResponse, nil
}
