package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gerokkos/clerk/api/models"
	"github.com/gorilla/mux"
	"github.com/nbio/st"
	"gopkg.in/go-playground/assert.v1"
	"gopkg.in/h2non/gock.v1"
)

func TestCreateUser(t *testing.T) {

	samples := []struct {
		inputJSON  string
		first      string
		last       string
		email      string
		cell       string
		picture    string
		registered time.Time
	}{
		{
			inputJSON: `[{"id":11454,"name":{"first":"Kara","last":"Bootsman"},"email":"kara.bootsman@example.com","cell":"(179)-338-5378","picture":{"medium":"https://randomuser.me/api/portraits/med/women/65.jpg"},"registered":{"date":"2019-09-25T21:34:11.424Z"}}]`,
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("GET", "/clerks?limit=1", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.Clerks)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Body.String(), v.inputJSON)

		assert.Equal(t, responseMap["nickname"], v.first)
		assert.Equal(t, responseMap["email"], v.email)

		assert.Equal(t, responseMap["error"], v.last)

	}
}

func TestGetUsers(t *testing.T) {

	req, err := http.NewRequest("GET", "/clerks", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Clerks)
	handler.ServeHTTP(rr, req)

	var users []models.User
	err = json.Unmarshal([]byte(rr.Body.String()), &users)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(users), 10)
}

func TestGetUserBy(t *testing.T) {

	userSample := []struct {
		id           string
		statusCode   int
		nickname     string
		email        string
		errorMessage string
	}{
		{
			nickname: "name",
			email:    "tsd",
		},
		{
			id:         "unknwon",
			statusCode: 400,
		},
	}
	for _, v := range userSample {

		req, err := http.NewRequest("GET", "/users", nil)
		if err != nil {
			t.Errorf("This is the error: %v\n", err)
		}
		req = mux.SetURLVars(req, map[string]string{})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.Clerks)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			log.Fatalf("Cannot convert to json: %v", err)
		}

		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 200 {
			assert.Equal(t, rr.Code, http.StatusOK)
			assert.Equal(t, len(responseMap), 2)
		}
	}
}

func TestClient(t *testing.T) {
	defer gock.Off()

	gock.New("http://foo.com").
		Reply(200).
		BodyString("foo foo")

	req, err := http.NewRequest("GET", "http://foo.com", nil)
	client := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(client)

	res, err := client.Do(req)
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode, 200)
	body, _ := ioutil.ReadAll(res.Body)
	st.Expect(t, string(body), "foo foso")

	// Verify that we don't have pending mocks
	st.Expect(t, gock.IsDone(), true)
}
