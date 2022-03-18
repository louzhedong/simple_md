<template>
  <div class="main">
    <Header />
    <div class="list">
      <UserInfo />
    </div>
  </div>
</template>

<script>
import Header from "./components/Header.vue";
import UserInfo from './components/UserInfo.vue';
export default {
  components: {
    Header,
    UserInfo
  },
};
</script>

<style lang="less">
.list {
  padding: 3rem 10rem;
}
</style>