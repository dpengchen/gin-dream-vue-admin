import request from '@/api/request.ts'
import type { Dict, DictQuery } from '@/interface/system/dict'

const baseURL = '/v1/system/dict'
//添加生成表
export const addDict = (data: Dict) => {
  return request({
    url: baseURL,
    method: 'POST',
    data: data,
  })
}

//删除生成表通过ID
export const removeDictById = (ids: string[] | number[]) => {
  return request({
    url: `${baseURL}`,
    method: 'DELETE',
    data: ids,
  })
}

//修改生成表
export const modifyDict = (data: Dict, id: number) => {
  return request({
    url: `${baseURL}/${id}`,
    method: 'PUT',
    data: data,
  })
}

//获取生成表通过ID
export const getDictList = (params: DictQuery) => {
  return request({
    url: `${baseURL}/list`,
    method: 'GET',
    params: params,
  })
}

//获取生成表通过ID
export const getDictById = (id: number) => {
  return request({
    url: `${baseURL}/${id}`,
    method: 'GET',
  })
}

//导出生成表 如果有选择ids则导出对应的记录，如果没有则导出符合搜索条件的所有记录
export const exportDict = (params: DictQuery, ids: number[] | string[]) => {
  return request({
    url: `${baseURL}/export`,
    method: 'POST',
    data: ids,
    params: params,
  })
}
