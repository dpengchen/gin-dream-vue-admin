<script setup lang="ts">
import logo from '@/assets/logo.png'
import { useSystemStore } from '@/stores/system.ts'
import MenuRender from '@/components/layout/MenuRender.vue'

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
    :class="{
      '!bg-white': systemStore.system.them === 'light',
      '!bg-[#07161a]': systemStore.system.them == 'dark',
      'shadow-gray-200/10': systemStore.system.them == 'dark',
    }"

  >
    <div
      class="logo flex bg-white items-center justify-center m-3 rounded border border-gray-100"
      v-if="systemStore.system.layout == '2'"
    >
      <img :src="logo" alt="logo" class="w-[40px] h-[40px]" />
      <span class="text-xl" v-if="!collapsed">DreamAdmin</span>
    </div>
    <MenuRender />


  </a-layout-sider>
</template>

<style scoped lang="scss"></style>
