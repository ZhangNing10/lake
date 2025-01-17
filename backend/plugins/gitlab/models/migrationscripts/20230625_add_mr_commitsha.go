/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package migrationscripts

import (
	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/helpers/migrationhelper"
)

type addMrCommitSha struct{}

type GitlabMergeRequest20230625 struct {
	MergeCommitSha  string `gorm:"type:varchar(255)"`
	SquashCommitSha string `gorm:"type:varchar(255)"`
}

func (GitlabMergeRequest20230625) TableName() string {
	return "_tool_gitlab_merge_requests"
}

func (*addMrCommitSha) Up(baseRes context.BasicRes) errors.Error {
	err := migrationhelper.AutoMigrateTables(
		baseRes,
		&GitlabMergeRequest20230625{},
	)
	if err != nil {
		return err
	}

	return nil
}

func (*addMrCommitSha) Version() uint64 {
	return 20230625110339
}

func (*addMrCommitSha) Name() string {
	return "add _tool_gitlab_issue_assignees table"
}
