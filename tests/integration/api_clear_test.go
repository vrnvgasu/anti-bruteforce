//go:build integration

package integration

import (
	"context"
	"vrnvgasu/anti-bruteforce/tests/grpc/pb"
)

func (s *SuiteIntegrationTest) TestClear() {
	s.redisIsEmpty()

	_, err := s.redisClient.Set(context.TODO(), loginDefault, 10, 0).Result()
	s.NoError(err)
	_, err = s.redisClient.Set(context.TODO(), passwordDefault, 10, 0).Result()
	s.NoError(err)
	_, err = s.redisClient.Set(context.TODO(), ipDefault, 10, 0).Result()
	s.NoError(err)

	s.checkLoginDefaultBucketCounter(10)
	s.checkPasswordDefaultBucketCounter(10)
	s.checkIPDefaultBucketCounter(10)

	res, err := s.grpcClient.Clear(context.TODO(), &pb.ClearRequest{
		Login: loginDefault,
		Ip:    ipDefault,
	})
	s.NoError(err)
	s.NotNil(res)

	bucketsCount, err := s.redisClient.Keys(context.TODO(), "*").Result()
	s.NoError(err)
	s.Len(bucketsCount, 1)

	s.checkLoginDefaultBucketCounter(0)
	s.checkPasswordDefaultBucketCounter(10)
	s.checkIPDefaultBucketCounter(0)
}
