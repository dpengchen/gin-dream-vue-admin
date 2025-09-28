<script setup lang="ts">
import LayoutContainer from '@/components/layout/LayoutContainer.vue'
import { reactive } from 'vue'
import { PlusOutlined } from '@ant-design/icons-vue'
import type { Dict } from '@/interface/system/dict'

const dictForm = reactive({
  form: {} as Dict, //表单数据
  rules: {}, //校验规则
  visible: false, //drawer 抽屉显示
  title: '新增字典', //标题
  reset: () => {
    //重置表单
    dictForm.form = {
      id: null,
      dictCode: '',
      dictName: '',
    }
  },
  edit: (record: any) => {
    //新增、修改
    if (record.id) {
      //修改
      dictForm.title = '修改字典'
      return
    }
    dictForm.title = '新增字典'
  },
  remove: (id: number) => {
    console.log(`删除id = ${id}`)
  },
})
const table = reactive({
  list: [] as Dict[],
  columns: [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
    },
    {
      title: '字典名称',
      dataIndex: 'dictName',
      key: 'dictName',
    },
    {
      title: '字典编码',
      dataIndex: 'dictCode',
      key: 'dictCode',
    },
    {
      title: '创建时间',
      dataIndex: 'createTime',
      key: 'createTime',
    },
    {
      title: '更新时间',
      dataIndex: 'updateTime',
      key: 'updateTime',
    },
    {
      title: '创建人',
      dataIndex: 'createByNickname',
      key: 'createByNickname',
    },
    {
      title: '更新人',
      dataIndex: 'updateByNickname',
      key: 'updateByNickname',
    },
    {
      title: '操作',
      dataIndex: 'operation',
      key: 'operation',
      width: 200,
      scopedSlots: { customRender: 'bodyCell' },
    },
  ],
})
</script>

<template>
  <div>
    <!--  搜索区域  -->
    <LayoutContainer></LayoutContainer>

    <!--  table区域 -->
    <LayoutContainer show-title>
      <template #right>
        <a-button @click="dictForm.edit">
          <template #icon>
            <PlusOutlined />
          </template>
          新增字典
        </a-button>
      </template>
      <!--      <a-table :columns="table.columns" :data-source="table.list" :scroll="{ x: 2000 }">-->
      <a-table :columns="table.columns" :data-source="table.list">
        <template #bodyCell="{ record, column }">
          <template v-if="column.key === 'operation'">
            <a-space>
              <a-button @click="dictForm.edit(record)">修改</a-button>
              <a-button @click="dictForm.remove(record.id)">删除</a-button>
            </a-space>
          </template>
        </template>
      </a-table>
    </LayoutContainer>
  </div>
</template>

<style scoped lang="scss"></style>
