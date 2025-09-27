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

    return { system, setSystem }
  },
  { persist: true },
)
