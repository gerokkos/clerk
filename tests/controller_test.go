package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gerokkos/clerk/api/controllers"
)

//function to compare two JSONs to be used in tests
func AreEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}
	return reflect.DeepEqual(o1, o2), nil
}

//seeding pre-test, the database with a mocked request from RandomUser API, with 15 users, saved on jsonkeeper
func TestSeedWithMockData(t *testing.T) {
	req, err := http.NewRequest("POST", "/populate?url=https://jsonkeeper.com/b/TXN5", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Populate)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//Test without optional params, we are getting the 10 users by desc registration date
func TestGetWithoutParams(t *testing.T) {
	req, err := http.NewRequest("GET", "/clerks", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Clerks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	data, err := ioutil.ReadFile("./user.json")
	datastring := string(data)
	if err != nil {
		fmt.Print(err)
	}
	areEqual, err := AreEqualJSON(datastring, rr.Body.String())

	if areEqual == false {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), datastring)
	}
}

//Optional param limit set as 1
func TestGetWithLimitEquals1(t *testing.T) {
	req, err := http.NewRequest("GET", "/clerks/?limit=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Clerks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"id":12,"name":{"first":"Jennifer","last":"Bishop"},"email":"jennifer.bishop@example.com","cell":"0765-442-350","picture":{"medium":"https://randomuser.me/api/portraits/med/women/2.jpg"},"registered":{"date":"2018-03-01T00:00:00Z"}}]`

	areEqual, err := AreEqualJSON(expected, rr.Body.String())

	if areEqual == false {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//limit set as 12
func TestGetWithLimitEquals12(t *testing.T) {
	req, err := http.NewRequest("GET", "/clerks?limit=12", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Clerks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"id":27,"name":{"first":"Jennifer","last":"Bishop"},"email":"jennifer.bishop@example.com","cell":"0765-442-350","picture":{"medium":"https://randomuser.me/api/portraits/med/women/2.jpg"},"registered":{"date":"2018-03-01T00:00:00Z"}},{"id":12,"name":{"first":"Jennifer","last":"Bishop"},"email":"jennifer.bishop@example.com","cell":"0765-442-350","picture":{"medium":"https://randomuser.me/api/portraits/med/women/2.jpg"},"registered":{"date":"2018-03-01T00:00:00Z"}},{"id":11,"name":{"first":"Jim","last":"Dunn"},"email":"jim.dunn@example.com","cell":"0467-825-071","picture":{"medium":"https://randomuser.me/api/portraits/med/men/54.jpg"},"registered":{"date":"2017-03-12T00:00:00Z"}},{"id":26,"name":{"first":"Jim","last":"Dunn"},"email":"jim.dunn@example.com","cell":"0467-825-071","picture":{"medium":"https://randomuser.me/api/portraits/med/men/54.jpg"},"registered":{"date":"2017-03-12T00:00:00Z"}},{"id":14,"name":{"first":"Alfredo","last":"Van Hoften"},"email":"alfredo.vanhoften@example.com","cell":"(847)-901-0801","picture":{"medium":"https://randomuser.me/api/portraits/med/men/90.jpg"},"registered":{"date":"2016-11-01T00:00:00Z"}},{"id":29,"name":{"first":"Alfredo","last":"Van Hoften"},"email":"alfredo.vanhoften@example.com","cell":"(847)-901-0801","picture":{"medium":"https://randomuser.me/api/portraits/med/men/90.jpg"},"registered":{"date":"2016-11-01T00:00:00Z"}},{"id":3,"name":{"first":"Artur","last":"Enger"},"email":"artur.enger@example.com","cell":"44172889","picture":{"medium":"https://randomuser.me/api/portraits/med/men/49.jpg"},"registered":{"date":"2016-09-22T00:00:00Z"}},{"id":18,"name":{"first":"Artur","last":"Enger"},"email":"artur.enger@example.com","cell":"44172889","picture":{"medium":"https://randomuser.me/api/portraits/med/men/49.jpg"},"registered":{"date":"2016-09-22T00:00:00Z"}},{"id":15,"name":{"first":"Mandy","last":"Gomez"},"email":"mandy.gomez@example.com","cell":"0787-709-991","picture":{"medium":"https://randomuser.me/api/portraits/med/women/59.jpg"},"registered":{"date":"2015-12-31T00:00:00Z"}},{"id":30,"name":{"first":"Mandy","last":"Gomez"},"email":"mandy.gomez@example.com","cell":"0787-709-991","picture":{"medium":"https://randomuser.me/api/portraits/med/women/59.jpg"},"registered":{"date":"2015-12-31T00:00:00Z"}},{"id":7,"name":{"first":"Solene","last":"Almeida"},"email":"solene.almeida@example.com","cell":"(91) 2095-8615","picture":{"medium":"https://randomuser.me/api/portraits/med/women/25.jpg"},"registered":{"date":"2015-03-14T00:00:00Z"}},{"id":22,"name":{"first":"Solene","last":"Almeida"},"email":"solene.almeida@example.com","cell":"(91) 2095-8615","picture":{"medium":"https://randomuser.me/api/portraits/med/women/25.jpg"},"registered":{"date":"2015-03-14T00:00:00Z"}}]`

	areEqual, err := AreEqualJSON(expected, rr.Body.String())

	if areEqual == false {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//setting the param starting after user id
func TestGetStartingAfter(t *testing.T) {
	req, err := http.NewRequest("GET", "/clerks?starting_after=5", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Clerks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"id":12,"name":{"first":"Jennifer","last":"Bishop"},"email":"jennifer.bishop@example.com","cell":"0765-442-350","picture":{"medium":"https://randomuser.me/api/portraits/med/women/2.jpg"},"registered":{"date":"2018-03-01T00:00:00Z"}},{"id":11,"name":{"first":"Jim","last":"Dunn"},"email":"jim.dunn@example.com","cell":"0467-825-071","picture":{"medium":"https://randomuser.me/api/portraits/med/men/54.jpg"},"registered":{"date":"2017-03-12T00:00:00Z"}},{"id":14,"name":{"first":"Alfredo","last":"Van Hoften"},"email":"alfredo.vanhoften@example.com","cell":"(847)-901-0801","picture":{"medium":"https://randomuser.me/api/portraits/med/men/90.jpg"},"registered":{"date":"2016-11-01T00:00:00Z"}},{"id":15,"name":{"first":"Mandy","last":"Gomez"},"email":"mandy.gomez@example.com","cell":"0787-709-991","picture":{"medium":"https://randomuser.me/api/portraits/med/women/59.jpg"},"registered":{"date":"2015-12-31T00:00:00Z"}},{"id":7,"name":{"first":"Solene","last":"Almeida"},"email":"solene.almeida@example.com","cell":"(91) 2095-8615","picture":{"medium":"https://randomuser.me/api/portraits/med/women/25.jpg"},"registered":{"date":"2015-03-14T00:00:00Z"}},{"id":8,"name":{"first":"Noah","last":"Mortensen"},"email":"noah.mortensen@example.com","cell":"17264150","picture":{"medium":"https://randomuser.me/api/portraits/med/men/26.jpg"},"registered":{"date":"2014-12-18T00:00:00Z"}},{"id":9,"name":{"first":"Daniela","last":"Delgado"},"email":"daniela.delgado@example.com","cell":"611-309-687","picture":{"medium":"https://randomuser.me/api/portraits/med/women/77.jpg"},"registered":{"date":"2014-07-20T00:00:00Z"}},{"id":13,"name":{"first":"Ilyès","last":"Vidal"},"email":"ilyes.vidal@example.com","cell":"06-58-69-27-40","picture":{"medium":"https://randomuser.me/api/portraits/med/men/98.jpg"},"registered":{"date":"2012-02-26T00:00:00Z"}},{"id":10,"name":{"first":"Rosina","last":"Henry"},"email":"rosina.henry@example.com","cell":"075 685 20 64","picture":{"medium":"https://randomuser.me/api/portraits/med/women/71.jpg"},"registered":{"date":"2004-12-03T00:00:00Z"}},{"id":6,"name":{"first":"Lauri","last":"Heikkila"},"email":"lauri.heikkila@example.com","cell":"041-690-62-21","picture":{"medium":"https://randomuser.me/api/portraits/med/men/23.jpg"},"registered":{"date":"2003-09-05T00:00:00Z"}}]`

	areEqual, err := AreEqualJSON(expected, rr.Body.String())

	if areEqual == false {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//starting after among with limit
func TestGetStartingAfterWithLimit(t *testing.T) {
	req, err := http.NewRequest("GET", "/clerks?starting_after=5&limit=2", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Clerks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"id":7,"name":{"first":"Solene","last":"Almeida"},"email":"solene.almeida@example.com","cell":"(91) 2095-8615","picture":{"medium":"https://randomuser.me/api/portraits/med/women/25.jpg"},"registered":{"date":"2015-03-14T00:00:00Z"}},{"id":6,"name":{"first":"Lauri","last":"Heikkila"},"email":"lauri.heikkila@example.com","cell":"041-690-62-21","picture":{"medium":"https://randomuser.me/api/portraits/med/men/23.jpg"},"registered":{"date":"2003-09-05T00:00:00Z"}}]`

	areEqual, err := AreEqualJSON(expected, rr.Body.String())

	if areEqual == false {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//ending before user id
func TestGetEndingBefore(t *testing.T) {
	req, err := http.NewRequest("GET", "/clerks?ending_before=10", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Clerks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"id":3,"name":{"first":"Artur","last":"Enger"},"email":"artur.enger@example.com","cell":"44172889","picture":{"medium":"https://randomuser.me/api/portraits/med/men/49.jpg"},"registered":{"date":"2016-09-22T00:00:00Z"}},{"id":7,"name":{"first":"Solene","last":"Almeida"},"email":"solene.almeida@example.com","cell":"(91) 2095-8615","picture":{"medium":"https://randomuser.me/api/portraits/med/women/25.jpg"},"registered":{"date":"2015-03-14T00:00:00Z"}},{"id":8,"name":{"first":"Noah","last":"Mortensen"},"email":"noah.mortensen@example.com","cell":"17264150","picture":{"medium":"https://randomuser.me/api/portraits/med/men/26.jpg"},"registered":{"date":"2014-12-18T00:00:00Z"}},{"id":9,"name":{"first":"Daniela","last":"Delgado"},"email":"daniela.delgado@example.com","cell":"611-309-687","picture":{"medium":"https://randomuser.me/api/portraits/med/women/77.jpg"},"registered":{"date":"2014-07-20T00:00:00Z"}},{"id":4,"name":{"first":"Edda","last":"Baum"},"email":"edda.baum@example.com","cell":"0170-0193364","picture":{"medium":"https://randomuser.me/api/portraits/med/women/50.jpg"},"registered":{"date":"2010-09-12T00:00:00Z"}},{"id":1,"name":{"first":"Amely","last":"Roelen"},"email":"amely.roelen@example.com","cell":"(610)-633-4301","picture":{"medium":"https://randomuser.me/api/portraits/med/women/9.jpg"},"registered":{"date":"2006-09-01T00:00:00Z"}},{"id":5,"name":{"first":"سینا","last":"گلشن"},"email":"syn.glshn@example.com","cell":"0986-479-1395","picture":{"medium":"https://randomuser.me/api/portraits/med/men/12.jpg"},"registered":{"date":"2005-01-23T00:00:00Z"}},{"id":2,"name":{"first":"Leah","last":"Coleman"},"email":"leah.coleman@example.com","cell":"081-428-0307","picture":{"medium":"https://randomuser.me/api/portraits/med/women/1.jpg"},"registered":{"date":"2005-01-21T00:00:00Z"}},{"id":6,"name":{"first":"Lauri","last":"Heikkila"},"email":"lauri.heikkila@example.com","cell":"041-690-62-21","picture":{"medium":"https://randomuser.me/api/portraits/med/men/23.jpg"},"registered":{"date":"2003-09-05T00:00:00Z"}}]`

	areEqual, err := AreEqualJSON(expected, rr.Body.String())

	if areEqual == false {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//ending before with limit
func TestGetEndingBeforeWithLimit(t *testing.T) {
	req, err := http.NewRequest("GET", "/clerks?ending_before=10&limit=2", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Clerks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"id":8,"name":{"first":"Noah","last":"Mortensen"},"email":"noah.mortensen@example.com","cell":"17264150","picture":{"medium":"https://randomuser.me/api/portraits/med/men/26.jpg"},"registered":{"date":"2014-12-18T00:00:00Z"}},{"id":9,"name":{"first":"Daniela","last":"Delgado"},"email":"daniela.delgado@example.com","cell":"611-309-687","picture":{"medium":"https://randomuser.me/api/portraits/med/women/77.jpg"},"registered":{"date":"2014-07-20T00:00:00Z"}}]`

	areEqual, err := AreEqualJSON(expected, rr.Body.String())

	if areEqual == false {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//get user by email
func TestGetWithEmail(t *testing.T) {
	req, err := http.NewRequest("GET", "/clerks?email=edda.baum@example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Clerks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"id":4,"name":{"first":"Edda","last":"Baum"},"email":"edda.baum@example.com","cell":"0170-0193364","picture":{"medium":"https://randomuser.me/api/portraits/med/women/50.jpg"},"registered":{"date":"2010-09-12T00:00:00Z"}},{"id":19,"name":{"first":"Edda","last":"Baum"},"email":"edda.baum@example.com","cell":"0170-0193364","picture":{"medium":"https://randomuser.me/api/portraits/med/women/50.jpg"},"registered":{"date":"2010-09-12T00:00:00Z"}}]`

	areEqual, err := AreEqualJSON(expected, rr.Body.String())

	if areEqual == false {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//get user by email case insensitive
func TestGetWithEmailUpperCase(t *testing.T) {
	req, err := http.NewRequest("GET", "/clerks?email=EDDA.BAUM@EXAMPLE.COM", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Clerks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"id":4,"name":{"first":"Edda","last":"Baum"},"email":"edda.baum@example.com","cell":"0170-0193364","picture":{"medium":"https://randomuser.me/api/portraits/med/women/50.jpg"},"registered":{"date":"2010-09-12T00:00:00Z"}},{"id":19,"name":{"first":"Edda","last":"Baum"},"email":"edda.baum@example.com","cell":"0170-0193364","picture":{"medium":"https://randomuser.me/api/portraits/med/women/50.jpg"},"registered":{"date":"2010-09-12T00:00:00Z"}}]`

	areEqual, err := AreEqualJSON(expected, rr.Body.String())

	if areEqual == false {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//Delete the mock data we created in pre-tes so the db is cleared
func TestDeleteMockData(t *testing.T) {
	controllers.Load(server.DB)
}
