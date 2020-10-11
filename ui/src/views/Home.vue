<template>
  <Layout class="layout">
    <Header>
      <Menu mode="horizontal" theme="dark" active-name="">
        <div class="layout-logo">小刀面板</div>
        <div class="layout-nav">
          <MenuItem name="terminal">
            <ion-icon name="terminal" class="ivu-icon"></ion-icon>
            终端
          </MenuItem>
          <MenuItem name="person">
            <ion-icon name="person" class="ivu-icon"></ion-icon>
            管理员
          </MenuItem>
        </div>
      </Menu>
    </Header>
    <Layout>
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
      <Layout :style="{padding: '0 24px 24px'}">
        <Content :style="{padding: '24px', minHeight: '280px', background: '#fff'}">
          <router-view/>
        </Content>
      </Layout>
    </Layout>
  </Layout>

</template>

<script>
export default {
  name: "Home",
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

.layout-logo {
  width: 100px;
  height: 30px;
  border-radius: 3px;
  float: left;
  position: relative;
  top: 15px;
  left: 20px;
  color: white;
  text-align: center;
  line-height: 30px;
}

.layout-nav {
  width: 200px;
  margin: 0 20px 0 auto;
}

</style>