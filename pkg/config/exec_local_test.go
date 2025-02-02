package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCreatesExecLocal(t *testing.T) {
	c := NewExecLocal("abc")

	assert.Equal(t, "abc", c.Name)
	assert.Equal(t, TypeExecLocal, c.Type)
}

func TestExecLocalCreatesCorrectly(t *testing.T) {
	c, _ := CreateConfigFromStrings(t, execLocalRelative)

	ex, err := c.FindResource("exec_local.setup_vault")
	assert.NoError(t, err)

	assert.Equal(t, "setup_vault", ex.Info().Name)
	assert.Equal(t, TypeExecLocal, ex.Info().Type)
	assert.Equal(t, PendingCreation, ex.Info().Status)
	assert.Equal(t, "./", ex.(*ExecLocal).WorkingDirectory)
	assert.True(t, ex.(*ExecLocal).Daemon)
}

func TestExecLocalSetsDisabled(t *testing.T) {
	c, _ := CreateConfigFromStrings(t, execLocalDisabled)

	ex, err := c.FindResource("exec_local.setup_vault")
	assert.NoError(t, err)

	assert.Equal(t, Disabled, ex.Info().Status)
}

var execLocalRelative = `
exec_local "setup_vault" {
  cmd = "./scripts/setup_vault.sh"
  args = [ "root", "abc" ] 
  working_directory = "./"
  daemon = true
}
`
var execLocalDisabled = `
exec_local "setup_vault" {
	disabled = true

  cmd = "./scripts/setup_vault.sh"
  args = [ "root", "abc" ] 
  working_directory = "./"
  daemon = true
}
`
