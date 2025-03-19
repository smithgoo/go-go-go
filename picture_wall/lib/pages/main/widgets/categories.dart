import 'package:flutter/material.dart';
import 'package:picture_wall/common/routers/names.dart';
import 'package:picture_wall/common/values/values.dart';
import 'package:get/get.dart';
import '../index.dart';

/// 分类导航
class NewsCategoriesWidget extends GetView<MainController> {
  NewsCategoriesWidget();

  @override
  Widget build(BuildContext context) {
    // 获取屏幕的高度
    double screenHeight = MediaQuery.of(context).size.height;

    // 假设顶部导航栏和底部导航栏的高度
    double topNavBarHeight = AppBar().preferredSize.height; // 顶部导航栏高度
    double bottomNavBarHeight = kBottomNavigationBarHeight; // 底部导航栏高度

    // 计算 GridView 的高度
    double gridViewHeight = screenHeight - topNavBarHeight - bottomNavBarHeight;

    return SizedBox(
      height: gridViewHeight, // 设置固定高度
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: Obx(
          () {
            // 检查 videosList 是否为空
            if (controller.videosList.isEmpty) {
              return const Center(
                child: Text(
                  'No Videos Available',
                  style: TextStyle(fontSize: 18, color: Colors.grey),
                ),
              );
            }

            // 使用 GridView.builder 构建网格视图
            return GridView.builder(
              gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                crossAxisCount: 3,
                crossAxisSpacing: 10.0,
                mainAxisSpacing: 10.0,
                childAspectRatio: 230 / 322, // 设置每个 item 的宽高比为 9:16
              ),
              itemCount: controller.videosList.length, // 使用 videosList 的长度
              itemBuilder: (context, index) {
                // 获取当前视频的名称
                String videoTitle =
                    controller.videosList[index].title as String;
                 String cover =
                    controller.videosList[index].cover as String;
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
                          style: const TextStyle(
                              color: Colors.white, fontSize: 16),
                        ),
                      ],
                    ),
                  ),
                );
              },
            );
          },
        ),
      ),
    );
  }
}
