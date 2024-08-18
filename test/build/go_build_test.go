package go_build_test

import (
	"testing"

	_ "github.com/aaomidi/xds/xds/data/orca/v3"
	_ "github.com/aaomidi/xds/xds/service/orca/v3"
	_ "github.com/aaomidi/xds/xds/type/v3"
)

func TestNoop(t *testing.T) {
	// Noop test that verifies the successful importation of xDS protos
}
