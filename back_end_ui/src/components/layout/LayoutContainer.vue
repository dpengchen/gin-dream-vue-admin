<script setup lang="ts">
import { currentBgColor, systemInfo } from '@/utils/theme.js'

defineProps({
  showTitle: {
    type: Boolean,
    default: false,
  },
  title: {
    type: String,
    default: '标题',
  },
  hideDivider: {
    type: Boolean,
    default: true,
  },
})
</script>

<template>
  <div class="m-3.5">
    <a-layout
      class="p-2.5 rounded shadow-xs"
      :class="{
        'shadow-gray-700': systemInfo.theme === 'dark',
      }"
      :style="{
        backgroundColor: currentBgColor,
      }"
    >
      <div class="flex justify-between items-center p-2" v-if="showTitle">
        <div class="left text-md">
          <slot name="left"></slot>
        </div>
        <div class="right">
          <slot name="right" />
        </div>
      </div>
      <div class="border-b border-b-gray-300 mb-5" v-if="!hideDivider"></div>
      <slot />
    </a-layout>
  </div>
</template>

<style scoped lang="scss"></style>
