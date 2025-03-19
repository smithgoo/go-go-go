import 'package:flutter/material.dart';
import 'package:picture_wall/common/apis/apis.dart';
import 'package:picture_wall/common/apis/videos.dart';
import 'package:picture_wall/common/entities/entities.dart';
import 'package:get/get.dart';
import 'package:picture_wall/common/entities/videos.dart';
import 'package:pull_to_refresh/pull_to_refresh.dart';

import 'index.dart';

class MainController extends GetxController {
  MainController();
  final ScrollController scrollController = ScrollController();
  bool isLoadingMore = false;

  final RefreshController refreshController = RefreshController(
    initialRefresh: true,
  );

  /// 响应式成员变量
  var videosList = <Videos>[].obs; // 使用 RxList 来管理视频列表
  final state = MainState();
  var page = 1;

  void onRefresh() {
    page = 1;
    videosList.clear();
    getVideosList();
  }

  void onLoading() async {
    var pp = page++;
    print(pp);
    VideosListResponseEntity listInfo = await VideosApi.videosList(
      page: pp,
    );
    if (listInfo.videos != null) {
      videosList.addAll(listInfo.videos!); // 更新 RxList
      refreshController.loadComplete();
    } else {
      refreshController.loadFailed();
    }
  }

  /// 生命周期

  /// 初始
  @override
  void onInit() {
    super.onInit();
  }

  /// 可 async 拉取数据
  /// 可触发展示交互组件
  /// onInit 之后
  @override
  void onReady() {
    super.onReady();
    // getVideosList();
  }

  /// 关闭页面
  /// 可以缓存数据，关闭各种控制器
  /// dispose 之前
  @override
  void onClose() {
    super.onClose();
  }

  /// 被释放
  /// 手动释放各种内存资源
  @override
  void dispose() {
    super.dispose();
  }

  /// 拉取视频列表
  getVideosList() async {
    VideosListResponseEntity listInfo = await VideosApi.videosList();
    if (listInfo.videos != null) {
      videosList.addAll(listInfo.videos!); // 更新 RxList
      refreshController.refreshCompleted(resetFooterState: true);
    } else {
      refreshController.loadFailed();
    }
  }
}
