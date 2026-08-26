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
	"time"

	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
)

var _ plugin.MigrationScript = (*addAiCommitsDomain)(nil)

type addAiCommitsDomain struct{}

type archivedAiCommit20260826 struct {
	Id           string `gorm:"primaryKey;type:varchar(255)"`
	ProjectName  string `gorm:"index;type:varchar(255)"`
	CommitSha    string `gorm:"index;type:varchar(40)"`
	RepoId       string `gorm:"index;type:varchar(255)"`
	AiTool       string `gorm:"type:varchar(100)"`
	AuthorName   string `gorm:"type:varchar(255)"`
	AuthoredDate time.Time
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}

func (archivedAiCommit20260826) TableName() string { return "ai_commits" }

func (*addAiCommitsDomain) Up(basicRes context.BasicRes) errors.Error {
	db := basicRes.GetDal()
	if err := db.AutoMigrate(&archivedAiCommit20260826{}); err != nil {
		return err
	}
	return nil
}

func (*addAiCommitsDomain) Version() uint64 {
	return 20260826000001
}

func (*addAiCommitsDomain) Name() string {
	return "add ai_commits domain table"
}
