package api

import (
	"github.com/ravelo-systematic-solutions/fwork/exceptions"
	"testing"
)

func TestGeneratePrivateKey_success(t *testing.T) {
	//given
	//when
	priv, err := GeneratePrivateKey(1024)

	//then
	if err != nil {
		t.Errorf("GeneratePrivateKey(), failed: %v", err)
	}

	if err := priv.Validate(); err != nil {
		t.Errorf("GeneratePrivateKey() validation failed: %s", err)
	}
}

func TestGeneratePrivateKey_error(t *testing.T) {
	//given
	//when
	_, err := GeneratePrivateKey(0)

	//then
	e := err.(*exceptions.Exception)

	if e.Code != exceptions.ResourceNotGeneratedCode {
		t.Errorf(
			"GeneratePrivateKey() invalid code:  got %s but want %s",
			e.Code,
			exceptions.ResourceNotGeneratedCode,
		)
	}

	if e.Message != exceptions.ResourceNotGeneratedMessage {
		t.Errorf(
			"GeneratePrivateKey() invalid code:  got %s but want %s",
			e.Message,
			exceptions.ResourceNotGeneratedMessage,
		)
	}
}
