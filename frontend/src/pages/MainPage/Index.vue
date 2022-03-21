<template>
  <div class="main">
    <Header />
    <div class="list">
      <UserInfo />
      <BriefList />
    </div>
  </div>
</template>

<script>
import Header from "./components/Header.vue";
import UserInfo from "./components/UserInfo.vue";
import BriefList from "./components/BriefList.vue";
export default {
  components: {
    Header,
    UserInfo,
    BriefList,
  },
};
</script>

<style lang="less">
.list {
  display: flex;
  padding-top: 50px;
  justify-content: center;
}
</style>