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
      const elem = document.getElementById("terminal")

      if (elem !== null) {
        //let term: Terminal;
        this.term = new Xterm(elem);
        const httpsEnabled = window.location.protocol === "https:";
        const url = (httpsEnabled ? 'wss://' : 'ws://') + window.location.host + '/tty/';
        const args = window.location.search;
        console.log("url:" + url)
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
  },
  mounted() {
    this.init();
  }
}
</script>

<style scoped>
#terminal {
  width: 100%;
  height: 100%;
}
</style>