//go:build integration

package integration

import (
	"context"
	"vrnvgasu/anti-bruteforce/tests/grpc/pb"
)

func (s *SuiteIntegrationTest) TestCheckSuccess() {
	s.redisIsEmpty()

	checkResponse, err := s.grpcClient.Check(context.TODO(), &pb.CheckRequest{
		Login:    loginDefault,
		Password: passwordDefault,
		Ip:       ipDefault,
	})
	s.NoError(err)
	s.NotNil(checkResponse)
	s.True(checkResponse.Result)

	bucketsCount, err := s.redisClient.Keys(context.TODO(), "*").Result()
	s.NoError(err)
	s.Len(bucketsCount, 3)

	s.checkLoginDefaultBucketCounter(1)
	s.checkPasswordDefaultBucketCounter(1)
	s.checkIPDefaultBucketCounter(1)
}
