<script setup lang="ts">
import LayoutContainer from '@/components/layout/LayoutContainer.vue'
import { reactive, ref } from 'vue'
import { Code as IconCode } from '@icon-park/vue-next'
import { PlusOutlined, SaveOutlined, EyeOutlined } from '@ant-design/icons-vue'
import IconPackAdapter from '@/components/IconPackAdapter.vue'

const table = reactive({
  list: [],
  remove: (index: number) => {
    table.list.splice(index, 1)
  },
  columns: [
    {
      title: '字段标签',
      dataIndex: 'columnLabel',
      key: 'columnLabel',
      fixed: true,
      align: 'center',
    },
    {
      title: '字段名',
      dataIndex: 'structName',
      key: 'structName',
      align: 'center',
    },
    {
      title: '字段类型',
      dataIndex: 'columnType',
      key: 'columnType',
      align: 'center',
    },
    {
      title: '字段长度',
      dataIndex: 'columnLen',
      key: 'columnLen',
      align: 'center',
    },
    {
      title: '新建/编辑',
      dataIndex: 'isEdit',
      width: 120,
      key: 'isEdit',
      align: 'center',
    },
    {
      title: '列表',
      dataIndex: 'isShow',
      width: 65,
      key: 'isShow',
      align: 'center',
    },
    {
      title: '必填',
      dataIndex: 'isRequired',
      width: 65,
      key: 'isRequired',
      align: 'center',
    },
    {
      title: '排序',
      dataIndex: 'isSort',
      width: 65,
      key: 'isSort',
      align: 'center',
    },
    {
      title: '排序类型',
      dataIndex: 'sortType',
      key: 'sortType',
      align: 'center',
    },
    {
      title: '查询',
      dataIndex: 'isQuery',
      width: 65,
      key: 'isQuery',
      align: 'center',
    },
    {
      title: '查询条件',
      dataIndex: 'queryType',
      key: 'queryType',
      align: 'center',
    },
    {
      title: '导出',
      dataIndex: 'isExport',
      width: 65,
      key: 'isExport',
      align: 'center',
    },
    {
      title: '操作',
      key: 'operation',
      fixed: 'right',
      align: 'center',
    },
  ],
})

const tableForm = reactive({
  form: {
    id: null,
    generateVersion: '',
    generateBasePath: '',
    structName: '',
    tableComment: '',
    softDelete: false,
    generateColumns: [],
    relation: '',
    generateTableId: null,
  },
  rules: {
    generateVersion: [{ required: true, message: '请输入版本号', trigger: 'blur' }],
    generateBasePath: [{ required: true, message: '请输入生成路径', trigger: 'blur' }],
    structName: [{ required: true, message: '请输入结构体名', trigger: 'blur' }],
    tableComment: [{ required: true, message: '请输入表注释', trigger: 'blur' }],
    softDelete: [{ required: true, message: '请选择是否软删除', trigger: 'blur' }],
    generateColumns: [{ required: true, message: '请选择字段', trigger: 'blur' }],
    relation: [{ required: true, message: '请选择关联关系', trigger: 'blur' }],

  },
})

const fieldFormRef = ref()
const fieldForm = reactive({
  index: null,
  visible: false,
  title: '添加字段',
  map: {
    text: ['string', 'varchar', true],
    rich_text: ['longtext'],
    md: ['string', 'longtext'],
    int: ['int', 'int', true],
    float: ['float', 'float', true],
    date: ['time.Time', 'date'],
    datetime: ['time.Time', 'datetime'],
    switch: ['bool', 'TINYINT(1)'],
    color: ['string', 'varchar(10)'],
    img: ['string', 'varchar(255)'],
    file: ['string', 'varchar(512)'],
    many_img: ['string', 'text'],
    many_file: ['string', 'text'],
  },
  form: {
    id: null,
    structName: '',
    jsonName: '',
    sqlName: '',
    columnLabel: '',
    columnType: '',
    inputType: '',
    columnLen: null,
    sqlType: '',
    dictId: null,
    isEdit: true,
    isExport: false,
    isShow: true,
    isQuery: false,
    queryType: '=',
    isSort: null,
    sortType: 'DESC',
    isRequired: false,
  },
  rules: {
    structName: [{ required: true, message: '请输入结构体名', trigger: 'blur' }],
    jsonName: [{ required: true, message: '请输入json名', trigger: 'blur' }],
    sqlName: [{ required: true, message: '请输入sql名', trigger: 'blur' }],
    columnLabel: [{ required: true, message: '请输入Label名', trigger: 'blur' }],
    inputType: [{ required: true, message: '请选择字段类型', trigger: 'blur' }],
  },
  reset: () => {
    fieldForm.form = {
      id: null,
      structName: '',
      jsonName: '',
      sqlName: '',
      columnLabel: '',
      columnType: '',
      inputType: '',
      columnLen: null,
      sqlType: '',
      dictId: null,
      isEdit: true,
      isExport: false,
      isShow: true,
      isQuery: false,
      queryType: '=',
      isSort: false,
      sortType: 'DESC',
      isRequired: false,
    }
  },
  edit: (record: any, index: number) => {
    if (index != undefined) {
      //存在id
      fieldForm.form = record
      fieldForm.index = index
      fieldForm.title = '编辑字段'
    } else {
      //不存在
      fieldForm.index = null
      fieldForm.reset()
      fieldForm.title = '添加字段'
    }
    fieldForm.visible = true
  },
  confirm: () => {
    fieldFormRef.value.validate().then(() => {
      fieldForm.visible = false
      fieldForm.form.sqlType = fieldForm.map[fieldForm.form.inputType][1]
      fieldForm.form.columnType = fieldForm.map[fieldForm.form.inputType][0]

      //判断是否要加上长度
      if (fieldForm.map[fieldForm.form.inputType].length > 2) {
        if (fieldForm.form.columnLen) {
          fieldForm.form.sqlType += '(' + fieldForm.form.columnLen + ')'
        }
      }
      console.log(fieldForm.form)
      if (fieldForm.index != undefined) {
        table.list[fieldForm.index] = fieldForm.form
      } else {
        table.list.push(fieldForm.form)
      }
    })
  },
  convertName: () => {
    const name = fieldForm.form.structName
    //转换为小写
    const allLowerName = name
      .replace(/([A-Z])/g, '_$1')
      .toLowerCase()
      .substring(1)

    fieldForm.form.sqlName = allLowerName
    fieldForm.form.jsonName = name[0].toLowerCase() + name.substring(1)
  },
})
</script>

<template>
  <div>
    <LayoutContainer show-title>
      <template #left>生成表配置</template>
      <template #right>
        <a-space>
          <a-button type="primary">
            <template #icon>
              <SaveOutlined />
            </template>
            保存</a-button
          >
          <a-button type="primary" class="!bg-orange-400">
            <template #icon>
              <IconPackAdapter>
                <IconCode />
              </IconPackAdapter>
            </template>
            生成</a-button
          >
          <a-button type="primary" class="!bg-green-500">
            <template #icon><EyeOutlined /></template>预览</a-button
          >
        </a-space>
      </template>
      <a-form
        ref="tableFormRef"
        :rules="tableForm.rules"
        :model="tableForm.form"
        :label-col="{ style: { width: '100px' } }"
        hide-required-mark
      >
        <a-row>
          <a-col :span="6">
            <a-form-item label="生成版本" name="generateVersion">
              <a-input v-model:value="tableForm.form.generateVersion"></a-input>
            </a-form-item>
          </a-col>
          <a-col :span="6">
            <a-form-item label="生成路径" name="generateBasePath">
              <a-input v-model:value="tableForm.form.generateBasePath"></a-input>
            </a-form-item>
          </a-col>
          <a-col :span="6">
            <a-form-item label="结构体名称" name="structName">
              <a-input v-model:value="tableForm.form.structName"></a-input>
            </a-form-item>
          </a-col>
          <a-col :span="6">
            <a-form-item label="表备注" name="tableComment">
              <a-input v-model:value="tableForm.form.tableComment"></a-input>
            </a-form-item>
          </a-col>
          <a-col :span="6">
            <a-form-item label="关联关系" name="relation">
              <a-select v-model:value="tableForm.form.relation">
                <a-select-option :value="null">无</a-select-option>
                <a-select-option value="many">一对多</a-select-option>
                <a-select-option value="many2many">多对多</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
          <a-col :span="6" v-if="tableForm.form.relation">
            <a-form-item label="关联表" name="generateTableId" :rules="[{ required: tableForm.form.relation, message: '请选择关联关系', trigger: 'blur' }]">
              <a-select v-model:value="tableForm.form.generateTableId"></a-select>
            </a-form-item>
          </a-col>
          <a-col :span="6">
            <a-form-item label="软删除" name="softDelete">
              <a-switch v-model:checked="tableForm.form.softDelete"></a-switch>
            </a-form-item>
          </a-col>
        </a-row>
      </a-form>
    </LayoutContainer>
    <LayoutContainer show-title>
      <template #left>生成表--字段</template>
      <template #right>
        <a-button @click="fieldForm.edit">
          <template #icon>
            <PlusOutlined />
          </template>
          添加字段
        </a-button>
      </template>
      <a-table :columns="table.columns" :data-source="table.list" :scroll="{ x: 2000 }">
        <template #bodyCell="{ column, record, index }">
          <template v-if="column.key.indexOf('is') != -1">
            <a-tag :color="record[column.key] ? '#108ee9' : '#cd201f'">{{
              record[column.key] ? 'yes' : 'no'
            }}</a-tag>
          </template>
          <template v-else-if="column.key == 'operation'">
            <a-space>
              <a-button type="link" class="!text-orange-400" @click="fieldForm.edit(record, index)"
                >修改</a-button
              >
              <a-button type="link" danger @click="table.remove(index)">删除</a-button>
            </a-space>
          </template>
        </template>
      </a-table>
    </LayoutContainer>

    <!--  表单区域  -->
    <a-drawer
      v-model:open="fieldForm.visible"
      :close-icon="null"
      :title="fieldForm.title"
      width="1200"
    >
      <template #extra>
        <a-space>
          <a-button @click="fieldForm.visible = false">取消</a-button>
          <a-button type="primary" @click="fieldForm.confirm">确认</a-button>
        </a-space>
      </template>
      <a-form
        ref="fieldFormRef"
        :label-col="{ style: { width: '100px' } }"
        :model="fieldForm.form"
        :rules="fieldForm.rules"
      >
        <a-row>
          <a-col :span="12" :gutter="20">
            <a-form-item label="结构体名：" name="structName">
              <a-row>
                <a-col :span="19">
                  <a-input v-model:value="fieldForm.form.structName"> </a-input>
                </a-col>
                <a-col :span="4" :offset="1">
                  <a-button @click="fieldForm.convertName">同步</a-button>
                </a-col>
              </a-row>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="json名：" name="jsonName">
              <a-input v-model:value="fieldForm.form.jsonName"></a-input>
            </a-form-item>
          </a-col>
        </a-row>
        <a-row>
          <a-col :span="12">
            <a-form-item label="sql名：" name="sqlName">
              <a-input v-model:value="fieldForm.form.sqlName"></a-input>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="Label名：" name="columnLabel">
              <a-input v-model:value="fieldForm.form.columnLabel"></a-input>
            </a-form-item>
          </a-col>
        </a-row>

        <a-row>
          <a-col :span="12">
            <a-form-item label="控件类型：" name="inputType">
              <a-select v-model:value.number="fieldForm.form.inputType">
                <a-select-option value="text">文本</a-select-option>
                <a-select-option value="rich_text">富文本</a-select-option>
                <a-select-option value="md">MD</a-select-option>
                <a-select-option value="int">整数</a-select-option>
                <a-select-option value="float">小数</a-select-option>
                <a-select-option value="date">日期</a-select-option>
                <a-select-option value="datetime">日期（包含时分秒）</a-select-option>
                <a-select-option value="switch">开关</a-select-option>
                <a-select-option value="color">颜色</a-select-option>
                <a-select-option value="img">图片上传</a-select-option>
                <a-select-option value="file">文件上传</a-select-option>
                <a-select-option value="many_img">图片上传（多）</a-select-option>
                <a-select-option value="many_file">文件上传（多）</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="类型长度">
              <a-input type="number" v-model:value="fieldForm.form.columnLen"></a-input>
            </a-form-item>
          </a-col>
        </a-row>
        <a-row>
          <a-col :span="12">
            <a-form-item label="字典：" name="dictId">
              <a-select v-model:value.number="fieldForm.form.columnType"></a-select>
            </a-form-item>
          </a-col>
        </a-row>

        <a-row>
          <a-col :span="12">
            <a-form-item label="编辑" name="isEdit">
              <a-switch v-model:checked="fieldForm.form.isEdit"></a-switch>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="导入/导出" name="isExport">
              <a-switch v-model:checked="fieldForm.form.isExport"></a-switch>
            </a-form-item>
          </a-col>
        </a-row>
        <a-row>
          <a-col :span="12">
            <a-form-item label="显示" name="isShow">
              <a-switch v-model:checked="fieldForm.form.isShow"></a-switch>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="必填" name="isRequired">
              <a-switch v-model:checked="fieldForm.form.isRequired"></a-switch>
            </a-form-item>
          </a-col>
        </a-row>
        <a-row>
          <a-col :span="12">
            <a-form-item label="查询" name="isQuery">
              <a-switch v-model:checked="fieldForm.form.isQuery"></a-switch>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="查询条件" name="queryType">
              <a-select v-model:value="fieldForm.form.queryType">
                <a-select-option value="LIKE">LIKE</a-select-option>
                <a-select-option value="=">=</a-select-option>
                <a-select-option value=">">&gt;</a-select-option>
                <a-select-option value="<">&lt;</a-select-option>
                <a-select-option value=">=">&gt;=</a-select-option>
                <a-select-option value="<=">&lt;=</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
        </a-row>
        <a-row>
          <a-col :span="12">
            <a-form-item label="排序" name="isOrder">
              <a-switch v-model:checked="fieldForm.form.isSort"></a-switch>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="排序类型" name="orderType">
              <a-select v-model:value="fieldForm.form.sortType">
                <a-select-option value="DESC">DESC</a-select-option>
                <a-select-option value="ASC">ASC</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
        </a-row>
      </a-form>
    </a-drawer>
  </div>
</template>

<style scoped lang="scss"></style>
