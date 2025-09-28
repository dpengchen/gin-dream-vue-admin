<script setup lang="ts">
import { UserOutlined, MenuFoldOutlined, ReloadOutlined } from '@ant-design/icons-vue'
import { Sun as IconSun, Moon as IconMoon } from '@icon-park/vue-next'
import { SettingOutlined } from '@ant-design/icons-vue'
import { reactive } from 'vue'
import logo from '@/assets/logo.png'
import MenuRender from '@/components/layout/MenuRender.vue'
import { systemInfo, currentBgColor, resetTheme } from '@/utils/theme.ts'

const systemForm = reactive({
  visible: false,
  form: systemInfo,
})
const emit = defineEmits(['toggleSide'])

const changeBgBaseColor = (e: Event) => {
  systemForm.form.bgBaseColor = (e.target! as HTMLInputElement).value
}
const changeDarkBgBaseColor = (e: Event) => {
  systemForm.form.darkBgBaseColor = (e.target! as HTMLInputElement).value
}
const changePrimaryColor = (e: Event) => {
  systemForm.form.colorPrimary = (e.target! as HTMLInputElement).value
}

const changeTheme = () => {
  systemForm.form.theme = systemForm.form.theme == 'light' ? 'dark' : 'light'
}
</script>

<template>
  <a-layout-header
    class="flex justify-between items-center bg-white shadow z-10"
    :class="{
      'flex-row-reverse': systemInfo.showSide && systemInfo.sidePosition == 'right',
      '!pl-5': systemInfo.layout == '1',
      'shadow-gray-700': systemInfo.theme === 'dark',
    }"
    :style="{ 'background-color': currentBgColor }"
  >
    <div class="left flex">
      <div
        class="logo flex items-center p-1 rounded cursor-pointer"
        v-if="systemInfo.layout == '1' || !systemInfo.showSide"
        @click="emit('toggleSide')"
      >
        <img :src="logo" alt="logo" class="w-[40px] h-[40px]" />
        <span class="text-xl">DreamAdmin</span>
      </div>
      <div v-else-if="systemInfo.showSide">
        <a-button @click="emit('toggleSide')">
          <template #icon>
            <MenuFoldOutlined />
          </template>
        </a-button>
      </div>
      <div class="ml-5">
        <MenuRender mode="horizontal" v-if="!systemInfo.showSide" />
      </div>
    </div>

    <div class="right flex items-center">
      <a-button shape="circle" @click="changeTheme">
        <template #icon>
          <div class="flex items-center justify-center">
            <IconSun theme="outline" size="18" fill="#fff" v-if="systemForm.form.theme == 'dark'" />
            <IconMoon theme="outline" size="18" fill="#141414" v-else />
          </div>
        </template>
      </a-button>
      <div class="w-3"></div>
      <a-button shape="circle" @click="systemForm.visible = true" class="h-fit" size="64">
        <template #icon>
          <SettingOutlined />
        </template>
      </a-button>
      <div class="w-3"></div>
      <a-dropdown class="mb-10">
        <a-avatar size="64">
          <a-icon>
            <UserOutlined />
          </a-icon>
        </a-avatar>
        <template #overlay>
          <a-menu>
            <a-menu-item>个人中心</a-menu-item>
            <a-menu-item>退出登录</a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
    </div>
  </a-layout-header>

  <!-- 抽屉系统设置 -->
  <a-drawer v-model:open="systemForm.visible" title="系统设置" :close-icon="null" width="500px">
    <template #extra>
      <a-button @click="resetTheme">
        <template #icon> <ReloadOutlined /></template>
        重置
      </a-button>
    </template>
    <a-form :label-col="{ style: { width: '80px' } }">
      <a-form-item label="背景基础色">
        亮色
        <input type="color" :value="systemForm.form.bgBaseColor" @change="changeBgBaseColor" />
        暗色
        <input
          type="color"
          :value="systemForm.form.darkBgBaseColor"
          @change="changeDarkBgBaseColor"
        />
      </a-form-item>
      <a-form-item label="主题色">
        <div
          class="bg-[#1677ff] inline-block w-[50px] h-[19px]"
          @click="systemForm.form.colorPrimary = '#1677ff'"
        ></div>
        <div
          class="bg-[#392f5a] inline-block w-[50px] h-[19px]"
          @click="systemForm.form.colorPrimary = '#392f5a'"
        ></div>
        <div
          class="bg-[#a8e6cf] inline-block w-[50px] h-[19px]"
          @click="systemForm.form.colorPrimary = '#a8e6cf'"
        ></div>
        <div class="ml-6 inline-block">自定义</div>
        <input type="color" :value="systemForm.form.colorPrimary" @change="changePrimaryColor" />
      </a-form-item>
      <a-form-item label="模式">
        <a-radio-group v-model:value="systemForm.form.theme">
          <a-radio value="light">亮色</a-radio>
          <a-radio value="dark">暗色</a-radio>
        </a-radio-group>
      </a-form-item>
      <a-form-item label="布局">
        <a-radio-group class="w-full" v-model:value="systemForm.form.layout">
          <div class="grid grid-cols-2 gap-2">
            <div class="item">
              <div>
                <div class="bg-[#7dbcea] h-[20px]"></div>
                <div class="flex h-[40px]">
                  <div class="bg-[#3ba0e9] h-full w-1/5"></div>
                  <div class="bg-[#108ee9] h-full flex-1"></div>
                </div>
                <div class="bg-[#7dbcea] h-[20px]"></div>
              </div>
              <a-radio value="1">布局1</a-radio>
            </div>
            <div class="item">
              <div class="flex">
                <div class="bg-[#3ba0e9] h-[80px] w-1/5"></div>
                <div class="flex-1">
                  <div class="bg-[#7dbcea] h-[20px]"></div>
                  <div class="bg-[#108ee9] h-[40px]"></div>
                  <div class="bg-[#7dbcea] h-[20px]"></div>
                </div>
              </div>
              <a-radio value="2">布局2</a-radio>
            </div>
          </div>
        </a-radio-group>
      </a-form-item>
      <a-form-item label="侧边栏">
        <a-radio-group v-model:value="systemForm.form.showSide">
          <a-radio :value="true">侧边显示</a-radio>
          <a-radio :value="false">顶部显示</a-radio>
        </a-radio-group>
      </a-form-item>
      <a-form-item label="(位置)" v-if="systemForm.form.showSide">
        <a-radio-group v-model:value="systemForm.form.sidePosition">
          <a-radio value="left">左侧</a-radio>
          <a-radio value="right">右侧</a-radio>
        </a-radio-group>
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<style scoped lang="scss"></style>
