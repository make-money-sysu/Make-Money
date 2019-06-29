状态码

| code |              |      |
| ---- | ------------ | ---- |
| 400  | bad request  |      |
| 401  | Unauthorized |      |
| 403  | Forbidden    |      |
|      |              |      |




## 1./
本地址用来测试服务器是否在正常运行，显示beego默认界面
## 2./login
**目前注册时没给session的，请登录获取**

| 方法 | 参数     | 返回值               |
| ---- | -------- | -------------------- |
| post | id       | status  string 状态  |
|      | password | msg  string 额外信息 |



## 3./user

返回的用户json格式：

| 字段         | 类型   |      |
| ------------ | ------ | ---- |
| id           | int    |      |
| real_name    | string |      |
| nick_name    | string |      |
| age          | uint   |      |
| gender       | string |      |
| head_picture | string |      |
| balance      | float  |      |
| profession   | string |      |
| grade        | string |      |
| phone        | string |      |
| email        | string |      |



| 方法   | 作用     | 参数 | 返回值                       |
| ------ | -------- | ---- | ---------------------------- |
| post   | 注册     |      |                              |
| delete | 注销账号 |      | status，msg                  |
| get    |          |      | status，msg, data (用户数组) |
| put    |          |      |                              |



<!-- 注： /?:id即为id不一定存在，并且id为变量，如上述地址同时匹配/user和/user/123 -->
此地址用来对用户表进行修改
符合RESTful设计支持GET,POST,PUT和DELETE
<!-- GET为获取用户信息，必须提供id(方式如上)： -->
get需要登录，不再需要id

输入body：raw
```
{
	"id" : 16340132,
	"password" : "dick20",
	"real_name" : "dick",
	"nick_name" : "dick",
	"age" : 18,
	"gender" : "1",
	"head_picture" : " blob not null1",
	"balance" : 100000,
	"profession" : "1",
	"grade" : "3",
	"phone" : "13680473185",
	"email" : "dic0k@qq.com"
}
```

返回为json格式，除了密码外此用户的所有信息都会返回，字段名同数据库中列名(见init.sql)
POST即为注册功能，http包body体中提供json，字段名同数据库即可，如果注册成功，则返回json:

```
{
	"status" : "success"
}
```
失败则返回json：
```
{
	"status" : "failed",
	"msg" : "报错信息"
}
```
PUT为修改用户信息功能，使用方法和POST类似，返回同上

## 4./survey/?:id

输入字段格式：

| 字段    | 类型   | 作用               |
| ------- | ------ | ------------------ |
| title   | string |                    |
| state   | int    | 0为未发布，1为发布 |
| checked | int    |                    |
| content | string |                    |



问卷格式：

| 字段         | 类型   | 作用               |
| ------------ | ------ | ------------------ |
| id           | int    |                    |
| publisher_id | int    |                    |
| title        | string |                    |
| state        | int    | 0为未发布，1为发布 |
| checked      | int    |                    |
| content      | string |                    |
| create_time  | string |                    |
|              |        |                    |



|        | 作用     | 参数                                                         |                             |
| ------ | -------- | ------------------------------------------------------------ | --------------------------- |
| post   | 发布问卷 | title，content                                               | msg,status,id (int, option) |
| put    | 修改问卷 | （/survey/[id]),  [需要修改的字段]                           | msg,status                  |
| get    | 获得问卷 | （/survey?[id=1] [&publisher_id=1] [&name=aaa] [&name=1] [&offset=1] [&limit=10]) | msg,status,data             |
| delete | 删除问卷 | （/survey/[id])                                              | msg,status                  |



## 5./friends

下面的人信息数组一般为以下字段：

id，real_name，nick_name，age，gender，head_picture，balance，

profession，grade，phone，email



|        |                  | 参数                               | 返回值                                          |
| ------ | ---------------- | ---------------------------------- | ----------------------------------------------- |
| get    | 获取好友请求列表 | method=request   支持offset  limit | "status":string，"msg":  string，"data": 人数组 |
| post   | 申请好友         |                                    |                                                 |
| get    | 获取好友列表     | method=friends                     |                                                 |
|        | 获取请求列表     | method=request                     |                                                 |
| delete | 删除好友         |                                    |                                                 |



此地址处理好友申请，同意，删除，获取
支持GET,POST,DELETE，不支持PUT

GET请求获取好友列表或者请求列表，需要在请求参数中添加id，如/friends?id=1,支持offset和limit
同时必须添加参数method,method可以是friends或者request,如果method为friends，则为获取好友列表，如果method为request，则为
返回json：

```
{
	"status" : "success"
	"data" : 用户数组，除了password字段不存在其他字段均与数据库相同
}
```

POST为接受好友请求或者添加好友
需要在body体中提供user1_id和user2_id(不是json),body中为类似user1_id=1&user2_id=2即可
user1_id为本用户id，user2为对方id，如果接收请求或者发送添加好友请求成功，则返回json:
```
{
	"status" : "success"
}
```
失败则返回json：
```
{
	"status" : "failed",
	"msg" : "报错信息"
}
```

DELETE为删除好友，同样需要提供user1_id和user2_id，删除成功，则返回json:
```
{
	"status" : "success"
}
```
失败则返回json：
```
{
	"status" : "failed",
	"msg" : "报错信息"
}
```
## /package

快递业务相关

**目前设定全部全体可见**



[可选条件]

包裹信息包含以下字段：

| 字段               | 类型   | 备注                                |
| ------------------ | ------ | ----------------------------------- |
| id                 | int    |                                     |
| owner_id           | int    |                                     |
| owner_real_name    | string |                                     |
| owner_nick_name    | string |                                     |
| owner_Phone        | string |                                     |
| receiver_id        | int    |                                     |
| receiver_real_name | string |                                     |
| receiver_nick_name | string |                                     |
| receiver_Phone     | string |                                     |
| create_time        | string |                                     |
| reward             | int    |                                     |
| state              | int    | 0为刚发布，1为已被接单，2为确认送达 |
| note               | string |                                     |



| 方法 | 作用         | 参数                                                         | 返回值                    |
| ---- | ------------ | ------------------------------------------------------------ | ------------------------- |
| get  | 获得快递列表 | [id (快递id)] [owner_id] [receiver_id] [state] [limit] [offset] | data数组，status 字符串   |
| post | 上传快递单   | reward 整数; note 字符串                                     | status 字符串，msg 字符串 |
| PUT  | 接单         | method:receive ; id (快递id)                                 | status 字符串，msg 字符串 |
| PUT  | 确认收货     | method: confirm ; id (快递id)                                | status 字符串，msg 字符串 |
|      |              |                                                              |                           |



PUT为接单和确认收货功能，在参数列表中必须提供method，method为receive则为接单，confirm则为确认收货，如果method为receive，就必须同时提供receiver_id，并且receiver_id必须是当前登录用户，成功则返回：
```
{
	"status" : "success"
}
```
失败则返回：
```
{
	"status" : "failed",
	"msg" : "报错信息"
}
```

----
# 消息机制

/msg

1. 系统消息，什么任务已完成，钱已到账的通知之类的
2. 私聊消息

注： 

- 好友间和非好友间沟通有没有什么不同？目前没有设定不同。

- 图片什么的还不确定，缓一下。自行转码就是了。

- 以下提到的一个信息的结构为：

  - mid，uint
  - fromid，uint，发送者id
  - from，string，发送者名字
  - toid，uint，接收者id
  - to，string，接收者名字
  - msg，string，信息主体
  - time，string，发送的时间

- 因为不能在表格中换行，所以下面的表格只能新建行。

  

| 方法   | 作用                                       | 参数                                                  | 返回值                                |
| ------ | ------------------------------------------ | ----------------------------------------------------- | ------------------------------------- |
| post   | 发送消息                                   | to,  int，对象id                                      | status,bool，结果是否正常             |
|        |                                            | msg,  string，信息主体，目前设定一次信息不超过140个字 | msg,string,结果信息                   |
|        |                                            |                                                       |                                       |
| get    | 获取消息变更，所有未读信息，发送的信息被读 | history,int,是否为查看历史信息，（为非0）             | status，bool，结果是否成功            |
|        |                                            |                                                       | msg，string，结果信息                 |
|        |                                            |                                                       | unreadData,json array，新收到信息数组 |
|        |                                            |                                                       | readData,json array,已被读信息数组    |
| get    | 获取消息，个体间历史消息                   | history,bool,是否为查看历史信息，（为0，继续          | status,bool，结果是否正常             |
|        |                                            | with，int，（系统为0                                  | msg,string,结果信息                   |
|        |                                            | 支持offset 和 limit                                   | data, 信息数组                        |
| delete | 撤回单条信息，已读的不能撤回               | mid，信息id                                           | status,bool，结果是否正常             |
|        |                                            |                                                       | msg,string,结果信息                   |





# 填写问卷

/do_survey

| 方法 | 作用         | 参数                                                         | 返回值                    |
| ---- | ------------ | ------------------------------------------------------------ | ------------------------- |
| get  | 获得填写列表 | [survey_id] [recipient_id] [content] [create_time]           | data数组，status 字符串   |
| post | 上传填写结果 | 无url参数，在body中以json格式发送survey_id，recipient_id和content | status 字符串，msg 字符串 |


