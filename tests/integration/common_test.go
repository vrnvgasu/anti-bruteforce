//go:build integration

package integration

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-redis/redis/v8"
)

const (
	redisAddrDefault = "127.0.0.1:6380"
	grpcAddrDefault  = "127.0.0.1:50051"

	rateLoginDefault     = "5"
	ratePasswordDefault  = "10"
	rateIPDefault        = "15"
	rateTimeLimitDefault = "1"

	loginDefault    = "login"
	passwordDefault = "pass"
	ipDefault       = "192.168.5.1"
)

func (s *SuiteIntegrationTest) redisIsEmpty() {
	result, err := s.redisClient.Keys(context.TODO(), "*").Result()
	s.NoError(err)
	s.Empty(result)
}

func (s *SuiteIntegrationTest) checkLoginDefaultBucketCounter(counter int) {
	s.checkBucketCounter(loginDefault, counter)
}

func (s *SuiteIntegrationTest) checkPasswordDefaultBucketCounter(counter int) {
	s.checkBucketCounter(passwordDefault, counter)
}

func (s *SuiteIntegrationTest) checkIPDefaultBucketCounter(counter int) {
	s.checkBucketCounter(ipDefault, counter)
}

func (s *SuiteIntegrationTest) checkBucketCounter(key string, counter int) {
	v, err := s.redisClient.Get(context.TODO(), key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			s.Fail(err.Error())
		}

		v = "0"
	}
	s.Equal(strconv.Itoa(counter), v)
}
