<script setup lang="ts">
import { useSystemStore } from '@/stores/system.ts'

const systemStore = useSystemStore()
</script>

<template>
  <a-layout-footer
    :class="{
      '!bg-[#081d29]': systemStore.system.them == 'dark',
    }"
  >
    我是底部版权部分
  </a-layout-footer>
</template>

<style scoped lang="scss"></style>
