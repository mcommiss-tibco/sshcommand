package sshcommand

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	// No settings needed for SSH command execution
}

type Input struct {
	SSHServerName string `md:"sshservername,required"`
	SSHServerPort int    `md:"sshserverport,required"`
	SSHUsername   string `md:"sshusername,required"`
	SSHPassword   string `md:"sshpassword,required"`
	SSHCommand    string `md:"sshcommand,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	var err error

	r.SSHServerName, err = coerce.ToString(values["sshservername"])
	if err != nil {
		return err
	}

	r.SSHServerPort, err = coerce.ToInt(values["sshserverport"])
	if err != nil {
		return err
	}

	r.SSHUsername, err = coerce.ToString(values["sshusername"])
	if err != nil {
		return err
	}

	r.SSHPassword, err = coerce.ToString(values["sshpassword"])
	if err != nil {
		return err
	}

	r.SSHCommand, err = coerce.ToString(values["sshcommand"])
	if err != nil {
		return err
	}

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"sshservername": r.SSHServerName,
		"sshserverport": r.SSHServerPort,
		"sshusername":   r.SSHUsername,
		"sshpassword":   r.SSHPassword,
		"sshcommand":    r.SSHCommand,
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
