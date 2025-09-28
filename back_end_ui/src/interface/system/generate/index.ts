import type { DreamModel } from '@/interface/core'

export interface GenerateTable extends DreamModel {
  generateVersion: string
  generateBasePath: string
  structName: string
  tableComment: string
  softDelete: boolean
  privateData: boolean
  generateColumns?: GenerateColumns[]
  relation?: string
  generateTableId?: number | null
  generateTable?: GenerateTable
}

export interface GenerateColumns {
  id?: number | null
  generateTableId?: number | null
  structName?: string | null
  jsonName?: string | null
  sqlName?: string | null
  columnLabel?: string | null
  columnType?: string | null
  inputType?: string | null
  sqlType?: string | null
  dictId?: number | null
  isEdit?: boolean | null
  isExport?: boolean | null
  isShow?: boolean | null
  isQuery?: boolean | null
  queryType?: string | null
  isSort?: boolean | null
  sortType?: string | null
  isRequired?: boolean | null
}
