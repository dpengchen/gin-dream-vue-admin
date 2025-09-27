<script setup lang="ts">
import { watchEffect, ref } from 'vue'
import { theme } from 'ant-design-vue'
import { useSystemStore } from '@/stores/system.ts'
import HeaderLayout from '@/components/layout/HeaderLayout.vue'
import SiderLayout from '@/components/layout/SiderLayout.vue'
import FooterLayout from '@/components/layout/FooterLayout.vue'

const systemStore = useSystemStore()
const collapsed = ref(false)
const toggleSide = () => {
  collapsed.value = !collapsed.value
}

const algorithm = ref(theme.defaultAlgorithm)

//监听系统主题
watchEffect(() => {
  algorithm.value =
    systemStore.system.them === 'dark' ? theme.darkAlgorithm : theme.defaultAlgorithm
  console.log('systemStore.system.them', systemStore.system.them)
})
</script>

<template>
  <a-config-provider
    :theme="{
      algorithm: algorithm,
      token: {
        colorPrimary: systemStore.system.colorPrimary || 'balck',
      },
    }"
  >
    <a-layout class="w-full h-full">
      <template v-if="systemStore.system.layout == '1'">
        <HeaderLayout @toggleSide="toggleSide" />
        <a-layout>
          <SiderLayout
            v-if="systemStore.system.showSide && systemStore.system.sidePosition === 'left'"
            :collapsed="collapsed"
          />
          <a-layout>
            <a-layout-content
              :class="{
                '!bg-[#081d29]': systemStore.system.them == 'dark',
              }"
            >
              <router-view />
            </a-layout-content>
            <FooterLayout />
          </a-layout>
          <SiderLayout
            v-if="systemStore.system.showSide && systemStore.system.sidePosition === 'right'"
            :collapsed="collapsed"
          />
        </a-layout>
      </template>
      <template v-else>
        <SiderLayout
          v-if="systemStore.system.showSide && systemStore.system.sidePosition === 'left'"
          :collapsed="collapsed"
        />
        <a-layout>
          <HeaderLayout @toggleSide="toggleSide" />
          <a-layout-content
            :class="{
              '!bg-[#081d29]': systemStore.system.them == 'dark',
            }"
          >
            <router-view />
          </a-layout-content>
          <FooterLayout />
        </a-layout>
        <SiderLayout
          v-if="systemStore.system.showSide && systemStore.system.sidePosition === 'right'"
          :collapsed="collapsed"
        />
      </template>
    </a-layout>
  </a-config-provider>
</template>

<style scoped lang="scss"></style>
