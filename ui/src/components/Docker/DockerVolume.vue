<template>
  <div>
    <Table :columns="columns" :data="data"></Table>
  </div>
</template>

<script>
import { listVolumes} from "@/service/DockerService";

export default {
  name: "DockerVolume",
  data() {
    return {
      columns: [
        {
          title: '名称',
          key: 'Name',
          align: 'center'
        },
        {
          title: '驱动',
          key: 'Driver',
          align: 'center'
        },
        {
          title: '挂载点',
          key: 'Mountpoint',
          align: 'center'
        },
      ],
      data: []
    }
  },
  methods: {
    handleListVolumes() {
      listVolumes().then(data => {
        this.data = data.Volumes;
      })
    }
  }
}
</script>

<style scoped>

</style>