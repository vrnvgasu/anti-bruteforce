//go:build integration

package integration

import (
	"context"
	"fmt"
	"time"
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

func (s *SuiteIntegrationTest) TestBlockByLoginSuccess() {
	s.redisIsEmpty()

	n := s.rateLogin

	for i := 1; i < n; i++ {
		checkResponse, err := s.grpcClient.Check(context.TODO(), &pb.CheckRequest{
			Login:    loginDefault,
			Password: fmt.Sprintf("pass_%d", i),
			Ip:       fmt.Sprintf("ip_%d", i),
		})
		s.NoError(err)
		s.True(checkResponse.Result)
	}

	checkResponse, err := s.grpcClient.Check(context.TODO(), &pb.CheckRequest{
		Login:    loginDefault,
		Password: fmt.Sprintf("pass_%d", n+1),
		Ip:       fmt.Sprintf("ip_%d", n+1),
	})
	s.NoError(err)
	s.False(checkResponse.Result)

	// общее кол-во бакетов в редисе
	bucketsCount, err := s.redisClient.Keys(context.TODO(), "*").Result()
	s.NoError(err)
	// бакетов логина - 1 (проверяем первым) - упали на проверке n
	// бакетов пароля - n-1 (проверяем вторым)
	// бакетов ip - n-1 (проверяем третьим)
	s.Len(bucketsCount, 1+(n-1)+(n-1))

	s.checkLoginDefaultBucketCounter(n) // дошли до лимита
	s.checkPasswordDefaultBucketCounter(0)
	s.checkIPDefaultBucketCounter(0)

	s.checkAfterUnblockByRateLimit()
}

func (s *SuiteIntegrationTest) TestBlockByPasswordSuccess() {
	s.redisIsEmpty()

	n := s.ratePassword

	for i := 1; i < n; i++ {
		checkResponse, err := s.grpcClient.Check(context.TODO(), &pb.CheckRequest{
			Login:    fmt.Sprintf("login_%d", i),
			Password: passwordDefault,
			Ip:       fmt.Sprintf("ip_%d", i),
		})
		s.NoError(err)
		s.True(checkResponse.Result)
	}

	checkResponse, err := s.grpcClient.Check(context.TODO(), &pb.CheckRequest{
		Login:    fmt.Sprintf("login_%d", n+1),
		Password: passwordDefault,
		Ip:       fmt.Sprintf("ip_%d", n+1),
	})
	s.NoError(err)
	s.False(checkResponse.Result)

	// общее кол-во бакетов в редисе
	bucketsCount, err := s.redisClient.Keys(context.TODO(), "*").Result()
	s.NoError(err)
	// бакетов логина - n (проверяем первым)
	// бакетов пароля - 1 (проверяем вторым) - упали на проверке n
	// бакетов ip - n-1 (проверяем третьим)
	s.Len(bucketsCount, n+1+(n-1))

	s.checkLoginDefaultBucketCounter(0)
	s.checkPasswordDefaultBucketCounter(n) // дошли до лимита
	s.checkIPDefaultBucketCounter(0)

	s.checkAfterUnblockByRateLimit()
}

func (s *SuiteIntegrationTest) TestBlockByIPSuccess() {
	s.redisIsEmpty()

	n := s.rateIP

	for i := 1; i < n; i++ {
		checkResponse, err := s.grpcClient.Check(context.TODO(), &pb.CheckRequest{
			Login:    fmt.Sprintf("login_%d", i),
			Password: fmt.Sprintf("password_%d", i),
			Ip:       ipDefault,
		})
		s.NoError(err)
		s.True(checkResponse.Result)
	}

	checkResponse, err := s.grpcClient.Check(context.TODO(), &pb.CheckRequest{
		Login:    fmt.Sprintf("login_%d", n+1),
		Password: fmt.Sprintf("password_%d", n+1),
		Ip:       ipDefault,
	})
	s.NoError(err)
	s.False(checkResponse.Result)

	// общее кол-во бакетов в редисе
	bucketsCount, err := s.redisClient.Keys(context.TODO(), "*").Result()
	s.NoError(err)
	// бакетов логина - n (проверяем первым)
	// бакетов пароля - n (проверяем вторым)
	// бакетов ip - 1 (проверяем третьим) - упали на проверке n
	s.Len(bucketsCount, n+n+1)

	s.checkLoginDefaultBucketCounter(0)
	s.checkPasswordDefaultBucketCounter(0)
	s.checkIPDefaultBucketCounter(n) // дошли до лимита

	s.checkAfterUnblockByRateLimit()
}

func (s *SuiteIntegrationTest) checkAfterUnblockByRateLimit() {
	time.Sleep(time.Duration(s.rateTimeLimit)*time.Second + time.Millisecond)
	s.redisIsEmpty()

	checkResponse, err := s.grpcClient.Check(context.TODO(), &pb.CheckRequest{
		Login:    loginDefault,
		Password: passwordDefault,
		Ip:       ipDefault,
	})
	s.NoError(err)
	s.True(checkResponse.Result)
}
