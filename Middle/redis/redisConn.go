package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"git.xiaojukeji.com/falcon/pope-action-executor/common/logger"
	"git.xiaojukeji.com/falcon/pope-action-executor/common/util"
	redigo "github.com/gomodule/redigo/redis"
)

type RedisConfig struct {
	Servers      []string
	SwitchOff    bool
	MaxIdle      int
	MaxActive    int
	IdleTimeout  time.Duration
	ConnTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Password     string
}

const (
	REDIS_AUTH_KEY = "AUTH"
)

func New(conf *RedisConfig) (*RedisClient, error) {
	return newRedisPool(conf)
}

type RedisClient struct {
	Pool map[string]*redigo.Pool
	Conf *RedisConfig
}

func newRedisPool(conf *RedisConfig) (*RedisClient, error) {
	if len(conf.Servers) == 0 {
		return nil, errors.New("invalid servers")
	}

	var client RedisClient
	client.Conf = conf
	client.Pool = make(map[string]*redigo.Pool, len(conf.Servers))
	for _, srv := range client.Conf.Servers {
		p := &redigo.Pool{
			MaxIdle:     conf.MaxIdle,
			MaxActive:   conf.MaxActive,
			IdleTimeout: conf.IdleTimeout,
			Dial: func(serverAddr string) func() (redigo.Conn, error) {
				return func() (redigo.Conn, error) {
					c, err := redigo.DialTimeout("tcp", serverAddr,
						conf.ConnTimeout,
						conf.ReadTimeout,
						conf.WriteTimeout)
					if err != nil {
						return nil, err
					}
					//在配置密码不为空的时候 用密码去鉴权
					password := strings.Trim(conf.Password, " ")
					if password != "" {
						if _, err := c.Do(REDIS_AUTH_KEY, password); err != nil {
							return nil, err
						}
					}
					return c, err
				}
			}(srv),
		}
		client.Pool[srv] = p
	}

	return &client, nil
}

func (r *RedisClient) do(ctx context.Context, commandName string, args ...interface{}) (interface{}, error) {
	// 若开关关闭，直接返回
	if r.Conf.SwitchOff {
		logger.Trace.Infof(ctx, logger.DLTagRedisSuccess, "%s%s",
			logger.BuildLogByMap(map[string]interface{}{"method": commandName}), "redis switch is turned off")
		return nil, ErrSwitchOff
	}

	// 随机获取连接
	index := rand.Intn(len(r.Conf.Servers))

	host := r.Conf.Servers[index]
	p, ok := r.Pool[host]
	if !ok {
		return nil, fmt.Errorf("get server fail. host: %s", host)
	}

	//ctx.Set("redis_start_time", time.Now())
	conn := p.Get()
	defer conn.Close()

	reply, err := conn.Do(commandName, args...)

	writeLog(ctx, host, commandName, reply, err, args...)

	if err != nil {
		return nil, err
	} else {
		return reply, nil
	}
}

// 如果需要redis的事务、watch等，可以提供该接口
func (r *RedisClient) getConn() (redigo.Conn, error) {
	host := r.Conf.Servers[0]
	if len(r.Conf.Servers) > 1 {
		index := rand.Intn(len(r.Conf.Servers))
		host = r.Conf.Servers[index]
	}
	p, ok := r.Pool[host]
	if !ok {
		return nil, fmt.Errorf("get server fail. host: %s", host)
	}
	conn := p.Get()
	return conn, nil
}

// 写redis操作日志
func writeLog(ctx context.Context, host, commandName string, reply interface{}, err error, args ...interface{}) {
	//didiCtx := middleware.GetDidiContext(ctx)
	//if didiCtx == nil || didiCtx.Logger == nil {
	//	return
	//}

	// 设置dltag标准参数
	//startTime := ctx.Get("redis_start_time").(time.Time)
	startTime := time.Now().UnixNano()

	logMap := map[string]interface{}{
		"method":    commandName,
		"proc_time": util.CalcCost(startTime),
	}
	ipPort := strings.Split(host, ":")
	logMap["host"] = ipPort[0]
	logMap["port"] = ""
	if len(ipPort) > 1 {
		logMap["port"] = ipPort[1]
	}
	var argSLice []interface{}
	for _, arg := range args {
		if argBytes, ok := arg.([]byte); ok {
			argSLice = append(argSLice, string(argBytes))
		} else {
			argSLice = append(argSLice, arg)
		}
	}
	argJson, e := json.Marshal(argSLice)
	if e == nil {
		logMap["args"] = fmt.Sprintf("%s", argJson)
	}

	if err != nil {
		logger.Trace.Infof(ctx, logger.DLTagRedisFailed, "%s%s", logger.BuildLogByMap(logMap), err.Error())
	} else {
		switch commandName {
		case "SET", "MSET", "SETEX", "HSET", "HMSET", "EXPIRE", "DEL", "SETNX":
			logMap["result"] = fmt.Sprint(reply)
		case "GET", "HGET":
			if reply == nil {
				logMap["result"] = fmt.Sprint(reply)
			} else if val, ok := reply.([]byte); ok {
				logMap["result"] = string(val)
			}
		case "MGET", "HMGET", "HGETALL":
			if arr, ok := reply.([]interface{}); ok {
				slice := make([]interface{}, 0, len(arr))
				for _, val := range arr {
					if val == nil {
						slice = append(slice, val)
					} else {
						slice = append(slice, string(val.([]byte)))
					}
				}
			}
		default:
			logMap["result"] = fmt.Sprint(reply)
		}
		logger.Trace.Infof(ctx, logger.DLTagRedisSuccess, "%s", logger.BuildLogByMap(logMap))
	}
}
