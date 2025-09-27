<script setup lang="ts">
import { UserOutlined, LockOutlined, WechatOutlined, GithubOutlined } from '@ant-design/icons-vue'
import logo from '@/assets/logo.png'
</script>

<template>
  <div class="w-full h-full flex items-center justify-center loginContainer">
    <div
      class="absolute origin-center 2xl:translate-x-[-55%] translate-x-[-60%] translate-y-[-50%] left-0 top-0 rotate-[15deg] w-[250vh] h-[250vh] bg-white hidden lg:block transition duration-[300]"
    ></div>
    <div
      class="flex flex-col items-center 2xl:-translate-x-full xl:-translate-x-1/2 transition duration-300 bg-white p-2 rounded py-4"
    >
      <div class="logo z-10">
        <img :src="logo" alt="logo" class="w-[80px] h-[80px] z-10" />
      </div>
      <div class="name text-2xl font-bold mb-4 z-10">Gin-Dream-Vue-Admin</div>
      <a-form class="w-[350px]">
        <a-form-item>
          <a-input>
            <template #prefix>
              <UserOutlined />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item>
          <a-input>
            <template #prefix>
              <LockOutlined />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item>
          <a-row :gutter="10">
            <a-col :span="12">
              <a-input></a-input>
            </a-col>
            <a-col :span="12">
              <img src="" alt="验证码" class="w-full" />
            </a-col>
          </a-row>
        </a-form-item>
        <a-form-item>
          <a-button class="w-full" type="primary">登录</a-button>
        </a-form-item>
      </a-form>
      <a-space>
        <a-button class="flex items-center">
          <template #icon>
            <GithubOutlined />
          </template>
          Github
        </a-button>
        <a-button class="flex items-center">
          <template #icon>
            <WechatOutlined />
          </template>
          Wechat
        </a-button>
      </a-space>
    </div>
  </div>
</template>

<style scoped lang="scss">
.loginContainer {
  background-image: url('@/assets/bg.png');
  background-size: cover;
  position: relative;
  overflow: hidden;
}
</style>
