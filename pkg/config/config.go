// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package config

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/koding/multiconfig"

	"openpitrix.io/openpitrix/pkg/logger"
)

type Config struct {
	Log       LogConfig
	Grpc      GrpcConfig
	Mysql     MysqlConfig
	Etcd      EtcdConfig
	Profiling ProfilingConfig
}

type LogConfig struct {
	Level string `default:"info"` // debug, info, warn, error, fatal
}

type GrpcConfig struct {
	ShowErrorCause bool `default:"false"` // show grpc error cause to frontend
}

type EtcdConfig struct {
	Endpoints string `default:"openpitrix-etcd:2379"` // Example: "localhost:2379,localhost:22379,localhost:32379"
}

type MysqlConfig struct {
	Host     string `default:"openpitrix-db"`
	Port     string `default:"3306"`
	User     string `default:"root"`
	Password string `default:"password"`
	Database string `default:"openpitrix"`
	Disable  bool   `default:"false"`
}

type ProfilingConfig struct {
	Enable bool `default:"false"`
}

func (m *MysqlConfig) GetUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", m.User, m.Password, m.Host, m.Port, m.Database)
}

func PrintUsage() {
	fmt.Fprintf(os.Stdout, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprint(os.Stdout, "\nSupported environment variables:\n")
	e := newLoader("openpitrix")
	e.PrintEnvs(new(Config))
	fmt.Println("")
}

func GetFlagSet() *flag.FlagSet {
	flag.CommandLine.Usage = PrintUsage
	return flag.CommandLine
}

func ParseFlag() {
	GetFlagSet().Parse(os.Args[1:])
}

var profilingServerStarted = false

func LoadConf() *Config {
	ParseFlag()

	config := new(Config)
	m := &multiconfig.DefaultLoader{}
	m.Loader = multiconfig.MultiLoader(newLoader("openpitrix"))
	m.Validator = multiconfig.MultiValidator(
		&multiconfig.RequiredValidator{},
	)
	err := m.Load(config)
	if err != nil {
		logger.Critical(nil, "Failed to load config: %+v", err)
		panic(err)
	}
	logger.SetLevelByString(config.Log.Level)
	logger.Debug(nil, "LoadConf: %+v", config)

	if config.Profiling.Enable && !profilingServerStarted {
		profilingServerStarted = true
		logger.Info(nil, "Profiling start...")
		go func() {
			mux := http.NewServeMux()
			mux.HandleFunc("/debug/pprof/", pprof.Index)
			mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
			mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
			mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
			mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
			logger.Info(nil, "Profiling log: %s", http.ListenAndServe("localhost:9300", mux))
		}()
	}

	return config
}
