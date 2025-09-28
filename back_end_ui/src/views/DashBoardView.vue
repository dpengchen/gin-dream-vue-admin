<script setup lang="ts">
import { watchEffect, ref } from 'vue'
import { theme } from 'ant-design-vue'
import HeaderLayout from '@/components/layout/HeaderLayout.vue'
import SiderLayout from '@/components/layout/SiderLayout.vue'
import FooterLayout from '@/components/layout/FooterLayout.vue'
import { systemInfo } from '@/utils/theme.ts'

const collapsed = ref(false)
const toggleSide = () => {
  collapsed.value = !collapsed.value
}

const algorithm = ref(theme.defaultAlgorithm)

//监听系统主题
watchEffect(() => {
  algorithm.value = systemInfo.theme === 'dark' ? theme.darkAlgorithm : theme.defaultAlgorithm
})
</script>

<template>
  <a-config-provider
    :theme="{
      algorithm: algorithm,
      token: {
        colorPrimary: systemInfo.colorPrimary || 'black',
        colorBgBase:
          systemInfo.theme == 'light' ? systemInfo.bgBaseColor : systemInfo.darkBgBaseColor,
      },
    }"
  >
    <a-layout class="w-full h-full">
      <template v-if="systemInfo.layout == '1'">
        <HeaderLayout @toggleSide="toggleSide" />
        <a-layout>
          <SiderLayout
            v-if="systemInfo.showSide && systemInfo.sidePosition === 'left'"
            :collapsed="collapsed"
          />
          <a-layout>
            <a-layout-content :class="{}">
              <router-view />
            </a-layout-content>
            <FooterLayout />
          </a-layout>
          <SiderLayout
            v-if="systemInfo.showSide && systemInfo.sidePosition === 'right'"
            :collapsed="collapsed"
          />
        </a-layout>
      </template>
      <template v-else>
        <SiderLayout
          v-if="systemInfo.showSide && systemInfo.sidePosition === 'left'"
          :collapsed="collapsed"
        />
        <a-layout>
          <HeaderLayout @toggleSide="toggleSide" />
          <a-layout-content :class="{}">
            <router-view />
          </a-layout-content>
          <FooterLayout />
        </a-layout>
        <SiderLayout
          v-if="systemInfo.showSide && systemInfo.sidePosition === 'right'"
          :collapsed="collapsed"
        />
      </template>
    </a-layout>
  </a-config-provider>
</template>

<style scoped lang="scss"></style>
