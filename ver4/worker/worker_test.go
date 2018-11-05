package worker

import (
	"io/ioutil"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/labstack/echo"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestGetID(t *testing.T) {
	t.Run("get id", func(t *testing.T) {
		e := echo.New()
		n := NewMessage()
		e.GET("/:id", n.GetID)

		s := httptest.NewServer(e)
		defer s.Close()

		var id int = 3

		client := s.Client()
		req, err := client.Get(s.URL + "/" + strconv.Itoa(id))
		if err != nil {
			t.Fatal(err)
		}

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			t.Fatal(err)
		}

		// got body, want query
		actualStr := string(body)
		actual, _ := strconv.Atoi(actualStr)
		if actual != id {
			t.Errorf("got: %s, want: %s", actualStr, strconv.Itoa(id))
		}
	})
}
