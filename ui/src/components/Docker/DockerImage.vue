<template>
  <div>
    <Table :columns="columns" :data="data"></Table>
  </div>
</template>

<script>
import {listImages} from "@/service/DockerService";
import {formatSize} from "@/utils/FileSizeUtils";

export default {
  name: "DockerImage",
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
          key: 'Names',
          align: 'center',
          "render": (h, params) => {
            let repo = params.row.RepoTags[0]
            return h('span', repo.split(":")[0])
          }
        },
        {
          title: '标签',
          key: 'State',
          align: 'center',
          "render": (h, params) => {
            let repo = params.row.RepoTags[0]
            return h('span', repo.split(":")[1])
          }
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
          title: '文件大小',
          key: 'Size',
          align: 'center',
          "render": (h, params) => {
            return h('span', formatSize( params.row.Size))
          }
        }
      ],
      data: []
    }
  },
  methods: {
    handleListImages() {
      listImages().then(data => {
        this.data = data;
      })
    }
  }
}
</script>

<style scoped>

</style>