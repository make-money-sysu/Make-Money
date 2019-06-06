<template>
    <div class="dialog" v-show="showMask">
        <div class="dialog-container">
            <div class="dialog-title">{{title}}</div>
            <div class="content">
                <form id="form-table" method="post" role="form">
                    <div class="form-group" id="row-4">
                        <label class="item" for="date">{{user_date_label}}</label>
                        <div id="col-4">
                            <input type="text" class="form-control" name="date" id="date" placeholder="请输入日期" v-model="dateToPost">
                            <span id="isCan" style="color: red;"></span>
                        </div>
                    </div>     
                    <div class="form-group" id="row-5">
                        <label class="item" for="money">{{user_money_label}}</label>
                        <div id="col-5">
                            <input type="text" class="form-control" name="money" id="money" placeholder="请输入赏金" v-model="moneyToPost">
                            <span id="isCan" style="color: red;"></span>
                        </div>
                    </div>  
                    <div id="row-6">
                        <div id="col-6">
                            <textarea type="text" name="comment" id="comment" placeholder="请输入备注" v-model="commentToPost"></textarea>
                            <span id="isCan" style="color: red;"></span>
                        </div>
                    </div> 
                    <div class="btns">
                        <div v-if="type != 'confirm'" class="default-btn" @click="closeBtn">
                            {{cancelText}}
                        </div>
                        <input type="submit" v-if="type == 'danger'" class="danger-btn" @click="dangerBtn">
                            {{dangerText}}
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
        },
    },
    data(){
        return{
            showMask: false,
            dateToPost: "",
            moneyToPost: 0,
            commentToPost: "",
            user_date_label: '日期',
            user_money_label: '赏金'        }
    },
    methods:{
        closeMask() {
            this.showMask = false;
        },
        closeBtn() {
            this.$emit('cancel');
            this.closeMask();
        },
        dangerBtn(){

            $.ajax({
                type: "post",
                dataType: "json",
                url: "http://139.199.166.124:8080/package",
                crossDomain: true,
                data: JSON.stringify({
                    "create_time": this.dateToPost,
                    "reward": this.moneyToPost,
                    "noted": this.commentToPost,
                    "state": 0
                }),
                success: function(data) {
                    console.log("Post success!");
                    console.log(data);
                },
                error: function(Request, status, msg) {
                    console.log(Request);
                    console.log(status);
                    console.log(msg);
                    console.log("Fail!")
                }
            });

            console.log("Submit data")
            this.closeMask();
            return true;
        },
        confirmBtn(){
            this.$emit('confirm');
            this.closeMask();
        },
        login_func() {
            $.ajax({
                type: "post",
                dataType: "json",
                url: "http://139.199.166.124:8080/login",
                crossDomain: true,
                data: JSON.stringify({
                    "id": 666,
                    "password": "123456"
                }),
                success: function(data) {
                    console.log("Login successfully!");
                },
                error: function(Request, status, msg) {
                    console.log(Request);
                    console.log(status);
                    console.log(msg);
                    console.log("Fail");
                }
            })
        }
    },
    created() {
        this.login_func();
    },
    mounted(){
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
</script>

<style lang="css" src="../style/dialog.css" scoped></style>