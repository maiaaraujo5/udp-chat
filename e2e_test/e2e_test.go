package e2e_test

import (
	"context"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/fx/module/runner"
	"github.com/stretchr/testify/suite"
	"net"
	"sync"
	"testing"
	"time"
)

type e2eTestSuite struct {
	suite.Suite
	redis *redis.Client
	conn  *net.UDPConn
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, new(e2eTestSuite))
}

func (s *e2eTestSuite) SetupSuite() {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	err = mr.StartAddr("127.0.0.1:6379")
	if err != nil {
		panic(err)
	}

	s.redis = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})

	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		defer wg.Done()
		err := runner.Run()
		if err != nil {
			panic(err)
		}
	}()

	wg.Wait()
	time.Sleep(1 * time.Second)
}

func (s *e2eTestSuite) SetupTest() {
	err := s.redis.FlushDB(context.Background()).Err()
	s.Assert().NoError(err)

	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:3000")
	s.Assert().NoError(err)

	conn, err := net.DialUDP("udp", nil, addr)
	s.Assert().NoError(err)

	s.conn = conn
}

func (s *e2eTestSuite) Test_EndToEnd_New_Connection_Without_Old_Messages() {
	_, err := s.conn.Write([]byte("{\"id\":\"123\",\"action\":\"NEW_CONNECTION\",\"message\":\"join the room\"}"))
	s.Assert().NoError(err)

	time.Sleep(10 * time.Millisecond)

	values, err := s.redis.LRange(context.Background(), "messages", 0, -1).Result()
	s.Assert().NoError(err)

	s.Assert().Len(values, 1)
	s.Assert().Equal(fmt.Sprintf("123-%s-join the room", s.conn.LocalAddr().String()), values[0])
}

func (s *e2eTestSuite) Test_EndToEnd_New_Connection_With_Old_Messages() {
	err := s.redis.RPush(s.redis.Context(), "messages", []string{"122-127.0.0.1:5315-Hello"}).Err()
	s.Assert().NoError(err)

	time.Sleep(10 * time.Millisecond)

	_, err = s.conn.Write([]byte("{\"id\":\"123\",\"action\":\"NEW_CONNECTION\",\"message\":\"join the room\"}"))
	s.Assert().NoError(err)

	b := make([]byte, 57)
	_, err = s.conn.Read(b)
	s.Assert().NoError(err)
	str := string(b)
	time.Sleep(10 * time.Millisecond)

	values, err := s.redis.LRange(context.Background(), "messages", 0, -1).Result()
	s.Assert().NoError(err)

	s.Assert().Equal("{\"id\":\"122\",\"user_id\":\"127.0.0.1:5315\",\"message\":\"Hello\"}", fmt.Sprint(str))
	s.Assert().Len(values, 2)
	s.Assert().Equal("122-127.0.0.1:5315-Hello", values[0])
	s.Assert().Equal(fmt.Sprintf("123-%s-join the room", s.conn.LocalAddr().String()), values[1])
}

func (s *e2eTestSuite) Test_EndToEnd_New_Receive_Message() {
	_, err := s.conn.Write([]byte("{\"id\":\"123\",\"action\":\"NEW_MESSAGE\",\"message\":\"Hello\"}"))
	s.Assert().NoError(err)

	time.Sleep(10 * time.Millisecond)

	values, err := s.redis.LRange(context.Background(), "messages", 0, -1).Result()
	s.Assert().NoError(err)

	s.Assert().Len(values, 1)
	s.Assert().Equal(fmt.Sprintf("123-%s-Hello", s.conn.LocalAddr().String()), values[0])
}

func (s *e2eTestSuite) Test_EndToEnd_Delete_Message() {
	err := s.redis.RPush(s.redis.Context(), "messages", []string{"122-127.0.0.1:5315-Hello",
		fmt.Sprintf("123-%s-Hey", s.conn.LocalAddr().String())}).Err()

	s.Assert().NoError(err)

	_, err = s.conn.Write([]byte("{\"action\":\"DELETE_MESSAGE\",\"message\":\"123\"}"))
	s.Assert().NoError(err)

	time.Sleep(10 * time.Millisecond)

	values, err := s.redis.LRange(context.Background(), "messages", 0, -1).Result()
	s.Assert().NoError(err)

	s.Assert().Len(values, 1)
	s.Assert().Equal("122-127.0.0.1:5315-Hello", values[0])
}
