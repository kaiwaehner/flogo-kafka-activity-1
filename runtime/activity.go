package kafka

import (
	"github.com/TIBCOSoftware/flogo-lib/flow/activity"
	"github.com/op/go-logging"

	"github.com/optiopay/kafka"
	"github.com/optiopay/kafka/proto"
)

// log is the default package logger
var log = logging.MustGetLogger("activity-tibco-kafka")

const (
	//	ivAcctSID   = "accountSID"
	//	ivAuthToken = "authToken"
	//	ivFrom      = "from"
	//	ivTo        = "to"
	//	ivMessage   = "message"

	topic     = "topic"
	message   = "message"
	partition = 0
)

var kafkaAddrs = []string{"localhost:9092", "localhost:9093"}

// KafkaActivity is a Kafka Activity implementation
type KafkaActivity struct {
	metadata *activity.Metadata
}

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(&KafkaActivity{metadata: md})
}

// Metadata implements activity.Activity.Metadata
func (a *KafkaActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *KafkaActivity) Eval(context activity.Context) (done bool, err error) {

	//	accountSID := context.GetInput(ivAcctSID).(string)
	//	authToken := context.GetInput(ivAuthToken).(string)
	//	from := context.GetInput(ivFrom).(string)
	//	to := context.GetInput(ivTo).(string)
	//	message := context.GetInput(ivMessage).(string)

	//	twilio := gotwilio.NewTwilioClient(accountSID, authToken)

	topicInput := context.GetInput(topic).(string)

	messageInput := context.GetInput(message).(string)

	conf := kafka.NewBrokerConf("test-client")
	conf.AllowTopicCreation = true
	// connect to kafka cluster
	broker, err := kafka.Dial(kafkaAddrs, conf)
	if err != nil {
		log.Fatalf("cannot connect to kafka cluster: %s", err)
	}
	defer broker.Close()

	producer := broker.Producer(kafka.NewProducerConf())
	//input := bufio.NewReader(os.Stdin)
	//	for {
	//		line, err := input.ReadString('\n')
	//		if err != nil {
	//			log.Fatalf("input error: %s", err)
	//		}
	//		line = strings.TrimSpace(line)
	//		if line == "" {
	//			continue
	//		}

	//		msg := &proto.Message{Value: []byte(line)}
	//		if _, err := producer.Produce(topic, partition, msg); err != nil {
	//			log.Fatalf("cannot produce message to %s:%d: %s", topic, partition, err)
	//		}
	//	}

	// resp, _, err := twilio.SendSMS(from, to, message, "", "")

	msg := &proto.Message{Value: []byte(messageInput)}

	resp, err := producer.Produce(topicInput, partition, msg)

	if err != nil {
		log.Error("Error sending message to Kafka broker:", err)
	}

	if log.IsEnabledFor(logging.DEBUG) {
		log.Debug("Response:", resp)
	}

	return true, nil
}
