<script>
import { Toast } from "vant";
import { getRequest } from "../../tools/axios";
import USShow from "../subViews/USShow.vue";
import MoviePage from "../subViews/MoviePage.vue";
import TV from "../subViews/TV.vue";
import AnimePage from "../subViews/AnimePage.vue";
import userStorage from "../../tools/storage";

export default {
  name: "UserLogin",
  data() {
    return {
      email: "",
      password: "",
      captchaImageSrc: "",
      captchaID: "",
      captchaAnswer: "",
      message: "",
      error: "",
      active: 0,
    };
  },
  methods: {
    goToLogin() {
      this.$router.push("/login");
    },
    goToSetting() {
      this.$router.push("/setting");
    },
    getHomePageInfo() {
      getRequest("/user/home", {}, true)
        .then((data) => {
          console.log("data:", data); // Debug log
          this.getAllList();
        })
        .catch((error) => {
          console.log(error);

          if (error.status == 401) {
            this.goToLogin();
            Toast.fail("Request need login" + error.status);
          } else {
            Toast.fail("Request fail " + error.status);
          }
        });
    },
    getAllList() {
      getRequest("/v1/products", {}, false)
        .then((response) => {
          console.log(response.data);
        })
        .catch((error) => {
          console.log(error.message);
        });
    },
    onChange(index) {
      // Update 'active' when tab is switched
      this.active = index;
      console.log(`Switched to tab ${index}`);
    },
  },
  components: {
    MoviePage,
    TV,
    USShow,
    AnimePage,
  },
  created() {
    if (userStorage.getItem("isAdmin") == 3) {
      return;
    }
    this.getHomePageInfo();
  },
  onMounted() {
    if (userStorage.getItem("isAdmin") == 3) {
      return;
    }
    this.getHomePageInfo();
  },
};
</script>

<template>
  <div class="tabs-container">
    <van-tabs :active="active" @change="onChange" class="tabs">
      <van-tab title="电影">
        <div class="tab-content">
          <MoviePage v-if="active === 0"></MoviePage>
        </div>
      </van-tab>
      <van-tab title="电视">
        <div class="tab-content">
          <TV v-if="active === 1"></TV>
        </div>
      </van-tab>
      <van-tab title="美剧">
        <div class="tab-content">
          <USShow v-if="active === 2"></USShow>
        </div>
      </van-tab>
      <van-tab title="动漫">
        <div class="tab-content">
          <AnimePage v-if="active === 3"></AnimePage>
        </div>
      </van-tab>
    </van-tabs>
  </div>
</template>

<style scoped>
.title-info {
  display: flex;
  flex-direction: column;
  text-align: center;
  color: black;
  font-size: 2vw;
}

.tabs-container {
  height: 100vh; /* 设置容器高度 */
  overflow: hidden; /* 隐藏超出部分 */
}

.tabs {
  height: auto; /* tabs 自身不滑动 */
}

.tab-content {
  height: calc(100vh - 48px); /* 计算内容区高度，48px 为 tab 的高度 */
  overflow-y: auto; /* 允许垂直滚动 */
  padding: 10px; /* 可选，添加内边距 */
}
</style>
