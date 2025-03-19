<template>
  <div class="container">
    <form @submit.prevent="login" id="loginForm">
      <h1>Welcome to Login</h1>
      <label for="email">Email or Phone:</label>
      <van-field
        v-model="email"
        placeholder="Input your email"
        border="false"
        style="margin-bottom: 20px"
      />
      <label for="password">Password:</label>
      <van-field
        type="password"
        v-model="password"
        placeholder="Input your password"
        border="false"
        style="margin-bottom: 20px"
      />
      <label for="captchaAnswer">Captcha:</label>
      <div id="captcha-container">
        <van-field
          type="text"
          placeholder="Input captcha code"
          v-model="captchaAnswer"
          id="captchaAnswer"
          name="captcha_answer"
          border="false"
          style="flex: 1; margin-right: 10px"
        />
        <img
          class="captcha-image"
          :src="captchaImageSrc"
          alt="Captcha Image"
          v-if="captchaImageSrc"
          @click="fetchCaptcha"
        />
      </div>
      <input
        type="hidden"
        v-model="captchaID"
        id="captchaID"
        name="captcha_id"
      />
      <div class="bottom-btn-container">
        <van-button type="primary" class="bottom-btn" @click="login"
          >Login</van-button
        >
        <van-button type="default" class="bottom-btn" @click="goToRegister"
          >Register</van-button
        >
      </div>
    </form>
  </div>
</template>

<script>
import axios from "../../tools/axios";
import userStorage from "../../tools/storage";
import { Toast } from "vant";
import "vant/lib/index.css";

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
    };
  },
  created() {
    this.fetchCaptcha();
  },
  methods: {
    async login() {
      try {
        const formData = new FormData();
        formData.append("email", this.email);
        formData.append("password", this.password);
        formData.append("captcha_id", this.captchaID);
        formData.append("captcha_answer", this.captchaAnswer);

        const response = await axios.post("/user/login", formData);
        this.message = "Login successful!";
        userStorage.setItem("user", "Bearer " + response.data.token);
        userStorage.setItem("account", this.email);
        userStorage.setItem("isAdmin", response.data.user.role_id);

        if (response.data.user.role_id != 3) {
          Toast.success(this.message);
          this.goToHome();
        } else {
          Toast.success("please contact admin for access");
        }
      } catch (error) {
        this.fetchCaptcha();
        this.error = "Login failed. Please check your credentials.";
        Toast.fail(this.error);
      }
    },
    async fetchCaptcha() {
      try {
        const response = await axios.get("/user/captcha");
        this.captchaImageSrc = response.data.captcha_image;
        this.captchaID = response.data.captcha_id;
      } catch (error) {
        this.error = "Error fetching captcha. Please try again later.";
      }
    },
    goToRegister() {
      this.$router.push("/register");
    },
    goToHome() {
      this.$router.push("/home");
    },
  },
};
</script>

<style scoped>
html,
body {
  height: 100%;
  margin: 0;
}

body {
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #f4f4f4;
}

.container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100vh; /* Ensure full viewport height */
}

form {
  background-color: #fff;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 600px;
  box-sizing: border-box;
  height: 100%;
}

h1 {
  text-align: center;
  margin-bottom: 20px;
}

label {
  display: block;
  margin-bottom: 8px;
}

.van-field {
  width: calc(100% - 20px);
}

#captcha-container {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.captcha-image {
  cursor: pointer;
  margin-left: 10px;
  height: 40px;
  width: 120px;
}

.bottom-btn-container {
  display: flex;
  flex-direction: row;
  width: 100%;
  justify-content: space-between;
  margin-top: 20px;
}

.bottom-btn {
  width: 45%;
  height: 40px;
}
</style>
