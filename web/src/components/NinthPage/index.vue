<template>
  <div class="v">
    <div class="a">
      <div class="b">#2023</div>
      <div class="c">年度聊天报告</div>
      <div class="d">关键词</div>
      <div class="e">{{ wordcount.content }}</div>
      <div class="f"></div>
      <div class="d">爱意词</div>
      <div class="e">{{ wordcount.loveword }}</div>
      <div class="f"></div>
      <div class="d">话痨月</div>
      <div class="e" style="font-size: 5vw">{{ wordcount.maxMonth }}月</div>
      <div class="f"></div>
      <div class="h">
        <img class="p" src="@/assets/images/result.png" />
        <img class="y" src="@/assets/images/cloud2.png" />
      </div>
      <div class="j">再见</div>
      <div class="i">2023</div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      currentTime: 0, // 初始化为0
      data: null, // 用于存储从后端获取的数据
      activeTime: null
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
            this.activeTime = response.data.ActiveTime; // Assuming the JSON object has a key 'monthGroup'

            // 更新当前时间戳（秒）
            console.log(this.data)
          })
          .catch(error => {
            // 请求失败，输出错误信息
            console.error('Error fetching data:', error);
          });
    }
  },
  computed: {
    wordcount() {
      if (!this.data) {
        return {}; // 如果数据还未加载，返回空对象
      }
      let mostusedword = Object.keys(this.data.MostUsedWord)[0]
      let count = this.data.MostUsedWord[mostusedword]
      console.log(mostusedword, count)

      count = 0
      let mostusedloveword = ""
      // 在LoveCount中找到使用最多的词
      for (let key in this.data.LoveCount) {
        if (this.data.LoveCount[key] > count) {
          mostusedloveword = key
          count = this.data.LoveCount[key]
        }
      }
      console.log(mostusedloveword, count)

      // 寻找交流最多的月份
      // console.log(this.activeTime)
      // console.log(typeof activeTime)
      // {"month_2023-01":123} 转换为 [[1, 123], ...]
      let result = [];
      for (const [key, value] of Object.entries(this.activeTime)) {
        if (key.includes("month")) {
          result.push([parseInt(key.split('-')[1]), value]);
        }
      }
      console.log(result)
      // this.arr = result;
      let max = result[0];
      for (let i of result) {
        if (i[1] > max[1]) {
          max = i;
        }
      }
      // console.log("max:"+max)

      return {
        content: mostusedword,
        loveword: mostusedloveword,
        maxMonth: max[0]
      };
    }
  }
};
</script>

<style scoped>
.v {
  background-image: linear-gradient(#f89d92, #fac9bf);
  height: 100%;
  color: #73e373;
}
.a {
  margin-top: 10vh;
  height: 60vh;
  border: 5px #73e373 solid;
  background: #fef9f9;
  padding: 5vw;
}
.b {
  font-size: 13vw;
  padding-bottom: unset;
}
.c {
  font-size: 5vw;
  font-weight: 700;
  padding-bottom: 5vh;
}
.d {
  color: black;
  font-weight: bold;
}
.e {
  color: black;
  /* font-weight: bold; */
  font-size: 12vw;
}
.f {
  height: 1px;
  background: #73e373;
  width: 30vw;
  margin-top: 2vh;
  margin-bottom: 2vh;
}
.p {
  position: absolute;
  width: 27vw;
  left: 55vw;
  top: 40vh;
  z-index: 2;
}
.y {
  position: absolute;
  width: 60vw;
  left: 36vw;
  top: 50vh;
}
.i,
.j {
  position: absolute;
  color: #3c4739;
}
.j {
  font-size: 15vw;
  top: 25vh;
  left: 50vw;
}
.i {
  font-size: 15vw;
  top: 32vh;
  left: 53vw;
}
</style>