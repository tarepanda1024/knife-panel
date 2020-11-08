<template>
  <Card style="height: 100%">
    <Tree :data="data" :load-data="handleLoadTables" :render="renderContent" @on-select-change="handleNodeClick">
    </Tree>
  </Card>
</template>

<script>
import {listDB, listDBTables} from "@/service/DatabaseService";

export default {
  name: "DatabaseNav",
  data() {
    return {
      data: []
    }
  },
  methods: {
    renderContent(h, {root, node, data}) {
      let icon;
      if (data.type === 'db') {
        icon = 'server'
      } else {
        icon = 'apps-sharp'
      }
      return h('span', {
        style: {
          display: 'inline-block',
          width: '100%'
        }
      }, [
        h('span', [
          h('ion-icon', {
            attrs: {
              name: icon
            },
            style: {
              marginRight: '8px'
            }
          }),
          h('span', data.title)
        ]),
      ]);
    },

    loadDatabase() {
      listDB().then(data => {
        if (data) {
          for (let i in data) {
            let item = data[i];
            this.data.push({
              title: item.Name,
              loading: false,
              id: item.ID,
              type: 'db',
              children: []
            })
          }
        }
      })
    },
    handleLoadTables(item, callback) {
      listDBTables({id: item.id}).then(data => {
        let tables = []
        for (let i in data) {
          tables.push({
            title: data[i],
            type: 'table',
            pid: item.id
          })
        }
        callback(tables)
      })
    },
    handleNodeClick(select, data) {
      if (data.type === 'table') {
        this.$emit('onTreeNodeChanged', data.pid, data.title)
      } else {
        this.$emit('onTreeNodeChanged', -1, data.title)
      }
    }
  },
  mounted() {
    this.loadDatabase()
  }
}
</script>

<style scoped>

</style>