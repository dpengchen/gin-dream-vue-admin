import type { DreamModel, PageParams } from '@/interface/core'

export interface Dict extends DreamModel {
  dictName: string // 字典名称
  dictCode: string // 字典编码
}

// DictValue 接口对应 Go 中的 DictValue 结构体
export interface DictValue extends DreamModel {
  dictId: number // 对应 DictID
  label: string // 标签
  dictValue: number // 值
}

// DictQuery 搜索参数
export interface DictQuery extends PageParams {
  dictName?: string // 字典名称
}
