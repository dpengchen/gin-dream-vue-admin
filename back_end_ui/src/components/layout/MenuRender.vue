<script setup lang="ts">
import { DesktopOutlined } from '@ant-design/icons-vue'
import { useSystemStore } from '@/stores/system.ts'

const systemStore = useSystemStore()
defineProps({
  mode: {
    type: String,
    default: 'vertical',
  },
})
</script>

<template>
  <div class="box-border">
    <a-menu
      :mode="mode"
      class="!border-0"
      :class="{
        '!bg-[#07161a]': systemStore.system.them == 'dark',
      }"
    >
      <a-menu-item>
        <DesktopOutlined />
        <span>控制台</span>
      </a-menu-item>
      <a-menu-item>
        <DesktopOutlined />
        <span>控制台</span>
      </a-menu-item>
      <a-sub-menu>
        <template #title>
          <DesktopOutlined />
          <span>个人中心</span>
        </template>
        <a-menu-item>
          <DesktopOutlined />
          <span>个人中心</span>
        </a-menu-item>
        <a-menu-item>
          <DesktopOutlined />
          <span>个人中心</span>
        </a-menu-item>
      </a-sub-menu>
    </a-menu>
  </div>
</template>

<style scoped lang="scss"></style>
