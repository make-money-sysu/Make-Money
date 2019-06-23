<template>
    <div class="dialog" v-show="showMask">
        <div class="dialog-container" v-if="type == 'danger'">
            <div class="dialog-title">{{title}}</div>
            <div class="content">

                <form id="form-table" role="form">

                    <label class="item" for="money">{{user_money_label}}</label>
                    <input type="number" min="0" class="form-control" name="money" id="money" placeholder="请输入赏金" v-model="moneyToPost"> 

                    <div id="comment-container">
                        <textarea type="text" name="comment" id="comment" placeholder="请输入备注" v-model="commentToPost"></textarea>
                    </div>


                    <div class="btns">
                        <input type="type" v-if="type != 'confirm'" class="default-btn" @click="closeBtn" value="取消">
                        <input type="submit" v-if="type == 'danger'" class="danger-btn" @click="dangerBtn">
                        <input type="accept" v-if="type == 'accept'" class="danger-btn">
                    </div>
                    <div class="close-btn" @click="closeMask"><i class="iconfont icon-close"></i></div>    
                                    
                </form>

            </div>
        </div>

        <div class="dialog-container-accept" v-if="type == 'accept'">
            <div class="dialog-title-accept">{{title}}</div>
            <input type="type" v-if="type != 'confirm'" class="accept-btn" @click="closeBtn" value="确定">
        </div>

        <div class="dialog-container-query" v-if="type == 'query'">
            <div class="dialog-title-query">{{title}}</div>
            <div id="query-content">
                <v-table
                :width="600"
                :height="540"
                :min-height="500"
                :columns="tableConfig.columns"
                :table-data="tableConfig.tableData"
                :title-rows="tableConfig.titleRows"
                row-hover-color="#eee"
                row-click-color="#edf7ff"
                :show-vertical-border="false"
                :is-loading="isLoading"
                :paging-index="(pageIndex-1)*pageSize"
                ></v-table>
            </div>
            <v-pagination @page-change="pageChange" @page-size-change="pageSizeChange" :total="20" :page-size="pageSize" :layout="['total', 'prev', 'pager', 'next', 'sizer', 'jumper']" id="vpage2">
            </v-pagination>
            <input type="type" v-if="type != 'confirm'" class="query-btn" @click="closeBtn" value="取消">
        </div>
    </div>
</template>


<script>
import axios from 'axios';

export default {
    props: {
        value: {},
        // 类型包括 defalut 默认， danger 危险， confirm 确认，
        type:{
            type: String,
            default: 'default'
        },
        title: {
            type: String,
            default: ''
        },
        cancelText: {
            type: String,
            default: '取消'
        },
        dangerText: {
            type: String,
            default: '删除'
        },
        confirmText: {
            type: String,
            default: '确认'
        }
    },
    data(){
        return {
            pageIndex:1,
            pageSize:5,
            showMask: false,
            moneyToPost: '',
            commentToPost: '',
            currentId: 0,
            user_money_label: '请输入赏金:',
            acceptSuccess: '接受成功',
            isLoading: true,
            tableConfig: {
                multipleSort: false,
                tableData: [],
                columns: [
                    {field: "seqNum", width: 150, columnAlign: "center"},
                    {field: "accepter", width: 150, columnAlign: "center"},
                    {field: "accphone", width: 150, columnAlign: "center"},
                    {field: "state", width: 145, columnAlign: "center"}
                ],
                titleRows: [
                    [
                        {fields: ["seqNum"], title: "单号", titleAlign: "center"},
                        {fields: ["accepter"], title: "接收者", titleAlign: "center"},
                        {fields: ["accphone"], title: "接收者电话", titleAlign: "center"},
                        {fields: ["state"], title: "包裹状态", titleAlign: "center"}
                    ]
                ],
            } 
        }
    },
    methods:{
        pageChange(pageIndex) {
            this.pageIndex = pageIndex;
            this.fetch_data();
        },

        pageSizeChange(pageSize) {
            this.pageIndex = 1;
            this.pageSize = pageSize;
            this.fetch_data();
        },

        closeMask() {
            this.showMask = false;
            this.commentToPost = '';
            this.moneyToPost = '';
        },
        closeBtn() {
            this.$emit('cancel');
            this.closeMask();
        },
        dangerBtn(){
            this.$emit('danger');
            let moneyValue = this.moneyToPost;
            let commentValue = this.commentToPost;

            
            if (commentValue == '') {
                alert("请输入包裹信息！");
                window.location.href = "/ExpressDelivery";
                return false;
            }

            if (parseFloat(moneyValue) <= 0.0) {
                alert("输入的金额必须为正!");
                window.location.href = "/ExpressDelivery";
                return false;
            } 

            
            const url = "http://139.199.166.124:8080/login"
            console.log(this.moneyToPost)
            axios.defaults.withCredentials=true;

            /*
            axios.post(url, JSON.stringify({
                "id": 16340129,
                "password": "123456"
            })).then(function(response) {
                console.log("Login successfully!");
                console.log(response.data);
            }).catch(function(err) {
                console.log(err);
            });*/

            /*
            $.ajax({
            type: "get",
            dataType: 'json',
            // url: "http://182.254.206.244:8080/user",
            url: "http://139.199.166.124:8080/user/", //lt
            xhrFields: {
            withCredentials: true // 要在这里设置上传cookie
            },
            crossDomain: true,
            success: function(data){
            console.log('success');
            data = data['data'];
            console.log(data['nick_name']);
            },
            error: function(Request, status, msg){
            // console.log(Request);
            // console.log(status);
            // console.log(msg);
            console.log('fail');
            }
            });*/


            let url_post = "http://139.199.166.124:8080/package";

            axios.post(url_post, JSON.stringify({
                "reward": parseFloat(moneyValue),
                "note": commentValue
              }))
              .then(response => {
                console.log(response.data)
                window.location.href = "/ExpressDelivery";
              }).catch(function(error) {
                console.log(error);
              });             
              console.log("Submit data");

            window.location.href = "/ExpressDelivery";
            this.closeMask();
        },
        confirmBtn(){
            this.$emit('confirm');
            this.closeMask();
        },
        login_func() {
            let url = "http://139.199.166.124:8080/package"

            axios.post(url, JSON.stringify({
                "id": 666,
                "password": "123456"
            })).then(function(response) {
                console.log("Login successfully!");
                console.log(response.data);
            }).catch(function(err) {
                console.log(err);
            });
        },
        fetch_data() {
            let url = "http://139.199.166.124:8080/package?owner_id=";
            let data = [];
            let pIndex = this.pageIndex;
            let pSize = this.pageSize;

            url = url + this.currentId
            console.log(url);
            axios.get(url)
              .then(response => {
                let temp = response.data.data
                temp.forEach(item => {
                    var j = {};
                    j.seqNum = item.id;
                    if (item.state == '0') {
                        j.accepter = "尚未接受";
                        j.accphone = "尚未接受";
                    }
                    else {
                        j.accepter = item.receiver_real_name;
                        j.accphone = item.receiver_Phone;                                 
                    }
                    if (item.state == '0') {
                        j.state = 'Release';
                    } else if (item.state == '1') {
                        j.state = 'Accepted';
                    } else if (item.state == '2') {
                        j.state = 'Finish';
                    }
                    data.push(j);  
                });
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
                    data = data["data"];
                    this.currentId = data["id"];
                    //this.$set(this.currentId, data["id"]);
                },
                error: (Request, status, msg) =>{
                    console.log('fail');
                    window.location.href = "/HomePage";
                }
            });
        },

        async iterate(arr) {
          let index = 0;
          while (index < arr.length - 1) {
            await arr[index]();
            index += 1;
          }
          return arr[index]();
        }
    },
    created() {
    },
    mounted(){
        //this.login_func();
        this.showMask = this.value;
        const arr = [
            () => {
                return new Promise((resolve, reject) => {
                    setTimeout(() => {
                        this.get_session();
                        resolve();
                    }, 300);
                });
            },

            () => {
                return new Promise((resolve, reject) => {
                    setTimeout(() => {
                        this.fetch_data();
                        resolve();
                    }, 300);
                });
            }   
        ];

        this.iterate(arr);
        // this.get_session();
        // this.fetch_data();
    },
    watch:{
        value(newVal, oldVal){
            this.showMask = newVal;
        },
        showMask(val) {
            this.$emit('input', val);
        }
    },
}


</script><style lang="css" src="../style/dialog.css" scoped></style>