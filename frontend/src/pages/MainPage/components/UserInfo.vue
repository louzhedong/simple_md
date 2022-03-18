<template>
  <div class="user-info">
    <div class="user-avatar">
      <img :src="avatar" alt="" class="avatar-img" />
    </div>
		<h3 class="user-nickname">Miracle</h3>
		<div class="social-icon">
			<a class="iconfont icon-github" href="https://github.com/louzhedong" target="__blank"></a>
			<i class="iconfont icon-weixin"></i>
		</div>
  </div>
</template>

<script>
import avatar from "@assets/cat.jpeg";

export default {
  data() {
    return {
      avatar,
    };
  },
};
</script>

<style lang="less">
.user-info {
	width: 150px;
  .user-avatar {
    width: 150px;
    height: 150px;
    border-radius: 10px;
    overflow: hidden;
    .avatar-img {
      width: 150px;
      height: 150px;
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
		}
	}
}
</style>

