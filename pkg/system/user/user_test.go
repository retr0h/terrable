// Copyright (c) 2021 John Dewey <john@dewey.ws>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package user

import (
	"fmt"
	"os/user"
	"testing"

	"github.com/retr0h/terraform-provider-terrable/pkg/exec"
	"github.com/stretchr/testify/assert"
)

func TestLocate(t *testing.T) {
	u, _ := user.Current()

	cases := []struct {
		Name     string
		UserName string
		Err      bool
	}{
		{
			Name:     "Existing user name",
			UserName: u.Username,
			Err:      false,
		},
		{
			Name:     "Non-existing user name",
			UserName: "invalid",
			Err:      true,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d-%s", i, tc.Name), func(t *testing.T) {
			u, err := Lookup(tc.UserName)

			if tc.Err == false {
				assert.Equal(t, u.Username, tc.UserName)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}

		})
	}
}

func TestAdd(t *testing.T) {
	fc := &exec.FakeCommander{}
	fuc := &exec.FakeUnsuccessfulCommander{}
	cases := []struct {
		Name      string
		User      *User
		Want      []string
		Err       bool
		Commander exec.CommanderDelegate
	}{
		{
			Name: "All fields",
			User: &User{
				Name:      "fake-name",
				Directory: "fake-dir",
				Shell:     "fake-shell",
				Groups: []string{
					"foo",
					"bar",
				},
				System: true,
				UID:    "1099",
				GID:    "1099",
			},
			Want: []string{
				"/usr/sbin/useradd",
				"-s", "fake-shell",
				"-m",
				"-d", "fake-dir",
				"-G", "foo,bar",
				"-r",
				"-u", "1099",
				"-g", "1099",
				"fake-name",
			},
			Err:       false,
			Commander: fc,
		},
		{
			Name: "Without optional fields",
			User: &User{
				Name:  "fake-name",
				Shell: "fake-shell",
			},
			Want: []string{
				"/usr/sbin/useradd",
				"-s", "fake-shell",
				"-m",
				"-d", "/home/fake-name",
				"fake-name",
			},
			Err:       false,
			Commander: fc,
		},
		{
			Name: "Returns an error",
			User: &User{
				Name:  "fake-name",
				Shell: "fake-shell",
			},
			Want:      []string(nil),
			Err:       true,
			Commander: fuc,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d-%s", i, tc.Name), func(t *testing.T) {
			u := tc.User
			err := u.Add(tc.Commander)

			got := tc.Commander.About()
			assert.Equal(t, tc.Want, got)

			if tc.Err == false {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	fc := &exec.FakeCommander{}
	fuc := &exec.FakeUnsuccessfulCommander{}
	cases := []struct {
		Name      string
		User      *User
		Want      []string
		Err       bool
		Commander exec.CommanderDelegate
	}{
		{
			Name: "Default",
			User: &User{
				Name: "fake-user",
			},
			Want: []string{
				"/usr/sbin/userdel",
				"fake-user",
			},
			Err:       false,
			Commander: fc,
		},
		{
			Name: "Returns an error",
			User: &User{
				Name: "fake-name",
			},
			Want:      []string(nil),
			Err:       true,
			Commander: fuc,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d-%s", i, tc.Name), func(t *testing.T) {
			u := tc.User
			err := u.Delete(tc.Commander)

			got := tc.Commander.About()
			assert.Equal(t, tc.Want, got)

			if tc.Err == false {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
