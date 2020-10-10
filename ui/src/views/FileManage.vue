<template>
  <div>
    <Table :columns="fileColumns" :data="files"></Table>
  </div>
</template>

<script>
export default {
  name: "FileManage",
  data() {
    return {
      fileColumns: [
        {
          "title": "文件名",
          "key": "name"
        },
        {
          "title": "大小",
          "key": "size"
        },
        {
          "title": "修改时间",
          "key": "modifyTime"
        }
      ],
      files: []
    }
  }
}
</script>

<style scoped>

</style>