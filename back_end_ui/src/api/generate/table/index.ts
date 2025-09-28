import request from '@/api/request.ts'
import type { GenerateTable } from '@/interface/system/generate'

//添加生成表
export const addGenerateTable = (data: GenerateTable) => {
  return request({
    url: '/v1/generate/table',
    method: 'POST',
    data: data,
  })
}

export const getGenerateTableById = (id: number) => {
  return request({
    url: '/v1/generate/table/' + id,
    method: 'GET',
  })
}
