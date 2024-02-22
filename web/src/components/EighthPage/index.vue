<template>
  <div>
    <div class="a">
      这一年我们最喜欢在
      <span class="value"> {{ maxMonth[0] }} </span>点聊天 说的话 有<span
        class="value"
      >
        {{ maxMonth[1] }} </span
      >条
    </div>
    <div ref="chart" class="chart"></div>
  </div>
</template>
<script>
import * as echarts from "echarts";
import axios from 'axios'; // Import Axios

export default {
  data() {
    return {
      myChart: undefined,
      arr: [], // Initialize arr as an empty array
    };
  },
  computed: {
    maxMonth() {
      let max = this.arr[0];
      for (let i of this.arr) {
        if (i[1] > max[1]) {
          max = i;
        }
      }
      console.log("max:"+max)
      return max;
    },
  },
  async mounted() {
    // Fetch data using Axios from the backend
    try {
      const response = await axios.get('http://172.20.10.3:9000/report'); // Replace 'backend-url' with the actual URL
      let activeTime = response.data.ActiveTime; // Assuming the JSON object has a key 'monthGroup'
      // console.log(activeTime)
      // console.log(typeof activeTime)
      // {"month_2023-01":123} 转换为 [[1, 123], ...]
      let result = [];
      for (const [key, value] of Object.entries(activeTime)) {
        if (key.includes("hour")) {
          result.push([parseInt(key.split('_')[1]), value]);
        }
      }
      this.arr = result;
      console.log("arr"+this.arr)
      // Use $nextTick to ensure DOM update
      this.$nextTick(() => {
        // Initialize ECharts after fetching data
        this.myChart = echarts.init(this.$refs.chart);
        this.myChart.setOption({
          tooltip: {},
          xAxis: {},
          yAxis: {
            type: "value",
          },
          series: [
            {
              type: "bar",
              data: this.arr,
              itemStyle: {
                color: "#ebf45f",
              },
              clip :false
            },
          ],
        });
      });
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  },
};
</script>

<style scoped>
.chart {
  width: 110vw;
  height: 60vh;
  animation-name: slide-top;
  animation-duration: 1s;
  animation-delay: 1s;
  animation-fill-mode: backwards;
}
.a {
  padding-top: 15vh;
  animation-name: slide-top;
  animation-duration: 1s;
  animation-delay: 0.5s;
  animation-fill-mode: backwards;
}
</style>