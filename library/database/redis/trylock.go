package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v8"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/11/25 17:48
 * @Desc:
 */

const (
	InternalCodeFailed     = 999 // 未取到锁
	InternalCodeErr        = 998 // 连接异常
	InternalCodeLockNotOwn = 997 // 不是自己的锁
	InternalCodeExLtLease  = 996 // 有效期必须大于续租+minInterval时间
)

const minInterval = 5 // 心跳和有效期最小间隔

type LockError struct {
	Code int
	Msg  string
}

func (le LockError) Error() string {
	return le.Msg
}

func (le LockError) GetCode() int {
	return le.Code
}

// TryLock 尝试获取锁
func TryLock(ctx context.Context, conn *redis.Client, key string, ex, lease int, leaseId snowflake.ID) (context.CancelFunc, error) {
	if ex < lease+minInterval {
		return nil, LockError{Code: InternalCodeExLtLease,
			Msg: fmt.Sprintf("参数lease 必须 小于 参数ex %d秒以上", minInterval)}
	}
	ss := conn.Do(ctx, "SET", key, leaseId.String(), "EX", ex, "NX").Err()
	if ss != nil {
		return nil, LockError{Code: InternalCodeFailed, Msg: "未取到锁"}
	}
	ctxCancel, cancel := context.WithCancel(ctx)
	go func(c context.Context, k string) {
		t := time.NewTicker(time.Duration(lease) * time.Second)
		for {
			select {
			case <-t.C:
				err := conn.Do(ctx, "EXPIRE", k, ex).Err()
				if err != nil {
					// 网络等原因，心跳失败，等下次心跳
				}
			case <-c.Done():
				// 心跳结束
				return
			}
		}

	}(ctxCancel, key)
	return cancel, nil
}

// 释放锁
func delLockKey(ctx context.Context, conn *redis.Client, key string, leaseId snowflake.ID) (bool, error) {
	s, err := conn.Do(ctx, "GET", key).Int64()
	if err != nil {
		return false, LockError{Code: InternalCodeErr, Msg: err.Error()}
	}
	if snowflake.ID(s) != leaseId {
		return false, LockError{Code: InternalCodeLockNotOwn, Msg: "锁释放失败:不是自己的锁"}
	}
	b, err := conn.Do(ctx, "DEL", key).Bool()
	if err != nil {
		return false, LockError{Code: InternalCodeErr, Msg: err.Error()}
	}
	return b, nil
}

// UnLock 取消分布式锁
func UnLock(ctx context.Context, conn *redis.Client, cancelFunc context.CancelFunc, key string, leaseId snowflake.ID) (bool, error) {
	cancelFunc()
	return delLockKey(ctx, conn, key, leaseId)
}
