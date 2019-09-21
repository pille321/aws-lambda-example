package main

import "testing"

func TestHandleLambdaEvent(t *testing.T) {

    event := MyEvent {
      Name: "Bob",
      Age: 56,
    }

    response := MyResponse {
      Message: "Bob is 56 years old!",
    }


    got, _ := HandleLambdaEvent(event)

    if got != response {
		t.Errorf("HandleLambdaEvent(%q) == %q, want %q", event, got, response)
	}


}
