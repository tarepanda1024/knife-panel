<template>
  <div>
    <Table :columns="columns" :data="data"></Table>
  </div>
</template>

<script>
import {listNets} from "@/service/DockerService";
import {formatSize} from "@/utils/FileSizeUtils";

export default {
  name: "DockerNetwork",
  data() {
    return {
      columns: [
        {
          title: 'ID',
          key: 'Id',
          align: 'center',
          "render": (h, params) => {
            let id = params.row.Id.substr(7,12)
            return h('span', id)
          }
        },
        {
          title: '名称',
          key: 'Name',
          align: 'center',
        },
        {
          title: '驱动',
          key: 'Driver',
          align: 'center',
        },
        {
          title: 'Attachable',
          key: 'Attachable',
          align: 'center',
        },
        {
          title: 'IPAM驱动',
          key: 'IPAM',
          align: 'center',
          "render": (h, params) => {
            let driver = params.row.IPAM.Driver
            return h('span', driver)
          }
        },
        {
          title: '创建时间',
          key: 'Created',
          align: 'center',
          "render": (h, params) => {
            let time = params.row.Created;
            let value = this.$moment.utc(time).local().format('YYYY-MM-DD HH:mm:ss')
            return h('span', value + ' ')
          }
        }
      ],
      data: []
    }
  },
  methods: {
    handleListNets() {
      listNets().then(data => {
        this.data = data;
      })
    }
  }
}
</script>

<style scoped>

</style>