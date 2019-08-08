package pair

import (
	"kv-storage/internal/pkg/db"
	"kv-storage/internal/pkg/models"
)

func GetPair(key string) (models.PairData, error) {
	output, err := db.PairGet(key)
	if err != nil {
		return models.PairData{}, err
	}

	return models.PairData{
		Key:   output.Key,
		Value: output.Value,
	}, nil
}

func CreatePair(input models.CreatePairData) (err error) {
	_, err = db.PairGet(input.Key)
	if err == nil {
		return models.AlreadyExistsError
	} else if err != models.NotFoundError {
		return
	}

	// Create user
	return db.PairCreate(db.PairData{
		Key:   input.Key,
		Value: input.Value,
	})
}

func UpdatePair(key string, input models.UpdatePairData) (err error) {
	_, err = db.PairGet(key)
	if err != nil {
		return
	}

	return db.PairUpdate(db.PairData{
		Key:   key,
		Value: input.Value,
	})
}

func RemovePair(key string) error {
	_, err := db.PairGet(key)
	if err != nil {
		return err
	}

	return db.PairRemove(key)
}
