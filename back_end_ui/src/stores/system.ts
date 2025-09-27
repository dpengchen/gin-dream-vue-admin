import { defineStore } from 'pinia'
import { ref } from 'vue'
import { SystemInfo } from '@/interface/store/system.d'

export const useSystemStore = defineStore(
  'system',
  () => {
    const system = ref(new SystemInfo())

    const setSystem = (info: SystemInfo) => {
      system.value = info
    }
    const getThemeColor = (dark: string, light: string) => {
      return system.value.them === 'dark' ? dark : light
    }

    return { system, setSystem, getThemeColor }
  },
  { persist: true },
)
