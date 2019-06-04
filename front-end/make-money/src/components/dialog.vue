<template>
    <div class="dialog" v-show="showMask">
        <div class="dialog-container">
            <div class="dialog-title">{{title}}</div>
            <div class="content">
                <form id="form-table" method="post" role="form">
                    <div class="form-group" id="row-1">
                        <label class="item" for="name">用户名：</label>
                        <div id="col-1">
                            <input type="text" class="form-control" name="name" id="name" placeholder="请输入用户名">
                            <span id="isCan" style="color: red;"></span>
                        </div>
                    </div>
                    <div class="form-group" id="row-2">
                        <label class="item" for="number">学号：</label>
                        <div id="col-2">
                            <input type="text" class="form-control" name="number" id="number" placeholder="请输入学号">
                            <span id="isCan" style="color: red;"></span>
                        </div>
                    </div>        
                    <div class="form-group" id="row-3">
                        <label class="item" for="phone">手机号：</label>
                        <div id="col-3">
                            <input type="text" class="form-control" name="phone" id="phone" placeholder="请输入手机号">
                            <span id="isCan" style="color: red;"></span>
                        </div>
                    </div>  
                    <div class="form-group" id="row-4">
                        <label class="item" for="date">日期：</label>
                        <div id="col-4">
                            <input type="text" class="form-control" name="date" id="date" placeholder="请输入日期">
                            <span id="isCan" style="color: red;"></span>
                        </div>
                    </div>     
                    <div class="form-group" id="row-5">
                        <label class="item" for="money">赏金：</label>
                        <div id="col-5">
                            <input type="text" class="form-control" name="money" id="money" placeholder="请输入赏金">
                            <span id="isCan" style="color: red;"></span>
                        </div>
                    </div>  
                    <div id="row-6">
                        <div id="col-6">
                            <textarea type="text" name="comment" id="comment" placeholder="请输入备注"></textarea>
                            <span id="isCan" style="color: red;"></span>
                        </div>
                    </div>                                                                                        
                </form>
            </div>
            <div class="btns">
                <div v-if="type != 'confirm'" class="default-btn" @click="closeBtn">
                    {{cancelText}}
                </div>
                <div v-if="type == 'danger'" class="danger-btn" @click="dangerBtn">
                    {{dangerText}}
                </div>
                <div v-if="type == 'confirm'" class="confirm-btn" @click="confirmBtn">
                    {{confirmText}}
                </div>
            </div>
            <div class="close-btn" @click="closeMask"><i class="iconfont icon-close"></i></div>
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
            text1: '姓名',
            text1: '学号',
            text1: '手机号',
            text1: '日期',
            text1: '赏金',
            text1: '备注'
        }
    },
    methods:{
        closeMask(){
            this.showMask = false;
        },
        closeBtn(){
            this.$emit('cancel');
            this.closeMask();
        },
        dangerBtn(){
            this.$emit('danger');
            this.closeMask();
        },
        confirmBtn(){
            this.$emit('confirm');
            this.closeMask();
        }
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