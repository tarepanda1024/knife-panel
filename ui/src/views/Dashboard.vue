<template>
  <div style="display: flex">
    <div id="cpuChart" style="width: 300px;height:300px;">
    </div>
    <div id="memChart" style="width: 300px;height:300px;">
    </div>
  </div>
</template>

<script>
import echartsLiquidfill from 'echarts-liquidfill'
export default {
  name: "Dashboard",
  data() {
    return {}
  },
  methods: {
    loadLoadChart() {
      const cpuChart = this.$echarts.init(document.getElementById('cpuChart'));
      const option = {
        tooltip: {
          formatter: '{a} <br/>{b} : {c}%'
        },
        series: [
          {
            name: 'cpu-usage',
            type: 'gauge',
            detail: {formatter: '{value}%'},
            data: [{value: 20, name: 'CPU使用率'}]
          },
        ]
      };
      cpuChart.setOption(option)
    },
    loadMemChart() {
      const memChart = this.$echarts.init(document.getElementById('memChart'));
      let option = {
        series: [{
          type: 'liquidFill',
          name: '内存使用率',
          radius: '65%',
          data: [0.3],
          label: {
            formatter: '{a}\n{c}%',
            fontSize: 18
          },
          center: ['50%', '45%'],
          outline: {
            show: false
          }
        }]
      };
      memChart.setOption(option)
    }
  },
  mounted() {
    this.loadLoadChart();
    this.loadMemChart()
  }
}
</script>

<style scoped>

</style>