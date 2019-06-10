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
		></v-table>

		<div class="mt20 mb20 bold" id="paging">
			<v-pagination @page-change="pageChange" @page-size-change="pageSizeChange" :total="50" :page-size="pageSize" :layout="['total', 'prev', 'pager', 'next', 'sizer', 'jumper']" id="vpage">
			</v-pagination>
			<button id="issue" class="btn" v-on:click="openMask">发布快递</button>
			<dialog-bar v-model="sendVal" type="danger" title="发布快递" v-on:cancel="clickCancel()" @danger="clickDanger()" @confirm="clickConfirm()" dangerText="提交"></dialog-bar>
			<button id="accept" class="btn">接受任务</button>
			<button id="submit" class="btn">确定快递</button>
			<button id="query" class="btn">快递查询</button>
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
				pageIndex:1,
				pageSize:12,
				sendVal: false,
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

			clickCancel() {
				console.log('点击了取消')
			},

			clickDanger() {
				console.log("这里是danger回调")
			},

			clickConfirm() {
				console.log("点击了confirm")
			},

			fetch_data() {
			    let url = "http://182.254.206.244:8080/package";
			    let data = [];
			    let pIndex = this.pageIndex;
			    let pSize = this.pageSize;

			    axios.get(url)
			      .then(response => {
			        let temp = response.data.data
			        console.log(temp)
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
			          	j.stuId = item.owner_Phone;
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
			          	data.push(j);
			          	
			        })
			        this.tableConfig.tableData = data.slice((pIndex-1)*pSize,(pIndex)*pSize);
			      });
			}

		},

		created: function() {
			// this.getTableData();
		},

		mounted: function() {
			this.fetch_data();
			console.log("Initialize successfully!");
		}
	}
</script>

<style lang="css" src="../style/ExpressDelivery.css"></style>