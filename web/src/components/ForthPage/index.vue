<template>
  <div>
    <div class="a">
      这一年 我们说过最长的一句话 <span class="value">{{ content.longmsg.length }}</span
      >字
    </div>
    <div class="b">
      {{ content.longmsg.substring(0, 100) }}
    </div>
    <div class="c">...</div>
    <div class="d">
      {{ content.time }}
    </div>
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
    content() {
      if (!this.data) {
        return {}; // 如果数据还未加载，返回空对象
      }
// 创建一个新的 Date 对象，并将时间戳传递给它
      const date = new Date(this.data.LongestMsg.CreateTime * 1000);

// 从 Date 对象中获取月、日、小时、分钟和秒
      const month = date.getMonth() + 1; // 月份从 0 开始，因此要加上 1
      const day = date.getDay();
      const hours = date.getHours();
      const minutes = date.getMinutes();
      const seconds = date.getSeconds();

// 格式化月份和日期，如果小于 10，则在前面加上 '0'
      const formattedMonth = month < 10 ? "0" + month : month;
      const formattedDay = day < 10 ? "0" + day : day;

// 将时间转换为字符串格式
      const formattedTime = formattedMonth + "-" + formattedDay + " " + hours + ":" + minutes + ":" + seconds;

      return {
        longmsg: this.data.LongestMsg.StrContent,
        time: formattedTime
      };
    }
  }
};
</script>

<style scoped>
.a {
  font-size: 5vw;
  padding-top: 15vh;
  animation-name: slide-top;
  animation-duration: 1s;
  animation-delay: 0.5s;
  animation-fill-mode: backwards;
}
.b {
  padding-top: 5vh;
  white-space: normal;
  text-overflow: ellipsis;
  overflow: hidden;
  line-height: 2;
  color: ebf45f;
  animation-name: slide-top;
  animation-duration: 1s;
  animation-delay: 1s;
  animation-fill-mode: backwards;
}
.c{
  animation-name: slide-top;
  animation-duration: 1s;
  animation-delay: 1s;
  animation-fill-mode: backwards;
}
.d {
  text-align: right;
  padding-top: 3vh;
  font-weight: bolder;
  animation-name: slide-top;
  animation-duration: 1s;
  animation-delay: 1.5s;
  animation-fill-mode: backwards;
}
</style>