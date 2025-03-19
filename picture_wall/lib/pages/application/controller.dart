import 'dart:async';

import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:picture_wall/common/routers/routes.dart';
import 'package:picture_wall/common/utils/utils.dart';
import 'package:picture_wall/common/values/values.dart';
import 'package:get/get.dart';
import 'package:uni_links/uni_links.dart';

import 'index.dart';

class ApplicationController extends GetxController {
  ApplicationController();

  /// 响应式成员变量

  final state = ApplicationState();

  /// 成员变量

  // tab 页标题
  late final List<String> tabTitles;

  // 页控制器
  late final PageController pageController;

  // 底部导航项目
  late final List<BottomNavigationBarItem> bottomTabs;

  /// 事件

  // tab栏动画
  void handleNavBarTap(int index) {
    pageController.animateToPage(index,
        duration: const Duration(milliseconds: 200), curve: Curves.ease);
  }

  // tab栏页码切换
  void handlePageChanged(int page) {
    state.page = page;
  }

  /// scheme 内部打开
  bool isInitialUriIsHandled = false;
  StreamSubscription? uriSub;

  // 第一次打开
  Future<void> handleInitialUri() async {
    if (!isInitialUriIsHandled) {
      isInitialUriIsHandled = true;
      try {
        final uri = await getInitialUri();
        if (uri == null) {
          print('no initial uri');
        } else {
          // 这里获取了 scheme 请求
          print('got initial uri: $uri');
        }
      } on PlatformException {
        print('falied to get initial uri');
      } on FormatException catch (err) {
        print('malformed initial uri, ' + err.toString());
      }
    }
  }

  // 程序打开时介入
  void handleIncomingLinks() {
    if (!kIsWeb) {
      uriSub = uriLinkStream.listen((Uri? uri) {
        // 这里获取了 scheme 请求
        print('got uri: $uri');

        // if (uri!.pathSegments[1].toLowerCase() == 'category') {
        if (uri != null && uri.path == '/notify/category') {
          Get.toNamed(AppRoutes.Category);
        }
      }, onError: (Object err) {
        print('got err: $err');
      });
    }
  }

  /// 生命周期

  @override
  void onInit() {
    super.onInit();

    // handleInitialUri();
    // handleIncomingLinks();

    // 准备一些静态数据
    tabTitles = ['Welcome', 'Cagegory', 'Bookmarks', 'Account'];
    bottomTabs = <BottomNavigationBarItem>[
      const BottomNavigationBarItem(
        icon: Icon(
          Icons.home,
          color: AppColors.tabBarElement,
        ),
        activeIcon: Icon(
          Icons.home,
          color: AppColors.secondaryElementText,
        ),
        label: 'main',
      ),
      const BottomNavigationBarItem(
        icon: Icon(
          Icons.gamepad,
          color: AppColors.tabBarElement,
        ),
        activeIcon: Icon(
          Icons.gamepad,
          color: AppColors.secondaryElementText,
        ),
        label: 'category',
      ),
      const BottomNavigationBarItem(
        icon: Icon(
          Icons.tag,
          color: AppColors.tabBarElement,
        ),
        activeIcon: Icon(
          Icons.tag,
          color: AppColors.secondaryElementText,
        ),
        label: 'tag',
      ),
      const BottomNavigationBarItem(
        icon: Icon(
          Icons.mail,
          color: AppColors.tabBarElement,
        ),
        activeIcon: Icon(
          Icons.mail,
          color: AppColors.secondaryElementText,
        ),
        label: 'my',
      ),
    ];

    pageController = new PageController(initialPage: state.page);
  }

  @override
  void onReady() {
    super.onReady();
  }

  @override
  void onClose() {
    super.onClose();
  }

  @override
  void dispose() {
    uriSub?.cancel();
    pageController.dispose();
    super.dispose();
  }
}
