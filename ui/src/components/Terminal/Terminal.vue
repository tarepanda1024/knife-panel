<template>
  <div id="terminal">

  </div>
</template>

<script>
import {Xterm} from "./xterm";
import {ConnectionFactory} from "./websocket";
import {protocols, WebTTY} from "./webtty";

export default {
  name: "Terminal",
  methods: {
    init() {
      console.log("begin init")
      const elem = document.getElementById("terminal")
      if (elem != null && this.term == null) {
        //TODO:需要判断链接是否关闭，如果关闭，需要重新连接
        this.term = new Xterm(elem);
        const httpsEnabled = window.location.protocol === "https:";
        const url = (httpsEnabled ? 'wss://' : 'ws://') + window.location.host + '/tty/';
        const args = window.location.search;
        const factory = new ConnectionFactory(url, protocols);
        const wt = new WebTTY(this.term, factory, args);
        const closer = wt.open();
        const tmpTerm = this.term
        window.addEventListener("unload", () => {
          closer();
          tmpTerm.close();
        });
      } else {
        console.log("ele is null")
      }
    }
  }
}
</script>

<style scoped>
#terminal {
  width: 100%;
  height: 100%;
  min-height: 500px;
}
</style>