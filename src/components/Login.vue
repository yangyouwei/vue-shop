<template>
    <div class="login_container">
        <div class="login_box">
            <div class="avatar_box">
                <img src="@/assets/logo.png" alt="">
            </div>
            <el-form ref="loginFormRef" :rules="rules" :model="form" label-width="0px" class="login_form">
                <el-form-item prop="name">
                    <el-input v-model="form.name" placeholder="请输入用户名" prefix-icon="el-icon-user"></el-input>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input v-model="form.password" type="password" placeholder="请输入密码" prefix-icon="el-icon-unlock"></el-input>
                </el-form-item>
                <el-form-item class="btns">
                    <el-button type="primary" @click="submitForm('loginFormRef')">登录</el-button>
                    <el-button type="info" @click="resetForm('loginFormRef')">重置</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script>
    export default {
        data(){
            return{
                form:{
                    name:'',
                    password:''
                },
                rules:{
                    name:[
                        { required: true, message: '请输入用户名', trigger: 'blur' },
                        { min: 3, max: 15, message: '长度在 3 到 15 个字符', trigger: 'blur' }
                    ],
                    password:[
                        { required: true, message: '请输入密码', trigger: 'blur' },
                        { min: 3, max: 12, message: '长度在 3 到 12 个字符', trigger: 'blur' }
                    ]
                }
            }
        },
        methods: {
            submitForm(formName){
                this.$refs[formName].validate(async (valid) => {
                    if (!valid) return;
                    const {data:res} = await this.$http.post('login',this.form)
                    console.log(res)
                });
            },
            resetForm(formName) {
                this.$refs[formName].resetFields();
            }
        }
    }
</script>

<style lang="less" scoped>
    .login_container {
        background-color: #2b4b6b;
        height: 100%;
    }

    .login_box {
        width: 450px;
        height: 300px;
        background-color: #fff;
        border-radius: 3px;
        position: absolute;
        left: 50%;
        top: 50%;
        transform: translate(-50%, -50%);
    }
    .avatar_box {
        height: 130px;
        width: 130px;
        border: 1px solid #eee;
        border-radius: 50%;
        padding: 10px;
        box-shadow: 0 0 10px #ddd;
        position: absolute;
        left: 50%;
        transform: translate(-50%, -50%);
        background-color: #fff;
    img {
        width: 100%;
        height: 100%;
        border-radius: 50%;
        background-color: #eee;
    }
    }
    .login_form{
        position: absolute;
        bottom: 0;
        width: 100%;
        padding: 0 20px;
        box-sizing: border-box;
    }
    .btns{
        display: flex;
        justify-content: flex-end;
    }
</style>