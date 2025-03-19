<template>
  <van-nav-bar
    :title="movie.title"
    left-text=""
    right-text=""
    left-arrow
    @click-left="onClickLeft"
    @click-right="onClickRight"
  />
  <div :class="['video-container', { fullscreen: isFullscreen }]">
    <video ref="videoPlayer" class="video-player" controls></video>
    <!-- <div class="controls" v-if="showControls">
      <button @click="togglePlay">{{ isPlaying ? "Pause" : "Play" }}</button>
      <button @click="fastForward">Fast Forward</button>
      <button @click="toggleFullscreen">
        {{ isFullscreen ? "Exit Fullscreen" : "Fullscreen" }}
      </button>
    </div> -->
  </div>
</template>

<script>
import Hls from "hls.js";

export default {
  name: "PlayerPage",
  data() {
    return {
      videoSrc: "", // 直接在这里定义 M3U8 地址
      isPlaying: false,
      isFullscreen: false,
      showControls: true,
      currentTime: 0,
      duration: 0,
      movie: {},
    };
  },
  created() {
    this.movie = JSON.parse(this.$route.query.movie);
    const videos = this.movie.address.split("#");
    this.videoSrc = videos[0];
    console.log("Movie Details:", videos);
    console.log("Movie Details:", this.videoSrc);
  },
  mounted() {
    const video = this.$refs.videoPlayer;
    if (Hls.isSupported()) {
      const hls = new Hls();
      hls.loadSource(this.videoSrc);
      hls.attachMedia(video);
      hls.on(Hls.Events.MANIFEST_PARSED, () => {
        video
          .play()
          .then(() => {
            this.isPlaying = true;
          })
          .catch((error) => {
            console.error("自动播放失败:", error);
          });
      });
    } else if (video.canPlayType("application/vnd.apple.mpegurl")) {
      // 对于 Safari 支持的情况
      video.src = this.videoSrc;
      video.addEventListener("loadedmetadata", () => {
        video
          .play()
          .then(() => {
            this.isPlaying = true;
          })
          .catch((error) => {
            console.error("自动播放失败:", error);
          });
      });
    }

    video.addEventListener("mouseenter", () => {
      this.showControls = true;
    });
    video.addEventListener("mouseleave", () => {
      this.showControls = false;
    });
  },
  methods: {
    onClickLeft() {
      console.log("Movie Details:", this.movie);
      this.$router.go(-1);
    },
    onClickRight() {
      console.log("Movie Details:", this.movie);
    },
    togglePlay() {
      const video = this.$refs.videoPlayer;
      if (video.paused) {
        video.play();
        this.isPlaying = true;
      } else {
        video.pause();
        this.isPlaying = false;
      }
    },
    fastForward() {
      const video = this.$refs.videoPlayer;
      video.currentTime += 10; // 快进10秒
    },
    toggleFullscreen() {
      const videoContainer = this.$el;
      if (!this.isFullscreen) {
        if (videoContainer.requestFullscreen) {
          videoContainer.requestFullscreen();
        } else if (videoContainer.webkitRequestFullscreen) {
          videoContainer.webkitRequestFullscreen();
        } else if (videoContainer.msRequestFullscreen) {
          videoContainer.msRequestFullscreen();
        }
      } else {
        if (document.exitFullscreen) {
          document.exitFullscreen();
        } else if (document.webkitExitFullscreen) {
          document.webkitExitFullscreen();
        } else if (document.msExitFullscreen) {
          document.msExitFullscreen();
        }
      }
      this.isFullscreen = !this.isFullscreen;
    },
    onLoadedMetadata() {
      const video = this.$refs.videoPlayer;
      this.duration = video.duration;
    },
    onTimeUpdate() {
      this.currentTime = this.$refs.videoPlayer.currentTime;
    },
  },
};
</script>

<style scoped>
.video-container {
  position: relative;
  width: 100%;
  height: 0;
  padding-bottom: 56.25%; /* 16:9 Aspect Ratio */
  overflow: hidden;
}

.video-player {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

.controls {
  position: absolute;
  bottom: 10px;
  left: 10px;
  background: rgba(0, 0, 0, 0.5);
  padding: 10px;
  border-radius: 5px;
}

.fullscreen {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 9999;
}
</style>
