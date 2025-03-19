import 'package:picture_wall/common/entities/entities.dart';
import 'package:picture_wall/common/entities/videos.dart';
import 'package:picture_wall/common/utils/utils.dart';
import 'package:picture_wall/common/values/values.dart';

/// 新闻
class VideosApi {
  static Future videosList({
    int page = 1,
    int size = 20,
    bool refresh = false,
    bool cacheDisk = false,
  }) async {
    var response = await HttpUtil().get(
      '/v2/getHomeAll',
      queryParameters: {
        'page': page,
        'size': size,
      }, // 传递页码作为参数
    );
    return VideosListResponseEntity.fromJson(response);
  }
}
