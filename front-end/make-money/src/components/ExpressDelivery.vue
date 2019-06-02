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
			:is-loading="isLoading"
		></v-table>

		<div class="mt20 mb20 bold" id="paging">
			<v-pagination @page-change="pageChange" @page-size-change="pageSizeChange" :total="50" :page-size="pageSize" :layout="['total', 'prev', 'pager', 'next', 'sizer', 'jumper']" id="vpage">
			</v-pagination>
			<button id="issue" class="btn">发布快递</button>
			<button id="accept" class="btn">接受快递</button>
			<button id="submit" class="btn">确定快递</button>
			<button id="query" class="btn">快递查询</button>
		</div>
	</div>
</template>

<script>
	import tableDate from './expressTest.js'

	export default {
		name: 'ExpressDelivery',
		data() {
			return {
				pageIndex:1,
				pageSize:10,
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
					] 
				}
			}
		},

		methods: {
			getTableData() {
				this.tableConfig.tableData = tableDate.slice((this.pageIndex-1)*this.pageSize,(this.pageIndex)*this.pageSize)
			},

			pageChange(pageIndex) {
				this.pageIndex = pageIndex;
				this.getTableData()
				console.log(pageIndex)
			},

			pageSizeChange(pageSize) {
				this.pageIndex = 1;
				this.pageSize = pageSize;
				this.getTableData();
			}
		},

		created: function() {
			this.getTableData();
			console.log("Initialize successfully!");
		}
	}
</script>

<style lang="css">
	@import '../style/ExpressDelivery';
</style>