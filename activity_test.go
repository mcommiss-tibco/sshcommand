package sshcommand

import (
"testing"

"github.com/project-flogo/core/activity"
"github.com/project-flogo/core/support/test"
"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

ref := activity.GetRef(&Activity{})
act := activity.Get(ref)

assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

act := &Activity{}
tc := test.NewActivityContext(act.Metadata())

input := &Input{
SSHServerName: "test-server.com",
SSHServerPort: 22,
SSHUsername:   "testuser",
SSHPassword:   "testpass",
SSHCommand:    "ls -la",
}

err := tc.SetInputObject(input)
assert.Nil(t, err)

done, err := act.Eval(tc)
assert.True(t, done)
assert.Nil(t, err)

// Test that the output parameters are set correctly
output := &Output{}
err = tc.GetOutputObject(output)
assert.Nil(t, err)
assert.Equal(t, 0, output.StatusCode)
assert.Contains(t, output.Result, "placeholder")
}
