<script>
import { getRequest, postRequest } from "../../tools/axios";
import { handleError } from "@/tools/errorHandler";
import userStorage from "../../tools/storage";
export default {
  name: "UserAdmin",
  data() {
    return {
      account: "",
      searchQuery: "", // 存储搜索框的输入值
      searchResults: [], // 存储搜索结果
      item: {
        role_id: 1, // or whatever value you need
      },
      message: "",
      list: [
        {
          type: "primary",
          size: "small",
          // style: {
          //   backgroundColor: this.getButtonColor(this.item.role_id, 1),
          // },
          role: "admin",
          index: 1,
        },
        {
          type: "primary",
          size: "small",
          // style: {
          //   backgroundColor: this.getButtonColor(this.item.role_id, 2),
          // },
          role: "viewer",
          index: 2,
        },
        {
          type: "primary",
          size: "small",
          // style: {
          //   backgroundColor: this.getButtonColor(this.item.role_id, 3),
          // },
          role: "user",
          index: 3,
        },
      ],
    };
  },
  methods: {
    getUsersList() {
      getRequest("/user/getAllUser", {}, true)
        .then((data) => {
          console.log(data.users);
          let arr = data.users;
          arr.forEach((item) => {
            console.log(item);
            if (item.email === userStorage.getItem("account")) {
              arr.splice(arr.indexOf(item), 1);
            }
          });
          this.list = arr;
        })
        .catch((error) => {
          console.log(error);
        });
    },
    editAction(item, newRole) {
      console.log(item.ID);
      console.log(newRole);
      this.item = item;
      postRequest(
        "/user/edit/role",
        {
          user_id: item.ID,
          role_id: newRole,
          is_admin: userStorage.getItem("isAdmin"),
        },
        true
      )
        .then((response) => {
          console.log(response);
          if (response) {
            this.list = [];
            this.getUsersList();
          }
        })
        .catch((error) => {
          console.log(error);
          handleError(error);
        });
    },
    getButtonColor(currentRoleId, buttonRoleId) {
      if (currentRoleId === buttonRoleId) {
        // 根据 role_id 设置不同颜色
        if (buttonRoleId === 1) return "#409eff"; // 蓝色 - admin
        if (buttonRoleId === 2) return "#67c23a"; // 绿色 - user
        if (buttonRoleId === 3) return "#ff6699"; // 橙色 - viewer
        if (buttonRoleId == 4) return "#f56c6c"; // 红色 - prohibit
      }
      return "#ccc"; // 默认灰色
    },
    getUserAccount() {
      return userStorage.getItem("account");
    },
    goToLogin() {
      userStorage.clear();
      this.$router.push("/login");
    },
    onSearch() {
      // 确保 dataList 是定义的
      if (!Array.isArray(this.dataList)) {
        console.error("dataList is not an array");
        return;
      }

      // 确保 searchQuery 是一个字符串
      const query = this.searchQuery || "";

      // 处理搜索逻辑，从本地数组中过滤结果
      this.searchResults = this.dataList.filter((item) =>
        item.name.toLowerCase().includes(query.toLowerCase())
      );
      this.list = this.searchResults;
    },
    onCancel() {
      // 清空搜索框和结果
      this.searchQuery = "";
      this.searchResults = [];
      this.list = this.searchResults;
    },
  },
  created() {
    this.getUsersList();
    this.account = this.getUserAccount() || "";
  },
};
</script>

<template>
  <h1>setting page</h1>

  <div class="userinfo-method">
    <van-search
      v-model="searchQuery"
      placeholder="input search keyword"
      @search="onSearch"
      @cancel="onCancel"
    />
    <h6>{{ account }}</h6>
    <van-button class="login-out" @click="goToLogin">logout</van-button>
  </div>
  <van-cell-group class="bottom-group-container">
    <!-- 使用 v-for 循环遍历 list -->
    <van-cell
      v-for="(item, index) in list"
      :key="index"
      :title="Email"
      :value="item.email"
      :label="item.phone"
    >
      <!-- 在右侧添加一个编辑按钮 -->
      <template #right-icon>
        <div class="van-group-style">
          <!-- <van-button
            type="primary"
            size="small"
            @click="editAction(item, 1)"
            :style="{ backgroundColor: getButtonColor(item.role_id, 1) }"
            >admin</van-button
          > -->
          <van-button
            type="primary"
            size="small"
            @click="editAction(item, 2)"
            :style="{ backgroundColor: getButtonColor(item.role_id, 2) }"
            >viewer</van-button
          >
          <van-button
            type="primary"
            size="small"
            @click="editAction(item, 3)"
            :style="{ backgroundColor: getButtonColor(item.role_id, 3) }"
            >prohibit</van-button
          >
        </div>
      </template>
    </van-cell>
  </van-cell-group>
</template>

<style scoped>
h1 {
  text-align: center;
  width: 100%;
  height: 40px;
}

.van-group-style {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  gap: 10px;
}
.bottom-group-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 50%;
  width: 50%;
  margin-left: 25%;
  background-color: #f9f9f9;
}

h6 {
  text-align: right;
  margin-right: 20px;
}

.userinfo-method {
  display: flex;
  flex-direction: row;
  margin-left: 80vw;
  height: 30px;
}

.login-out {
  display: flex;
  height: 30px;
  justify-content: center;
  align-items: center;
}
</style>
