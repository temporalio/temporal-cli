package activity

import (
	"github.com/temporalio/temporal-cli/common"
	"github.com/urfave/cli/v2"
)

func NewActivityCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "complete",
			Usage: "Complete an activity",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     common.FlagWorkflowID,
					Aliases:  common.FlagWorkflowIDAlias,
					Usage:    "Workflow Id",
					Required: true,
					Category: common.CategoryMain,
				},
				&cli.StringFlag{
					Name:     common.FlagRunID,
					Aliases:  common.FlagRunIDAlias,
					Usage:    "Run Id",
					Required: true,
					Category: common.CategoryMain,
				},
				&cli.StringFlag{
					Name:     common.FlagActivityID,
					Usage:    "The Activity Id to complete",
					Required: true,
					Category: common.CategoryMain,
				},
				&cli.StringFlag{
					Name:     common.FlagResult,
					Usage:    "Set the result value of completion",
					Required: true,
					Category: common.CategoryMain,
				},
				&cli.StringFlag{
					Name:     common.FlagIdentity,
					Usage:    "Specify operator's identity",
					Required: true,
					Category: common.CategoryMain,
				},
			},
			Action: func(c *cli.Context) error {
				return CompleteActivity(c)
			},
		},
		{
			Name:  "fail",
			Usage: "Fail an activity",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     common.FlagWorkflowID,
					Aliases:  common.FlagWorkflowIDAlias,
					Usage:    "Workflow Id",
					Required: true,
					Category: common.CategoryMain,
				},
				&cli.StringFlag{
					Name:     common.FlagRunID,
					Aliases:  common.FlagRunIDAlias,
					Usage:    "Run Id",
					Required: true,
					Category: common.CategoryMain,
				},
				&cli.StringFlag{
					Name:     common.FlagActivityID,
					Usage:    "The Activity Id to fail",
					Required: true,
					Category: common.CategoryMain,
				},
				&cli.StringFlag{
					Name:     common.FlagReason,
					Usage:    "Reason to fail the Activity",
					Required: true,
					Category: common.CategoryMain,
				},
				&cli.StringFlag{
					Name:     common.FlagDetail,
					Usage:    "Detail to fail the Activity",
					Required: true,
					Category: common.CategoryMain,
				},
				&cli.StringFlag{
					Name:     common.FlagIdentity,
					Usage:    "Specify operator's identity",
					Required: true,
					Category: common.CategoryMain,
				},
			},
			Action: func(c *cli.Context) error {
				return FailActivity(c)
			},
		},
	}
}
