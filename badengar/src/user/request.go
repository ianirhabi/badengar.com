package user

type Requestlogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Respons struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
type RequestRegis struct {
	Nama     string `json:nama`
	Username string `json:username`
	Password string `json:password`
	Simsatu  string `json:sim1`
	Simdua   string `json:sim2`
	Tanggal  string `json:tanggal`
	Jam      string `json:jam`
}
