//go:build integration

package integration

import (
	"context"
	"log"
	"os"
	"strconv"
	"testing"
	"vrnvgasu/anti-bruteforce/tests/grpc/pb"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SuiteIntegrationTest struct {
	suite.Suite

	grpcClient  pb.AntiBruteforceClient
	redisClient redis.UniversalClient

	rateLogin     int
	ratePassword  int
	rateIP        int
	rateTimeLimit int
}

func TestSuit(t *testing.T) {
	suite.Run(t, &SuiteIntegrationTest{})
}

func (s *SuiteIntegrationTest) SetupSuite() {
	redisAddr, ok := os.LookupEnv("TEST_REDIS_ADDR")
	if !ok {
		redisAddr = redisAddrDefault
	}
	s.redisClient = redis.NewClient(&redis.Options{Addr: redisAddr})
	if _, err := s.redisClient.Ping(context.TODO()).Result(); err != nil {
		log.Fatal("Failed to connect to redis")
	}

	grpcAddr, ok := os.LookupEnv("TEST_GRPC_ADDR")
	if !ok {
		grpcAddr = grpcAddrDefault
	}
	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to create grpc client")
	}
	s.grpcClient = pb.NewAntiBruteforceClient(conn)

	s.rateLogin = rateValue("TEST_RATE_LOGIN", rateLoginDefault)
	s.ratePassword = rateValue("TEST_RATE_PASSWORD", ratePasswordDefault)
	s.rateIP = rateValue("TEST_RATE_IP", rateIPDefault)
	s.rateTimeLimit = rateValue("TEST_RATE_TIME_LIMIT", rateTimeLimitDefault)
}

func (s *SuiteIntegrationTest) TearDownTest() {
	_, err := s.redisClient.FlushDB(context.TODO()).Result()
	if err != nil {
		return
	}
	require.NoError(s.T(), err)
}

func rateValue(envKey string, defaultEnvValue string) int {
	envValue, ok := os.LookupEnv(envKey)
	if !ok {
		envValue = defaultEnvValue
	}
	v, err := strconv.Atoi(envValue)
	if err != nil {
		log.Fatal(err)
	}

	return v
}
