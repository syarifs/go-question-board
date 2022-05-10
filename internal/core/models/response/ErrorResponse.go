package response

type Error struct {
	Error interface{} `json:"error"`
	Message string `json:"message"`
}
