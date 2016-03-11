package watch

import (
	"github.com/otsimo/api/apipb"
	redis "gopkg.in/redis.v3"
	"encoding/base64"
	"github.com/Sirupsen/logrus"
)

const (
	channelName = "AnalytisEvent"
)

type RedisClient struct {
	client *redis.Client
	pubsub *redis.PubSub
}

func NewRedisClient(config *Config) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
	ps, err := client.Subscribe(channelName)

	if err != nil {
		return nil, err
	}
	return &RedisClient{
		client: client,
		pubsub:ps,
	}, nil
}

func (r *RedisClient) Emit(in *apipb.EmitRequest) {
	data, err := in.Marshal()
	if err != nil {
		logrus.Errorf("failed to marshall request error=%+v", err)
		return
	}
	if err := r.client.Publish(channelName, base64.StdEncoding.EncodeToString(data)).Err(); err != nil {
		logrus.Errorf("failed to publish emitrequest, %+v", err)
	}
}

func (r *RedisClient) Receive() {
	for {
		mes, err := r.pubsub.ReceiveMessage()
		if err != nil {
			logrus.Errorf("Failed to receive message, error=%+v", err)
			continue
		}
		data, err := base64.StdEncoding.DecodeString(mes.String())
		if err != nil {
			logrus.Errorf("Failed to receive message, error=%+v", err)
			continue
		}
		req := apipb.EmitRequest{}
		err = req.Unmarshal(data)
		if err != nil {
			logrus.Errorf("Failed to Unmarshal byte data")
			continue
		}
		logrus.Debugf("broadcast emit request come from redis %+v", req)
		h.broadcast <- &req
	}
}