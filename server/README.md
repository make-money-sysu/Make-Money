<p align="center">
<img src="https://i.loli.net/2019/06/27/5d14dd80d4baa78383.png" alt="ScoreChain" title="ScoreChain" length = "1400" width="500"/><br/>

</p>
<p align="center">
<a href="https://make-money-sysu.github.io/"><img src="https://img.shields.io/badge/club-M%24%20-red.svg"></a>
<a href="https://github.com/make-money-sysu"><img src="https://img.shields.io/badge/coverage-A-%2300ccee.svg"></a>
<a href="https://logojoy.com/"><img src="https://img.shields.io/badge/logo-logojoy-FFD700.svg"></a>
<a href="https://github.com/make-money-sysu"><img src="https://img.shields.io/appveyor/ci/gruntjs/grunt.svg"></a><br/>
<a href="https://github.com/make-money-sysu"><img src="https://img.shields.io/badge/platform-Win 10-388E8E.svg"></a>
<a href="https://github.com/make-money-sysu"><img src="https://img.shields.io/badge/copyright-M$ Club-blue.svg"></a>
<a href="https://github.com/make-money-sysu"><img src="https://img.shields.io/badge/date-3~7, 2019-66a033.svg"></a>
<a href="https://996.icu"><img src="https://img.shields.io/badge/link-996.icu-8855dd.svg" alt="996.icu"></a>
<br/><br/>
</p>

We, **[M$ Club](https://github.com/make-money-sysu)**, are aiming to construct a cloud platform for college students to earn pocket money by doing simple tasks. 

**Make-Money project** is an operation-oriented service software and it's also a professional "crowdsourcing" system for college students. In our system, `Cloud Service Center` provides constantly improving business. Each student has the privilege to own a `M$ Client`. Some `Agencies`, called cows, arrange platforms for their tasks. Basic business includes assigning requirements and offering salaries by system, obtaining the extra money and using money to release the task or withdraw cash by clients. Our system does not support free transactions for the moment.

<p align="center">
<img src="https://ws2.sinaimg.cn/bmiddle/006tKfTcgy1g11alobgjij318n0u0gq9.jpg" alt="ScoreChain" title="ScoreChain" length = "1400" width="500"/><br/>
</p>

**Make-Money System** is the integration of many complex systems, including student management system, task management system, transaction management system, account management system, etc. The difficulty of our system lies in operation, that is, the discovery and delivery of tasks (business).


### 更新日志
#### 2019-4-27 添加cookie认证
在登录后会在本地存储cookie，里面存储了sessionid，在每次进行只有本用户才可以进行的操作，如进行查询好友，通过好友，创建问卷等操作时，会自动验证是否是本用户登录后所为，若验证失败api会返回Login expired

#### 2019-5-13 添加快递相关接口
详情请看api.md，数据库结构参考init.sql

#### 2019-6-2 解决跨域 && bug修正
解决跨域问题， 修正friend，package部分，补充相关api文档。乐观测试两个部分的api。

#### 2019-6-3 增加pakage返回信息
根据文档，pakage的返回值添加前端需要的一些信息

#### 2019-6-4 添加消息机制
实现消息机制，文档已更新，完成一次简单测试。

#### 2019-6-7 添加状态码信息 && 修改问卷系统字段信息 && 等等
添加状态码信息方便前端使用回调；修改问卷系统字段等信息，增加时间和状态等字段信息；删除绝大部分要提供自己id的操作，从session获取。

#### 2019-6-26 添加注销功能
添加注销功能

#### 2019-6-27 修正修改用户信息时的错误
修正修改用户信息时的错误，现在不需要提供所有的字段。不需要提供用户的id
