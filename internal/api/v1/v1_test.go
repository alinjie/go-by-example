package v1_test

import (
	"github.com/stretchr/testify/assert"

	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/alinjie/go-by-example/internal/api/v1"
	"github.com/alinjie/go-by-example/internal/mocks"
)

func TestListDog(t *testing.T) {
	s := mocks.NewMockDogStorage(t)

	ts := httptest.NewServer(v1.Server(s))

	defer ts.Close()

	dogs := []v1.DogDTO{
		{
			ID:   1,
			Name: "Halvor",
		},
	}

	s.On("ListDogs").Return(dogs, nil)

	r, err := http.Get(ts.URL + "/api/v1/dogs")

	if err != nil {
		t.Fatal(err)
	}

	rsp, err := io.ReadAll(r.Body)
	assert.NoError(t, err)

	assert.Equal(t, string(rsp), "[{\"id\":1,\"name\":\"Halvor\"}]")
}
func TestGetDog(t *testing.T) {
	s := mocks.NewMockDogStorage(t)

	dog := v1.DogDTO{
		ID:   1,
		Name: "Halvor",
	}

	s.On("GetDog", 1).Return(dog, nil)

	ts := httptest.NewServer(v1.Server(s))

	defer ts.Close()

	r, err := http.Get(ts.URL + "/api/v1/dogs/1")

	if err != nil {
		t.Fatal(err)
	}

	rsp, err := io.ReadAll(r.Body)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "{\"id\":1,\"name\":\"Halvor\"}", string(rsp))
}
