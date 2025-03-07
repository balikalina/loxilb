/*
 * Copyright (c) 2022 NetLOX Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"github.com/loxilb-io/loxilb/api/restapi"
	"github.com/loxilb-io/loxilb/api/restapi/handler"
	"github.com/loxilb-io/loxilb/api/restapi/operations"
	cmn "github.com/loxilb-io/loxilb/common"
	"github.com/loxilb-io/loxilb/options"
	tk "github.com/loxilb-io/loxilib"
	"log"
	"os"
	"runtime/debug"
	"time"
)

var (
	ApiReady bool
)

// RegisterAPIHooks - routine to register interface for api
func RegisterAPIHooks(hooks cmn.NetHookInterface) {
	handler.ApiHooks = hooks
}

// WaitAPIServerReady - routine to wait till api server is up
func WaitAPIServerReady() {
	for {
		if ApiReady {
			time.Sleep(2 * time.Second)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

// RunAPIServer - routine to start API server
func RunAPIServer() {

	// Stack trace logger
	defer func() {
		if e := recover(); e != nil {
			tk.LogIt(tk.LogCritical, "%s: %s", e, debug.Stack())
		}
	}()

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewLoxilbRestAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Loxilb Rest API"
	parser.LongDescription = "Loxilb REST API for Baremetal Scenarios"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()
	// API server host list
	server.Host = options.Opts.Host
	server.TLSHost = options.Opts.TLSHost
	server.TLSCertificateKey = options.Opts.TLSCertificateKey
	server.TLSCertificate = options.Opts.TLSCertificate
	server.Port = options.Opts.Port
	server.TLSPort = options.Opts.TLSPort
	ApiReady = true

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
