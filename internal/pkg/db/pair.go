package db

type PairData struct {
	Key   string      `json:"key, string" example:"test"`
	Value interface{} `json:"value" example:"{SOME ARBITRARY JSON}"`
}

func PairCreate(data PairData) error {
	return create([]interface{}{data.Key, data.Value})
}

func PairGet(key string) (data PairData, err error) {
	output, err := getOne("primary", key)
	if err != nil {
		return data, err
	}

	return PairData{
		Key:   output[0].(string),
		Value: getStrMap(output[1]),
	}, nil
}

func PairUpdate(data PairData) error {
	return update([]interface{}{data.Key, data.Value})
}

func PairRemove(key string) error {
	return remove("primary", key)
}
