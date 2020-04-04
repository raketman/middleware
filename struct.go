package middleware

const StatusSuccess = "success"
const StatusError = "error"

type Client struct {
	Code 		string // код клиента, соответствует iss
	Secret		string
	Alg			string
	Default 	bool
}

type Token struct {
	Token string
}

type Response struct {
	Status string // статус операции
	Payload string // полезная нагрузка
	Message string // текст статуса
}

type Error struct {
	Message string // текст ошибки
}

func (e *Error) Error() string {
	return e.Message
}
