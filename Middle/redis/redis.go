package redis

import (
	"context"
	"encoding/json"
	"errors"
	redigo "github.com/gomodule/redigo/redis"
	"math"
	"reflect"
	"strconv"
	"time"

	"git.xiaojukeji.com/falcon/pope-action-executor/common/logger"
)

var ErrNotFound = errors.New("key not found")
var ErrSwitchOff = errors.New("switch off")

func (r *RedisClient) Get(ctx context.Context, key string, data interface{}) error {
	//logger := middleware.GetDidiContext(ctx).Logger

	reply, err := r.do(ctx, "GET", key)
	if err != nil {
		return err
	}

	if reply == nil {
		logger.Trace.Errorf(ctx, logger.DLTagRedisFailed, "%s%s",
			logger.BuildLogByMap(map[string]interface{}{"key": key}), "the key is not found in redis.")
		return ErrNotFound
	}

	err = parseReply(reply, data)
	if err != nil {
		logger.Trace.Errorf(ctx, logger.DLTagRedisFailed, "%s%s",
			logger.BuildLogByMap(map[string]interface{}{"key": key, "err": err.Error()}), "parse redis reply fail.")
		return err
	}

	return nil
}

func (r *RedisClient) MGet(ctx context.Context, keys ...interface{}) ([]string, error) {
	reply, err := r.do(ctx, "MGET", keys...)
	if err != nil {
		return nil, err
	}

	res, err := Strings(reply, err)
	if err != nil {
		logger.Trace.Errorf(ctx, logger.DLTagRedisFailed, "%s%s",
			logger.BuildLogByMap(map[string]interface{}{"err": err.Error()}), "call Strings fail")
		return nil, err
	}

	return res, nil
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}) error {
	_, err := r.do(ctx, "SET", key, value)
	return err
}

func (r *RedisClient) SetEx(ctx context.Context, key string, value interface{}, expire time.Duration) error {
	_, err := r.do(ctx, "SETEX", key, int32(expire/time.Second), value)
	return err
}

func (r *RedisClient) Do(ctx context.Context, args ...interface{}) (interface{}, error) {
	return r.do(ctx, "SET", args...)
}

func (r *RedisClient) HMSET(ctx context.Context, key ...interface{}) error {
	if 3 > len(key) || len(key)%2 == 0 {
		return errors.New("the command HMSET must go with two param at least")
	}

	_, err := r.do(ctx, "HMSET", key...)

	return err
}

func (r *RedisClient) HMGET(ctx context.Context, key ...interface{}) ([]string, error) {
	if 3 > len(key) {
		return nil, errors.New("the command HMGET must go with two param at least")
	}

	reply, err := r.do(ctx, "HMGET", key...)
	res, err := Strings(reply, err)
	if err != nil {
		logger.Trace.Errorf(ctx, logger.DLTagRedisFailed, "%s%s",
			logger.BuildLogByMap(map[string]interface{}{"err": err.Error()}), "call Strings fail")
		return nil, err
	}

	return res, nil
}

func (r *RedisClient) HSET(ctx context.Context, key ...interface{}) error {
	if 3 != len(key) {
		return errors.New("the command HSET must go with three param at most")
	}

	_, err := r.do(ctx, "HSET", key...)

	return err
}

func (r *RedisClient) HGET(ctx context.Context, key ...interface{}) (string, error) {
	if 2 != len(key) {
		return "", errors.New("the command HGET must go with two param at most")
	}

	reply, err := r.do(ctx, "HGET", key)
	res, err := String(reply, err)
	if err != nil {
		logger.Trace.Errorf(ctx, logger.DLTagRedisFailed, "%s%s",
			logger.BuildLogByMap(map[string]interface{}{"err": err.Error()}), "call Strings fail")
		return "", err
	}

	return res, nil
}

func (r *RedisClient) HGETALL(ctx context.Context, key interface{}) (map[string]string, error) {
	reply, err := r.do(ctx, "HGETALL", key)

	res, err := StringMap(reply, err)
	if err != nil {
		logger.Trace.Errorf(ctx, logger.DLTagRedisFailed, "%s%s",
			logger.BuildLogByMap(map[string]interface{}{"err": err.Error()}), "call Strings fail")
		return nil, err
	}

	return res, nil
}

func (r *RedisClient) SETNX(ctx context.Context, key, val string, expire int64) (bool, error) {
	v, err := r.do(ctx, "SET", key, val, "ex", expire, "nx")
	if err != nil {
		logger.Trace.Errorf(ctx, logger.DLTagRedisFailed, "%s%s",
			logger.BuildLogByMap(map[string]interface{}{"err": err.Error()}), "call Strings fail")
		return false, err
	} else {
		if v == nil {
			return false, nil
		} else {
			return true, nil
		}
	}
}

func (r *RedisClient) Delete(ctx context.Context, key ...interface{}) (int, error) {
	if 0 == len(key) {
		return 0, errors.New("the command DEL must go with one param at least")
	}

	reply, err := r.do(ctx, "DEL", key...)
	if err != nil {
		return 0, err
	}
	res, err := Int(reply, err)
	if err != nil {
		logger.Trace.Errorf(ctx, logger.DLTagRedisFailed, "%s%s",
			logger.BuildLogByMap(map[string]interface{}{"err": err.Error()}), "call Int fail")
		return 0, err
	}

	return res, nil
}

func (r *RedisClient) Expire(ctx context.Context, key string, expire time.Duration) error {
	_, err := r.do(ctx, "EXPIRE", key, int32(expire/time.Second))

	return err
}

func (r *RedisClient) Eval(ctx context.Context, script *redigo.Script, keyAndArgs ...interface{}) (interface{}, error) {
	conn, err := r.getConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	reply, err := script.Do(conn, keyAndArgs...)
	writeLog(ctx, "", "EVAL", reply, err, keyAndArgs...)
	return reply, err
}

func (r *RedisClient) AllowN(ctx context.Context, key string, burst int, rate int, n int) (allowed int, retryAfter time.Duration, resetAfter time.Duration, err error) {
	period := time.Second.Seconds()
	v, err := r.Eval(ctx, allowN, key, burst, rate, period, n)
	if err != nil {
		return 0, 0, 0, err
	}
	values, ok := v.([]interface{})
	if !ok {
		return 0, 0, 0, errors.New("value cast error")
	}

	retryAfterF64, err := parseFloat64(values[2])
	if err != nil {
		return 0, 0, 0, err
	}

	resetAfterF64, err := parseFloat64(values[3])
	if err != nil {
		return 0, 0, 0, err
	}

	allowed = int(values[0].(int64))
	retryAfter = dur(retryAfterF64)
	resetAfter = dur(resetAfterF64)
	return
}

func dur(f float64) time.Duration {
	if f == -1 {
		return -1
	}
	return time.Duration(f * float64(time.Second))
}

func parseFloat64(reply interface{}) (f float64, err error) {
	switch v := reply.(type) {
	case []byte:
		s := string(v)
		f, err = strconv.ParseFloat(s, 64)
		return
	case string:
		f, err = strconv.ParseFloat(v, 64)
		return
	default:
		return math.NaN(), errors.New("invalid type")
	}
}

func parseReply(reply, data interface{}) error {
	objV := reflect.ValueOf(data)
	objT := objV.Type()

	// 判断obj的类型，必须是指针
	if objT.Kind() != reflect.Ptr {
		return errors.New("data is not a pointer")
	}

	objV = objV.Elem()
	objT = objT.Elem()

	if !objV.CanSet() {
		return errors.New("data can not be set")
	}

	switch objT.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		x, err := Int(reply, nil)
		if err != nil {
			return err
		}

		objV.SetInt(int64(x))

	case reflect.Int64:
		x, err := Int64(reply, nil)
		if err != nil {
			return err
		}

		objV.SetInt(x)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		x, err := Uint64(reply, nil)
		if err != nil {
			return err
		}

		objV.SetUint(x)

	case reflect.String:
		x, err := String(reply, nil)
		if err != nil {
			return err
		}

		objV.SetString(x)

	case reflect.Struct:
		x, err := Bytes(reply, nil)
		if err != nil {
			return err
		}

		err = json.Unmarshal(x, data)
		if err != nil {
			return err
		}
	}

	return nil
}
