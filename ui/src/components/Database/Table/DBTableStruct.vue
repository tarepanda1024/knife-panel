<template>
  <div>
    <Table :columns="columns" :data="data"></Table>
  </div>
</template>

<script>

import {listTableColumns} from "@/service/DatabaseService";

export default {
  name: "DBTableStruct",
  data() {
    return {
      columns: [
        {
          title: '名称',
          key: 'Name',
          align: 'center'
        },
        {
          title: '字段类型',
          key: 'Type',
          align: 'center',
        },
        {
          title: '允许为空',
          key: 'NotNull',
          align: 'center'
        },
        {
          title: '默认值',
          key: 'DefaultValue',
          align: 'center'
        },
        {
          title: '主键',
          key: 'PK',
          align: 'center'
        },
      ],
      data: []
    }
  },
  methods: {
    handleListColumns(id, name) {
      let params = {
        id: id,
        tblName: name
      }
      listTableColumns(params).then(data => {
        this.data = data;
      })
    }
  }
}
</script>

<style scoped>

</style>