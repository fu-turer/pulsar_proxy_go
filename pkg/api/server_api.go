// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package api

import (
	pb "github.com/paashzj/pulsar_proto_go"
	"github.com/panjf2000/gnet"
	"net"
)

type PulsarServer interface {
	Connect(addr net.Addr, connect *pb.CommandConnect) (*pb.CommandConnected, error)

	Lookup(addr net.Addr, lookup *pb.CommandLookupTopic) (*pb.CommandLookupTopicResponse, error)

	DirectHandler(addr net.Addr, command *pb.BaseCommand) ([]byte, gnet.Action)
}
