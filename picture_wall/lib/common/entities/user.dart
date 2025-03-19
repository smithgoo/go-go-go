// 注册请
// 登录请求
class UserLoginRequestEntity {
  String email;
  String password;
  String captcha_id;
  String captcha_answer;

  UserLoginRequestEntity({
    required this.email,
    required this.password,
    required this.captcha_id,
    required this.captcha_answer,
  });

  factory UserLoginRequestEntity.fromJson(Map<String, dynamic> json) =>
      UserLoginRequestEntity(
        email: json["email"],
        password: json["password"],
        captcha_id: json["captcha_id"],
        captcha_answer: json["captcha_answer"],
      );

  Map<String, dynamic> toJson() => {
        "email": email,
        "password": password,
        "captcha_id": captcha_id,
        "captcha_answer": captcha_answer,
      };
}

class UserRegisterRequestEntity {
  String email;
  String password;
  String captcha_id;
  String captcha_answer;
  String phone;

  UserRegisterRequestEntity({
    required this.email,
    required this.phone,
    required this.password,
    required this.captcha_id,
    required this.captcha_answer,
  });

  factory UserRegisterRequestEntity.fromJson(Map<String, dynamic> json) =>
      UserRegisterRequestEntity(
        email: json["email"],
        phone: json["phone"],
        password: json["password"],
        captcha_id: json["captcha_id"],
        captcha_answer: json["captcha_answer"],
      );

  Map<String, dynamic> toJson() => {
        "email": email,
        "phone": phone,
        "password": password,
        "captcha_id": captcha_id,
        "captcha_answer": captcha_answer,
      };
}

class UserCaptchaResponseEntity {
  String captcha_id;
  String captcha_image;

  UserCaptchaResponseEntity({
    required this.captcha_id,
    required this.captcha_image,
  });

  factory UserCaptchaResponseEntity.fromJson(Map<String, dynamic> json) =>
      UserCaptchaResponseEntity(
        captcha_id: json["captcha_id"],
        captcha_image: json["captcha_image"],
      );

  Map<String, dynamic> toJson() => {
        "captcha_id": captcha_id,
        "captcha_image": captcha_image,
      };
}

// 登录返回
class UserLoginResponseEntity {
  String? token;
  User? user;

  UserLoginResponseEntity({this.token, this.user});

  UserLoginResponseEntity.fromJson(Map<String, dynamic> json) {
    token = json['token'];
    user = json['user'] != null ? new User.fromJson(json['user']) : null;
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['token'] = this.token;
    if (this.user != null) {
      data['user'] = this.user!.toJson();
    }
    return data;
  }
}

class User {
  int? iD;
  String? createdAt;
  String? updatedAt;
  Null? deletedAt;
  String? email;
  String? phone;
  String? password;
  int? roleId;
  Role? role;

  User(
      {this.iD,
      this.createdAt,
      this.updatedAt,
      this.deletedAt,
      this.email,
      this.phone,
      this.password,
      this.roleId,
      this.role});

  User.fromJson(Map<String, dynamic> json) {
    iD = json['ID'];
    createdAt = json['CreatedAt'];
    updatedAt = json['UpdatedAt'];
    deletedAt = json['DeletedAt'];
    email = json['email'];
    phone = json['phone'];
    password = json['password'];
    roleId = json['role_id'];
    role = json['role'] != null ? new Role.fromJson(json['role']) : null;
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['ID'] = this.iD;
    data['CreatedAt'] = this.createdAt;
    data['UpdatedAt'] = this.updatedAt;
    data['DeletedAt'] = this.deletedAt;
    data['email'] = this.email;
    data['phone'] = this.phone;
    data['password'] = this.password;
    data['role_id'] = this.roleId;
    if (this.role != null) {
      data['role'] = this.role!.toJson();
    }
    return data;
  }
}

class Role {
  int? iD;
  String? createdAt;
  String? updatedAt;
  Null? deletedAt;
  String? name;

  Role({this.iD, this.createdAt, this.updatedAt, this.deletedAt, this.name});

  Role.fromJson(Map<String, dynamic> json) {
    iD = json['ID'];
    createdAt = json['CreatedAt'];
    updatedAt = json['UpdatedAt'];
    deletedAt = json['DeletedAt'];
    name = json['name'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['ID'] = this.iD;
    data['CreatedAt'] = this.createdAt;
    data['UpdatedAt'] = this.updatedAt;
    data['DeletedAt'] = this.deletedAt;
    data['name'] = this.name;
    return data;
  }
}
