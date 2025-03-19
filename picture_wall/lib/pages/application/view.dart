import 'package:flutter/material.dart';
import 'package:picture_wall/common/values/values.dart';
import 'package:picture_wall/common/widgets/widgets.dart';
import 'package:picture_wall/pages/category/index.dart';
import 'package:picture_wall/pages/main/index.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:get/get.dart';

import 'index.dart';

class ApplicationPage extends GetView<ApplicationController> {
  // 顶部导航
  AppBar _buildAppBar() {
    return transparentAppBar(
        title: Obx(() => Text(
              controller.tabTitles[controller.state.page],
              style: TextStyle(
                color: AppColors.primaryText,
                fontFamily: 'Montserrat',
                fontSize: 18.sp,
                fontWeight: FontWeight.w600,
              ),
            )),
        actions: <Widget>[
          IconButton(
            icon: Icon(
              Icons.search,
              color: AppColors.primaryText,
            ),
            onPressed: () {},
          )
        ]);
  }

  // 内容页
  Widget _buildPageView() {
    return PageView(
      physics: const NeverScrollableScrollPhysics(),
      controller: controller.pageController,
      onPageChanged: controller.handlePageChanged,
      children: <Widget>[
        MainPage(),
        CategoryPage(),
        const Text('BookmarksPage'),
        const Text('AccountPage'),
      ],
    );
  }

  // 底部导航
  Widget _buildBottomNavigationBar() {
    return Obx(() => BottomNavigationBar(
          items: controller.bottomTabs,
          currentIndex: controller.state.page,
          fixedColor: AppColors.primaryElement,
          type: BottomNavigationBarType.fixed,
          onTap: controller.handleNavBarTap,
          showSelectedLabels: true,
          showUnselectedLabels: true,
        ));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: _buildAppBar(),
      body: _buildPageView(),
      bottomNavigationBar: _buildBottomNavigationBar(),
    );
  }
}
