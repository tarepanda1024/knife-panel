<template>
  <div >
    <Row>
      <Col span="6">
        <Card >
            <span slot="title">
              基本信息
            </span>
        </Card>
      </Col>
      <Col span="18">
        <Card style="margin-left: 16px">
            <span slot="title">
              资源概览
            </span>
        </Card>
      </Col>
    </Row>
  </div>
</template>

<script>
import {overview} from "@/service/DockerService";

export default {
  name: "DockerOverview",
  data(){
    return {
      data:{}
    }
  },
  methods:{
    loadData(){
      overview().then(data => {
        this.data = data;
      })
    }
  }
}
</script>

<style scoped>

</style>