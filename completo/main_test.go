package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	os.Exit(m.Run())
}

func TestGetPeople(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name   string
		args   args
		resp   string
		status int
	}{
		{
			name: "test getall",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("DELETE", "http:localhost:8000", nil),
			},
			resp:   `[{"id":"1","firstname":"John","lastname":"Doe","address":{"city":"City X","state":"State X"}},{"id":"2","firstname":"Koko","lastname":"Doe","address":{"city":"City Z","state":"State Y"}}]`,
			status: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetPeople(tt.args.w, tt.args.r)

			resp := tt.args.w.Result()
			body, _ := ioutil.ReadAll(resp.Body)

			if tt.status != resp.StatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					resp.StatusCode, tt.status)
			}

			if strings.TrimRight(string(body), "\n") != strings.TrimRight(tt.resp, "\n") {
				t.Errorf("handler returned unexpected body: got %v want %v",
					string(body), tt.resp)
			}
		})
	}
}

func TestAllHandler(t *testing.T) {
	type args struct {
		w      *httptest.ResponseRecorder
		r      *http.Request
		param  map[string]string
		hander http.HandlerFunc
	}
	tests := []struct {
		name   string
		args   args
		resp   string
		status int
	}{
		{
			name: "test GetPeople",
			args: args{
				w:      httptest.NewRecorder(),
				r:      httptest.NewRequest("GET", "http:localhost:8000", nil),
				hander: GetPeople,
				param:  make(map[string]string),
			},
			resp:   `[{"id":"1","firstname":"John","lastname":"Doe","address":{"city":"City X","state":"State X"}},{"id":"2","firstname":"Koko","lastname":"Doe","address":{"city":"City Z","state":"State Y"}}]`,
			status: http.StatusOK,
		},
		{
			name: "test GetPerson",
			args: args{
				w:      httptest.NewRecorder(),
				r:      httptest.NewRequest("GET", "/contato/1", nil),
				hander: GetPerson,
				param: map[string]string{
					"id": "1",
				},
			},
			resp:   `{"id":"1","firstname":"John","lastname":"Doe","address":{"city":"City X","state":"State X"}}`,
			status: http.StatusOK,
		},
		{
			name: "test GetPerson 2",
			args: args{
				w:      httptest.NewRecorder(),
				r:      httptest.NewRequest("GET", "/contato/2", nil),
				hander: GetPerson,
				param: map[string]string{
					"id": "2",
				},
			},
			resp:   `{"id":"2","firstname":"Koko","lastname":"Doe","address":{"city":"City Z","state":"State Y"}}`,
			status: http.StatusOK,
		},
		{
			name: "test delete 1",
			args: args{
				w:      httptest.NewRecorder(),
				r:      httptest.NewRequest("GET", "/contato/2", nil),
				hander: DeletePerson,
				param: map[string]string{
					"id": "1",
				},
			},
			resp:   `[{"id":"2","firstname":"Koko","lastname":"Doe","address":{"city":"City Z","state":"State Y"}}]`,
			status: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := mux.SetURLVars(tt.args.r, tt.args.param)
			tt.args.hander(tt.args.w, r)

			resp := tt.args.w.Result()
			body, _ := ioutil.ReadAll(resp.Body)

			if tt.status != resp.StatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					resp.StatusCode, tt.status)
			}
			if strings.TrimRight(string(body), "\n") != strings.TrimRight(tt.resp, "\n") {
				t.Errorf("handler returned unexpected body: got %v want %v",
					string(body), tt.resp)
			}
		})
	}
}
