package lihatsms

type Requestsms struct {
	Userid  string `json:"id"`
	Number  string `json:"number"`
	Message string `json:"message"`
}

type Respons struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
