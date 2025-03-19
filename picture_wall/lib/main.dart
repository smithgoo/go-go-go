import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:picture_wall/common/langs/translation_service.dart';
import 'package:picture_wall/common/routers/pages.dart';
import 'package:picture_wall/common/store/config.dart';
import 'package:picture_wall/common/style/theme.dart';
import 'package:picture_wall/common/utils/logger.dart';
import 'package:picture_wall/global.dart';
import 'package:get/get.dart';

Future<void> main() async {
  await Global.init();
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    
    return ScreenUtilInit(
      designSize: const Size(375, 812),
      builder: (context, child) => GetMaterialApp(
        title: 'News',
        theme: AppTheme.light,
        debugShowCheckedModeBanner: false,
        initialRoute: AppPages.INITIAL,
        getPages: AppPages.routes,
        builder: EasyLoading.init(),
        translations: TranslationService(),
        navigatorObservers: [AppPages.observer],
        // localizationsDelegates: [
        //   GlobalMaterialLocalizations.delegate,
        //   GlobalWidgetsLocalizations.delegate,
        //   GlobalCupertinoLocalizations.delegate,
        // ],
        supportedLocales: ConfigStore.to.languages,
        locale: ConfigStore.to.locale,
        fallbackLocale: Locale('en', 'US'),
        enableLog: true,
        logWriterCallback: Logger.write,
      ),
    );
  }
}
