import 'dart:convert';
import 'dart:typed_data';

import 'package:flutter/material.dart';
import 'package:picture_wall/common/apis/apis.dart';
import 'package:picture_wall/common/entities/entities.dart';
import 'package:picture_wall/common/routers/routes.dart';
import 'package:picture_wall/common/store/store.dart';
import 'package:picture_wall/common/utils/utils.dart';
import 'package:picture_wall/common/widgets/widgets.dart';
import 'package:get/get.dart';

import 'index.dart';

class SignInController extends GetxController {
  final state = SignInState();
  var captchaId = ''.obs;
  var base64Image = ''.obs;

  SignInController();

  // email的控制器
  final TextEditingController emailController = TextEditingController();
  // 密码的控制器
  final TextEditingController passController = TextEditingController();

// 验证码的控制器
  final TextEditingController captchaController = TextEditingController();
  // final MyRepository repository;
  // SignInController({@required this.repository}) : assert(repository != null);

  // 跳转 注册界面
  handleNavSignUp() {
    Get.toNamed(AppRoutes.SIGN_UP);
  }

  // 忘记密码
  handleFogotPassword() {
    toastInfo(msg: '忘记密码');
  }

  // 执行登录操作
  handleSignIn() async {
    // if (!duIsEmail(_emailController.value.text)) {
    //   toastInfo(msg: '请正确输入邮件');
    //   return;
    // }
    // if (!duCheckStringLength(_passController.value.text, 6)) {
    //   toastInfo(msg: '密码不能小于6位');
    //   return;
    // }
    try {
        UserLoginRequestEntity params = UserLoginRequestEntity(
        email: emailController.value.text,
        password: duSHA256(passController.value.text),
        captcha_id: captchaId.value,
        captcha_answer: captchaController.value.text,
      );

      UserLoginResponseEntity userProfile = await UserAPI.login(
        params: params,
      );
      UserStore.to.saveProfile(userProfile);

      Get.offAndToNamed(AppRoutes.Application);
    } catch (e) {
      getCaptchaImageInfo();
    }
    
  }

  @override
  void onReady() {
    super.onReady();
    getCaptchaImageInfo();
  }

  @override
  void dispose() {
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
