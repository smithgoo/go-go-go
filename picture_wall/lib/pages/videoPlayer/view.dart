import 'package:flutter/material.dart';
import 'package:get/get.dart';

import 'index.dart';
import 'widgets/widgets.dart';

class VideoplayerPage extends GetView<VideoplayerController> {
  const VideoplayerPage({Key? key}) : super(key: key);

  // 主视图
  Widget _buildView() {
    return const HelloWidget();
  }

  @override
  Widget build(BuildContext context) {
    return GetBuilder<VideoplayerController>(
      builder: (_) {
        return Scaffold(
          appBar: AppBar(title: const Text("videoplayer")),
          body: SafeArea(
            child: _buildView(),
          ),
        );
      },
    );
  }
}
