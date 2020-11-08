<template>
  <Card style="margin-left: 16px;height: 100%">
    <Tabs value="DataQuery" @on-click="handleTabChanged">
      <TabPane label="数据查询" name="DataQuery">
        <DBDataQuery ref="dataQuery"/>
      </TabPane>
      <TabPane label="表结构" name="TableStruct">
        <DBTableStruct ref="tableStruct"/>
      </TabPane>
      <TabPane label="索引管理" name="TableIndex">
        <DBIndex ref="dbIndex"/>
      </TabPane>
    </Tabs>
  </Card>
</template>

<script>
import DBDataQuery from "@/components/Database/Table/DBDataQuery";
import DBTableStruct from "@/components/Database/Table/DBTableStruct";
import DBIndex from "@/components/Database/Table/DBIndex";

export default {
  name: "DatabaseContent",
  props: {
    id: 0,
    name: ''
  },
  components: {DBIndex, DBTableStruct, DBDataQuery},
  methods: {
    handleTabChanged(name) {
      if (this.id < 1 || this.name === '') {
        return
      }
      switch (name) {
        case 'DataQuery':
          break
        case 'TableStruct':
          this.$refs.tableStruct.handleListColumns(this.id, this.name)
          break
        case 'TableIndex':
          break
      }
    }
  },
  watch: {
    name: function (val) {
      if (this.id > 0) {
        this.$refs.tableStruct.handleListColumns(this.id, this.name)
      }
    }
  }
}
</script>

<style scoped>

</style>