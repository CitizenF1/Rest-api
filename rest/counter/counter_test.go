package counter

// var redisServer *miniredis.Miniredis

// func Test(m *testing.M) {
// 	// setup()
// 	// defer redisServer.Close()
// 	mr, err := miniredis.Run()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	RedisConnector = redis.NewClient(&redis.Options{
// 		Addr: mr.Addr(),
// 	})
// 	code := m.Run()
// 	os.Exit(code)
// 	//add := counteradd("10")
// 	//assert.Equal(t, 10, add)
// }

// func TestCounterAdd(m *testing.M) {
// 	exp := time.Duration(0)
// 	mock := redismock.NewNiceMock(RedisConnector)
// 	mock.On("set", "key", "val", exp).Return(redis.NewStatusResult("", nil))
// }

// func setup() {
// 	redisServer = mockRedis()
// 	RedisConnector = redis.NewClient(&redis.Options{
// 		Addr: redisServer.Addr(),
// 	})
// }

// func mockRedis() *miniredis.Miniredis {
// 	s, err := miniredis.Run()

// 	if err != nil {
// 		panic(err)
// 	}

// 	return s
// }
