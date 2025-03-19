import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:get/get.dart';
import 'package:picture_wall/common/routers/names.dart';
import 'package:pull_to_refresh/pull_to_refresh.dart';

import 'index.dart';
import 'widgets/widgets.dart';

class MainPage extends GetView<MainController> {
  const MainPage();

  @override
  Widget build(BuildContext context) {
    return GetX<MainController>(
      init: controller,
      builder: (controller) => SmartRefresher(
        enablePullUp: true,
        controller: controller.refreshController,
        onRefresh: controller.onRefresh,
        onLoading: controller.onLoading,
        child:GridView.builder(
      gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
        crossAxisCount: 3,
        crossAxisSpacing: 10.0,
        mainAxisSpacing: 10.0,
        childAspectRatio: 230 / 322, // 设置每个 item 的宽高比为 9:16
      ),
      itemCount: controller.videosList.length, // 使用 videosList 的长度
      itemBuilder: (context, index) {
        // 获取当前视频的名称
        String videoTitle = controller.videosList[index].title as String;
        String cover = controller.videosList[index].cover as String;
        return GestureDetector(
          onTap: () {
            // 点击后跳转到视频播放器页面
            Get.toNamed(AppRoutes.VideoPlayer);
          },
          child: Container(
            color: Colors.blueAccent,
            alignment: Alignment.center,
            child: Stack(
              children: [
                Image.network(cover),
                Text(
                  videoTitle, // 显示视频的名称
                  style: const TextStyle(color: Colors.white, fontSize: 16),
                ),
              ],
            ),
          ),
        );
      },
    ),),);
  }
}
