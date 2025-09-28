<script setup lang="ts">
import { DesktopOutlined, SettingOutlined } from '@ant-design/icons-vue'
import { Share as IconShare, Cpu as IconCPU, Notebook as IconNotebook } from '@icon-park/vue-next'
import IconPackAdapter from '@/components/IconPackAdapter.vue'
import router from '@/router'
import { currentBgColor, systemInfo } from '@/utils/theme.ts'

defineProps({
  mode: {
    type: String,
    default: 'inline',
  },
})

const selectMenu = ({ key, selectedKeys }: any) => {
  systemInfo.selectedKeys = selectedKeys
  router.push({
    path: key,
  })
}
</script>

<template>
  <div class="box-border">
    <a-menu
      :mode="mode"
      class="!border-0"
      v-model:selected-keys="systemInfo.selectedKeys"
      @select="selectMenu"
      :style="{ 'background-color': currentBgColor }"
    >
      <a-menu-item key="1" title="控制台">
        <DesktopOutlined />
        <span>控制台</span>
      </a-menu-item>
      <a-sub-menu key="1">
        <template #title>
          <span>
            <SettingOutlined />
            <span>系统工具</span>
          </span>
        </template>
        <a-menu-item key="/dict">
          <IconPackAdapter>
            <IconNotebook />
          </IconPackAdapter>
          <span>字典管理</span>
        </a-menu-item>
        <a-menu-item key="/generate">
          <IconPackAdapter>
            <IconCPU />
          </IconPackAdapter>
          <span>代码生成</span>
        </a-menu-item>
      </a-sub-menu>
      <a-menu-item key="3" title="关于我们">
        <IconPackAdapter>
          <IconShare size="14" />
        </IconPackAdapter>
        <span>关于我们</span>
      </a-menu-item>
    </a-menu>
  </div>
</template>

<style scoped lang="scss"></style>
