import 'dart:convert';

import 'package:picture_wall/common/apis/apis.dart';
import 'package:picture_wall/common/entities/entities.dart';
import 'package:picture_wall/common/services/services.dart';
import 'package:picture_wall/common/values/values.dart';
import 'package:get/get.dart';

class UserStore extends GetxController {
  static UserStore get to => Get.find();

  // 是否登录
  final _isLogin = false.obs;
  // 令牌 token
  String token = '';
  // 用户 profile
  final _profile = UserLoginResponseEntity().obs;

  bool get isLogin => _isLogin.value;
  UserLoginResponseEntity get profile => _profile.value;
  bool get hasToken => token.isNotEmpty;

  @override
  void onInit() {
    super.onInit();
    token = StorageService.to.getString(STORAGE_USER_TOKEN_KEY);
    var profileOffline = StorageService.to.getString(STORAGE_USER_PROFILE_KEY);
    if (profileOffline.isNotEmpty) {
      _profile(UserLoginResponseEntity.fromJson(jsonDecode(profileOffline)));
    }
  }

  // 保存 token
  Future<void> setToken(String value) async {
    await StorageService.to.setString(STORAGE_USER_TOKEN_KEY, value);
    token = value;
  }

  // 获取 profile
  Future<Map<String, dynamic>> getProfile() async {
  if (token.isEmpty) {
    return {}; // 如果 token 为空，返回一个空的 Map
  }

  var result = await UserAPI.profile();
  _profile(result); // 假设 _profile 是处理用户信息的方法
  _isLogin.value = true;

  // 从 StorageService 中获取用户信息
  var info = StorageService.to.getString(STORAGE_USER_PROFILE_KEY);
  if (info.isNotEmpty) {
    var response = jsonDecode(info);
    return response; // 直接返回解码后的 JSON 数据
  } else {
    return {}; // 如果没有用户信息，返回一个空的 Map
  }
}


  // 保存 profile
  Future<void> saveProfile(UserLoginResponseEntity profile) async {
    _isLogin.value = true;
    // StorageService.to.setString(STORAGE_IS_LOGIN, true as String);
    StorageService.to.setString(STORAGE_USER_PROFILE_KEY, jsonEncode(profile));
  }

  // 注销
  Future<void> onLogout() async {
    if (_isLogin.value) await UserAPI.logout();
    await StorageService.to.remove(STORAGE_USER_TOKEN_KEY);
    _isLogin.value = false;
    token = '';
  }
}
