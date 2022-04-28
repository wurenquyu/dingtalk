/*
 * Copyright 2020 zhaoyunxing.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dingtalk

//
//import (
//	"github.com/wurenquyu/dingtalk/v2/constant"
//	"github.com/wurenquyu/dingtalk/v2/domain"
//	"net/http"
//)
//
////发送工作通知返回结果
//type WorkNotifyRep struct {
//	domain.Response
//	RequestId string `json:"request_id"` //请求的id
//	TaskId    int    `json:"task_id"`    //创建的异步发送任务id
//}
//
////获取工作通知消息的发送进度返回
//type WorkNotifyProgressRsp struct {
//	domain.Response
//	workNotifyProgress `json:"progress"`
//}
//
//type WorkNotifyResultRsp struct {
//	domain.Response
//	sendResult `json:"send_result"` //返回结果。
//}
//
////工作通知进度
//type workNotifyProgress struct {
//	Percent int `json:"progress_in_percent"` //取值0~100，表示处理的百分比
//	Status  int `json:"status"`              //任务执行状态，0：未开始1：处理中 2：处理完毕
//}
//
////发送返回s
//type sendResult struct {
//	InvalidUser   []string        `json:"invalid_user_id_list"`   //无效的用户id
//	InvalidDept   []int           `json:"invalid_dept_id_list"`   //无效部门id
//	ForbiddenUser []string        `json:"forbidden_user_id_list"` //因发送消息过于频繁或超量而被流控过滤后实际未发送的UserID。
//	FailedUser    []string        `json:"failed_user_id_list"`    //发送失败的userid。
//	ReadUser      []string        `json:"read_user_id_list"`      //已读消息的userid。
//	UnReadUser    []string        `json:"unread_user_id_list"`    //未读消息的userid。
//	ForbiddenList []forbiddenList `json:"forbidden_list"`         //推送被禁止的具体原因。
//}
//
////推送被禁止的具体原因。
//// code:143105表示企业自建应用每日推送给用户的消息超过上限,
////      143106表示企业自建应用推送给用户的消息重复
//type forbiddenList struct {
//	Code   string `json:"code"`   //流控code
//	Count  int    `json:"count"`  //流控阀值。
//	UserId string `json:"userid"` //被流控员工的userid。
//}
//
////发送普通文本工作通知
////部门id或用户id重复会剔除
////发送工作通知
//func (ding *DingTalk) SendWorkNotify(res domain.WorkNotifyRes) (resp WorkNotifyRep, err error) {
//	//组装部门、用户
//	res.AssembleDept()
//	res.AssembleUser()
//	res.SetAgentId(ding.id)
//
//	err = ding.Request(http.MethodPost, constant.SendCorpConversationKey, nil, res, &resp)
//
//	return resp, err
//}
//
////获取工作通知消息的发送进度，仅支持查询24小时内的任务。
////taskId:工作通知id
//func (ding *DingTalk) GetWorkNotifyProgress(taskId int) (rsp WorkNotifyProgressRsp, err error) {
//	form := map[string]interface{}{
//		"agent_id": ding.id,
//		"task_id":  taskId,
//	}
//	err = ding.Request(http.MethodPost, constant.GetSendProgressKey, nil, form, &rsp)
//
//	return rsp, err
//}
//
////获取工作通知消息的发送结果
////taskId:工作通知id
//func (ding *DingTalk) GetWorkNotifySendResult(taskId int) (rsp WorkNotifyResultRsp, err error) {
//	form := map[string]interface{}{
//		"agent_id": ding.id,
//		"task_id":  taskId,
//	}
//	err = ding.Request(http.MethodPost, constant.GetSendResultKey, nil, form, &rsp)
//
//	return rsp, err
//}
//
////撤回工作通知消息
////taskId:工作通知id
//func (ding *DingTalk) RecallWorkNotifySendResult(taskId int) (rsp domain.Response, err error) {
//	form := map[string]interface{}{
//		"agent_id":    ding.id,
//		"msg_task_id": taskId,
//	}
//	err = ding.Request(http.MethodPost, constant.RecallCorpConversationKey, nil, form, &rsp)
//
//	return rsp, err
//}
