import 'package:picture_wall/common/entities/entities.dart';
import 'package:picture_wall/common/utils/utils.dart';

/// 用户
class UserAPI {
  /// 登录
  static Future<UserLoginResponseEntity> login({
    UserLoginRequestEntity? params,
  }) async {
    var response = await HttpUtil().postFormData(
      '/user/login',
      data: params?.toJson(),
    );
    print(response);
    return UserLoginResponseEntity.fromJson(response);
  }

  /// 注册
  static Future register({
    UserRegisterRequestEntity? params,
  }) async {
    var response = await HttpUtil().postFormData(
      '/user/register',
      data: params?.toJson(),
    );
    print(response);
    return response;
  }

  /// Profile
  static Future<UserLoginResponseEntity> profile() async {
    var response = await HttpUtil().post(
      '/user/profile',
    );
    return UserLoginResponseEntity.fromJson(response);
  }

  /// Logout
  static Future logout() async {
    return await HttpUtil().post(
      '/user/logout',
    );
  }

  /// captcha
  static Future captcha() async {
    var response = await HttpUtil().get(
      '/user/captcha',
    );
    return UserCaptchaResponseEntity.fromJson(response);
  }
}
