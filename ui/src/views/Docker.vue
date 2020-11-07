<template>
  <Tabs value="overview" @on-click="handleTableChanged">
    <TabPane label="总览" name="overview">
      <DockerOverview ref="dockerOverview"/>
    </TabPane>
    <TabPane label="容器" name="container">
      <DockerContainer ref="dockerContainer"/>
    </TabPane>
    <TabPane label="镜像" name="image">
      <DockerImage ref="dockerImage"/>
    </TabPane>
    <TabPane label="网络" name="network">
      <DockerNetwork ref="dockerNetwork"/>
    </TabPane>
    <TabPane label="存储" name="volume">
      <DockerVolume ref="dockerVolume"/>
    </TabPane>
  </Tabs>
</template>

<script>
import DockerContainer from "@/components/Docker/DockerContainer";
import DockerImage from "@/components/Docker/DockerImage";
import DockerNetwork from "@/components/Docker/DockerNetwork";
import DockerVolume from "@/components/Docker/DockerVolume";
import DockerOverview from "@/components/Docker/DockerOverview";

export default {
  name: "Docker",
  components: {DockerOverview, DockerVolume, DockerNetwork, DockerImage, DockerContainer},
  methods: {
    handleTableChanged(name) {
      switch (name) {
        case 'overview':
          this.$refs.dockerOverview.loadData();
          break
        case 'container':
          this.$refs.dockerContainer.handleListContainers()
          break
        case 'image':
          this.$refs.dockerImage.handleListImages()
          break
        case 'network':
          this.$refs.dockerNetwork.handleListNets()
          break
        case 'volume':
          this.$refs.dockerVolume.handleListVolumes()
          break
      }
    }
  }
}
</script>

<style scoped>

</style>