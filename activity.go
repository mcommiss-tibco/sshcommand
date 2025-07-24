package sshcommand

import (
	"github.com/project-flogo/core/activity"
)

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

// Activity is an SSH command execution activity
type Activity struct{}

func init() {
	_ = activity.Register(&Activity{})
}

// New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {
	// No settings needed for SSH command execution
	act := &Activity{}
	return act, nil
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Executes SSH command on remote server
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	logger := ctx.Logger()
	logger.Info("SSH Command Activity - Starting evaluation")
	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	logger.Info("SSH Command Activity - Connecting to: ", input.SshServername, ":", input.SshServerPort)
	logger.Debug("SSH Command Activity - Username: ", input.SshUsername)
	logger.Debug("SSH Command Activity - Command: ", input.SshCommand)
	var password string = input.SshPassword
	logger.Debug("SSH Command Activity - Password: ", password)

	// TODO: Implement SSH connection and command execution logic here

	// For now, return placeholder values
	output := &Output{
		StatusCode: 0,                                             // Success status code
		Result:     "Command executed successfully (placeholder)", // Placeholder result
	}

	ctx.Logger().Info("SSH Command Activity - Command completed with status: ", output.StatusCode)

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
