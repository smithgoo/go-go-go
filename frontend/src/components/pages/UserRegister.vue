<template>
  <div class="register-container">
    <div class="register-form">
      <h1>Register</h1>
      <van-form @submit.prevent="register">
        <van-field
          v-model="email"
          label="Email"
          type="email"
          placeholder="Enter your email"
          required
          class="form-field"
        />
        <van-field
          v-model="phone"
          label="Phone"
          type="tel"
          placeholder="Enter your phone"
          required
          class="form-field"
        />
        <van-field
          v-model="password"
          label="Password"
          type="password"
          placeholder="Enter your password"
          required
          class="form-field"
        />
        <van-field
          v-model="captchaAnswer"
          label="Captcha"
          type="text"
          placeholder="Enter the captcha"
          required
          class="form-field"
        >
          <template #button>
            <img
              :src="captchaImageSrc"
              alt="Captcha Image"
              @click="fetchCaptcha"
              class="captcha-image"
            />
          </template>
        </van-field>
        <input type="hidden" v-model="captchaID" name="captcha_id" />
        <van-button
          type="primary"
          native-type="submit"
          class="register-btn"
          @click="register"
          >Register</van-button
        >
      </van-form>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import { Toast } from "vant";
import { handleError } from "../../tools/errorHandler";

export default {
  name: "UserRegister",
  data() {
    return {
      email: "",
      phone: "",
      password: "",
      captchaImageSrc: "",
      captchaID: "",
      captchaAnswer: "",
    };
  },
  created() {
    this.fetchCaptcha();
  },
  methods: {
    async register() {
      try {
        const formData = new FormData();
        formData.append("email", this.email);
        formData.append("phone", this.phone);
        formData.append("password", this.password);
        formData.append("captcha_id", this.captchaID);
        formData.append("captcha_answer", this.captchaAnswer);

        const response = await axios.post("/user/register", formData);
        Toast.success("Registration Successful!");

        if (response) {
          this.$router.push("/login");
        }
      } catch (error) {
        handleError(error);
      }
    },
    async fetchCaptcha() {
      try {
        const response = await axios.get("/user/captcha");
        this.captchaImageSrc = response.data.captcha_image;
        this.captchaID = response.data.captcha_id;
      } catch (error) {
        handleError(error);
      }
    },
  },
};
</script>

<style scoped>
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f5f5f5;
  margin: 0;
}

.register-form {
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #fff;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  width: 90%;
  max-width: 600px;
  height: 100%;
}

h1 {
  text-align: center;
  margin-bottom: 20px;
}

.form-field {
  width: 100%;
  margin-bottom: 5px;
  height: 40px;
}

.captcha-image {
  height: 25px;
  margin-left: 10px;
}

.register-btn {
  width: 100%;
  padding: 10px;
  border-radius: 5px;
  background-color: #4caf50;
  color: white;
  border: none;
  cursor: pointer;
}

.register-btn:hover {
  background-color: #45a049;
}

@media (max-width: 768px) {
  .register-form {
    width: 90%;
    max-width: none;
  }
}
</style>
