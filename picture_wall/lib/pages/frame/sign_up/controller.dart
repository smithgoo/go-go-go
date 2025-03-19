import 'package:flutter/material.dart';
import 'package:picture_wall/common/apis/apis.dart';
import 'package:picture_wall/common/entities/entities.dart';
import 'package:picture_wall/common/utils/utils.dart';
import 'package:picture_wall/common/widgets/widgets.dart';
import 'package:picture_wall/common/routers/routes.dart';
import 'package:get/get.dart';

import 'index.dart';

class SignUpController extends GetxController {
  var captchaId = ''.obs;
  var base64Image = ''.obs;
  SignUpController();

  /// obs 响应式变量 才写入 state
  final state = SignUpState();

  /// 业务变量

  // 验证码的控制器
  final TextEditingController captchaController = TextEditingController();
  final TextEditingController emailController = TextEditingController();
  final TextEditingController passController = TextEditingController();
  final TextEditingController phoneController = TextEditingController();

  /// 业务事件

  // 返回上一页
  handleNavPop() {
    Get.back();
  }

  // 提示信息
  handleTip() {
    toastInfo(msg: '这是注册界面');
  }

  // 忘记密码
  handleFogotPassword() {
    toastInfo(msg: '忘记密码');
  }

  // 执行注册操作
  handleSignUp() async {
    // if (!duCheckStringLength(fullnameController.value.text, 5)) {
    //   toastInfo(msg: '用户名不能小于5位');
    //   return;
    // }
    // if (!duIsEmail(emailController.value.text)) {
    //   toastInfo(msg: '请正确输入邮件');
    //   return;
    // }
    // if (!duCheckStringLength(passController.value.text, 6)) {
    //   toastInfo(msg: '密码不能小于6位');
    //   return;
    // }

    UserRegisterRequestEntity params = UserRegisterRequestEntity(
      email: emailController.value.text,
      password: duSHA256(passController.value.text),
      captcha_id: captchaId.value,
      captcha_answer: captchaController.value.text,
      phone: phoneController.value.text,
    );

    var userProfile = await UserAPI.register(
      params: params,
    );
    print(userProfile);
    if (userProfile != null) {
      toastInfo(msg: '注册成功');
    }
    Get.offAndToNamed(AppRoutes.SIGN_IN);
    // Get.back();
  }

  /// 生命周期

  // 页面载入完成
  @override
  void onReady() {
    super.onReady();
    getCaptchaImageInfo();
  }

  @override
  void dispose() {
    captchaController.dispose();
    emailController.dispose();
    passController.dispose();
    super.dispose();
  }

  getCaptchaImageInfo() async {
    UserCaptchaResponseEntity captchainfo = await UserAPI.captcha();
    captchaId.value = captchainfo.captcha_id;
    String base64Str = captchainfo.captcha_image.split(',').last;
    base64Image.value = base64Str; // 更新 observable 数据
  }
}
