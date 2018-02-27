// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package app

import (
	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/db"
	"openpitrix.io/openpitrix/pkg/manager"

	"google.golang.org/grpc"

	"openpitrix.io/openpitrix/pkg/config"
	"openpitrix.io/openpitrix/pkg/pb"
)

type Server struct {
	db *db.Database
}

func Serve(cfg *config.Config) {
	m := manager.GrpcServer{
		ServiceName: "app-manager",
		Port:        constants.AppManagerPort,
		MysqlConfig: cfg.Mysql,
	}
	m.Serve(func(server *grpc.Server, db *db.Database) {
		pb.RegisterAppManagerServer(server, &Server{db})
	})
}