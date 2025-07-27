package core

import (
	"context"
	"github.com/bandanascripts/phantompay/pkg/service/redis"
	twowaykey "github.com/bandanascripts/phantompay/pkg/service/two_way_key"
	"github.com/bandanascripts/phantompay/pkg/utils"
)

func GenAndStoreKey(ctx context.Context, privateKeyId, publicKey string, ttls int) error {

	pemPrivKey, pemPubKey, err := twowaykey.GenAndEncode()

	if err != nil {
		return err
	}

	activeKeyId, err := utils.NewId()

	if err != nil {
		return err
	}

	if err = redis.SetToRedis(ctx, "RSA:ACTIVEKEY", activeKeyId, ttls); err != nil {
		return err
	}

	if err = redis.SetToRedis(ctx, privateKeyId + activeKeyId, pemPrivKey, ttls); err != nil {
		return err
	}

	if err = redis.SetToRedis(ctx, publicKey + activeKeyId, pemPubKey, ttls); err != nil {
		return err
	}

	return nil
}
