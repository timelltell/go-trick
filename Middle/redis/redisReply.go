package redis

import (
	redigo "github.com/gomodule/redigo/redis"
)

func Int(reply interface{}, err error) (int, error) {
	return redigo.Int(reply, err)
}

func Int64(reply interface{}, err error) (int64, error) {
	return redigo.Int64(reply, err)
}

func Uint64(reply interface{}, err error) (uint64, error) {
	return redigo.Uint64(reply, err)
}

func Float64(reply interface{}, err error) (float64, error) {
	return redigo.Float64(reply, err)
}

func String(reply interface{}, err error) (string, error) {
	return redigo.String(reply, err)
}

func Bytes(reply interface{}, err error) ([]byte, error) {
	return redigo.Bytes(reply, err)
}

func Bool(reply interface{}, err error) (bool, error) {
	return redigo.Bool(reply, err)
}

func MultiBulk(reply interface{}, err error) ([]interface{}, error) {
	return redigo.MultiBulk(reply, err)
}

func Values(reply interface{}, err error) ([]interface{}, error) {
	return redigo.Values(reply, err)
}

func Strings(reply interface{}, err error) ([]string, error) {
	return redigo.Strings(reply, err)
}

func ByteSlices(reply interface{}, err error) ([][]byte, error) {
	return redigo.ByteSlices(reply, err)
}

func Ints(reply interface{}, err error) ([]int, error) {
	return redigo.Ints(reply, err)
}

func StringMap(result interface{}, err error) (map[string]string, error) {
	return redigo.StringMap(result, err)
}

func IntMap(result interface{}, err error) (map[string]int, error) {
	return redigo.IntMap(result, err)
}

func Int64Map(result interface{}, err error) (map[string]int64, error) {
	return redigo.Int64Map(result, err)
}
