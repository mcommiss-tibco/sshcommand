package sshcommand

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	// No settings needed for SSH command execution
}

type Input struct {
	SshServername string `md:"sshservername,required"`
	SshServerPort string `md:"sshserverport,required"`
	SshUsername   string `md:"sshusername,required"`
	SshPassword   string `md:"sshpassword"`
	SshCommand    string `md:"sshcommand,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	var err error

	r.SshServername, err = coerce.ToString(values["sshservername"])
	if err != nil {
		return err
	}

	r.SshServerPort, err = coerce.ToString(values["sshserverport"])
	if err != nil {
		return err
	}

	r.SshUsername, err = coerce.ToString(values["sshusername"])
	if err != nil {
		return err
	}

	r.SshPassword, err = coerce.ToString(values["sshpassword"])
	if err != nil {
		return err
	}

	r.SshCommand, err = coerce.ToString(values["sshcommand"])
	if err != nil {
		return err
	}

	return nil
}

func (r *Input) ToMap() map[string]interface{} {

	return map[string]interface{}{
		"sshservername": r.SshServername,
		"sshserverport": r.SshServerPort,
		"sshusername":   r.SshUsername,
		"sshpassword":   r.SshPassword,
		"sshcommand":    r.SshCommand,
	}
}

type Output struct {
	StatusCode int    `md:"statuscode"`
	Result     string `md:"result"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	var err error

	o.StatusCode, err = coerce.ToInt(values["statuscode"])
	if err != nil {
		return err
	}

	o.Result, err = coerce.ToString(values["result"])
	if err != nil {
		return err
	}

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"statuscode": o.StatusCode,
		"result":     o.Result,
	}
}
