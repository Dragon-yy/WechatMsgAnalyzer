<template>
  <div>
    <div style="padding-top: 15vh" class="a">
      这一年 我们说过最多的词是

      <span style="font-size: 5vw" class="value"> {{ mostusedword.content }} </span>
      居然有
      <span class="value"> {{ mostusedword.count }}次</span>
    </div>
    <div class="b">
      <VueWordcloud :data="wordCloudData" :options="wordCloudOptions" />
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import VueWordcloud from 'vue-wordcloud';

export default {
  components: {
    VueWordcloud,
  },
  data() {
    return {
      currentTime: 0, // 初始化为0
      data: null, // 用于存储从后端获取的数据
      wordCloudData: [], // 用于存储词云数据
      wordCloudOptions: {
        fontFamily: ['Arial', 'sans-serif'],
        fontWeight: 'bold',
        minFontSize: 20,
        maxFontSize: 80,
        rotationAngles: [-45, 45],
        rotationSteps: 0.1,
        colors: ['#FF0000', '#00FF00', '#0000FF'],
        backgroundColor: '#FFFFFF',
        shape: 'circle',
        padding: 10,
        scale: 1,
        spiral: 'archimedean',
        showTooltip: false,
      }
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
            // 更新当前时间戳（秒）
            this.currentTime = Math.floor(Date.now() / 1000);
            this.wordCloudData = response.data.TopWords.map((word) => {
              // console.log(word["Word"], word["Count"])
              return {
                name: word["Word"],//字符串
                value: word["Count"],
              };
            });
            // console.log(this.wordCloudData)

          })
          .catch(error => {
            // 请求失败，输出错误信息
            console.error('Error fetching data:', error);
          });
    }
  },
  computed: {
    mostusedword() {
      if (!this.data) {
        return {}; // 如果数据还未加载，返回空对象
      }
      let mostusedword = Object.keys(this.data.MostUsedWord)[0]
      let count = this.data.MostUsedWord[mostusedword]
      console.log(mostusedword, count)
      return {
        content: mostusedword,
        count: count,
      };
    }
  }
};
</script>

<style scoped>
.a{
  animation-name: slide-top;
  animation-duration: 1s;
  animation-delay: 0.5s;
  animation-fill-mode: backwards;
}
.b img {
  width: 100%;
  animation-name: slide-top;
  animation-duration: 1s;
  animation-delay: 1s;
  animation-fill-mode: backwards;
}
</style>