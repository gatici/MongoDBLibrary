// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package main

import (
	"github.com/omec-project/MongoDBLibrary"
	"log"
)

// TODO : take DB name from helm chart
// TODO : inbuild shell commands to

func main() {
	log.Println("dbtestapp started")

	// connect to mongoDB
	MongoDBLibrary.SetMongoDB("sdcore", "mongodb://mongodb-arbiter-headless")

	initDrsm("ngapid")

	//blocking
	http_server()
}
