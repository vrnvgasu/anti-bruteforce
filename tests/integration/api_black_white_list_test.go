//go:build integration

package integration

import (
	"context"
	"vrnvgasu/anti-bruteforce/tests/grpc/pb"
)

func (s *SuiteIntegrationTest) TestBlackList() {
	s.redisIsEmpty()

	// ADD TO BLACK LIST
	res, err := s.grpcClient.AddToList(context.TODO(), &pb.ListRequest{
		Subnet: &pb.Subnet{
			Ip:   subnetIPDefault,
			Mask: subnetMaskDefault,
		},
		Type: blackList,
	})
	s.NoError(err)
	s.NotNil(res)

	checkResponse, err := s.grpcClient.Check(context.TODO(), &pb.CheckRequest{
		Login:    loginDefault,
		Password: passwordDefault,
		Ip:       ipDefault,
	})
	s.NoError(err)
	s.NotNil(checkResponse)
	s.False(checkResponse.Result)

	// не доходим до создания бакетов
	s.checkLoginDefaultBucketCounter(0)
	s.checkPasswordDefaultBucketCounter(0)
	s.checkIPDefaultBucketCounter(0)

	// REMOVE TO BLACK LIST
	res, err = s.grpcClient.RemoveFromList(context.TODO(), &pb.ListRequest{
		Subnet: &pb.Subnet{
			Ip:   subnetIPDefault,
			Mask: subnetMaskDefault,
		},
		Type: blackList,
	})
	s.NoError(err)
	s.NotNil(res)
	checkResponse, err = s.grpcClient.Check(context.TODO(), &pb.CheckRequest{
		Login:    loginDefault,
		Password: passwordDefault,
		Ip:       ipDefault,
	})
	s.NoError(err)
	s.NotNil(checkResponse)
	s.True(checkResponse.Result)

	// создали бакеты
	s.checkLoginDefaultBucketCounter(1)
	s.checkPasswordDefaultBucketCounter(1)
	s.checkIPDefaultBucketCounter(1)
}

func (s *SuiteIntegrationTest) TestWhiteList() {
	s.redisIsEmpty()

	_, err := s.redisClient.Set(context.TODO(), loginDefault, s.rateLogin, 0).Result()
	s.NoError(err)
	_, err = s.redisClient.Set(context.TODO(), passwordDefault, s.ratePassword, 0).Result()
	s.NoError(err)
	_, err = s.redisClient.Set(context.TODO(), ipDefault, s.rateIP, 0).Result()
	s.NoError(err)

	s.checkLoginDefaultBucketCounter(s.rateLogin)
	s.checkPasswordDefaultBucketCounter(s.ratePassword)
	s.checkIPDefaultBucketCounter(s.rateIP)

	// ADD TO WHITE LIST
	res, err := s.grpcClient.AddToList(context.TODO(), &pb.ListRequest{
		Subnet: &pb.Subnet{
			Ip:   subnetIPDefault,
			Mask: subnetMaskDefault,
		},
		Type: whiteList,
	})
	s.NoError(err)
	s.NotNil(res)

	checkResponse, err := s.grpcClient.Check(context.TODO(), &pb.CheckRequest{
		Login:    loginDefault,
		Password: passwordDefault,
		Ip:       ipDefault,
	})
	s.NoError(err)
	s.NotNil(checkResponse)
	s.True(checkResponse.Result)

	// не доходим до изменения бакетов
	s.checkLoginDefaultBucketCounter(s.rateLogin)
	s.checkPasswordDefaultBucketCounter(s.ratePassword)
	s.checkIPDefaultBucketCounter(s.rateIP)

	// REMOVE TO WHITE LIST
	res, err = s.grpcClient.RemoveFromList(context.TODO(), &pb.ListRequest{
		Subnet: &pb.Subnet{
			Ip:   subnetIPDefault,
			Mask: subnetMaskDefault,
		},
		Type: whiteList,
	})
	s.NoError(err)
	s.NotNil(res)
	checkResponse, err = s.grpcClient.Check(context.TODO(), &pb.CheckRequest{
		Login:    loginDefault,
		Password: passwordDefault,
		Ip:       ipDefault,
	})
	s.NoError(err)
	s.NotNil(checkResponse)
	s.False(checkResponse.Result)

	// упали на проверке логина
	s.checkLoginDefaultBucketCounter(s.rateLogin + 1)
	s.checkPasswordDefaultBucketCounter(s.ratePassword)
	s.checkIPDefaultBucketCounter(s.rateIP)
}
