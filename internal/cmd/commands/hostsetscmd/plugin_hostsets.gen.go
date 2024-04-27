// Code generated by "make cli"; DO NOT EDIT.
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hostsetscmd

import (
	"errors"
	"fmt"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/hostsets"
	"github.com/hashicorp/boundary/internal/cmd/base"
	"github.com/hashicorp/boundary/internal/cmd/common"
	"github.com/hashicorp/go-secure-stdlib/strutil"
	"github.com/mitchellh/cli"
	"github.com/posener/complete"
)

func initPluginFlags() {
	flagsOnce.Do(func() {
		extraFlags := extraPluginActionsFlagsMapFunc()
		for k, v := range extraFlags {
			flagsPluginMap[k] = append(flagsPluginMap[k], v...)
		}
	})
}

var (
	_ cli.Command             = (*PluginCommand)(nil)
	_ cli.CommandAutocomplete = (*PluginCommand)(nil)
)

type PluginCommand struct {
	*base.Command

	Func string

	plural string

	extraPluginCmdVars
}

func (c *PluginCommand) AutocompleteArgs() complete.Predictor {
	initPluginFlags()
	return complete.PredictAnything
}

func (c *PluginCommand) AutocompleteFlags() complete.Flags {
	initPluginFlags()
	return c.Flags().Completions()
}

func (c *PluginCommand) Synopsis() string {
	if extra := extraPluginSynopsisFunc(c); extra != "" {
		return extra
	}

	synopsisStr := "host set"

	synopsisStr = fmt.Sprintf("%s %s", "plugin-type", synopsisStr)

	return common.SynopsisFunc(c.Func, synopsisStr)
}

func (c *PluginCommand) Help() string {
	initPluginFlags()

	var helpStr string
	helpMap := common.HelpMap("host set")

	switch c.Func {

	default:

		helpStr = c.extraPluginHelpFunc(helpMap)

	}

	// Keep linter from complaining if we don't actually generate code using it
	_ = helpMap
	return helpStr
}

var flagsPluginMap = map[string][]string{

	"create": {"host-catalog-id", "name", "description", "attributes", "attr", "string-attr", "bool-attr", "num-attr"},

	"update": {"id", "name", "description", "version", "attributes", "attr", "string-attr", "bool-attr", "num-attr"},
}

func (c *PluginCommand) Flags() *base.FlagSets {
	if len(flagsPluginMap[c.Func]) == 0 {
		return c.FlagSet(base.FlagSetNone)
	}

	set := c.FlagSet(base.FlagSetHTTP | base.FlagSetClient | base.FlagSetOutputFormat)
	f := set.NewFlagSet("Command Options")
	common.PopulateCommonFlags(c.Command, f, "plugin-type host set", flagsPluginMap, c.Func)

	f = set.NewFlagSet("Attribute Options")
	attrsInput := common.CombinedSliceFlagValuePopulationInput{
		FlagSet:                          f,
		FlagNames:                        flagsPluginMap[c.Func],
		FullPopulationFlag:               &c.FlagAttributes,
		FullPopulationInputName:          "attributes",
		PiecewisePopulationFlag:          &c.FlagAttrs,
		PiecewisePopulationInputBaseName: "attr",
	}
	common.PopulateCombinedSliceFlagValue(attrsInput)

	extraPluginFlagsFunc(c, set, f)

	return set
}

func (c *PluginCommand) Run(args []string) int {
	initPluginFlags()

	switch c.Func {
	case "":
		return cli.RunResultHelp

	}

	c.plural = "plugin-type host set"
	switch c.Func {
	case "list":
		c.plural = "plugin-type host sets"
	}

	f := c.Flags()

	if err := f.Parse(args); err != nil {
		c.PrintCliError(err)
		return base.CommandUserError
	}

	if strutil.StrListContains(flagsPluginMap[c.Func], "id") && c.FlagId == "" {
		c.PrintCliError(errors.New("ID is required but not passed in via -id"))
		return base.CommandUserError
	}

	if c.FlagAttributes != "" && len(c.FlagAttrs) > 0 {
		c.PrintCliError(errors.New("-attributes flag cannot be used along with the following flags: attr, bool-attr, num-attr, string-attr"))
		return base.CommandUserError
	}

	var opts []hostsets.Option

	if strutil.StrListContains(flagsPluginMap[c.Func], "host-catalog-id") {
		switch c.Func {

		case "create":
			if c.FlagHostCatalogId == "" {
				c.PrintCliError(errors.New("HostCatalog ID must be passed in via -host-catalog-id or BOUNDARY_HOST_CATALOG_ID"))
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
	hostsetsClient := hostsets.NewClient(client)

	switch c.FlagName {
	case "":
	case "null":
		opts = append(opts, hostsets.DefaultName())
	default:
		opts = append(opts, hostsets.WithName(c.FlagName))
	}

	switch c.FlagDescription {
	case "":
	case "null":
		opts = append(opts, hostsets.DefaultDescription())
	default:
		opts = append(opts, hostsets.WithDescription(c.FlagDescription))
	}

	if c.FlagFilter != "" {
		opts = append(opts, hostsets.WithFilter(c.FlagFilter))
	}

	var version uint32

	switch c.Func {

	case "update":
		switch c.FlagVersion {
		case 0:
			opts = append(opts, hostsets.WithAutomaticVersioning(true))
		default:
			version = uint32(c.FlagVersion)
		}

	}

	if err := common.HandleAttributeFlags(
		c.Command,
		"attr",
		c.FlagAttributes,
		c.FlagAttrs,
		func() {
			opts = append(opts, hostsets.DefaultAttributes())
		},
		func(in map[string]interface{}) {
			opts = append(opts, hostsets.WithAttributes(in))
		}); err != nil {
		c.PrintCliError(fmt.Errorf("Error evaluating attribute flags to: %s", err.Error()))
		return base.CommandCliError
	}

	if ok := extraPluginFlagsHandlingFunc(c, f, &opts); !ok {
		return base.CommandUserError
	}

	var resp *api.Response
	var item *hostsets.HostSet

	var createResult *hostsets.HostSetCreateResult

	var updateResult *hostsets.HostSetUpdateResult

	switch c.Func {

	case "create":
		createResult, err = hostsetsClient.Create(c.Context, c.FlagHostCatalogId, opts...)
		if exitCode := c.checkFuncError(err); exitCode > 0 {
			return exitCode
		}
		resp = createResult.GetResponse()
		item = createResult.GetItem()

	case "update":
		updateResult, err = hostsetsClient.Update(c.Context, c.FlagId, version, opts...)
		if exitCode := c.checkFuncError(err); exitCode > 0 {
			return exitCode
		}
		resp = updateResult.GetResponse()
		item = updateResult.GetItem()

	}

	resp, item, err = executeExtraPluginActions(c, resp, item, err, hostsetsClient, version, opts)
	if exitCode := c.checkFuncError(err); exitCode > 0 {
		return exitCode
	}

	output, err := printCustomPluginActionOutput(c)
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

func (c *PluginCommand) checkFuncError(err error) int {
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
	extraPluginActionsFlagsMapFunc = func() map[string][]string { return nil }
	extraPluginSynopsisFunc        = func(*PluginCommand) string { return "" }
	extraPluginFlagsFunc           = func(*PluginCommand, *base.FlagSets, *base.FlagSet) {}
	extraPluginFlagsHandlingFunc   = func(*PluginCommand, *base.FlagSets, *[]hostsets.Option) bool { return true }
	executeExtraPluginActions      = func(_ *PluginCommand, inResp *api.Response, inItem *hostsets.HostSet, inErr error, _ *hostsets.Client, _ uint32, _ []hostsets.Option) (*api.Response, *hostsets.HostSet, error) {
		return inResp, inItem, inErr
	}
	printCustomPluginActionOutput = func(*PluginCommand) (bool, error) { return false, nil }
)