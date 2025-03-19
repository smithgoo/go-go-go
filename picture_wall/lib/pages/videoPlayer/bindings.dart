import 'package:get/get.dart';

import 'controller.dart';

class VideoplayerBinding implements Bindings {
  @override
  void dependencies() {
    Get.lazyPut<VideoplayerController>(() => VideoplayerController());
  }
}
