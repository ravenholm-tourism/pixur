//go:generate protoc config.proto --go_out=.
package config

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
)

var (
	defaultValues = &Config{
		DbName:                "sqlite3",
		DbConfig:              ":memory:",
		ListenAddress:         ":8889",
		ListenNetwork:         "tcp",
		PixPath:               "pix",
		SessionPrivateKeyPath: "",
		SessionPublicKeyPath:  "",
		TokenSecret:           "",
	}
	Conf = mergeParseConfigFlag(defaultValues)
)

func init() {
	_ = flag.String("configbe", ".configbe.pb.txt", "The default configuration file")
	flag.StringVar(&Conf.ListenAddress, "listen_address", Conf.ListenAddress, "Listening Address")
	flag.StringVar(&Conf.ListenNetwork, "listen_network", Conf.ListenNetwork, "Listening Network")
	flag.StringVar(&Conf.PixPath, "pix_path", Conf.PixPath, "Default picture storage directory")
	flag.StringVar(&Conf.DbName, "db_name", Conf.DbName, "Database Name")
	flag.StringVar(&Conf.DbConfig, "db_config", Conf.DbConfig, "Database Configuration")
}

func mergeParseConfigFlag(defaults *Config) *Config {
	conf, err := parseConfigFlag()
	if err != nil {
		glog.Fatal(err)
	}
	merged := *defaults
	proto.Merge(&merged, conf)
	return &merged
}

func parseConfigFlag() (*Config, error) {
	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	configPath := fs.String("configbe", envOrDefault("PIXUR_BE_CONFIG", ".configbe.pb.txt"), "")
	if err := fs.Parse(os.Args[1:]); err != nil && err != flag.ErrHelp {
		// ignore, the next parse call will find it.
		glog.V(2).Info(err)
	}
	var config = new(Config)
	f, err := os.Open(*configPath)
	if os.IsNotExist(err) {
		glog.Warning("Unable to open config file, using defaults", err)
		return config, nil
	} else if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	if err := proto.UnmarshalText(string(data), config); err != nil {
		return nil, err
	}

	if !filepath.IsAbs(config.PixPath) {
		dir := filepath.Dir(f.Name())
		config.PixPath = filepath.Join(dir, config.PixPath)
	}

	return config, nil
}

func envOrDefault(name, defaultVal string) string {
	if val, ok := os.LookupEnv(name); ok {
		return val
	}
	return defaultVal
}
