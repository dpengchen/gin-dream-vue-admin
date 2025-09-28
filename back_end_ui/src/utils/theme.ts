import { useSystemStore } from '@/stores/system.ts'
import { computed } from 'vue'
import { SystemInfo } from '@/interface/store/system.d'

const systemStore = useSystemStore()

//当前背景基础颜色颜色
export const currentBgColor = computed(() => {
  if (systemStore.system.theme === 'dark') {
    return systemStore.system.darkBgBaseColor
  }
  return systemStore.system.bgBaseColor
})

//系统信息
export const systemInfo = systemStore.system
//获取当前模式颜色
export const getThemeColor = systemStore.getThemeColor

export const resetTheme = () => {
  systemStore.system = {
    ...new SystemInfo(),
  }
  location.reload()
}
