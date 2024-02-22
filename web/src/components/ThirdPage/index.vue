<template>
  <div>
    <div class="a">
      今年我们一共发了
      <span class="value"> {{ count.totalCount }} </span>
      条消息
    </div>
    <div class="b">
      其中有
      <span class="value"> {{ count.imgCount }} </span>
      是图片&nbsp; 其中有
      <span class="value"> {{ count.emojiCount }} </span>
      是表情包
    </div>
    <div class="c">
      <div v-for="(count, word, index) in count.loveWord" :key="index">
        说{{ word }} <span class="value"> {{ count }} </span>次
      </div>
      <div>...</div>
    </div>
    <div class="d">比起网络<br />我们更喜欢在现实表达爱意</div>
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
    count() {
      if (!this.data) {
        return {}; // 如果数据还未加载，返回空对象
      }
      // 遍历
     return {
        totalCount:this.data.SpeakCount,
        imgCount:this.data.ImgCount,
        emojiCount:this.data.EmojiCount,
        loveWord: this.data.LoveCount

     }

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
  animation-delay: .5s;
  animation-fill-mode: backwards;
}
.b{
  animation-name: slide-top;
  animation-duration: 1s;
  animation-delay: 1s;
  animation-fill-mode: backwards;
}
.c{
  animation-name: slide-top;
  animation-duration: 1s;
  animation-delay: 2s;
  animation-fill-mode: backwards;
}
.d {
  padding-top: 10vh;
  font-size: 5vw;
  animation-name: slide-top;
  animation-duration: 1s;
  animation-delay: 2.5s;
  animation-fill-mode: backwards;
}

</style>