// Code generated. DO NOT EDIT.

package temporalcli

import (
	"github.com/mattn/go-isatty"

	"github.com/spf13/cobra"

	"github.com/spf13/pflag"

	"os"

	"time"
)

var hasHighlighting = isatty.IsTerminal(os.Stdout.Fd())

type TemporalCommand struct {
	Command                 cobra.Command
	Env                     string
	EnvFile                 string
	LogLevel                StringEnum
	LogFormat               string
	Output                  StringEnum
	TimeFormat              StringEnum
	Color                   StringEnum
	NoJsonShorthandPayloads bool
}

func NewTemporalCommand(cctx *CommandContext) *TemporalCommand {
	var s TemporalCommand
	s.Command.Use = "temporal"
	s.Command.Short = "Temporal command-line interface and development server"
	if hasHighlighting {
		s.Command.Long = "The Temporal CLI manages, monitors, and debugs Temporal apps. It lets you\nrun a local Temporal Service, start Workflow Executions, pass messages to\nrunning Workflows, inspects state, and more.\n\n* Start a local development service:\n      \x1b[1mtemporal server start-dev\x1b[0m\n* View help: pass --help to any command:\n      \x1b[1mtemporal activity complete --help\x1b[0m"
	} else {
		s.Command.Long = "The Temporal CLI manages, monitors, and debugs Temporal apps. It lets you\nrun a local Temporal Service, start Workflow Executions, pass messages to\nrunning Workflows, inspects state, and more.\n\n* Start a local development service:\n      `temporal server start-dev`\n* View help: pass --help to any command:\n      `temporal activity complete --help`"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.AddCommand(&NewTemporalActivityCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalBatchCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalEnvCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalOperatorCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalScheduleCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalServerCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalTaskQueueCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalWorkflowCommand(cctx, &s).Command)
	s.Command.PersistentFlags().StringVar(&s.Env, "env", "default", "Active environment name (`ENV`).")
	cctx.BindFlagEnvVar(s.Command.PersistentFlags().Lookup("env"), "TEMPORAL_ENV")
	s.Command.PersistentFlags().StringVar(&s.EnvFile, "env-file", "", "Path to environment settings file. (defaults to `$HOME/.config/temporalio/temporal.yaml`).")
	s.LogLevel = NewStringEnum([]string{"debug", "info", "warn", "error", "never"}, "info")
	s.Command.PersistentFlags().Var(&s.LogLevel, "log-level", "Log level. Default is \"info\" for most commands and \"warn\" for `server start-dev`. Accepted values: debug, info, warn, error, never.")
	s.Command.PersistentFlags().StringVar(&s.LogFormat, "log-format", "text", "Log format. Accepted values: text, json.")
	s.Output = NewStringEnum([]string{"text", "json", "jsonl", "none"}, "text")
	s.Command.PersistentFlags().VarP(&s.Output, "output", "o", "Non-logging data output format. Accepted values: text, json, jsonl, none.")
	s.TimeFormat = NewStringEnum([]string{"relative", "iso", "raw"}, "relative")
	s.Command.PersistentFlags().Var(&s.TimeFormat, "time-format", "Time format. Accepted values: relative, iso, raw.")
	s.Color = NewStringEnum([]string{"always", "never", "auto"}, "auto")
	s.Command.PersistentFlags().Var(&s.Color, "color", "Output coloring. Accepted values: always, never, auto.")
	s.Command.PersistentFlags().BoolVar(&s.NoJsonShorthandPayloads, "no-json-shorthand-payloads", false, "Raw payload output, even if they are JSON.")
	s.initCommand(cctx)
	return &s
}

type TemporalActivityCommand struct {
	Parent  *TemporalCommand
	Command cobra.Command
	ClientOptions
}

func NewTemporalActivityCommand(cctx *CommandContext, parent *TemporalCommand) *TemporalActivityCommand {
	var s TemporalActivityCommand
	s.Parent = parent
	s.Command.Use = "activity"
	s.Command.Short = "Complete or fail an Activity"
	if hasHighlighting {
		s.Command.Long = "Update an Activity's state to as completed or failed. This marks an\nActivity as successfully finished or as having encountered an error:\n\n\x1b[1mtemporal activity complete \\\n    --activity-id=YourActivityId \\\n    --workflow-id=YourWorkflowId \\\n    --result='{\"YourResultKey\": \"YourResultValue\"}'\x1b[0m"
	} else {
		s.Command.Long = "Update an Activity's state to as completed or failed. This marks an\nActivity as successfully finished or as having encountered an error:\n\n```\ntemporal activity complete \\\n    --activity-id=YourActivityId \\\n    --workflow-id=YourWorkflowId \\\n    --result='{\"YourResultKey\": \"YourResultValue\"}'\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.AddCommand(&NewTemporalActivityCompleteCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalActivityFailCommand(cctx, &s).Command)
	s.ClientOptions.buildFlags(cctx, s.Command.PersistentFlags())
	return &s
}

type TemporalActivityCompleteCommand struct {
	Parent  *TemporalActivityCommand
	Command cobra.Command
	WorkflowReferenceOptions
	ActivityId string
	Result     string
	Identity   string
}

func NewTemporalActivityCompleteCommand(cctx *CommandContext, parent *TemporalActivityCommand) *TemporalActivityCompleteCommand {
	var s TemporalActivityCompleteCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "complete [flags]"
	s.Command.Short = "Complete an Activity"
	if hasHighlighting {
		s.Command.Long = "Complete an Activity, marking it as successfully finished. Specify the \nActivity ID and include a JSON result for the returned value:\n\n\x1b[1mtemporal activity complete \\\n    --activity-id YourActivityId \\\n    --workflow-id YourWorkflowId \\\n    --result '{\"YourResultKey\": \"YourResultVal\"}'\x1b[0m"
	} else {
		s.Command.Long = "Complete an Activity, marking it as successfully finished. Specify the \nActivity ID and include a JSON result for the returned value:\n\n```\ntemporal activity complete \\\n    --activity-id YourActivityId \\\n    --workflow-id YourWorkflowId \\\n    --result '{\"YourResultKey\": \"YourResultVal\"}'\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.WorkflowReferenceOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Flags().StringVar(&s.ActivityId, "activity-id", "", "Activity ID to complete.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "activity-id")
	s.Command.Flags().StringVar(&s.Result, "result", "", "Result `JSON` to return.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "result")
	s.Command.Flags().StringVar(&s.Identity, "identity", "", "Identity of the user submitting this request.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalActivityFailCommand struct {
	Parent  *TemporalActivityCommand
	Command cobra.Command
	WorkflowReferenceOptions
	ActivityId string
	Detail     string
	Identity   string
	Reason     string
}

func NewTemporalActivityFailCommand(cctx *CommandContext, parent *TemporalActivityCommand) *TemporalActivityFailCommand {
	var s TemporalActivityFailCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "fail [flags]"
	s.Command.Short = "Fail an Activity"
	if hasHighlighting {
		s.Command.Long = "Fail an Activity, marking it as having encountered an error. Specify the\nActivity and Workflow IDs:\n\n\x1b[1mtemporal activity fail \\\n    --activity-id YourActivityId \\\n    --workflow-id YourWorkflowId\x1b[0m"
	} else {
		s.Command.Long = "Fail an Activity, marking it as having encountered an error. Specify the\nActivity and Workflow IDs:\n\n```\ntemporal activity fail \\\n    --activity-id YourActivityId \\\n    --workflow-id YourWorkflowId\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.WorkflowReferenceOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Flags().StringVar(&s.ActivityId, "activity-id", "", "Activity ID to fail.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "activity-id")
	s.Command.Flags().StringVar(&s.Detail, "detail", "", "Reason for failing the Activity (JSON).")
	s.Command.Flags().StringVar(&s.Identity, "identity", "", "Identity of the user submitting this request.")
	s.Command.Flags().StringVar(&s.Reason, "reason", "", "Reason for failing the Activity.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalBatchCommand struct {
	Parent  *TemporalCommand
	Command cobra.Command
	ClientOptions
}

func NewTemporalBatchCommand(cctx *CommandContext, parent *TemporalCommand) *TemporalBatchCommand {
	var s TemporalBatchCommand
	s.Parent = parent
	s.Command.Use = "batch"
	s.Command.Short = "Manage batch jobs"
	if hasHighlighting {
		s.Command.Long = "A batch job executes a command on multiple Workflow Executions at once. \nUse an SQL-like \x1b[1m--query\x1b[0m to select which Workflow Executions to include:\n\n\x1b[1mtemporal batch workflow cancel \\\n  --query 'ExecutionStatus = \"Running\" AND WorkflowType=\"YourWorkflow\"' \\\n  --reason \"Testing\"\x1b[0m\n\nThe \x1b[1mbatch\x1b[0m keyword is optional when using a \x1b[1m--query\x1b[0m. The following command\nis functionally identical to the previous one:\n\n\x1b[1mtemporal workflow cancel \\\n  --query 'ExecutionStatus = \"Running\" AND WorkflowType=\"YourWorkflow\"' \\\n  --reason \"Testing\"\x1b[0m"
	} else {
		s.Command.Long = "A batch job executes a command on multiple Workflow Executions at once. \nUse an SQL-like `--query` to select which Workflow Executions to include:\n\n```\ntemporal batch workflow cancel \\\n  --query 'ExecutionStatus = \"Running\" AND WorkflowType=\"YourWorkflow\"' \\\n  --reason \"Testing\"\n```\n\nThe `batch` keyword is optional when using a `--query`. The following command\nis functionally identical to the previous one:\n\n```\ntemporal workflow cancel \\\n  --query 'ExecutionStatus = \"Running\" AND WorkflowType=\"YourWorkflow\"' \\\n  --reason \"Testing\"\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.AddCommand(&NewTemporalBatchDescribeCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalBatchListCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalBatchTerminateCommand(cctx, &s).Command)
	s.ClientOptions.buildFlags(cctx, s.Command.PersistentFlags())
	return &s
}

type TemporalBatchDescribeCommand struct {
	Parent  *TemporalBatchCommand
	Command cobra.Command
	JobId   string
}

func NewTemporalBatchDescribeCommand(cctx *CommandContext, parent *TemporalBatchCommand) *TemporalBatchDescribeCommand {
	var s TemporalBatchDescribeCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "describe [flags]"
	s.Command.Short = "Show batch job progress"
	if hasHighlighting {
		s.Command.Long = "Show the progress of an ongoing batch job. Pass a valid job ID to display\nits information:\n\n\x1b[1mtemporal batch describe \\\n    --job-id YourJobId\x1b[0m"
	} else {
		s.Command.Long = "Show the progress of an ongoing batch job. Pass a valid job ID to display\nits information:\n\n```\ntemporal batch describe \\\n    --job-id YourJobId\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringVar(&s.JobId, "job-id", "", "Batch job ID.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "job-id")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalBatchListCommand struct {
	Parent  *TemporalBatchCommand
	Command cobra.Command
	Limit   int
}

func NewTemporalBatchListCommand(cctx *CommandContext, parent *TemporalBatchCommand) *TemporalBatchListCommand {
	var s TemporalBatchListCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "list [flags]"
	s.Command.Short = "List all batch jobs"
	if hasHighlighting {
		s.Command.Long = "Return a list of batch jobs on the Service or within a \nsingle Namespace. For example, list the batch jobs for\n\"YourNamespace\":\n\n\x1b[1mtemporal batch list \\\n    --namespace YourNamespace\x1b[0m"
	} else {
		s.Command.Long = "Return a list of batch jobs on the Service or within a \nsingle Namespace. For example, list the batch jobs for\n\"YourNamespace\":\n\n```\ntemporal batch list \\\n    --namespace YourNamespace\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().IntVar(&s.Limit, "limit", 0, "Maximum number of batch jobs to display.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalBatchTerminateCommand struct {
	Parent  *TemporalBatchCommand
	Command cobra.Command
	JobId   string
	Reason  string
}

func NewTemporalBatchTerminateCommand(cctx *CommandContext, parent *TemporalBatchCommand) *TemporalBatchTerminateCommand {
	var s TemporalBatchTerminateCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "terminate [flags]"
	s.Command.Short = "Forcefully end a batch job"
	if hasHighlighting {
		s.Command.Long = "Terminate a batch job with the provided job ID. You must provide a reason\nfor the termination. The Service stores this explanation as\nmetadata for the termination event for later reference:\n\n\x1b[1mtemporal batch terminate \\\n    --job-id YourJobId \\\n    --reason YourTerminationReason\x1b[0m"
	} else {
		s.Command.Long = "Terminate a batch job with the provided job ID. You must provide a reason\nfor the termination. The Service stores this explanation as\nmetadata for the termination event for later reference:\n\n```\ntemporal batch terminate \\\n    --job-id YourJobId \\\n    --reason YourTerminationReason\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringVar(&s.JobId, "job-id", "", "Job Id to terminate.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "job-id")
	s.Command.Flags().StringVar(&s.Reason, "reason", "", "Reason for terminating the batch job.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "reason")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalEnvCommand struct {
	Parent  *TemporalCommand
	Command cobra.Command
}

func NewTemporalEnvCommand(cctx *CommandContext, parent *TemporalCommand) *TemporalEnvCommand {
	var s TemporalEnvCommand
	s.Parent = parent
	s.Command.Use = "env"
	s.Command.Short = "Manage environments"
	if hasHighlighting {
		s.Command.Long = "Environments manage key-value presets, auto-configuring Temporal CLI options\nfor you. You can set up distinct environments like \"dev\" and \"prod\" for \nconvenience:\n\n\x1b[1mtemporal env set \\\n    --env prod \\\n    --key address \\\n    --value production.f45a2.tmprl.cloud:7233\x1b[0m\n\nEach environment is isolated. Changes to \"prod\" presets won't affect \"dev\".\n\nFor easiest use, set a \x1b[1mTEMPORAL_ENV\x1b[0m environment variable in your shell.\nThe Temporal CLI checks for an \x1b[1m--env\x1b[0m option first, then checks the shell.\nIf neither is set, most commands use \x1b[1mdefault\x1b[0m environment presets if\navailable."
	} else {
		s.Command.Long = "Environments manage key-value presets, auto-configuring Temporal CLI options\nfor you. You can set up distinct environments like \"dev\" and \"prod\" for \nconvenience:\n\n```\ntemporal env set \\\n    --env prod \\\n    --key address \\\n    --value production.f45a2.tmprl.cloud:7233\n```\n\nEach environment is isolated. Changes to \"prod\" presets won't affect \"dev\".\n\nFor easiest use, set a `TEMPORAL_ENV` environment variable in your shell.\nThe Temporal CLI checks for an `--env` option first, then checks the shell.\nIf neither is set, most commands use `default` environment presets if\navailable."
	}
	s.Command.Args = cobra.NoArgs
	s.Command.AddCommand(&NewTemporalEnvDeleteCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalEnvGetCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalEnvListCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalEnvSetCommand(cctx, &s).Command)
	return &s
}

type TemporalEnvDeleteCommand struct {
	Parent  *TemporalEnvCommand
	Command cobra.Command
	Key     string
}

func NewTemporalEnvDeleteCommand(cctx *CommandContext, parent *TemporalEnvCommand) *TemporalEnvDeleteCommand {
	var s TemporalEnvDeleteCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "delete [flags]"
	s.Command.Short = "Delete an environment or environment property"
	if hasHighlighting {
		s.Command.Long = "Remove a presets environment entirely _or_ remove a key-value pair within\nan environment. If you don't specify an environment (with \x1b[1m--env\x1b[0m or by\nsetting the \x1b[1mTEMPORAL_ENV\x1b[0m variable), this command updates the default \nenvironment:\n\n\x1b[1mtemporal env delete \\\n    --env YourEnvironment\x1b[0m\n\nor\n\n\x1b[1mtemporal env delete \\\n    --env prod \\\n    --key tls-key-path\x1b[0m"
	} else {
		s.Command.Long = "Remove a presets environment entirely _or_ remove a key-value pair within\nan environment. If you don't specify an environment (with `--env` or by\nsetting the `TEMPORAL_ENV` variable), this command updates the default \nenvironment:\n\n```\ntemporal env delete \\\n    --env YourEnvironment\n```\n\nor\n\n```\ntemporal env delete \\\n    --env prod \\\n    --key tls-key-path\n```"
	}
	s.Command.Args = cobra.MaximumNArgs(1)
	s.Command.Flags().StringVarP(&s.Key, "key", "k", "", "Property name.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalEnvGetCommand struct {
	Parent  *TemporalEnvCommand
	Command cobra.Command
	Key     string
}

func NewTemporalEnvGetCommand(cctx *CommandContext, parent *TemporalEnvCommand) *TemporalEnvGetCommand {
	var s TemporalEnvGetCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "get [flags]"
	s.Command.Short = "Show environment properties"
	if hasHighlighting {
		s.Command.Long = "List the properties for a given environment:\n\n\x1b[1mtemporal env get \\\n    --env YourEnvironment\x1b[0m\n\nPrint a single property:\n\n\x1b[1mtemporal env get \\\n    --env YourEnvironment \\\n    --key YourPropertyKey\x1b[0m\n\nIf you don't specify an environment (with \x1b[1m--env\x1b[0m or by setting the \n\x1b[1mTEMPORAL_ENV\x1b[0m variable), this command lists properties of the default \nenvironment."
	} else {
		s.Command.Long = "List the properties for a given environment:\n\n```\ntemporal env get \\\n    --env YourEnvironment\n```\n\nPrint a single property:\n\n```\ntemporal env get \\\n    --env YourEnvironment \\\n    --key YourPropertyKey\n```\n\nIf you don't specify an environment (with `--env` or by setting the \n`TEMPORAL_ENV` variable), this command lists properties of the default \nenvironment."
	}
	s.Command.Args = cobra.MaximumNArgs(1)
	s.Command.Flags().StringVarP(&s.Key, "key", "k", "", "Property name.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalEnvListCommand struct {
	Parent  *TemporalEnvCommand
	Command cobra.Command
}

func NewTemporalEnvListCommand(cctx *CommandContext, parent *TemporalEnvCommand) *TemporalEnvListCommand {
	var s TemporalEnvListCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "list [flags]"
	s.Command.Short = "Show environment names"
	s.Command.Long = "List the environments you have set up on your local computer. Environments \nare stored in \"$HOME/.config/temporalio/temporal.yaml\"."
	s.Command.Args = cobra.NoArgs
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalEnvSetCommand struct {
	Parent  *TemporalEnvCommand
	Command cobra.Command
	Key     string
	Value   string
}

func NewTemporalEnvSetCommand(cctx *CommandContext, parent *TemporalEnvCommand) *TemporalEnvSetCommand {
	var s TemporalEnvSetCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "set [flags]"
	s.Command.Short = "Set environment properties"
	if hasHighlighting {
		s.Command.Long = "Assign a value to a property key and store it to an environment:\n\n\x1b[1mtemporal env set \\\n    --env environment \\\n    --key property \\\n    --value value\x1b[0m\n\nIf you don't specify an environment (with \x1b[1m--env\x1b[0m or by setting the \n\x1b[1mTEMPORAL_ENV\x1b[0m variable), this command sets properties in the default \nenvironment.\n\nStoring keys with CLI option names lets the CLI automatically set those\noptions for you. This reduces effort and helps avoid typos when issuing\ncommands."
	} else {
		s.Command.Long = "Assign a value to a property key and store it to an environment:\n\n```\ntemporal env set \\\n    --env environment \\\n    --key property \\\n    --value value\n```\n\nIf you don't specify an environment (with `--env` or by setting the \n`TEMPORAL_ENV` variable), this command sets properties in the default \nenvironment.\n\nStoring keys with CLI option names lets the CLI automatically set those\noptions for you. This reduces effort and helps avoid typos when issuing\ncommands."
	}
	s.Command.Args = cobra.MaximumNArgs(2)
	s.Command.Flags().StringVarP(&s.Key, "key", "k", "", "Property name.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "key")
	s.Command.Flags().StringVarP(&s.Value, "value", "v", "", "Property value.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "value")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalOperatorCommand struct {
	Parent  *TemporalCommand
	Command cobra.Command
	ClientOptions
}

func NewTemporalOperatorCommand(cctx *CommandContext, parent *TemporalCommand) *TemporalOperatorCommand {
	var s TemporalOperatorCommand
	s.Parent = parent
	s.Command.Use = "operator"
	s.Command.Short = "Manage Temporal deployments"
	if hasHighlighting {
		s.Command.Long = "Operator commands manage and fetch information about Namespace, \nSearch Attributes, and Temporal Services:\n\n\x1b[1mtemporal operator [command] [subcommand] [options]\x1b[0m\n\nFor example, to show information about the Temporal Service at the\ndefault address (localhost):\n\n\x1b[1mtemporal operator cluster describe\x1b[0m"
	} else {
		s.Command.Long = "Operator commands manage and fetch information about Namespace, \nSearch Attributes, and Temporal Services:\n\n```\ntemporal operator [command] [subcommand] [options]\n```\n\nFor example, to show information about the Temporal Service at the\ndefault address (localhost):\n\n```\ntemporal operator cluster describe\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.AddCommand(&NewTemporalOperatorClusterCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalOperatorNamespaceCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalOperatorSearchAttributeCommand(cctx, &s).Command)
	s.ClientOptions.buildFlags(cctx, s.Command.PersistentFlags())
	return &s
}

type TemporalOperatorClusterCommand struct {
	Parent  *TemporalOperatorCommand
	Command cobra.Command
}

func NewTemporalOperatorClusterCommand(cctx *CommandContext, parent *TemporalOperatorCommand) *TemporalOperatorClusterCommand {
	var s TemporalOperatorClusterCommand
	s.Parent = parent
	s.Command.Use = "cluster"
	s.Command.Short = "Manage a Temporal Cluster"
	if hasHighlighting {
		s.Command.Long = "Perform operator actions on Temporal Services (also known as Clusters).\n\n\x1b[1mtemporal operator cluster [subcommand] [options]\x1b[0m\n\nFor example to check Service/Cluster health:\n\n\x1b[1mtemporal operator cluster health\x1b[0m"
	} else {
		s.Command.Long = "Perform operator actions on Temporal Services (also known as Clusters).\n\n```\ntemporal operator cluster [subcommand] [options]\n```\n\nFor example to check Service/Cluster health:\n\n```\ntemporal operator cluster health\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.AddCommand(&NewTemporalOperatorClusterDescribeCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalOperatorClusterHealthCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalOperatorClusterListCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalOperatorClusterRemoveCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalOperatorClusterSystemCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalOperatorClusterUpsertCommand(cctx, &s).Command)
	return &s
}

type TemporalOperatorClusterDescribeCommand struct {
	Parent  *TemporalOperatorClusterCommand
	Command cobra.Command
	Detail  bool
}

func NewTemporalOperatorClusterDescribeCommand(cctx *CommandContext, parent *TemporalOperatorClusterCommand) *TemporalOperatorClusterDescribeCommand {
	var s TemporalOperatorClusterDescribeCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "describe [flags]"
	s.Command.Short = "Show Temporal Cluster information"
	if hasHighlighting {
		s.Command.Long = "View information about a Temporal Cluster (Service), including Cluster Name,\npersistence store, and visibility store. Add \x1b[1m--detail\x1b[0m for additional info:\n\n\x1b[1mtemporal operator cluster describe [--detail]\x1b[0m"
	} else {
		s.Command.Long = "View information about a Temporal Cluster (Service), including Cluster Name,\npersistence store, and visibility store. Add `--detail` for additional info:\n\n```\ntemporal operator cluster describe [--detail]\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().BoolVar(&s.Detail, "detail", false, "Show history shard count and Cluster/Service version information.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalOperatorClusterHealthCommand struct {
	Parent  *TemporalOperatorClusterCommand
	Command cobra.Command
}

func NewTemporalOperatorClusterHealthCommand(cctx *CommandContext, parent *TemporalOperatorClusterCommand) *TemporalOperatorClusterHealthCommand {
	var s TemporalOperatorClusterHealthCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "health [flags]"
	s.Command.Short = "Check Temporal Service health"
	if hasHighlighting {
		s.Command.Long = "View information about the health of a Temporal Service:\n\n\x1b[1mtemporal operator cluster health\x1b[0m"
	} else {
		s.Command.Long = "View information about the health of a Temporal Service:\n\n```\ntemporal operator cluster health\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalOperatorClusterListCommand struct {
	Parent  *TemporalOperatorClusterCommand
	Command cobra.Command
	Limit   int
}

func NewTemporalOperatorClusterListCommand(cctx *CommandContext, parent *TemporalOperatorClusterCommand) *TemporalOperatorClusterListCommand {
	var s TemporalOperatorClusterListCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "list [flags]"
	s.Command.Short = "Show Temporal Clusters"
	if hasHighlighting {
		s.Command.Long = "Print a list of Temporal Clusters (Services) registered to this\nsystem. Report details include the Cluster's name, ID, address, History Shard\ncount, Failover version, and availability:\n\n\x1b[1mtemporal operator cluster list [--limit max-count]\x1b[0m"
	} else {
		s.Command.Long = "Print a list of Temporal Clusters (Services) registered to this\nsystem. Report details include the Cluster's name, ID, address, History Shard\ncount, Failover version, and availability:\n\n```\ntemporal operator cluster list [--limit max-count]\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().IntVar(&s.Limit, "limit", 0, "Maximum number of Clusters to display.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalOperatorClusterRemoveCommand struct {
	Parent  *TemporalOperatorClusterCommand
	Command cobra.Command
	Name    string
}

func NewTemporalOperatorClusterRemoveCommand(cctx *CommandContext, parent *TemporalOperatorClusterCommand) *TemporalOperatorClusterRemoveCommand {
	var s TemporalOperatorClusterRemoveCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "remove [flags]"
	s.Command.Short = "Remove a Temporal Cluster"
	if hasHighlighting {
		s.Command.Long = "Remove a registered Temporal Cluster (Service) from this system.\n\n\x1b[1mtemporal operator cluster remove \\\n    --name YourClusterName\x1b[0m"
	} else {
		s.Command.Long = "Remove a registered Temporal Cluster (Service) from this system.\n\n```\ntemporal operator cluster remove \\\n    --name YourClusterName\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringVar(&s.Name, "name", "", "Cluster/Service name.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "name")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalOperatorClusterSystemCommand struct {
	Parent  *TemporalOperatorClusterCommand
	Command cobra.Command
}

func NewTemporalOperatorClusterSystemCommand(cctx *CommandContext, parent *TemporalOperatorClusterCommand) *TemporalOperatorClusterSystemCommand {
	var s TemporalOperatorClusterSystemCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "system [flags]"
	s.Command.Short = "Show Temporal Cluster info"
	if hasHighlighting {
		s.Command.Long = "Show Temporal Server information for Temporal Clusters (Service):\nServer version, scheduling support, and more. This information helps diagnose\nproblems with the Temporal Server.\n\nThe command defaults to the local Service. Otherwise, use the \n\x1b[1m--frontend-address\x1b[0m option to specify a Cluster (Service) endpoint:\n\n\x1b[1mtemporal operator cluster system \\ \n    --frontend-address \"YourRemoteEndpoint:YourRemotePort\"\x1b[0m"
	} else {
		s.Command.Long = "Show Temporal Server information for Temporal Clusters (Service):\nServer version, scheduling support, and more. This information helps diagnose\nproblems with the Temporal Server.\n\nThe command defaults to the local Service. Otherwise, use the \n`--frontend-address` option to specify a Cluster (Service) endpoint:\n\n```\ntemporal operator cluster system \\ \n    --frontend-address \"YourRemoteEndpoint:YourRemotePort\"\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalOperatorClusterUpsertCommand struct {
	Parent           *TemporalOperatorClusterCommand
	Command          cobra.Command
	FrontendAddress  string
	EnableConnection bool
}

func NewTemporalOperatorClusterUpsertCommand(cctx *CommandContext, parent *TemporalOperatorClusterCommand) *TemporalOperatorClusterUpsertCommand {
	var s TemporalOperatorClusterUpsertCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "upsert [flags]"
	s.Command.Short = "Add/update a Temporal Cluster"
	if hasHighlighting {
		s.Command.Long = "Add, remove, or update a registered (\"remote\") Temporal Cluster (Service).\n\n\x1b[1mtemporal operator cluster upsert [options]\x1b[0m\n\nFor example:\n\n\x1b[1mtemporal operator cluster upsert \\\n    --frontend-address \"YourRemoteEndpoint:YourRemotePort\"\n    --enable-connection false\x1b[0m"
	} else {
		s.Command.Long = "Add, remove, or update a registered (\"remote\") Temporal Cluster (Service).\n\n```\ntemporal operator cluster upsert [options]\n```\n\nFor example:\n\n```\ntemporal operator cluster upsert \\\n    --frontend-address \"YourRemoteEndpoint:YourRemotePort\"\n    --enable-connection false\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringVar(&s.FrontendAddress, "frontend-address", "", "Remote endpoint.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "frontend-address")
	s.Command.Flags().BoolVar(&s.EnableConnection, "enable-connection", false, "Set the connection to \"enabled\".")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalOperatorNamespaceCommand struct {
	Parent  *TemporalOperatorCommand
	Command cobra.Command
}

func NewTemporalOperatorNamespaceCommand(cctx *CommandContext, parent *TemporalOperatorCommand) *TemporalOperatorNamespaceCommand {
	var s TemporalOperatorNamespaceCommand
	s.Parent = parent
	s.Command.Use = "namespace"
	s.Command.Short = "Namespace operations"
	if hasHighlighting {
		s.Command.Long = "Manage Temporal Cluster (Service) Namespaces:\n\n\x1b[1mtemporal operator namespace [command] [command options]\x1b[0m\n\nFor example:\n\n\x1b[1mtemporal operator namespace create \\\n    --namespace YourNewNamespaceName\x1b[0m"
	} else {
		s.Command.Long = "Manage Temporal Cluster (Service) Namespaces:\n\n```\ntemporal operator namespace [command] [command options]\n```\n\nFor example:\n\n```\ntemporal operator namespace create \\\n    --namespace YourNewNamespaceName\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.AddCommand(&NewTemporalOperatorNamespaceCreateCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalOperatorNamespaceDeleteCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalOperatorNamespaceDescribeCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalOperatorNamespaceListCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalOperatorNamespaceUpdateCommand(cctx, &s).Command)
	return &s
}

type TemporalOperatorNamespaceCreateCommand struct {
	Parent                  *TemporalOperatorNamespaceCommand
	Command                 cobra.Command
	ActiveCluster           string
	Cluster                 []string
	Data                    string
	Description             string
	Email                   string
	Global                  bool
	HistoryArchivalState    StringEnum
	HistoryUri              string
	Retention               time.Duration
	VisibilityArchivalState StringEnum
	VisibilityUri           string
}

func NewTemporalOperatorNamespaceCreateCommand(cctx *CommandContext, parent *TemporalOperatorNamespaceCommand) *TemporalOperatorNamespaceCreateCommand {
	var s TemporalOperatorNamespaceCreateCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "create [flags]"
	s.Command.Short = "Register a new Namespace"
	if hasHighlighting {
		s.Command.Long = "Create a new Namespace on the Temporal Service:\n\n\x1b[1mtemporal operator namespace create \\\n    --namespace YourNewNamespaceName \\\n    [options]\x1b[0m`\n\nCreate a Namespace with multi-region data replication:\n\n\x1b[1mtemporal operator namespace create \\\n    --global \\\n    --namespace YourNewNamespaceName\x1b[0m\n\nConfigure settings like retention and Visibility Archival State as needed.\nFor example, the Visibility Archive can be set on a separate URI:\n\n\x1b[1mtemporal operator namespace create \\\n    --retention 5d \\\n    --visibility-archival-state enabled \\\n    --visibility-uri YourURI \\\n    --namespace YourNewNamespaceName\x1b[0m\n\nNote: URI values for archival states can't be changed once enabled."
	} else {
		s.Command.Long = "Create a new Namespace on the Temporal Service:\n\n```\ntemporal operator namespace create \\\n    --namespace YourNewNamespaceName \\\n    [options]\n````\n\nCreate a Namespace with multi-region data replication:\n\n```\ntemporal operator namespace create \\\n    --global \\\n    --namespace YourNewNamespaceName\n```\n\nConfigure settings like retention and Visibility Archival State as needed.\nFor example, the Visibility Archive can be set on a separate URI:\n\n```\ntemporal operator namespace create \\\n    --retention 5d \\\n    --visibility-archival-state enabled \\\n    --visibility-uri YourURI \\\n    --namespace YourNewNamespaceName\n```\n\nNote: URI values for archival states can't be changed once enabled."
	}
	s.Command.Args = cobra.MaximumNArgs(1)
	s.Command.Flags().StringVar(&s.ActiveCluster, "active-cluster", "", "Active Cluster (Service) name.")
	s.Command.Flags().StringArrayVar(&s.Cluster, "cluster", nil, "Cluster (Service) names for Namespace creation. Can be passed multiple times.")
	s.Command.Flags().StringVar(&s.Data, "data", "", "Namespace data.")
	s.Command.Flags().StringVar(&s.Description, "description", "", "Namespace description. Comma-separated 'KEY=\"VALUE\"' string pairs.")
	s.Command.Flags().StringVar(&s.Email, "email", "", "Owner email.")
	s.Command.Flags().BoolVar(&s.Global, "global", false, "Enable multi-region data replication.")
	s.HistoryArchivalState = NewStringEnum([]string{"disabled", "enabled"}, "disabled")
	s.Command.Flags().Var(&s.HistoryArchivalState, "history-archival-state", "History archival state. Accepted values: disabled, enabled.")
	s.Command.Flags().StringVar(&s.HistoryUri, "history-uri", "", "Archive history to this `URI`. Once enabled, can't be changed.")
	s.Command.Flags().DurationVar(&s.Retention, "retention", 259200000*time.Millisecond, "Time to preserve closed Workflows before deletion.")
	s.VisibilityArchivalState = NewStringEnum([]string{"disabled", "enabled"}, "disable")
	s.Command.Flags().Var(&s.VisibilityArchivalState, "visibility-archival-state", "Visibility archival state. Accepted values: disabled, enabled.")
	s.Command.Flags().StringVar(&s.VisibilityUri, "visibility-uri", "", "Archive visibility data to this `URI`. Once enabled, can't be changed.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalOperatorNamespaceDeleteCommand struct {
	Parent  *TemporalOperatorNamespaceCommand
	Command cobra.Command
	Yes     bool
}

func NewTemporalOperatorNamespaceDeleteCommand(cctx *CommandContext, parent *TemporalOperatorNamespaceCommand) *TemporalOperatorNamespaceDeleteCommand {
	var s TemporalOperatorNamespaceDeleteCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "delete [flags] [namespace]"
	s.Command.Short = "Delete Namespace"
	if hasHighlighting {
		s.Command.Long = "Removes a Namespace from the Service.\n\n\x1b[1mtemporal operator namespace delete [options]\x1b[0m\n\nFor example:\n\n\x1b[1mtemporal operator namespace delete \\\n    --namespace YourNamespaceName\x1b[0m"
	} else {
		s.Command.Long = "Removes a Namespace from the Service.\n\n```\ntemporal operator namespace delete [options]\n```\n\nFor example:\n\n```\ntemporal operator namespace delete \\\n    --namespace YourNamespaceName\n```"
	}
	s.Command.Args = cobra.MaximumNArgs(1)
	s.Command.Flags().BoolVarP(&s.Yes, "yes", "y", false, "Request confirmation before deletion.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalOperatorNamespaceDescribeCommand struct {
	Parent      *TemporalOperatorNamespaceCommand
	Command     cobra.Command
	NamespaceId string
}

func NewTemporalOperatorNamespaceDescribeCommand(cctx *CommandContext, parent *TemporalOperatorNamespaceCommand) *TemporalOperatorNamespaceDescribeCommand {
	var s TemporalOperatorNamespaceDescribeCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "describe [flags] [namespace]"
	s.Command.Short = "Describe a Namespace"
	if hasHighlighting {
		s.Command.Long = "Provide long-form information about a Namespace identified by its ID or name:\n\n\x1b[1mtemporal operator namespace describe \\\n    --namespace-id YourNamespaceId\x1b[0m\n\nor\n\n\x1b[1mtemporal operator namespace describe \\\n    --namespace YourNamespaceName\x1b[0m"
	} else {
		s.Command.Long = "Provide long-form information about a Namespace identified by its ID or name:\n\n```\ntemporal operator namespace describe \\\n    --namespace-id YourNamespaceId\n```\n\nor\n\n```\ntemporal operator namespace describe \\\n    --namespace YourNamespaceName\n```"
	}
	s.Command.Args = cobra.MaximumNArgs(1)
	s.Command.Flags().StringVar(&s.NamespaceId, "namespace-id", "", "Namespace ID.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalOperatorNamespaceListCommand struct {
	Parent  *TemporalOperatorNamespaceCommand
	Command cobra.Command
}

func NewTemporalOperatorNamespaceListCommand(cctx *CommandContext, parent *TemporalOperatorNamespaceCommand) *TemporalOperatorNamespaceListCommand {
	var s TemporalOperatorNamespaceListCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "list [flags]"
	s.Command.Short = "List Namespaces"
	if hasHighlighting {
		s.Command.Long = "Display a detailed listing for all Namespaces on the Service:\n\n\x1b[1mtemporal operator namespace list\x1b[0m"
	} else {
		s.Command.Long = "Display a detailed listing for all Namespaces on the Service:\n\n```\ntemporal operator namespace list\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalOperatorNamespaceUpdateCommand struct {
	Parent                  *TemporalOperatorNamespaceCommand
	Command                 cobra.Command
	ActiveCluster           string
	Cluster                 []string
	Data                    []string
	Description             string
	Email                   string
	PromoteGlobal           bool
	HistoryArchivalState    StringEnum
	HistoryUri              string
	Retention               time.Duration
	VisibilityArchivalState StringEnum
	VisibilityUri           string
}

func NewTemporalOperatorNamespaceUpdateCommand(cctx *CommandContext, parent *TemporalOperatorNamespaceCommand) *TemporalOperatorNamespaceUpdateCommand {
	var s TemporalOperatorNamespaceUpdateCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "update [flags]"
	s.Command.Short = "Update Namespace"
	if hasHighlighting {
		s.Command.Long = "Update a Namespace using properties you specify.\n\n\x1b[1mtemporal operator namespace update [options]\x1b[0m \n\nAssign a Namespace's active Cluster (Service):\n\n\x1b[1mtemporal operator namespace update \\\n    --namespace YourNamespaceName \\\n    --active-cluster NewActiveCluster\x1b[0m\n\nPromote a Namespace for multi-region data replication:\n\n\x1b[1mtemporal operator namespace update \\\n    --namespace YourNamespaceName \\\n    --promote-global\x1b[0m\n\nYou may update archives that were previously enabled or disabled.\nNote: URI values for archival states can't be changed once enabled.\n\n\x1b[1mtemporal operator namespace update \\\n    --namespace YourNamespaceName \\\n    --history-archival-state enabled \\\n    --visibility-archival-state disabled\x1b[0m"
	} else {
		s.Command.Long = "Update a Namespace using properties you specify.\n\n```\ntemporal operator namespace update [options]\n``` \n\nAssign a Namespace's active Cluster (Service):\n\n```\ntemporal operator namespace update \\\n    --namespace YourNamespaceName \\\n    --active-cluster NewActiveCluster\n```\n\nPromote a Namespace for multi-region data replication:\n\n```\ntemporal operator namespace update \\\n    --namespace YourNamespaceName \\\n    --promote-global\n```\n\nYou may update archives that were previously enabled or disabled.\nNote: URI values for archival states can't be changed once enabled.\n\n```\ntemporal operator namespace update \\\n    --namespace YourNamespaceName \\\n    --history-archival-state enabled \\\n    --visibility-archival-state disabled\n```"
	}
	s.Command.Args = cobra.MaximumNArgs(1)
	s.Command.Flags().StringVar(&s.ActiveCluster, "active-cluster", "", "Active Cluster (S) name.")
	s.Command.Flags().StringArrayVar(&s.Cluster, "cluster", nil, "Cluster names.")
	s.Command.Flags().StringArrayVar(&s.Data, "data", nil, "Set a 'KEY=\"VALUE\"' string pair in Namespace data. KEY is a string, VALUE is JSON. Can be passed multiple times.")
	s.Command.Flags().StringVar(&s.Description, "description", "", "Namespace description.")
	s.Command.Flags().StringVar(&s.Email, "email", "", "Owner email.")
	s.Command.Flags().BoolVar(&s.PromoteGlobal, "promote-global", false, "Enable multi-region data replication.")
	s.HistoryArchivalState = NewStringEnum([]string{"disabled", "enabled"}, "")
	s.Command.Flags().Var(&s.HistoryArchivalState, "history-archival-state", "History archival state. Accepted values: disabled, enabled.")
	s.Command.Flags().StringVar(&s.HistoryUri, "history-uri", "", "Archive history to this `URI`. Once enabled, can't be changed.")
	s.Command.Flags().DurationVar(&s.Retention, "retention", 0, "Length of time a closed Workflow is preserved before deletion")
	s.VisibilityArchivalState = NewStringEnum([]string{"disabled", "enable"}, "")
	s.Command.Flags().Var(&s.VisibilityArchivalState, "visibility-archival-state", "Visibility archival state. Accepted values: disabled, enable.")
	s.Command.Flags().StringVar(&s.VisibilityUri, "visibility-uri", "", "Archive visibility data to this `URI`. Once enabled, can't be changed.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalOperatorSearchAttributeCommand struct {
	Parent  *TemporalOperatorCommand
	Command cobra.Command
}

func NewTemporalOperatorSearchAttributeCommand(cctx *CommandContext, parent *TemporalOperatorCommand) *TemporalOperatorSearchAttributeCommand {
	var s TemporalOperatorSearchAttributeCommand
	s.Parent = parent
	s.Command.Use = "search-attribute"
	s.Command.Short = "Search Attribute operations"
	if hasHighlighting {
		s.Command.Long = "Create, list, or remove Search Attributes\nfields stored in a Workflow Execution's metadata:\n\n\x1b[1mtemporal operator search-attribute create \\\n    --name YourAttributeName \\\n    --type Keyword\x1b[0m\n\nSupported types include: Text, Keyword, Int, Double, Bool, Datetime, and\nKeywordList."
	} else {
		s.Command.Long = "Create, list, or remove Search Attributes\nfields stored in a Workflow Execution's metadata:\n\n```\ntemporal operator search-attribute create \\\n    --name YourAttributeName \\\n    --type Keyword\n```\n\nSupported types include: Text, Keyword, Int, Double, Bool, Datetime, and\nKeywordList."
	}
	s.Command.Args = cobra.NoArgs
	s.Command.AddCommand(&NewTemporalOperatorSearchAttributeCreateCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalOperatorSearchAttributeListCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalOperatorSearchAttributeRemoveCommand(cctx, &s).Command)
	return &s
}

type TemporalOperatorSearchAttributeCreateCommand struct {
	Parent  *TemporalOperatorSearchAttributeCommand
	Command cobra.Command
	Name    []string
	Type    []string
}

func NewTemporalOperatorSearchAttributeCreateCommand(cctx *CommandContext, parent *TemporalOperatorSearchAttributeCommand) *TemporalOperatorSearchAttributeCreateCommand {
	var s TemporalOperatorSearchAttributeCreateCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "create [flags]"
	s.Command.Short = "Add custom Search Attributes"
	if hasHighlighting {
		s.Command.Long = "Add one or more custom Search Attributes:\n\n\x1b[1mtemporal operator search-attribute create \\\n    --name YourAttributeName \\\n    --type Keyword\x1b[0m"
	} else {
		s.Command.Long = "Add one or more custom Search Attributes:\n\n```\ntemporal operator search-attribute create \\\n    --name YourAttributeName \\\n    --type Keyword\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringArrayVar(&s.Name, "name", nil, "Search Attribute name. Required")
	s.Command.Flags().StringArrayVar(&s.Type, "type", nil, "Search Attribute type. Accepted values: Text, Keyword, Int, Double, Bool, Datetime, KeywordList.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "type")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalOperatorSearchAttributeListCommand struct {
	Parent  *TemporalOperatorSearchAttributeCommand
	Command cobra.Command
}

func NewTemporalOperatorSearchAttributeListCommand(cctx *CommandContext, parent *TemporalOperatorSearchAttributeCommand) *TemporalOperatorSearchAttributeListCommand {
	var s TemporalOperatorSearchAttributeListCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "list [flags]"
	s.Command.Short = "List Search Attributes"
	if hasHighlighting {
		s.Command.Long = "Display a list of active Search Attributes that can be \nassigned or used with Workflow Queries. You can manage this list and add\nattributes as needed:\n\n\x1b[1mtemporal operator search-attribute list\x1b[0m"
	} else {
		s.Command.Long = "Display a list of active Search Attributes that can be \nassigned or used with Workflow Queries. You can manage this list and add\nattributes as needed:\n\n```\ntemporal operator search-attribute list\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalOperatorSearchAttributeRemoveCommand struct {
	Parent  *TemporalOperatorSearchAttributeCommand
	Command cobra.Command
	Name    []string
	Yes     bool
}

func NewTemporalOperatorSearchAttributeRemoveCommand(cctx *CommandContext, parent *TemporalOperatorSearchAttributeCommand) *TemporalOperatorSearchAttributeRemoveCommand {
	var s TemporalOperatorSearchAttributeRemoveCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "remove [flags]"
	s.Command.Short = "Remove custom Search Attributes"
	if hasHighlighting {
		s.Command.Long = "Remove custom Search Attributes from the options that can be\nassigned or used with Workflow Queries.\n\n\x1b[1mtemporal operator search-attribute remove \\\n    --name YourAttributeName\x1b[0m\n\nRemove attributes without confirmation:\n\n\x1b[1mtemporal operator search-attribute remove \\\n    --name YourAttributeName \\\n    --yes\x1b[0m"
	} else {
		s.Command.Long = "Remove custom Search Attributes from the options that can be\nassigned or used with Workflow Queries.\n\n```\ntemporal operator search-attribute remove \\\n    --name YourAttributeName\n```\n\nRemove attributes without confirmation:\n\n```\ntemporal operator search-attribute remove \\\n    --name YourAttributeName \\\n    --yes\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringArrayVar(&s.Name, "name", nil, "Search Attribute name.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "name")
	s.Command.Flags().BoolVarP(&s.Yes, "yes", "y", false, "Don't prompt to confirm removal.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalScheduleCommand struct {
	Parent  *TemporalCommand
	Command cobra.Command
	ClientOptions
}

func NewTemporalScheduleCommand(cctx *CommandContext, parent *TemporalCommand) *TemporalScheduleCommand {
	var s TemporalScheduleCommand
	s.Parent = parent
	s.Command.Use = "schedule"
	s.Command.Short = "Perform operations on Schedules"
	if hasHighlighting {
		s.Command.Long = "Create, use, and update Schedules that allow \nWorkflow Executions to be created at specified times:\n\n\x1b[1mtemporal schedule [commands] [options]\x1b[0m\n\nFor example:\n\n\x1b[1mtemporal schedule describe \\\n    --schedule-id \"YourScheduleId\"\x1b[0m"
	} else {
		s.Command.Long = "Create, use, and update Schedules that allow \nWorkflow Executions to be created at specified times:\n\n```\ntemporal schedule [commands] [options]\n```\n\nFor example:\n\n```\ntemporal schedule describe \\\n    --schedule-id \"YourScheduleId\"\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.AddCommand(&NewTemporalScheduleBackfillCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalScheduleCreateCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalScheduleDeleteCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalScheduleDescribeCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalScheduleListCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalScheduleToggleCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalScheduleTriggerCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalScheduleUpdateCommand(cctx, &s).Command)
	s.ClientOptions.buildFlags(cctx, s.Command.PersistentFlags())
	return &s
}

type OverlapPolicyOptions struct {
	OverlapPolicy StringEnum
}

func (v *OverlapPolicyOptions) buildFlags(cctx *CommandContext, f *pflag.FlagSet) {
	v.OverlapPolicy = NewStringEnum([]string{"Skip", "BufferOne", "BufferAll", "CancelOther", "TerminateOther", "AllowAll"}, "Skip")
	f.Var(&v.OverlapPolicy, "overlap-policy", "Overlap policy. Accepted values: Skip, BufferOne, BufferAll, CancelOther, TerminateOther, AllowAll.")
}

type ScheduleIdOptions struct {
	ScheduleId string
}

func (v *ScheduleIdOptions) buildFlags(cctx *CommandContext, f *pflag.FlagSet) {
	f.StringVarP(&v.ScheduleId, "schedule-id", "s", "", "Schedule ID.")
	_ = cobra.MarkFlagRequired(f, "schedule-id")
}

type TemporalScheduleBackfillCommand struct {
	Parent  *TemporalScheduleCommand
	Command cobra.Command
	OverlapPolicyOptions
	ScheduleIdOptions
	EndTime   Timestamp
	StartTime Timestamp
}

func NewTemporalScheduleBackfillCommand(cctx *CommandContext, parent *TemporalScheduleCommand) *TemporalScheduleBackfillCommand {
	var s TemporalScheduleBackfillCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "backfill [flags]"
	s.Command.Short = "Backfill past actions"
	if hasHighlighting {
		s.Command.Long = "Batch-execute actions that would have run during a specified time interval.\nUse this command to fill in Workflow Runs from when a Schedule\nwas paused, before a Schedule was created, from the future, or to re-process a\npreviously executed interval.\n\nBackfills require a Schedule ID and the time period covered by the request.\nIt's best to use the \x1b[1mBufferAll\x1b[0m or \x1b[1mAllowAll\x1b[0m policies to avoid conflicts\nand ensure no Workflow Executions are\nskipped.\n\nFor example:\n\n\x1b[1m  temporal schedule backfill \\\n    --schedule-id \"YourScheduleId\" \\\n    --start-time \"2022-05-01T00:00:00Z\" \\\n    --end-time \"2022-05-31T23:59:59Z\" \\\n    --overlap-policy BufferAll\x1b[0m\n\nThe policies include:\n\n* **AllowAll**: Allow unlimited concurrent Workflow Executions. This\n  significantly speeds up the backfilling process on systems that support\n  concurrency. Ensure running Workflow Executions do not interfere with each\n  other.\n* **BufferAll**: Buffer all incoming Workflow Executions while waiting for the\n  running Workflow Execution to complete.\n* **Skip**: If a previous Workflow Execution is still running, discard new\n  Workflow Executions.\n* **BufferOne**: Same as 'Skip' but buffer a single Workflow Execution to be\n  run after the previous Execution completes. Discard other Workflow Executions.\n* **CancelOther**: Cancel the running Workflow Execution and replace it with\n  the incoming new Workflow Execution.\n* **TerminateOther**: Terminate the running Workflow Execution and replace\n  it with the incoming new Workflow Execution."
	} else {
		s.Command.Long = "Batch-execute actions that would have run during a specified time interval.\nUse this command to fill in Workflow Runs from when a Schedule\nwas paused, before a Schedule was created, from the future, or to re-process a\npreviously executed interval.\n\nBackfills require a Schedule ID and the time period covered by the request.\nIt's best to use the `BufferAll` or `AllowAll` policies to avoid conflicts\nand ensure no Workflow Executions are\nskipped.\n\nFor example:\n\n```\n  temporal schedule backfill \\\n    --schedule-id \"YourScheduleId\" \\\n    --start-time \"2022-05-01T00:00:00Z\" \\\n    --end-time \"2022-05-31T23:59:59Z\" \\\n    --overlap-policy BufferAll\n```\n\nThe policies include:\n\n* **AllowAll**: Allow unlimited concurrent Workflow Executions. This\n  significantly speeds up the backfilling process on systems that support\n  concurrency. Ensure running Workflow Executions do not interfere with each\n  other.\n* **BufferAll**: Buffer all incoming Workflow Executions while waiting for the\n  running Workflow Execution to complete.\n* **Skip**: If a previous Workflow Execution is still running, discard new\n  Workflow Executions.\n* **BufferOne**: Same as 'Skip' but buffer a single Workflow Execution to be\n  run after the previous Execution completes. Discard other Workflow Executions.\n* **CancelOther**: Cancel the running Workflow Execution and replace it with\n  the incoming new Workflow Execution.\n* **TerminateOther**: Terminate the running Workflow Execution and replace\n  it with the incoming new Workflow Execution."
	}
	s.Command.Args = cobra.NoArgs
	s.OverlapPolicyOptions.buildFlags(cctx, s.Command.Flags())
	s.ScheduleIdOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Flags().Var(&s.EndTime, "end-time", "Backfill end time.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "end-time")
	s.Command.Flags().Var(&s.StartTime, "start-time", "Backfill start time.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "start-time")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type ScheduleConfigurationOptions struct {
	Calendar                []string
	CatchupWindow           time.Duration
	Cron                    []string
	EndTime                 Timestamp
	Interval                []string
	Jitter                  time.Duration
	Notes                   string
	Paused                  bool
	PauseOnFailure          bool
	RemainingActions        int
	StartTime               Timestamp
	TimeZone                string
	ScheduleSearchAttribute []string
	ScheduleMemo            []string
}

func (v *ScheduleConfigurationOptions) buildFlags(cctx *CommandContext, f *pflag.FlagSet) {
	f.StringArrayVar(&v.Calendar, "calendar", nil, "[Calendar specification](https://docs.temporal.io/workflows#spec) in JSON. For example: `{\"dayOfWeek\":\"Fri\",\"hour\":\"17\",\"minute\":\"5\"}`.")
	f.DurationVar(&v.CatchupWindow, "catchup-window", 0, "Maximum catch-up time for when the Service is unavailable.")
	f.StringArrayVar(&v.Cron, "cron", nil, "Calendar specification in [cron string format](https://docs.temporal.io/workflows#cron-schedules). For example: `\"30 12 * * Fri\"`.")
	f.Var(&v.EndTime, "end-time", "Schedule end time")
	f.StringArrayVar(&v.Interval, "interval", nil, "[Interval duration](https://docs.temporal.io/workflows#spec). For example, 90m, or 60m/15m to include phase offset.")
	f.DurationVar(&v.Jitter, "jitter", 0, "Max difference in time from the specification. Vary the start time randomly within this amount.")
	f.StringVar(&v.Notes, "notes", "", "Initial notes field value.")
	f.BoolVar(&v.Paused, "paused", false, "Pause the Schedule immediately on creation.")
	f.BoolVar(&v.PauseOnFailure, "pause-on-failure", false, "Pause schedule after Workflow failures.")
	f.IntVar(&v.RemainingActions, "remaining-actions", 0, "Total allowed actions. Default is zero (unlimited).")
	f.Var(&v.StartTime, "start-time", "Schedule start time.")
	f.StringVar(&v.TimeZone, "time-zone", "", "Interpret calendar specs with the `TZ` time zone. For a list of time zones, see: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones.")
	f.StringArrayVar(&v.ScheduleSearchAttribute, "schedule-search-attribute", nil, "Set schedule Search Attributes using 'KEY=\"VALUE\"' format. KEY is a string, VALUE is JSON. Can be passed multiple times.")
	f.StringArrayVar(&v.ScheduleMemo, "schedule-memo", nil, "Set a schedule memo using 'KEY=\"VALUE\"' format. KEY is a string, VALUE is JSON. Can be passed multiple times.")
}

type TemporalScheduleCreateCommand struct {
	Parent  *TemporalScheduleCommand
	Command cobra.Command
	ScheduleConfigurationOptions
	ScheduleIdOptions
	OverlapPolicyOptions
	SharedWorkflowStartOptions
	PayloadInputOptions
}

func NewTemporalScheduleCreateCommand(cctx *CommandContext, parent *TemporalScheduleCommand) *TemporalScheduleCreateCommand {
	var s TemporalScheduleCreateCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "create [flags]"
	s.Command.Short = "Create a new Schedule"
	if hasHighlighting {
		s.Command.Long = "Create a new Schedule on the Temporal Service. A Schedule automatically\nstarts new Workflow Executions at the times you specify.\n\nFor example:\n\n\x1b[1m  temporal schedule create \\\n    --schedule-id \"YourScheduleId\" \\\n    --calendar '{\"dayOfWeek\":\"Fri\",\"hour\":\"3\",\"minute\":\"30\"}' \\\n    --workflow-id \"YourBaseWorkflowId\" \\\n    --task-queue \"YourTaskQueue\" \\\n    --workflow-type \"YourWorkflowType\"\x1b[0m\n\nSchedules support any combination of \x1b[1m--calendar\x1b[0m, \x1b[1m--interval\x1b[0m, and \x1b[1m--cron\x1b[0m:\n\n* Shorthand \x1b[1m--interval\x1b[0m strings.\n  For example: 45m (every 45 minutes) or 6h/5h (every 6 hours, at the top of\n  the 5th hour).\n* JSON \x1b[1m--calendar\x1b[0m, as in the\n  preceding example.\n* Unix-style \x1b[1m--cron\x1b[0m strings and robfig\n  declarations (@daily/@weekly/@every X/etc). For example, every Friday at \n  12:30 PM: \x1b[1m30 12 * * Fri\x1b[0m."
	} else {
		s.Command.Long = "Create a new Schedule on the Temporal Service. A Schedule automatically\nstarts new Workflow Executions at the times you specify.\n\nFor example:\n\n```\n  temporal schedule create \\\n    --schedule-id \"YourScheduleId\" \\\n    --calendar '{\"dayOfWeek\":\"Fri\",\"hour\":\"3\",\"minute\":\"30\"}' \\\n    --workflow-id \"YourBaseWorkflowId\" \\\n    --task-queue \"YourTaskQueue\" \\\n    --workflow-type \"YourWorkflowType\"\n```\n\nSchedules support any combination of `--calendar`, `--interval`, and `--cron`:\n\n* Shorthand `--interval` strings.\n  For example: 45m (every 45 minutes) or 6h/5h (every 6 hours, at the top of\n  the 5th hour).\n* JSON `--calendar`, as in the\n  preceding example.\n* Unix-style `--cron` strings and robfig\n  declarations (@daily/@weekly/@every X/etc). For example, every Friday at \n  12:30 PM: `30 12 * * Fri`."
	}
	s.Command.Args = cobra.NoArgs
	s.ScheduleConfigurationOptions.buildFlags(cctx, s.Command.Flags())
	s.ScheduleIdOptions.buildFlags(cctx, s.Command.Flags())
	s.OverlapPolicyOptions.buildFlags(cctx, s.Command.Flags())
	s.SharedWorkflowStartOptions.buildFlags(cctx, s.Command.Flags())
	s.PayloadInputOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalScheduleDeleteCommand struct {
	Parent  *TemporalScheduleCommand
	Command cobra.Command
	ScheduleIdOptions
}

func NewTemporalScheduleDeleteCommand(cctx *CommandContext, parent *TemporalScheduleCommand) *TemporalScheduleDeleteCommand {
	var s TemporalScheduleDeleteCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "delete [flags]"
	s.Command.Short = "Remove a Schedule"
	if hasHighlighting {
		s.Command.Long = "Deletes a Schedule on the front end Service:\n\n\x1b[1mtemporal schedule delete \\\n    --schedule-id YourScheduleId\x1b[0m\n\nRemoving a Schedule won't affect the Workflow Executions it started\nthat are still running. To cancel or terminate these Workflow Executions, use \n\x1b[1mtemporal workflow delete\x1b[0m with the \x1b[1mTemporalScheduledById\x1b[0m Search Attribute instead."
	} else {
		s.Command.Long = "Deletes a Schedule on the front end Service:\n\n```\ntemporal schedule delete \\\n    --schedule-id YourScheduleId\n```\n\nRemoving a Schedule won't affect the Workflow Executions it started\nthat are still running. To cancel or terminate these Workflow Executions, use \n`temporal workflow delete` with the `TemporalScheduledById` Search Attribute instead."
	}
	s.Command.Args = cobra.NoArgs
	s.ScheduleIdOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalScheduleDescribeCommand struct {
	Parent  *TemporalScheduleCommand
	Command cobra.Command
	ScheduleIdOptions
}

func NewTemporalScheduleDescribeCommand(cctx *CommandContext, parent *TemporalScheduleCommand) *TemporalScheduleDescribeCommand {
	var s TemporalScheduleDescribeCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "describe [flags]"
	s.Command.Short = "Get Schedule configuration and current state"
	if hasHighlighting {
		s.Command.Long = "Show a Schedule configuration, including information about past,\ncurrent, and future Workflow runs:\n\n\x1b[1mtemporal schedule describe \\\n    --schedule-id YourScheduleId\x1b[0m"
	} else {
		s.Command.Long = "Show a Schedule configuration, including information about past,\ncurrent, and future Workflow runs:\n\n```\ntemporal schedule describe \\\n    --schedule-id YourScheduleId\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.ScheduleIdOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalScheduleListCommand struct {
	Parent     *TemporalScheduleCommand
	Command    cobra.Command
	Long       bool
	ReallyLong bool
}

func NewTemporalScheduleListCommand(cctx *CommandContext, parent *TemporalScheduleCommand) *TemporalScheduleListCommand {
	var s TemporalScheduleListCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "list [flags]"
	s.Command.Short = "Display hosted Schedules"
	if hasHighlighting {
		s.Command.Long = "Lists the Schedules hosted by a Namespace:\n\n\x1b[1mtemporal schedule list \\\n    --namespace YourNamespace\x1b[0m"
	} else {
		s.Command.Long = "Lists the Schedules hosted by a Namespace:\n\n```\ntemporal schedule list \\\n    --namespace YourNamespace\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().BoolVarP(&s.Long, "long", "l", false, "Show detailed information")
	s.Command.Flags().BoolVar(&s.ReallyLong, "really-long", false, "Show extensive information in non-table form.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalScheduleToggleCommand struct {
	Parent  *TemporalScheduleCommand
	Command cobra.Command
	ScheduleIdOptions
	Pause   bool
	Reason  string
	Unpause bool
}

func NewTemporalScheduleToggleCommand(cctx *CommandContext, parent *TemporalScheduleCommand) *TemporalScheduleToggleCommand {
	var s TemporalScheduleToggleCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "toggle [flags]"
	s.Command.Short = "Pause or unpause a Schedule"
	if hasHighlighting {
		s.Command.Long = "Pause or unpause a Schedule by passing a flag with your\ndesired state:\n\n\x1b[1mtemporal schedule toggle \\\n    --schedule-id \"YourScheduleId\" \\\n    --pause \\\n    --reason \"YourReason\"\x1b[0m\n\nand\n\n\x1b[1mtemporal schedule toggle\n    --schedule-id \"YourScheduleId\" \\\n    --unpause \\\n    --reason \"YourReason\"\x1b[0m\n\nThe \x1b[1m--reason\x1b[0m text updates the Schedule's \x1b[1mnotes\x1b[0m field for operations \ncommunication. It defaults to \"(no reason provided)\" if omitted. This field\nis also visible on the Service Web UI."
	} else {
		s.Command.Long = "Pause or unpause a Schedule by passing a flag with your\ndesired state:\n\n```\ntemporal schedule toggle \\\n    --schedule-id \"YourScheduleId\" \\\n    --pause \\\n    --reason \"YourReason\"\n```\n\nand\n\n```\ntemporal schedule toggle\n    --schedule-id \"YourScheduleId\" \\\n    --unpause \\\n    --reason \"YourReason\"\n```\n\nThe `--reason` text updates the Schedule's `notes` field for operations \ncommunication. It defaults to \"(no reason provided)\" if omitted. This field\nis also visible on the Service Web UI."
	}
	s.Command.Args = cobra.NoArgs
	s.ScheduleIdOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Flags().BoolVar(&s.Pause, "pause", false, "Pause the schedule.")
	s.Command.Flags().StringVar(&s.Reason, "reason", "\"(no reason provided)", "Reason for pausing or unpausing a Schedule.")
	s.Command.Flags().BoolVar(&s.Unpause, "unpause", false, "Unpause the schedule.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalScheduleTriggerCommand struct {
	Parent  *TemporalScheduleCommand
	Command cobra.Command
	ScheduleIdOptions
	OverlapPolicyOptions
}

func NewTemporalScheduleTriggerCommand(cctx *CommandContext, parent *TemporalScheduleCommand) *TemporalScheduleTriggerCommand {
	var s TemporalScheduleTriggerCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "trigger [flags]"
	s.Command.Short = "Immediately run a Schedule"
	if hasHighlighting {
		s.Command.Long = "Trigger a Schedule to run immediately:\n\n\x1b[1mtemporal schedule trigger \\\n    --schedule-id \"YourScheduleId\"\x1b[0m"
	} else {
		s.Command.Long = "Trigger a Schedule to run immediately:\n\n```\ntemporal schedule trigger \\\n    --schedule-id \"YourScheduleId\"\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.ScheduleIdOptions.buildFlags(cctx, s.Command.Flags())
	s.OverlapPolicyOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalScheduleUpdateCommand struct {
	Parent  *TemporalScheduleCommand
	Command cobra.Command
	ScheduleConfigurationOptions
	ScheduleIdOptions
	OverlapPolicyOptions
	SharedWorkflowStartOptions
	PayloadInputOptions
}

func NewTemporalScheduleUpdateCommand(cctx *CommandContext, parent *TemporalScheduleCommand) *TemporalScheduleUpdateCommand {
	var s TemporalScheduleUpdateCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "update [flags]"
	s.Command.Short = "Update Schedule details"
	if hasHighlighting {
		s.Command.Long = "Update an existing Schedule with new configuration details, including\nspec, action, and policies:\n\n\x1b[1mtemporal schedule update \ntemporal schedule update \\  \n --schedule-id \"YourScheduleId\"   \n --workflow-type \"NewWorkflowType\"\x1b[0m"
	} else {
		s.Command.Long = "Update an existing Schedule with new configuration details, including\nspec, action, and policies:\n\n```\ntemporal schedule update \ntemporal schedule update \\  \n --schedule-id \"YourScheduleId\"   \n --workflow-type \"NewWorkflowType\"\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.ScheduleConfigurationOptions.buildFlags(cctx, s.Command.Flags())
	s.ScheduleIdOptions.buildFlags(cctx, s.Command.Flags())
	s.OverlapPolicyOptions.buildFlags(cctx, s.Command.Flags())
	s.SharedWorkflowStartOptions.buildFlags(cctx, s.Command.Flags())
	s.PayloadInputOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalServerCommand struct {
	Parent  *TemporalCommand
	Command cobra.Command
}

func NewTemporalServerCommand(cctx *CommandContext, parent *TemporalCommand) *TemporalServerCommand {
	var s TemporalServerCommand
	s.Parent = parent
	s.Command.Use = "server"
	s.Command.Short = "Run Temporal Server"
	if hasHighlighting {
		s.Command.Long = "Run a development Temporal Server on your local system.\nView the Web UI for the default configuration at http://localhost:8233:\n\n\x1b[1mtemporal server start-dev\x1b[0m\n\nAdd persistence for Workflow Executions across runs:\n\n\x1b[1mtemporal server start-dev \\\n    --db-filename path-to-your-local-persistent-store\x1b[0m\n\nSet the port from the front-end gRPC Service (7233 default):\n\n\x1b[1mtemporal server start-dev \\\n    --port 7234 \\\n    --ui-port 8234 \\\n    --metrics-port 57271\x1b[0m\n\nUse a custom port for the Web UI. The default is the gRPC port (7233 default)\nplus 1000 (8233):\n\n\x1b[1mtemporal server start-dev \\\n    --ui-port 3000\x1b[0m"
	} else {
		s.Command.Long = "Run a development Temporal Server on your local system.\nView the Web UI for the default configuration at http://localhost:8233:\n\n```\ntemporal server start-dev\n```\n\nAdd persistence for Workflow Executions across runs:\n\n```\ntemporal server start-dev \\\n    --db-filename path-to-your-local-persistent-store\n```\n\nSet the port from the front-end gRPC Service (7233 default):\n\n```\ntemporal server start-dev \\\n    --port 7234 \\\n    --ui-port 8234 \\\n    --metrics-port 57271\n```\n\nUse a custom port for the Web UI. The default is the gRPC port (7233 default)\nplus 1000 (8233):\n\n```\ntemporal server start-dev \\\n    --ui-port 3000\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.AddCommand(&NewTemporalServerStartDevCommand(cctx, &s).Command)
	return &s
}

type TemporalServerStartDevCommand struct {
	Parent             *TemporalServerCommand
	Command            cobra.Command
	DbFilename         string
	Namespace          []string
	Port               int
	HttpPort           int
	MetricsPort        int
	UiPort             int
	Headless           bool
	Ip                 string
	UiIp               string
	UiAssetPath        string
	UiCodecEndpoint    string
	SqlitePragma       []string
	DynamicConfigValue []string
	LogConfig          bool
}

func NewTemporalServerStartDevCommand(cctx *CommandContext, parent *TemporalServerCommand) *TemporalServerStartDevCommand {
	var s TemporalServerStartDevCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "start-dev [flags]"
	s.Command.Short = "Start Temporal development server"
	if hasHighlighting {
		s.Command.Long = "Run a development Temporal Server on your local system.\nView the Web UI for the default configuration at http://localhost:8233:\n\n\x1b[1mtemporal server start-dev\x1b[0m\n\nAdd persistence for Workflow Executions across runs:\n\n\x1b[1mtemporal server start-dev \\\n    --db-filename path-to-your-local-persistent-store\x1b[0m\n\nSet the port from the front-end gRPC Service (7233 default):\n\n\x1b[1mtemporal server start-dev \\\n    --port 7000\x1b[0m\n\nUse a custom port for the Web UI. The default is the gRPC port (7233 default)\nplus 1000 (8233):\n\n\x1b[1mtemporal server start-dev \\\n    --ui-port 3000\x1b[0m"
	} else {
		s.Command.Long = "Run a development Temporal Server on your local system.\nView the Web UI for the default configuration at http://localhost:8233:\n\n```\ntemporal server start-dev\n```\n\nAdd persistence for Workflow Executions across runs:\n\n```\ntemporal server start-dev \\\n    --db-filename path-to-your-local-persistent-store\n```\n\nSet the port from the front-end gRPC Service (7233 default):\n\n```\ntemporal server start-dev \\\n    --port 7000\n```\n\nUse a custom port for the Web UI. The default is the gRPC port (7233 default)\nplus 1000 (8233):\n\n```\ntemporal server start-dev \\\n    --ui-port 3000\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringVarP(&s.DbFilename, "db-filename", "f", "", "Path to file for persistent Temporal state store. By default, Workflow Executions are lost when the server process dies.")
	s.Command.Flags().StringArrayVarP(&s.Namespace, "namespace", "n", nil, "Namespaces to be created at launch. The \"default\" Namespace is always created automatically.")
	s.Command.Flags().IntVarP(&s.Port, "port", "p", 7233, "Port for the frontend gRPC Service.")
	s.Command.Flags().IntVar(&s.HttpPort, "http-port", 0, "Port for the HTTP API service. Default is off.")
	s.Command.Flags().IntVar(&s.MetricsPort, "metrics-port", 0, "Port for '/metrics'. Default is off.")
	s.Command.Flags().IntVar(&s.UiPort, "ui-port", 0, "Port for the Web UI. Default is '--port' value + 1000.")
	s.Command.Flags().BoolVar(&s.Headless, "headless", false, "Disable the Web UI.")
	s.Command.Flags().StringVar(&s.Ip, "ip", "localhost", "IP address bound to the frontend Service.")
	s.Command.Flags().StringVar(&s.UiIp, "ui-ip", "", "IP address bound to the WebUI. Default is same as '--ip' value.")
	s.Command.Flags().StringVar(&s.UiAssetPath, "ui-asset-path", "", "UI custom assets path.")
	s.Command.Flags().StringVar(&s.UiCodecEndpoint, "ui-codec-endpoint", "", "UI remote codec HTTP endpoint.")
	s.Command.Flags().StringArrayVar(&s.SqlitePragma, "sqlite-pragma", nil, "SQLite pragma statements in \"PRAGMA=VALUE\" format.")
	s.Command.Flags().StringArrayVar(&s.DynamicConfigValue, "dynamic-config-value", nil, "Dynamic configuration value in 'KEY=\"VALUE\"' format. KEY is a string, VALUE is JSON. For example, \"YourKey=\\\"YourStringValue\\\"\".")
	s.Command.Flags().BoolVar(&s.LogConfig, "log-config", false, "Log the server config to stderr.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalTaskQueueCommand struct {
	Parent  *TemporalCommand
	Command cobra.Command
	ClientOptions
}

func NewTemporalTaskQueueCommand(cctx *CommandContext, parent *TemporalCommand) *TemporalTaskQueueCommand {
	var s TemporalTaskQueueCommand
	s.Parent = parent
	s.Command.Use = "task-queue"
	s.Command.Short = "Manage Task Queues"
	if hasHighlighting {
		s.Command.Long = "Inspect and update Task Queues, the queues that\nWorkers poll for Workflow and Activity tasks:\n\n\x1b[1mtemporal task-queue [command] [command options]\x1b[0m\n\nFor example:\n\n\x1b[1mtemporal task-queue describe \\\n    --task-queue \"YourTaskQueueName\"\x1b[0m"
	} else {
		s.Command.Long = "Inspect and update Task Queues, the queues that\nWorkers poll for Workflow and Activity tasks:\n\n```\ntemporal task-queue [command] [command options]\n```\n\nFor example:\n\n```\ntemporal task-queue describe \\\n    --task-queue \"YourTaskQueueName\"\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.AddCommand(&NewTemporalTaskQueueDescribeCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalTaskQueueGetBuildIdReachabilityCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalTaskQueueGetBuildIdsCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalTaskQueueListPartitionCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalTaskQueueUpdateBuildIdsCommand(cctx, &s).Command)
	s.ClientOptions.buildFlags(cctx, s.Command.PersistentFlags())
	return &s
}

type TemporalTaskQueueDescribeCommand struct {
	Parent        *TemporalTaskQueueCommand
	Command       cobra.Command
	TaskQueue     string
	TaskQueueType StringEnum
	Partitions    int
}

func NewTemporalTaskQueueDescribeCommand(cctx *CommandContext, parent *TemporalTaskQueueCommand) *TemporalTaskQueueDescribeCommand {
	var s TemporalTaskQueueDescribeCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "describe [flags]"
	s.Command.Short = "Show active Workers"
	if hasHighlighting {
		s.Command.Long = "Display a list of active Workers that have recently polled\na Task Queue. The Temporal Server records each poll\nrequest time. A \x1b[1mLastAccessTime\x1b[0m over one minute may indicate the Worker is\nat capacity or has shut down. Temporal Workers are removed if 5 minutes\nhave passed since the last poll request:\n\n\x1b[1mtemporal task-queue describe \\\n   --task-queue \"YourTaskQueueName\"\x1b[0m\n\nWorkflow and Activity polling use separate Task Queues:\n\n\x1b[1mtemporal task-queue describe \\\n    --task-queue YourTaskQueue \\\n    --task-queue-type \"activity\"`\x1b[0m"
	} else {
		s.Command.Long = "Display a list of active Workers that have recently polled\na Task Queue. The Temporal Server records each poll\nrequest time. A `LastAccessTime` over one minute may indicate the Worker is\nat capacity or has shut down. Temporal Workers are removed if 5 minutes\nhave passed since the last poll request:\n\n```\ntemporal task-queue describe \\\n   --task-queue \"YourTaskQueueName\"\n```\n\nWorkflow and Activity polling use separate Task Queues:\n\n```\ntemporal task-queue describe \\\n    --task-queue YourTaskQueue \\\n    --task-queue-type \"activity\"`\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringVarP(&s.TaskQueue, "task-queue", "t", "", "Task queue name.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "task-queue")
	s.TaskQueueType = NewStringEnum([]string{"workflow", "activity"}, "workflow")
	s.Command.Flags().Var(&s.TaskQueueType, "task-queue-type", "Task Queue type. Accepted values: workflow, activity.")
	s.Command.Flags().IntVar(&s.Partitions, "partitions", 0, "Query partitions 1 through `N`. Experimental/Temporary feature. Default: 1")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalTaskQueueGetBuildIdReachabilityCommand struct {
	Parent           *TemporalTaskQueueCommand
	Command          cobra.Command
	BuildId          []string
	ReachabilityType StringEnum
	TaskQueue        []string
}

func NewTemporalTaskQueueGetBuildIdReachabilityCommand(cctx *CommandContext, parent *TemporalTaskQueueCommand) *TemporalTaskQueueGetBuildIdReachabilityCommand {
	var s TemporalTaskQueueGetBuildIdReachabilityCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "get-build-id-reachability [flags]"
	s.Command.Short = "Build ID availability"
	if hasHighlighting {
		s.Command.Long = "Show if a given Build ID can be used for new, existing, or closed workflows\nin Namespaces that support Worker versioning:\n\n\x1b[1mtemporal task-queue get-build-id-reachability \\\n    --build-id \"2.0\"\x1b[0m\n\nYou can specify the \x1b[1m--build-id\x1b[0m and \x1b[1m--task-queue\x1b[0m flags multiple times.\nIf \x1b[1m--task-queue\x1b[0m is omitted, the command checks Build ID reachability\nagainst all Task Queues."
	} else {
		s.Command.Long = "Show if a given Build ID can be used for new, existing, or closed workflows\nin Namespaces that support Worker versioning:\n\n```\ntemporal task-queue get-build-id-reachability \\\n    --build-id \"2.0\"\n```\n\nYou can specify the `--build-id` and `--task-queue` flags multiple times.\nIf `--task-queue` is omitted, the command checks Build ID reachability\nagainst all Task Queues."
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringArrayVar(&s.BuildId, "build-id", nil, "One or more Build ID strings. Can be passed multiple times.")
	s.ReachabilityType = NewStringEnum([]string{"open", "closed", "existing"}, "existing")
	s.Command.Flags().Var(&s.ReachabilityType, "reachability-type", "Reachability filter. `open`: reachable by one or more open workflows. `closed`: reachable by one or more closed workflows. `existing`: reachable by either. New Workflow Executions reachable by a Build ID are always reported. Accepted values: open, closed, existing.")
	s.Command.Flags().StringArrayVarP(&s.TaskQueue, "task-queue", "t", nil, "Search only the specified task queue(s). Can be passed multiple times.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalTaskQueueGetBuildIdsCommand struct {
	Parent    *TemporalTaskQueueCommand
	Command   cobra.Command
	TaskQueue string
	MaxSets   int
}

func NewTemporalTaskQueueGetBuildIdsCommand(cctx *CommandContext, parent *TemporalTaskQueueCommand) *TemporalTaskQueueGetBuildIdsCommand {
	var s TemporalTaskQueueGetBuildIdsCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "get-build-ids [flags]"
	s.Command.Short = "Build ID versions"
	if hasHighlighting {
		s.Command.Long = "Fetch sets of compatible Build IDs for specified Task Queues and\ndisplay their information:\n\n\x1b[1mtemporal task-queue get-build-ids \\\n    --task-queue \"YourTaskQueue\"\x1b[0m\n\nThis command is limited to Namespaces that enable Worker versioning."
	} else {
		s.Command.Long = "Fetch sets of compatible Build IDs for specified Task Queues and\ndisplay their information:\n\n```\ntemporal task-queue get-build-ids \\\n    --task-queue \"YourTaskQueue\"\n```\n\nThis command is limited to Namespaces that enable Worker versioning."
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringVarP(&s.TaskQueue, "task-queue", "t", "", "Task queue name.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "task-queue")
	s.Command.Flags().IntVar(&s.MaxSets, "max-sets", 0, "Max return count. Use 1 for default major version. Use 0 for all sets.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalTaskQueueListPartitionCommand struct {
	Parent    *TemporalTaskQueueCommand
	Command   cobra.Command
	TaskQueue string
}

func NewTemporalTaskQueueListPartitionCommand(cctx *CommandContext, parent *TemporalTaskQueueCommand) *TemporalTaskQueueListPartitionCommand {
	var s TemporalTaskQueueListPartitionCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "list-partition [flags]"
	s.Command.Short = "List partitions"
	if hasHighlighting {
		s.Command.Long = "Display a Task Queue's partition list with assigned matching nodes:\n\n\x1b[1mtemporal task-queue list-partition \\\n    --task-queue \"YourTaskQueue\"\x1b[0m"
	} else {
		s.Command.Long = "Display a Task Queue's partition list with assigned matching nodes:\n\n```\ntemporal task-queue list-partition \\\n    --task-queue \"YourTaskQueue\"\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringVarP(&s.TaskQueue, "task-queue", "t", "", "Task Queue name. Required")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalTaskQueueUpdateBuildIdsCommand struct {
	Parent  *TemporalTaskQueueCommand
	Command cobra.Command
}

func NewTemporalTaskQueueUpdateBuildIdsCommand(cctx *CommandContext, parent *TemporalTaskQueueCommand) *TemporalTaskQueueUpdateBuildIdsCommand {
	var s TemporalTaskQueueUpdateBuildIdsCommand
	s.Parent = parent
	s.Command.Use = "update-build-ids"
	s.Command.Short = "Manage build IDs"
	s.Command.Long = "Add or change a Task Queue's compatible Build IDs for Namespaces using\nWorker versioning."
	s.Command.Args = cobra.NoArgs
	s.Command.AddCommand(&NewTemporalTaskQueueUpdateBuildIdsAddNewCompatibleCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalTaskQueueUpdateBuildIdsAddNewDefaultCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalTaskQueueUpdateBuildIdsPromoteIdInSetCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalTaskQueueUpdateBuildIdsPromoteSetCommand(cctx, &s).Command)
	return &s
}

type TemporalTaskQueueUpdateBuildIdsAddNewCompatibleCommand struct {
	Parent                    *TemporalTaskQueueUpdateBuildIdsCommand
	Command                   cobra.Command
	BuildId                   string
	TaskQueue                 string
	ExistingCompatibleBuildId string
	SetAsDefault              bool
}

func NewTemporalTaskQueueUpdateBuildIdsAddNewCompatibleCommand(cctx *CommandContext, parent *TemporalTaskQueueUpdateBuildIdsCommand) *TemporalTaskQueueUpdateBuildIdsAddNewCompatibleCommand {
	var s TemporalTaskQueueUpdateBuildIdsAddNewCompatibleCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "add-new-compatible [flags]"
	s.Command.Short = "Add compatible build ID"
	if hasHighlighting {
		s.Command.Long = "Add a compatible Build ID to a Task Queue's existing version set.\nProvide an existing Build ID and a new Build ID:\n\n\x1b[1mtemporal task-queue update-build-ids add-new-compatible \\\n    --task-queue YourTaskQueue \\\n    --existing-compatible-build-id YourExistingBuildId \\\n    --build-id YourNewBuildId\x1b[0m\n\nThe new ID is stored in the set containing the existing ID and becomes the\nnew default for that set.\n\nThis command is limited to Namespaces that support Worker versioning."
	} else {
		s.Command.Long = "Add a compatible Build ID to a Task Queue's existing version set.\nProvide an existing Build ID and a new Build ID:\n\n```\ntemporal task-queue update-build-ids add-new-compatible \\\n    --task-queue YourTaskQueue \\\n    --existing-compatible-build-id YourExistingBuildId \\\n    --build-id YourNewBuildId\n```\n\nThe new ID is stored in the set containing the existing ID and becomes the\nnew default for that set.\n\nThis command is limited to Namespaces that support Worker versioning."
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringVar(&s.BuildId, "build-id", "", "Build Id to be added. Required")
	s.Command.Flags().StringVarP(&s.TaskQueue, "task-queue", "t", "", "Task Queue name.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "task-queue")
	s.Command.Flags().StringVar(&s.ExistingCompatibleBuildId, "existing-compatible-build-id", "", "Pre-existing Build Id in this Task Queue.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "existing-compatible-build-id")
	s.Command.Flags().BoolVar(&s.SetAsDefault, "set-as-default", false, "Set the expanded Build Id set as the Task Queue default. Defaults to false.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalTaskQueueUpdateBuildIdsAddNewDefaultCommand struct {
	Parent    *TemporalTaskQueueUpdateBuildIdsCommand
	Command   cobra.Command
	BuildId   string
	TaskQueue string
}

func NewTemporalTaskQueueUpdateBuildIdsAddNewDefaultCommand(cctx *CommandContext, parent *TemporalTaskQueueUpdateBuildIdsCommand) *TemporalTaskQueueUpdateBuildIdsAddNewDefaultCommand {
	var s TemporalTaskQueueUpdateBuildIdsAddNewDefaultCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "add-new-default [flags]"
	s.Command.Short = "Establish new default BuildID set."
	if hasHighlighting {
		s.Command.Long = "Create a new Task Queue Build ID set, add a Build ID to it, and make it the\noverall Task Queue default. The new set will be incompatible with previous \nsets and versions.\n\n\x1b[1mtemporal task-queue update-build-ids add-new-default \\\n    --task-queue YourTaskQueue \\\n    --build-id YourNewBuildId\x1b[0m\n\nThis command is limited to Namespaces that support Worker versioning."
	} else {
		s.Command.Long = "Create a new Task Queue Build ID set, add a Build ID to it, and make it the\noverall Task Queue default. The new set will be incompatible with previous \nsets and versions.\n\n```\ntemporal task-queue update-build-ids add-new-default \\\n    --task-queue YourTaskQueue \\\n    --build-id YourNewBuildId\n```\n\nThis command is limited to Namespaces that support Worker versioning."
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringVar(&s.BuildId, "build-id", "", "Build Id to be added. Required")
	s.Command.Flags().StringVarP(&s.TaskQueue, "task-queue", "t", "", "Task Queue name.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "task-queue")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalTaskQueueUpdateBuildIdsPromoteIdInSetCommand struct {
	Parent    *TemporalTaskQueueUpdateBuildIdsCommand
	Command   cobra.Command
	BuildId   string
	TaskQueue string
}

func NewTemporalTaskQueueUpdateBuildIdsPromoteIdInSetCommand(cctx *CommandContext, parent *TemporalTaskQueueUpdateBuildIdsCommand) *TemporalTaskQueueUpdateBuildIdsPromoteIdInSetCommand {
	var s TemporalTaskQueueUpdateBuildIdsPromoteIdInSetCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "promote-id-in-set [flags]"
	s.Command.Short = "use Build ID as set default"
	if hasHighlighting {
		s.Command.Long = "Establish an existing Build ID as the default in its Task Queue set. New\ntasks compatible with this set will now be dispatched to this ID:\n\n\x1b[1mtemporal task-queue update-build-ids promote-id-in-set \\\n    --task-queue YourTaskQueue \\\n    --build-id YourBuildId\x1b[0m\n\nThis command is limited to Namespaces that support Worker versioning."
	} else {
		s.Command.Long = "Establish an existing Build ID as the default in its Task Queue set. New\ntasks compatible with this set will now be dispatched to this ID:\n\n```\ntemporal task-queue update-build-ids promote-id-in-set \\\n    --task-queue YourTaskQueue \\\n    --build-id YourBuildId\n```\n\nThis command is limited to Namespaces that support Worker versioning."
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringVar(&s.BuildId, "build-id", "", "Build Id to set as default. Required")
	s.Command.Flags().StringVarP(&s.TaskQueue, "task-queue", "t", "", "Task Queue name.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "task-queue")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalTaskQueueUpdateBuildIdsPromoteSetCommand struct {
	Parent    *TemporalTaskQueueUpdateBuildIdsCommand
	Command   cobra.Command
	BuildId   string
	TaskQueue string
}

func NewTemporalTaskQueueUpdateBuildIdsPromoteSetCommand(cctx *CommandContext, parent *TemporalTaskQueueUpdateBuildIdsCommand) *TemporalTaskQueueUpdateBuildIdsPromoteSetCommand {
	var s TemporalTaskQueueUpdateBuildIdsPromoteSetCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "promote-set [flags]"
	s.Command.Short = "Promote Build ID set"
	if hasHighlighting {
		s.Command.Long = "Promote a Build ID set to be the default on a Task Queue. Identify the\nset by providing a Build ID within it. If the set is already the default,\nthis command has no effect:\n\n\x1b[1mtemporal task-queue update-build-ids promote-set \\\n    --task-queue YourTaskQueue \\\n    --build-id YourBuildId\x1b[0m\n\nThis command is limited to Namespaces that support Worker versioning."
	} else {
		s.Command.Long = "Promote a Build ID set to be the default on a Task Queue. Identify the\nset by providing a Build ID within it. If the set is already the default,\nthis command has no effect:\n\n```\ntemporal task-queue update-build-ids promote-set \\\n    --task-queue YourTaskQueue \\\n    --build-id YourBuildId\n```\n\nThis command is limited to Namespaces that support Worker versioning."
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringVar(&s.BuildId, "build-id", "", "Build Id within the promoted set. Required")
	s.Command.Flags().StringVarP(&s.TaskQueue, "task-queue", "t", "", "Task Queue name.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "task-queue")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type ClientOptions struct {
	Address                    string
	Namespace                  string
	ApiKey                     string
	GrpcMeta                   []string
	Tls                        bool
	TlsCertPath                string
	TlsCertData                string
	TlsKeyPath                 string
	TlsKeyData                 string
	TlsCaPath                  string
	TlsCaData                  string
	TlsDisableHostVerification bool
	TlsServerName              string
	CodecEndpoint              string
	CodecAuth                  string
}

func (v *ClientOptions) buildFlags(cctx *CommandContext, f *pflag.FlagSet) {
	f.StringVar(&v.Address, "address", "127.0.0.1:7233", "Temporal Service gRPC endpoint.")
	cctx.BindFlagEnvVar(f.Lookup("address"), "TEMPORAL_ADDRESS")
	f.StringVarP(&v.Namespace, "namespace", "n", "default", "Temporal Service Namespace.")
	cctx.BindFlagEnvVar(f.Lookup("namespace"), "TEMPORAL_NAMESPACE")
	f.StringVar(&v.ApiKey, "api-key", "", "API key for request.")
	cctx.BindFlagEnvVar(f.Lookup("api-key"), "TEMPORAL_API_KE")
	f.StringArrayVar(&v.GrpcMeta, "grpc-meta", nil, "HTTP headers for requests. Format as comma-separated \"KEY=VALUE\" pairs.")
	f.BoolVar(&v.Tls, "tls", false, "Enable base TLS encryption. Does not have additional options like mTLS or client certs.")
	cctx.BindFlagEnvVar(f.Lookup("tls"), "TEMPORAL_TLS")
	f.StringVar(&v.TlsCertPath, "tls-cert-path", "", "Path to x509 certificate. Can't be used with --tls-cert-data.")
	cctx.BindFlagEnvVar(f.Lookup("tls-cert-path"), "TEMPORAL_TLS_CERT")
	f.StringVar(&v.TlsCertData, "tls-cert-data", "", "Data for x509 certificate. Can't be used with --tls-cert-path.")
	cctx.BindFlagEnvVar(f.Lookup("tls-cert-data"), "TEMPORAL_TLS_CERT_DATA")
	f.StringVar(&v.TlsKeyPath, "tls-key-path", "", "Path to x509 private key. Can't be used with --tls-key-data.")
	cctx.BindFlagEnvVar(f.Lookup("tls-key-path"), "TEMPORAL_TLS_KEY")
	f.StringVar(&v.TlsKeyData, "tls-key-data", "", "Private certificate key data. Can't be used with --tls-key-path.")
	cctx.BindFlagEnvVar(f.Lookup("tls-key-data"), "TEMPORAL_TLS_KEY_DATA")
	f.StringVar(&v.TlsCaPath, "tls-ca-path", "", "Path to server CA certificate. Can't be used with --tls-ca-data.")
	cctx.BindFlagEnvVar(f.Lookup("tls-ca-path"), "TEMPORAL_TLS_CA")
	f.StringVar(&v.TlsCaData, "tls-ca-data", "", "Data for server CA certificate. Can't be used with --tls-ca-path.")
	cctx.BindFlagEnvVar(f.Lookup("tls-ca-data"), "TEMPORAL_TLS_CA_DATA")
	f.BoolVar(&v.TlsDisableHostVerification, "tls-disable-host-verification", false, "Disable TLS host-name verification.")
	cctx.BindFlagEnvVar(f.Lookup("tls-disable-host-verification"), "TEMPORAL_TLS_DISABLE_HOST_VERIFICATION")
	f.StringVar(&v.TlsServerName, "tls-server-name", "", "Override target TLS server name.")
	cctx.BindFlagEnvVar(f.Lookup("tls-server-name"), "TEMPORAL_TLS_SERVER_NAME")
	f.StringVar(&v.CodecEndpoint, "codec-endpoint", "", "Remote Codec Server endpoint.")
	cctx.BindFlagEnvVar(f.Lookup("codec-endpoint"), "TEMPORAL_CODEC_ENDPOINT")
	f.StringVar(&v.CodecAuth, "codec-auth", "", "Authorization header for Codec Server requests.")
	cctx.BindFlagEnvVar(f.Lookup("codec-auth"), "TEMPORAL_CODEC_AUTH")
}

type TemporalWorkflowCommand struct {
	Parent  *TemporalCommand
	Command cobra.Command
	ClientOptions
}

func NewTemporalWorkflowCommand(cctx *CommandContext, parent *TemporalCommand) *TemporalWorkflowCommand {
	var s TemporalWorkflowCommand
	s.Parent = parent
	s.Command.Use = "workflow"
	s.Command.Short = "Start, list, and operate on Workflows"
	if hasHighlighting {
		s.Command.Long = "Workflow commands perform operations on Workflow Executions:\n\n\x1b[1mtemporal workflow [command] [options]\x1b[0m\n\nFor example:\n\n\x1b[1mtemporal workflow list\x1b[0m"
	} else {
		s.Command.Long = "Workflow commands perform operations on Workflow Executions:\n\n```\ntemporal workflow [command] [options]\n```\n\nFor example:\n\n```\ntemporal workflow list\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.AddCommand(&NewTemporalWorkflowCancelCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalWorkflowCountCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalWorkflowDeleteCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalWorkflowDescribeCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalWorkflowExecuteCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalWorkflowFixHistoryJsonCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalWorkflowListCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalWorkflowQueryCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalWorkflowResetCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalWorkflowShowCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalWorkflowSignalCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalWorkflowStackCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalWorkflowStartCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalWorkflowTerminateCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalWorkflowTraceCommand(cctx, &s).Command)
	s.Command.AddCommand(&NewTemporalWorkflowUpdateCommand(cctx, &s).Command)
	s.ClientOptions.buildFlags(cctx, s.Command.PersistentFlags())
	return &s
}

type TemporalWorkflowCancelCommand struct {
	Parent  *TemporalWorkflowCommand
	Command cobra.Command
	SingleWorkflowOrBatchOptions
}

func NewTemporalWorkflowCancelCommand(cctx *CommandContext, parent *TemporalWorkflowCommand) *TemporalWorkflowCancelCommand {
	var s TemporalWorkflowCancelCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "cancel [flags]"
	s.Command.Short = "Send cancellation to Workflow Execution"
	if hasHighlighting {
		s.Command.Long = "Canceling a running Workflow Execution records a \n\x1b[1mWorkflowExecutionCancelRequested\x1b[0m event in the Event History. The Service\nschedules a new Command Task, and the Workflow Execution performs any cleanup\nwork supported by its implementation.\n\nUse the Workflow ID to cancel an Execution:\n\n\x1b[1mtemporal workflow cancel \\\n    --workflow-id YourWorkflowId\x1b[0m\n\nA visibility query lets you send bulk cancellations to Workflow Executions\nmatching the results:\n\n\x1b[1mtemporal workflow cancel \\\n    --query YourQuery\x1b[0m"
	} else {
		s.Command.Long = "Canceling a running Workflow Execution records a \n`WorkflowExecutionCancelRequested` event in the Event History. The Service\nschedules a new Command Task, and the Workflow Execution performs any cleanup\nwork supported by its implementation.\n\nUse the Workflow ID to cancel an Execution:\n\n```\ntemporal workflow cancel \\\n    --workflow-id YourWorkflowId\n```\n\nA visibility query lets you send bulk cancellations to Workflow Executions\nmatching the results:\n\n```\ntemporal workflow cancel \\\n    --query YourQuery\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.SingleWorkflowOrBatchOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalWorkflowCountCommand struct {
	Parent  *TemporalWorkflowCommand
	Command cobra.Command
	Query   string
}

func NewTemporalWorkflowCountCommand(cctx *CommandContext, parent *TemporalWorkflowCommand) *TemporalWorkflowCountCommand {
	var s TemporalWorkflowCountCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "count [flags]"
	s.Command.Short = "Number of Workflow Executions"
	if hasHighlighting {
		s.Command.Long = "Show a count of Workflow Executions, regardless of execution\nstate (running, terminated, etc). Use \x1b[1m--query\x1b[0m to select a subset of Workflow\nExecutions:\n\n\x1b[1mtemporal workflow count \\\n    --query YourQuery\x1b[0m"
	} else {
		s.Command.Long = "Show a count of Workflow Executions, regardless of execution\nstate (running, terminated, etc). Use `--query` to select a subset of Workflow\nExecutions:\n\n```\ntemporal workflow count \\\n    --query YourQuery\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringVarP(&s.Query, "query", "q", "", "Content for an SQL-like `QUERY` List Filter.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalWorkflowDeleteCommand struct {
	Parent  *TemporalWorkflowCommand
	Command cobra.Command
	SingleWorkflowOrBatchOptions
}

func NewTemporalWorkflowDeleteCommand(cctx *CommandContext, parent *TemporalWorkflowCommand) *TemporalWorkflowDeleteCommand {
	var s TemporalWorkflowDeleteCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "delete [flags]"
	s.Command.Short = "Remove Workflow Execution"
	if hasHighlighting {
		s.Command.Long = "Delete a Workflow Executions and its Event History:\n\n\x1b[1mtemporal workflow delete \\\n    --workflow-id YourWorkflowId\x1b[0m\n\nThe removal executes asynchronously. If the Execution is Running, the Service\nterminates it before deletion."
	} else {
		s.Command.Long = "Delete a Workflow Executions and its Event History:\n\n```\ntemporal workflow delete \\\n    --workflow-id YourWorkflowId\n```\n\nThe removal executes asynchronously. If the Execution is Running, the Service\nterminates it before deletion."
	}
	s.Command.Args = cobra.NoArgs
	s.SingleWorkflowOrBatchOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type WorkflowReferenceOptions struct {
	WorkflowId string
	RunId      string
}

func (v *WorkflowReferenceOptions) buildFlags(cctx *CommandContext, f *pflag.FlagSet) {
	f.StringVarP(&v.WorkflowId, "workflow-id", "w", "", "Workflow ID.")
	_ = cobra.MarkFlagRequired(f, "workflow-id")
	f.StringVarP(&v.RunId, "run-id", "r", "", "Run ID.")
}

type TemporalWorkflowDescribeCommand struct {
	Parent  *TemporalWorkflowCommand
	Command cobra.Command
	WorkflowReferenceOptions
	ResetPoints bool
	Raw         bool
}

func NewTemporalWorkflowDescribeCommand(cctx *CommandContext, parent *TemporalWorkflowCommand) *TemporalWorkflowDescribeCommand {
	var s TemporalWorkflowDescribeCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "describe [flags]"
	s.Command.Short = "Show Workflow Execution info"
	if hasHighlighting {
		s.Command.Long = "Display information about a specific Workflow Execution:\n\n\x1b[1mtemporal workflow describe \\\n    --workflow-id YourWorkflowId\x1b[0m\n\nShow the Workflow Execution's auto-reset points:\n\n\x1b[1mtemporal workflow describe \\\n    --workflow-id YourWorkflowId \\\n    --reset-points true\x1b[0m"
	} else {
		s.Command.Long = "Display information about a specific Workflow Execution:\n\n```\ntemporal workflow describe \\\n    --workflow-id YourWorkflowId\n```\n\nShow the Workflow Execution's auto-reset points:\n\n```\ntemporal workflow describe \\\n    --workflow-id YourWorkflowId \\\n    --reset-points true\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.WorkflowReferenceOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Flags().BoolVar(&s.ResetPoints, "reset-points", false, "Show auto-reset points only.")
	s.Command.Flags().BoolVar(&s.Raw, "raw", false, "Print properties without changing their format. Defaults to true.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalWorkflowExecuteCommand struct {
	Parent  *TemporalWorkflowCommand
	Command cobra.Command
	SharedWorkflowStartOptions
	WorkflowStartOptions
	PayloadInputOptions
	EventDetails bool
}

func NewTemporalWorkflowExecuteCommand(cctx *CommandContext, parent *TemporalWorkflowCommand) *TemporalWorkflowExecuteCommand {
	var s TemporalWorkflowExecuteCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "execute [flags]"
	s.Command.Short = "Start a new Workflow Execution"
	if hasHighlighting {
		s.Command.Long = "Establish a new Workflow Execution and direct its\nprogress to stdout. The command blocks and returns when the Workflow Execution\ncompletes. If your Workflow requires input, pass valid JSON:\n\n\x1b[1mtemporal workflow execute\n\t\t--workflow-id YourWorkflowId \\\n\t\t--type YourWorkflow \\\n\t\t--task-queue YourTaskQueue \\\n\t\t--input '{\"Input\": \"As-JSON\"}'\x1b[0m\n\nUse \x1b[1m--event-details\x1b[0m to relay updates to the command-line output in JSON\nformat. When using JSON output (\x1b[1m--output json\x1b[0m), this includes the entire\n\"history\" JSON key for the run."
	} else {
		s.Command.Long = "Establish a new Workflow Execution and direct its\nprogress to stdout. The command blocks and returns when the Workflow Execution\ncompletes. If your Workflow requires input, pass valid JSON:\n\n```\ntemporal workflow execute\n\t\t--workflow-id YourWorkflowId \\\n\t\t--type YourWorkflow \\\n\t\t--task-queue YourTaskQueue \\\n\t\t--input '{\"Input\": \"As-JSON\"}'\n```\n\nUse `--event-details` to relay updates to the command-line output in JSON\nformat. When using JSON output (`--output json`), this includes the entire\n\"history\" JSON key for the run."
	}
	s.Command.Args = cobra.NoArgs
	s.SharedWorkflowStartOptions.buildFlags(cctx, s.Command.Flags())
	s.WorkflowStartOptions.buildFlags(cctx, s.Command.Flags())
	s.PayloadInputOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Flags().BoolVar(&s.EventDetails, "event-details", false, "Show event details during run.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalWorkflowFixHistoryJsonCommand struct {
	Parent  *TemporalWorkflowCommand
	Command cobra.Command
	Source  string
	Target  string
}

func NewTemporalWorkflowFixHistoryJsonCommand(cctx *CommandContext, parent *TemporalWorkflowCommand) *TemporalWorkflowFixHistoryJsonCommand {
	var s TemporalWorkflowFixHistoryJsonCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "fix-history-json [flags]"
	s.Command.Short = "Updates an event history JSON file"
	if hasHighlighting {
		s.Command.Long = "Reserialize an Event History JSON file:\n\n\x1b[1mtemporal workflow fix-history-json \\\n\t--source /path/to/original.json \\\n\t--target /path/to/reserialized.json\x1b[0m"
	} else {
		s.Command.Long = "Reserialize an Event History JSON file:\n\n```\ntemporal workflow fix-history-json \\\n\t--source /path/to/original.json \\\n\t--target /path/to/reserialized.json\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringVarP(&s.Source, "source", "s", "", "Path to the original file.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "source")
	s.Command.Flags().StringVarP(&s.Target, "target", "t", "", "Path to the results file. When omitted, output is sent to stdout.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalWorkflowListCommand struct {
	Parent   *TemporalWorkflowCommand
	Command  cobra.Command
	Query    string
	Archived bool
	Limit    int
}

func NewTemporalWorkflowListCommand(cctx *CommandContext, parent *TemporalWorkflowCommand) *TemporalWorkflowListCommand {
	var s TemporalWorkflowListCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "list [flags]"
	s.Command.Short = "Show Workflow Executions"
	if hasHighlighting {
		s.Command.Long = "List Workflow Executions. By default, this command returns\nup to 10 closed Workflow Executions. The optional \x1b[1m--query\x1b[0m limits the output\nto Workflows matching a Query:\n\n\x1b[1mtemporal workflow list \\\n    --query YourQuery`\x1b[0m\n\nView a list of archived Workflow Executions:\n\n\x1b[1mtemporal workflow list \\\n    --archived\x1b[0m"
	} else {
		s.Command.Long = "List Workflow Executions. By default, this command returns\nup to 10 closed Workflow Executions. The optional `--query` limits the output\nto Workflows matching a Query:\n\n```\ntemporal workflow list \\\n    --query YourQuery`\n```\n\nView a list of archived Workflow Executions:\n\n```\ntemporal workflow list \\\n    --archived\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringVarP(&s.Query, "query", "q", "", "Content for an SQL-like `QUERY` List Filter.")
	s.Command.Flags().BoolVar(&s.Archived, "archived", false, "Limit output to archived Workflow Executions.")
	s.Command.Flags().IntVar(&s.Limit, "limit", 0, "Maximum number of Workflow Executions to display.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalWorkflowQueryCommand struct {
	Parent  *TemporalWorkflowCommand
	Command cobra.Command
	PayloadInputOptions
	WorkflowReferenceOptions
	Type            string
	RejectCondition StringEnum
}

func NewTemporalWorkflowQueryCommand(cctx *CommandContext, parent *TemporalWorkflowCommand) *TemporalWorkflowQueryCommand {
	var s TemporalWorkflowQueryCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "query [flags]"
	s.Command.Short = "Synchronously retrieve Workflow Execution state"
	if hasHighlighting {
		s.Command.Long = "Send a Query to a Workflow Execution by\nWorkflow ID to retrieve its state. This synchronous\noperation exposes the internal state of a running Workflow Execution, which\nconstantly changes. You can query both running and completed Workflow\nExecutions:\n\n\x1b[1mtemporal workflow query \\\n    --workflow-id YourWorkflowId\n    --type YourQueryType\n    --input '{\"YourInputKey\": \"YourInputValue\"}'\x1b[0m\n\nQuery implementations must never mutate Workflow Execution state and must\nnot contain blocking code."
	} else {
		s.Command.Long = "Send a Query to a Workflow Execution by\nWorkflow ID to retrieve its state. This synchronous\noperation exposes the internal state of a running Workflow Execution, which\nconstantly changes. You can query both running and completed Workflow\nExecutions:\n\n```\ntemporal workflow query \\\n    --workflow-id YourWorkflowId\n    --type YourQueryType\n    --input '{\"YourInputKey\": \"YourInputValue\"}'\n```\n\nQuery implementations must never mutate Workflow Execution state and must\nnot contain blocking code."
	}
	s.Command.Args = cobra.NoArgs
	s.PayloadInputOptions.buildFlags(cctx, s.Command.Flags())
	s.WorkflowReferenceOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Flags().StringVar(&s.Type, "type", "", "Query Type/Name.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "type")
	s.RejectCondition = NewStringEnum([]string{"not_open", "not_completed_cleanly"}, "")
	s.Command.Flags().Var(&s.RejectCondition, "reject-condition", "Optional flag to reject Queries based on Workflow state. Accepted values: not_open, not_completed_cleanly.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalWorkflowResetCommand struct {
	Parent      *TemporalWorkflowCommand
	Command     cobra.Command
	WorkflowId  string
	RunId       string
	EventId     int
	Reason      string
	ReapplyType StringEnum
	Type        StringEnum
	BuildId     string
	Query       string
	Yes         bool
}

func NewTemporalWorkflowResetCommand(cctx *CommandContext, parent *TemporalWorkflowCommand) *TemporalWorkflowResetCommand {
	var s TemporalWorkflowResetCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "reset [flags]"
	s.Command.Short = "Move a Workflow Execution to an older history point"
	if hasHighlighting {
		s.Command.Long = "Reset a Workflow Execution so it can resume from a point in\nits Event History without losing its progress up to that point:\n\n\x1b[1mtemporal workflow reset \\\n    --workflow-id YourWorkflowId\n    --event-id YourLastEvent\x1b[0m\n\nStart from where the Workflow Execution last continued as new:\n\n\x1b[1mtemporal workflow reset \\\n    --workflow-id YourWorkflowId \\\n    --type LastContinuedAsNew\x1b[0m\n\nFor batch resets, limit your resets to FirstWorkflowTask, LastWorkflowTask,\nor BuildId. Do not use Workflow IDs, run IDs, or event IDs with this\ncommand."
	} else {
		s.Command.Long = "Reset a Workflow Execution so it can resume from a point in\nits Event History without losing its progress up to that point:\n\n```\ntemporal workflow reset \\\n    --workflow-id YourWorkflowId\n    --event-id YourLastEvent\n```\n\nStart from where the Workflow Execution last continued as new:\n\n```\ntemporal workflow reset \\\n    --workflow-id YourWorkflowId \\\n    --type LastContinuedAsNew\n```\n\nFor batch resets, limit your resets to FirstWorkflowTask, LastWorkflowTask,\nor BuildId. Do not use Workflow IDs, run IDs, or event IDs with this\ncommand."
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringVarP(&s.WorkflowId, "workflow-id", "w", "", "Workflow ID. Required for non-batch reset operations.")
	s.Command.Flags().StringVarP(&s.RunId, "run-id", "r", "", "Run ID.")
	s.Command.Flags().IntVarP(&s.EventId, "event-id", "e", 0, "Event ID to reset to. Event must occur after `WorkflowTaskStarted`. `WorkflowTaskCompleted`, `WorkflowTaskFailed`, etc are valid.")
	s.Command.Flags().StringVar(&s.Reason, "reason", "", "Reason for reset.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "reason")
	s.ReapplyType = NewStringEnum([]string{"All", "Signal", "None"}, "All")
	s.Command.Flags().Var(&s.ReapplyType, "reapply-type", "Types of events to re-apply after reset point. Accepted values: All, Signal, None.")
	s.Type = NewStringEnum([]string{"FirstWorkflowTask", "LastWorkflowTask", "LastContinuedAsNew", "BuildId"}, "")
	s.Command.Flags().VarP(&s.Type, "type", "t", "The event type for the reset. Accepted values: FirstWorkflowTask, LastWorkflowTask, LastContinuedAsNew, BuildId.")
	s.Command.Flags().StringVar(&s.BuildId, "build-id", "", "A Build ID. Use only with the BuildId `--type`. Resets the first Workflow task processed by this ID. By default, this reset may be in a prior run, earlier than a Continue as New point.")
	s.Command.Flags().StringVarP(&s.Query, "query", "q", "", "Content for an SQL-like `QUERY` List Filter.")
	s.Command.Flags().BoolVarP(&s.Yes, "yes", "y", false, "Don't prompt to confirm. Only allowed when `--query` is present.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalWorkflowShowCommand struct {
	Parent  *TemporalWorkflowCommand
	Command cobra.Command
	WorkflowReferenceOptions
	EventDetails bool
	Follow       bool
}

func NewTemporalWorkflowShowCommand(cctx *CommandContext, parent *TemporalWorkflowCommand) *TemporalWorkflowShowCommand {
	var s TemporalWorkflowShowCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "show [flags]"
	s.Command.Short = "Display Event History"
	if hasHighlighting {
		s.Command.Long = "Show a Workflow Execution's Event History. \nWhen using JSON output (\x1b[1m--output JSON\x1b[0m), you may pass the results to an SDK\nto perform a replay:\n\n\x1b[1mtemporal workflow show \\\n    --workflow-id YourWorkflowId\n    --output json\x1b[0m"
	} else {
		s.Command.Long = "Show a Workflow Execution's Event History. \nWhen using JSON output (`--output JSON`), you may pass the results to an SDK\nto perform a replay:\n\n```\ntemporal workflow show \\\n    --workflow-id YourWorkflowId\n    --output json\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.WorkflowReferenceOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Flags().BoolVar(&s.EventDetails, "event-details", false, "Show event details during run.")
	s.Command.Flags().BoolVarP(&s.Follow, "follow", "f", false, "Direct Workflow Execution progress to stdout in real time. Does not apply when JSON output is selected.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type SingleWorkflowOrBatchOptions struct {
	WorkflowId string
	Query      string
	RunId      string
	Reason     string
	Yes        bool
}

func (v *SingleWorkflowOrBatchOptions) buildFlags(cctx *CommandContext, f *pflag.FlagSet) {
	f.StringVarP(&v.WorkflowId, "workflow-id", "w", "", "Workflow ID. You must set either --workflow-id or --query.")
	f.StringVarP(&v.Query, "query", "q", "", "Content for an SQL-like `QUERY` List Filter. You must set either --workflow-id or --query.")
	f.StringVarP(&v.RunId, "run-id", "r", "", "Run ID. Only use with --workflow-id. Cannot use with --query.")
	f.StringVar(&v.Reason, "reason", "", "Reason to perform batch. Only use with --query. Defaults to user name.")
	f.BoolVarP(&v.Yes, "yes", "y", false, "Don't prompt to confirm signalling. Only allowed when --query is present.")
}

type TemporalWorkflowSignalCommand struct {
	Parent  *TemporalWorkflowCommand
	Command cobra.Command
	PayloadInputOptions
	Name string
	SingleWorkflowOrBatchOptions
}

func NewTemporalWorkflowSignalCommand(cctx *CommandContext, parent *TemporalWorkflowCommand) *TemporalWorkflowSignalCommand {
	var s TemporalWorkflowSignalCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "signal [flags]"
	s.Command.Short = "Send a message to a Workflow Execution"
	if hasHighlighting {
		s.Command.Long = "Send an asynchronous notification (Signal) to a running Workflow Execution \nby its Workflow ID. The Signal is written to the\nHistory. When you include \x1b[1m--input\x1b[0m, that data is available for the Workflow\nExecution to consume:\n\n\x1b[1mtemporal workflow signal \\\n\t\t--workflow-id YourWorkflowId \\\n\t\t--name YourSignal \\\n\t\t--input '{\"YourInputKey\": \"YourInputValue\"}'\x1b[0m"
	} else {
		s.Command.Long = "Send an asynchronous notification (Signal) to a running Workflow Execution \nby its Workflow ID. The Signal is written to the\nHistory. When you include `--input`, that data is available for the Workflow\nExecution to consume:\n\n```\ntemporal workflow signal \\\n\t\t--workflow-id YourWorkflowId \\\n\t\t--name YourSignal \\\n\t\t--input '{\"YourInputKey\": \"YourInputValue\"}'\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.PayloadInputOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Flags().StringVar(&s.Name, "name", "", "Signal name.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "name")
	s.SingleWorkflowOrBatchOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalWorkflowStackCommand struct {
	Parent  *TemporalWorkflowCommand
	Command cobra.Command
	WorkflowReferenceOptions
	RejectCondition StringEnum
}

func NewTemporalWorkflowStackCommand(cctx *CommandContext, parent *TemporalWorkflowCommand) *TemporalWorkflowStackCommand {
	var s TemporalWorkflowStackCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "stack [flags]"
	s.Command.Short = "Trace a Workflow Execution"
	if hasHighlighting {
		s.Command.Long = "Perform a Query on a Workflow Execution using a \x1b[1m__stack_trace\x1b[0m-type\nQuery. Display a stack trace of the threads and routines currently in use\nby the Workflow for troubleshooting:\n\n\x1b[1mtemporal workflow stack \\\n    --workflow-id YourWorkflowId\x1b[0m"
	} else {
		s.Command.Long = "Perform a Query on a Workflow Execution using a `__stack_trace`-type\nQuery. Display a stack trace of the threads and routines currently in use\nby the Workflow for troubleshooting:\n\n```\ntemporal workflow stack \\\n    --workflow-id YourWorkflowId\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.WorkflowReferenceOptions.buildFlags(cctx, s.Command.Flags())
	s.RejectCondition = NewStringEnum([]string{"not_open", "not_completed_cleanly"}, "")
	s.Command.Flags().Var(&s.RejectCondition, "reject-condition", "Optional flag to reject Queries based on Workflow state. Accepted values: not_open, not_completed_cleanly.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type SharedWorkflowStartOptions struct {
	WorkflowId       string
	Type             string
	TaskQueue        string
	RunTimeout       time.Duration
	ExecutionTimeout time.Duration
	TaskTimeout      time.Duration
	SearchAttribute  []string
	Memo             []string
}

func (v *SharedWorkflowStartOptions) buildFlags(cctx *CommandContext, f *pflag.FlagSet) {
	f.StringVarP(&v.WorkflowId, "workflow-id", "w", "", "Workflow ID. If not supply, the Service generates a unique ID.")
	f.StringVar(&v.Type, "type", "", "Workflow Type name.")
	_ = cobra.MarkFlagRequired(f, "type")
	f.StringVarP(&v.TaskQueue, "task-queue", "t", "", "Workflow Task queue.")
	_ = cobra.MarkFlagRequired(f, "task-queue")
	f.DurationVar(&v.RunTimeout, "run-timeout", 0, "Fail a Workflow Run if it takes longer than `DURATION`.")
	f.DurationVar(&v.ExecutionTimeout, "execution-timeout", 0, "Fail a WorkflowExecution if it takes longer than `DURATION`. This time-out includes retries and ContinueAsNew tasks.")
	f.DurationVar(&v.TaskTimeout, "task-timeout", 10000*time.Millisecond, "Fail a Workflow Task if it takes longer than `DURATION`. This is the Start-to-close timeout for a Workflow Task.")
	f.StringArrayVar(&v.SearchAttribute, "search-attribute", nil, "Search Attribute in \"KEY='VALUE'\" format. KEY is a string, VALUE is JSON.")
	f.StringArrayVar(&v.Memo, "memo", nil, "Memo in 'KEY=\"VALUE\"' format. KEY is a string, VALUE is JSON.")
}

type WorkflowStartOptions struct {
	Cron          string
	FailExisting  bool
	StartDelay    time.Duration
	IdReusePolicy string
}

func (v *WorkflowStartOptions) buildFlags(cctx *CommandContext, f *pflag.FlagSet) {
	f.StringVar(&v.Cron, "cron", "", "Cron schedule for the Workflow. Deprecated. Use Schedules instead.")
	f.BoolVar(&v.FailExisting, "fail-existing", false, "Fail if the Workflow already exists.")
	f.DurationVar(&v.StartDelay, "start-delay", 0, "Delay before starting the Workflow Execution. Can't be used with cron schedules. If the Workflow receives a signal or update prior to this time, the Workflow Execution starts immediately.")
	f.StringVar(&v.IdReusePolicy, "id-reuse-policy", "", "Re-use policy for the Workflow ID in new Workflow Executions. Accepted values: AllowDuplicate, AllowDuplicateFailedOnly, RejectDuplicate, TerminateIfRunning.")
}

type PayloadInputOptions struct {
	Input       []string
	InputFile   []string
	InputMeta   []string
	InputBase64 bool
}

func (v *PayloadInputOptions) buildFlags(cctx *CommandContext, f *pflag.FlagSet) {
	f.StringArrayVarP(&v.Input, "input", "i", nil, "Input value. Use JSON content or set --input-meta to override. Can't be combined with --input-file. Can be passed multiple times.")
	f.StringArrayVar(&v.InputFile, "input-file", nil, "A path or paths for input file(s). Use JSON content or set --input-meta to override. Can't be combined with --input. Can be passed multiple times to concatenate the file contents.")
	f.StringArrayVar(&v.InputMeta, "input-meta", nil, "Input payload metadata as a \"KEY=VALUE\" pair. When the KEY is \"encoding\", this overrides the default (\"json/plain\"). Can be passed multiple times.")
	f.BoolVar(&v.InputBase64, "input-base64", false, "Assume inputs are base64-encoded and attempt to decode them.")
}

type TemporalWorkflowStartCommand struct {
	Parent  *TemporalWorkflowCommand
	Command cobra.Command
	SharedWorkflowStartOptions
	WorkflowStartOptions
	PayloadInputOptions
}

func NewTemporalWorkflowStartCommand(cctx *CommandContext, parent *TemporalWorkflowCommand) *TemporalWorkflowStartCommand {
	var s TemporalWorkflowStartCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "start [flags]"
	s.Command.Short = "Initiate a Workflow Execution"
	if hasHighlighting {
		s.Command.Long = "Start a new Workflow Execution. Returns the Workflow-\nand Run-IDs.\n\n\x1b[1mtemporal workflow start \\\n\t\t--workflow-id YourWorkflowId \\\n\t\t--type YourWorkflow \\\n\t\t--task-queue YourTaskQueue \\\n\t\t--input '{\"Input\": \"As-JSON\"}'\x1b[0m"
	} else {
		s.Command.Long = "Start a new Workflow Execution. Returns the Workflow-\nand Run-IDs.\n\n```\ntemporal workflow start \\\n\t\t--workflow-id YourWorkflowId \\\n\t\t--type YourWorkflow \\\n\t\t--task-queue YourTaskQueue \\\n\t\t--input '{\"Input\": \"As-JSON\"}'\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.SharedWorkflowStartOptions.buildFlags(cctx, s.Command.Flags())
	s.WorkflowStartOptions.buildFlags(cctx, s.Command.Flags())
	s.PayloadInputOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalWorkflowTerminateCommand struct {
	Parent     *TemporalWorkflowCommand
	Command    cobra.Command
	WorkflowId string
	Query      string
	RunId      string
	Reason     string
	Yes        bool
}

func NewTemporalWorkflowTerminateCommand(cctx *CommandContext, parent *TemporalWorkflowCommand) *TemporalWorkflowTerminateCommand {
	var s TemporalWorkflowTerminateCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "terminate [flags]"
	s.Command.Short = "Forcefully end a Workflow Execution"
	if hasHighlighting {
		s.Command.Long = "Terminate a Workflow Execution:\n\n\x1b[1mtemporal workflow terminate \\\n    --reason YourReasonForTermination \\\n    --workflow-id YourWorkflowId\x1b[0m\n\nThe reason is optional and defaults to the current user's name. The reason\nis stored in the Event History as part of the \x1b[1mWorkflowExecutionTerminated\x1b[0m\nevent. This becomes the closing Event in the Workflow Execution's history.\n\nExecutions may be terminated in bulk via a visibility query list filter:\n\n\x1b[1mtemporal workflow terminate \\\n    --query YourQuery \\\n    --reason YourReasonForTermination\x1b[0m\n\nWorkflow code cannot see or respond to terminations. To perform clean-up work\nin your Workflow code, use \x1b[1mtemporal workflow cancel\x1b[0m instead."
	} else {
		s.Command.Long = "Terminate a Workflow Execution:\n\n```\ntemporal workflow terminate \\\n    --reason YourReasonForTermination \\\n    --workflow-id YourWorkflowId\n```\n\nThe reason is optional and defaults to the current user's name. The reason\nis stored in the Event History as part of the `WorkflowExecutionTerminated`\nevent. This becomes the closing Event in the Workflow Execution's history.\n\nExecutions may be terminated in bulk via a visibility query list filter:\n\n```\ntemporal workflow terminate \\\n    --query YourQuery \\\n    --reason YourReasonForTermination\n```\n\nWorkflow code cannot see or respond to terminations. To perform clean-up work\nin your Workflow code, use `temporal workflow cancel` instead."
	}
	s.Command.Args = cobra.NoArgs
	s.Command.Flags().StringVarP(&s.WorkflowId, "workflow-id", "w", "", "Workflow ID. You must set either --workflow-id or --query.")
	s.Command.Flags().StringVarP(&s.Query, "query", "q", "", "Content for an SQL-like `QUERY` List Filter. You must set either --workflow-id or --query.")
	s.Command.Flags().StringVarP(&s.RunId, "run-id", "r", "", "Run ID. Can only be set with --workflow-id. Do not use with --query.")
	s.Command.Flags().StringVar(&s.Reason, "reason", "", "Reason for termination. Defaults to message with the current user's name.")
	s.Command.Flags().BoolVarP(&s.Yes, "yes", "y", false, "Don't prompt to confirm termination. Can only be used with --query.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalWorkflowTraceCommand struct {
	Parent  *TemporalWorkflowCommand
	Command cobra.Command
	WorkflowReferenceOptions
	Fold        []string
	NoFold      bool
	Depth       int
	Concurrency int
}

func NewTemporalWorkflowTraceCommand(cctx *CommandContext, parent *TemporalWorkflowCommand) *TemporalWorkflowTraceCommand {
	var s TemporalWorkflowTraceCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "trace [flags]"
	s.Command.Short = "Workflow Execution live progress"
	if hasHighlighting {
		s.Command.Long = "Display the progress of a Workflow Execution and its child\nworkflows with a real-time trace. This view helps you understand how Workflows\nare proceeding:\n\n\x1b[1mtemporal workflow trace \\\n    --workflow-id YourWorkflowId\x1b[0m"
	} else {
		s.Command.Long = "Display the progress of a Workflow Execution and its child\nworkflows with a real-time trace. This view helps you understand how Workflows\nare proceeding:\n\n```\ntemporal workflow trace \\\n    --workflow-id YourWorkflowId\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.WorkflowReferenceOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Flags().StringArrayVar(&s.Fold, "fold", nil, "Status for folding away Child Workflows. Case-insensitive. Ignored if --no-fold supplied. Available values: running, completed, failed, canceled, terminated, timedout, continueasnew. Can be passed multiple times. Each fold reduces the amount of information fetched and displayed.")
	s.Command.Flags().BoolVar(&s.NoFold, "no-fold", false, "Disable folding. Fetch and display Child Workflows within set depth.")
	s.Command.Flags().IntVar(&s.Depth, "depth", -1, "Set depth for Child Workflow fetches. Pass -1 to fetch child workflows at any depth.")
	s.Command.Flags().IntVar(&s.Concurrency, "concurrency", 10, "Number of Workflow Histories to fetch at a time.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}

type TemporalWorkflowUpdateCommand struct {
	Parent  *TemporalWorkflowCommand
	Command cobra.Command
	PayloadInputOptions
	Name                string
	WorkflowId          string
	UpdateId            string
	RunId               string
	FirstExecutionRunId string
}

func NewTemporalWorkflowUpdateCommand(cctx *CommandContext, parent *TemporalWorkflowCommand) *TemporalWorkflowUpdateCommand {
	var s TemporalWorkflowUpdateCommand
	s.Parent = parent
	s.Command.DisableFlagsInUseLine = true
	s.Command.Use = "update [flags]"
	s.Command.Short = "Synchronously run a Workflow update handler"
	if hasHighlighting {
		s.Command.Long = "Send a message to a Workflow Execution to invoke an\nupdate handler. An update can change the\nstate of a Workflow Execution and return a response:\n\n\x1b[1mtemporal workflow update \\\n\t\t--workflow-id YourWorkflowId \\\n\t\t--name YourUpdate \\\n\t\t--input '{\"Input\": \"As-JSON\"}'\x1b[0m"
	} else {
		s.Command.Long = "Send a message to a Workflow Execution to invoke an\nupdate handler. An update can change the\nstate of a Workflow Execution and return a response:\n\n```\ntemporal workflow update \\\n\t\t--workflow-id YourWorkflowId \\\n\t\t--name YourUpdate \\\n\t\t--input '{\"Input\": \"As-JSON\"}'\n```"
	}
	s.Command.Args = cobra.NoArgs
	s.PayloadInputOptions.buildFlags(cctx, s.Command.Flags())
	s.Command.Flags().StringVar(&s.Name, "name", "", "Handler method name.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "name")
	s.Command.Flags().StringVarP(&s.WorkflowId, "workflow-id", "w", "", "Workflow ID.")
	_ = cobra.MarkFlagRequired(s.Command.Flags(), "workflow-id")
	s.Command.Flags().StringVar(&s.UpdateId, "update-id", "", "Update ID. If unset, defaults to a UUID. Must be unique per Workflow Execution.")
	s.Command.Flags().StringVarP(&s.RunId, "run-id", "r", "", "Run ID. If unset, updates the currently-running Workflow Execution.")
	s.Command.Flags().StringVar(&s.FirstExecutionRunId, "first-execution-run-id", "", "Parent Run ID. The update is sent to the last Workflow Execution in the chain started with this Run ID.")
	s.Command.Run = func(c *cobra.Command, args []string) {
		if err := s.run(cctx, args); err != nil {
			cctx.Options.Fail(err)
		}
	}
	return &s
}
