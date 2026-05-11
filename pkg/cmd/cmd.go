package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	"k8s.io/cli-runtime/pkg/resource"
)

type PatchAllOptions struct {
	configFlags *genericclioptions.ConfigFlags
	selector    string
	dryRun      bool
	args        []string
	genericiooptions.IOStreams
}

func NewPatchAllOptions(streams genericiooptions.IOStreams) *PatchAllOptions {
	return &PatchAllOptions{
		configFlags: genericclioptions.NewConfigFlags(true),
		IOStreams:   streams,
	}
}

func NewPatchAllCmd(streams genericiooptions.IOStreams) *cobra.Command {
	o := NewPatchAllOptions(streams)
	cmd := &cobra.Command{
		Use:   "kubectl patch-all <resource> [flags]",
		Short: "Batch-patch Kubernetes resources by applying a single patch across many objects",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.Complete(cmd, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			return o.Run()
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	o.configFlags.AddFlags(cmd.Flags())
	cmd.Flags().StringVarP(&o.selector, "selector", "l", "",
		"Selector (label query) to filter on, supports '=', '==', and '!='.")
	cmd.Flags().BoolVarP(&o.dryRun, "dry-run", "d", false,
		"Don't patch — just print the resources that would be patched.")

	return cmd
}

func (o *PatchAllOptions) Complete(cmd *cobra.Command, args []string) error {
	o.args = args
	return nil
}

func (o *PatchAllOptions) Validate() error {
	if len(o.args) == 0 {
		return fmt.Errorf("at least one resource type is required (e.g. pods, deploy)")
	}
	return nil
}

func (o *PatchAllOptions) Run() error {
	if o.dryRun {
		return o.runDryRun()
	}
	return fmt.Errorf("patch operation not yet implemented; rerun with --dry-run to preview matching resources")
}

func (o *PatchAllOptions) runDryRun() error {
	ns, _, err := o.configFlags.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return err
	}

	r := resource.NewBuilder(o.configFlags).
		Unstructured().
		NamespaceParam(ns).DefaultNamespace().
		ResourceTypeOrNameArgs(true, o.args...).
		LabelSelectorParam(o.selector).
		ContinueOnError().
		Latest().
		Flatten().
		Do()

	if err := r.Err(); err != nil {
		return err
	}

	return r.Visit(func(info *resource.Info, err error) error {
		if err != nil {
			return err
		}
		fmt.Fprintf(o.Out, "%s\t%s\t%s\n",
			info.Mapping.Resource.Resource, info.Namespace, info.Name)
		return nil
	})
}

func Run(ctx context.Context, streams genericiooptions.IOStreams) int {
	cmd := NewPatchAllCmd(streams)
	if err := cmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return 1
	}
	return 0
}
