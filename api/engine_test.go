package api

import (
	"context"
	"fwork/exceptions"
	"log"
	"net/http"
	"testing"
)

func TestEngine_Run_success(t *testing.T) {
	//given
	api := &engine{
		server: http.Server{Addr: ":90000"},
		config: Config{Service: Service{
			Id:       "i1",
			Name:     "n1",
			Internal: ":90000",
			External: "http://localhost:90000",
		}},
	}

	go func(api *engine) {
		if err := api.server.Shutdown(context.TODO()); err != nil {
			log.Panicf("unable to shutdown: [err: %v]", err)
		}
	}(api)

	//when
	err := api.Run()

	//then
	if err == nil {
		t.Errorf("Run() expected resource closed error")
	}

	e := err.(*exceptions.Exception)
	if e.Code != exceptions.ResourceClosedCode {
		t.Errorf("Run() got %v but want %v",
			e.Code,
			exceptions.ResourceClosedCode,
		)
	}

	e = err.(*exceptions.Exception)
	if e.Message != exceptions.ResourceClosedMessage {
		t.Errorf("Run() got %v but want %v",
			e.Code,
			exceptions.ResourceClosedMessage,
		)
	}
}
