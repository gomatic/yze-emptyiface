package emptyiface_test

import (
	"testing"

	emptyiface "github.com/gomatic/yze-go-emptyiface"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestEmptyInterfaceIsReportedAndFixed(t *testing.T) {
	analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), emptyiface.Analyzer, "a")
}

func TestRegistrationIsWellFormed(t *testing.T) {
	assert.NoError(t, emptyiface.Registration.Validate())
	assert.Equal(t, "yze/go/emptyiface", emptyiface.Registration.RuleID())
	assert.Same(t, emptyiface.Analyzer, emptyiface.Registration.Analyzer)
}
