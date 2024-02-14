package goutbra

import (
	"bytes"
	"testing"

	"github.com/drewstinnett/gout/v2"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestDefaultBind(t *testing.T) {
	require.NoError(t, Bind(&cobra.Command{}))
}

func TestCmd(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			Cmd(cmd)
			gout.SetWriter(b)
			return gout.Print([]string{"foo", "bar"})
		},
	}
	require.NoError(t, Bind(cmd))

	// Default yaml exporter
	require.NoError(t, cmd.Execute())
	require.Equal(t, "- foo\n- bar\n", b.String())

	// Try it with json
	b = bytes.NewBufferString("")
	cmd.SetArgs([]string{"--format", "json"})
	require.NoError(t, cmd.Execute())
	require.Equal(t, "[\"foo\",\"bar\"]\n", b.String())
}

func TestBindOptions(t *testing.T) {
	cmd := &cobra.Command{}
	require.NoError(t, Bind(cmd,
		WithField("output"),
		WithDefault("json"),
		WithDefaultTemplate("some-template"),
	))
	u := cmd.UsageString()
	require.Contains(t, u, " --output ")
	require.Contains(t, u, " --output-template ")
	require.Contains(t, u, `default "json"`)
	require.Contains(t, u, `default "some-template"`)
}

func TestCustomHelp(t *testing.T) {
	cmd := &cobra.Command{}
	require.NoError(t, Bind(cmd,
		WithHelp("help for the format field"),
		WithHelpTemplate("help for the template field"),
	))
	u := cmd.UsageString()
	require.Contains(t, u, "help for the format field")
	require.Contains(t, u, "help for the template field")
}
