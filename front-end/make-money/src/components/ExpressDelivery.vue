<template>
	<div id="DeliveryTable">
		<v-table
			:is-vertical-resize="true"
			:columns="tableConfig.columns"
			:title-rows="tableConfig.titleRows"
			:table-data="tableConfig.tableData"
			:show-vertical-border="false"
			:width="1295"
			:height="540"
			:min-height="500"
			row-hover-color="#eee"
			row-click-color="#edf7ff"
			:paging-index="(pageIndex-1)*pageSize"
			:select-all="selectAll"
			:select-change="selectChange"
			:select-group-change="selectGroupChange"
			:is-loading="isLoading"
		></v-table>

		<div class="mt20 mb20 bold" id="paging">
			<v-pagination @page-change="pageChange" @page-size-change="pageSizeChange" :total="50" :page-size="pageSize" :layout="['total', 'prev', 'pager', 'next', 'sizer', 'jumper']" id="vpage">
			</v-pagination>
			<button id="issue" class="btn" v-on:click="openMask">发布快递</button>
			<dialog-bar v-model="sendVal" type="danger" title="发布快递" v-on:cancel="clickCancel()" @danger="clickDanger()" @confirm="clickConfirm()" dangerText="提交"></dialog-bar>
			<button id="accept" class="btn" v-on:click="openAccept">接受任务</button>
			<dialog-bar v-model="AcceptedVal" type="accept" v-on:cancel="clickCancel()" dangerText="提交" title="接受成功"></dialog-bar>
			<button id="submit" class="btn" v-on:click="submitdata">确定快递</button>
			<button id="query" class="btn" v-on:click="openQuery">快递查询</button>
			<dialog-bar v-model="queryVal" type="query" title="查询快递" v-on:cancel="clickCancel()" @danger="clickDanger()" @confirm="clickConfirm()"></dialog-bar>
		</div>
	</div>
</template>

<script>
	import dialogBar from './dialog.vue'
	// import tableDate from './expressTest.js'
	import axios from 'axios';

	export default {
		name: 'ExpressDelivery',
		components: {
			'dialog-bar': dialogBar,
		},
		data() {
			return {
				isLoading: true,
				pageIndex:1,
				pageSize:12,
				sendVal: false,
				AcceptedVal: false,
				queryVal: false,
				currentUser: "",
				currentPhone: 0,
				currentId: "",
				tableConfig: {
					multipleSort: false,
					tableData: [],
					columns: [
						{field: "option", width: 75, titleAlign: "center", columnAlign: "center", type: "selection"},
						{field: "seqNum", width: 125, columnAlign: "center"},
						{field: "name", width: 125, columnAlign: "center"},
						{field: "stuId", width: 125, columnAlign: "center"},
						{field: "tel", width: 150, columnAlign: "center"},
						{field: "date", width: 150, columnAlign: "center"},
						{field: "reward", width: 125, columnAlign: "center"},
						{field: "state", width: 125, columnAlign: "center"},
						{field: "comment", width: 290, columnAlign: "center"}
					],
					titleRows: [
						[
							{fields: ["option"], titleAlign: "center", type:"selection"},
							{fields: ["seqNum"], title: "单号", titleAlign: "center"},
							{fields: ["name"], title: "姓名", titleAlign: "center"},
							{fields: ["stuId"], title: "学号", titleAlign: "center"},
							{fields: ["tel"], title: "手机号", titleAlign: "center"},
							{fields: ["date"], title: "日期", titleAlign: "center"},
							{fields: ["reward"], title: "赏金", titleAlign: "center"},
							{fields: ["state"], title: "状态", titleAlign: "center"},
							{fields: ["comment"], title: "备注", titleAlign: "center"}
						]
					],
				}
			}
		},

		methods: {
			/*
			getTableData() {
				this.fetch_data();
				// console.log(user_data);
				// console.log(user_data.slice(0, 5));
				// 
				// this.tableConfig.tableData = user_data.slice((this.pageIndex-1)*this.pageSize,(this.pageIndex)*this.pageSize);
			},*/

			pageChange(pageIndex) {
				this.pageIndex = pageIndex;
				this.fetch_data();
			},

			pageSizeChange(pageSize) {
				this.pageIndex = 1;
				this.pageSize = pageSize;
				this.fetch_data();
			},

			openMask(index) {
				this.sendVal = true;
			},

			putDataToServer(url, data) {
				axios.put(url, JSON.stringify({
					"method" : "received",
					"id" : data.seqNum
				})).then(res => {
					for (var i = 0; i < this.tableConfig.tableData.length; i++) {
						if (this.tableConfig.tableData[i].seqNum == data.seqNum) {
							this.tableConfig.tableData[i].state = "Accepted";
							break;
						}
					}
					console.log(res.data);
				}).catch((err) => {
					console.log(err);
				});
			},

			openAccept() {
				let putdata = [];

				for (var i = 0; i < this.tableConfig.tableData.length; i++) {
										console.log("Start");
					if (this.tableConfig.tableData[i]._checked == true) {
						if (this.currentId == this.tableConfig.tableData[i].stuId) {
							alert("不能选择自己的快递！");
							return false;
						}
						if (this.tableConfig.tableData[i].state == "Accepted") {
							alert("已经有人接受此任务，请重新选择！");
							return false;
						}
						console.log("push success!");
						putdata.push(this.tableConfig.tableData[i]);
					}
				}
				axios.defaults.withCredentials=true;
				console.log("show");
				console.log(putdata[0].seqNum);
				this.AcceptedVal = true;

				for (var i = 0; i < putdata.length; i++) {
					var url = "http://139.199.166.124:8080/package?method=receive&id=";
					url = url + putdata[i].seqNum;
					console.log(url);
					axios.put(url).then(res => {
						for (let i = 0; i < this.tableConfig.tableData.length; i++) {
							for (let j = 0; j < putdata.length; j++) {
								// console.log(this.tableConfig.tableData[i].seqNum);
								// console.log(putdata[i]);
								if (this.tableConfig.tableData[i].seqNum == putdata[j].seqNum) {
									this.tableConfig.tableData[i].state = "Accepted";
									break;
								}						
							}
						}
						// console.log(res.data);
						window.location.href = "/ExpressDelivery";
					}).catch((err) => {
						console.log(err);
					});
				}
			},

			openQuery() {
				this.queryVal = true;
			},

			selectAll(selection) {
				console.log("H");
			},

			selectChange(selection, rawData) {
				console.log(selection);
				if (rawData._checked == true) {
					for (var i = 0; i < this.tableConfig.tableData.length; i++) {
						if (rawData.seqNum == this.tableConfig.tableData[i].seqNum) {
							this.tableConfig.tableData[i]._checked = false;
						}
					}
				}
				else {
					for (var i = 0; i < this.tableConfig.tableData.length; i++) {
						if (rawData.seqNum == this.tableConfig.tableData[i].seqNum) {
							this.tableConfig.tableData[i]._checked = true;
						}
					}
				}
			},

			selectGroupChange(selection) {
				console.log("Happy");
			},

			clickCancel() {
				console.log('点击了取消')
			},

			clickDanger() {
				console.log("这里是danger回调")
			},

			clickConfirm() {
				console.log("点击了confirm")
			},

			submitdata() {
				let putdata = [];
				console.log(this.currentUser);
				console.log(this.currentId);
				console.log(this.currentPhone);

				for (var i = 0; i < this.tableConfig.tableData.length; i++) {
					if (this.tableConfig.tableData[i]._checked == true) {
						console.log("Start");
						console.log(this.currentUser);
						console.log(this.currentId);
						console.log(this.currentPhone);
						//console.log(this.tableConfig.tableData[i].stuId);
						if (this.currentId != this.tableConfig.tableData[i].stuId) {
							alert("不能帮其他人确定快递！");
							return false;
						}
						if (this.tableConfig.tableData[i].state == "Release") {
							alert("快递还没有人帮忙接受，无法确认！");
							return false;
						}
						if (this.tableConfig.tableData[i].state == "Finish") {
							alert("快递已完成！");
							return false;
						}					
						putdata.push(this.tableConfig.tableData[i]);
						console.log("push success!");
					}
				}
				axios.defaults.withCredentials=true;	

				for (var i = 0; i < putdata.length; i++) {
					var url = "http://139.199.166.124:8080/package?method=confirm&id=";
					url = url + putdata[i].seqNum;
					console.log(url);
					axios.put(url).then(res => {
						for (let i = 0; i < this.tableConfig.tableData.length; i++) {
							for (let j = 0; j < putdata.length; j++) {
								// console.log(this.tableConfig.tableData[i].seqNum);
								// console.log(putdata[i]);
								if (this.tableConfig.tableData[i].seqNum == putdata[j].seqNum) {
									this.tableConfig.tableData[i].state = "Finish";
									break;
								}						
							}
						}
						// console.log(res.data);
						window.location.href = "/ExpressDelivery";
					}).catch((err) => {
						console.log(err);
					});
				}	
			},

			fetch_data() {
			    let url = "http://139.199.166.124:8080/package";
			    let data = [];
			    let pIndex = this.pageIndex;
			    let pSize = this.pageSize;

			    axios.get(url)
			      .then(response => {
			        let temp = response.data.data
			        // console.log(temp)
			        // Handle list
			        temp.forEach(item => {
			        	var j = {};
			          	// console.log(item.id); // 获取num
			          	// console.log(item.owner_real_name);
			          	// console.log(item.owner_id)
			          	// console.log(item.owner_Phone); //获取手机
			          	// console.log(item.create_time); // 获得时间
			          	// console.log(item.reward); // 获取赏金
			          	// console.log(item.state); // 获得状态
			          	// console.log(item.note); // 获得备注
			          	j.seqNum = item.id;
			          	j.name = item.owner_real_name;
			          	j.stuId = item.owner_id;
			          	j.tel = item.owner_Phone;
			          	j.date = item.create_time;
			          	j.reward = item.reward;
			          	if (item.state == '0') {
			          		j.state = 'Release';
			          	} else if (item.state == '1') {
			          		j.state = 'Accepted';
			          	} else if (item.state == '2') {
			          		j.state = 'Finish';
			          	}
			          	j.comment = item.note;
			          	j._checked = false;
			          	data.push(j);
			          	
			        })
			        this.isLoading = false;
			        this.tableConfig.tableData = data.slice((pIndex-1)*pSize,(pIndex)*pSize);
			      });
			},

			get_session() {
				$.ajax({
					type: "get",
					dataType: 'json',
					// url: "http://182.254.206.244:8080/user",
					url: "http://139.199.166.124:8080/user/", //lt
					xhrFields: {
						withCredentials: true // 要在这里设置上传cookie
					},
					crossDomain: true,
					success: (data) => {
						console.log('success');
						data = data["data"];
						this.currentId = data["id"];
						this.currentUser = data["real_name"];
						this.currentPhone = data["phone"];
						console.log(this.currentUser);
						console.log(this.currentId);
						console.log(this.currentPhone);
						//this.$set(this.currentId, data["id"]);
					},
					error: (Request, status, msg) =>{
						console.log('fail');
						window.location.href = "/HomePage";
					}
				});
			}

		},

		created: function() {
			// this.getTableData();
		},

		mounted: function() {
			this.fetch_data();
			this.get_session();
			console.log("Initialize successfully!");
		}
	}
</script>

<style lang="css" src="../style/ExpressDelivery.css"></style>