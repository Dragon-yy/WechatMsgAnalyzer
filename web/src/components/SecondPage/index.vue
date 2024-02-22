<template>
  <div>
    <div class="a">
      我们在一起了
      <span class="value">{{ timeDifference.days }}</span> 天
      <span class="value">{{ timeDifference.hours }}</span> 小时
      <span class="value">{{ timeDifference.minutes }}</span> 分钟
      <span class="value">{{ timeDifference.seconds }}</span> 秒
    </div>

    <div class="b">年年有我，岁岁有我</div>
    <div class="c">出发去看看我们的2023吧~</div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      currentTime: 0, // 初始化为0
      data: null // 用于存储从后端获取的数据
    };
  },
  created() {
    // 在组件创建时发送HTTP请求获取后端数据
    this.fetchData();
  },
  methods: {
    fetchData() {
      // 使用axios发送GET请求获取后端数据
      axios.get("http://172.20.10.3:9000/report") //
          .then(response => {
            // 请求成功，将后端返回的数据保存到data属性中
            this.data = response.data;
            console.log(this.data)
            // 更新当前时间戳（秒）
            this.currentTime = Math.floor(Date.now() / 1000);
          })
          .catch(error => {
            // 请求失败，输出错误信息
            console.error('Error fetching data:', error);
          });
    }
  },
  computed: {
    timeDifference() {
      if (!this.data) {
        return {}; // 如果数据还未加载，返回空对象
      }

      // 解析接收到的 JSON 数据
      const timestampDifference = this.currentTime - this.data.FirstTime;
      const day = 24 * 60 * 60; // 一天的秒数
      const hour = 60 * 60; // 一小时的秒数
      const minute = 60; // 一分钟的秒数

      // 计算天、小时、分钟和秒
      const days = Math.floor(timestampDifference / day);
      const hours = Math.floor((timestampDifference % day) / hour);
      const minutes = Math.floor((timestampDifference % hour) / minute);
      const seconds = timestampDifference % minute;

      return {
        days: days,
        hours: hours,
        minutes: minutes,
        seconds: seconds
      };
    }
  }
};
</script>
<style scoped>
.a,
.b,
.c {
  text-align: center;
  color: white;
}
.a {
  font-size: 8vw;
  padding-top: 15vh;
  animation-name: slide-top;
  animation-duration: 1s;
  animation-delay: .5s;
  animation-fill-mode: backwards;
}
.b {
  font-size: 6vw;
  padding-top: 15vh;
  animation-name: slide-top;
  animation-duration: 1s;
  animation-delay: 1s;
  animation-fill-mode: backwards;
}
.c {
  font-size: 5vw;
  font-weight: lighter;
  animation-name: slide-top;
  animation-duration: 1s;
  animation-delay: 1.5s;
  animation-fill-mode: backwards;
}
</style>