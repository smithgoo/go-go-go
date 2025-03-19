import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:picture_wall/common/routers/routes.dart';
import 'package:picture_wall/common/services/storage.dart';
import 'package:picture_wall/common/store/store.dart';

import 'package:get/get.dart';
import 'package:picture_wall/common/values/storage.dart';

/// 检查是否登录
class RouteAuthMiddleware extends GetMiddleware {
  // priority 数字小优先级高
  @override
  int? priority = 0;

  RouteAuthMiddleware({required this.priority});

  @override
  RouteSettings? redirect(String? route) {
    if (UserStore.to.isLogin ||
        route == AppRoutes.SIGN_IN ||
        route == AppRoutes.SIGN_UP ||
        route == AppRoutes.INITIAL) {
      return null;
    } else {
      // var info = StorageService.to.getString(STORAGE_USER_PROFILE_KEY);
      // var islogin = false;

      // if (info != null && info.isNotEmpty) {
      //   var tmpinfo = jsonDecode(info);

      //   if (tmpinfo['token'] != null && tmpinfo['token'].isNotEmpty) {
      //     islogin = true;
      //   }
      // }
      // if (islogin) {
      //   return const RouteSettings(name: AppRoutes.Application);
      // } else {
        Future.delayed(
            Duration(seconds: 1), () => Get.snackbar("提示", "登录过期,请重新登录"));
        return const RouteSettings(name: AppRoutes.SIGN_IN);
      // }
    }
  }
}
