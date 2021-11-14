/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dingtalk

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"testing"
)

func TestDingTalk_CreateUser(t *testing.T) {

	user, err := client.CreateUser(request.NewCreateUser("张三", "15669019211", 554656655).Build())
	assert.Nil(t, err)
	assert.NotNil(t, user.UserId)
	t.Log(user.UserId)
}

func TestDingTalk_UpdateUser(t *testing.T) {

	res, err := client.UpdateUser(
		request.NewUpdateUser("1948546245774889").
			SetName("李四").
			//SetDept(1, 560935057).
			//SetTelephone("").
			//SetForceUpdateFields("title,userid,title,userid").
			//SetDeptTitle(1, "普通员工").
			//SetDeptTitle(554656655, "设计人员").
			//SetJobNumber("").
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res.Code)
}

func TestDingTalk_DeleteUser(t *testing.T) {

	res, err := client.DeleteUser("1948546245774889")

	assert.Nil(t, err)
	assert.NotNil(t, res.Code)
	assert.NotNil(t, res.RequestId)
}

func TestDingTalk_GetUserDetail(t *testing.T) {

	detail, err := client.GetUserDetail(request.NewUserDetail("1948546245774889").Build())

	assert.Nil(t, err)
	assert.NotNil(t, detail.Code)
	assert.NotNil(t, detail.RequestId)
	assert.NotNil(t, detail.Name)
	assert.NotNil(t, detail.UnionId)
	assert.NotNil(t, detail.UserId)
}

func TestDingTalk_GetUserCount(t *testing.T) {

	req, err := client.GetUserCount(request.NewUserCount(true))

	assert.Nil(t, err)
	assert.NotNil(t, req.Count)
}

func TestDingTalk_GetInactiveUser(t *testing.T) {

	req, err := client.GetInactiveUser(
		request.NewInactiveUser(true, 0, 10, "20211103").
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, req)
}

func TestDingTalk_GetUserIdByMobile(t *testing.T) {

	req, err := client.GetUserIdByMobile(request.NewMobileGetUserId("18357154439"))

	assert.Nil(t, err)
	assert.NotNil(t, req.UserId)
}

func TestDingTalk_GetUserIdByUnionId(t *testing.T) {

	req, err := client.GetUserIdByUnionId(request.NewUnionIdGetUserId("DmJWoRm5zkjT9kWWgVpm0giEiE"))

	assert.Nil(t, err)
	assert.NotNil(t, req.UserId)
}

func TestDingTalk_GetOrgAdminUser(t *testing.T) {

	req, err := client.GetOrgAdminUser()

	assert.Nil(t, err)
	assert.NotNil(t, req.AdminUser)
}

func TestDingTalk_GetOrgAdminScope(t *testing.T) {

	req, err := client.GetOrgAdminScope(request.NewAdminUserScope("manager164"))

	assert.Nil(t, err)
	assert.NotNil(t, req.DeptIds)
}

func TestDingTalk_GetUserCanAccessApplet(t *testing.T) {

	req, err := client.GetUserCanAccessApplet(1244553273, "manager164")

	assert.Nil(t, err)
	assert.NotNil(t, req.Access)
}

func TestDingTalk_GetUserInfoByCode(t *testing.T) {

	req, err := isv.GetUserInfoByCode("ed61dbdbd5733577b7fe9e7d993b0a00")

	assert.Nil(t, err)
	assert.NotNil(t, req)
}

func TestDingTalk_GetSSOUserInfo(t *testing.T) {

	req, err := client.GetSSOUserInfo("689635c61c2430b187410fc9ccfda645")

	assert.Nil(t, err)
	assert.NotNil(t, req)
}
