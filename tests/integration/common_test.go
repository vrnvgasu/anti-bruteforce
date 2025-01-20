//go:build integration

package integration

import (
	"context"
	"strconv"
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
	s.NoError(err)
	s.Equal(strconv.Itoa(counter), v)
}
