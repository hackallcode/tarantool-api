package models

// 100-199 - success answer
// 200-299 - incorrect answer because of the user
// 300-399 - incorrect answer because developers

//easyjson:json
type MessageAnswer struct {
	Status  int    `json:"status, int" example:"100"`
	Message string `json:"message, string" example:"ok"`
}

/* SUCCESS ANSWERS */

func GetSuccessAnswer(message string) *MessageAnswer {
	return &MessageAnswer{
		Status:  100,
		Message: message,
	}
}

var PairCreatedAnswer = MessageAnswer{
	Status:  101,
	Message: "pair created",
}

var PairUpdatedAnswer = MessageAnswer{
	Status:  102,
	Message: "pair updated",
}

var PairRemovedAnswer = MessageAnswer{
	Status:  103,
	Message: "pair removed",
}

/* USERS ERRORS */

func GetUserErrorAnswer(error string) *MessageAnswer {
	return &MessageAnswer{
		Status:  200,
		Message: error,
	}
}

var IncorrectJsonAnswer = MessageAnswer{
	Status:  201,
	Message: "incorrect JSON",
}

var IncorrectParamAnswer = MessageAnswer{
	Status:  202,
	Message: "incorrect param",
}

var KeyExistsAnswer = MessageAnswer{
	Status:  203,
	Message: "key already exists",
}

var KeyNotFoundAnswer = MessageAnswer{
	Status:  204,
	Message: "key not found",
}

/* DEVELOPERS ERRORS */

func GetDeveloperErrorAnswer(error string) *MessageAnswer {
	return &MessageAnswer{
		Status:  300,
		Message: error,
	}
}
