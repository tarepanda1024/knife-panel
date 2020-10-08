<template>
  <div id="app">
    <Termianl />
  </div>
</template>

<script>
import Termianl from './components/Terminal/Terminal.vue'

export default {
  name: 'app',
  components: {
    Termianl
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
