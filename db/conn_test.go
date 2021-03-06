// Copyright 2015 tsuru-autoscale authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package db

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type S struct{}

var _ = check.Suite(&S{})

func (s *S) TestEvents(c *check.C) {
	strg, err := Conn()
	c.Assert(err, check.IsNil)
	event := strg.Events()
	eventc := strg.Collection("events")
	c.Assert(event, check.DeepEquals, eventc)
}

func (s *S) TestConfigs(c *check.C) {
	strg, err := Conn()
	c.Assert(err, check.IsNil)
	config := strg.Configs()
	configc := strg.Collection("configs")
	c.Assert(config, check.DeepEquals, configc)
}

func (s *S) TestInstances(c *check.C) {
	strg, err := Conn()
	c.Assert(err, check.IsNil)
	instance := strg.Instances()
	instancec := strg.Collection("instances")
	c.Assert(instance, check.DeepEquals, instancec)
}

func (s *S) TestDataSources(c *check.C) {
	strg, err := Conn()
	c.Assert(err, check.IsNil)
	datasource := strg.DataSources()
	datasourcec := strg.Collection("datasources")
	c.Assert(datasource, check.DeepEquals, datasourcec)
}

func (s *S) TestAlarms(c *check.C) {
	strg, err := Conn()
	c.Assert(err, check.IsNil)
	alarm := strg.Alarms()
	alarmc := strg.Collection("alarms")
	c.Assert(alarm, check.DeepEquals, alarmc)
}

func (s *S) TestActions(c *check.C) {
	strg, err := Conn()
	c.Assert(err, check.IsNil)
	action := strg.Actions()
	actionc := strg.Collection("actions")
	c.Assert(action, check.DeepEquals, actionc)
}
