//go:build integration
// +build integration

package main

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerWithRealDependencies(t *testing.T) {
	s, err := NewServer(&MongoDB{}, &Redis{})
	assert.NoError(t, err)

	srv := httptest.NewServer(s)
	defer srv.Close()

	testServer(srv, t)
}
