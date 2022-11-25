package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v8"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/11/25 17:53
 * @Desc:
 */

var leaseId snowflake.ID

func TestTryLock(t *testing.T) {
	// prepare
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       10,
	})
	_, err := redisClient.Ping(context.TODO()).Result()
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})

	// 测试tryLock
	key := "lockKey"
	node, err := snowflake.NewNode(1)
	if err != nil {
		t.Fatal(err)
	}
	leaseId = node.Generate()
	cancelFunc, err := TryLock(context.TODO(), redisClient, key, 10, 4, leaseId)
	if err != nil {
		Println(err)
		return
	} else {
		fmt.Println("上锁成功, 睡眠30秒")
		time.Sleep(30 * time.Second)

		fmt.Println("开始释放锁")
		b, err := UnLock(context.TODO(), redisClient, cancelFunc, key, leaseId)
		Println(err)
		fmt.Println("释放锁成功", leaseId, b)
		go func() {
			done <- struct{}{}
		}()
	}
	<-done
}

func Println(err error) {
	if err != nil {
		switch e := err.(type) {
		case LockError:
			fmt.Println(e.GetCode(), e.Error())
		default:
			fmt.Println(e.Error())
		}
	}
	return
}

// 测试释放别人的锁
func TestUnlock(t *testing.T) {
	// prepare
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       10,
	})
	_, err := redisClient.Ping(context.TODO()).Result()
	if err != nil {
		t.Fatal(err)
	}

	_, cancelFunc := context.WithCancel(context.TODO())
	key := "lockKey"
	_, err = UnLock(context.TODO(), redisClient, cancelFunc, key, leaseId)
	Println(err)

	time.Sleep(5 * time.Second)
}
