package kafka

import (
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/flow/activity"
	"github.com/TIBCOSoftware/flogo-lib/flow/test"
)

func TestRegistered(t *testing.T) {
	act := activity.Get("tibco-kafka")

	if act == nil {
		t.Error("Activity Not Registered")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	md := activity.NewMetadata(jsonMetadata)
	act := &KafkaActivity{metadata: md}

	tc := test.NewTestActivityContext(md)

	//setup attrs
	//	tc.SetInput(ivAcctSID, "A...9")
	//	tc.SetInput(ivAuthToken, "f...4")
	//	tc.SetInput(ivTo, "+1...")
	//	tc.SetInput(ivFrom, "+12016901385")
	//	tc.SetInput(ivMessage, "Go Flogo")

	tc.SetInput(topic, "my-messages")
	tc.SetInput(message, "test message by kai for kafka")

	act.Eval(tc)

	//check result attr

	// TODO
}
