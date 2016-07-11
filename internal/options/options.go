package options

import (
	"github.com/jessevdk/go-flags"

	"github.com/hirakiuc/ec2s/internal/common"
	"github.com/hirakiuc/ec2s/internal/config"
)

var logger *common.Logger

func init() {
	logger = common.GetLogger()
}

// Options describe global options of ec2s command.
type Options struct {
	ConfigPath string `short:"c" long:"config-path" description:"<string> config file path" default:"~/.ec2s.toml"`
	Verbose    bool   `short:"v" long:"verbose" description:"show verbose messages"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

func init() {
	options = Options{}
}

// ParseOptions parse global options.
func ParseOptions() ([]string, error) {
	return parser.Parse()
}

// GetOptions return pointer to options object.
func GetOptions() *Options {
	return &options
}

// AddCommand invoke AddCommand function of flags.Parser.
func AddCommand(command string, shortDescription string, longDescription string, data interface{}) (*flags.Command, error) {
	return parser.AddCommand(command, shortDescription, longDescription, data)
}

// Validate validate global options.
func (opts Options) Validate() error {
	if _, err := config.LoadConfig(options.ConfigPath); err != nil {
		logger.Error("Can't load config file.\n")
		return err
	}

	return nil
}
