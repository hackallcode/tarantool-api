package models

/********************
 *    IN MODELS     *
 ********************/

/* SIGN UP DATA */

//easyjson:json
type CreatePairData struct {
	Key   string `json:"key, string" example:"test"`
	Value interface{} `json:"value" example:"{SOME ARBITRARY JSON}"`
}

func (v CreatePairData) Validate() bool {
	return v.Key != "" && v.Value != nil
}

/* UPDATE USER DATA */

//easyjson:json
type UpdatePairData struct {
	Value interface{} `json:"value" example:"{SOME ARBITRARY JSON}"`
}

func (v UpdatePairData) Validate() bool {
	return v.Value != nil
}

/********************
 *    OUT MODELS    *
 ********************/

/* USER DATA */

//easyjson:json
type PairData struct {
	Key   string `json:"key, string" example:"test"`
	Value interface{} `json:"value" example:"{SOME ARBITRARY JSON}"`
}
