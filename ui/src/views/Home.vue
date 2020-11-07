<template>
  <Layout class="layout">
    <KnifeHeader/>
    <Layout>
      <!--侧边栏-->
      <Sider hide-trigger :style="{background: '#fff'}">
        <Menu theme="light" width="auto">
          <template v-for="item in homeRoutes">
            <Submenu v-if="item.children && item.children.length > 0" :name="item.name">
              <template slot="title">
                <ion-icon v-if="item.meta.icon" :name="item.meta.icon"></ion-icon>
                {{ item.meta.title }}
              </template>
              <template v-for="childItem in item.children">
                <MenuItem :name="childItem.name" :to="{'name':childItem.name}">
                  {{ childItem.meta.title }}
                </MenuItem>
              </template>
            </Submenu>
            <MenuItem v-else :name="item.name" :to="{'name':item.name}">
              <ion-icon v-if="item.meta.icon" :name="item.meta.icon"></ion-icon>
              {{ item.meta.title }}
            </MenuItem>
          </template>
        </Menu>
      </Sider>
      <!--内容区-->
      <Layout :style="{padding: '0 24px 24px'}">
        <Content :style="{padding: '24px', height: '100%', background: '#fff'}">
          <router-view/>
        </Content>
      </Layout>
    </Layout>
  </Layout>

</template>

<script>
import KnifeHeader from "@/components/KnifeHeader/KnifeHeader.vue";

export default {
  name: "Home",
  components: {
    KnifeHeader,

  },
  data() {
    return {
      homeRoutes: []
    }
  },
  mounted() {
    for (let index in this.$router.options.routes) {
      let routeItem = this.$router.options.routes[index];
      if (routeItem.name === 'Home') {
        this.homeRoutes = routeItem.children;
        break;
      }
    }
    console.log(this.homeRoutes)
  }
}
</script>
<style scoped>
.layout {
  width: 100%;
  height: 100%;
  border: 1px solid #d7dde4;
  background: #f5f7f9;
  position: relative;
  border-radius: 4px;
  overflow: hidden;
}

</style>