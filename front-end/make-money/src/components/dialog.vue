<template>
    <div class="dialog" v-show="showMask">
        <div class="dialog-container">
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
                        <div v-if="type == 'confirm'" class="confirm-btn" @click="confirmBtn">
                            {{confirmText}}
                        </div>
                    </div>
                    <div class="close-btn" @click="closeMask"><i class="iconfont icon-close"></i></div>    
                                    
                </form>

            </div>
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
            showMask: false,
            moneyToPost: '',
            commentToPost: '',
            user_money_label: '请输入赏金:'        
        }
    },
    methods:{
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
            return true;
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
        }
    },
    created() {
    },
    mounted(){
        //this.login_func();
        this.showMask = this.value;
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