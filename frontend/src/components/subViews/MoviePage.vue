<template>
  <div class="main-container">
    <div class="movie-list">
      <div
        v-for="movie in movies"
        :key="movie.ID"
        class="movie-card"
        @click="openPlayer(movie)"
      >
        <img :src="movie.cover" alt="poster" class="movie-poster" />
        <span class="movie-title">{{ movie.title }}</span>
      </div>
    </div>
  </div>
</template>

<script>
import { getRequest } from "../../tools/axios";

export default {
  name: "MoviePage",
  data() {
    return {
      movies: [],
    };
  },
  methods: {
    openPlayer(movie) {
      console.log("open player for movie: ", movie);
      this.$router.push({
        name: "player",
        query: { movie: JSON.stringify(movie) },
      });
    },
    getAllVideosList() {
      getRequest("/v2/getHomeAll", { page: 1, size: 20 }, false)
        .then((response) => {
          console.log("API response data:", response.videos); // 检查返回的数据
          if (Array.isArray(response.videos)) {
            this.movies = response.videos; // 确保 movies 被赋值
          } else {
            console.log("Unexpected data structure:", response.videos);
          }
        })
        .catch((error) => {
          console.log("Error fetching data:", error.message);
        });
    },
  },
  created() {
    console.log("MoviePage component is rendered");
    this.getAllVideosList();
  },
};
</script>

<style>
.main-container {
  width: 100vw;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

.movie-list {
  display: flex;
  flex-wrap: wrap;
  justify-content: flex-start;
  width: 100%; /* 减去间隔 */
  height: 100%;
}

.movie-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: calc((100% / 3) - 10px); /* 三列，减去间隔 */
  height: (calc((100% / 3) * (16 / 9)) + 20px); /* 计算 9:16 比例 */
  margin-bottom: 10px; /* 纵向间隔 */
  background-color: white;
  margin-left: 5px;
}

.movie-poster {
  width: 100%; /* 使图片填满卡片 */
  height: auto; /* 保持比例 */
}

.movie-title {
  font-size: 14px;
  width: 100%;
  text-align: center;
  margin-top: 5px;
}
</style>
