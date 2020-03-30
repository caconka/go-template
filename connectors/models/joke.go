package models

type Joke struct {
	ID      string `json:"id" binding:"required"`
	Joke    string `json:"joke,omitempty"`
	Status  int    `json:"status,omitempty" binding:"required"`
	Message string `json:"message,omitempty"`
}
