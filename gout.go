/*
Package goutbra integrate gout with cobra commands
*/
package goutbra

import (
	"fmt"
	"strings"

	"github.com/drewstinnett/gout/v2"
	"github.com/drewstinnett/gout/v2/config"
	"github.com/drewstinnett/gout/v2/formats/gotemplate"
	"github.com/spf13/cobra"
)

// Config defines what fields the formatting values are stored in
type Config struct {
	FormatField     string
	FormatDefault   string
	FormatHelp      string
	TemplateDefault string
	TemplateHelp    string
}

// Option is a functional option to pass in to a new Config
type Option func(*Config)

// WithField sets the name of the field for the 'format' in a cobra.Command
//
// Example: WithField("output") would produce the "--output" and "--output-template" as the fields instead of "--format" and "--format-template"
func WithField(s string) Option {
	return func(c *Config) {
		c.FormatField = s
	}
}

// WithDefault sets the default format to produce.
func WithDefault(s string) Option {
	return func(c *Config) {
		c.FormatDefault = s
	}
}

// WithDefaultTemplate sets the default template string
func WithDefaultTemplate(s string) Option {
	return func(c *Config) {
		c.TemplateDefault = s
	}
}

// WithHelp defines wht the help for the format itself should be
func WithHelp(s string) Option {
	return func(c *Config) {
		c.FormatHelp = s
	}
}

// WithHelpTemplate defines what the help for the format template should be
func WithHelpTemplate(s string) Option {
	return func(c *Config) {
		c.TemplateHelp = s
	}
}

// newConfig generates a new CobraCmdConfig with some sane defaults
func newConfig(opts ...Option) *Config {
	// Set up a default config
	c := &Config{
		FormatField:     "format",
		FormatDefault:   "yaml",
		FormatHelp:      "Format to use for output",
		TemplateDefault: "{{ . }}",
		TemplateHelp:    "Template to use when using the gotemplate format",
	}

	// Override that stuff
	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Bind creates a new set of flags for Cobra that can be used to
// configure Gout
// func Bind(cmd *cobra.Command, config *Config) error {
func Bind(cmd *cobra.Command, opts ...Option) error {
	config := newConfig(opts...)
	keys := []string{}
	for k := range gout.BuiltInFormatters {
		keys = append(keys, k)
	}
	cmd.PersistentFlags().String(
		config.FormatField,
		config.FormatDefault,
		config.FormatHelp+" ("+strings.Join(keys, "|")+")",
	)
	cmd.PersistentFlags().String(
		config.FormatField+"-template",
		config.TemplateDefault,
		config.TemplateHelp,
	)
	return nil
}

// Cmd sets up the the built-in Gout client using options from the cobra.Command
func Cmd(cmd *cobra.Command, opts ...Option) error {
	return CmdGout(cmd, gout.GetGout(), opts...)
}

// CmdGout sets up a given gout using options from the cobra.Command
func CmdGout(cmd *cobra.Command, g *gout.Gout, opts ...Option) error {
	return apply(g, cmd, opts...)
}

// apply uses settings from cobra.Command to set up a given Gout instance
func apply(g *gout.Gout, cmd *cobra.Command, opts ...Option) error {
	conf := newConfig(opts...)

	format, err := cmd.Flags().GetString(conf.FormatField)
	if err != nil {
		return err
	}
	if format == "gotemplate" {
		t, err := cmd.PersistentFlags().GetString(conf.FormatField + "-template")
		if err != nil {
			return err
		}
		g.SetFormatter(gotemplate.Formatter{
			Opts: config.FormatterOpts{
				"template": t,
			},
		})
	} else {
		if fr, ok := gout.BuiltInFormatters[format]; ok {
			g.SetFormatter(fr)
		} else {
			return fmt.Errorf("could not find the format %v", format)
		}
	}
	return nil
}

/*
// NewWithCobraCmd creates a pointer to a new writer with options from a cobra.Command
func NewWithCobraCmd(cmd *cobra.Command, opts ...Option) (*gout.Gout, error) {
	// Default this writer to stdout
	c, err := gout.New()
	if err != nil {
		return nil, err
	}

	if err := apply(c, cmd, opts...); err != nil {
		return nil, err
	}

	return c, nil
}
*/
