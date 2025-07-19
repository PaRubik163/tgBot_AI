package models
//Здесь все структруты (запросы, ответы)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Body struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Response struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}