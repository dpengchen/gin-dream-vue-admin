<script setup lang="ts">
import logo from '@/assets/logo.png'
import { useSystemStore } from '@/stores/system.ts'
import MenuRender from '@/components/layout/MenuRender.vue'
import { currentBgColor, systemInfo } from '@/utils/theme.ts'

const systemStore = useSystemStore()
defineProps({
  collapsed: {
    type: Boolean,
    default: false,
    required: true,
  },
})
</script>

<template>
  <a-layout-sider
    class="shadow"
    :collapsed="collapsed"
    :class="{ 'shadow-gray-700': systemInfo.theme === 'dark' }"
    :style="{ 'background-color': currentBgColor }"
  >
    <div
      class="logo flex items-center justify-center m-3 rounded cursor-pointer"
      v-if="systemStore.system.layout == '2'"
    >
      <img :src="logo" alt="logo" class="w-[40px] h-[40px]" />
      <span class="text-xl" v-if="!collapsed">DreamAdmin</span>
    </div>
    <MenuRender />
  </a-layout-sider>
</template>

<style scoped lang="scss"></style>
