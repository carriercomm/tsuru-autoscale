// Copyright 2015 tsuru-autoscale authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package datasource

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/tsuru/tsuru-autoscale/db"
	"github.com/tsuru/tsuru/db/dbtest"
	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type S struct {
	conn *db.Storage
}

func (s *S) SetUpSuite(c *check.C) {
	err := os.Setenv("MONGODB_DATABASE_NAME", "tsuru_autoscale_datasource")
	c.Assert(err, check.IsNil)
	s.conn, err = db.Conn()
	c.Assert(err, check.IsNil)
}

func (s *S) TearDownTest(c *check.C) {
	dbtest.ClearAllCollections(s.conn.DataSources().Database)
}

func (s *S) TearDownSuite(c *check.C) {
	err := os.Unsetenv("MONGODB_DATABASE_NAME")
	c.Assert(err, check.IsNil)
}

var _ = check.Suite(&S{})

type testHandler struct{}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	w.Write(body)
}

func (s *S) TestHttpDataSourceGet(c *check.C) {
	h := testHandler{}
	ts := httptest.NewServer(&h)
	defer ts.Close()
	ds := DataSource{Method: "POST", URL: ts.URL, Body: `{"Name": "{app}"}`}
	type dataType struct {
		Name string
	}
	data := dataType{}
	result, err := ds.Get("Paul")
	c.Assert(err, check.IsNil)
	err = json.Unmarshal([]byte(result), &data)
	c.Assert(err, check.IsNil)
	c.Assert(data.Name, check.Equals, "Paul")
}

func (s *S) TestNew(c *check.C) {
	dsConfigTests := []struct {
		conf *DataSource
		err  error
	}{
		{&DataSource{URL: "http://tsuru.io", Method: "GET"}, nil},
		{&DataSource{URL: "http://tsuru.io"}, errors.New("datasource: method required")},
		{&DataSource{Method: ""}, errors.New("datasource: url required")},
	}
	for _, tt := range dsConfigTests {
		err := New(tt.conf)
		c.Check(err, check.DeepEquals, tt.err)
	}
}

func (s *S) TestGet(c *check.C) {
	ds := DataSource{
		Name:    "xpto",
		Headers: nil,
	}
	s.conn.DataSources().Insert(&ds)
	instance, err := Get(ds.Name)
	c.Assert(err, check.IsNil)
	c.Assert(instance.Name, check.Equals, ds.Name)
}

func (s *S) TestAll(c *check.C) {
	ds := DataSource{
		Name:    "xpto",
		Headers: nil,
	}
	s.conn.DataSources().Insert(&ds)
	ds = DataSource{
		Name:    "xpto2",
		Headers: nil,
	}
	s.conn.DataSources().Insert(&ds)
	all, err := All()
	c.Assert(err, check.IsNil)
	c.Assert(all, check.HasLen, 2)
}

func (s *S) TestRemove(c *check.C) {
	ds := DataSource{
		Name:    "xpto",
		Headers: map[string]string{},
	}
	s.conn.DataSources().Insert(&ds)
	err := Remove(&ds)
	c.Assert(err, check.IsNil)
	_, err = Get(ds.Name)
	c.Assert(err, check.NotNil)
}
