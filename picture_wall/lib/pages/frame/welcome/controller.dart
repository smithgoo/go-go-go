import 'package:picture_wall/common/routers/routes.dart';
import 'package:picture_wall/common/store/store.dart';
import 'package:get/get.dart';

import 'index.dart';

class WelcomeController extends GetxController {
  final state = WelcomeState();

  WelcomeController();

  // 跳转 注册界面
  handleNavSignIn() async {
    await ConfigStore.to.saveAlreadyOpen();
    Get.offAndToNamed(AppRoutes.SIGN_IN);
  }

  @override
  void onReady() {
    super.onReady();
  }
}
