<script setup lang="ts">
import { UserOutlined, MenuFoldOutlined } from '@ant-design/icons-vue'
import { Sun as IconSun, Moon as IconMoon } from '@icon-park/vue-next'
import { SettingOutlined } from '@ant-design/icons-vue'
import { reactive } from 'vue'
import { useSystemStore } from '@/stores/system.ts'
import logo from '@/assets/logo.png'
import MenuRender from '@/components/layout/MenuRender.vue'

const systemStore = useSystemStore()
const systemForm = reactive({
  visible: false,
  form: systemStore.system,
})
const emit = defineEmits(['toggleSide'])

const changePrimaryColor = (e: Event) => {
  systemForm.form.colorPrimary = (e.target! as HTMLInputElement).value
}

const changeTheme = () => {
  systemForm.form.them = systemForm.form.them == 'light' ? 'dark' : 'light'
}
</script>

<template>
  <a-layout-header
    class="flex justify-between items-center bg-white shadow z-10"
    :class="{
      '!pl-5': systemStore.system.layout == '1',
      '!bg-white': systemStore.system.them == 'light',
      '!bg-[#07161a]': systemStore.system.them == 'dark',
      'shadow-gray-200/10': systemStore.system.them == 'dark',
    }"
  >
    <div class="left flex">
      <div
        class="logo flex items-center p-1 rounded"
        v-if="systemStore.system.layout == '1' || !systemStore.system.showSide"
        @click="emit('toggleSide')"
      >
        <img :src="logo" alt="logo" class="w-[40px] h-[40px]" />
        <span class="text-xl">DreamAdmin</span>
      </div>
      <div v-else-if="systemStore.system.showSide">
        <a-button @click="emit('toggleSide')">
          <template #icon>
            <MenuFoldOutlined />
          </template>
        </a-button>
      </div>
      <div class="ml-5">
        <MenuRender mode="horizontal" v-if="!systemStore.system.showSide" />
      </div>
    </div>

    <div class="right flex items-center">
      <a-button shape="circle" @click="changeTheme">
        <template #icon>
          <div class="flex items-center justify-center">
            <IconSun
              theme="outline"
              size="18"
              :fill="systemStore.getThemeColor('#fff', '#141414')"
              v-if="systemForm.form.them == 'dark'"
            />
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
  <a-drawer v-model:open="systemForm.visible" title="系统设置" :close-icon="null">
    <a-form :label-col="{ style: { width: '60px' } }">
      <a-form-item label="主题色">
        <input type="color" :value="systemForm.form.colorPrimary" @change="changePrimaryColor" />
      </a-form-item>
      <a-form-item label="模式">
        <a-radio-group v-model:value="systemForm.form.them">
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
      <a-form-item label="(位置)">
        <a-radio-group v-model:value="systemForm.form.sidePosition">
          <a-radio value="left">左侧</a-radio>
          <a-radio value="right">右侧</a-radio>
        </a-radio-group>
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<style scoped lang="scss"></style>
