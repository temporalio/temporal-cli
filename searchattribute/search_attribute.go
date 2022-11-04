// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package searchattribute

import (
	"fmt"

	"github.com/temporalio/cli/common"
	"github.com/temporalio/tctl-kit/pkg/output"
	"github.com/urfave/cli/v2"
	enumspb "go.temporal.io/api/enums/v1"
)

func NewSearchAttributeCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "create",
			Usage: "Add custom search attributes",
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:     common.FlagName,
					Required: true,
					Usage:    "Search attribute name",
				},
				&cli.StringSliceFlag{
					Name:     common.FlagType,
					Required: true,
					Usage:    fmt.Sprintf("Search attribute type: %v", common.AllowedEnumValues(enumspb.IndexedValueType_name)),
				},
				&cli.BoolFlag{
					Name:    common.FlagYes,
					Aliases: common.FlagYesAlias,
					Usage:   "Confirm all prompts",
				},
			},
			Action: func(c *cli.Context) error {
				return AddSearchAttributes(c)
			},
		},
		{
			Name:  "list",
			Usage: "List search attributes that can be used in list workflow query",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    output.FlagOutput,
					Aliases: common.FlagOutputAlias,
					Usage:   output.UsageText,
					Value:   string(output.Table),
				},
			},
			Action: func(c *cli.Context) error {
				return ListSearchAttributes(c)
			},
		},
		{
			Name:  "remove",
			Usage: "Remove custom search attributes metadata only (Elasticsearch index schema is not modified)",
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:     common.FlagName,
					Required: true,
					Usage:    "Search attribute name",
				},
				&cli.BoolFlag{
					Name:    common.FlagYes,
					Aliases: common.FlagYesAlias,
					Usage:   "Confirm all prompts",
				},
			},
			Action: func(c *cli.Context) error {
				return RemoveSearchAttributes(c)
			},
		},
	}
}