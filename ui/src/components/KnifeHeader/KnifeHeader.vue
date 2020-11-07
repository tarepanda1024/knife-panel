<template>
  <Header>
    <Menu mode="horizontal" theme="dark" active-name="" @on-select="handleMenuSelect">
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

    <Modal
        v-model="terminalConf.showModal"
        :fullscreen=terminalConf.fullScreen
        width="800"
        title="终端"
        footer-hide
    >
      <Terminal ref="terminal"/>
    </Modal>
  </Header>
</template>

<script>
import Terminal from "@/components/Terminal/Terminal";

export default {
  name: "KnifeHeader",
  components: {Terminal},
  data() {
    return {
      terminalConf: {
        showModal: false,
        fullScreen: true
      }
    }
  },
  methods: {
    handleMenuSelect(name) {
      switch (name) {
        case 'terminal':
          this.terminalConf.showModal = true;
          let that = this;
          let t;
          clearTimeout(t)
          t = setTimeout(function () {
            //延时执行，确保页面渲染完毕
            that.$refs.terminal.init()
          }, 100)
          break
      }
    }
  }
}
</script>

<style scoped>

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