<template>
  <Row style="height: 100%">
    <Col span="4" style="height: 100%">
      <DatabaseNav @onTreeNodeChanged="handleTreeNodeChanged"/>
    </Col>
    <Col span="20" style="height: 100%">
      <DatabaseContent :id="dbId" :name="dbname" />
    </Col>
  </Row>
</template>

<script>
import DatabaseNav from "@/components/Database/DatabaseNav";
import DatabaseContent from "@/components/Database/DatabaseContent";

export default {
  name: "Database",
  components: {DatabaseContent, DatabaseNav},
  data() {
    return {
      dbId: 0,
      dbname: ''
    }
  },
  methods: {
    handleTreeNodeChanged(id, tblname) {
      this.dbId = id;
      this.dbname = tblname;
      console.log(this.dbId + " " + this.dbname)
    }
  }
}
</script>

<style scoped>

</style>