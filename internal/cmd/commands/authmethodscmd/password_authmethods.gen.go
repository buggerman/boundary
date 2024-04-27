// Code generated by "make cli"; DO NOT EDIT.
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authmethodscmd

import (
	"errors"
	"fmt"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/authmethods"
	"github.com/hashicorp/boundary/internal/cmd/base"
	"github.com/hashicorp/boundary/internal/cmd/common"
	"github.com/hashicorp/go-secure-stdlib/strutil"
	"github.com/mitchellh/cli"
	"github.com/posener/complete"
)

func initPasswordFlags() {
	flagsOnce.Do(func() {
		extraFlags := extraPasswordActionsFlagsMapFunc()
		for k, v := range extraFlags {
			flagsPasswordMap[k] = append(flagsPasswordMap[k], v...)
		}
	})
}

var (
	_ cli.Command             = (*PasswordCommand)(nil)
	_ cli.CommandAutocomplete = (*PasswordCommand)(nil)
)

type PasswordCommand struct {
	*base.Command

	Func string

	plural string

	extraPasswordCmdVars
}

func (c *PasswordCommand) AutocompleteArgs() complete.Predictor {
	initPasswordFlags()
	return complete.PredictAnything
}

func (c *PasswordCommand) AutocompleteFlags() complete.Flags {
	initPasswordFlags()
	return c.Flags().Completions()
}

func (c *PasswordCommand) Synopsis() string {
	if extra := extraPasswordSynopsisFunc(c); extra != "" {
		return extra
	}

	synopsisStr := "auth method"

	synopsisStr = fmt.Sprintf("%s %s", "password-type", synopsisStr)

	return common.SynopsisFunc(c.Func, synopsisStr)
}

func (c *PasswordCommand) Help() string {
	initPasswordFlags()

	var helpStr string
	helpMap := common.HelpMap("auth method")

	switch c.Func {

	default:

		helpStr = c.extraPasswordHelpFunc(helpMap)

	}

	// Keep linter from complaining if we don't actually generate code using it
	_ = helpMap
	return helpStr
}

var flagsPasswordMap = map[string][]string{

	"create": {"scope-id", "name", "description"},

	"update": {"id", "name", "description", "version"},
}

func (c *PasswordCommand) Flags() *base.FlagSets {
	if len(flagsPasswordMap[c.Func]) == 0 {
		return c.FlagSet(base.FlagSetNone)
	}

	set := c.FlagSet(base.FlagSetHTTP | base.FlagSetClient | base.FlagSetOutputFormat)
	f := set.NewFlagSet("Command Options")
	common.PopulateCommonFlags(c.Command, f, "password-type auth method", flagsPasswordMap, c.Func)

	extraPasswordFlagsFunc(c, set, f)

	return set
}

func (c *PasswordCommand) Run(args []string) int {
	initPasswordFlags()

	switch c.Func {
	case "":
		return cli.RunResultHelp

	}

	c.plural = "password-type auth method"
	switch c.Func {
	case "list":
		c.plural = "password-type auth methods"
	}

	f := c.Flags()

	if err := f.Parse(args); err != nil {
		c.PrintCliError(err)
		return base.CommandUserError
	}

	if strutil.StrListContains(flagsPasswordMap[c.Func], "id") && c.FlagId == "" {
		c.PrintCliError(errors.New("ID is required but not passed in via -id"))
		return base.CommandUserError
	}

	var opts []authmethods.Option

	if strutil.StrListContains(flagsPasswordMap[c.Func], "scope-id") {
		switch c.Func {

		case "create":
			if c.FlagScopeId == "" {
				c.PrintCliError(errors.New("Scope ID must be passed in via -scope-id or BOUNDARY_SCOPE_ID"))
				return base.CommandUserError
			}

		}
	}

	client, err := c.Client()
	if c.WrapperCleanupFunc != nil {
		defer func() {
			if err := c.WrapperCleanupFunc(); err != nil {
				c.PrintCliError(fmt.Errorf("Error cleaning kms wrapper: %w", err))
			}
		}()
	}
	if err != nil {
		c.PrintCliError(fmt.Errorf("Error creating API client: %w", err))
		return base.CommandCliError
	}
	authmethodsClient := authmethods.NewClient(client)

	switch c.FlagName {
	case "":
	case "null":
		opts = append(opts, authmethods.DefaultName())
	default:
		opts = append(opts, authmethods.WithName(c.FlagName))
	}

	switch c.FlagDescription {
	case "":
	case "null":
		opts = append(opts, authmethods.DefaultDescription())
	default:
		opts = append(opts, authmethods.WithDescription(c.FlagDescription))
	}

	switch c.FlagRecursive {
	case true:
		opts = append(opts, authmethods.WithRecursive(true))
	}

	if c.FlagFilter != "" {
		opts = append(opts, authmethods.WithFilter(c.FlagFilter))
	}

	var version uint32

	switch c.Func {

	case "update":
		switch c.FlagVersion {
		case 0:
			opts = append(opts, authmethods.WithAutomaticVersioning(true))
		default:
			version = uint32(c.FlagVersion)
		}

	}

	if ok := extraPasswordFlagsHandlingFunc(c, f, &opts); !ok {
		return base.CommandUserError
	}

	var resp *api.Response
	var item *authmethods.AuthMethod

	var createResult *authmethods.AuthMethodCreateResult

	var updateResult *authmethods.AuthMethodUpdateResult

	switch c.Func {

	case "create":
		createResult, err = authmethodsClient.Create(c.Context, "password", c.FlagScopeId, opts...)
		if exitCode := c.checkFuncError(err); exitCode > 0 {
			return exitCode
		}
		resp = createResult.GetResponse()
		item = createResult.GetItem()

	case "update":
		updateResult, err = authmethodsClient.Update(c.Context, c.FlagId, version, opts...)
		if exitCode := c.checkFuncError(err); exitCode > 0 {
			return exitCode
		}
		resp = updateResult.GetResponse()
		item = updateResult.GetItem()

	}

	resp, item, err = executeExtraPasswordActions(c, resp, item, err, authmethodsClient, version, opts)
	if exitCode := c.checkFuncError(err); exitCode > 0 {
		return exitCode
	}

	output, err := printCustomPasswordActionOutput(c)
	if err != nil {
		c.PrintCliError(err)
		return base.CommandUserError
	}
	if output {
		return base.CommandSuccess
	}

	switch c.Func {

	}

	switch base.Format(c.UI) {
	case "table":
		c.UI.Output(printItemTable(item, resp))

	case "json":
		if ok := c.PrintJsonItem(resp); !ok {
			return base.CommandCliError
		}
	}

	return base.CommandSuccess
}

func (c *PasswordCommand) checkFuncError(err error) int {
	if err == nil {
		return 0
	}
	if apiErr := api.AsServerError(err); apiErr != nil {
		c.PrintApiError(apiErr, fmt.Sprintf("Error from controller when performing %s on %s", c.Func, c.plural))
		return base.CommandApiError
	}
	c.PrintCliError(fmt.Errorf("Error trying to %s %s: %s", c.Func, c.plural, err.Error()))
	return base.CommandCliError
}

var (
	extraPasswordActionsFlagsMapFunc = func() map[string][]string { return nil }
	extraPasswordSynopsisFunc        = func(*PasswordCommand) string { return "" }
	extraPasswordFlagsFunc           = func(*PasswordCommand, *base.FlagSets, *base.FlagSet) {}
	extraPasswordFlagsHandlingFunc   = func(*PasswordCommand, *base.FlagSets, *[]authmethods.Option) bool { return true }
	executeExtraPasswordActions      = func(_ *PasswordCommand, inResp *api.Response, inItem *authmethods.AuthMethod, inErr error, _ *authmethods.Client, _ uint32, _ []authmethods.Option) (*api.Response, *authmethods.AuthMethod, error) {
		return inResp, inItem, inErr
	}
	printCustomPasswordActionOutput = func(*PasswordCommand) (bool, error) { return false, nil }
)