<script setup lang="ts">
import { DesktopOutlined } from '@ant-design/icons-vue'
import { useSystemStore } from '@/stores/system.ts'
import { Code as IconCode ,Share as IconShare} from '@icon-park/vue-next'
import IconPackAdapter from '@/components/IconPackAdapter.vue'

const systemStore = useSystemStore()
defineProps({
  mode: {
    type: String,
    default: 'inline',
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
      <a-menu-item key="1">
        <DesktopOutlined />
        <span>控制台</span>
      </a-menu-item>
      <a-menu-item key="2">
        <IconPackAdapter>
          <IconCode size="14"/>
        </IconPackAdapter>
        <span>代码生成</span>
      </a-menu-item>
      <a-menu-item key="3">
        <IconPackAdapter>
          <IconShare size="14"/>
        </IconPackAdapter>
        <span>关于我们</span>
      </a-menu-item>
    </a-menu>
  </div>
</template>

<style scoped lang="scss"></style>
