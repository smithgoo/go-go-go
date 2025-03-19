import 'package:get/get.dart';

class VideoplayerState {
  // title
  final _title = "".obs;
  set title(value) => _title.value = value;
  get title => _title.value;
}
