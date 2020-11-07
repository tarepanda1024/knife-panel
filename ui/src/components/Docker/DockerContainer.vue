<template>
  <div>
    <Table :columns="columns" :data="data"></Table>
  </div>
</template>

<script>
import {listContainers} from "@/service/DockerService";

export default {
  name: "DockerContainer",
  data() {
    return {
      columns: [
        {
          title: 'ID',
          key: 'Id',
          align: 'center',
          "render": (h, params) => {
            let id = params.row.Id.substr(0,12)
            return h('span', id)
          }
        },
        {
          title: '名称',
          key: 'Names',
          align: 'center',
          "render": (h, params) => {
            let name = params.row.Names[0].substring(1)
            return h('span', name)
          }
        },
        {
          title: '状态',
          key: 'State',
          align: 'center',
        },
        {
          title: '镜像',
          key: 'Image',
          align: 'center',
        },
        {
          title: '创建时间',
          key: 'Created',
          align: 'center',
          "render": (h, params) => {
            let time = params.row.Created;
            let value = this.$moment.unix(time).format('YYYY-MM-DD HH:mm:ss')
            return h('span', value + ' ')
          }
        },
        {
          title: '启动时间',
          key: 'Status',
          align: 'center',
        }
      ],
      data: []
    }
  },
  methods: {
    handleListContainers() {
      listContainers().then(data => {
        this.data = data;
      })
    }
  }
}
</script>

<style scoped>

</style>