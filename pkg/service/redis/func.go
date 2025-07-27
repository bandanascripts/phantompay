package redis

import (
	"context"
	"fmt"
	"time"
)

func SetToRedis(ctx context.Context, key, value string, ttls int) error {

	if err := GlobalRedisClient.Set(ctx, key, value, time.Duration(ttls)*time.Second).Err(); err != nil {
		return fmt.Errorf("failed to set data to redis : %w", err)
	}

	return nil
}

func GetFromRedis(ctx context.Context, key string) (string, error) {

	result, err := GlobalRedisClient.Get(ctx, key).Result()

	if err != nil {
		return "", fmt.Errorf("failed to get data from redis : %w", err)
	}

	return result, nil
}
