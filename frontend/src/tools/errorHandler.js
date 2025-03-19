// utils/errorHandler.js
import { Toast } from "vant";

export function handleError(error, fetchCaptchaCallback) {
  // 检查是否为 401 错误
  if (error.response && error.response.status === 401) {
    // 401 Unauthorized 错误处理
    Toast.fail("Session expired. Please login again.");
    // 可选择将用户重定向到登录页面
    window.location.href = "/user/login"; // 或者使用 Vue 路由：this.$router.push('/login');
    return;
  }

  // 提取其他错误信息
  const errorMsg =
    (error.response && error.response.data && error.response.data.error) ||
    (error.request && JSON.parse(error.request.response).error) ||
    "An unexpected error occurred. Please try again.";

  // 如果有刷新验证码的回调函数，调用它
  if (typeof fetchCaptchaCallback === "function") {
    fetchCaptchaCallback();
  }

  // 使用 Toast 弹出错误信息
  Toast.fail(errorMsg);
}
