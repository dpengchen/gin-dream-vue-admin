import request from '@/api/request.ts'
import type { GenerateTable } from '@/interface/generate'
import type { PageParams } from '@/interface/core'

const baseURL = '/v1/generate/table'
//添加生成表
export const addGenerateTable = (data: GenerateTable) => {
  return request({
    url: baseURL,
    method: 'POST',
    data: data,
  })
}

//删除生成表通过ID
export const removeGenerateTableById = (id: number) => {
  return request({
    url: `${baseURL}/${id}`,
    method: 'DELETE',
  })
}

//修改生成表
export const modifyGenerateTable = (data: GenerateTable, id: number) => {
  return request({
    url: `${baseURL}/${id}`,
    method: 'PUT',
    data: data,
  })
}

//获取生成表列表
export const getGenerateTableList = (params: PageParams) => {
  return request({
    url: `${baseURL}/list`,
    method: 'GET',
    params: params,
  })
}

//获取生成表通过ID
export const getGenerateTableById = (id: number) => {
  return request({
    url: `${baseURL}/${id}`,
    method: 'GET',
  })
}

//生成代码
export const generateTable=(id:number)=>{
  return request({
    url: `${baseURL}/generate/${id}`,
    method: 'GET',
  })
}
