package d

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	R "github.com/gomodule/redigo/redis"
	"log"
	"server/pkg/setting"
	"time"
)

type User struct {
	Id      int64     `json:"id"`
	Uname   string    `json:"uname"`
	Upwd    string    `json:"upwd"`
	Pnumber string    `json:"pnumber"`
	Ctime   time.Time `json:"ctime"`
	Ltime   time.Time `json:"ltime"`
}

type Task struct {
	Id       int64     `json:"id"`
	Createby string    `json:"createby"`
	Tittle   string    `ison:"tittle"`
	Category int       `json:"category"`
	Detail   string    `json:"detail"`
	Question string    `json:"question"`
	Tperiod  int8      `json:"tperiod"`
	Addition string    `json:"addition"`
	Ctime    time.Time `json:"ctime"`
	Utime    time.Time `json:"utime"`
}

var db *sql.DB
var redis *R.Pool

func init() {
	err := dbinit()
	if err != nil {
		log.Println("db err: " + err.Error())
	}
	err = redisinit()
	if err != nil {
		log.Println("reids err: " + err.Error())
	}
}

func dbinit() error {
	var (
		err                          error
		dbType, user, password, host string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		return err
	}

	dbType = sec.Key("TYPE").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()

	//不指定数据库的名称
	//查询：stmt,err := db.Prepare("select * from upwork.user where pnumber=? and upwd = ? limit 1;")
	db, err = sql.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8&parseTime=true",
		user,
		password,
		host))

	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(10000)

	return nil
}

func redisinit() error {
	var (
		err         error
		Host        string
		Password    string
		MaxIdle     int
		MaxActive   int
		IdleTimeout int
	)
	sec, err := setting.Cfg.GetSection("redis")
	if err != nil {
		return err
	}

	Host = sec.Key("HOST").String()
	Password = sec.Key("PASSWORD").String()
	MaxIdle, _ = sec.Key("MAXIDLE").Int()
	MaxActive, _ = sec.Key("MAXACTIVE").Int()
	IdleTimeout, _ = sec.Key("IDLETIMEOUT").Int()

	redis = &R.Pool{
		MaxIdle:     MaxIdle,
		MaxActive:   MaxActive,
		IdleTimeout: time.Duration(IdleTimeout) * time.Second, //注意time.Duration(IdleTimeout)
		Dial: func() (R.Conn, error) {
			c, err := R.Dial("tcp", Host)
			if err != nil {
				return nil, err
			}
			if Password != "" {
				if _, err = c.Do("AUTH", Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c R.Conn, t time.Time) error {
			_, err = c.Do("PING")
			return err
		},
	}
	return nil
}

func redisSet(key string, data interface{}, time int) error {
	conn := redis.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

func redisExists(key string) bool {
	conn := redis.Get()
	defer conn.Close()

	exists, err := R.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func redisGet(key string) (string, error) {
	conn := redis.Get()
	defer conn.Close()

	reply, err := R.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return reply, nil
}

func redisDelete(key string) (bool, error) {
	conn := redis.Get()
	defer conn.Close()
	return R.Bool(conn.Do("DEL", key))
}
