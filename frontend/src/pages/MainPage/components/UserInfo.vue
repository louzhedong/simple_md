<template>
  <div class="user-info">
    <div class="user-avatar">
      <img :src="avatar" alt="" class="avatar-img" />
    </div>
    <h3 class="user-nickname">Miracle</h3>
    <div class="social-icon">
      <a
        class="iconfont icon-github"
        href="https://github.com/louzhedong"
        target="__blank"
      ></a>
      <i
        class="iconfont icon-weixin"
        @mouseover="handleShowWechat"
        @mouseleave="handleHideWechat"
      >
        <img :src="wechat" alt="" class="wechat" v-if="showWechat" />
      </i>
    </div>
  </div>
</template>

<script>
import avatar from "@assets/cat.jpeg";
import wechat from "@assets/wechat.png";

export default {
  data() {
    return {
      avatar,
      wechat,
      showWechat: false,
    };
  },
  methods: {
    handleShowWechat() {
      this.showWechat = true;
    },
    handleHideWechat() {
      this.showWechat = false;
    },
  },
};
</script>

<style lang="less">
.user-info {
  width: 200px;
  .user-avatar {
    width: 200px;
    height: 200px;
    border-radius: 10px;
    overflow: hidden;
    .avatar-img {
      width: 200px;
      height: 200px;
    }
  }
  .user-nickname {
    font-size: 30px;
    margin-top: 20px;
    text-align: center;
  }
  .social-icon {
    .iconfont {
      font-size: 24px;
      margin: 4px;
      color: #777;
      position: relative;
    }
    .wechat {
      width: 200px;
      height: 200px;
      border-radius: 10px;
      position: absolute;
      bottom: 20px;
      left: 10px;
    }
  }
}
</style>

