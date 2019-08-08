package db

import (
	"errors"
	"time"

	"github.com/tarantool/go-tarantool"
	"kv-storage/internal/pkg/config"
	"kv-storage/internal/pkg/logger"
	"kv-storage/internal/pkg/models"
)

/********************/
/*      ERRORS      */
/********************/

var (
	AlreadyInitError = errors.New("db already initialized")
	NotInitError     = errors.New("db wasn't initialized")
)

/********************/
/*  BASE FUNCTIONS  */
/********************/

var client *tarantool.Connection

func Ping() error {
	if client == nil {
		return NotInitError
	}

	_, err := client.Ping()
	return err
}

func Open() (err error) {
	if client != nil {
		return AlreadyInitError
	}

	server := config.Db.Host + ":" + config.Db.Port
	opts := tarantool.Opts{
		Timeout:       500 * time.Millisecond,
		Reconnect:     1 * time.Second,
		MaxReconnects: 3,
		User:          config.Db.Username,
		Pass:          config.Db.Password,
	}
	client, err = tarantool.Connect(server, opts)
	if err != nil {
		return err
	}

	return Ping()
}

func Close() error {
	if client == nil {
		return NotInitError
	}

	return client.Close()
}

func getOne(index, key interface{}) ([]interface{}, error) {
	if client == nil {
		return nil, NotInitError
	}

	logger.Debug("Select query with index = %v, key = %v", index, key)
	resp, err := client.Select(config.Db.Space, index, 0, 1, tarantool.IterEq, []interface{}{key})
	if err != nil || resp == nil {
		return nil, err
	}
	if len(resp.Data) == 0 {
		return nil, models.NotFoundError
	}
	return resp.Data[0].([]interface{}), nil
}

func create(tuple []interface{}) error {
	if client == nil {
		return NotInitError
	}

	logger.Debug("Insert query with tuple = %v", tuple)
	_, err := client.Insert(config.Db.Space, tuple)
	return err
}

func update(tuple []interface{}) error {
	if client == nil {
		return NotInitError
	}

	logger.Debug("Update query with tuple = %v", tuple)
	_, err := client.Replace(config.Db.Space, tuple)
	return err
}

func remove(index, key interface{}) error {
	if client == nil {
		return NotInitError
	}

	logger.Debug("Delete query with index = %v, key = %v", index, key)
	_, err := client.Delete(config.Db.Space, index, []interface{}{key})
	return err
}

func getStrMap(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = getStrMap(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = getStrMap(v)
		}
	}
	return i
}