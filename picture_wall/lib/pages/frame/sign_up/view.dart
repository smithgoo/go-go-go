import 'dart:convert';
import 'dart:typed_data';

import 'package:flutter/material.dart';
import 'package:picture_wall/common/values/values.dart';
import 'package:picture_wall/common/widgets/widgets.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:get/get.dart';

import 'index.dart';

class SignUpPage extends GetView<SignUpController> {
  // logo
  Widget _buildLogo() {
    return Container(
      margin: EdgeInsets.only(top: 50.h),
      child: Text(
        "Sign up",
        textAlign: TextAlign.center,
        style: TextStyle(
          color: AppColors.primaryText,
          fontFamily: "Montserrat",
          fontWeight: FontWeight.w600,
          fontSize: 24.sp,
          height: 1,
        ),
      ),
    );
  }

  // 注册表单
  Widget _buildInputForm() {
    return Container(
      width: 295.w,
      // height: 204,
      margin: EdgeInsets.only(top: 49.h),
      child: Column(
        children: [
          // email input
          inputTextEdit(
            controller: controller.emailController,
            keyboardType: TextInputType.emailAddress,
            hintText: "Email",
          ),
          inputTextEdit(
            controller: controller.phoneController,
            keyboardType: TextInputType.phone,
            hintText: "Phone",
          ),
          // password input
          inputTextEdit(
            controller: controller.passController,
            keyboardType: TextInputType.visiblePassword,
            hintText: "Password",
            isPassword: true,
          ),
          Container(
            height: 44.h,
            margin: EdgeInsets.only(top: 15.h),
            child: Row(
              children: [
                // 注册
                Container(
                  width: 150,
                  height: 44,
                  child: inputTextEdit(
                    controller: controller.captchaController,
                    keyboardType: TextInputType.number,
                    hintText: "Captcha",
                    marginTop: 0,
                    // autofocus: true,
                  ),
                ),
                Spacer(),
                // 登录
                GestureDetector(
                  onTap: () {
                    // 点击事件触发时执行的逻辑
                    controller.getCaptchaImageInfo(); // 刷新验证码图片
                  },
                  child: SizedBox(
                    width: 150,
                    height: 44,
                    child: Center(
                      child: Obx(() {
                        if (controller.base64Image.isEmpty) {
                          return const CircularProgressIndicator(); // 加载动画
                        } else {
                          Uint8List bytes =
                              base64Decode(controller.base64Image.value);
                          return Image.memory(bytes); // 显示解码后的图片
                        }
                      }),
                    ),
                  ),
                ),
              ],
            ),
          ),

          // 创建
          Container(
            height: 44.h,
            margin: EdgeInsets.only(top: 15.h),
            child: btnFlatButtonWidget(
              onPressed: controller.handleSignUp,
              width: 295.w,
              fontWeight: FontWeight.w600,
              title: "Create an account",
            ),
          ),
          // Spacer(),

          // Fogot password
          Padding(
            padding: EdgeInsets.only(top: 8.0),
            child: TextButton(
              onPressed: controller.handleFogotPassword,
              child: Text(
                "Fogot password?",
                textAlign: TextAlign.center,
                style: TextStyle(
                  color: AppColors.secondaryElementText,
                  fontFamily: "Avenir",
                  fontWeight: FontWeight.w400,
                  fontSize: 16.sp,
                  height: 1, // 设置下行高，否则字体下沉
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }

  // 第三方
  Widget _buildThirdPartyLogin() {
    return Container(
      width: 295.w,
      margin: EdgeInsets.only(bottom: 40.h),
      child: Column(
        children: <Widget>[
          // title
          Text(
            "Or sign in with social networks",
            textAlign: TextAlign.center,
            style: TextStyle(
              color: AppColors.primaryText,
              fontFamily: "Avenir",
              fontWeight: FontWeight.w400,
              fontSize: 16.sp,
            ),
          ),
          // 按钮
          Padding(
            padding: EdgeInsets.only(top: 20.h),
            child: Row(
              children: <Widget>[
                btnFlatButtonBorderOnlyWidget(
                  onPressed: () {},
                  width: 88,
                  iconFileName: "twitter",
                ),
                Spacer(),
                btnFlatButtonBorderOnlyWidget(
                  onPressed: () {},
                  width: 88,
                  iconFileName: "google",
                ),
                Spacer(),
                btnFlatButtonBorderOnlyWidget(
                  onPressed: () {},
                  width: 88,
                  iconFileName: "facebook",
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  // 有账号
  Widget _buildHaveAccountButton() {
    return Container(
      margin: EdgeInsets.only(bottom: 20.h),
      child: btnFlatButtonWidget(
        onPressed: controller.handleNavPop,
        width: 294,
        gbColor: AppColors.secondaryElement,
        fontColor: AppColors.primaryText,
        title: "I have an account",
        fontWeight: FontWeight.w500,
        fontSize: 16,
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      resizeToAvoidBottomInset: false,
      appBar: transparentAppBar(
        leading: IconButton(
          icon: Icon(
            Icons.arrow_back,
            color: AppColors.primaryText,
          ),
          onPressed: controller.handleNavPop,
        ),
        actions: <Widget>[
          IconButton(
            icon: Icon(
              Icons.info_outline,
              color: AppColors.primaryText,
            ),
            onPressed: controller.handleTip,
          )
        ],
      ),
      body: Center(
        child: Column(
          children: <Widget>[
            Divider(height: 1),
            _buildLogo(),
            _buildInputForm(),
            Spacer(),
            _buildThirdPartyLogin(),
            _buildHaveAccountButton(),
          ],
        ),
      ),
    );
  }
}
