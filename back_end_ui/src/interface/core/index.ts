export interface DreamModel {
  id: number | null
  createTime?: string | null
  updateTime?: string | null
  createBy?: number | null
  updateBy?: number | null
  isDel?: number | null
  deleteBy?: number | null
  deleteTime?: string | null
}

export interface PageParams {
  page?:number
  size?:number
}

export interface PageResult<T> {
  list: T[]
  total: number
  totalPage: number,
  page: number,
  size: number
}
