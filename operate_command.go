// Copyright 2013-2014 Aerospike, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aerospike

type OperateCommand struct {
	ReadCommand

	policy     *WritePolicy
	operations []*Operation
}

func NewOperateCommand(cluster *Cluster, policy *WritePolicy, key *Key, operations []*Operation) *OperateCommand {
	if policy == nil {
		policy = NewWritePolicy(0, 0)
	}
	return &OperateCommand{
		ReadCommand: *NewReadCommand(cluster, policy, key, nil),
		policy:      policy,
		operations:  operations,
	}
}

func (this *OperateCommand) writeBuffer(ifc Command) error {
	return this.SetOperate(this.policy, this.key, this.operations)
}

func (this *OperateCommand) Execute() error {
	return this.execute(this)
}